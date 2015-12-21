package jvmbind_test

import (
	"testing"
	"github.com/ghaskins/jvmbind"
)

func TestLaunch(t *testing.T) {
	instance, err := jvmbind.Launch("./test/helloworld.jar")
	if err != nil {
		t.Error(err)
	}

	instance.Wait()
}
