## 安装 nfs server
```
$ sudo apt-get update
 
$ sudo apt-get install -y nfs-kernel-server
```
### 2.创建目录,配置 IP 共享目录绑定
```
$ vim /etc/exports
 
# 新增
 
/data/nfs *(rw,sync,no_root_squash,no_subtree_check)
 
 
配置说明：
 
/data/nfs：是共享的数据目录
*：表示任何人都有权限连接，当然也可以是一个网段，一个 IP，也可以是域名
rw：读写的权限
sync：表示文件同时写入硬盘和内存
no_root_squash：当登录 NFS 主机使用共享目录的使用者是 root 时，其权限将被转换成为匿名使用者，通常它的 UID 与 GID，都会变成 nobody 身份
no_subtree_check:不做子目录检查
```
### 3.重启并查看服务
```
$ systemctl restart nfs-server
 
$ systemctl status nfs-server
 
● nfs-server.service - NFS server and services
     Loaded: loaded (/lib/systemd/system/nfs-server.service; enabled; vendor preset: enabled)
    Drop-In: /run/systemd/generator/nfs-server.service.d
             └─order-with-mounts.conf
     Active: active (exited) since Tue 2021-04-27 03:48:09 UTC; 11min ago
    Process: 737365 ExecStartPre=/usr/sbin/exportfs -r (code=exited, status=0/SUCCESS)
    Process: 737366 ExecStart=/usr/sbin/rpc.nfsd $RPCNFSDARGS (code=exited, status=0/SUCCESS)
   Main PID: 737366 (code=exited, status=0/SUCCESS)
 
Apr 27 03:48:08 minikube systemd[1]: Starting NFS server and services...
Apr 27 03:48:09 minikube systemd[1]: Finished NFS server and services.
```
### 4.验证
```
$ showmount -e  # 本机验证
 
Export list for minikube:
/data/nfs *
 
 
$ showmount -e 192.168.79.2 # 客户端验证 这里服务端与客户端是一个
 
 
Export list for 192.168.79.2:
/data/nfs *
```
