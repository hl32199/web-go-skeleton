package glog

import (
	"os"
)

type stdout struct {
	file           *os.File
}

func NewStdout() *stderr {
	output := stderr{
		file:os.Stdout,
	}

	return &output
}

func (s *stdout) Write(msg []byte) error  {
	_,err := s.file.Write(msg)
	return err
}

func (s *stdout) Close()  {
	s.file.Close()
}