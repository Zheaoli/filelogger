package filelogger

import (
	"testing"
)

func Test_DefaultLogger(t *testing.T) {
	Log.Info("Info")
	Log.WithFields(map[string]interface{}{
		"fielda": "a",
		"fieldb": "b",
	}).Info("Info fields")

	type StructA struct {
		A int `json:"a"`
		B int `json:"b"`
	}

	sa := StructA{
		A: 1,
		B: 2,
	}

	Log.Error(sa)
}

func Test_NewLogger(t *testing.T) {
	logger, err := NewLogger("./testdata", "test.log", "debug")
	if err != nil {
		t.Fail()
	}

	logger.Info("Info")
	logger.WithFields(map[string]interface{}{
		"fielda": "a",
		"fieldb": "b",
	}).Info("Info fields")

	type StructA struct {
		A int `json:"a"`
		B int `json:"b"`
	}

	sa := StructA{
		A: 1,
		B: 2,
	}

	logger.Error(sa)
}

