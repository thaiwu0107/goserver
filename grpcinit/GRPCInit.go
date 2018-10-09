package grpcinit

import (
	Mysim "git.com.ggttoo44/src/service/mysim"
	grpc "google.golang.org/grpc"
)

// RegisterAllServer 在這裡註冊所有的GRPC Server
func RegisterAllServer(s *grpc.Server) {
	Mysim.RegisterMysimServer(s, &Mysim.Server{})
}
