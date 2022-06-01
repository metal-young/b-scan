package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wujiangweiphp/go-curl"
	"log"
)

func HttpGet(url string, queries map[string]string) (string, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	req := curl.NewRequest()
	resp, err := req.
		SetUrl(url).
		SetHeaders(headers).
		SetQueries(queries).
		Get()
	if err != nil {
		return "", err
	} else {
		if resp.IsOk() {
			return resp.Body, nil
		} else {
			log.Printf("%v\n", resp.Raw)
			return "", errors.New("请求失败")
		}
	}
}

func main() {
	fmt.Print("                                                                              \n")
	fmt.Print("███╗   ██████████████████╗█████╗██╗ ██╗   ██╗██████╗██╗   █████╗   ██╗██████╗\n")
	fmt.Print("████╗ ██████╔════╚══██╔══██╔══████║ ╚██╗ ██╔██╔═══████║   ██████╗  ████╔════╝\n")
	fmt.Print("██╔████╔███████╗    ██║  █████████║  ╚████╔╝██║   ████║   ████╔██╗ ████║  ███╗\n")
	fmt.Print("██║╚██╔╝████╔══╝    ██║  ██╔══████║   ╚██╔╝ ██║   ████║   ████║╚██╗████║   ██║\n")
	fmt.Print("██║╚██╔╝████╔══╝    ██║  ██╔══████║   ╚██╔╝ ██║   ████║   ████║╚██╗████║   ██║\n")
	fmt.Print("██║ ╚═╝ █████████╗  ██║  ██║  ███████████║  ╚██████╔╚██████╔██║ ╚████╚██████╔╝\n")
	fmt.Print("╚═╝ 小  ╚═╚══════╝雨╚═╝  ╚═╝  ╚═╚══════╚═╝青 ╚═════╝ ╚═════╝╚═╝年╚═══╝╚═════╝ \n")
	fmt.Print("                                                                              \n")

	//个人信息
	//url := "https://m.bilibili.com/space/7979698"
	//https://api.bilibili.com/x/space/upstat?mid=7979698
	//{"code":0,"message":"0","ttl":1,"data":{"archive":{"view":42486973},"article":{"view":1560936},"likes":750746}} 点赞数据
	url := "https://api.bilibili.com/x/relation/stat" //关注粉丝数据
	queries := map[string]string{
		"vmid": "10462362",
	}
	res, err := HttpGet(url, queries)
	if err != nil {
		fmt.Println(err)
		return
	}
	var stat interface{}
	err = json.Unmarshal([]byte(res), &stat)
	if err != nil {
		fmt.Println(err)
		return
	}
	ad := stat.(map[string]interface{})
	fmt.Printf("粉丝数：%.0f", ad["data"].(map[string]interface{})["follower"])
}
