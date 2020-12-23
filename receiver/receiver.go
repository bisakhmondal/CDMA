package receiver

import (
	"log"
)

type Receiver struct{
	C2R chan []int
	walshTable [][]int
	NUMSENDER int
}

func NewReceiver(c chan []int, table [][]int, nodes int) *Receiver{
	return &Receiver{
		C2R: c,
		walshTable: table,
		NUMSENDER: nodes,
	}
}

func (r * Receiver) Decode(bitarr []int){
	for i:= 0;i<r.NUMSENDER;i++{
		data :=0
		for j:= 0;j<len(bitarr);j++{
			data += r.walshTable[i][j]*bitarr[j]
		}
		data /= r.NUMSENDER

		if data==0{
			log.Println("Received from: Sender: ", i+1," | Silence")
		}else if data==1{
			log.Println("Received from: Sender: ", i+1," | Bit: 1")
		}else{
			log.Println("Received from: Sender: ", i+1," | Bit: 0")
		}
	}
}

func (r * Receiver)Init(){
	for {
		bytestream := <- r.C2R

		r.Decode(bytestream)
		log.Println("**********************************************************")
	}
}
