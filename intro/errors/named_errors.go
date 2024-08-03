package main
import (
	"errors"
	"log"
	"time"
)

var ranOnce bool

var (
	ErrNetwork = errors.New("network error")
	ErrInput = errors.New("Input Error")
)

func nameErrorFunc(data string) error{
	if !ranOnce {
		ranOnce = true
		return ErrNetwork
	}
	return ErrInput
}

func main() {
	for {
		err := nameErrorFunc("data")
		if err == nil {
			break
		}
		if errors.Is(err, ErrNetwork) {
			log.Println("this is network error")
			time.Sleep(1 * time.Second)
			continue
		} else if errors.Is(err, ErrInput) {
			log.Println("this is input error")
			break
		} else {
			log.Println("this is unknown error")
			break
		}

	}
}