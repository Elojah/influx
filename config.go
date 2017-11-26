package couchbase

import (
	"errors"
)

// Config is couchbase structure config.
type Config struct {
	URL      string `json:"url"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Bucket   string `json:"bucket"`
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	return c.URL == rhs.URL &&
		c.Port == rhs.Port &&
		c.User == rhs.User &&
		c.Password == rhs.Password
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.New("namespace empty")
	}
	cURL, ok := fconf["url"]
	if !ok {
		return errors.New("missing key url")
	}
	if c.URL, ok = cURL.(string); !ok {
		return errors.New("key url invalid. must be string")
	}
	cPort, ok := fconf["port"]
	if !ok {
		return errors.New("missing key port")
	}
	if c.Port, ok = cPort.(string); !ok {
		return errors.New("key port invalid. must be string")
	}
	cUser, ok := fconf["user"]
	if !ok {
		return errors.New("missing key user")
	}
	if c.User, ok = cUser.(string); !ok {
		return errors.New("key user invalid. must be string")
	}
	cPassword, ok := fconf["password"]
	if !ok {
		return errors.New("missing key password")
	}
	if c.Password, ok = cPassword.(string); !ok {
		return errors.New("key password invalid. must be string")
	}
	cBucket, ok := fconf["bucket"]
	if !ok {
		return errors.New("missing key bucket")
	}
	if c.Bucket, ok = cBucket.(string); !ok {
		return errors.New("key bucket invalid. must be string")
	}
	return nil
}
