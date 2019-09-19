# Swoole：面向生产环境的 PHP 异步网络通信引擎
使 PHP 开发人员可以编写高性能的异步并发 TCP、UDP、Unix Socket、HTTP，WebSocket 服务。

Swoole 可以广泛应用于互联网、移动通信、企业软件、云计算、网络游戏、物联网（IOT）、车联网、智能家居等领域。 

使用 PHP + Swoole 作为网络通信框架，可以使企业 IT 研发团队的效率大大提升，更加专注于开发创新产品。

## 特性
Swoole 使用纯 C 语言编写，提供了 PHP 语言的异步多线程服务器，异步 TCP/UDP 网络客户端，异步 MySQL，异步 Redis，数据库连接池，AsyncTask，消息队列，毫秒定时器，异步文件读写，异步DNS查询。 Swoole内置了Http/WebSocket服务器端/客户端、Http2.0服务器端。

除了异步 IO 的支持之外，Swoole 为 PHP 多进程的模式设计了多个并发数据结构和IPC通信机制，可以大大简化多进程并发编程的工作。其中包括了并发原子计数器，并发 HashTable，Channel，Lock，进程间通信IPC等丰富的功能特性。

Swoole2.0 支持了类似 Go 语言的协程，可以使用完全同步的代码实现异步程序。PHP 代码无需额外增加任何关键词，底层自动进行协程调度，实现异步。


- 事件驱动的异步编程模式
- 异步TCP/UDP/HTTP/WebSocket/HTTP2协议的服务器端/客户端
- 支持IPv4/IPv6/UnixSocket/TCP/UDP
- 支持SSL/TLS隧道加密
- 支持并发百万TCP长连接
- 支持毫秒定时器
- 支持异步/同步/协程
- 支持CPU亲和性设置/守护进程

官网：https://www.swoole.com/
