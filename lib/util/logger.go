package util

import (
	"log"
	"runtime"
)

func PrintStack(err error) {
	var buf [1024]byte
    n := runtime.Stack(buf[:], true)
	log.Println(string(buf[:]), n)
}