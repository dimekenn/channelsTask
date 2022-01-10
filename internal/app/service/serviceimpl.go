package service

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ServiceImpl struct {
}

func NewService() Service {
	return &ServiceImpl{}
}

func (s ServiceImpl) FromNChannelsToOneChannel(channelCount, randomIn int) []int {
	var chArr []chan int
	for i:=0;i<channelCount;i++{
		ch := make(chan int)
		fmt.Printf("channel %d created\n", i+1)
		arr := []int{rand.New(rand.NewSource(time.Now().UnixNano())).Intn(randomIn), rand.New(rand.NewSource(time.Now().UnixNano())).Intn(randomIn), rand.New(rand.NewSource(time.Now().UnixNano())).Intn(randomIn)}
		go func() {
			for _, v := range arr{
				ch <- v
			}
			defer close(ch)
		}()
		chArr = append(chArr, ch)
	}

	var intArr []int

	for num := range joinChannels(chArr) {
		intArr = append(intArr, num)
		fmt.Println("num: ",num)
	}
	return intArr
}


func joinChannels(channels []chan int) <-chan int {
	totalCh := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}

		wg.Add(len(channels))

		for _, ch := range channels {
			go func(ch <-chan int, wg *sync.WaitGroup) {
				for id := range ch {
					totalCh <- id
				}
				wg.Done()
			}(ch, wg)

		}

		wg.Wait()
		close(totalCh)
	}()

	return totalCh
}
