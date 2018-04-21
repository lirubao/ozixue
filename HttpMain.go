package main

import (
	"log"      //调试
	"net/http" //go标准库http内核
	"time"
)

func main() {
	//搭建简单的http服务
	//version1()
	//version2()
	version3()
}

type myHandler struct{}

//myHandler实现Handler接口中的ServerHTTP方法
func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//使用r.URL.String()查看未注册的地址
	w.Write([]byte("Hello v2 ,the request URL is:" + r.URL.String()))
}

func version3() {
	server := &http.Server{
		Addr:         ":80",           //端口
		WriteTimeout: 2 * time.Second, //3秒
	}
	//使用自定义handleFunc
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye3)
	server.Handler = mux
	//打印日记
	log.Println("Staring server... v3")
	//如果http服务在监听时出错执行或者退出http服务时执行
	log.Fatal(server.ListenAndServe())
}
func sayBye3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bye bye , this is version 3!"))
}

func version2() {
	//使用自定义handleFunc
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye2)
	//打印日记
	log.Println("Staring server... v2")
	//如果http服务在监听时出错执行或者退出http服务时执行
	log.Fatal(http.ListenAndServe(":80", mux))
}
func sayBye2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bye bye , this is version 2!"))
}

func version1() {
	//设置路由1
	http.HandleFunc("/", handler)
	//设置路由2
	http.HandleFunc("/bye", sayBye1)
	//打印日记
	log.Println("Staring server... v1")
	//如果http服务在监听时出错执行或者退出http服务时执行
	log.Fatal(http.ListenAndServe(":80", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello ,this is version 1!"))
}

func sayBye1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bye bye , this is version 1!"))
}
