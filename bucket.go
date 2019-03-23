package randstring

import (
	"errors"
	"math"
)

// SetBucket set the bucket and calculate bucket related values
func (r *RandString) SetBucket(bucket string) error {
	if len(bucket) == 0 {
		return errors.New("Bucket is empty")
	}

	bucketBit := uint(math.Ceil(math.Log2(float64(len(bucket)))))
	if bucketBit > 64 {
		return errors.New("Bucket exceed maximum length")
	}

	r.bucket = bucket
	r.bucketLength = uint64(len(bucket))
	r.bucketBit = bucketBit
	r.bucketMask = 1<<r.bucketBit - 1
	r.bucketMax = 64 / r.bucketBit

	return nil
}

// AddBucket add additional bucket to current bucket and calculate bucket related values
func (r *RandString) AddBucket(bucket string) error {
	return r.SetBucket(r.bucket + bucket)
}

// AddNumberBucket add numbers to current bucket and calculate bucket related values
func (r *RandString) AddNumberBucket() error {
	return r.SetBucket(r.bucket + "0123456789")
}

// AddUpperLetterBucket add upper letters to current bucket and calculate bucket related values
func (r *RandString) AddUpperLetterBucket() error {
	return r.SetBucket(r.bucket + "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

// AddLowerLetterBucket add lower letters to current bucket and calculate bucket related values
func (r *RandString) AddLowerLetterBucket() error {
	return r.SetBucket(r.bucket + "abcdefghijklmnopqrstuvwxyz")
}

// AddSpecialCharacterBucket add special characters to current bucket and calculate bucket related values
func (r *RandString) AddSpecialCharacterBucket() error {
	return r.SetBucket(r.bucket + " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")
}

// GetBucket get current bucket
func (r *RandString) GetBucket() string {
	return r.bucket
}
