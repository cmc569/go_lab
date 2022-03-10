package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println													//指定以 p() 取代 fmt.Println()
	
	now := time.Now()
	p(now.Format(time.RFC3339))											//指定以 RFC3339 的日期顯示格式輸出

	day := time.Date(2022, 11, 10, 15, 32, 00, 651387237, time.UTC)		//指定日期時間與時區
	p(day)
}
