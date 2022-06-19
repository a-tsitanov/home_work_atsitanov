package main

import (
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type MyTelnetClient struct {
	conn    net.Conn
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
}

func (tc *MyTelnetClient) Receive() error {
	_, err := io.Copy(tc.out, tc.conn)
	return err
}

func (tc *MyTelnetClient) Send() error {
	_, err := io.Copy(tc.conn, tc.in)
	return err
}

func (tc *MyTelnetClient) Close() error {
	err := tc.conn.Close()
	return err
}

func (tc *MyTelnetClient) Connect() error {
	var err error
	tc.conn, err = net.DialTimeout("tcp", tc.address, tc.timeout)
	return err
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &MyTelnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

// Place your code here.
// P.S. Author's solution takes no more than 50 lines.
