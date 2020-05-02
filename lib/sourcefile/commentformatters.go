package sourcefile

import (
	"commentor-backend/lib/function"
	"fmt"
	"strings"
)

const (
	commentSearch = "CommentSearch"
	commentStart  = "CommentStart"
	funcStart     = "FunctionStart"
	funcEnd       = "FunctionEnd"
)

var (
	formatters = map[string](func(string) string){
		"go": goComment,
	}

	parsers = map[string](func(string) map[uint64]*function.Function){
		"go": ParseGo,
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

func ParseGo(code string) (functions map[uint64]*function.Function) {

	codeLines := strings.Split(code, "\n")

	functions = make(map[uint64]*function.Function)

	var (
		startLine       uint64
		endLine         uint64
		comment         string
		functionContent string
		state           = commentSearch
	)

	for idx, line := range codeLines {
		// Searching for comment or "func" keyword
		strings.ReplaceAll(line, "\r", "")

		// We found a comment. Transition state to commentStart
		if strings.HasPrefix(line, "//") {
			state = commentStart

		} else if strings.Contains(line, "func") {

			// we found the function keyword so we transition to funcStart state
			if state == commentSearch {
				// If we're coming from commentSearch, that means that we didn't have a comment so we set startLine to idx
				startLine = uint64(idx)
			}
			// otherwise, we're coming from commentStart, that means that we had a comment so we leave startLine as it is
			state = funcStart
		} else if strings.HasPrefix(line, "}") {

			state = funcEnd

		} else if !(strings.HasPrefix(line, "//")) && state != funcStart {
			state = commentSearch
			comment = ""
		}

		// fmt.Println("--------------------------------------------------")
		// fmt.Printf("'%v' ----> %v\n", line, state)

		switch state {
		case commentSearch:
			continue
		case commentStart:
			startLine = uint64(idx)
			comment += fmt.Sprintf("%v\n", line)
		case funcStart:
			startLine = uint64(idx)
			functionContent += fmt.Sprintf("%v\n", line)

		case funcEnd:
			// add the closing brace
			functionContent += fmt.Sprintf("%v\n", line)

			// create a new function object with the information we got
			f := function.NewFunction(comment, functionContent, "noNameYet", 0, startLine, endLine)

			// add that to our map
			functions[uint64(idx)] = f

			fmt.Println(f)
			// reset our state machine
			startLine = 0
			comment = ""
			functionContent = ""
			state = commentSearch

		default:
			continue
		}

	}

	return
}
