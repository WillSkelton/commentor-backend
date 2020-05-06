package sourcefile

import (
	"fmt"
	"strings"
)

var (
	formatters = map[string](func(string) string){
		"go":  goComment,
		"c":   cComment,
		"cpp": cppComment,
		"py":  pyComment,
	}
)

func goComment(str string) (comment string) {
	commentLines := strings.Split(strings.ReplaceAll(str, "\r", ""), "\n")
	for _, line := range commentLines {
		if line == "" {
			continue
		}

		comment += fmt.Sprintf("// %v\n", line)
	}
	return
}

func cComment(str string) (comment string) {
	fmt.Println("Hello World!")
	return
}
func cppComment(str string) (comment string) {
	return
}
func pyComment(str string) (comment string) {
	return
}
