### Canal
> 阿里巴巴 MySQL binlog 增量订阅&消费组件
> https://github.com/alibaba/canal

#### MySQL Canal Golang Client
```
github.com/go-mysql-org/go-mysql
```

#### MySQL Canal to ElasticSearch
```
https://github.com/go-mysql-org/go-mysql-elasticsearch
git clone https://github.com/siddontang/go-mysql-elasticsearch
cd go-mysql-elasticsearch && make
vi etc/river.toml
启动服务：./bin/go-mysql-elasticsearch -config=etc/river.toml
```

#### Binlog
```
# 查看binlog是否开启   
show VARIABLES like '%log_bin%'

# 查看my.cnf配置   
mysql --help | grep my.cnf

# 在相关文件夹新建my.cnf （如/etc/mysql/my.cnf）

[mysqld]
log-bin = mysql-bin #开启binlog
binlog-format = ROW #选择row模式
server_id = 1 #配置mysql replication需要定义，不能和canal的slaveId重复

# 重启MySQL后，查看binlog日志状态
show master status

# 刷新之后会新建一个新的Binlog日志
flush logs

# 清空日志文件
reset master

# 查看binlog内容
show binlog events in 'mysql-bin.000002' limit 0,100

# 修改mysql数据文件夹权限
sudo chmod -R a+rwx /usr/local/mysql/data

# 终端查看内容
mysqlbinlog  /usr/local/mysql/data/mysql-bin.000002 | more

## 授权Canal账号权限
CREATE USER canal IDENTIFIED BY 'canal';  
GRANT SELECT, REPLICATION SLAVE, REPLICATION CLIENT ON *.* TO 'canal'@'%';
-- GRANT ALL PRIVILEGES ON *.* TO 'canal'@'%' ;
FLUSH PRIVILEGES;
```