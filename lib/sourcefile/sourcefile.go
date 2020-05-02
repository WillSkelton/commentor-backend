package sourcefile

import (
	"commentor-backend/lib/function"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	sfIDTracker = uint64(0)
)

// SourceFile contains all the information about a source code file
type SourceFile struct {
	Path      string
	FileName  string
	lang      language
	FileID    uint64
	functions map[uint64]*function.Function
}

type language struct {
	extension  string
	formatFunc (func(string) string)
}

// NewSourceFile Creates a new NewSourceFile object
func NewSourceFile(path string) (sf *SourceFile, err error) {
	sf = &SourceFile{}
	sf.Path = path

	sf.FileName = filepath.Base(path)
	sf.FileName = sf.FileName[:len(sf.FileName)-len(filepath.Ext(sf.FileName))]

	sf.lang.extension = filepath.Ext(sf.Path)
	sf.lang.formatFunc = formatters[sf.lang.extension]

	sf.FileID = sfIDTracker
	sf.functions = make(map[uint64]*function.Function, 0)
	if err = sf.GatherFunctions(); err != nil {
		return nil, err
	}

	sfIDTracker++
	return
}

// GatherFunctions will traverse a file and fill the file's function map with function objects
func (sf *SourceFile) GatherFunctions() (err error) {
	fmt.Println(sf)

	if _, err := os.Stat(sf.Path); os.IsNotExist(err) {
		return err
	}

	var data []byte
	if data, err = ioutil.ReadFile(sf.Path); err != nil {
		return err
	}

	fmt.Println(string(data))

	return nil
}

// String() implements the Stinger interface so a SourceFile can be all dolled up before it's printed
func (sf SourceFile) String() string {
	str := ""
	str += fmt.Sprintf("---------------------------------------\n")
	str += fmt.Sprintf("Path: %v\n", sf.Path)
	str += fmt.Sprintf("FileName: %v\n", sf.FileName)
	str += fmt.Sprintf("extension: %v\n", sf.lang.extension)
	str += fmt.Sprintf("FileID: %v\n", sf.FileID)
	str += fmt.Sprintf("functions: %v\n", len(sf.functions))
	str += fmt.Sprintf("---------------------------------------\n")

	return str
}
