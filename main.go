package main

import (
	"fmt"
	"wechatFeeds/download"
	"wechatFeeds/utils"
)

func main() {
	path, fileBytes, err := download.Download("https://github.com/hellodword/wechat-feeds/raw/main/list.csv", "list.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println("保存csv:", path)
	feeds := utils.CSVFormat(fileBytes)

	jsonPath, err := utils.SaveJSON(feeds)
	if err != nil {
		panic(err)
	}
	fmt.Println("保存为Json：", jsonPath)
}
