package channel

import (
	. "CDMA/utils"
	"log"
	"sync"
	"time"
)


type Channel struct {
	C2S chan [] int
	C2R chan [] int
	Buffer []int
	monitor * StatsMonitor
}

func NewChannel(chc2s, chc2r chan [] int, Monitor *StatsMonitor) * Channel{
	return &Channel{
		C2S: chc2s,
		C2R: chc2r,
		Buffer: []int{},
		monitor: Monitor,
	}
}


func (c * Channel)Init(){
	log.Println("Channel Initialized")
	var wgi sync.WaitGroup
	wgi.Add(1)
	//go Teller()
	go c.send2receiver(&wgi)
	c.receivefromSender()

}
var lock sync.Mutex
func (c *Channel)CompressBuff(buffer [][]int, length int) {

	if length==0{
		return
	}
	newbuf := make([]int, len(buffer[0]))
	for j:=0;j<len(buffer[0]);j++ {
		for i := 0; i < length; i++ {
			newbuf[j] += buffer[i][j]
		}
	}

	lock.Lock()
	c.Buffer = newbuf
	lock.Unlock()
}
func (c* Channel)receivefromSender(){
	for {
		//make the collision flag unset
		//SetNotCollided()

		if SenseMedium() == CANWRITE {
			//if more than one frames in the pipe then collision
			var buffer [][]int
			cnt := 0
			if len(c.C2S) == 0 {
				time.Sleep(150 * time.Microsecond)
			}

			if len(c.C2S) != 0 {
				//collision
				time.Sleep(150 * time.Microsecond)
				SetMediumNWRABLE()

				//log into monitor for throughput and efficiency
				c.monitor.Transmitted(len(c.C2S))

				c.monitor.Success()
				//free the pipe, data is unusable
				for len(c.C2S) > 0 {
					code := <-c.C2S
					buffer = append(buffer, code)
					cnt++
				}

			}
			c.CompressBuff(buffer, cnt)
			//compression takes time
			time.Sleep(300* time.Microsecond)
		}

		SetMediumWRABLE()

	}
}

func (c *Channel)send2receiver(wg  *sync.WaitGroup){
	defer wg.Done()
	for {

		if len(c.Buffer)==0{
			time.Sleep(100*time.Microsecond)
			continue
		}
		lock.Lock()
		bytestream := c.Buffer
		c.Buffer = []int{}
		c.C2R <- bytestream
		lock.Unlock()
	}
	//log.Println("exited from s2r")

}