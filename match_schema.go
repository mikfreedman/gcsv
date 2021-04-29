package gcsv

import (
	"encoding/csv"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/onsi/gomega/types"
)

const errorMsg = "(field at position %d of row %d should have been of type %s)"

type RepresentSchemaOption func(*representSchemaMatcher)

func IgnoreHeaderRow() RepresentSchemaOption {
	return func(matcher *representSchemaMatcher) {
		matcher.ignoreHeaderRow = true
	}
}

func WithHeaders(headers ...string) RepresentSchemaOption {
	return func(matcher *representSchemaMatcher) {
		matcher.withHeaders = headers
	}
}

func RepresentSchema(expected []interface{}, opts ...RepresentSchemaOption) types.GomegaMatcher {
	m := &representSchemaMatcher{
		expected: expected,
	}

	for _, v := range opts {
		v(m)
	}

	return m
}

type representSchemaMatcher struct {
	expected        []interface{}
	ignoreHeaderRow bool
	withHeaders     []string
	lastError       string
}

func (matcher *representSchemaMatcher) Match(actual interface{}) (success bool, err error) {
	response, ok := actual.(string)
	if !ok {
		return false, fmt.Errorf("representSchema matcher expects a string, got %T instead", actual)
	}

	lines, err := csv.NewReader(strings.NewReader(response)).ReadAll()

	if err != nil {
		return false, err
	}

	for r, v := range lines {
		if r == 0 {
			switch {
			case matcher.ignoreHeaderRow:
				continue
			case len(matcher.withHeaders) > 0:
				if reflect.DeepEqual(v, matcher.withHeaders) {
					continue
				}

				matcher.lastError = "headers do not match"
				return false, nil
			}
		}

		for i, e := range matcher.expected {
			switch e.(type) {
			case string:
			case int:
				_, err := strconv.Atoi(v[i])
				if err != nil {
					matcher.lastError = fmt.Sprintf(errorMsg, i+1, r+1, "int")
					return false, nil
				}
			case bool:
				_, err := strconv.ParseBool(v[i])
				if err != nil {
					matcher.lastError = fmt.Sprintf(errorMsg, i+1, r+1, "bool")
					return false, nil
				}
			case float64:
				_, err := strconv.ParseFloat(v[i], 64)
				if err != nil {
					matcher.lastError = fmt.Sprintf(errorMsg, i+1, r+1, "float64")
					return false, nil
				}
			default:
				return false, fmt.Errorf("type of field at position %d of actual is unable to be handled", i)
			}
		}
	}
	return true, nil
}

func (matcher *representSchemaMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nto contain a CSV matching the schema of\n\t%#v\n%s", actual, matcher.expected, matcher.lastError)
}

func (matcher *representSchemaMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nnot to contain a CSV matching the schema of\n\t%#v\n%s", actual, matcher.expected, matcher.lastError)
}
