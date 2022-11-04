package tcp

import (
	"context"
	"crypto/tls"
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/outputs"
	"github.com/elastic/beats/v7/libbeat/outputs/codec"
	"github.com/elastic/beats/v7/libbeat/publisher"
	"github.com/elastic/elastic-agent-libs/config"
	"github.com/elastic/elastic-agent-libs/logp"
	"github.com/pkg/errors"
	"net"
	"time"
)

const (
	networkTCP = "tcp"
)

func init() {
	outputs.RegisterType("tcp", makeTcp)
}

type tcpOut struct {
	*logp.Logger
	connection net.Conn

	address   *net.TCPAddr
	sslEnable bool
	sslConfig *tls.Config

	lineDelimiter []byte
	codec         codec.Codec

	observer outputs.Observer
	index    string
}

var count int64

func makeTcp(
	_ outputs.IndexManager,
	beat beat.Info,
	observer outputs.Observer,
	cfg *config.C,
) (outputs.Group, error) {
	config := defaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return outputs.Fail(err)
	}
	logger := logp.NewLogger("tcp")
	// disable bulk support in publisher pipeline
	err := cfg.SetInt("bulk_max_size", -1, -1)
	if err != nil {
		logger.Info("cfg.SetInt failed with: %v", err)
	}
	clis := make([]outputs.NetworkClient, 0, config.Worker)
	for i := 0; i < config.Worker; i++ {
		enc, err := codec.CreateEncoder(beat, config.Codec)
		if err != nil {
			return outputs.Fail(err)
		}
		t, err := newTcpOut(logger, beat.Beat, config, observer, enc)
		if err != nil {
			return outputs.Fail(err)
		}
		clis = append(clis, outputs.WithBackoff(t, config.Backoff.Init, config.Backoff.Max))
	}
	go func() {
		for {
			select {
			case <-time.After(time.Second):
				logger.Infof("count: %v/s", count)
				count = 0
			}
		}
	}()
	return outputs.SuccessNet(true, -1, 0, clis)
}

func newTcpOut(logger *logp.Logger, index string, c tcpConfig, observer outputs.Observer, codec codec.Codec) (*tcpOut, error) {
	t := &tcpOut{
		Logger:        logger,
		sslEnable:     c.SSLEnable,
		lineDelimiter: []byte(c.LineDelimiter),
		observer:      observer,
		index:         index,
		codec:         codec,
	}
	addr, err := net.ResolveTCPAddr(networkTCP, net.JoinHostPort(c.Host, c.Port))
	if err != nil {
		return nil, errors.Wrap(err, "resolve tcp addr failed")
	}
	t.address = addr

	if c.SSLEnable {
		var cert tls.Certificate
		cert, err := tls.LoadX509KeyPair(c.SSLCertPath, c.SSLKeyPath)
		if err != nil {
			return nil, errors.Wrap(err, "load tls cert failed")
		}
		t.sslConfig = &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	}

	t.Info("new tcp output, address=%v", t.address)
	return t, nil
}
func (t *tcpOut) Connect() (err error) {
	//if t.sslEnable {
	//	t.connection, err = tls.Dial(networkTCP, t.address.String(), t.sslConfig)
	//} else {
	//	t.connection, err = net.DialTCP(networkTCP, nil, t.address)
	//}
	return err
}

func (t *tcpOut) Close() error {
	t.Info("tcp output connection %v close", t.address)
	//_ = t.connection.Close()
	//t.connection = nil
	return nil
}

func (t *tcpOut) Publish(ctx context.Context, batch publisher.Batch) error {

	events := batch.Events()

	t.observer.NewBatch(len(events))

	dropped := 0
	//buf := net.Buffers{}
	//for _ := range events {
	//	//serializedEvent, err := t.codec.Encode(t.index, &events[i].Content)
	//	//if err != nil {
	//	//	dropped++
	//	//	continue
	//	//}
	//	//buf = append(buf, append(serializedEvent, t.lineDelimiter...))
	//	count += int64(len(batch.Events()))
	//	//t.observer.WriteBytes(n)
	//}

	count += int64(len(batch.Events()))
	//n, err := buf.WriteTo(t.connection)
	//if err != nil {
	//	t.observer.WriteError(err)
	//	dropped = len(events)
	//}
	t.observer.Dropped(dropped)
	t.observer.Acked(len(events) - dropped)
	batch.ACK()
	return nil
}

func (t *tcpOut) String() string {
	return "tcp(" + t.address.String() + ")"
}
