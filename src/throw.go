package src

import (
	"errors"
	"strings"
)

func throw(msg ...string) {
	var build strings.Builder
	build.WriteString("")

	for _, s := range msg {
		build.WriteString(s)
	}

	panic(errors.New(build.String()))
}
