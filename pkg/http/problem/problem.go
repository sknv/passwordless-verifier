package problem

import (
	"fmt"
)

const _defaultType = "about:blank"

// InvalidParam describes an invalid parameter.
type InvalidParam struct {
	Name     string   `json:"name"`
	Messages []string `json:"messages"`
}

type invalidParams struct {
	InvalidParams []InvalidParam `json:"invalidParams"`
}

// Problem defines a problem details object (see RFC-7807).
type Problem struct {
	// Type contains a URI that identifies the problem type. This URI will,
	// ideally, contain human-readable documentation for the problem when de-referenced.
	Type string `json:"type"`

	// Title is a short, human-readable summary of the problem type. This title
	// SHOULD NOT change from occurrence to occurrence of the problem, except for purposes of localization.
	Title string `json:"title"`

	// Status is an HTTP status code for this occurrence of the problem.
	Status int `json:"status"`

	// Detail is a human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail,omitempty"`

	// Data holds any additional information.
	Data any `json:"data,omitempty"`
}

func New(status int, title string) *Problem {
	return &Problem{
		Status: status,
		Title:  title,
		Type:   _defaultType,
	}
}

func (p *Problem) Error() string {
	return fmt.Sprintf("status = %d, title = %s, detail = %s", p.Status, p.Title, p.Detail)
}
