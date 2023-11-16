package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("カレントディレクトリの取得に失敗しました: %v\n", err)
		return
	}
	dir := filepath.Dir(exePath)
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("ディレクトリの読み込みに失敗しました: %v", err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		info, err := entry.Info()
		if err != nil {
			log.Fatalf("ファイル情報の取得に失敗しました: %v", err)
		}
		fmt.Printf("%v\n", info.Name())
	}
}
