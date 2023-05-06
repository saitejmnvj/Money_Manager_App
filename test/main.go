package main

import (
	"fmt"
	"reflect"
	"test/files"
)

var (
	file_path = "../sample.json"
)

// func read_file_err(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }
func main() {
	json_map := files.Parsejson(file_path)
	fmt.Println("This is json Map", files.Parsejson(file_path))
	fmt.Println("This is type of file", reflect.TypeOf(file_path))
}
