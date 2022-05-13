package problem

import (
	"fmt"
	"strings"
)

// InvalidParam describes an invalid parameter.
type InvalidParam struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

// Problem defines a problem details object.
type Problem struct {
	// Err holds a wrapped error. The field is not rendered.
	Err error `json:"-"`

	// Type contains a URI that identifies the problem type. This URI will,
	// ideally, contain human-readable documentation for the issue when de-referenced.
	Type string `json:"type,omitempty"`

	// Title is a short, human-readable summary of the problem type. This title
	// SHOULD NOT change from occurrence to occurrence of the issue, except for purposes of localization.
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
	}
}

func (p *Problem) Error() string {
	fields := []string{
		fmt.Sprintf("status = %d", p.Status),
		fmt.Sprintf("title = %s", p.Title),
	}

	if p.Type != "" {
		fields = append(fields, fmt.Sprintf("type = %s", p.Type))
	}

	if p.Err != nil {
		fields = append(fields, fmt.Sprintf("err = %s", p.Err))
	}

	return strings.Join(fields, ", ")
}

func (p *Problem) Unwrap() error {
	return p.Err
}
