package cluster

import "gopkg.in/nelsam/influxdb.v0/protocol"

// ResponseChannel is a processor for Responses as opposed to Series
// like `engine.Processor'
type ResponseChannel interface {
	Yield(r *protocol.Response) bool
	Name() string
}
