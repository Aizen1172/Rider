package logger

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func Logger() {
	SetLogger()
	//simpleHttpGetUrl("http://localhost:8080/")
}

func SetLogger() {
	path := time.Now().Format("2006-01-02")
	//创建日志文件
	logFile, err := os.OpenFile("./log/"+path+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Println("创建失败。", err)
	}
	//设置日志输出的信息   日期 /          时间      /      文件信息
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	//输出到日志文件      实现多目的的传送，输出到日志文件和控制台
	log.SetOutput(io.MultiWriter(logFile, os.Stdout)) 
}

func simpleHttpGetUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("错误："+url, err.Error())
		return
	} else {
		fmt.Println(resp)
		log.Println("成功："+url, resp.Status)
		_ = resp.Body.Close()
	}

}
