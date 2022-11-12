package api_test

import (
	"fmt"
	"os"
	"task_list/driver"
	"task_list/route"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	routes *gin.Engine
	orm    *gorm.DB
)

func setup() {
	routes = route.Init()
	gin.SetMode(gin.TestMode)
	orm = driver.InitGorm()
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}

func TestMain(m *testing.M) {
	setup()
	fmt.Println("Test begins....")
	code := m.Run()
	teardown()
	os.Exit(code)
}
