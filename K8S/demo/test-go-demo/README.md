# 任务
```
镜像名：foxliang/go:v
镜像内容：二进制文件，跑起来是个http服务，监听8080端口，接到请求会打印 This is version:{v} running in pod {hostname}，其中v是可变参数代表版本（字符串1.0这种），可以通过传参数进去
```

## 构建Docker镜像并打包上传
```
$ docker build  -t go:v1 .
 
$ docker tag go:v1 foxliang/go:v1
 
$ docker push foxliang/go:v1
```

## 部署服务
```
$ kubectl apply -f go.yaml
```

### 查询本地服务
```
$ kubectl get pods
NAME                             READY   STATUS    RESTARTS   AGE
go-deployment-8675897977-6r9rd   1/1     Running   0          3m25s
go-deployment-8675897977-c7kv2   1/1     Running   0          3m47s
go-deployment-8675897977-f7dn2   1/1     Running   0          3m47s

$ kubectl get service            
NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
go-service   NodePort    10.98.181.145   <none>        8080:31080/TCP   5m9s


$ minikube service list
|----------------------|---------------------------|--------------|---------------------------|
|      NAMESPACE       |           NAME            | TARGET PORT  |            URL            |
|----------------------|---------------------------|--------------|---------------------------|
| default              | go-service                | go-port/8080 | http://192.168.79.2:31080 |

```

#### 请求本地服务
```
for i in `seq 1 10`     
do
  curl http://192.168.79.2:31080
  echo " "
# echo $i
done

This is version:1.3 running in pod go-deployment-8675897977-6r9rd 
This is version:1.3 running in pod go-deployment-8675897977-c7kv2 
This is version:1.3 running in pod go-deployment-8675897977-f7dn2 
This is version:1.3 running in pod go-deployment-8675897977-f7dn2 
This is version:1.3 running in pod go-deployment-8675897977-c7kv2 
This is version:1.3 running in pod go-deployment-8675897977-c7kv2 
This is version:1.3 running in pod go-deployment-8675897977-f7dn2 
This is version:1.3 running in pod go-deployment-8675897977-f7dn2 
This is version:1.3 running in pod go-deployment-8675897977-6r9rd 
This is version:1.3 running in pod go-deployment-8675897977-6r9rd 
```

