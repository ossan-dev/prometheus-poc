package todo

import (
	"fmt"
	"net"
	"time"
)

var graphiteNetConn net.Conn

func setGraphiteConn() (err error) {
	graphiteHost := "localhost"
	graphitePort := "2003"
	graphiteNetConn, err = net.Dial("tcp", net.JoinHostPort(graphiteHost, graphitePort))
	return
}

func writeMetric(name string, value float64) (err error) {
	_, err = fmt.Fprintf(graphiteNetConn, "%s %f %d\n", name, value, time.Now().Unix())
	return
}
