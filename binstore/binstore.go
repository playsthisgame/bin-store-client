package binstore

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/playsthisgame/bin-store/types"
)

type Config struct {
	Host string
	Port uint16
}

type BinStoreClient struct {
	conn *net.TCPConn
}

func Connect(conf *Config) (*BinStoreClient, error) {
	addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)

	server, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		slog.Error("Error resolving server:", "error", err)
		return nil, err
	}

	conn, err := net.DialTCP("tcp", nil, server)
	if err != nil {
		slog.Error("Error dialing server:", "error", err)
		return nil, err
	}

	return &BinStoreClient{
		conn: conn,
	}, nil
}

func (c *BinStoreClient) Close() {
	c.conn.Close()
}

func (c *BinStoreClient) Write(key int64, data *[]byte) error {

	cmd := &types.TCPCommand{
		Command: 1,
		Data:    *data,
		Key:     key,
	}

	return sendCommand(c, cmd)
}

func (c *BinStoreClient) Read(key int64) (*[]byte, error) {
	cmd := &types.TCPCommand{
		Command: 2,
		Key:     key,
	}

	err := sendCommand(c, cmd)
	if err != nil {
		return nil, err
	}

	received := make([]byte, 1024) // TODO: this probably wont return everything right?
	_, err = c.conn.Read(received)
	if err != nil {
		slog.Error("Read data failed:", "error", err.Error())
		return nil, err
	}
	return &received, nil
}

func (c *BinStoreClient) Store(name string) error {
	cmd := &types.TCPCommand{
		Command: 3,
		Data:    []byte(name),
		Key:     0,
	}

	return sendCommand(c, cmd)
}

func (c *BinStoreClient) Load(name string) error {
	cmd := &types.TCPCommand{
		Command: 4,
		Data:    []byte(name),
		Key:     0,
	}

	return sendCommand(c, cmd)
}

func (c *BinStoreClient) Merge(name string) error {
	cmd := &types.TCPCommand{
		Command: 5,
		Data:    []byte(name),
		Key:     0,
	}

	return sendCommand(c, cmd)
}

func (c *BinStoreClient) Clear() error {
	cmd := &types.TCPCommand{
		Command: 6,
		Key:     0,
	}

	return sendCommand(c, cmd)
}

func sendCommand(c *BinStoreClient, cmd *types.TCPCommand) error {
	data, err := cmd.MarshalBinary()
	if err != nil {
		slog.Error("Error marshalling data:", "error", err)
		return err
	}

	_, err = c.conn.Write(data)
	if err != nil {
		slog.Error("Error writing to server", "error", err)
		return err
	}
	return nil
}
