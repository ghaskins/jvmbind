package jvmbind

import (
	"os/exec"
	"log"
	"io"
)

type Instance struct {
	cmd* exec.Cmd
	Stdout io.ReadCloser
	Stdin io.WriteCloser
}

func Launch(jar string) (*Instance, error)  {
	cmd := exec.Command("java", "-jar", jar)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	instance := &Instance{cmd: cmd, Stdout: stdout, Stdin: stdin}

	return instance, nil
}

func (self* Instance) Wait() {
	self.cmd.Wait()
}

