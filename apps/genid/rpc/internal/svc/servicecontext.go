package svc

import (
	"github.com/ac-dcz/lightRW/apps/genid/rpc/internal/config"
	"github.com/sony/sonyflake"
	"time"
)

type ServiceContext struct {
	Config config.Config
	Flake  *sonyflake.Sonyflake
}

func NewServiceContext(c config.Config) *ServiceContext {
	startTime, err := time.Parse(time.DateTime, c.FlakeConf.StartTime)
	if err != nil {
		panic("start time parse err")
	}
	return &ServiceContext{
		Config: c,
		Flake: sonyflake.NewSonyflake(sonyflake.Settings{
			StartTime: startTime,
			MachineID: func() (uint16, error) {
				return c.FlakeConf.MachineID, nil
			},
			CheckMachineID: func(u uint16) bool {
				return true
			},
		}),
	}
}
