package transaction

import (
	"fmt"
)

func transaction_main() {
	fmt.Println("Hello")
}

type Transaction struct {
	amount              float32
	kind, account, date string
	//Need to change date type to another format
	//Do we need a boolean
}

/*
*************Transaction(struct type is basically a blue pront of the data it will hold)
 */

//I have intiated this struct with capital letter, may be this will be used in other packages,
//WIll have to check how to do that

func Help() {
}

// This service
//How can I add console output in browser
//Ofc that can't be done, as browser only interacts with front end

//Things defined with a starting capital letter can be exported...
