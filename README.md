# trpc

[![License](https://img.shields.io/:license-apache%202-blue.svg)](https://opensource.org/licenses/Apache-2.0)

`trpc` is a distributed RPC framework, focus on service governance

## 功能

- 基于ProtoBuf生成客户端，服务端
- 基于TCP长连接， 底层传输协议可扩展
- 服务发现
- 负载均衡
- 可用性：限流，熔断，重试
- 插件化

## 快速开始

## 代码生成

### 下载protoc

[下载链接](https://developers.google.com/protocol-buffers/docs/downloads)

### 安装工具

```shell
go install .\cmd\protoc-gen-trpc\...
```

### 生成代码

```shell
protoc --trpc_out=plugins=trpc:. xxx.proto
```