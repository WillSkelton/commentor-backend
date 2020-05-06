package sourcefile

import (
	"commentor-backend/lib/function"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	sfIDTracker = uint64(0)
)

// SourceFile contains all the information about a source code file
type SourceFile struct {
	Path      string
	FileName  string
	Lang      Language
	FileID    uint64
	Functions map[uint64]*function.Function
}

type Language struct {
	Extension  string
	formatFunc (func(string) string)
	parseFunc  (func(string) map[uint64]*function.Function)
}

// NewSourceFile Creates a new NewSourceFile object
func NewSourceFile(path string) (sf *SourceFile, err error) {

	// allocate memory
	sf = &SourceFile{}

	// set path
	sf.Path = path

	// get file extension
	ext := strings.Replace(filepath.Ext(sf.Path), ".", "", 1)
	sf.Lang.Extension = ext

	// get filename
	sf.FileName = filepath.Base(path)

	// strip off ext
	sf.FileName = strings.ReplaceAll(sf.FileName, filepath.Ext(sf.FileName), "")

	// set Functions
	sf.Lang.formatFunc = Formatters[sf.Lang.Extension]
	sf.Lang.parseFunc = Parsers[sf.Lang.Extension]

	sf.FileID = sfIDTracker

	sf.Functions = make(map[uint64]*function.Function)
	if err = sf.GatherFunctions(); err != nil {
		return nil, err
	}

	sfIDTracker++
	return
}

// GatherFunctions will traverse a file and fill the file's function map with function objects
func (sf *SourceFile) GatherFunctions() (err error) {
	// fmt.Println(sf)

	if _, err := os.Stat(sf.Path); os.IsNotExist(err) {
		return err
	}

	var data []byte
	if data, err = ioutil.ReadFile(sf.Path); err != nil {
		return err
	}

	sf.Functions = sf.Lang.parseFunc(string(data))

	return nil
}

// String() implements the Stinger interface so a SourceFile can be all dolled up before it's printed
func (sf SourceFile) String() string {
	str := ""
	str += fmt.Sprintf("---------------------------------------\n")
	str += fmt.Sprintf("Path: %v\n", sf.Path)
	str += fmt.Sprintf("FileName: %v\n", sf.FileName)
	str += fmt.Sprintf("Extension: %v\n", sf.Lang.Extension)
	str += fmt.Sprintf("FileID: %v\n", sf.FileID)
	str += fmt.Sprintf("Functions: %v\n", len(sf.Functions))
	str += fmt.Sprintf("---------------------------------------\n")

	return str
}
