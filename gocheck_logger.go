package gomolgocheck

import (
	"time"

	"github.com/aphistic/gomol"
	check "gopkg.in/check.v1"
)

type GocheckLogger struct {
	base *gomol.Base
	c    *check.C
}

func NewGocheckLogger() *GocheckLogger {
	return &GocheckLogger{}
}

func (gl *GocheckLogger) SetC(c *check.C) {
	gl.c = c
}

func (gl *GocheckLogger) SetBase(base *gomol.Base) {
	gl.base = base
}

func (gl *GocheckLogger) InitLogger() error {
	return nil
}

func (gl *GocheckLogger) ShutdownLogger() error {
	return nil
}

func (gl *GocheckLogger) IsInitialized() bool {
	return true
}

func (gl *GocheckLogger) Logm(ts time.Time, level gomol.LogLevel, attrs map[string]interface{}, msg string) error {
	if gl.c != nil {
		gl.c.Log(msg)
	}
	return nil
}
