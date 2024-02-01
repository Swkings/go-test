module mq-test

go 1.18

replace (
	gitlab.uisee.ai/cloud/sdk/go2sky-plugin => gitlab.uisee.ai/cloud/sdk/go2sky-plugin.git v1.2.17
	gitlab.uisee.ai/cloud/sdk/utc-prom-client => gitlab.uisee.ai/cloud/sdk/utc-prom-client.git v0.1.0
	gitlab.uisee.ai/cloud/sdk/utclogger => gitlab.uisee.ai/cloud/sdk/utclogger.git v0.2.4
	gitlab.uisee.ai/gkl10385/utc-utils => gitlab.uisee.ai/gkl10385/utc-utils.git v0.50.2
)

require (
	gitlab.uisee.ai/gkl10385/scheduler-general-mq-protocol v0.16.1
	gitlab.uisee.ai/gkl10385/utc-utils v0.50.2
)

require (
	bou.ke/monkey v1.0.2 // indirect
	github.com/SkyAPM/go2sky v1.3.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rabbitmq/amqp091-go v1.5.0 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/streadway/amqp v1.1.0 // indirect
	gitlab.uisee.ai/cloud/sdk/utclogger v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20211013025323-ce878158c4d4 // indirect
	google.golang.org/grpc v1.41.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gorm.io/gorm v1.21.10 // indirect
	skywalking.apache.org/repo/goapi v0.0.0-20210628073857-a95ba03d3c7a // indirect
)
