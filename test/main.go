package main

import (
	"fmt"
	"reflect"
	"test/files"
)

var (
	file_path = "../sample.json"
	spending  float64
	balance   float64
)

/*
######################################
If there's variable is an interface,
To use it in some operation always assert that variable,
Can't even iterate interface unless you the variable is asserted into
iteratable type
######################################

*/

// func read_file_err(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }
func main() {
	json_map := files.Parsejson(file_path)
	//fmt.Println("This is json Map", files.Parsejson(file_path))
	//fmt.Println("This is type of file", reflect.TypeOf(file_path))
	fmt.Println("Type of JSON_MAP", reflect.TypeOf(json_map))
	fmt.Println("This is json parsed data:", json_map)

	for k, v := range json_map {
		//Change this whole logic in the next commit
		//Try using Switch Case
		//How json is parsed totally depends on the json format,
		//So play around with json to write a bug free logic
		if k == "Balance" {
			balance = v.(float64)
			fmt.Println("Opening Balance for this month is :", v)
		} else if k == "Debit" {
			//fmt.Printf("This is type of value %T", v)
			if value, ok := v.([]interface{}); ok {
				//fmt.Println(value)
				for _, val := range value {
					if str, ok1 := val.(map[string]interface{}); ok1 {
						for i, j := range str {
							fmt.Printf("You have spent %g on %s \n", j, i)
							spending = spending + j.(float64)
						}
					}
				}
			} else {
				fmt.Println("Not a map")
			}
		}

	}

	// for k, v := range json_map {
	// 	switch i := v.(type) {
	// 	case map[string]interface{}:
	// 		fmt.Println("Hello")
	// 	case []interface{}:
	// 		fmt.Println("Hello[]interface{}")

	// 	}

	// }

	fmt.Println("This is the total spending for the month x : ", spending)
	fmt.Println("This is the current Balance: ", balance-spending)
	/*
		**************
		Range doesn't iterate over a map in an order, it can iterate in any order
		**************
		##############
		PS C:\Users\SAITEJ\myWorkspace\projects\Money_Manager_App\test> go run .
		This is a key Debit
		This is a value [map[Body Shop:300 Rent:796 Travel:100.6]]
		This is a key Balance
		This is a value 40000
		PS C:\Users\SAITEJ\myWorkspace\projects\Money_Manager_App\test> go run .
		This is a key Balance
		This is a value 40000
		This is a key Debit
		This is a value [map[Body Shop:300 Rent:796 Travel:100.6]]
		PS C:\Users\SAITEJ\myWorkspace\projects\Money_Manager_App\test>
		##############
		As you can see once the compiler printed the map in one order, next time it printed it in
		completely opposite order

	*/
	//json_map["Debit"] is interface type , We cannot iterate over interface type using range
	//.\main.go:23:14: cannot range over json_map["Debit"] (type interface {})

	/*
		$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
		WHat is []interface{}
		It is a slice of elements of any type(interface)
		THe type of that is a slice, it's not interface so this cannot be asserted to anything

		example :
		items := []interface{}{true,"Hello",3.14}
		Here items is a slice of interface where the type of elements are
		true -> bool
		Hello -> String
		3.14 -> Float

		%%%%
		Don't confuse interface with struct as the above example creates some confusion,
		A struct of similar type can be created with bool, string and float as fields in it
		But Interface can be any type , it can be a struct , but a struct can't be an interface
		Example: items := []interface{}{"Hello", "Hi"} , this can just be a slice of string elements
		%%%%
		$#$##############################$$$$$$$$$$$$$$$$$$$$$$$$$
	*/
}

/*
I need to write a switch case to handle different kind of json datas that come, I have written for
[]interface{}, for JSON arrays
map[string]interface{}, for JSON objects

json_map is map[string]interface{}
v in Line 47 is []interface{}
*/
