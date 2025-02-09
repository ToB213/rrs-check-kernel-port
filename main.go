package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	dir := "maps"
	target := "common.cfg"

	// ポート番号を更新するかどうかを確認
	var flag int
	fmt.Println("Do you want to update the port number? (1: Yes, 2: No)")
	fmt.Scan(&flag)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 走査する必要がないディレクトリをスキップ
		if info.IsDir() && (info.Name() == "map" || info.Name() == "comms") {
			return filepath.SkipDir
		}

		if !info.IsDir() && strings.Contains(path, target) {
			fmt.Println("Processing", path)
			// 対象ファイルを操作する関数
			err := processFile(path, flag)
			if err != nil {
				fmt.Println("Error Processing", err)
				return err
			}
			fmt.Println("============================================================")
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}

// ファイルを操作する関数
func processFile(path string, flag int) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	// ファイルの内容をスライスに格納
	// 1行ずつ読み込む
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	modifiedLines := make([]string, len(lines))
	for i, line := range lines {
		// ポート番号を置換する関数
		// modifiedLinesに置換後の行を格納
		modifiedLines[i] = replaceKernelPort(line)
	}

	if flag == 1 {
		fmt.Println("Replace Port: 27931")

		// 書き込み用の関数
		err = writeFileContents(path, modifiedLines)
		if err != nil {
			return err
		}
	}

	return nil
}

// ポート番号を置換する関数
func replaceKernelPort(line string) string {
	targetLine := "kernel.port"
	rex := regexp.MustCompile("[0-9]+")

	// lineにtargetLineが含まれている場合，文字列内の数字を検出し，"27931"に置換
	if strings.Contains(line, targetLine) {
		currentPort := rex.FindString(line)
		fmt.Println("Current Port:", currentPort)

		return rex.ReplaceAllString(line, "27931")
	}
	return line
}

// ファイルの内容を書き換える関数
// 一時ファイルで編集後の内容を書き込み，元のファイルにリネーム（上書き）
func writeFileContents(path string, lines []string) error {
	// 一時ファイルを作成
	tmpFile, err := os.CreateTemp(filepath.Dir(path), "temp")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	// 一時ファイルに編集後の内容を書き込む
	w := bufio.NewWriter(tmpFile)
	for _, line := range lines {
		_, err := fmt.Fprintln(w, line)
		if err != nil {
			return err
		}
	}
	if err := w.Flush(); err != nil {
		return err
	}
	if err := tmpFile.Close(); err != nil {
		return err
	}

	// 一時ファイルを元のファイルにリネーム
	err = os.Rename(tmpFile.Name(), path)
	if err != nil {
		return err
	}

	fmt.Println("Updated", path)

	return nil
}
