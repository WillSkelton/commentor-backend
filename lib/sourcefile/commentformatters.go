package sourcefile

import (
	"fmt"
	"strings"
)

var (
	formatters = map[string](func(string) string){
		"go": goComment,
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
