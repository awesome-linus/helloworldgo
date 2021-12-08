package helloworldgo

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Kou"
	actual := Hello(name)

	want := fmt.Sprintf("Hello World! %s", name)
	if want != actual {
		t.Fatal("Error!")
	}
}
