package core

import (
	"testing"
)

func TestGormDBInit(t *testing.T) {
	ViperInit()
	ZapLoggerInit()
	GormDBInit()
}
