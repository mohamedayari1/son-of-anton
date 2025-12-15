package file

import (
	"errors"
)


type Page struct {
	bytes []byte
}


func NewPage(n int) *Page {
	return &Page{
		bytes : make([]byte, n),
	}
}


func (page *Page) Write(offset int, data []byte) (int, error) {
	if offset +len(data) > len(page.bytes) {
		return 0, errors.New("write exceeds page size")
	}

	n := copy(page.bytes[offset:], data)
	
	return n, nil  
}



func (page *Page) Read(offset int, dst []byte) int {

	return copy(dst, page.bytes[offset:])	

}

func (page *Page) Bytes() []byte {

	return page.bytes 
}


func (page *Page) Size() int {

	return len(page.bytes) 
}