package src

import "fmt"

//first block for constants
const (
	globalConstOne = 4.6
	globalConstTwo = "yoopo"

	globalConstantVariableOne = iota + 3
	globalConstantVariableTwo = iota
)

//second block for constants
const (
	globalConstantInSecondBlock = iota
)

func main() {

	//variable declaration and initialization
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Samvid"
	fmt.Println(firstName)

	// inbuilt functions to deal with complex numbers
	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)

	//pointers
	var lastName *string = new(string)

	*lastName = "Zare"

	fmt.Println("variable is ", *lastName)
	fmt.Println("address of variable is ", lastName)

	middleName := "MiddleNameee" //dereferrencing
	ptr := &middleName
	fmt.Println(ptr, *ptr)
	middleName = "ChangedMiddleName"

	fmt.Println(ptr, *ptr)

	//constants
	const pi = 3
	fmt.Println(pi)
	fmt.Println(c + 3)
	fmt.Println(c + 1.2) //implicit declaration syntax for c
	fmt.Println(globalConstOne, globalConstTwo)
	fmt.Println(globalConstantVariableOne, globalConstantVariableTwo, globalConstantInSecondBlock) //global constant

	palindromeChecker("yoyoy")
}
