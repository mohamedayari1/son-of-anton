package file


import (
	"testing"
)



func TestPageWriteRead(t *testing.T) {
	page := NewPage(16)
	data := []byte("Hello, World!")
	n, err := page.Write(0, data)

	if err != nil {
		t.Fatal("unexpected error on write:", err)
	}

	if n != len(data) {
		t.Fatalf("expected to write %d bytes, wrote %d", len(data), n)
	}

	got := make([]byte, len(data))
	page.Read(0, got)
	want := data
	
	if string(got) != string(want) {
		t.Fatalf("expected to read %q, got %q", want, got)
	}



	data = []byte("SimpleDB!")
	n, err = page.Write(7, data)

	if err != nil {
		t.Fatalf("Write returned : %v", err)
	}

	if n != len(data) {
		t.Fatalf("expected to write %d bytes, wrote %d", len(data), n)
	}

	got = page.Bytes()
	want = []byte("Hello, SimpleDB!")

	if string(got) != string(want) {
		t.Fatalf("expected to read %q, got %q", want, got)
	}



}







