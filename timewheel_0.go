package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	// "github.com/ouqiang/timewheel"
	"github.com/wgliang/timewheel"
)

type Wheel struct {
	mux   *sync.Mutex
	Cache map[string]string
}

func NewWheel() *Wheel {
	wh := &Wheel{
		mux: new(sync.Mutex),
	}
	wh.Cache = make(map[string]string, 0)
	return wh
}

func (wh *Wheel) Get(key string) string {
	wh.mux.Lock()
	defer wh.mux.Unlock()

	return wh.Cache[key]
}

func (wh *Wheel) Add(key, value string) {
	wh.mux.Lock()
	defer wh.mux.Unlock()

	wh.Cache[key] = value
}

func (wh *Wheel) Remove(key string) {
	wh.mux.Lock()
	defer wh.mux.Unlock()

	delete(wh.Cache, key)
}

func goValue(tw *timewheel.TimeWheel, wh *Wheel) {
	ti := time.Tick(time.Second * time.Duration(1))
	key := 0
	for {
		select {
		case <-ti:
			if key == 20 {
				return
			}

			key++
			value := time.Now().Format("2006-01-02 15:04:05")
			// 业务中对资源的管理
			wh.Add(strconv.Itoa(key), value)
			fmt.Println("add to Wheel...", value)
			// 不要忘记同时在时间轮里也要做改变，原则就是业务中的改变记得通知时间轮，但时间轮做的工作我们无需关心
			tw.Add(strconv.Itoa(key))
			fmt.Println("add to TimeWheel...", key)
		}
	}
}

func printWheel(wh *Wheel) {
	ti := time.Tick(time.Second * time.Duration(1))
	key := 0
	for {
		select {
		case <-ti:
			if key == 30 {
				return
			}
			key++
			wh.mux.Lock()
			fmt.Println(wh.Cache)
			wh.mux.Unlock()
		}
	}
}

func WheelTest_0() {
	// 初始化你的资源和接口
	wh := NewWheel()
	// 传入你要对你的资源要做的操作，以及传入回调函数
	wheel := timewheel.NewTimeWheel(time.Second*time.Duration(1), 3600, func(w interface{}, key interface{}) {
		w.(*Wheel).Remove(key.(string))
	}, wh)
	// 开启我们的时间轮
	wheel.Start()
	// 在这里代表你的项目中对资源做的改变，增加等等，然后剩下的就交给时间轮管理吧
	go goValue(wheel, wh)
	go printWheel(wh)

	time.Sleep(time.Duration(30) * time.Second)
}
