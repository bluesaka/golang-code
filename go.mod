module go-code

go 1.15

require (
	github.com/bwmarrin/snowflake v0.3.0
	github.com/coreos/bbolt v1.3.5 // indirect
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/gojektech/heimdall/v6 v6.1.0
	github.com/golang/protobuf v1.4.2
	github.com/gomodule/redigo v1.8.4
	github.com/google/uuid v1.2.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.14.5 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/json-iterator/go v1.1.10
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/onsi/ginkgo v1.12.1 // indirect
	github.com/onsi/gomega v1.10.0 // indirect
	github.com/prometheus/common v0.10.0 // indirect
	github.com/prometheus/procfs v0.2.0 // indirect
	github.com/samuel/go-zookeeper v0.0.0-20201211165307-7117e9ea2414
	github.com/sirupsen/logrus v1.6.0 // indirect
	github.com/spf13/cast v1.3.1
	github.com/stretchr/testify v1.5.1
	github.com/tal-tech/go-zero v1.1.6
	github.com/tmc/grpc-websocket-proxy v0.0.0-20201229170055-e5319fda7802 // indirect
	github.com/valyala/fasthttp v1.23.0
	github.com/zieckey/etcdsync v0.0.0-20180810020013-cd5b26bc05a1
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/lint v0.0.0-20201208152925-83fdc39ff7b5 // indirect
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	golang.org/x/tools v0.1.0 // indirect
	google.golang.org/protobuf v1.26.0-rc.1 // indirect
)

replace github.com/coreos/bbolt v1.3.5 => go.etcd.io/bbolt v1.3.5

replace google.golang.org/grpc v1.36.0 => google.golang.org/grpc v1.29.1
