package glog

import (
	"os"
)

type stderr struct {
	file           *os.File
}

func NewStderr() *stderr {
	output := stderr{
		file:os.Stderr,
	}

	return &output
}

func (s *stderr) Write(msg []byte) error  {
	_,err := s.file.Write(msg)
	return err
}

func (s *stderr) Close()  {
	s.file.Close()
}