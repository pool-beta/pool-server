package simple

import (

)

/* 
	Simple contains the simple api for working with pools and streams

	All GET/POST requests will come thru here at one point or another if it needs to interact with pools/streams
*/

type Simple interface {
	CreatePool() ()
	CreateStream()
}

type User interface {
	
}

type Pool interface {

}
