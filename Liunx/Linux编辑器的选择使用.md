在执行 crontab -e的时候 发现编辑器不对 用的是nano 

立刻

```
sudo select-editor
```

然后终端返回

```
Select an editor. To change later, run 'select-editor'.
1. /bin/ed
2. /bin/nano <---- easiest
3. /usr/bin/vim.basic
4. /usr/bin/vim.tiny

Choose 1-4 [2]: 
```

然后选择4 就可以指定用vim编辑器了
