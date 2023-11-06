package main

import (
	"bytes"
	"fmt"
	"github.com/unidoc/unipdf/v3/common/license"
	"log"
	"os"

	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

func init() {
	// The customer name needs to match the entry that is embedded in the signed key.
	//customerName := "dandan"
	offlineLicenseKey := "2875f02ab968144525eb8e22681a06b54644f360eae6bbda929fd930f7953ebe"
	// Good to load the license key in `init`. Needs to be done prior to using the library, otherwise operations
	// will result in an error.
	err := license.SetMeteredKey(offlineLicenseKey)
	if err != nil {
		panic(err)
	}
}
func main() {
	// 打开 PDF 文件
	lk := license.GetLicenseKey()
	if lk == nil {
		fmt.Printf("Failed retrieving license key")
		return
	}
	fmt.Printf("License: %s\n", lk.ToString())
	pdfPath := "test.pdf"
	f, err := os.ReadFile(pdfPath)
	if err != nil {
		log.Fatalf("无法打开 PDF 文件：%v", err)
	}

	// 创建 PDF 解析器
	pdfReader, err := model.NewPdfReader(bytes.NewReader(f))
	if err != nil {
		log.Fatalf("无法创建 PDF 解析器：%v", err)
	}

	// 获取 PDF 页面数量
	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		log.Fatalf("无法获取 PDF 页面数量：%v", err)
	}

	// 遍历每个页面并提取文本和图片
	for pageNum := 1; pageNum <= numPages; pageNum++ {
		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			log.Printf("无法获取第 %d 页：%v", pageNum, err)
			continue
		}

		extractor, err := extractor.New(page)
		if err != nil {
			log.Printf("无法创建页面解析器：%v", err)
			continue
		}

		// 提取文本内容
		text, err := extractor.ExtractText()
		if err != nil {
			log.Printf("无法提取文本内容：%v", err)
		} else {
			fmt.Printf("第 %d 页文本内容：\n%s\n", pageNum, text)
		}
		// 提取图片内容
		//images, err := extractor.ExtractImages()
		//if err != nil {
		//	log.Printf("无法提取图片内容：%v", err)
		//} else {
		//	for _, img := range images {
		//		fmt.Printf("第 %d 页图片：%s\n", pageNum, img)
		//	}
		//}
	}
}
