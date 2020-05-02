package function

import "fmt"

var (
	funcIDTracker = uint64(0)
)

// Function keeps track of a function as it appears in a file
type Function struct {
	comment      string
	contents     string
	functionName string
	FuncID       uint64
	location     uint64
	startLine    uint64
	endLine      uint64
}

// NewFunction creates a new function object
func NewFunction(comment, contents, functionName string, location, startLine, endLine uint64) (f *Function) {
	f = &Function{}

	f.startLine = startLine
	f.endLine = endLine
	f.comment = comment
	f.contents = contents
	f.functionName = functionName
	f.FuncID = funcIDTracker
	f.location = location

	funcIDTracker++

	return
}

func (f Function) String() string {
	str := ""
	str += fmt.Sprintf("---------------------------------------\n")
	str += fmt.Sprintf("comment: %v\n", f.comment)
	str += fmt.Sprintf("contents: %v\n", f.contents)
	str += fmt.Sprintf("functionName: %v\n", f.functionName)
	str += fmt.Sprintf("FuncID: %v\n", f.FuncID)
	str += fmt.Sprintf("location: %v\n", f.location)
	str += fmt.Sprintf("startLine: %v\n", f.startLine)
	str += fmt.Sprintf("endLine: %v\n", f.endLine)
	str += fmt.Sprintf("---------------------------------------\n")

	return str
}
