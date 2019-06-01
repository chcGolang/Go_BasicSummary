package main

import (
	"fmt"
	"testing"
)

func TestBatchPushQueue(t *testing.T) {
	err := BatchPushQueue("test:22", []string{"22", "33"})
	if err != nil {
		panic(err)
	}
}

func TestPopQueue(t *testing.T) {
	for {
		data, err := PopQueue("test:22", 0)
		if err != nil {
			panic(err)
		}
		fmt.Println(data)
	}

}
