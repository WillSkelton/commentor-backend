package function

import "fmt"

var (
	funcIDTracker = uint64(0)
)

// Function keeps track of a function as it appears in a file
type Function struct {
	Comment      string
	Contents     string
	FunctionName string
	FuncID       uint64
	Location     uint64
	StartLine    uint64
	EndLine      uint64
}

// NewFunction creates a new function object
func NewFunction(comment, contents, functionName string, location, startLine, endLine uint64) (f *Function) {
	f = &Function{}

	f.StartLine = startLine
	f.EndLine = endLine
	f.Comment = comment
	f.Contents = contents
	f.FunctionName = functionName
	f.FuncID = funcIDTracker
	f.Location = location

	funcIDTracker++

	return
}

func (f Function) String() string {
	str := ""
	str += fmt.Sprintf("---------------------------------------\n")
	str += fmt.Sprintf("Comment: %v\n", f.Comment)
	str += fmt.Sprintf("Contents: %v\n", f.Contents)
	str += fmt.Sprintf("FunctionName: %v\n", f.FunctionName)
	str += fmt.Sprintf("FuncID: %v\n", f.FuncID)
	str += fmt.Sprintf("Location: %v\n", f.Location)
	str += fmt.Sprintf("StartLine: %v\n", f.StartLine)
	str += fmt.Sprintf("EndLine: %v\n", f.EndLine)
	str += fmt.Sprintf("---------------------------------------\n")

	return str
}
