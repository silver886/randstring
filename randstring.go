package randstring

import (
	"errors"
	"math/rand"
	"time"
)

// RandString is a source of random strings
type RandString struct {
	src *rand.Rand

	bucket       string
	bucketLength uint64
	bucketBit    uint
	bucketMask   uint64
	bucketMax    uint
}

// Config is the config of RandString
type Config struct {
	Source *rand.Source
	Bucket string
}

// New create a random string generator
func New(cfg *Config) (*RandString, error) {
	if cfg == nil {
		return nil, errors.New("Config is nil")
	}

	randString := new(RandString)

	if err := randString.SetBucket(cfg.Bucket); err != nil {
		return nil, err
	}

	if cfg.Source != nil {
		randString.src = rand.New(*cfg.Source)
	} else {
		randString.src = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	return randString, nil
}
