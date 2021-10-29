package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func main() {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		process(c)
	})
	r.Run(":5701")
}

func process(c *gin.Context) {
	dataReader := c.Request.Body
	rawdata, _ := ioutil.ReadAll(dataReader)
	Mess := gjson.GetMany(string(rawdata), "post_type", "message_type", "group_id", "user_id", "message")
	fmt.Println(Mess)
	if Mess[0].String() == "message" && Mess[1].String() == "group" {

		reaction(Mess[4].String())
	}
}
func reaction(message string) {
	fmt.Println(message)
	if message == "你好,高性能干饭机器人" {
		url := "http://127.0.0.1:5700/send_group_msg?group_id=438377875&message=" + "你好"
		Get(url)
	}
}

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}
