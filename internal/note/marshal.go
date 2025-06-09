package note

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/Sn0wo2/QuickNote/pkg/compress"
)

const (
	fieldTitle   = 1
	fieldContent = 2
	magicVersion = 1
)

// magicHeader First 2 bytes of SHA-256 hash of "QuickNote"
var magicHeader = [...]byte{0x34, 0x19}

func (n *Note) Encode() error {
	if n == nil {
		return errors.New("nil note")
	}

	buf := bytes.NewBuffer(make([]byte, 0, 256))
	buf.Grow(len(magicHeader) + 1 + 1 + 4 + len(n.Title) + 1 + 4 + len(n.Content))

	buf.Write(magicHeader[:])
	buf.WriteByte(magicVersion)

	writeField := func(id byte, data []byte) {
		buf.WriteByte(id)
		_ = binary.Write(buf, binary.LittleEndian, uint32(len(data)))
		buf.Write(data)
	}

	writeField(fieldTitle, n.Title)
	writeField(fieldContent, n.Content)

	var err error
	n.Data, err = compress.FlateCompress(buf.Bytes())
	n.Title, n.Content = nil, nil

	return err
}

func (n *Note) Decode(data []byte) error {
	if n == nil {
		return errors.New("nil note")
	}

	decompressed, err := compress.FlateDecompress(data)
	if err != nil {
		return fmt.Errorf("decompress: %w", err)
	}

	if len(decompressed) < len(magicHeader)+1 {
		return errors.New("data too short")
	}

	if !bytes.Equal(decompressed[:len(magicHeader)], magicHeader[:]) {
		return errors.New("invalid magic header")
	}

	r := bytes.NewReader(decompressed[len(magicHeader):])
	if v, err := r.ReadByte(); err != nil || v != magicVersion {
		return fmt.Errorf("invalid version: %w", err)
	}

	buf := make([]byte, 4)
	for {
		id, err := r.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("read field id: %w", err)
		}

		if _, err := io.ReadFull(r, buf); err != nil {
			return fmt.Errorf("read field %d length: %w", id, err)
		}

		length := binary.LittleEndian.Uint32(buf)
		value := make([]byte, length)
		if _, err := io.ReadFull(r, value); err != nil {
			return fmt.Errorf("read field %d: %w", id, err)
		}

		switch id {
		case fieldTitle:
			n.Title = value
		case fieldContent:
			n.Content = value
		default:
			return fmt.Errorf("unknown field: %d", id)
		}
	}

	n.Data = nil
	return nil
}
