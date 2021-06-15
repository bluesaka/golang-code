## k8s

#### 相关定义
- Pod 容器组，运行的最小单位
- Replica Sets (rs) 副本集
- Deployment 无状态应用
- StatefulSet 有状态应用
- DaemonSet 守护进程集
- Job 任务
- CronJob 定时任务
- Service 服务
- Ingress 路由
- PersistentVolume 存储卷
- PersistentVolumeClaim 存储声明
- ConfigMap 配置项
- Secret 保密字典

#### 相关命令
```
kubectl -n xxx get pods --show-labels
kubectl -n xxx get rs
kubectl -n xxx get deployments
kubectl -n xxx get svc
kubectl -n xxx get ingress
kubectl -n xxx get pvc

kubectl -n xxx set image deployment/topic-rpc ropic-rpc=harbor.xxx.com/topic-go/topic-rpc:v1.0.1 

kubectl get namespace
kubectl create namespace xxx

kubectl apply -f xxx.yaml
kubectl -n xxx describe pods/pod_name
kubectl exec -it xxx -- /bin/bash
```

#### YAML在线编辑工具
```
https://www.bejson.com/validators/yaml_editor/
```

