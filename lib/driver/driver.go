package driver

import (
	"commentor-backend/lib/sourcefile"
	"os"
	"path/filepath"
	"strings"
)

// Driver is the singleton that keeps track of the whole directory
type Driver struct {
	WorkingDirectory string
	FileManager      map[uint64]*sourcefile.SourceFile
}

// NewDriver makes a new Driver object
func NewDriver(wd string) (d *Driver) {
	d = &Driver{}
	d.WorkingDirectory = wd

	d.FileManager = make(map[uint64]*sourcefile.SourceFile)

	return
}

// GatherFiles will walk the working directory of a file and pick out the accepted files from it
func (d *Driver) GatherFiles() (err error) {

	if err = filepath.Walk(d.WorkingDirectory, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			extension := strings.ReplaceAll(filepath.Ext(info.Name()), ".", "")

			if _, exists := sourcefile.Formatters[extension]; exists {
				sf := &sourcefile.SourceFile{}
				if sf, err = sourcefile.NewSourceFile(path); err != nil {
					return err
				}

				d.FileManager[sf.FileID] = sf
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
