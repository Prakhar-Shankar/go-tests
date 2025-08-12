package dependencyinjection

import (
	"bytes"
	"fmt"
)

//left side - functions
func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer,"Hello, %s", name)
}