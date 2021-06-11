#!/bin/bash

#设置网络名
network_name=etcd_network

#创建网络
docker network create --driver bridge --subnet=10.3.36.0/16 --gateway=10.3.1.1 ${network_name}

#设置结点名
node1=etcd_node1
node1_ip=10.3.36.1

node2=etcd_node2
node2_ip=10.3.36.2

node3=etcd_node3
node3_ip=10.3.36.3

#设置集群口令
cluster_token=etcd_cluster


#创建节点1
docker run -d --name ${node1} \
	--network ${network_name} \
	--publish 12379:2379 \
	--publish 12380:2380 \
	--ip ${node1_ip} \
	--env ALLOW_NONE_AUTHENTICATION=yes \
	--env ETCD_NAME=${node1} \
	--env ETCD_ADVERTISE_CLIENT_URLS=http://${node1_ip}:2379 \
	--env ETCD_INITIAL_ADVERTISE_PEER_URLS=http://${node1_ip}:2380 \
	--env ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:2379 \
	--env ETCD_LISTEN_PEER_URLS=http://0.0.0.0:2380 \
	--env ETCD_INITIAL_CLUSTER_TOKEN=${cluster_token} \
	--env ETCD_INITIAL_CLUSTER=${node1}=http://${node1_ip}:2380,${node2}=http://${node2_ip}:2380,${node3}=http://${node3_ip}:2380 \
	--env ETCD_INITIAL_CLUSTER_STATE=new \
	bitnami/etcd:latest

#创建节点2
docker run -d --name ${node2} \
	--network ${network_name} \
	--publish 22379:2379 \
	--publish 22380:2380 \
	--ip ${node2_ip} \
	--env ALLOW_NONE_AUTHENTICATION=yes \
	--env ETCD_NAME=${node2} \
	--env ETCD_ADVERTISE_CLIENT_URLS=http://${node2_ip}:2379 \
	--env ETCD_INITIAL_ADVERTISE_PEER_URLS=http://${node2_ip}:2380 \
	--env ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:2379 \
	--env ETCD_LISTEN_PEER_URLS=http://0.0.0.0:2380 \
	--env ETCD_INITIAL_CLUSTER_TOKEN=${cluster_token} \
	--env ETCD_INITIAL_CLUSTER=${node1}=http://${node1_ip}:2380,${node2}=http://${node2_ip}:2380,${node3}=http://${node3_ip}:2380 \
	--env ETCD_INITIAL_CLUSTER_STATE=new \
	bitnami/etcd:latest

#创建节点3
docker run -d --name ${node3} \
	--network ${network_name} \
	--publish 32379:2379 \
	--publish 32380:2380 \
	--ip ${node3_ip} \
	--env ALLOW_NONE_AUTHENTICATION=yes \
	--env ETCD_NAME=${node3} \
	--env ETCD_ADVERTISE_CLIENT_URLS=http://${node3_ip}:2379 \
	--env ETCD_INITIAL_ADVERTISE_PEER_URLS=http://${node3_ip}:2380 \
	--env ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:2379 \
	--env ETCD_LISTEN_PEER_URLS=http://0.0.0.0:2380 \
	--env ETCD_INITIAL_CLUSTER_TOKEN=${cluster_token} \
	--env ETCD_INITIAL_CLUSTER=${node1}=http://${node1_ip}:2380,${node2}=http://${node2_ip}:2380,${node3}=http://${node3_ip}:2380 \
	--env ETCD_INITIAL_CLUSTER_STATE=new \
	bitnami/etcd:latest