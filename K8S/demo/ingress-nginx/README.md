# 目标:

- 实现一个服务，通过这个服务的某个端口可以访问内部不同的其他服务

### 这里使用helm安装ingress

```
helm repo add nginx-stable https://helm.nginx.com/stable
```
### 启动ingress

```
helm install gateway nginx-stable/nginx-ingress \
  --set controller.service.type=NodePort \
  --set controller.service.httpPort.nodePort=30080 \
  --set controller.service.httpsPort.nodePort=30443
  
$ helm list
NAME   	NAMESPACE	REVISION	UPDATED                                	STATUS  	CHART              	APP VERSION
gateway	default  	1       	2021-06-08 10:51:50.625881046 +0800 CST	deployed	nginx-ingress-0.9.3	1.11.3   
```

### 部署yaml文件

```
kubectl apply -f ingress.yaml
```

### 查询pod,svc

```
$ kubectl get pods
NAME                                     READY   STATUS    RESTARTS   AGE
gateway-nginx-ingress-6c6987b565-cgn2f   1/1     Running   0          3h37m
myapp-deploy-7f6c4894c9-6cbk5            1/1     Running   0          3h29m
myapp-deploy2-77b66545f9-xh978           1/1     Running   0          3h28m

$ kubectl get svc 
NAME                    TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                      AGE
gateway-nginx-ingress   NodePort    10.105.195.40    <none>        80:30080/TCP,443:30443/TCP   3h37m
ingress-nginx           NodePort    10.107.41.63     <none>        8080:30082/TCP               3h29m
ingress-nginx2          NodePort    10.110.105.100   <none>        8080:30081/TCP               3h29m

$ kubectl get ingress
NAME            CLASS    HOSTS              ADDRESS   PORTS   AGE
ingress-myapp   <none>   test.ingress.com             80      3h29m

$ minikube service list
|-------------|-----------------------|---------------|---------------------------|
|  NAMESPACE  |         NAME          |  TARGET PORT  |            URL            |
|-------------|-----------------------|---------------|---------------------------|
| default     | gateway-nginx-ingress | http/80       | http://192.168.49.2:30080 |
|             |                       | https/443     | http://192.168.49.2:30443 |
| default     | ingress-nginx         | go-port/8080  | http://192.168.49.2:30082 |
| default     | ingress-nginx2        | go-port2/8080 | http://192.168.49.2:30081 |
```

### 请求访问

需要先配置hosts minikubeip test.ingress.com

```
$ curl test.ingress.com:30080/1

/1 This is version:v1.1.0 running in pod myapp-deploy-7f6c4894c9-6cbk5root

$ curl test.ingress.com:30080/2

/2 This is version:v3.1.0 running in pod myapp-deploy2-77b66545f9-xh978root
```
