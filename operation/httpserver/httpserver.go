package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
        //"time"
        "io/ioutil"
)

func index(w http.ResponseWriter, r *http.Request){
	//3.设置version
	os.Setenv("VERSION", "v1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	fmt.Printf("os version: %s \n", version)

	//2.将requst中的header放到reponse中
	for k, v := range r.Header {
		for _, vv := range v {
			fmt.Printf("%s: %s\n", k, v)
			w.Header().Set(k, vv)
		}
	}
	//4.记录日志并输出
	clientip := ClientIP(r)
        uri := RequestUri(r)
	//fmt.Println(r.RemoteAddr)
        if _,err := ioutil.ReadAll(r.Body);err == nil {
                //t1 := time.Now().Unix()
                //date := time.Unix(t1, 0).String()
	        log.Printf("ClientIP:%s Uri:%s Response code:%d", clientip,uri,200)
        }
}

//5.健康检查
func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "code: 200\n")
}

/*
func getCurrentIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		// 当请求头不存在即不存在代理时直接获取ip
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}
*/

// 解析RemoteAddr、X-Real-IP和X-Forwarded-For
func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func RequestUri(r *http.Request) string {
	url := r.RequestURI
        return url

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz", healthz)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Start http server failed, error: %s\n", err.Error())
	}
}
