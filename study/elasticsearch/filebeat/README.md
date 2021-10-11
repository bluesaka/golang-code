## Filebeat

### mac安装
```
brew tap elastic/tap
brew install filebeat
```

### 配置
```
配置文件目录 /usr/local/etc/filebeat
日志路径目录 /usr/local/var/log/filebeat
文件状态(如采集的offset等) /usr/local/var/lib/filebeat/registry/filebeat/log.json
```

### filebeat.yml
```
filebeat.inputs:
  - type: log
    enabled: true
    paths:
      - /data/logs/test-*.log

output:
  elasticsearch:
    enabled: true
    hosts: ["localhost:9200"]
    index: "test-%{+yyyy.MM.dd}"

setup:
  ilm:
    enabled: false
  template:
    enabled: true
    name: "filebeat"
    pattern: "filebeat-*"

processors:
  - decode_json_fields:
      fields: ["message"]
      target: ""
  - drop_fields:
      fields: [ "input", "host", "agent", "ecs", "log", "message"]
      ignore_missing: true
```

### 启动
```
brew services start filebeat
或者
filebeat -e -c /usr/local/etc/filebeat/filebeat.yml
```

### Docker部署
```
FROM ubuntu:18.04

WORKDIR /usr/share/filebeat

COPY ./filebeat-7.12.0-linux-x86_64.tar.gz /usr/share
ADD ./docker-entrypoint.sh /usr/bin/

RUN cd /usr/share && \
    tar -xzf filebeat-7.12.0-linux-x86_64.tar.gz -C /usr/share/filebeat --strip-components=1 && \
    rm -f filebeat-7.12.0-linux-x86_64.tar.gz && \
    chmod +x /usr/share/filebeat && \
    chmod +x /usr/bin/docker-entrypoint.sh \ && 
    mkdir -p /etc/filebeat

ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["/usr/share/filebeat/filebeat", "-e", "-c", "/etc/filebeat/filebeat.yml"]
```

### K8s
```
K8s使用ConfiMap配置项，挂载filebeat.yml配置
```