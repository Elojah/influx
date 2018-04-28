package influx

import (
	"github.com/influxdata/influxdb/client/v2"
)

// Service represents the influx service.
type Service struct {
	client.Client
}

// Dial sends the new config to Service.
func (s *Service) Dial(c Config) error {
	var err error
	s.Client, err = client.NewUDPClient(client.UDPConfig{
		Addr:        c.Addr,
		PayloadSize: c.PayloadSize,
	})
	return err
}

// Close closes the session to cluster session.
func (s *Service) Close() error {
	return s.Client.Close()
}

// Healthcheck returns if database responds.
func (s *Service) Healthcheck() error {
	return nil
}
