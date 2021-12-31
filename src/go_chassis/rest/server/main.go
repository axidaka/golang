package main

import (
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/openlog"
	"golang/src/go_chassis/myschemas"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/rest/server/

func main() {
	chassis.RegisterSchema("rest", &myschemas.RestFulHello{})
	if err := chassis.Init(); err != nil {
		openlog.Fatal("Init failed." + err.Error())
		return
	}
	chassis.Run()
}
