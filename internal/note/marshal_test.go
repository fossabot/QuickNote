package note

import (
	"bytes"
	"math/rand"
	"testing"
)

func generateRandomNote() *Note {
	titleSize := rand.Intn(100) + 50
	contentSize := rand.Intn(1000) + 500

	return &Note{
		Title:   bytes.Repeat([]byte{'a'}, titleSize),
		Content: bytes.Repeat([]byte{'b'}, contentSize),
	}
}

func BenchmarkEncode(b *testing.B) {
	note := generateRandomNote()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = note.Encode()
	}
}

func BenchmarkDecode(b *testing.B) {
	note := generateRandomNote()
	if err := note.Encode(); err != nil {
		b.Fatalf("Failed to encode note for benchmark: %v", err)
	}
	encodedData := note.Data

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		newNote := &Note{Data: encodedData}
		_ = newNote.Decode(encodedData)
	}
}

func BenchmarkEncodeDecode(b *testing.B) {
	note := generateRandomNote()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = note.Encode()
		newNote := &Note{Data: note.Data}
		_ = newNote.Decode(note.Data)
	}
}

func BenchmarkEncodeLargeNote(b *testing.B) {
	largeNote := &Note{
		Title:   bytes.Repeat([]byte{'a'}, 10*1024),  // 10KB title
		Content: bytes.Repeat([]byte{'b'}, 100*1024), // 100KB content
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = largeNote.Encode()
	}
}

func BenchmarkDecodeLargeNote(b *testing.B) {
	largeNote := &Note{
		Title:   bytes.Repeat([]byte{'a'}, 10*1024),  // 10KB title
		Content: bytes.Repeat([]byte{'b'}, 100*1024), // 100KB content
	}
	if err := largeNote.Encode(); err != nil {
		b.Fatalf("Failed to encode large note for benchmark: %v", err)
	}
	encodedData := largeNote.Data

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		newNote := &Note{Data: encodedData}
		_ = newNote.Decode(encodedData)
	}
}
