package driver

import (
	"commentor-backend/lib/sourcefile"
	"os"
	"path/filepath"
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

	d.gatherFiles()

	return
}

func (d *Driver) gatherFiles() {
	filepath.Walk(d.WorkingDirectory, func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() {
			sf := &sourcefile.SourceFile{}
			sf = sourcefile.NewSourceFile(path)

			d.FileManager[sf.FileID] = sf
		}
		return nil
	})

}
