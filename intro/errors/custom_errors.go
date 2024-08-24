package main
import (
	"errors"
	"fmt"
	"log"
	"time"
)
var ranOnce bool = false

const (
	UnknownCode     = 0
	UnreachableCode = 1
	AuthFailureCode = 2
)

// ErrNetwork is a custom network error that had error codes
// and a message.

type ErrNetwork struct {
	Code int
	Msg string
}

// this Error method of ErrNetwork register itself to error interface
func (e ErrNetwork) Error() string {
	return fmt.Sprintf("network error(%d): %s", e.Code, e.Msg)
}

// someFunc is a standin for some function that takes in some data.
// On the first call it will return an ErrNetwork and the second just
// a standard error.

//func can take any numbers of data with any types
func someFunc(data... interface {}) error {
	if !ranOnce {
		ranOnce = true
		return ErrNetwork{
			Code: AuthFailureCode,
			Msg: "user unrecognized",
		}
	}
	return fmt.Errorf("another error")
}

func main() {
	// This loop will continue as long as we have network errors and time within 10 seconds that are
	// not code AuthFailureCode and will  terminate on any other error or success.
	for i := 0; i< 10; i++ {
		if err := someFunc("data"); err != nil {
			var netErr ErrNetwork
			if errors.As(err, &netErr) {
				if netErr.Code == AuthFailureCode {
					log.Println("auth failure! Danger")
					return 
				}
				log.Println("network error: ", err)
				time.Sleep(1 * time.Second)
				continue
				
			}
			log.Println("unrecoverable: ", err)
		}
		break
	}
}