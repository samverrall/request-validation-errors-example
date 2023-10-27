package errsx

import (
	"errors"
	"fmt"
	"strings"
)

// Map represents a collection of errors keyed by name.
type Map map[string]error

// Get will return the error string for the given key.
func (m Map) Get(key string) string {
	if err := m[key]; err != nil {
		return err.Error()
	}

	return ""
}

func (m *Map) Has(key string) bool {
	_, ok := (*m)[key]

	return ok
}

// Set associates the given error with the given key.
// The map is lazily instantiated if it is nil.
func (m *Map) Set(key string, msg any) {
	if *m == nil {
		*m = make(Map)
	}

	var err error
	switch msg := msg.(type) {
	case error:
		if msg == nil {
			return
		}

		err = msg

	case string:
		err = errors.New(msg)

	default:
		panic("want error or string message")
	}

	(*m)[key] = err
}

func (m Map) Error() string {
	if m == nil {
		return "<nil>"
	}

	pairs := make([]string, len(m))
	i := 0
	for key, err := range m {
		pairs[i] = fmt.Sprintf("%v: %v", key, err)

		i++
	}

	return strings.Join(pairs, "; ")
}

func (m Map) String() string {
	return m.Error()
}

// MarshalJSON implements the json.Marshaler interface.
func (m Map) MarshalJSON() ([]byte, error) {
	errs := make([]string, 0, len(m))
	for key, err := range m {
		errs = append(errs, fmt.Sprintf("%q:%q", key, err.Error()))
	}

	return []byte(fmt.Sprintf("{%v}", strings.Join(errs, ", "))), nil
}
