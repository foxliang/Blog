## 构建Docker镜像并打包上传
```
$ docker build  -t go:v1 .
 
$ docker tag go:v1 foxliang/go:v1
 
$ docker push foxliang/go:v1
```

## 部署服务
```
kubectl apply -f go.yaml
```

