package controllers

import "sync"

// 使用WaitGroup来等待所有goroutine完成
var Wg sync.WaitGroup
