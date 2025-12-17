package file

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileMgrWriteRead(t *testing.T) {
	dataDir := "testdata"
	blockSize := 16
	testFile := "writetestfile"

	mgr := NewFileMgr(dataDir, blockSize)

	t.Cleanup(func() {
		mgr.Close()
		os.Remove(filepath.Join(dataDir, testFile))
	})

	blockZero := &BlockID{
		Filename: testFile,
		Number: 0,
	}
	data := "aaaabbbbccccdddd"
	checkWrite(t, mgr, blockZero, data)
	checkRead(t, mgr, blockZero, data)
	checkFileContent(t, filepath.Join(dataDir, testFile), data)

}


func checkWrite(t *testing.T, mgr *FileMgr, blockID *BlockID, data string) {
	page := NewPage(mgr.blockSize)
	page.Write(0, []byte(data))

	n, err := mgr.Write(blockID, page)
	if err != nil {
		t.Fatalf("Write returned error: %v", err)
	}
	if n != len(data) {
		t.Fatalf("expected to write %d bytes, wrote %d", len(data), n)
	}
	
}

func checkRead(t *testing.T, mgr *FileMgr, blockID *BlockID, want string) {
	page := NewPage(mgr.blockSize)
	n, err := mgr.Read(blockID, page)
	if err != nil {
		t.Fatalf("Read failed: %v", err)
	}
	if n != mgr.blockSize {
		t.Fatalf("expected to read %d bytes, read %d", mgr.blockSize, n)
	}	
	if string(page.Bytes()) != want {
		t.Fatalf("expected to read %q, got %q", want, string(page.Bytes()))
	}
	
}



func checkFileContent(t *testing.T, path string, want string) {
	file, err	 := os.Open(path)
	if err != nil {
		t.Fatalf("failed to open file %s: %v", path, err)
	}
	got := make([]byte, len(want))
	_, err = file.Read([]byte(got))
	
	if err != nil {
		t.Fatalf("Failed to read file %v", err)

	}
	defer file.Close()

	if string(got) != want {
		t.Fatalf("expected file content %q, got %q", want, string(got))
	}

}