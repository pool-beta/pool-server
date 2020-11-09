package pool

import (

)

type Tank interface {
	Pull(Drop)
}

type tank struct {
	
}