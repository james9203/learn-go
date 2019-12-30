package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

func main() {
	fmt.Println("OSS GO SDK VERSION:",oss.Version)

	// 创建OSSClient实例。
	client, err := oss.New("http://oss-cn-shanghai.aliyuncs.com", "", "")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}


	// 获取存储空间。
	bucket, err := client.Bucket("customer-product")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}


	// 下载文件到本地文件。
	err = bucket.GetObjectToFile("20190321182223864250/38bdc7b67d89425cb2763d1aa175bffd.jpg", "d:/e.jpg")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

}
