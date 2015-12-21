package jvmbind_test

import (
	"testing"
	"strings"
	"github.com/ghaskins/jvmbind"
)

func TestHello(t *testing.T) {
	instance, err := jvmbind.Launch("./test/hello-world.jar")
	if err != nil {
		t.Error(err)
	}

	_msg := make([]byte, 256)

	stdout := instance.Stdout
	stdout.Read(_msg)
	msg := string(_msg)

	if strings.Contains(msg, "Hello World!") == false {
		t.Errorf("Unexpected response: %s", msg)
	}

	instance.Wait()
}
