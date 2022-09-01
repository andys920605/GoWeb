package commons

import (
	"GoWeb/infras/configs"
	infras_itf "GoWeb/infras/interface"
)

type Options struct {
	// fixed config data
	Config *configs.Config
	// logger
	Logger infras_itf.IApiLogger
}
