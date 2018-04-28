package influx

import (
	"errors"
)

// Config is influx structure config.
type Config struct {
	Addr        string `json:"addr"`
	PayloadSize int    `json:"payload_size"`
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	return c == rhs
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {

	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.New("namespace empty")
	}

	cAddr, ok := fconf["addr"]
	if !ok {
		return errors.New("missing key addr")
	}
	if c.Addr, ok = cAddr.(string); !ok {
		return errors.New("key addr invalid. must be string")
	}

	cPayloadSize, ok := fconf["payload_size"]
	if !ok {
		return errors.New("missing key payload_size")
	}
	cPayloadSizeFloat64, ok := cPayloadSize.(float64)
	if !ok {
		return errors.New("key payload_size invalid. must be int")
	}
	c.PayloadSize = int(cPayloadSizeFloat64)

	return nil
}
