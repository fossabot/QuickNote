package note

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/Sn0wo2/QuickNote/compress"
	"github.com/Sn0wo2/QuickNote/helper"
)

const (
	fieldTitle   = 1
	fieldContent = 2
	magicHeader  = "NOTE"
	magicVersion = 1
)

func (n *Note) Encode() error {
	var buf bytes.Buffer

	buf.WriteString(magicHeader)
	buf.WriteByte(magicVersion)

	for _, field := range []struct {
		id   byte
		data []byte
	}{
		{fieldTitle, n.Title},
		{fieldContent, n.Content},
	} {
		if err := buf.WriteByte(field.id); err != nil {
			return err
		}

		if err := binary.Write(&buf, binary.LittleEndian, uint32(len(field.data))); err != nil {
			return err
		}

		if _, err := buf.Write(field.data); err != nil {
			return err
		}
	}

	compressData, err := compress.FlateCompress(buf.Bytes())
	if err != nil {
		return err
	}

	n.Data = compressData
	n.Title = nil
	n.Content = nil

	return nil
}

func (n *Note) Decode(data []byte) error {
	var err error

	data, err = compress.FlateDecompress(data)
	if err != nil {
		return fmt.Errorf("decompression failed: %w", err)
	}

	if !bytes.HasPrefix(data, helper.StringToBytes(magicHeader)) {
		return errors.New("invalid magic header")
	}

	r := bytes.NewReader(data[len(magicHeader):])

	version, err := r.ReadByte()
	if err != nil || version != magicVersion {
		return fmt.Errorf("unsupported or invalid version: %d", version)
	}

	defer func() {
		n.Data = nil
	}()

	for {
		id, err := r.ReadByte()
		if err == io.EOF {
			break
		}

		if err != nil {
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
		if _, err = io.ReadFull(r, val); err != nil {
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
