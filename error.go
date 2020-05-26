package gen

import "strings"

type MultiError struct {
	errors []error
}

func NewMultiError(errors ...error) *MultiError {
	return &MultiError{errors}
}

func (m *MultiError) Add(err error) bool {
	if err != nil {
		m.errors = append(m.errors, err)
	}
	return err != nil
}

func (m *MultiError) Error() string {
	result := make([]string, len(m.errors))
	for i, e := range m.errors {
		result[i] = e.Error()
	}
	return strings.Join(result, "\n")
}

func (m *MultiError) Errors() []error {
	return m.errors
}
