package couchbase

import (
	gocb "gopkg.in/couchbase/gocb.v1"
)

// Service represents the couchbase service.
type Service struct {
	cluster *gocb.Cluster
	bucket  *gocb.Bucket
}

// Dial sends the new config to Service.
func (s *Service) Dial(c Config) error {
	var err error

	if s.cluster, err = gocb.Connect(c.URL); err != nil {
		return err
	}
	if s.bucket, err = s.cluster.OpenBucket(c.Bucket, ""); err != nil {
		return err
	}
	return nil
}

// Healthcheck returns if database responds.
func (s *Service) Healthcheck() error {
	q := gocb.NewN1qlQuery(`SELECT CLOCK_STR()`)
	_, err := s.bucket.ExecuteN1qlQuery(q, nil)
	return err
}
