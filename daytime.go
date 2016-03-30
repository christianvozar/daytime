// Copyright © 2015-2016 Christian R. Vozar
// MIT Licensed.

// Daytime server. See RFC 867
// https://tools.ietf.org/html/rfc867

package daytime

import (
	"net"
	"time"
)

// Server defines parameters for running a Daytime server.
type Server struct {
	Address string // TCP address to listen on, RFC 867 default of ":13" if empty
}

// ListenAndServe listens on the TCP network address srv.Address and handles
// requests on incoming connections.
func (srv *Server) ListenAndServe() error {
	addr := srv.Address
	if addr == "" {
		addr = ":13"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer ln.Close()

	for {
		// Wait for connection.
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		// Handle connection in new goroutine.
		// Loop returns to accepting so multiple connections are served concurrently.
		go func(c net.Conn) {
			// Write out current date and time in RFC850 format.
			t := time.Now().Format(time.RFC850)
			c.Write([]byte(t))
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}

// ListenAndServe listens on the TCP network address addr and serves Daytime
// protocol requests in a new goroutine.
func ListenAndServe(addr string) error {
	server := &Server{Address: addr}
	return server.ListenAndServe()
}
