package note

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

const (
	fieldTitle   = 1
	fieldContent = 2
	magicHeader  = "NOTE" // 4 bytes
	magicVersion = 1
)

func (n *Note) Encode() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(magicHeader)
	buf.WriteByte(magicVersion)

	writeField := func(id byte, data []byte) error {
		if err := buf.WriteByte(id); err != nil {
			return err
		}
		if err := binary.Write(&buf, binary.LittleEndian, uint32(len(data))); err != nil {
			return err
		}
		_, err := buf.Write(data)
		return err
	}

	if err := writeField(fieldTitle, n.Title); err != nil {
		return nil, err
	}
	if err := writeField(fieldContent, n.Content); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (n *Note) Decode(data []byte) error {
	if len(data) < len(magicHeader)+1 || !bytes.Equal(data[:len(magicHeader)], []byte(magicHeader)) {
		return errors.New("invalid magic header")
	}

	r := bytes.NewReader(data[len(magicHeader):])

	versionByte, err := r.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read version: %w", err)
	}
	if versionByte != magicVersion {
		return fmt.Errorf("unsupported version: %d", versionByte)
	}

	for {
		id, err := r.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		var length uint32
		if err := binary.Read(r, binary.LittleEndian, &length); err != nil {
			return err
		}
		if int64(length) > int64(r.Len()) {
			return fmt.Errorf("field %d length %d exceeds remaining %d", id, length, r.Len())
		}

		val := make([]byte, length)
		if _, err := io.ReadFull(r, val); err != nil {
			return err
		}

		switch id {
		case fieldTitle:
			n.Title = val
		case fieldContent:
			n.Content = val
		default:
			return fmt.Errorf("unknown field id: %d", id)
		}
	}
	return nil
}
