程序中时常有发送邮件的需求。有异常情况了需要通知管理员和负责人，用户下单后可能需要通知订单信息，电商平台、中国移动和联通都有每月账单，这些都可以通过邮件来推送。还有我们平时收到的垃圾邮件大都也是通过这种方式发送的😭。

## 那么如何在 Go 语言发送邮件？本文我们介绍一下email库的使用。

先安装库，这个自不必说：
```
$ go get github.com/jordan-wright/email
```
我们需要额外一些工作。我们知道邮箱使用SMTP/POP3/IMAP等协议从邮件服务器上拉取邮件。邮件并不是直接发送到邮箱的，而是邮箱请求拉取的。

所以，我们需要配置SMTP/POP3/IMAP服务器。从头搭建固然可行，而且也有现成的开源库，但是比较麻烦。现在一般的邮箱服务商都开放了SMTP/POP3/IMAP服务器。

我这里拿 QQ 邮箱来举例，使用SMTP服务器。

首先，登录邮箱；


点开顶部的设置，选择POP3/SMTP/IMAP；

点击开启IMAP/SMTP服务，按照步骤开启即可，有个密码设置，记住这个密码，后面有用。

然后就可以编码了：

```
package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func main() {
	e := email.NewEmail()
	e.From = "dj <xx@qq.com>"
	e.Cc = []string{"xx@qq.com"}
	e.To = []string{"xx@qq.com"}
	e.Subject = "Go 每日一库"
	e.HTML = []byte(`
  	<ul>
		test
	</ul>
  	`)
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "xx@qq.com", "xx", "smtp.qq.com"))
	if err != nil {
		fmt.Println("failed to send email:", err)
	}
}
```
有时会被当作垃圾邮件，可能需要到垃圾箱查看

结果：

![image](https://github.com/foxliang/Blog/blob/master/images/go-%E9%82%AE%E4%BB%B6%E5%8F%91%E9%80%81%E7%BB%93%E6%9E%9C.png)
