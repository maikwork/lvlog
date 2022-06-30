package logger

import (
	"flag"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func NewLog() *logrus.Logger {
	wf := flag.Bool("wf", false, "write to file")
	l := flag.String("l", "ALL", "logging level")
	flag.Parse()

	log := logrus.New()

	if *l == "ALL" && *wf {
		log.Debug("For the flag -l the flag is not -wf not implemented")
		return log
	}

	f, lvl, err := getOptionLog(*wf, *l)
	if err != nil {
		log.Debug(err)
	}

	log.SetOutput(f)

	if *wf {
		log.SetLevel(lvl)
	}

	return log
}

func getOptionLog(y bool, name string) (io.Writer, logrus.Level, error) {
	var f io.Writer
	var err error
	var lvl logrus.Level

	f, err = getFileLog(y, "all")
	switch name {
	case "DEBUG":
		f, err = getFileLog(y, "debug")
		lvl = logrus.DebugLevel
	case "ERROR":
		f, err = getFileLog(y, "error")
		lvl = logrus.ErrorLevel
	case "INFO":
		f, err = getFileLog(y, "info")
		lvl = logrus.InfoLevel
	case "WARNING":
		f, err = getFileLog(y, "warning")
		lvl = logrus.WarnLevel
	}

	return f, lvl, err
}

func getFileLog(y bool, name string) (io.Writer, error) {
	var f io.Writer
	var err error

	if name == "error" {
		f = os.Stderr
	} else {
		f = os.Stdout
	}

	if y && name != "all" {
		f, err = os.OpenFile("./log/"+name+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logrus.Debug(err)
			return nil, err
		}
	}

	return f, nil
}
