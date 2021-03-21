# Re-go
A streamlined interface for Golang's `regexp` package. Re-go is based on Python's `re` module.

# Struct Reference
Struct Name|Description|Fields
-|-|-
`Rego`|A wrapper around Golang's built-in `Regexp` struct| `regex *regexp.Regexp`
`Match`|Contains information about a specific regex match. Based on Python's `Match` class from the `re` module.|`source string`,<br>`Groups []string`,<br>`NamedGroups map[string]string`,<br>`Start int`,<br>`End int`,<br>`Span [2]int`

# Function Reference
*WIP*
Function Name|Description
-|-