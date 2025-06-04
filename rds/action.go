package rds

import (
	"context"
	"errors"
	"time"

	"github.com/Sn0wo2/QuickNote/compress"
)

func SetCompressedBytes(key string, value []byte, expire time.Duration) error {
	if len(value) == 0 {
		return errors.New("empty value")
	}

	value, err := compress.FlateCompress(value)
	if err != nil {
		return err
	}

	return Instance.Get().Set(context.Background(), key, value, expire).Err()
}

func GetCompressedBytes(key string) ([]byte, error) {
	value, err := Instance.Get().Get(context.Background(), key).Bytes()
	if err != nil {
		return nil, err
	}

	value, err = compress.FlateDecompress(value)
	if err != nil {
		return nil, err
	}

	return value, nil
}
