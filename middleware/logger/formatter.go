package logger

import (
	"strconv"
	"strings"
)

type (
	FormatterAware interface {
		format(context []Context) string
	}

	Formatter struct {
		FormatterAware
	}
)

func NewFormatter() *Formatter {
	return &Formatter{}
}

func (self *Formatter) format(context []Context) string {
	context = self.prepend(" ▶ ", context)
	output := []string{}
	for _, value := range context {
		switch value.(type) {
		case string:
			output = append(output, value.(string))
		case int:
			output = append(output, strconv.Itoa(value.(int)))
			break
		}
	}

	return strings.Join(output, " ")
}

func (self *Formatter) prepend(value string, context []Context) []Context {
	context = append(context, 0)
	copy(context[1:], context[0:])
	context[0] = value

	return context
}
