package tcp

import (
	"fmt"
	"os"
	"time"

	"github.com/elastic/beats/v7/libbeat/outputs/codec"
	"github.com/pkg/errors"
)

type tcpConfig struct {
	Host string `config:"host"`
	Port string `config:"port"`

	SSLEnable   bool   `config:"ssl.enable"`
	SSLCertPath string `config:"ssl.cert_path"`
	SSLKeyPath  string `config:"ssl.key_path"`

	LineDelimiter string       `config:"line_delimiter"`
	Codec         codec.Config `config:"codec"`

	Backoff Backoff `config:"backoff"`

	Worker int `config:"worker"`
}

type Backoff struct {
	Init time.Duration `config:"init"`
	Max  time.Duration `config:"max"`
}

var defaultConfig = tcpConfig{
	Worker:        3,
	LineDelimiter: "\n",
	Backoff: Backoff{
		Init: 1 * time.Second,
		Max:  60 * time.Second,
	},
}

func (c *tcpConfig) Validate() error {
	if c.SSLEnable {
		if _, err := os.Stat(c.SSLCertPath); os.IsNotExist(err) {
			return errors.New(fmt.Sprintf("certificate %s not found", c.SSLCertPath))
		}
		if _, err := os.Stat(c.SSLKeyPath); os.IsNotExist(err) {
			return errors.New(fmt.Sprintf("key %s not found", c.SSLKeyPath))
		}
	}
	return nil
}
