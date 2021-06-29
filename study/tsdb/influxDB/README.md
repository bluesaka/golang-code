### InfluxDB

#### 简介：
InfluxDB是使用Golang编写的时序数据库(Time Series Database、TSDB)，时序数据库就是一串按时间维度索引的数据。
时序数据库特点：
- 持续高并发的写入、无更新
- 数据压缩存储
- 低查询延时


#### 与MySQL的对比：
InfluxDB | MySQL
--- | ---
Database | Database
Measurement | Table |
Point | Row

Point：
- timestamp 唯一主键，自动生成
- tag 各种有索引的属性，例如地区等
- field 各种记录值，不带索引，例如温度等

Series
> Series是一些数据的集合，在同一database中，retention policy、measurement、tag sets 完全相同的数据属于一个series，
> 同一个series的数据会按照时间顺序存储在一起


#### 常用InfluxQL
```
-- 查看所有的数据库
show databases;

-- 使用特定的数据库
use database_name;

-- 查看所有的measurement
show measurements;

-- 查询10条数据
select * from measurement_name limit 10;

-- 数据中的时间字段默认显示的是一个纳秒时间戳，改成可读格式
precision rfc3339; -- 之后再查询，时间就是rfc3339标准格式

-- 或可以在连接数据库的时候，直接带该参数
influx -precision rfc3339

-- 查看一个measurement中所有的tag key 
show tag keys

-- 查看一个measurement中所有的field key 
show field keys

-- 查看一个measurement中所有的保存策略(可以有多个，一个标识为default)
show retention policies;
```

#### 其他
```
我本地mac安装的是1.8.3版本，注意与2.0版本的区别
https://docs.influxdata.com/influxdb/v1.8/administration/config/#using-the-configuration-file

# 启动服务
influxd -config /usr/local/etc/influxdb.conf

# 查看配置
influxd config
influxd config -config /usr/local/etc/influxdb.conf | grep flux-enabled

# 修改配置
如 flux.enabeld = true

# 启动客户端
influx -precision rfc3339

服务端口 localhost:8086
```
