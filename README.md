# 概述

该项目是etcd + docker + grpc的demo项目

该项目构建的镜像同时生成了服务端和客户端代码，并且默认运行了grpc服务。你可以进入客户端容器调用该服务。

docker网络使用默认的bridge网络，通过默认的docker0网桥进行通信。

# 步骤

1. 构建镜像 docker build -t reverse .   (镜像名称为reverse)
2. 运行容器 
   1. 服务端 docker run -ti --rm -p 5300:5300 --name reverse-server reverse
   2. 客户端 docker run -ti --rm --name reverse-client reverse /bin/bash
3. 访问服务 ./cli 123456789
4. 观察结果 如果无误的话客户端会收到响应987654321

# 注意点

1. etcd镜像和容器在[这里](https://cloud.docker.com/u/breakinferno/repository/docker/breakinferno/etcd-goreman)
直接pull下来run 即可。注意需要指定CLIENT_ADDR环境变量。表示自己客户端成员的ip.比如 docker run -itd -e CLIENT_ADDR=127.0.0.1 -p 2379:2379 -p 2380:2380 -p 22379:22379 -p 22380:22380
 -p 32380:32380 -p 32379:32379 --name etcd etcd-goreman:v3.3.13。
2. 这里使用了golang v1.11以上的mod管理机制。可以降级为vendor机制，此时需要设置Dockerfile里面环境变量`ENV GOFLAGS=-mod=vendor`
3. 由于使用的是桥接网络，所以为了服务访问的稳定性，这里采取docker容器访问宿主机端口的方式进行容器间通行。而非容器间直接通信


