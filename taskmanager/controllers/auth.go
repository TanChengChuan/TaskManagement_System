package controllers

import (
	"os"
	"os/exec"
	"sync"
	"time"
)

// 使用WaitGroup来等待所有goroutine完成
var Wg sync.WaitGroup

func ClearScreen() { //清个屏，稍微好看一点
	time.Sleep(200 * time.Millisecond)
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
