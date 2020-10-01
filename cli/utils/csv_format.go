package utils

import (
	"bytes"
	"encoding/csv"
	"io"
)

// WechatFeeds 类型
// name,bizid,description
type WechatFeeds struct {
	Name        string `json:"name"`
	BizID       string `json:"bizid"`
	Description string `json:"description"`
}

// CSVFormat 格式化csv
func CSVFormat(fileBytes []byte) (feeds []WechatFeeds) {
	reader := csv.NewReader(bytes.NewReader(fileBytes))

	// var feeds []WechatFeeds
	// 参考：https://www.kancloud.cn/mutouzhang/gocookbook/688578
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		if record[1] == "bizid" || record[1] == "" {
			continue
		}

		f := WechatFeeds{record[0], record[1], record[2]}
		feeds = append(feeds, f)
	}

	return
}
