package simplevalues

import "fmt"


func main() {

	//simple valude
	//init
	fmt.Println(1+1)
	//string
	fmt.Println("Hello World")
	//boolean
	fmt.Println(true)
	//float
	fmt.Println(3.14)
	//complex
	fmt.Println(1 + 2i)	
	//rune
	fmt.Println('a')
	//byte
	fmt.Println(byte(65))									
	//pointer
	var x int = 42
	var p *int = &x
	fmt.Println(*p)
	//nil pointer
	var np *int = nil
	fmt.Println(np)	
	
}