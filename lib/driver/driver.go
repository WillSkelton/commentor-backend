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
func NewDriver(wd string) (d *Driver, err error) {
	d = &Driver{}
	d.WorkingDirectory = wd

	d.FileManager = make(map[uint64]*sourcefile.SourceFile)

	if err = d.gatherFiles(); err != nil {
		return nil, err
	}
	return d, nil
}

func (d *Driver) gatherFiles() (err error) {
	if err = filepath.Walk(d.WorkingDirectory, func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() {
			sf := &sourcefile.SourceFile{}
			if sf, err = sourcefile.NewSourceFile(path); err != nil {
				return err
			}

			d.FileManager[sf.FileID] = sf
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
