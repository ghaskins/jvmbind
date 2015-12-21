package jvmbind

import (
	"os/exec"
	"log"
	"io"
	"encoding/binary"
	"errors"
	"fmt"
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

func (self *Instance) Send(payload []byte) {
	stream := self.Stdin

	header := make([]byte, 4)
	binary.BigEndian.PutUint32(header, uint32(len(payload)))

	stream.Write(header)
	stream.Write(payload)
}

func (self *Instance) Recv(payload []byte) error {
	stream := self.Stdout

	header := make([]byte, 4)
	hLen, err := stream.Read(header)
	if err != nil || hLen != 4 {
		return err
	}

	iLen := binary.BigEndian.Uint32(header)
	if int(iLen) > len(payload) {
		return errors.New("MTU violation")
	}

	pLen, err := stream.Read(payload[:iLen])
	if err != nil {
		return err
	}
	if pLen != int(iLen) {
		return errors.New(fmt.Sprintf("Read error: expected %d bytes, got %d", iLen, pLen))
	}

	return nil
}


func (self* Instance) Wait() {
	self.cmd.Wait()
}

