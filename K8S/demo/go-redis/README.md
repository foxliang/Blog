# 搭建一个api和一个redis服务

- 能够接收用户传的数据并且存在redis里面

- 数据需要持久化


部署

```
kubectl apply -f pv.yaml

kubectl apply -f pvc.yaml

kubectl apply -f redis.yaml
```

查看服务

```
$ kubectl get pods
NAME                                    READY   STATUS    RESTARTS   AGE
redis-fox-deployment-7b79f5755c-sg5rc   1/1     Running   0          5m24s

$ kubectl get svc 
NAME         TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
redis-fox    NodePort    10.98.96.220     <none>        6379:31379/TCP   5m26s

$ kubectl get pv 
NAME         CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM               STORAGECLASS    REASON   AGE
redis-pv     100Mi      RWO            Retain           Bound       default/redis-pvc   nfs                      5m28s

$ kubectl get pvc
NAME        STATUS   VOLUME     CAPACITY   ACCESS MODES   STORAGECLASS   AGE
redis-pvc   Bound    redis-pv   100Mi      RWO            nfs            5m30s

$ kubectl get configmap
NAME               DATA   AGE
redis.conf         1      102m

$ minikube service list
|----------------------|---------------------------|---------------------|---------------------------|
|      NAMESPACE       |           NAME            |     TARGET PORT     |            URL            |
|----------------------|---------------------------|---------------------|---------------------------|
| default              | redis-fox                 | redis-fox-port/6379 | http://192.168.79.2:31379 |
```

连接redis

```
$ redis-cli -h 192.168.79.2 -p 31379                            
192.168.79.2:31379> CONFIG GET maxmemory
1) "maxmemory"
2) "2097152"
192.168.79.2:31379> CONFIG GET maxmemory-policy
1) "maxmemory-policy"
2) "allkeys-lru"
192.168.79.2:31379> get test
"1"
192.168.79.2:31379> set test 2
OK
192.168.79.2:31379> get test
"2"
```
