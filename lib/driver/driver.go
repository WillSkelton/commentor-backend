package driver

import "commentor-backend/lib/sourcefile"

// Driver is the singleton that keeps track of the whole directory
type Driver struct {
	WorkingDirectory string
	FileManager      map[uint64]*sourcefile.SourceFile
}

// NewDriver makes a new Driver object
func NewDriver(wd string) (d *Driver) {
	d = &Driver{}
	d.WorkingDirectory = wd

	d.gatherFiles()

	return
}

func (d *Driver) gatherFiles() {

}
