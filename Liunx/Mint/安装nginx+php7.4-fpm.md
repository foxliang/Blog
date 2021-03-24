### 1.先禁用apache
```
sudo systemctl disable --now apache2
```
### 2.安装nginx,php7.4-fpm 及其他扩展

```
apt install nginx php7.4-fpm php7.4-dev  php-pear php7.4-mysql php7.4-curl php7.4-json php7.4-mbstring php7.4-xml php7.4-intl
```
之后可以用pecl安装php扩展

### 栗子:

```
pecl install yaf
```

安装成功后把 extension=yaf.so 加到php.ini中

访问localhost 可以看到nginx信息

配置nginx.conf 可访问php文件
