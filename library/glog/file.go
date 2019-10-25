package glog

import (
	"os"
	"sync"
	"time"
	"path/filepath"
	"log"
	"fmt"
)

const DateFormat = "2006-01-02"
const FileExt = ".log"
const perm = 0744

type fileOutput struct {
	lock           sync.Mutex
	path           string
	fileNamePrefix string
	date           string
	file           *os.File
	//curFileName    string
}

type FileOutputConfig struct {
	Path           string
	FileNamePrefix string
}

func NewFileOutput(config FileOutputConfig) *fileOutput {
	fileOutput := fileOutput{
		path:           config.Path,
		fileNamePrefix: config.FileNamePrefix,
	}

	return &fileOutput
}

func (f *fileOutput) Write(msg []byte) error {
	f.lock.Lock()
	defer f.lock.Unlock()

	if f.file == nil {
		f.newFile()
	}

	if time.Now().Format(DateFormat) != f.date {
		err := f.rotate()
		if err != nil {
			log.Print(err)
		}
	}

	_,err := f.file.Write(msg)
	if err != nil {
		return fmt.Errorf("write log to file failed:%s",err)
	}
	return nil
}

func (f *fileOutput) getFileName() string {
	return filepath.Join(f.path, f.fileNamePrefix+FileExt)
}

func (f *fileOutput) rotate() error {
	//f.lock.Lock()
	//defer f.lock.Unlock()

	f.file.Close()
	fileName := f.getFileName()
	oldFileName := filepath.Join(f.path, f.fileNamePrefix+"-"+f.date+FileExt)
	os.Rename(fileName,oldFileName)

	err := f.newFile()
	if err != nil{
		return fmt.Errorf("rotate failed:%s",err)
	}
	return nil
}


func (f *fileOutput) newFile() error {
	err := os.MkdirAll(f.path, perm)
	if err != nil {
		return fmt.Errorf("can't make directories：%s", err)
	}
	fileName := f.getFileName()
	logfile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, perm)
	if err != nil {
		return fmt.Errorf("can't open file %s：%s",fileName, err)
	}

	f.file = logfile
	f.date = time.Now().Format(DateFormat)
	return nil
}

func (f *fileOutput) Close()  {
	f.file.Close()
}
