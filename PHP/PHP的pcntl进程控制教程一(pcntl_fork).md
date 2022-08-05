# pcntl 简介
PHP的进程控制支持实现了Unix方式的进程创建, 程序执行, 信号处理以及进程的中断。 进程控制不能被应用在Web服务器环境，当其被用于Web服务环境时可能会带来意外的结果。

这份文档用于阐述每个进程控制函数的通常用法。关于Unix进程控制的更多信息建议您查阅 系统文档中关于fork（2）,waitpid（2），signal（2）等的部分或更全面的参考资料比如 《Unix环境高级编程》（作者：W. Richard Stevens，Addison-Wesley出版）。

PCNTL现在使用了ticks作为信号处理的回调机制，ticks在速度上远远超过了之前的处理机制。 这个变化与“用户ticks”遵循了相同的语义。您可以使用declare() 语句在程序中指定允许发生回调的位置。这使得我们对异步事件处理的开销最小化。在编译PHP时 启用pcntl将始终承担这种开销，不论您的脚本中是否真正使用了pcntl。

有一个调整是PHP 4.3.0之前的所有pcntl脚本要使其工作，要么在期望允许回调的（代码）部分使用 declare() ，要么使用declare()新的全局语法 使其在整个脚本范围有效。

Note: 此扩展在 Windows 平台上不可用。

官方文档

pcntl官方文档(https://www.php.net/manual/zh/book.pcntl.php)


代码

```
$count = 10000;
$total_limit = 5000;
// 总数量
$process_count = ceil($count / $total_limit);
$childs = [];
for ($i = 0; $i < $process_count; $i++) {
    $pid = pcntl_fork();
    if ($pid == -1) {
        die("could not fork\n");
    } else if ($pid) {
        var_dump($i, $pid);
        var_dump(microtime(true));
        $childs[] = $pid;
    } else {
        var_dump($i, 11111);
        var_dump(microtime(true));
        // 最后一个线程，计算需要处理的总记录数
        if ($i == ($process_count - 1)) {
            $total_limit = $count % $total_limit;
        }
        $log_prefix = '[线程号: ' . $i . '] [current: %d/' . $total_limit . '] ';
        var_dump($log_prefix);
        exit(0);
    }
}

while (count($childs) > 0) {
    foreach ($childs as $key => $pid) {
        $res = pcntl_waitpid($pid, $status, WNOHANG);

        //-1代表error, 大于0代表子进程已退出,返回的是子进程的pid,非阻塞时0代表没取到退出子进程
        if ($res == -1 || $res > 0)
            unset($childs[$key]);
    }
    sleep(1);
}
exit(0);
```

#### fork首先会执行父进程逻辑再执行子进程逻辑
