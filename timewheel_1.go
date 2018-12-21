package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/ouqiang/timewheel"
)

func addData(tw *timewheel.TimeWheel) {
	ti := time.Tick(time.Millisecond * time.Duration(1))
	key := 0
	for {
		select {
		case <-ti:
			if key == 10000 {
				return
			}

			key++
			addtime := time.Now().Format("2006-01-02 15:04:05")

			skey := fmt.Sprintf("timer_key_%d", key)
			tw.AddTimer(time.Second*time.Duration(1), skey, timewheel.TaskData{"key": skey, "uid": key, "addtime": addtime})

			fmt.Println("add to TimeWheel...", key)
		}
	}
}

var (
	tw *timewheel.TimeWheel
)

func WheelTest_1() {
	tw = timewheel.New(time.Second*time.Duration(1), 3600, func(data timewheel.TaskData) {
		fmt.Println(data["key"], data["uid"], data["addtime"])

		addtime := time.Now().Format("2006-01-02 15:04:05")
		tw.AddTimer(time.Second*time.Duration(1), data["key"], timewheel.TaskData{"key": data["key"], "uid": data["uid"], "addtime": addtime})
	})

	tw.Start()

	go addData(tw)

	// tw.RemoveTimer("aaaaa")
	// tw.Stop()

	select {}

	for {
		runtime.Gosched()
	}
}
