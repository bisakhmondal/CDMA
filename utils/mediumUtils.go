package utils

const (
	CANWRITE = 1
	CANTWRITE = 0
)
var addcount = make(chan int) // set current medium status
var counter = make(chan int) //  get current medium status

func SetMediumWRABLE() { addcount <- CANWRITE }
func SetMediumNWRABLE() { addcount <- CANTWRITE }

func SenseMedium() int { return <-counter }


func TellerMedium() {
	var curStatus int = CANTWRITE // Status of medium is confined to only Teller goroutine
	for {
		select {
		case curStatus = <-addcount:
		case counter <- curStatus:
		}
	}
}
