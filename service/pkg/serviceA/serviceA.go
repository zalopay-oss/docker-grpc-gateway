package serviceA

import (
	"context"
	service_proto "service/pkg/api"
	"time"

	"github.com/prometheus/common/log"
)

//ServiceA struct
type ServiceA struct {
}

//NewServiceA create service
func NewServiceA() service_proto.ServiceAServer {
	return &ServiceA{}
}

//getTimestamp get timestamp
func getTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

//Ping ping
func (service *ServiceA) Ping(ctx context.Context, msgPing *service_proto.MessagePing) (*service_proto.MessagePong, error) {
	log.Info(msgPing)
	return &service_proto.MessagePong{
		Timestamp:   getTimestamp(),
		ServiceName: "service A",
	}, nil
}
