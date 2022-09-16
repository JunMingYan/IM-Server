package test

import (
	"fmt"
	"testing"
	"time"
)

func TestSystemUser(t *testing.T) {
	t1 := time.Now().Unix()
	time.Sleep(time.Second * 60)
	t2 := time.Now().Unix()

	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t2 - t1)
}
