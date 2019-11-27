package omp

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"net"
	"sync"
	"time"

	"go.uber.org/zap"
)

// Connector .
type Connector struct {
	conn   net.Conn
	logger *zap.Logger

	reqLock *sync.Mutex
}

type connectorOption func(*Connector) error

func withLogger(l *zap.Logger) connectorOption {
	return func(c *Connector) error {
		c.logger = l
		return nil
	}
}

// newConnector .
func newConnector(addr string, opts ...connectorOption) (*Connector, error) {
	tc := &tls.Config{
		InsecureSkipVerify: true,
	}
	dialer := &net.Dialer{
		Timeout: time.Second * 10,
	}
	conn, err := tls.DialWithDialer(dialer, "tcp", addr, tc)
	if err != nil {
		return nil, err
	}

	c := &Connector{
		conn:    conn,
		reqLock: &sync.Mutex{},
	}

	for _, opt := range opts {
		err = opt(c)
		if err != nil {
			return nil, err
		}
	}

	if c.logger == nil {
		logger, err := zap.NewDevelopment()
		if err != nil {
			return nil, err
		}
		c.logger = logger
	}

	return c, nil
}

func (c *Connector) doRequest(req requestIface, resp responseIface) error {
	c.reqLock.Lock()
	defer c.reqLock.Unlock()

	b, err := xml.MarshalIndent(req, "", "    ")
	if err != nil {
		return err
	}
	c.logger.Debug("send xml request", zap.String("content", string(b)))

	_, err = c.conn.Write(b)
	if err != nil {
		return err
	}
	buf := make([]byte, 0, 1024)
	tmpSize := 1024
	tempBuf := make([]byte, tmpSize)
	for {
		// 假设 OMP 是每次发送全部应答，而不是每次只发送部分数据
		// 否则，这里 n 小于 tempSize 就会提前退出
		n, err := c.conn.Read(tempBuf)
		if err != nil {
			return err
		}
		// FIXME: 未验证效率，应该有更好的方法
		buf = append(buf, tempBuf[:n]...)
		if n < tmpSize {
			break
		}
	}
	c.logger.Debug("receive xml response", zap.String("content", string(buf)))
	err = xml.Unmarshal(buf, resp)
	if err != nil {
		return err
	}
	if resp.GetStatus() >= 400 {
		return fmt.Errorf("response error: %v", resp.GetStatusText())
	}
	return nil
}

// Ping .
func (c *Connector) Ping() {
}

// Auth .
func (c *Connector) Auth(username, password string) error {
	req := requestAuth{}
	req.Credentials.Username = username
	req.Credentials.Password = password
	resp := &responseAuth{}
	err := c.doRequest(req, resp)
	if err != nil {
		return err
	}
	return nil
}

// GetVersion .
func (c *Connector) GetVersion() (string, error) {
	req := requestVersion{}
	resp := &responseVersion{}
	err := c.doRequest(req, resp)
	if err != nil {
		return "", err
	}
	return resp.Version, nil
}
