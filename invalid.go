package checron

import "fmt"

type Invalid struct {
	raw  string
	line int
}

func (er *Invalid) Type() Type {
	return TypeInvalid
}

func (er *Invalid) Err() error {
	return fmt.Errorf("invalid line: %q", er.raw)
}

func (er *Invalid) Raw() string {
	return er.raw
}

func (er *Invalid) Line() int {
	return er.line
}