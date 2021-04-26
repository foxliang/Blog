# 部署kafka单机服务

1.快速部署
```
# kubectl create namespace kafka
# kubectl create -f 'https://strimzi.io/install/latest?namespace=kafka' -n kafka
# kubectl apply -f kafka-zk.yaml -n kafka
# kubectl wait kafka/my-cluster --for=condition=Ready --timeout=300s -n kafka
```

2.查看服务

```
$ kubectl get pods -n kafka
^[[ANAME                                         READY   STATUS    RESTARTS   AGE
kafka-producer                               1/1     Running   0          3d20h
my-cluster-entity-operator-98c779b75-ngqph   3/3     Running   0          3d20h
my-cluster-kafka-0                           1/1     Running   0          3d20h
my-cluster-zookeeper-0                       1/1     Running   0          3d20h
strimzi-cluster-operator-957688b5c-wvdd9     1/1     Running   0          3d20h

$ kubectl get svc -n kafka
NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                      AGE
my-cluster-kafka-0                    NodePort    10.100.17.184    <none>        9094:31359/TCP               3d20h
my-cluster-kafka-bootstrap            ClusterIP   10.101.231.131   <none>        9091/TCP,9092/TCP,9093/TCP   3d20h
my-cluster-kafka-brokers              ClusterIP   None             <none>        9091/TCP,9092/TCP,9093/TCP   3d20h
my-cluster-kafka-external-bootstrap   NodePort    10.104.128.99    <none>        9094:32422/TCP               3d20h
my-cluster-zookeeper-client           ClusterIP   10.101.34.148    <none>        2181/TCP                     3d20h
my-cluster-zookeeper-nodes            ClusterIP   None             <none>        2181/TCP,2888/TCP,3888/TCP   3d20h

$ minikube service list -n kafka
|-----------|-------------------------------------|-------------------|---------------------------|
| NAMESPACE |                NAME                 |    TARGET PORT    |            URL            |
|-----------|-------------------------------------|-------------------|---------------------------|
| kafka     | my-cluster-kafka-0                  | tcp-external/9094 | http://192.168.79.2:31359 |
| kafka     | my-cluster-kafka-bootstrap          | No node port      |
| kafka     | my-cluster-kafka-brokers            | No node port      |
| kafka     | my-cluster-kafka-external-bootstrap | tcp-external/9094 | http://192.168.79.2:32422 |
| kafka     | my-cluster-zookeeper-client         | No node port      |
| kafka     | my-cluster-zookeeper-nodes          | No node port      |
|-----------|-------------------------------------|-------------------|---------------------------|
```
3.创建生产者,消费者

#eg1
```
kubectl -n kafka run kafka-producer -ti --image=quay.io/strimzi/kafka:0.22.1-kafka-2.7.0 --rm=true --restart=Never -- bin/kafka-console-producer.sh --broker-list my-cluster-kafka-bootstrap:9092 --topic my-topic

If you don't see a command prompt, try pressing enter.
>1
>2
>3

kubectl -n kafka run kafka-consumer -ti --image=quay.io/strimzi/kafka:0.22.1-kafka-2.7.0 --rm=true --restart=Never -- bin/kafka-console-consumer.sh --bootstrap-server my-cluster-kafka-bootstrap:9092 --topic my-topic --from-beginning

If you don't see a command prompt, try pressing enter.
1
2
3

```

#eg2
```
$ kubectl -n kafka run kafka-producer1 -ti --image=quay.io/strimzi/kafka:0.22.1-kafka-2.7.0 --rm=true --restart=Never -- bin/kafka-console-producer.sh --broker-list 192.168.79.2:31359 --topic my-topic

If you don't see a command prompt, try pressing enter.

>>
>2
>54
>test

$ kubectl -n kafka run kafka-consumer -ti --image=quay.io/strimzi/kafka:0.22.1-kafka-2.7.0 --rm=true --restart=Never -- bin/kafka-console-consumer.sh --bootstrap-server 192.168.79.2:31359  --topic my-topic --from-beginning

If you don't see a command prompt, try pressing enter.

2
54
test

```
