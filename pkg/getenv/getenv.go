package getenv

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	URL string
}
func GetDBEnv() (*Config,error){
	var config Config
	err := envconfig.Process("DB",&config)
	if err != nil {
		fmt.Printf("error getting environment variables:%v",err)
		return nil,err
	}
	return &config,nil
}
func GetRabbitEnv()(*Config,error){
	var config Config
	err := envconfig.Process("RABBITMQ",&config)
	if err != nil {
		fmt.Printf("error getting environment variables:%v",err)
		return nil,err
	}
	return &config,nil
}
