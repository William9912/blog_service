package middleware

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var lock1 sync.Mutex
var lock2 sync.Mutex
var needCount bool = true
var x int64 = 0

func TimeCounter(conut time.Duration) {
	time.Sleep(conut)
	WriteFileAndSet0()
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true

}

func WriteFileAndSet0() {
	lock1.Lock()
	var wireteString = fmt.Sprintf("%v ----qps is :%v \n", time.Now().Format("2006-01-02 15:04:05"), x)
	lock1.Unlock()
	var filename = "storage\\logs\\qps.txt"
	var f *os.File
	/***************************** 第一种方式: 使用 io.WriteString 写入文件 ***********************************************/
	if Exists(filename) { //如果文件存在
		f, _ = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
	} else {
		f, _ = os.Create(filename) //创建文件
	}
	io.WriteString(f, wireteString) //写入文件(字符串)
	lock1.Lock()
	x = 0
	lock1.Unlock()
	lock2.Lock()
	needCount = true
	lock2.Unlock()
}

func QPS() gin.HandlerFunc {
	return func(c *gin.Context) {
		lock1.Lock()
		x++
		lock1.Unlock()
		//一秒之后 输出 x 归零
		lock2.Lock()
		if needCount {
			needCount = false
			go TimeCounter(1 * time.Second)
		}
		lock2.Unlock()
	}
}
