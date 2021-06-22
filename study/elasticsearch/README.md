### ElasticSearch

#### 简介

ElasticSearch是一个基于Lucene的搜索服务器。它提供了一个分布式多用户能力的全文搜索引擎，基于RESTful web接口。

#### cluster、node、sharding分片，分布式集群

```
cluster health:  curl -XGET 'localhost:9200/_cat/health?v&pretty'
cluster nodes list:   curl -XGET 'localhost:9200/_cat/nodes?v'
index: curl -XGET 'localhost:9200/_cat/indices?v''
create index: curl -XPUT 'localhost:9200/index-name?pretty&pretty'
```

#### 端口
```
elasticsearch默认端口9200
kibana默认端口5601
```

#### CURD index documents 索引数据增删改查
```
#### 若user的index不存在，会自动创建，es7已经废弃了指定type，es8开始不支持，默认为_doc
#### es5.0开始，移除string类型，使用text+keyword替代，text用于全文搜索，keyword用于关键词搜索及聚合

新增/更新: curl -XPUT 'localhost:9200/user/_doc/1?pretty&pretty' -H 'Content-Type: application/json' -d'{"name": "John Doe"}' (必须指定id值，换name再执行就是更新)
新增: curl -XPOST 'localhost:9200/user/_doc/2?pretty&pretty' -d '{"name": "AAA"}'  (新增id为2的索引数据，id值可以不指定，elastic会随机生成一个id)
更新: curl -XPOST 'localhost:9200/user/_doc/2/_update?pretty&pretty' -d '{"doc": {"name": "BBB", "age": 20}}'
查看: curl -XGET 'localhost:9200/user/_doc/1?pretty'
删除: curl -XDELETE 'localhost:9200/user?pretty&pretty'
```

#### 批量bulk

```
curl -XPOST 'localhost:9200/customer/_doc/_bulk?pretty&pretty' -d'
{"index":{"_id":"1"}}
{"name": "CCC" }
{"index":{"_id":"2"}}
{"name": "DDD" }
'

curl -XPOST 'localhost:9200/user/_doc/_bulk?pretty&pretty' -d'
{"update":{"_id":"1"}}
{"doc": { "name": "DDD becomes new DDD" } }
{"delete":{"_id":"2"}}
'
```

#### 导入json文件

```
curl -XPOST 'localhost:9200/bank/_doc/_bulk?pretty&refresh' --data-binary "@accounts.json" //路径正确即可，如在/data/accounts.json, 则 "@/data/accounts.json"
```

#### 查询 match

```
GET index
GET index/_search
GET index/_mapping

curl -XGET 'localhost:9200/bank/_search?q=*&sort=account_name:asc&pretty'

#### match_all 匹配所有
curl -XGET 'localhost:9200/bank/_search?pretty' -H 'Content-Type: application/json' -d'
{
  "query": { "match_all": {} },
  "sort": [
    { "account_number": "asc" }
  ]
}
'

curl -XGET 'localhost:9200/bank/_search?pretty' -d'
{
	"query": {"match_all": {}},
	"from": 10,
	"size": 3,
	"sort": {"balance": {"order": "desc"}},
	"_source": ["account_number", "balance"]
}'

#### match 绝对匹配
curl -XGET 'localhost:9200/bank/_search?pretty' -d'
{
	"query": {"match": {"account_number": 20}}
}'

#### match_phrase 模糊匹配
curl -XGET 'localhost:9200/bank/_search?pretty' -d'
{
	"query": {"match_phrase": {"address": "mill"}}
}'
```

#### 查询 bool

```
#### bool must  所有都要满足（and）
curl -XGET 'localhost:9200/bank/_search?pretty' -d'
{
  "query": {
    "bool": {
      "must": [
        { "match": { "address": "mill" } },
        { "match": { "address": "lane" } }
      ]
    }
  }
}
'

#### bool should 满足其中之一 （or）
curl -XGET 'localhost:9200/bank/_search?pretty' -d'
{
  "query": {
    "bool": {
      "should": [
        { "match": { "address": "mill" } },
        { "match": { "address": "lane" } }
      ]
    }
  }
}
'

#### bool must_not 不满足其中任何一个
curl -XGET 'localhost:9200/bank/_search?pretty' -d'
{
  "query": {
    "bool": {
      "must_not": [
        { "match": { "address": "mill" } },
        { "match": { "address": "lane" } }
      ]
    }
  }
}
'
```

#### 查询过滤  fliter

```
curl -XGET 'localhost:9200/bank/_search?pretty' -d'
{
  "query": {
    "bool": {
      "must": { "match_all": {} },
      "filter": {
        "range": {
          "balance": {
            "gte": 20000,
            "lte": 30000
          }
        }
      }
    }
  }
}
'
```

#### 查询聚合 aggregation (like MySQL group by,  但聚合会返回聚合结果和所有的列表)

```
curl -XGET 'localhost:9200/bank/_search?pretty' -d'
{
    "aggs": {
        "group_by_lastname_adam":{
            "terms": {
                "field": "firstname"
           }
       }
   },
   "query": {
        "match": {
            "lastname": "Heath"
       }
   }
}
'

curl -XGET 'localhost:9200/bank/_search?pretty' -d'
{
    "aggs": {
        "all_state":{
            "terms": {"field": "state"},
            "aggs": {"avg_balance": {"avg": {"field": "balance"}}}
       }
   }
}
'
```

#### nested查询
```
#### 创建mapping
PUT nest_index
{
  "mappings": {
    "properties": {
      "userId": {
        "type": "long"
      },
      "name": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "orders": {
        "properties": {
          "status": {
            "type": "long"
          },
          "amount": {
            "type": "long"
          }
        }
      }
    }
  }
}

#### 添加数据
POST nest_index/_doc/1
{
  "userId": 1,
  "name": "adam",
  "orders":[{
    "status":1,
    "amount":100
  },
  {
    "status":2,
    "amount":200
  }]
}

POST nest_index/_bulk
{"index":{"_id":3}}
{"userId":3, "name":"mike", "orders":[{"status":4, "amount":411}]}
{"index":{"_id":4}}
{"userId":4, "name":"tracy", "orders":[{"status":4, "amount":458}]}

#### 查询
GET nest_index/_search
{
  "query": {
    "term": {
      "orders.status": {
        "value": "4"
      }
    }
  }
}
```

#### join查询
```
#### 创建mapping
PUT join_test
{
  "mappings": {
    "properties": {
      "parent":{
        "properties": {
          "id":{
            "type": "long"
          },
          "parentName":{
            "type":"text"
          }
        }
      },
      "child":{
        "properties": {
          "id": {
            "type":"long"
          },
          "childName": {
            "type": "text"
          },
          "parentId": {
            "type": "long"
          }
        }
      },
      
      "relation": {
        "type":"join",
        "relations": {
          "parent":["child"]
        }
      }
    }
  }
}

#### 添加父parent数据
POST join_test/_doc/1
{
  "parent.id": 1,
  "relation":"parent",
  "parent.parentName":"父亲1"
}

POST join_test/_doc/2
{
  "parent.id":2,
  "relation":"parent",
  "parent.parentName":"父亲2"
}

#### 添加子child数据
POST join_test/_doc/101?routing=1
{
  "child.id":101,
  "child.Name":"父亲1的儿子",
  "child.parentId": 1,
  "relation":{
    "name":"child",
    "parent":1
  }
}
POST join_test/_doc/102?routing=1
{
  "child.id":102,
  "child.Name":"父亲1的女儿",
  "child.parentId": 1,
  "relation":{
    "name":"child",
    "parent":1
  }
}
POST join_test/_doc/201?routing=2
{
  "child.id":201,
  "child.Name":"父亲2的儿子",
  "child.parentId": 2,
  "relation":{
    "name":"child",
    "parent":2
  }
}
POST join_test/_doc/202?routing=2
{
  "child.id":202,
  "child.Name":"父亲2的女儿",
  "child.parentId": 2,
  "relation":{
    "name":"child",
    "parent":2
  }
}
POST join_test/_doc/203?routing=2
{
  "child.id":203,
  "child.Name":"父亲1的儿子，但parentId=2",
  "child.parentId": 2,
  "relation":{
    "name":"child",
    "parent":1
  }
}

#### 以子查父
GET join_test/_search
{
  "query": {
    "has_child": {
      "type": "child",
      "query": {
        "term": {
          "child.parentId": {
            "value": "1"
          }
        }
      }
    }
  }
}

#### 带inner_hits 以子查父，查出父子级联信息
GET join_test/_search
{
  "query": {
    "has_child": {
      "type": "child",
      "query": {
        "term": {
          "child.parentId": {
            "value": "1"
          }
        }
      },
      "inner_hits": {}
    }
  }
}

#### 以父查子
GET join_test/_search
{
  "query": {
    "has_parent": {
      "parent_type": "parent",
      "query": {
        "term": {
          "parent.id": {
            "value": "1"
          }
        }
      }
    }
  }
}
```

#### analyzer分析器
```
ik分词器，https://github.com/medcl/elasticsearch-analysis-ik/releases
拼音分词器 https://github.com/medcl/elasticsearch-analysis-pinyin/releases
下载对应版本的包，解压到 {path}/libexec/plugins文件夹中，重启es即可
扩展本地词库:在plugins/config目录下新增 new1.dic和new2.dic，IKAnalyzer.cfg.xml编辑<entry key="ext_dict">new1.dic;new2.dic</entry>

GET _analyze
{
  "analyzer": "standard",
  "text": "我爱中华人民共和国"
}

GET _analyze
{
  "analyzer": "ik_max_word",
  "text": "我爱中华人民共和国"
}

GET _analyze
{
  "analyzer": "ik_smart",
  "text": "我爱中华人民共和国"
}

分词器主要在两中情况下使用：
1. 在插入文档时，text类型的字段会被分词插入倒排索引 (analyzer)
2. 在查询文档时，对要查询的text类型的输入做分词，再去倒排索引搜索 (search_analyzer)

#### 创建mapping时指定分词器
PUT an_index
{
  "mappings": { 
    "properties": {
      "content":{
        "type":"text",
        "analyzer": "ik_max_word",
        "search_analyzer": "ik_smart"
      }
    }
  }
}

#### 添加数据
POST an_index/_bulk
{"index":{"_id":1}}
{"content":"美国留给伊拉克的是个烂摊子吗"}
{"index":{"_id":2}}
{"content":"公安部：各地校车将享最高路权"}
{"index":{"_id":3}}
{"content":"中韩渔警冲突调查：韩警平均每天扣1艘中国渔船"}
{"index":{"_id":4}}
{"content":"中国驻洛杉矶领事馆遭亚裔男子枪击 嫌犯已自首"}

#### highlight高亮查询
GET an_index/_search
{
  "query": {
    "match": {
      "content": {
        "query": "中国"
      }
    }
  },
  "highlight": {
    "pre_tags" : ["<tag1>", "<tag2>"],
    "post_tags" : ["</tag1>", "</tag2>"],
    "fields": {
      "content": {}
    }
  }
}

#### wildcard通配符匹配
GET an_index/_search
{
  "query": {
    "wildcard": {
      "content": "中*"
    }
  },
  "highlight": {
    "pre_tags" : ["<tag1>", "<tag2>"],
    "post_tags" : ["</tag1>", "</tag2>"],
    "fields": {
      "content": {}
    }
  }
}

#### fuzzy regex prefix等查询
```


#### es head工具
```
es head包 或者 下载chrome插件
```

#### 多标签搜索
```
#### 添加数据
POST tag_test/_bulk
{ "index": { "_id": 1 }}
{ "tag_id" : [1,3] }
{ "index": { "_id": 2 }}
{ "tag_id" : [11,2] }
{ "index": { "_id": 3 }}
{ "tag_id" : [22,4] }

#### 查看mapping，tag_id.type=long
GET tag_test/_mapping

#### 查询，tag_id in (1,2)
GET tag_test/_search
{
  "query": {
    "terms": {
      "tag_id": [1,2]
    }
  }
}

#### 查询，tag_id同时包含1和2
GET tag_test/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "term": {
            "tag_id": {
              "value": "1"
            }
          }
        },
        {
          "term": {
            "tag_id": {
              "value": "2"
            }
          }
        }
      ]
    }
  }
}
```

#### 自定义逗号分词器搜索
```
#### 设置index
PUT tag_test2
{
  "settings": {
    "analysis": {
      "analyzer": {
        "my_comma": {
          "type":"pattern",
          "pattern": ","
        }
      }
    }
  },
  "mappings": {
    "properties": {
      "tag_id":{
        "type": "text",
        "analyzer": "my_comma",
        "search_analyzer": "my_comma",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      }
    }
  }
}

#### 添加数据
POST tag_test2/_bulk
{ "index": { "_id": 1 }}
{ "tag_id" : "1,2" }
{ "index": { "_id": 2 }}
{ "tag_id" : "11,2" }
{ "index": { "_id": 3 }}
{ "tag_id" : "22,4" }
{ "index": { "_id": 4 }}
{ "tag_id" : "1,22" }

#### 分析
GET tag_test2/_analyze
{
  "field": "tag_id",
  "text": "1,22",
  "analyzer": "my_comma"
}

#### 查询
GET tag_test2/_search
{
  "query": {
    "terms": {
      "tag_id": [1]
    }
  }
}

GET tag_test2/_search
{
  "query": {
    "match": {
      "tag_id": {
        "query": "1",
        "analyzer": "my_comma"
      }
    }
  }
}
```