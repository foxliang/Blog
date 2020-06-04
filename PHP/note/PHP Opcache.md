# PCache 介绍
OPCache 是Zend官方出品的，开放自由的 opcode 缓存扩展，还具有代码优化功能，省去了每次加载和解析 PHP 脚本的开销。

PHP 5.5.0 及后续版本中已经绑定了 OPcache 扩展。

缓存两类内容:

- OPCode

- Interned String，如注释、变量名等

## OPCache 原理
OPCache缓存的机制主要是：将编译好的操作码放入共享内存，提供给其他进程访问。

这里就涉及到内存共享机制，另外所有内存资源操作都有锁的问题，我们一一解读。

#### 共享内存
UNIX/Linux 系统提供很多种进程间内存共享的方式：

System-V shm API: System V共享内存,

sysv shm是持久化的，除非被一个进程明确的删除，否则它始终存在于内存里，直到系统关机；

mmap API：

mmap映射的内存在不是持久化的，如果进程关闭，映射随即失效，除非事先已经映射到了一个文件上

内存映射机制mmap是POSIX标准的系统调用，有匿名映射和文件映射两种

mmap的一大优点是把文件映射到进程的地址空间

避免了数据从用户缓冲区到内核page cache缓冲区的复制过程；

当然还有一个优点就是不需要频繁的read/write系统调用


POSIX API： System V 的共享内存是过时的, POSIX共享内存提供了使用更简单、设计更合理的API.

Unix socket API

OPCache 使用了前三个共享内存机制，根据配置或者默认mmap 内存共享模式。

依据PHP字节码缓存的场景，OPCache的内存管理设计非常简单，快速读写，不释放内存，过期数据置为Wasted。

当Wasted内存大于设定值时，自动重启OPCache机制，清空并重新生成缓存。

#### 互斥锁

任何内存资源的操作，都涉及到锁的机制。

共享内存：一个单位时间内，只允许一个进程执行写操作，允许多个进程执行读操作；

写操作同时，不阻止读操作，以至于很少有锁死的情况。

这就引发另外一个问题：新代码、大流量场景，进程排队执行缓存opcode操作；重复写入，导致资源浪费。

### OPCache 缓存解读
OPCache 是官方的Opcode 缓存解决方案，在PHP5.5版本之后，已经打包到PHP源码中一起发布。

它将PHP编译产生的字节码以及数据缓存到共享内存中, 在每次请求，从缓存中直接读取编译后的opcode，进行执行。

通过节省脚本的编译过程，提高PHP的运行效率。

如果正在使用APC扩展，做同样的工作，现在强烈推荐OPCache来代替，尤其是PHP7中。

### OPCode 缓存
Opcache 会缓存OPCode以及如下内容：

- PHP脚本涉及到的函数

- PHP脚本中定义的Class

- PHP脚本文件路径

- PHP脚本OPArray

- PHP脚本自身结构/内容

### OPCache 更新策略
是缓存，都存在过期，以及更新策略等。

而OPCache的更新策略非常简单，到期数据置为Wasted，达到设定值，清空缓存，重建缓存。

这里需要注意：在高流量的场景下，重建缓存是一件非常耗费资源的事儿。

OPCache 在创建缓存时并不会阻止其他进程读取。

这会导致大量进程反复新建缓存。所以，不要设置OPCache过期时间

每次发布新代码时，都会出现反复新建缓存的情况。如何避免呢？

不要在高峰期发布代码，这是任何情况下都要遵守的规则

代码预热，比如使用脚本批量调PHP 访问URL，

或者使用OPCache 暴露的API 如opcache_compile_file() 进行编译缓存

### OPCache 的配置
##### 内存配置
opcache.preferred_memory_model="mmap" OPcache 首选的内存模块。如果留空，OPcache 会选择适用的模块， 通常情况下，自动选择就可以满足需求。可选值包括： mmap，shm, posix 以及 win32。

opcache.memory_consumption=64 OPcache 的共享内存大小，以兆字节为单位，默认64M

opcache.interned_strings_buffer=4 用来存储临时字符串的内存大小，以兆字节为单位，默认4M

opcache.max_wasted_percentage=5 浪费内存的上限，以百分比计。 如果达到此上限，那么 OPcache 将产生重新启动续发事件。默认5

#### 允许缓存的文件数量以及大小
opcache.max_accelerated_files=2000 OPcache 哈希表中可存储的脚本文件数量上限。 真实的取值是在质数集合 { 223, 463, 983, 1979, 3907, 7963, 16229, 32531, 65407, 130987 } 中找到的第一个大于等于设置值的质数。 设置值取值范围最小值是 200，最大值在 PHP 5.5.6 之前是 100000，PHP 5.5.6 及之后是 1000000。默认值2000

opcache.max_file_size=0 以字节为单位的缓存的文件大小上限。设置为 0 表示缓存全部文件。默认值0

#### 注释相关的缓存
opcache.load_commentsboolean 如果禁用，则即使文件中包含注释，也不会加载这些注释内容。 本选项可以和 opcache.save_comments 一起使用，以实现按需加载注释内容。

opcache.fast_shutdown boolean 如果启用，则会使用快速停止续发事件。 所谓快速停止续发事件是指依赖 Zend 引擎的内存管理模块 一次释放全部请求变量的内存，而不是依次释放每一个已分配的内存块。

#### 二级缓存的配置
opcache.file_cache 配置二级缓存目录并启用二级缓存。 启用二级缓存可以在 SHM 内存满了、服务器重启或者重置 SHM 的时候提高性能。 默认值为空字符串 ""，表示禁用基于文件的缓存。

opcache.file_cache_onlyboolean 启用或禁用在共享内存中的 opcode 缓存。

opcache.file_cache_consistency_checksboolean 当从文件缓存中加载脚本的时候，是否对文件的校验和进行验证。

opcache.file_cache_fallbackboolean 在 Windows 平台上，当一个进程无法附加到共享内存的时候， 使用基于文件的缓存，也即：opcache.file_cache_only=1。 需要显示的启用文件缓存。

### 配置信息
```
zend_extension=opcache.so

; Zend Optimizer + 的开关, 关闭时代码不再优化.
opcache.enable=1

; Determines if Zend OPCache is enabled for the CLI version of PHP
opcache.enable_cli=1


; Zend Optimizer + 共享内存的大小, 总共能够存储多少预编译的 PHP 代码(单位:MB)
; 推荐 128
opcache.memory_consumption=64

; Zend Optimizer + 暂存池中字符串的占内存总量.(单位:MB)
; 推荐 8
opcache.interned_strings_buffer=4


; 最大缓存的文件数目 200  到 100000 之间
; 推荐 4000
opcache.max_accelerated_files=2000

; 内存“浪费”达到此值对应的百分比,就会发起一个重启调度.
opcache.max_wasted_percentage=5

; 开启这条指令, Zend Optimizer + 会自动将当前工作目录的名字追加到脚本键上,
; 以此消除同名文件间的键值命名冲突.关闭这条指令会提升性能,
; 但是会对已存在的应用造成破坏.
opcache.use_cwd=0


; 开启文件时间戳验证 
opcache.validate_timestamps=1


; 2s检查一次文件更新 注意:0是一直检查不是关闭
; 推荐 60
opcache.revalidate_freq=2

; 允许或禁止在 include_path 中进行文件搜索的优化
;opcache.revalidate_path=0


; 是否保存文件/函数的注释   如果apigen、Doctrine、 ZF2、 PHPUnit需要文件注释
; 推荐 0
opcache.save_comments=1

; 是否加载文件/函数的注释
;opcache.load_comments=1


; 打开快速关闭, 打开这个在PHP Request Shutdown的时候会收内存的速度会提高
; 推荐 1
opcache.fast_shutdown=1

;允许覆盖文件存在（file_exists等）的优化特性。
;opcache.enable_file_override=0


; 定义启动多少个优化过程
;opcache.optimization_level=0xffffffff


; 启用此Hack可以暂时性的解决”can’t redeclare class”错误.
;opcache.inherited_hack=1

; 启用此Hack可以暂时性的解决”can’t redeclare class”错误.
;opcache.dups_fix=0

; 设置不缓存的黑名单
; 不缓存指定目录下cache_开头的PHP文件. /png/www/example.com/public_html/cache/cache_ 
;opcache.blacklist_filename=


; 通过文件大小屏除大文件的缓存.默认情况下所有的文件都会被缓存.
;opcache.max_file_size=0

; 每 N 次请求检查一次缓存校验.默认值0表示检查被禁用了.
; 由于计算校验值有损性能,这个指令应当紧紧在开发调试的时候开启.
;opcache.consistency_checks=0

; 从缓存不被访问后,等待多久后(单位为秒)调度重启
;opcache.force_restart_timeout=180

; 错误日志文件名.留空表示使用标准错误输出(stderr).
;opcache.error_log=


; 将错误信息写入到服务器(Apache等)日志
;opcache.log_verbosity_level=1

; 内存共享的首选后台.留空则是让系统选择.
;opcache.preferred_memory_model=

; 防止共享内存在脚本执行期间被意外写入, 仅用于内部调试.
;opcache.protect_memory=0
```
