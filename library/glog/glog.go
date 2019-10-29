package glog

import (
	"sync"
	"reflect"
	"errors"
	"runtime"
	"time"
	"log"
	"fmt"
)

const (
	//日志级别
	LevelDebug   = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
)

const (
	Ldate         = 1 << iota // the date in the local time zone: 2009/01/23
	Ltime                     // the time in the local time zone: 01:23:23
	Lmicroseconds             // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                 // full file name and line number: /a/b/c/d.go:23
	Lshortfile                // final file name element and line number: d.go:23. overrides Llongfile
	LUTC
)

var ErrRegisterLoggerRepeat = errors.New("register logger repeated") //重复注册logger
var logger = &Glogger{
	level:   LevelDebug,
	flag:    Ldate | Ltime | Lshortfile,
	outputs: make(map[string]Output),
}
var levelPrefix = [...]string{"[debug]", "[info]", "[warning]", "[error]", "[fatal]"}

type Glogger struct {
	lock    sync.Mutex
	level   int
	flag    int
	outputs map[string]Output
}

type Output interface {
	Write([]byte) error
	Close()
}

func SetLevel(level int) {
	//这里不用对level范围做检测了
	logger.level = level
}

func SetFlag(flag int) {
	logger.flag = flag
}

func (l *Glogger) register(output Output) error {
	loggerType := reflect.TypeOf(output).String()
	if _, ok := l.outputs[loggerType]; ok {
		return ErrRegisterLoggerRepeat
	}

	l.lock.Lock()
	defer l.lock.Unlock()

	if _, ok := l.outputs[loggerType]; ok {
		return ErrRegisterLoggerRepeat
	}

	l.outputs[loggerType] = output
	return nil
}

func (l *Glogger) write(level int, msg string) {
	buf := l.formatHeader(level)
	buf = append(buf, msg...)
	if buf[len(buf)-1] != '\n' {
		buf = append(buf, '\n')
	}

	for outputType, output := range l.outputs {
		err := output.Write(buf)
		if err != nil {
			log.Printf("log write to output failed,output_type:%s,err:%s",outputType,err)
		}
	}
}

func (l *Glogger) formatHeader(level int) []byte {
	var prefix string
	if level < 0 || level >= len(levelPrefix) {
		prefix = "???"
	} else {
		prefix = levelPrefix[level]
	}

	buf := []byte(prefix)
	//日志输出时间
	if l.flag&(Ldate|Ltime|Lmicroseconds) > 0 {
		now := time.Now()
		if l.flag&LUTC != 0 {
			now = now.UTC()
		}

		year, month, day := now.Date()
		itoa(&buf, year, 4)
		buf = append(buf, '/')
		itoa(&buf, int(month), 2)
		buf = append(buf, '/')
		itoa(&buf, day, 2)
		buf = append(buf, ' ')

		if l.flag&(Ltime|Lmicroseconds) > 0 {
			hour, min, sec := now.Clock()
			itoa(&buf, hour, 2)
			buf = append(buf, ':')
			itoa(&buf, min, 2)
			buf = append(buf, ':')
			itoa(&buf, sec, 2)

			if l.flag&Lmicroseconds > 0 {
				buf = append(buf, '.')
				itoa(&buf, now.Nanosecond()/1e3, 6)
			}
			buf = append(buf, ' ')
		}
	}

	//文件名 行号
	if l.flag&(Lshortfile|Llongfile) > 0 {
		callDepth := 3
		_, file, line, ok := runtime.Caller(callDepth)
		if !ok {
			file = "???"
			line = 0
		}

		if l.flag&Lshortfile > 0 {
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					file = file[i+1:]
					break
				}
			}
		}

		buf = append(buf, file...)
		buf = append(buf, ':')
		itoa(&buf, line, -1)
	}

	buf = append(buf, ": "...)
	return buf
}

/**
整型转化为字符串，写入buf，如果wid超出i的本来长度，多出的长度在前面补0。
buf 写入目标字节数组
i 要转行的整型
wid 转化后的字符串宽度
 */
func itoa(buf *[]byte, i int, wid int) {
	// Assemble decimal in reverse order.
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10) //'0'为字符串"0"的字节编码，对应48
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}

func RegisterOutput(output Output) error {
	return logger.register(output)
}

func Debug(v...interface{}) {
	logger.write(LevelDebug,fmt.Sprintln(v...))
}

func Debugf(format string,v...interface{}) {
	logger.write(LevelDebug,fmt.Sprintf(format,v...))
}

func Info(v...interface{}) {
	logger.write(LevelInfo,fmt.Sprintln(v...))
}

func Infof(format string,v...interface{}) {
	logger.write(LevelInfo,fmt.Sprintf(format,v...))
}

func Warning(v...interface{}) {
	logger.write(LevelWarning,fmt.Sprintln(v...))
}

func Warningf(format string,v...interface{}) {
	logger.write(LevelWarning,fmt.Sprintf(format,v...))
}

func Error(v...interface{}) {
	logger.write(LevelError,fmt.Sprintln(v...))
}

func Errorf(format string,v...interface{}) {
	logger.write(LevelError,fmt.Sprintf(format,v...))
}

func Fatal(v...interface{}) {
	logger.write(LevelFatal,fmt.Sprintln(v...))
}

func Fatalf(format string,v...interface{}) {
	logger.write(LevelFatal,fmt.Sprintf(format,v...))
}