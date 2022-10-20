package src

import (
	"errors"
	"strings"
)

func Throw(msg ...string) {
	var build strings.Builder
	build.WriteString("")

	for _, s := range msg {
		build.WriteString(s)
	}

	panic(errors.New(build.String()))
}
