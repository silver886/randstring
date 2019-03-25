package randstring

import "errors"

// Get generates a random string
func (r *RandString) Get(n int) (string, error) {
	if r.src == nil {
		return "", errors.New("Random source is nil")
	}
	if r.bucketLength == 0 {
		return "", errors.New("Bucket is empty")
	}

	b := make([]byte, n)
	for i, cache, remain := n-1, r.src.Uint64(), r.bucketMax; i >= 0; {
		if remain == 0 {
			cache, remain = r.src.Uint64(), r.bucketMax
		}
		if idx := cache & r.bucketMask; idx < r.bucketLength {
			b[i] = r.bucket[idx]
			i--
		}
		cache >>= r.bucketBit
		remain--
	}

	return string(b), nil
}
