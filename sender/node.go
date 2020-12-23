package sender

import (
	"log"
	"math/rand"
	"sync"
)

//maximumm collision before a frame drops
const MAX_COLLISIONS = 12
const TIMESLOTS = 500 //at least 500 milliseconds

type Sender struct {
	buffer [] int
	S2C chan <- []int
	collisions int
	NodeNo int
	srcMac int
	destMac int
	walshCode []int
}

func NewSenderNode(nodeno int, destMac int, chc2s chan <- []int, code []int) *Sender{
	return &Sender{
		buffer: []int{},
		collisions: 0,
		S2C : chc2s,
		NodeNo: nodeno,
		destMac: destMac,
		srcMac: rand.Intn(1<<7),
		walshCode: code,
	}
}

func (s * Sender)Init(MAXSIMULATETIME, Arrival_rate int, wg *sync.WaitGroup){
	//notify about the sender thread completion
	defer wg.Done()
	//generate frames
	s.GenerateFrames(MAXSIMULATETIME, Arrival_rate)
	log.Println("Frame gen", s.buffer)
	s.Send()
	log.Println("Exiting sender", s.NodeNo)

}

func (s* Sender)GenerateFrames(MAXSIMULATETIME, ARRIVAL_RATE int){
	timeNow :=0.0
	for timeNow < float64(MAXSIMULATETIME) {
		//Time penalty for frame generation
		timeNow += rand.Float64() / float64(ARRIVAL_RATE)
		//appending the frames
		data := rand.Intn(1<<15) & 1

		//idle
		if rand.Float64() < 0.2{
			data = -1
		}
		//genFrame := MakeFrame(s.srcMac, s.destMac, data, s.NodeNo)
		s.buffer = append(s.buffer, data)
	}
}