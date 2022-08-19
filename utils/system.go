package utils

var ConfigPath string = "local" // default win-env

// Get config path for local or docker
func GetConfigPath() string {
	if ConfigPath == "docker" {
		return "configs/config-docker"
	} else if ConfigPath == "local" {
		return "configs/config" // localhost
	} else if ConfigPath == "example" {
		return "configs/config-local"
	}
	return "config-local" // release - win-env
}
