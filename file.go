package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")	//開啟檔案
	if err != nil {						//判斷是否成功開啟
		return
	}
	defer file.Close()					//程式結束前關閉檔案

	stat, err := file.Stat()			//取得檔案相關資訊以指定變數(slice)容量需求
	if err != nil {
		return
	}

	bs := make([]byte, stat.Size())		//宣告bs變數(slice)
	count, err := file.Read(bs)			//count = 內容文字總數
	if err != nil {
		return
	}

	fmt.Println(count)

	str := string(bs)					//將讀取內容轉成文字格式
	fmt.Println(str)
}