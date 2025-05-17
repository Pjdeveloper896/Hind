// main.go
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Function struct {
	params []string
	body   []string
}

var variables = make(map[string]string)
var functions = make(map[string]Function)

func getValue(s string) string {
	if val, ok := variables[s]; ok {
		return val
	}
	return s
}

func evalArithmetic(op string, a, b string) string {
	x, err1 := strconv.Atoi(getValue(a))
	y, err2 := strconv.Atoi(getValue(b))
	if err1 != nil || err2 != nil {
		return "Error: Invalid number"
	}
	switch op {
	case "jod":
		return strconv.Itoa(x + y)
	case "ghat":
		return strconv.Itoa(x - y)
	case "guna":
		return strconv.Itoa(x * y)
	case "bhaag":
		if y == 0 {
			return "Error: Division by zero"
		}
		return strconv.Itoa(x / y)
	}
	return "Unknown operation"
}

func evalCondition(cond []string) bool {
	if len(cond) < 3 {
		return false
	}
	left := getValue(cond[0])
	op := cond[1]
	right := getValue(cond[2])

	switch op {
	case "==":
		return left == right
	case "!=":
		return left != right
	case ">":
		li, err1 := strconv.Atoi(left)
		ri, err2 := strconv.Atoi(right)
		if err1 == nil && err2 == nil {
			return li > ri
		}
	case "<":
		li, err1 := strconv.Atoi(left)
		ri, err2 := strconv.Atoi(right)
		if err1 == nil && err2 == nil {
			return li < ri
		}
	case ">=":
		li, err1 := strconv.Atoi(left)
		ri, err2 := strconv.Atoi(right)
		if err1 == nil && err2 == nil {
			return li >= ri
		}
	case "<=":
		li, err1 := strconv.Atoi(left)
		ri, err2 := strconv.Atoi(right)
		if err1 == nil && err2 == nil {
			return li <= ri
		}
	}
	return false
}

func interpretLines(lines []string) (string, bool) {
	i := 0
	for i < len(lines) {
		line := strings.TrimSpace(lines[i])
		if line == "" || strings.HasPrefix(line, "#") {
			i++
			continue
		}
		tokens := strings.Fields(line)
		if len(tokens) == 0 {
			i++
			continue
		}

		switch tokens[0] {
		case "chhapo":
			if len(tokens) < 2 {
				fmt.Println("Kya chhapo? Argument do.")
			} else {
				fmt.Println(getValue(tokens[1]))
			}

		case "pucho":
			if len(tokens) < 2 {
				fmt.Println("Kisse puchna hai?")
			} else {
				fmt.Print(tokens[1] + ": ")
				reader := bufio.NewReader(os.Stdin)
				input, _ := reader.ReadString('\n')
				variables[tokens[1]] = strings.TrimSpace(input)
			}

		case "naam":
			if len(tokens) >= 4 && tokens[2] == "=" {
				variables[tokens[1]] = getValue(tokens[3])
			} else {
				fmt.Println("Sahi syntax nahi hai: naam x = 10")
			}

		case "jod", "ghat", "guna", "bhaag":
			if len(tokens) >= 3 {
				res := evalArithmetic(tokens[0], tokens[1], tokens[2])
				fmt.Println(res)
			} else {
				fmt.Println("Galat arithmetic syntax")
			}

		case "jodo":
			if len(tokens) >= 3 {
				fmt.Println(getValue(tokens[1]) + getValue(tokens[2]))
			} else {
				fmt.Println("Syntax: jodo x y")
			}

		case "agar":
			toIndex := -1
			for idx, tok := range tokens {
				if tok == "to" {
					toIndex = idx
					break
				}
			}
			if toIndex == -1 || toIndex == 1 {
				fmt.Println("Syntax error in agar statement")
				i++
				continue
			}
			condition := tokens[1:toIndex]

			warnaIndex := -1
			khatamIndex := -1
			for j := i + 1; j < len(lines); j++ {
				if strings.TrimSpace(lines[j]) == "warna" {
					warnaIndex = j
				} else if strings.TrimSpace(lines[j]) == "khatam" {
					khatamIndex = j
					break
				}
			}
			if khatamIndex == -1 {
				fmt.Println("khatam missing for agar block")
				i++
				continue
			}

			if evalCondition(condition) {
				end := khatamIndex
				if warnaIndex != -1 {
					end = warnaIndex
				}
				interpretLines(lines[i+1 : end])
			} else {
				if warnaIndex != -1 {
					interpretLines(lines[warnaIndex+1 : khatamIndex])
				}
			}
			i = khatamIndex

		case "ghoom":
			if len(tokens) < 3 || tokens[2] != "baar" {
				fmt.Println("Syntax: ghoom <count> baar")
				i++
				continue
			}
			count, err := strconv.Atoi(getValue(tokens[1]))
			if err != nil {
				fmt.Println("Ghoom mein number do")
				i++
				continue
			}
			khatamIndex := -1
			for j := i + 1; j < len(lines); j++ {
				if strings.TrimSpace(lines[j]) == "khatam" {
					khatamIndex = j
					break
				}
			}
			if khatamIndex == -1 {
				fmt.Println("khatam missing for ghoom block")
				i++
				continue
			}
			for c := 0; c < count; c++ {
				interpretLines(lines[i+1 : khatamIndex])
			}
			i = khatamIndex

		case "banao":
			if len(tokens) < 2 {
				fmt.Println("Function naam do")
				i++
				continue
			}
			funcName := tokens[1]
			paramsStart := strings.Index(funcName, "(")
			paramsEnd := strings.Index(funcName, ")")
			params := []string{}
			if paramsStart != -1 && paramsEnd != -1 && paramsEnd > paramsStart {
				paramList := funcName[paramsStart+1 : paramsEnd]
				if len(paramList) > 0 {
					params = strings.Split(paramList, ",")
					for idx := range params {
						params[idx] = strings.TrimSpace(params[idx])
					}
				}
				funcName = funcName[:paramsStart]
			} else {
				fmt.Println("Function syntax: banao functionName(param1,param2)")
				i++
				continue
			}
			khatamIndex := -1
			for j := i + 1; j < len(lines); j++ {
				if strings.TrimSpace(lines[j]) == "khatam" {
					khatamIndex = j
					break
				}
			}
			if khatamIndex == -1 {
				fmt.Println("Function body khatam missing")
				i++
				continue
			}
			body := lines[i+1 : khatamIndex]
			functions[funcName] = Function{params: params, body: body}
			i = khatamIndex

		case "wapis", "wapis karo":
			if len(tokens) >= 2 {
				return getValue(tokens[1]), true
			}
			return "", true

		case "server":
			if len(tokens) >= 3 && tokens[1] == "shuru" && tokens[2] == "port" {
				port := "8080"
				if len(tokens) >= 4 {
					port = tokens[3]
				}
				go func() {
					fmt.Println("Server chalu port pe:", port)
					http.ListenAndServe(":"+port, nil)
				}()
			}

		case "jab":
			if len(tokens) >= 5 && tokens[3] == "to" {
				method := strings.ToUpper(tokens[1])
				path := tokens[2]
				response := strings.Join(tokens[4:], " ")
				http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
					if r.Method == method {
						fmt.Fprintln(w, getValue(response))
					} else {
						http.NotFound(w, r)
					}
				})
			}

		default:
			if fn, ok := functions[tokens[0]]; ok {
				if len(tokens[1:]) != len(fn.params) {
					fmt.Printf("Function %s ko %d args chahiye\n", tokens[0], len(fn.params))
					i++
					continue
				}
				oldVars := make(map[string]string)
				for k, v := range variables {
					oldVars[k] = v
				}
				for idx, param := range fn.params {
					variables[param] = getValue(tokens[idx+1])
				}
				res, returned := interpretLines(fn.body)
				if returned {
					fmt.Println("Wapas:", res)
				}
				variables = oldVars
			} else {
				fmt.Println("Samjha nahi:", tokens[0])
			}
		}
		i++
	}
	return "", false
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Script file ka naam do")
		return
	}
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("File nahi mila:", err)
		return
	}
	lines := strings.Split(string(file), "\n")
	interpretLines(lines)
}
