package queue_module

import(
	"driver_module"
	)


type Queue_t struct(

	var queue [12]queue_post
	var order_lights [4][3]bool
	var backup [4]queue_backup

	)

func (queue * Queue_t) init(){

	var j driver_module.Elev_button_type_t

	queue.queue = Init_queue()

	for i :=0; i<driver_module.N_FLOORS; i++{
		for j = 0; j<driver_module.N_BUTTONS; j++{
			order_lights[i][j] = false
			driver_module.Elev_set_button_lamp(j,i,0)
		}
	}
}

type queue_post struct{

	var floor int
	var button_type driver_module.Elev_button_type_t

}

type queue_backup struct{

	var ip string
	var queue [12]queue_post

}

func (backup * queue_backup) init(ip string){

	backup.queue = Init_queue()
	backup.ip = ip
}