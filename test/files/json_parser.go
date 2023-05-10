package files

import (
	"encoding/json"
	"io/ioutil"
	"os"
	//"fmt"
)

/*	/*file_data, e := os.ReadFile(file_path)
	//ReadFile takes file path as string
	read_file_err(e)
	fmt.Println(string(file_data))
	fmt.Print("THis is type of value that holds file contents \n", reflect.TypeOf(file_data))
	fmt.Printf("%T\n", file_data) //Both the above lins perform same function : give the type of value
	// THis is type of value that holds file contents (type of file_data)
	// []uint8
	//How does ReadFile work at os level? : ReadFile copy entire content of file into memory in form of variable
*/

//***********************************************************************************************
/*
	Byte arrays are commonly used in programming languages to store and manipulate binary data,
	like files or network communication,
	 and to represent strings of characters in a more memory-efficient way.
*/
//***********************************************************************************************
// file_data, e := os.Open(file_path)
// read_file_err(e)

// byteValue, _ := ioutil.ReadAll(file_data)
// //fmt.Println(byteValue)

// defer file_data.Close()
// var json_data map[string]interface{}
// json.Unmarshal([]byte(byteValue), &json_data)

// fmt.Println(json_data["Debit"])
/*
	Unmarshal json:
	json array are converted into either slices or Go Arrays : ["2", "3"]
	json objects into Go Maps -> { "1": "Hello", "2": false}
*/
/*
############################################################################
	Slices are most commonly used inplace of arrays, as Slices are dynamic
	whereas arrays have fixed length
############################################################################
*/

//fmt.Printf("%T", json_data)

func read_file_err(e error) {
	if e != nil {
		panic(e)
	}
}

func Parsejson(file_path string) map[string]interface{} {
	file_data, e := os.Open(file_path)
	read_file_err(e)
	byteValue, _ := ioutil.ReadAll(file_data)
	defer file_data.Close()
	var json_data map[string]interface{}
	json.Unmarshal([]byte(byteValue), &json_data)

	return json_data
}

// Returning map with string as key and interface as value

//What is interface?

// There might be a need to separate read file and parse json as two functions in the future

/*
IMPORTANT IMPORTANT
############################################################
The defaults that the json package will decode into when the type isn't declared are:

bool, for JSON booleans
float64, for JSON numbers
string, for JSON strings
[]interface{}, for JSON arrays
map[string]interface{}, for JSON objects
nil for JSON null
############################################################
*******
This will come very handy when working with different types of json files.
You can use "switch case" to extract different kinds of json data
switch c := v.type() {
	case string:
		do something
	case float64
		do something else
}
*******
*/
