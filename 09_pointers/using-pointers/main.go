package main

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	fmt.Println(b)
	fmt.Println(&b)

	*b = 42 // says "the value at this address, change it to 42"
	fmt.Println(a) // 42
}

// this is useful
// we can pass a memory address instead of a bunch of values (we can pass a reference)
// and then we can still change the value of whatever is stored at that memory address
// this makes our programs more performant
// we don't have to pass around large amounts of data

// everything is PASS BY VALUE in go 
// when we pass a memory address, we are passing a value
