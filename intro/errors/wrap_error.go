package main
import (
	"errors"
	"fmt"
	"log"
	"time"
)
var ranOnce bool

const (
	UnknownCode     = 0
	UnreachableCode = 1
	AuthFailureCode = 2
)

// ErrNetwork is a custom network error that had error codes
// and a message.
type ErrNetwork struct {
	Code int
	Msg  string
}

func (e ErrNetwork) Error() string {
	return fmt.Sprintf("network error(%d): %s", e.Code, e.Msg)
}
func someFunc(data string) error {
	if !ranOnce {
		ranOnce = true
		return ErrNetwork{
			Code: AuthFailureCode,
			Msg:  "user unrecognized",
		}
	}
	return fmt.Errorf("another error")
}

func restCall(data string) error {
	if err := someFunc(data); err != nil {
		return fmt.Errorf("restcall(%s) had an error: %w", data, err)
	}
	return nil
}

func main() {
	for i:=0; i<5; i++ {
		if err := restCall("data"); err != nil {
			var netErr ErrNetwork
			if errors.As(err, &netErr) {
				log.Println("network error: ", err)
				time.Sleep(1 * time.Second)
				continue
			}
			log.Println("unrecoverable: ", err)
		}
	}
}
//output
/*
2024/08/04 20:45:08 network error:  restcall(data) had an error: network error(2): user unrecognized
2024/08/04 20:45:09 unrecoverable:  restcall(data) had an error: another error
2024/08/04 20:45:09 unrecoverable:  restcall(data) had an error: another error
2024/08/04 20:45:09 unrecoverable:  restcall(data) had an error: another error
2024/08/04 20:45:09 unrecoverable:  restcall(data) had an error: another error
*/

// without %w output will look

/*
2024/08/04 20:44:59 unrecoverable:  restcall(data) had an error
2024/08/04 20:44:59 unrecoverable:  restcall(data) had an error
2024/08/04 20:44:59 unrecoverable:  restcall(data) had an error
2024/08/04 20:44:59 unrecoverable:  restcall(data) had an error
2024/08/04 20:44:59 unrecoverable:  restcall(data) had an error
*/