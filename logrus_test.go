package xlog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"os"
	"sync"
	"testing"
)

func TestLogrus(t *testing.T) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Debugf("debug log")
	logrus.Infof("info log")
	logrus.Errorf("error log")
	p := &parent{}
	p.Test()

	fmt.Println("=========================")

	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Debugf("debug log")
	logrus.Infof("info log")
	logrus.Errorf("error log")
	p = &parent{}
	p.Test()

	SetUp(true)
	fmt.Println("=========================")

	logrus.Debugf("debug log")
	logrus.Infof("info log")
	logrus.Errorf("error log")

	p.Test()
}

type parent struct {
}

func (p *parent) Test() {
	logrus.Debugf("debug log")
	logrus.Infof("info log")
	logrus.Errorf("error log")
}

func TestLogrusInfoLevel(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	logrus.Debugf("debug log")
	logrus.Infof("info log")
	logrus.Errorf("error log")

	SetUp(true)

	logrus.Debugf("debug log")
	logrus.Infof("info log")
	logrus.Errorf("error log")

}

func TestLogrusWriteFile(t *testing.T) {
	// You could set this to any `io.Writer` such as a file
	filename := "logrus.log"
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		assert.NoErrorf(t, err, "Failed to log to file, using default stderr")
	}
	logrus.SetOutput(file)

	defer func() {
		err = os.Remove(filename)
		assert.NoErrorf(t, err, "Failed to remove file: %s", filename)
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		TestLogrus(t)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()

		TestLogrusInfoLevel(t)
	}()
	wg.Wait()

	content, err := os.ReadFile(filename)
	assert.NoErrorf(t, err, "Failed to read file: %s", filename)
	assert.Containsf(t, string(content), "debug log", "Failed to read file: %s", filename)
	assert.Containsf(t, string(content), "info log", "Failed to read file: %s", filename)
	assert.Containsf(t, string(content), "error log", "Failed to read file: %s", filename)
}

func TestLogrusWriteFile2(t *testing.T) {
	log := logrus.New()
	filename := "logrus.log"
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		assert.NoErrorf(t, err, "Failed to log to file, using default stderr")
	}
	log.SetOutput(file)

	defer func() {
		err = os.Remove(filename)
		assert.NoErrorf(t, err, "Failed to remove file: %s", filename)
	}()

	log.Debugf("debug log")
	log.Infof("info log")
	log.Errorf("error log")

	SetUpWithLogger(log, true)

	log.Debugf("debug log")
	log.Infof("info log")
	log.Errorf("error log")
}
