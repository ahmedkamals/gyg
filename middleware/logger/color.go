package logger

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type (
	Colorable struct {
		output io.Writer
	}

	ColorValue int
)

const (
	ESCAPE = "\x1b"
)

const (
	RESET ColorValue = iota
	BOLD
	FAINT
	ITALIC
	BLINK_SLOW
	BLINK_RAPID
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground colors
const (
	BLACK ColorValue = iota + 30
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
)

// Background text colors
const (
	BG_BLACK ColorValue = iota + 40
	BG_RED
	BG_GREEN
	BG_YELLOW
	BG_BLUE
	BG_MAGENTA
	BG_CYAN
	BG_WHITE
)

func NewColorable(output io.Writer) *Colorable {
	return &Colorable{
		output: output,
	}
}

func (self *Colorable) Set(color ...ColorValue) *Colorable {
	fmt.Fprintf(self.output, self.format(color...))
	return self
}

func (self *Colorable) Unset() {
	fmt.Fprintf(self.output, "%s[%dm", ESCAPE, RESET)
}

func (self *Colorable) Wrap(string string, color ...ColorValue) string {
	return self.format(color...) + string + self.resetFormat()
}

func (self *Colorable) format(color ...ColorValue) string {
	return fmt.Sprintf("%s[%sm", ESCAPE, self.sequence(color...))
}

func (self *Colorable) resetFormat() string {
	return fmt.Sprintf("%s[%dm", ESCAPE, RESET)
}

func (self *Colorable) sequence(color ...ColorValue) string {
	format := make([]string, len(color))
	for key, value := range color {
		format[key] = strconv.Itoa(int(value))
	}

	return strings.Join(format, ";")
}
