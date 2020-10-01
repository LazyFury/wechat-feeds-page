package download

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"wechatFeeds/utils"
)

// Download 下载文件
func Download(url string, fileName string) (path string, fileBytes []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	fileBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if fileName == "" {
		fileName = "temp_" + time.Now().String()
	}

	dir, err := utils.GetDir("./static", "")
	if err != nil {
		return
	}

	filePath := filepath.Join(dir, fileName)
	file, err := os.Create(filePath)
	_, err = file.Write(fileBytes)
	if err != nil {
		return
	}

	path = filePath
	return
}
