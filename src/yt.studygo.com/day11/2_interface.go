package main

import (
	"flag"
	"fmt"
)

type Celsius float32
type Fahrenheit float32

func FToc(f Fahrenheit) Celsius {
	return Celsius((f - 32.0) * 5.0 / 9.0)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g℃", c)
}

type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	// 在标记处理过程中根据输入的值更新celsius
	var unit string
	var value float32
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C":
		f.Celsius = Celsius(value)
		return nil
	case "F":
		f.Celsius = FToc(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage) // 返回一个flagset
	return &f.Celsius
}

/* func main() {
	var temp = CelsiusFlag("temp", 20.0, "the temperature")
	flag.Parse()
	fmt.Printf("%v \n", temp)
}
*/
