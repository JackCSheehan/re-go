// Example usage
package main

import (
	"fmt"
	"github.com/jackcsheehan/re-go/rego"
)

func main() {
	// Text to find patterns in
	t := `https://www.google.comhttps://www.nasa.gov`

	// Compile regular expression
	//r, _ := rego.Compile(`(?P<protocol>https|http)://(?P<prefix>www).(?P<domain>\w+).(?P<toplevel>com|net|gov)`)
	r, _ := rego.Compile(`https`)

	// Demo of built-in functions
	fmt.Println(r.IsMatch(t))
	fmt.Println(r.FindAll(t))
	fmt.Println(r.FindAllGroups(t))
	fmt.Println(r.FindAllNamedGroups(t))
	fmt.Println(r.Matches(t))
	fmt.Println(r.Replace(t, "URL,"))
	fmt.Println(r.ReadValidate())
	fmt.Println(r.ReadSearch())
	fmt.Println(r.Split(t))
}