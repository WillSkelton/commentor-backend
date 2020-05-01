package sourcefile

import (
	"commentor-backend/lib/function"
	"fmt"
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
	fileID    uint64
	functions map[uint64]*function.Function
}

type language struct {
	extension  string
	formatFunc (func(string) string)
}

// NewSourceFile Creates a new NewSourceFile object
func NewSourceFile(path, filename string) (sf *SourceFile) {
	sf.Path = path
	sf.FileName = filename

	sf.lang.extension = filepath.Ext(sf.Path)
	sf.lang.formatFunc = formatters[sf.lang.extension]

	sf.fileID = sfIDTracker
	sf.GatherFunctions()

	sfIDTracker++
	return
}

// GatherFunctions will traverse a file and fill the file's function map with function objects
func (sf *SourceFile) GatherFunctions() {
	fmt.Println()
}
