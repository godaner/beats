package tcp

import (
	"bufio"
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
)

const (
	networkTCP = "tcp4"
)

func init() {
	outputs.RegisterType("tcp", MakeTcp)
}

type tcpOut struct {
	*logp.Logger
	connection net.Conn
	bw         *bufio.Writer
	buf        net.Buffers

	address      *net.TCPAddr
	writevEnable bool
	sslEnable    bool
	sslConfig    *tls.Config

	lineDelimiter []byte
	codec         codec.Codec

	observer outputs.Observer
	index    string
}

func MakeTcp(
	_ outputs.IndexManager,
	beat beat.Info,
	observer outputs.Observer,
	cfg *config.C,
) (outputs.Group, error) {
	config := defaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return outputs.Fail(err)
	}
	logger := logp.NewLogger("beats-output-tcp")
	// disable bulk support in publisher pipeline
	err := cfg.SetInt("bulk_max_size", -1, -1)
	if err != nil {
		logger.Info("cfg.SetInt failed with: %v", err)
	}

	enc, err := codec.CreateEncoder(beat, config.Codec)
	if err != nil {
		return outputs.Fail(err)
	}
	t, err := newTcpOut(logger, beat.Beat, config, observer, enc)
	if err != nil {
		return outputs.Fail(err)
	}
	return outputs.Success(-1, 0, outputs.WithBackoff(t, config.Backoff.Init, config.Backoff.Max))
}

func newTcpOut(logger *logp.Logger, index string, c Config, observer outputs.Observer, codec codec.Codec) (*tcpOut, error) {
	t := &tcpOut{
		Logger:        logger,
		writevEnable:  c.WritevEnable,
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

	if t.writevEnable {
		t.buf = make([][]byte, c.BufferSize)
	} else {
		t.bw = bufio.NewWriterSize(t.connection, c.BufferSize)
	}

	t.Info("new tcp output, address=%v", t.address)
	return t, nil
}
func (t *tcpOut) Connect() (err error) {
	if t.sslEnable {
		t.connection, err = tls.Dial(networkTCP, t.address.String(), t.sslConfig)
	} else {
		t.connection, err = net.DialTCP(networkTCP, nil, t.address)
	}
	if err != nil {
		return err
	}
	if !t.writevEnable {
		t.bw.Reset(t.connection)
	}
	return nil
}

func (t *tcpOut) Close() error {
	t.Info("tcp output connection %v close", t.address)
	_ = t.connection.Close()
	t.connection = nil
	return nil
}

func (t *tcpOut) Publish(ctx context.Context, batch publisher.Batch) error {
	if t.writevEnable {
		return t.publishWritev(batch)
	}
	return t.publish(batch)
}

func (t *tcpOut) publish(batch publisher.Batch) (err error) {
	events := batch.Events()
	t.observer.NewBatch(len(events))

	bulkSize := 0
	dropped := 0
	for i := range events {
		serializedEvent, err := t.codec.Encode(t.index, &events[i].Content)
		if err != nil {
			dropped++
			continue
		}
		_, err = t.bw.Write(serializedEvent)
		if err != nil {
			t.observer.WriteError(err)
			dropped++
			continue
		}
		_, err = t.bw.Write(t.lineDelimiter)
		if err != nil {
			t.observer.WriteError(err)
			dropped++
			continue
		}
		bulkSize += len(serializedEvent) + 1
	}
	err = t.bw.Flush()
	if err != nil {
		t.observer.WriteError(err)
		dropped = len(events)
	}

	t.observer.WriteBytes(bulkSize)
	t.observer.Dropped(dropped)
	t.observer.Acked(len(events) - dropped)
	batch.ACK()
	return err
}

func (t *tcpOut) publishWritev(batch publisher.Batch) (err error) {
	events := batch.Events()
	t.observer.NewBatch(len(events))

	dropped := 0
	t.buf = t.buf[:0]
	for i := range events {
		serializedEvent, err := t.codec.Encode(t.index, &events[i].Content)
		if err != nil {
			dropped++
			continue
		}
		t.buf = append(t.buf, append(serializedEvent, t.lineDelimiter...))
	}
	var n int64
	n, err = t.buf.WriteTo(t.connection)
	if err != nil {
		t.observer.WriteError(err)
		dropped = len(events)
	}

	t.observer.WriteBytes(int(n))
	t.observer.Dropped(dropped)
	t.observer.Acked(len(events) - dropped)
	batch.ACK()
	return err
}

func (t *tcpOut) String() string {
	return "tcp(" + t.address.String() + ")"
}
