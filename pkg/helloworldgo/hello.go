package helloworldgo

import (
	"fmt"
)

func Hello(name string) string {
	return fmt.Sprintf("Hello World! %s", name)
}
