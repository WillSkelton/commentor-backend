package sourcefile

import (
	"commentor-backend/lib/function"
	"fmt"
	"regexp"
	"strings"
)

const (
	commentSearch = "CommentSearch"
	commentStart  = "CommentStart"
	funcStart     = "FunctionStart"
	funcEnd       = "FunctionEnd"
)

var (
	cFuncMatch, _ = regexp.Compile("(?m)[a-zA-Z0-9\\[\\]\\*]+ [a-zA-Z0-9\\[\\]\\*]+\\(([a-zA-Z0-9\\[\\]\\,*]+ [a-zA-Z0-9\\[\\]\\*]+)*\\)")
	// word = [a-zA-Z0-9\\[\\]\\*]+
	// this matches: word word((word word,?)*) the ',' is optional with following ?
	// word word(word word, word *word[], word word, ..., word word)
	// word word(word word)
	// word word()

	Formatters = map[string](func(string) string){
		"go":  dblSlashComment,
		"c":   dblSlashComment,
		"cpp": dblSlashComment,
		"js":  dblSlashComment,
		"jsx": dblSlashComment,
	}

	Parsers = map[string](func(string) map[uint64]*function.Function){
		"go":  ParseGo,
		"c":   ParseC,
		"cpp": ParseC,
		"js":  ParseJs,
		"jsx": ParseJs,
	}
)

func dblSlashComment(str string) (comment string) {
	commentLines := strings.Split(strings.ReplaceAll(str, "\r", ""), "\n")
	for _, line := range commentLines {
		if line == "" {
			continue
		}

		comment += fmt.Sprintf("// %v\n", line)
	}
	return
}

// ParseGo : parses go code
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
		lineIdx := uint64(idx + 1)
		// Searching for comment or "func"/"type" keywords
		strings.ReplaceAll(line, "\r", "")

		// We found a comment. Transition state to commentStart
		if strings.HasPrefix(line, "//") && state != commentStart {
			state = commentStart
			startLine = lineIdx

		} else if strings.Contains(line, "func") || strings.Contains(line, "type") {

			// we found the function keyword so we transition to funcStart state
			if state == commentSearch {
				// If we're coming from commentSearch, that means that we didn't have a comment so we set startLine to idx
				startLine = lineIdx

			}
			// otherwise, we're coming from commentStart, that means that we had a comment so we leave startLine as it is
			state = funcStart
		} else if strings.HasPrefix(line, "}") {
			state = funcEnd
			endLine = lineIdx

		} else if !(strings.HasPrefix(line, "//")) && state != funcStart {
			state = commentSearch
			comment = ""
			startLine = 0
			endLine = 0

		}

		switch state {
		case commentSearch:
			continue
		case commentStart:
			comment += fmt.Sprintf("%v\n", line)
		case funcStart:
			functionContent += fmt.Sprintf("%v\n", line)

		case funcEnd:
			// add the closing brace
			functionContent += fmt.Sprintf("%v\n", line)
			endLine = uint64(idx)

			// create a new function object with the information we got
			f := function.NewFunction(comment, functionContent, "noNameYet", 0, startLine, endLine)

			// add that to our map
			functions[uint64(f.FuncID)] = f

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

// ParseC : parses C code
func ParseC(code string) (functions map[uint64]*function.Function) {

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
		if strings.HasPrefix(line, "//") {
			state = commentStart
		} else if cFuncMatch.MatchString(line) ||
			(strings.Contains(line, "template") && strings.Contains(line, "typename")) {
			if state == commentSearch {
				// If we're coming from commentSearch, that means that we didn't have a comment so we set startLine to idx
				startLine = uint64(idx + 1)
			}
			state = funcStart
		} else if strings.Contains(line, "struct") && strings.Contains(line, "{") {
			if state == commentSearch {
				startLine = uint64(idx + 1)
			}
			state = funcStart
		} else if strings.HasPrefix(line, "}") {
			state = funcEnd
		} else if !(strings.HasPrefix(line, "//")) && state != funcStart {
			state = commentSearch
			comment = ""
		}

		switch state {
		case commentSearch:
			continue
		case commentStart:
			fmt.Printf("[ParseC]: commentStart=%d\n", idx+1)
			startLine = uint64(idx + 1)
			comment += fmt.Sprintf("%s\n", line)
		case funcStart:
			fmt.Printf("[ParseC]: funcStart=%d\n", idx+1)
			functionContent += fmt.Sprintf("%v\n", line)
		case funcEnd:
			fmt.Printf("[ParseC]: funcEnd=%d\n", idx+1)
			endLine = uint64(idx + 1)
			// add the closing brace
			functionContent += fmt.Sprintf("%v\n", line)
			// create a new function object with the information we got
			f := function.NewFunction(comment, functionContent, "noNameYet", 0, startLine, endLine)
			// add that to our map
			functions[uint64(f.FuncID)] = f
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

// ParseJs parses js
func ParseJs(code string) (functions map[uint64]*function.Function) {

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

		} else if strings.Contains(line, "function") || strings.Contains(line, "=>") {

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

		switch state {
		case commentSearch:
			continue
		case commentStart:
			fmt.Printf("[ParseC]: commentStart=%d\n", idx+1)
			startLine = uint64(idx)
			comment += fmt.Sprintf("%v\n", line)
		case funcStart:
			fmt.Printf("[ParseC]: funcStart=%d\n", idx+1)
			startLine = uint64(idx)
			functionContent += fmt.Sprintf("%v\n", line)

		case funcEnd:
			fmt.Printf("[ParseC]: funcEnd=%d\n", idx+1)
			// add the closing brace
			functionContent += fmt.Sprintf("%v\n", line)
			endLine = uint64(idx)

			// create a new function object with the information we got
			f := function.NewFunction(comment, functionContent, "noNameYet", 0, startLine, endLine)

			// add that to our map
			functions[uint64(f.FuncID)] = f

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
