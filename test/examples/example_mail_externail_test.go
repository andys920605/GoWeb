package examples_test

import (
	"GoWeb/infras/configs"
	ext "GoWeb/repository/externals"
	rep_interface "GoWeb/repository/interface"
	"GoWeb/utils"
	"fmt"
	"log"
)

var (
	mailExt rep_interface.IMailRep
)

func init() {
	utils.ConfigPath = "example"
	buffer, err := configs.LoadConfig(utils.GetConfigPath())
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	cfgTemp, err := configs.ParseConfig(buffer)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}
	mailExt = ext.NewMailRep(cfgTemp)
}

func ExampleMailExternal_Send() {
	ok := mailExt.Send("test123")
	fmt.Println(ok)
	// Output:
	// true
}
