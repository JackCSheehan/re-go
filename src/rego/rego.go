package rego

import (
	"regexp"
	//"fmt"
	"bufio"
	"os"
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
	Start		int					// Start index of this match
	End			int					// End index of this match
}

// Constructs new Rego struct instance
func Compile(t string) Rego {
	return Rego{regexp.MustCompile(t)}
}

// Finds all capture groups and returns them as a string slice 
func (r Rego) FindAllGroups(t string) [][]string {
	// Find capture groups using built-in submatch function
	submatches := r.regex.FindAllStringSubmatch(t, -1)

	// Slice to hold groups taken from submatches found using built-in submatch function
	var groups [][]string

	// Iterate through submatches and parse capture groups from them
	for _, slice := range(submatches) {
		groups = append(groups, slice[1:])
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

	// Index of matches and corresponding name
	index := 1

	// Map for the current match being parsed
	currentMatch := make(map[string]string)

	// Iterate through matches to map the names to the text
	for _, slice := range(submatches) {
		// Iterate through slice
		for index < len(slice) {
			// Map group name to corresponding capture group
			currentMatch[names[index]] = slice[index]
			index++
		}

		// When end of names/matches has been reached
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
func (r Rego) Replace(src, repl string) string {
	return r.regex.ReplaceAllString(src, repl)
}

// Returns new slice of Match struct populated with data gatherd from given string
func (r Rego) Matches(t string) []Match {
	// Slice of matches parsed from given text
	var matches []Match

	// Find data needed for matches slice
	groups := r.FindAllGroups(t)
	namedGroups := r.FindAllNamedGroups(t)
	indices := r.regex.FindAllStringIndex(t, -1)

	// Index to iterate through above collected data slices
	index := 0

	// Iterate through data collected above and create match instances
	for index < len(namedGroups) {
		matches = append(matches,Match {t, groups[index], namedGroups[index], indices[index][0], indices[index][1]})
		index++
	}

	return matches
}

// Tries to read given regex from stdin. Returns input and whether or not input matches regex
func (r Rego) ReadValidate() (string, bool) {
	// Create reader to read from console
	reader := bufio.NewReader(os.Stdin)

	// Get user input
	input, _ := reader.ReadString('\n')

	// Return true if the user's input matches regex
	if r.IsMatch(input) {
		return input, true
	}

	return input, false
}

// Reads any input from stdin. Returns input and slice of Match structs made from input
func (r Rego) ReadSearch() (string, []Match) {
	// Create reader to read from console
	reader := bufio.NewReader(os.Stdin)

	// Get user input
	input, _ := reader.ReadString('\n')

	// Get list of matches from input
	matches := r.Matches(input)

	return input, matches
}