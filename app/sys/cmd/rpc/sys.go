package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/service"
	"google.golang.org/grpc/reflection"

	"aiops/app/sys/cmd/rpc/internal/config"
	"aiops/app/sys/cmd/rpc/internal/server"
	"aiops/app/sys/cmd/rpc/internal/svc"
	"aiops/app/sys/cmd/rpc/pb"
	"aiops/common/interceptor/rpcserver"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/sys.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewSysServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterSysServer(grpcServer, srv)
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	//rpc log
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)

	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
