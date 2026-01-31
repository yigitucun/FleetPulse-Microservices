package discovery

import (
	"time"

	"github.com/hudl/fargo"
)

func Register(eurekaAddr string) *fargo.Instance {
	fargoConn := fargo.NewConn(eurekaAddr)
	instance := &fargo.Instance{
		HostName: "localhost",
		App: "INGESTION-SERVICE",
		IPAddr: "127.0.0.1",
		Port: 8081,
		Status:           fargo.UP,
		DataCenterInfo:   fargo.DataCenterInfo{Name: fargo.MyOwn},
		LeaseInfo:        fargo.LeaseInfo{RenewalIntervalInSecs: 30},
	}
	go func() {
		for {
			err := fargoConn.HeartBeatInstance(instance)
			if err != nil {
				err := fargoConn.RegisterInstance(instance)
				if err != nil {
					return 
				}
			}
			time.Sleep(30 * time.Second)
		}
	}()
	return instance
}
