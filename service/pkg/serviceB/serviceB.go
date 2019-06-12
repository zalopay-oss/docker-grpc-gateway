package serviceB

import (
	"context"
	service_proto "service/pkg/api"
	"time"

	"github.com/prometheus/common/log"
)

//ServiceB struct
type ServiceB struct {
}

//NewServiceB create service
func NewServiceB() service_proto.ServiceBServer {
	return &ServiceB{}
}

//getTimestamp get timestamp
func getTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

//Ping ping
func (service *ServiceB) Ping(ctx context.Context, msgPing *service_proto.MessagePing) (*service_proto.MessagePong, error) {
	log.Info(msgPing)
	return &service_proto.MessagePong{
		Timestamp:   getTimestamp(),
		ServiceName: "service B",
	}, nil
}
