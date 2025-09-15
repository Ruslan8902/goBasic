package api

import (
	"fmt"
	"gobasics/config"
)

func f() {
	configStuct := config.NewConfig()
	fmt.Println(configStuct)
}
