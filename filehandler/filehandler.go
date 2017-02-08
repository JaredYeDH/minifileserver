package filehandler

import (
	"errors"
	"io/ioutil"
	"strings"
	"sync"

	. "github.com/ecofast/sysutils"
)

type FileHandler struct {
	filePath   string
	mutex      sync.Mutex
	fileCaches map[string][]byte
}

func (fh *FileHandler) Initialize(filePath string) {
	fh.filePath = filePath
	fh.fileCaches = make(map[string][]byte)
}

func (fh *FileHandler) GetFile(filename string) ([]byte, error) {
	s := strings.ToLower(filename)
	fh.mutex.Lock()
	defer fh.mutex.Unlock()
	buf, ok := fh.fileCaches[s]
	if !ok {
		s = fh.filePath + filename
		if FileExists(s) {
			data, err := ioutil.ReadFile(s)
			if err != nil {
				return nil, err
			}
			fh.add(filename, data)
			return data, nil
		}
		return nil, errors.New("The required file does not exists: " + filename)
	}
	return buf, nil
}

func (fh *FileHandler) add(filename string, filebytes []byte) {
	fh.fileCaches[filename] = filebytes
}
