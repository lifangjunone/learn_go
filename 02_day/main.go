package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// ClearAndWrite 文件读写
func ClearAndWrite() {
	filePath := "/home/lifangjun/Desktop/go_study/learn_go/02_day/1.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err = %v\n", err)
		return
	}
	//　close file stream
	defer file.Close()
	// write

	wStr := "您好！！今晚我想吃鱼了."
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(wStr)
	}
	// 因为　writer 是带缓存的，需要通过 flush 到磁盘
	writer.Flush()
	file.Seek(0, 0)

	// read
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('.')
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}
}

// CopyFileToNewFile 拷贝文件内容到新的文件中
func CopyFileToNewFile() {
	oldFilePath := "/home/lifangjun/Desktop/go_study/learn_go/02_day/1.txt"
	newFilePath := "/home/lifangjun/Desktop/go_study/learn_go/02_day/2.txt"
	data, err := ioutil.ReadFile(oldFilePath)
	if err != nil {
		fmt.Printf("read file err = %v", err)
	}
	err = ioutil.WriteFile(newFilePath, data, 0666)
	if err != nil {
		fmt.Printf("write file err=%v", err)
	}
}

// JudgePathIsExists 判断文件或者文件夹是否存在
func JudgePathIsExists() (bool, error) {
	myPath := "/home/lifangjun/Desktop/go_study/learn_go/02_day/1.txt"
	_, err := os.Stat(myPath)
	if err == nil {
		fmt.Printf("当前文件/文件夹存在\n")
		return true, nil
	}
	if os.IsNotExist(err) {
		fmt.Printf("当前文件/文件夹不存在\n")
		return false, nil
	}
	return false, nil
}

// FileCopy 文件拷贝
func FileCopy() {

}

// 解析命令行参数
func parseCommandLineParams() {
	paramLen := len(os.Args)
	fmt.Println(paramLen)
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v \n", i, v)
	}
}

func parseCommandLineParams2() {
	var (
		user string
		pwd  string
		host string
		port string
	)
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，　默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为空")
	flag.StringVar(&port, "port", "3306", "端口号，　默认为空")
	flag.Parse()
	fmt.Printf("user=%v pwd=%v host=%v port=%v \n", user, pwd, host, port)
}

func main() {
	// ClearAndWrite()
	//CopyFileToNewFile()
	//JudgePathIsExists()
	parseCommandLineParams2()
}
