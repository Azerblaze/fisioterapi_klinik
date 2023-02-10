package main

import (
	"projek_fisioterapi/configs"
	"projek_fisioterapi/routes"
	v1 "projek_fisioterapi/routes/v1"
)

func main() {

	configs.InitConfig()
	configs.InitDatabase()

	routePayload := &routes.Payload{
		DBGorm: configs.DB,
		Config: configs.Cfg,
	}

	routePayload.InitUserService()

	e, trace := v1.InitRoute(routePayload)
	defer trace.Close()

	e.Logger.Fatal(e.Start(configs.Cfg.APIPort))
}
