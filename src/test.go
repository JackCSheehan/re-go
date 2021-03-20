// Example usage
package main

import (
	"fmt"
	"github.com/jackcsheehan/re-go/rego"
)

func main() {
	t := `https://www.google.comhttps://www.nasa.gov`

	r := rego.Compile(`(?P<protocol>https|http)://(?P<prefix>www).(?P<domain>\w+).(?P<toplevel>com|net|gov)`)
	fmt.Println(r.IsMatch(t))
	fmt.Println(r.FindAll(t))
	fmt.Println(r.FindAllGroups(t))
	fmt.Println(r.FindAllNamedGroups(t))
	fmt.Println(r.Matches(t))
}