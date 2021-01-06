package main

import (
	"fmt"
	"time"
)

//节点
type Node struct {
	time    int64
	counter int
}

func NewSlidingWindow(max, maxRequest int, interval int64) *SlidingWindow {
	return &SlidingWindow{
		max,
		0,
		[]Node{},
		interval,
		maxRequest,
	}
}

type SlidingWindow struct {
	Max        int
	Len        int
	arr        []Node
	interval   int64
	maxRequest int
}
//有思考并发加减的情况，应该加锁还是原子替换，不知道怎么整。
func (this *SlidingWindow) Count() bool {

	//如果数组长度是0 那么是刚启动
	if this.Len == 0 {
		this.arr = append(this.arr, Node{time.Now().UnixNano() / 1e6, 1})
		this.Len++
		return true
	}
	//如果当前时间减去最近一个格子的时间小于间隔时间并且整个窗口的最大请求书不超过规定的最大请求数则最近一个格子的计数器就加1
	if (time.Now().UnixNano()/1e6)-this.arr[this.Len-1].time < this.interval {
		nowRequest := 0
		for _, v := range this.arr {
			nowRequest += v.counter
		}
		//如果当前窗口内的所有请求数+1已经超过最大允许单位内的请求数则返回false限流
		if nowRequest+1 > this.maxRequest {
			return false
		}
		this.arr[this.Len-1].counter++
		return true
	}
	this.arr = append(this.arr, Node{time.Now().UnixNano() / 1e6, 1})
	this.Len++
	//数组长度是6的一个滑动窗口，如果当前数组大于6个就把第一个干掉
	if this.Len > this.Max {
		this.arr = this.arr[1:]
		this.Len--
	}

	return true
}

func main() {
	ticker := &time.Ticker{}
	IsCurrentLimiting := false
	window := NewSlidingWindow(6, 600, 1000)
	for i := 0; i < 32000; i++ {
		s := window.Count()
		if s == false && IsCurrentLimiting == false {
			//开始限流
			IsCurrentLimiting = true
			fmt.Println("限流开始冷却")
			ticker = time.NewTicker(time.Second * 5)
			break
		}
		time.Sleep(time.Nanosecond * 1000)
	}
	if IsCurrentLimiting {
		fmt.Println(window.arr)
		fmt.Println("已限流")
		select {
		case <-ticker.C:
			fmt.Println("限流冷却结束")
		}
	} else {
		fmt.Println(window.arr)
		fmt.Println("未限流")
	}
}
