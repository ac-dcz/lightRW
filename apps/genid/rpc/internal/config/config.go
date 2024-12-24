package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	FlakeConf struct {
		StartTime string
		MachineID uint16
	}
}
