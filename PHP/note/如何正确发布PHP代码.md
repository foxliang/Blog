几乎每一个 PHP 程序员都发布过代码，可能是通过 ftp 或者 rsync 同步的，也可能是通过 svn 或者 git 更新的。一个活跃的项目可能每天都要发布若干次代码，但是现实却是很少有人注意其中的细节，实际上这里面有好多坑，很可能你就在坑中却浑然不知。


一个正确实现的发布系统至少应该支持原子发布。如果说每一个版本都表示一个独立的状态的话，那么在发布期间，任何一次请求只能在单一状态下被执行。如此称之为支持原子发布；反之如果在发布期间，一次请求跨越不同的状态，那么就不能称之为原子发布。我们不妨举个例子来说明一下：假设一次请求需要 include 两个 PHP 文件，分别是 a.php 和 b.php，当 include a.php 完成后，发布代码，接着 include b.php，如果处理不当的话，那么就可能会导致旧版本的 a.php 和新版本的 b.php 同时存在于同一个请求之中，换句话说就是没有实现原子发布。

开源世界里有很多不错的发布代码工具，比如 ruby 社区的 capistrano，其流程大致就是发布代码到一个全新的目录，然后再软链接到真正的发布目录。
```
.
├── current -> releases/v1
└── releases
    ├── v1
    │   ├── foo.php
    │   └── bar.php
    └── v2
        ├── foo.php
        └── bar.php 
```
不过鉴于 PHP 本身的特殊性，如果只是简单套用上面的流程，那么将很难实现真正的原子发布。要理清个中缘由，还需要了解一下 PHP 中的两个 Cache 的概念：

- opcode cache
- realpath cache

先聊聊 opcode cache，基本就是 apc 或者 zend opcode，关于它的作用，大家都已经很熟悉，不必多言，需要注意的是 apc 的 bug 很多，比如开启了 apc.enable_cli 配置后就会有很多灵异问题，所以说 opcode cache 还是尽可能使用 zend opcache 吧，如果需要缓存数据，可以用 apcu。此外 apc 和 zend opcode 对缓存键的选择有所差异：apc 选择的是文件的 inode，zend opcode 选择的是文件的 path。

再聊聊 realpath cache，它的作用是缓冲获取文件信息的 IO 操作，大多数时候它对我们而言是透明的，以至于很多人都不知道它的存在，需要注意的是 realpath cache 是进程级别的，也就是说，每一个 php-fpm 进程都有自己独立的 realpath cache。

假设在发布代码期间，opcode cache 或者 realpath cache 里的数据出现过期，那么就会出现一部分缓存是旧文件，一部分缓存是新文件的非原子发布的情况，为了避免出现这种情况，我们应该保证缓存过期时间足够长，最好是除非我们手动刷新，否则永远不过期，对应到配置上就是：关闭 apc.stat、opcache.validate_timestamps 配置，设置足够大的 realpath_cache_size、realpath_cache_ttl 配置，必要的监控总是有好处的。

相关的技术细节特别琐碎，建议大家仔细阅读如下资料：

[realpath_cache](http://jpauli.github.io/2014/06/30/realpath-cache.html)
[PHP’s OPCache extension review](http://jpauli.github.io/2015/03/05/opcache.html)
[Atomic deploys at Etsy](https://codeascraft.com/2013/07/01/atomic-deploys-at-etsy/)
[Cache invalidation for scripts in symlinked folders](https://github.com/zendtech/ZendOptimizerPlus/issues/126)

在采用软链接发布代码的时候，通常遇到的第一个问题多半是新代码不生效！即便调用了 apc_clear_cache 或者 opcache_reset 方法也无效，重启 php-fpm 自然是能够解决问题，不过对脚本语言来说重启太重了！难道除了重启就没有别的办法了么？

事实上之所以会出现这样的问题，主要是因为 opcode cache 是通过 realpath cache 获取文件信息，即便软链接已经指向了新位置，但是如果 realpath cache 里还保存着旧数据的话，opcode cache 依然无法知道新代码的存在，缺省情况下，realpath_cache_ttl 缓存有效期是两分钟，这意味着发布代码后，可能要两分钟才能生效。为了让发布尽快生效，需要以进程为单位清除 realpath cache：
```
<?php

$key = 'php.pid_' . getmypid();

if (($rev = apc_fetch($key)) != DEPLOY_VERSION) {
    if($rev < DEPLOY_VERSION) {
        apc_store($key, DEPLOY_VERSION);
    }
    
    clearstatcache(true);
}

?>
```
如此在 apc 环境下基本就能工作了，但是在 zend opcode 环境下还可能有问题。因为在缺省情况下 opcache.revalidate_path 是关闭的，此时会缓存未解析的符号链接的值，这会导致即便软链接指向修改了，也无法生效，所以在使用 zend opcode 的时候，如果使用了软链接，视情况可能需要把 opcache.revalidate_path 激活。

详细介绍参考：PHP’s OPCache extension review。

BTW：如果需要手动重置 opcode cache，需要注意的是因为它是基于 SAPI 的概念，所以不能直接在命令行下调用 apc_clear_cache 或者 opcache_reset 方法来重置缓存，当然办法总是有的，那就是使用 CacheTool 在命令行下模拟 fastcgi 请求。

分析到这里，我们不妨反思一下：在 PHP 中原子发布之所以是一个棘手的问题，归根结底是因为软链接和缓存之间的的矛盾。不管是 opcode cache 还是 realpath cache，都是 PHP 固有的缓存特性，基于客观需要无法绕开，如此说来是否有办法绕开软链接，使其成为马奇诺防线呢？答案是 nginx 的 $realpath_root：

fastcgi_param SCRIPT_FILENAME $realpath_root$fastcgi_script_name;
fastcgi_param DOCUMENT_ROOT $realpath_root;
有了 $realpath_root，即便 DOCUMENT_ROOT 目录中含有软链接，nginx 也会把软链接指向的真正的路径发给 PHP，也就是说，对 PHP 而言，软链接已经不存在了！不过作为代价，每一次请求，nginx 都要通过相对昂贵的 IO 操作获取 $realpath_root 的值，通过 strace 命令我们能监控这一过程，下图从 current 到 foo 的过程：


realpath

在本例中，压测发现使用 $realpath_root 后，性能下降了大约 5% 左右，不过明眼人一下就能发现，虽然 $realpath_root 导致了 lstat 和 readlink 操作，但是 lstat 操作的次数是和目录深度成正比的，也就是说目录越深，执行的 lstat 次数越多，性能下降也就越大。如果能够降低发布目录的深度，那么可以预计还能降低一些性能损耗。

结尾介绍一下 Deployer，它是 PHP 中做得比较好的工具，有很多特色，比如支持并行发布，具体演示如下图，左边是串行，右边是并行，使用「vvv」能得到更详细信息：


deploy

不过 Deployer 在原子发布上有一点瑕疵，具体见 release/symlink 代码：
```
<?php

// deploy:release
run("cd {{deploy_path}} && if [ -h release ]; then rm release; fi");
run("ln -s $releasePath {{deploy_path}}/release");
// deploy:symlink
run("cd {{deploy_path}} && ln -sfn {{release_path}} current");
run("cd {{deploy_path}} && rm release");

?>
```
在 release 的时候，它是先删除再创建，是一个两步的非原子操作，在 symlink 的时候，看上去「ln -sfn」是单步原子操作，实际上也是错误的：
```
shell> strace ln -sfn releases/foo current
symlink("releases/foo", "current")      = -1 EEXIST (File exists)
unlink("current")                       = 0
symlink("releases/foo", "current")      = 0
通过 strace 我们能清晰的看到，虽然表面上使用「ln -sfn」是一步操作，但是内部依然是按照先删除再创建的逻辑执行的，实际上这里应该搭配使用「ln & mv」：

shell> ln -sfn releases/foo current.tmp
shell> mv -fT current.tmp current
```
先通过 ln 创建一个临时的软链接，再通过 mv 实现原子操作，此时如果使用 strace 监控，会发现 mv 的「T」选项实际上仅仅执行了一个 rename 操作，所以是原子的。

BTW：在使用「ln -sfn」前后，如果使用 stat 查看新旧文件的 inode 的话，可能会发现它们拥有一样的 inode 值，看上去和我们的结论相悖，其实不然，实际上只是复用删除值而已（如果想验证，注意 Linux 会复用，Mac 不会复用）。

据说一千个人的心中就有一千个哈姆雷特，不过我希望所有的 PHP 程序员在发布 PHP 代码的时候都能采用一种方法，那就是本文介绍的方法，正确的方法。
