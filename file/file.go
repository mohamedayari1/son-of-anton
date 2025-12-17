package file

import (
	"sync"
	"os"
	"path/filepath"
	"fmt"


)



type FileMgr struct {
	blockSize int 
	dataDir string 
	openedFiles map[string]*os.File

	mu sync.Mutex
}

func NewFileMgr(dataDir string, blockSize int) *FileMgr {
	return &FileMgr{
		blockSize: blockSize,
		dataDir: dataDir,
		openedFiles: make(map[string]*os.File),
	}
}

func (fm *FileMgr) Read(blockID *BlockID, p *Page) (int, error) {
	fm.mu.Lock()
	defer fm.mu.Unlock()

	file, err := fm.getFile(blockID.Filename)
	if err != nil {
		return 0, err
	}

	n, err := file.ReadAt(p.Bytes(), int64(blockID.Number*fm.blockSize))
	if err != nil && err.Error() != "EOF" {
		return n, fmt.Errorf("failed to read block %v: %v", blockID, err)
	}
	return n, nil 
}


func (fm *FileMgr) Write(blockID *BlockID, p *Page) (int, error) {
	fm.mu.Lock()
	defer fm.mu.Unlock()

	file, err := fm.getFile(blockID.Filename)
	if err != nil {
		return 0, err
	}

	n, err := file.WriteAt(p.Bytes(), int64(blockID.Number*fm.blockSize))
	if err != nil && err.Error() != "EOF" {
		return n, fmt.Errorf("failed to write block %v: %v", blockID, err)
	}
	return n, nil 
}


func (fm *FileMgr) Close() error {
	fm.mu.Lock()
	defer fm.mu.Unlock()
	
	for _, f := range fm.openedFiles {
		err := f.Close()
		if err != nil {
			return fmt.Errorf("failed to close file: %v", err)
		}
	}
	
	return nil
}



func (fm *FileMgr) getFile(fileName string) (*os.File, error) {
	path := filepath.Join(fm.dataDir, fileName)

	file, ok := fm.openedFiles[fileName]
	var err error
	if !ok {
		file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_SYNC, 0666)
		if err != nil {
			return nil, fmt.Errorf("failed to open or create file %s: %v", path, err)
		}
		fm.openedFiles[fileName] = file
	}
	return file, nil
}




