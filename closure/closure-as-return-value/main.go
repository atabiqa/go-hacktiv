package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	var studentList = []string{"airel", "nanda", "mailo", "scannel", "marco"}
// 	var find = findStudent(studentList)
// 	fmt.Println(find("airel"))
// }

// func findStudent(students []string) func(string) string {

// 	return func(s string) string {
// 		var student string
// 		var position int

// 		for i, v := range students {
// 			if strings.ToLower(v) == strings.ToLower(s) {
// 				student = v
// 				position = i
// 				break
// 			}
// 		}
// 		if student == "" {
// 			return fmt.Sprintf("%s does'nt exist!!", s)
// 		}
// 		return fmt.Sprintf("we found %s at position %d", s, position+1)
// 	}
// }
