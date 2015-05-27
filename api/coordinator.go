package api

import (
	"gopkg.in/nelsam/influxdb.v0/cluster"
	cmn "gopkg.in/nelsam/influxdb.v0/common"
	"gopkg.in/nelsam/influxdb.v0/engine"
	"gopkg.in/nelsam/influxdb.v0/protocol"
)

// The following are the api that is accessed by any api

type Coordinator interface {
	// This is only used in the force compaction http endpoint
	ForceCompaction(cmn.User) error

	// Data related api
	RunQuery(cmn.User, string, string, engine.Processor) error
	WriteSeriesData(cmn.User, string, []*protocol.Series) error

	// Administration related api
	CreateDatabase(cmn.User, string) error
	ListDatabases(cmn.User) ([]*cluster.Database, error)
	DropDatabase(cmn.User, string) error
}
