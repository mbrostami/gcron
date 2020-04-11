package validators

import "regexp"

// Regex struct for regex validation
type Regex struct {
	pattern string
}

// NewRegex make new regex struct
func NewRegex(pattern string) *Regex {
	return &Regex{pattern: pattern}
}

// Validate check input value with given regex
func (r *Regex) Validate(value string) bool {
	match, _ := regexp.MatchString(r.pattern, value)
	return match
}
