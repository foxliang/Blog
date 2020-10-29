防火墙是一组规则。当数据包进出受保护的网络区域时，进出内容（特别是关于其来源、目标和使用的协议等信息）会根据防火墙规则进行检测，以确定是否允许其通过。下面是一个简单的例子:

![images](https://github.com/foxliang/Blog/blob/master/images/Linux%E6%80%9D%E7%BB%B4%E8%84%91%E5%9B%BE/%E9%98%B2%E7%81%AB%E5%A2%99.jpg)

- iptables 是 Linux 机器上管理防火墙规则的工具。

-，firewalld 也是 Linux 机器上管理防火墙规则的工具。

你有什么问题吗？如果我告诉你还有另外一种工具，叫做 nftables，这会不会糟蹋你的美好一天呢？

好吧，我承认整件事确实有点好笑，所以让我来解释一下。这一切都从 Netfilter 开始，它在 Linux 内核模块级别控制访问网络栈。几十年来，管理 Netfilter 钩子的主要命令行工具是 iptables 规则集。

因为调用这些规则所需的语法看起来有点晦涩难懂，所以各种用户友好的实现方式，如 ufw 和 firewalld 被引入，作为更高级别的 Netfilter 解释器。然而，ufw 和 firewalld 主要是为解决单独的计算机所面临的各种问题而设计的。构建全方面的网络解决方案通常需要 iptables，或者从 2014 年起，它的替代品 nftables (nft 命令行工具)。

iptables 没有消失，仍然被广泛使用着。事实上，在未来的许多年里，作为一名管理员，你应该会使用 iptables 来保护的网络。但是 nftables 通过操作经典的 Netfilter 工具集带来了一些重要的崭新的功能。

从现在开始，我将通过示例展示 firewalld 和 iptables 如何解决简单的连接问题。

## 使用 firewalld 配置 HTTP 访问

正如你能从它的名字中猜到的，firewalld 是 systemd 家族的一部分。firewalld 可以安装在 Debian/Ubuntu 机器上，不过，它默认安装在 RedHat 和 CentOS 上。如果您的计算机上运行着像 Apache 这样的 web 服务器，您可以通过浏览服务器的 web 根目录来确认防火墙是否正在工作。如果网站不可访问，那么 firewalld 正在工作。

你可以使用 firewall-cmd 工具从命令行管理 firewalld 设置。添加 –state 参数将返回当前防火墙的状态:
```
# firewall-cmd --state
running
```
默认情况下，firewalld 处于运行状态，并拒绝所有传入流量，但有几个例外，如 SSH。这意味着你的网站不会有太多的访问者，这无疑会为你节省大量的数据传输成本。然而，这不是你对 web 服务器的要求，你希望打开 HTTP 和 HTTPS 端口，按照惯例，这两个端口分别被指定为 80 和 443。firewalld 提供了两种方法来实现这个功能。一个是通过 –add-port 参数，该参数直接引用端口号及其将使用的网络协议（在本例中为TCP）。 另外一个是通过 –permanent 参数，它告诉 firewalld 在每次服务器启动时加载此规则：
```
# firewall-cmd --permanent --add-port=80/tcp
# firewall-cmd --permanent --add-port=443/tcp
```
–reload 参数将这些规则应用于当前会话：
```
# firewall-cmd --reload
```
查看当前防火墙上的设置，运行 –list-services：

```
# firewall-cmd --list-services
dhcpv6-client http https ssh
```
假设您已经如前所述添加了浏览器访问，那么 HTTP、HTTPS 和 SSH 端口现在都应该是和 dhcpv6-client 一样开放的 —— 它允许 Linux 从本地 DHCP 服务器请求 IPv6 IP 地址。

## 使用 iptables 配置锁定的客户信息亭

我相信你已经看到了信息亭——它们是放在机场、图书馆和商务场所的盒子里的平板电脑、触摸屏和 ATM 类电脑，邀请顾客和路人浏览内容。大多数信息亭的问题是，你通常不希望用户像在自己家一样，把他们当成自己的设备。它们通常不是用来浏览、观看 YouTube 视频或对五角大楼发起拒绝服务攻击的。因此，为了确保它们没有被滥用，你需要锁定它们。

一种方法是应用某种信息亭模式，无论是通过巧妙使用 Linux 显示管理器还是控制在浏览器级别。但是为了确保你已经堵塞了所有的漏洞，你可能还想通过防火墙添加一些硬性的网络控制。在下一节中，我将讲解如何使用iptables 来完成。

关于使用 iptables，有两件重要的事情需要记住：你给出的规则的顺序非常关键；iptables 规则本身在重新启动后将无法保持。我会一次一个地在解释这些。

信息亭项目

为了说明这一切，让我们想象一下，我们为一家名为 BigMart 的大型连锁商店工作。它们已经存在了几十年；事实上，我们想象中的祖父母可能是在那里购物并长大的。但是如今，BigMart 公司总部的人可能只是在数着亚马逊将他们永远赶下去的时间。

尽管如此，BigMart 的 IT 部门正在尽他们最大努力提供解决方案，他们向你发放了一些具有 WiFi 功能信息亭设备，你在整个商店的战略位置使用这些设备。其想法是，登录到 BigMart.com 产品页面，允许查找商品特征、过道位置和库存水平。信息亭还允许进入 bigmart-data.com，那里储存着许多图像和视频媒体信息。

除此之外，您还需要允许下载软件包更新。最后，您还希望只允许从本地工作站访问 SSH，并阻止其他人登录。
