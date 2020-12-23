package sender

import (
	. "CDMA/utils"
	"time"
)

func (s *Sender)getCode(bit int) []int{
	arr := make([]int, len(s.walshCode))
	//silent
	if bit==-1 {
		return arr
	}

	copy(arr, s.walshCode)

	if bit==0{
		bit = -1
	}
	for i := range arr{
		arr[i]*=bit
	}
	return arr
}

func (s * Sender)Send(){
	curBit := s.buffer[0]

	for len(s.buffer)!=0 {

		if SenseMedium() == CANWRITE {

				//sending the code
				s.S2C <- s.getCode(curBit)

				//wait for collision signal during propagation
				time.Sleep(500 * time.Microsecond)

				//transmission successful
				s.buffer = s.buffer[1:]

				if len(s.buffer) != 0 {
					curBit = s.buffer[0]
				}

			}else{
				time.Sleep(100* time.Microsecond)
			}

	}
}

