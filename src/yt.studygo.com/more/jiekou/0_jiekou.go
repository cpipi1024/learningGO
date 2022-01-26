package main

type byteCounter int

func (c *byteCounter) Write(b []byte) (int, error) {
	*c += byteCounter(len(b))
	return len(b), nil
}

/* func main() {
	var c byteCounter
	c.Write([]byte("hello"))
	//fmt.Println(c)

	c = 0

	var name = "dolly"

	fmt.Fprintf(os.Stdout, "hello, %s\n", name)
	//fmt.Println(c)
} */
