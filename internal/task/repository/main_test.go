package repository

import (
	"os"
	"task_list/config"
	"testing"
)

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	config.InitConfig()
}
