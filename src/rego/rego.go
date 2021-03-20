package rego

import (
	"regexp"
)

// Struct that allows interfacing with Regexp struct
type Rego struct {
	regex	*regexp.Regexp
}

// Contains data from a match
type Match struct {
	Source		string				// String searched using regex
	Groups		[]string			// Slice of capture groups found
	NamedGroups	map[string]string	// Maps capture group names -> group contents	
	Start		map[int]int			// Maps group index -> start index of that group
	End			map[int]int			// Maps group index -> end index of that group
}

// Constructs new Rego struct instance
func Compile(t string) Rego {
	return Rego{regexp.MustCompile(t)}
}

// Finds all capture groups and returns them as a string slice 
func (r Rego) FindAllGroups(t string) []string {
	// Find capture groups using built-in submatch function
	submatches := r.regex.FindAllStringSubmatch(t, -1)

	// Slice to hold groups taken from submatches found using built-in submatch function
	var groups []string

	// Iterate through submatches and parse capture groups from them
	for _, slice := range(submatches) {
		groups = append(groups, slice[1])
	}

	return groups
}

// Finds all capture groups and maps them to their names
func (r Rego) FindAllNamedGroups(t string) []map[string]string {
	// Find capture groups using built-in submatch function
	submatches := r.regex.FindAllStringSubmatch(t, -1)

	// Get names of capture groups
	names := r.regex.SubexpNames()

	// Map of capture group names -> capture groups
	var namedGroups []map[string]string

	// Index of names. Starts at 1 since 0 is always blank
	index := 1

	// Map for the current match being parsed
	currentMatch := make(map[string]string)

	// Iterate through matches to map the names to the text
	for _, slice := range(submatches) {
		// Add current name/slice pair to current match map
		currentMatch[names[index]] = slice[1]
		index++

		// When end of names has been reached
		if index == len(names) {
			// Reset index
			index = 1

			// Add current match map to named group slice
			namedGroups = append(namedGroups, currentMatch)

			// Reset current match
			currentMatch = make(map[string]string)
		}
	}

	return namedGroups
}

// Finds all matches in given text and returns matches as string slice
func (r Rego) FindAll(t string) []string {
	return r.regex.FindAllString(t, -1)
}

// Returns true if there is a match in the given text
func (r Rego) IsMatch(t string) bool {
	return r.regex.Match([]byte(t))
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