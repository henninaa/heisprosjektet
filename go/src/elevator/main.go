package main

import(
	"driver_module"
	"control_module"
	"os"
	"os/signal"
	)

func main(){
	
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	driver_module.Elev_init()
	
	go control_module.Elevator_main_control()

	<- c
	driver_module.Elev_stop_engine()
}

