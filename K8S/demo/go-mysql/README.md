# 搭建一个api和一个mysql服务

- 能够接收用户传的数据并且存在mysql里面

- 数据需要持久化


### 1.部署

```
kubectl apply -f mysql-config.yaml
kubectl apply -f mysql-pv.yaml
kubectl apply -f mysql-pvc.yaml
kubectl apply -f mysql.yaml
```

### 2.查看服务

```
$ kubectl get configmap
NAME               DATA   AGE
mysql-config       1      76m

$ kubectl get pv       
NAME         CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM               STORAGECLASS    REASON   AGE
mysql-pv     100Mi      RWO            Retain           Bound       default/mysql-pvc   nfs                      76m

$ kubectl get pvc
NAME        STATUS   VOLUME     CAPACITY   ACCESS MODES   STORAGECLASS   AGE
mysql-pvc   Bound    mysql-pv   100Mi      RWO            nfs            76m

$ kubectl get pods
NAME                                    READY   STATUS    RESTARTS   AGE
mysql-fox-deployment-7676755855-q4ghq   1/1     Running   0          6m59s

$ kubectl get svc 
NAME         TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
mysql-fox    NodePort    10.110.211.206   <none>        3306:31306/TCP   7m2s

$ minikube service list
|----------------------|---------------------------|---------------------|---------------------------|
|      NAMESPACE       |           NAME            |     TARGET PORT     |            URL            |
|----------------------|---------------------------|---------------------|---------------------------|
| default              | mysql-fox                 | mysql-fox-port/3306 | http://192.168.79.2:31306 |
```
### 3.连接mysql

```
# mysql -h192.168.79.2 -P31306 -uroot -p
123456

```

### 4.测试容器,节点停止再重启看数据是否存在

