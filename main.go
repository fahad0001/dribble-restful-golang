package main

import (
	"gallery-app/configs"
	"gallery-app/controller"
	"fmt"
)

func main() {
	if err := configs.LoadConfig("./configs"); err != nil {
		panic(err)
	}
	controller.InitializeApp(configs.Config.DataSource)
	controller.GetDribbleData(configs.Config.DataSource, configs.Config.ImageStorePath, configs.Config.ClientAccessToken)

	// attach api with application
	fmt.Println("Server is Listening to API Request!")
	controller.APIBuilder(configs.Config.DataSource)
}