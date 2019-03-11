package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Input struct {
	Resources []string
	Times     int
	Timeout   int
}

var mutex = &sync.Mutex{}
var timeout int
var replyTime []float64
var noReply int

func main() {
	data, err := readFile("input.json")
	if err != nil {
		panic(err)
	}

	var input Input
	json.Unmarshal(data, &input)
	timeout = input.Timeout

	var wg sync.WaitGroup
	wg.Add(10)

	start := time.Now()
	for _, resource := range input.Resources {
		makeRequests(resource, input.Times, &wg)
	}
	wg.Wait()
	elapsed := time.Since(start).Seconds()

	var min, max, avg float64
	for _, reply := range replyTime {
		if min == 0 || min > reply {
			min = reply
		}

		if max < reply {
			max = reply
		}

		avg += reply
	}
	avg /= (10 - float64(noReply))

	fmt.Printf("Min request time: %f\n", min)
	fmt.Printf("Max request time: %f\n", max)
	fmt.Printf("Avg request time: %f\n", avg)
	fmt.Printf("Total time: %f\n", elapsed)
	fmt.Printf("Timed out requests: %d\n", noReply)
}

func makeRequests(resource string, times int, wg *sync.WaitGroup) {
	for i := 1; i <= times; i++ {
		go func() {
			makeRequest(resource, wg)
		}()
	}
}

func makeRequest(resource string, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()

	c := &http.Client{
		Timeout: time.Duration(timeout) * time.Millisecond,
	}
	result, err := c.Get(resource)

	if err != nil {
		mutex.Lock()
		noReply++
		mutex.Unlock()
	} else {
		defer result.Body.Close()
		elapsed := time.Since(start).Seconds()

		mutex.Lock()
		replyTime = append(replyTime, elapsed)
		mutex.Unlock()
	}
}

func readFile(fileName string) ([]byte, error) {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	return data, nil
}
