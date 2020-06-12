package rpcServer

import (
	"context"
	"fmt"
)

type ApiController struct{
}

func (ctrl *ApiController)SayHello(ctx context.Context, in *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{
		Message: fmt.Sprintf("%s", in.Name),
	},nil
}

func (ctrl *ApiController)LotsOfReplies(in *HelloRequest, stream Hello_LotsOfRepliesServer) error {
	for i:=0; i<5; i++ {
		stream.Send(&HelloResponse{
			//in.Name Reply [第i次]
			Message: fmt.Sprintf("%s %s %d", in.Name, "Reply", i),
		})
	}
	return nil
}

