# Re-go
A streamlined interface for Golang's `regexp` package. Re-go is based on Python's `re` module.

# Struct Reference
Struct Name|Description|Fields
-|-|-
`Rego`|A wrapper around Golang's built-in `Regexp` struct| `regex *regexp.Regexp`
`Match`|Contains information about a specific regex match. Based on Python's `Match` class from the `re` module.|`source string`,<br>`Groups []string`,<br>`NamedGroups map[string]string`,<br>`Start int`,<br>`End int`,<br>`Span [2]int`

# Function Reference
Function Name|Description
-|-
`Compile(t string) (Rego, error)`|Constructs a new Rego struct instance. Corresponds directly to `Regexp.Compile`.
`(r Rego) FindAllGroups(t string) [][]string`|Finds all capture groups and returns them as a slice of string slices where each slice contains the groups found in each regex match.
`(r Rego) FindAllNamedGroups(t string) []map[string]string`|Finds all named capture groups and maps their names to them.
`(r Rego) FindAll(t string) []string`|Finds all matches in given text and returns matches as a string slice. Corresponds to `Regexp.FindAllString`.
`(r Rego) HasMatch(t string) bool`|Returns true if there is a match in the given text. Corresponds to `Regexp.Match`.
`(r Rego) Replace(src, repl string) string`|Replaces source string found by regex with given replacement string. Corresponds to `Regexp.ReplaceAllString()`.
`(r Rego) Split(t string) []string`|Splits given string using regex and returns that split string. Corresponds to `Regexp.Split`.
`(r Rego) Matches(t string) []Match`|Returns new slice of `Match` structs populated with data gathered from given string using regex. Based on Python's `re.findall` function.
`(r Rego) ReadValidate() (string, bool)`|Reads input from stdin. Returns read string and whether or not the read string contains a regex match.
`(r Rego) ReadSearch() (string, []Match)`|Reads input from stdin and returns both the read string and a slice of `Match` structs for every regex match found in input.

