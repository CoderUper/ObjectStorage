# ObjectStorage   分布式对象存储项目开发

## 简介
项目主要分为apiServer端和dataServer端，apiServer向外提供restful接口，负责接受客户的连接请求。
dataServer负责存储数据。rabbitmq用于apiServer和dataServer端之间心跳、文件定位等。Elasticearch用于存储元数据。

## 主要特点
### 数据校验和去重
主要思想，计算文件的sha256哈希值，以哈希值代替文件名进行存储，元数据服务器存储<filename,hash>的映射。客户上传数据时会将哈希作为http头部数据一起上传，服务器接受时会计算哈希，并对比。同时
由于哈希的唯一性，重复数据只保留一份。
### 数据冗余和即时修复
用Reed Solomon纠删码进行冗余，每份数据分成若干份，每份存储在互不相同的结点上。读取的时候进行文件完整性检查和修复。
修复

### 断点续传
支持断点下载和上传。下载时http头部包含偏移量。上传为了支持哈希校验，在上传完成后才开始校验。

### gzip压缩
节省空间，数据服务器存储存储压缩数据，下载时解压之后再发送出去。

## 改进
加入Nginx进行负载均衡
