package rego

import (
	"regexp"
)

// Struct that allows interfacing with Regexp struct
type Rego struct {
	r	*regexp.Regexp
}

// Contains data from a match
type Match struct {
	Source		string				// String searched using regex
	Groups		[]string			// Slice of capture groups found
	NamedGroups	map[string]string	// Maps capture group names -> group contents	
	Start		map[int]int			// Maps group index -> start index of that group
	End			map[int]int			// Maps group index -> end index of that group
}

func Compile(t string) Rego {
	return Rego{regexp.MustCompile(t)}
}

// Finds all capture groups and returns them as a string slice 
func (r Rego) FindAllGroups(t string) []string {
	return nil
}

// Finds all matches in given text and returns matches as string slice
func (r Rego) FindAll(t string) []string {
	return nil
}

// Returns true if there is a match in the given text
func (r Rego) IsMatch(t string) bool {
	return false
}

// Replaces source string found by regex with given replacement string
func (r Rego) Replace(src, repl string) error {
	return nil
}

// Returns new Match struct populated with data gatherd from given string
func (r Rego) Match(t string) Match {
	return Match{"", nil, nil, nil, nil}
}

// Tries to read given regex from stdin. Returns input and whether or not input matches regex
func (r Rego) ReadValidate() (string, bool) {
	return "", false
}

// Reads any input from stdin. Returns input and Match struct made from input
func (r Rego) ReadSearch() (string, Match) {
	return "", Match{"", nil, nil, nil, nil}
}