package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// 创建sleeper接口
type Sleeper interface {
	sleep()
}

// 创建监视器 可以记录依赖关系是怎么被使用的
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) sleep() {
	s.Calls++
}

type ConfigurableSleeper struct {
	duration time.Duration
}

func (scfg *ConfigurableSleeper) sleep() {
	time.Sleep(1 * scfg.duration)
}
func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second}
	CountDown(os.Stdout, sleeper)
}

func CountDown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.sleep()
	}
	sleeper.sleep()
	fmt.Fprint(out, "go!")
}
