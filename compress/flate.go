package compress

import (
	"bytes"
	"compress/flate"
	"io"
)

func CompressFlate(data []byte) ([]byte, error) {
	var b bytes.Buffer
	// level: flate.DefaultCompression, flate.BestSpeed, etc.
	writer, err := flate.NewWriter(&b, flate.BestCompression)
	if err != nil {
		return nil, err
	}

	_, err = writer.Write(data)
	if err != nil {
		return nil, err
	}

	writer.Close()
	return b.Bytes(), nil
}

func DecompressFlate(compressed []byte) ([]byte, error) {
	reader := flate.NewReader(bytes.NewReader(compressed))
	defer reader.Close()

	var out bytes.Buffer
	_, err := io.Copy(&out, reader)
	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}
