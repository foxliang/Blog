Xdebug 在开发过程中可以帮我们查看具体的运行和步骤，以及每行代码执行的结果，在学习和解决代码问题的时候可以提供非常大的便利。PHPStorm 也可以进行 Xdebug 调试，VScode 也可以进行配置调试，且比 PHPStorm 的配置简单很多，不用每次去创建一个 Server，再创建一个 web page 服务。相比之下，VSCode 的界面好看，且简单方便，值得学习一下。

使用了一段时间，但是偶尔还是会出现一些问题，故而进行了整理总结。

一. 插件准备

二.进行配置
下载xdebug 扩展，根据phpinfo 中的信息去下载不同的版本

xdebug扩展





放到php/ext文件中



配置php.ini

[XDebug]
zend_extension="D:\xampp\php\ext\php_xdebug.dll"
xdebug.auto_trace=1
xdebug.collect_params=1
xdebug.collect_return=1
xdebug.trace_output_dir ="D:\xampp\htdocs\xdebug"
xdebug.profiler_output_dir ="D:\xampp\htdocs\xdebug"
xdebug.profiler_output_name = "cachegrind.out.%t.%p"
xdebug.remote_enable = 1
xdebug.remote_autostart = 1
xdebug.remote_handler = "dbgp"
xdebug.remote_host = "127.0.0.1"
# 设置端口号，默认是9000，此处因为本地环境端口冲突故设置为9002（在vscode配置中需要用到）
xdebug.remote_port = 9002
重新启动php，打开phpinfo，查看xdebug



3.查看 vscode 中 debug 页面 进行配置


 

还需要在settings.json 文件中配置php环境

    "php.validate.executablePath"      : "D:\\xampp\\php\\php.exe",
 

最后，在vscode里断点好后。按F5，等待请求，即可享受图形化的调试乐趣
