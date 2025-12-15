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


func (fm *FileMgr) Read(blockID *BlockID, p *Page) (int, error) {
	fm.mu.Lock()
	defer fm.mu.Unlock()

	file, err := fm.getFile(blockID.Filename)
	if err != nil {
		return 0, err
	}

	n, err := file.ReadAt(p.Bytes(), int64(blockID.Number*fm.blockSize))

}


func (fm *FileMgr) getFile(fileName string) (*os.File, error) {
	path := filepath.Join(fm.dataDir, fileName)

	file, ok := fm.openedFiles[fileName]
	if !ok {
		var err error
		file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			return nil, fmt.Errorf("failed to open or create file %s: %v", path, err)
		}
		fm.openedFiles[fileName] = file
	}
	return file, nil
}




