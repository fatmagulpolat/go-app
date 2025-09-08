package couchbase

import (
	"time"

	"github.com/couchbase/gocb/v2"
)

func ConnectCluster(host string, username string, password string, connBufferSize uint) *gocb.Cluster {
	cluster, err := gocb.Connect(host, gocb.ClusterOptions{
		Username: username,
		Password: password,
		TimeoutsConfig: gocb.TimeoutsConfig{
			KVTimeout: 20 * time.Second,
		},
		InternalConfig: gocb.InternalConfig{
			ConnectionBufferSize: connBufferSize,
		},
	})
	if err != nil {
		panic("Db connection failed")
	}

	return cluster
}
