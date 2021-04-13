
- namespace配置文件 ns.yaml

- elasticsearch配置文件 elastic.yaml

- 配置文件：kibana.yaml

- 这里设置fluentd的基于角色的访问控制配置文件，fluentd-rbac.yaml

- 配置fluentd-daemonset.yaml



## 依次部署服务

```
kubectl apply -f ns.yaml

kubectl apply -f elastic.yaml

kubectl apply -f kibana.yaml

kubectl apply -f fluentd-rbac.yaml

kubectl apply -f fluentd-daemonset.yaml

```

## 查看服务

```
kubectl get all -n kube-system

NAME                                   READY   STATUS    RESTARTS   AGE
pod/coredns-74ff55c5b-zcrjq            1/1     Running   0          26d
pod/etcd-minikube                      1/1     Running   0          26d
pod/fluentd-g76q4                      1/1     Running   0          77m
pod/kube-apiserver-minikube            1/1     Running   0          26d
pod/kube-controller-manager-minikube   1/1     Running   0          26d
pod/kube-proxy-ksc8m                   1/1     Running   0          26d
pod/kube-scheduler-minikube            1/1     Running   0          26d
pod/storage-provisioner                1/1     Running   0          26d

NAME               TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)                  AGE
service/kube-dns   ClusterIP   10.96.0.10   <none>        53/UDP,53/TCP,9153/TCP   26d

NAME                        DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR            AGE
daemonset.apps/fluentd      1         1         1       1            1           <none>                   77m
daemonset.apps/kube-proxy   1         1         1       1            1           kubernetes.io/os=linux   26d

NAME                      READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/coredns   1/1     1            1           26d

NAME                                DESIRED   CURRENT   READY   AGE
replicaset.apps/coredns-74ff55c5b   1         1         1       26d
```

#### es和kibana 都可以通过下面的url进行访问
```
minikube service list

|----------------------|---------------------------|------------------|---------------------------|
|      NAMESPACE       |           NAME            |   TARGET PORT    |            URL            |
|----------------------|---------------------------|------------------|---------------------------|
| default              | go-log-service            | go-log-port/8089 | http://192.168.79.2:31089 |
| default              | kubernetes                | No node port     |
| kube-system          | kube-dns                  | No node port     |
| kubernetes-dashboard | dashboard-metrics-scraper | No node port     |
| kubernetes-dashboard | kubernetes-dashboard      | No node port     |
| logging              | elasticsearch             |             9200 | http://192.168.79.2:31200 |
| logging              | kibana                    |             5601 | http://192.168.79.2:31601 |
|----------------------|---------------------------|------------------|---------------------------|
```

