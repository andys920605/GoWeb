package examples_test

import (
	"GoWeb/infras/configs"
	model_com "GoWeb/models/commons"
	models_ext "GoWeb/models/externals"
	ext "GoWeb/repository/externals"
	rep_interface "GoWeb/repository/interface"
	"GoWeb/utils"
	"fmt"
	"log"
)

var (
	mailExt rep_interface.IMailExt
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
	options := &model_com.Options{
		Config: cfgTemp,
	}
	mailExt = ext.NewMailExt(options)
}

func ExampleMailExt_Send() {
	ok := mailExt.Send(createEmail())
	fmt.Println(ok)
	// Output:
	// true
}

// region private function
func createEmail() *models_ext.SendMail {
	return &models_ext.SendMail{
		TargetAddress: "andys920605@gmail.com",
		Title:         "Example.Test",
		Body:          "Success",
	}
}

// endregion
