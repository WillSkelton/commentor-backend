package function

var (
	funcIDTracker = uint64(0)
)

// Function keeps track of a function as it appears in a file
type Function struct {
	comment      string
	contents     string
	functionName string
	funcID       uint64
	location     uint64
}

// NewFunction creates a new function object
func NewFunction(comment, contents, functionName string, location uint64) (f *Function) {
	f = &Function{}

	f.comment = comment
	f.contents = contents
	f.functionName = functionName
	f.funcID = funcIDTracker
	f.location = location

	funcIDTracker++

	return
}
