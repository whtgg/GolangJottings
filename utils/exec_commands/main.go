package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	// basicExec()
	// fileExec()
	// execBuffer()
	// execCombinedOutput()
	// execStd()
	// runServer()

	//data := []byte("hello world")
	//compressed, _ := bzipCompress(data)
	//r := bzip2.NewReader(bytes.NewBuffer(compressed))
	//decompressed, _ := ioutil.ReadAll(r)
	//fmt.Println(string(decompressed))

	execExistCommand()
}

// 检查命令是否存在
func execExistCommand() {
	if cmd, err := exec.LookPath("ls"); err != nil {
		fmt.Println("command is not exist", err.Error())
	} else {
		fmt.Println("command is", cmd)
	}
}

// 从标准输入获取参数
func bzipCompress(d []byte) ([]byte, error) {
	var out bytes.Buffer
	cmd := exec.Command("bzip2", "-c", "-9")
	cmd.Stdin = bytes.NewBuffer(d)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed: %v\n", err)
	}

	return out.Bytes(), nil
}

// 分别获取标准输出和标准错误
func execStd() {
	var stdOut, stdErr bytes.Buffer
	cmd := exec.Command("cal")
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed: %v\n", err)
	}

	fmt.Println(stdOut.String())
}

// CombinedOutput
func execCombinedOutput() {
	cmd := exec.Command("cal")
	if res, err := cmd.CombinedOutput(); err != nil {
		fmt.Println("cmd run error", err)
	} else {
		fmt.Println(string(res))
	}
}

// 保存到内存对象中
func execBuffer() {
	buf := bytes.NewBuffer(nil)
	cmd := exec.Command("cal")
	cmd.Stdout = buf
	cmd.Stderr = buf
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed: %v\n", err)
	}

	fmt.Println(buf.String())
}

// 简单服务器
func runServer() {
	http.HandleFunc("/cal", httpHandle)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("server error", err)
	}
}

// 发送到网络
func httpHandle(w http.ResponseWriter, r *http.Request) {
	year := r.URL.Query().Get("year")
	month := r.URL.Query().Get("month")

	f, err := os.OpenFile("out.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("open file error", err.Error())
	}
	defer f.Close()

	buf := bytes.NewBuffer(nil)
	mw := io.MultiWriter(w, f, buf)

	cmd := exec.Command("cal", month, year)
	cmd.Stdout = mw
	cmd.Stderr = mw

	if err := cmd.Run(); err != nil {
		fmt.Println("command run error", err.Error())
	}
	fmt.Println(buf.String())
}

// 输出到文件
func fileExec() {
	f, err := os.OpenFile("open.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("open error", err.Error())
	}
	defer f.Close()
	cmd := exec.Command("cal")
	cmd.Stdout = f
	cmd.Stderr = f
	if err := cmd.Run(); err != nil {
		fmt.Println("cmd run error", err.Error())
	}

}

// 显示到标准输出
func basicExec() {
	cmd := exec.Command("cal")
	// 如果不使用 不会显示命令执行结果
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdin
	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("cal command")
}
