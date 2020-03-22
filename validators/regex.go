package validators

import "regexp"

type Regex struct {
	pattern string
}

func NewRegex(pattern string) *Regex {
	return &Regex{pattern: pattern}
}

func (r Regex) Validate(value string) bool {
	match, _ := regexp.MatchString(r.pattern, value)
	return match
}
