package main

import (
	"log"
	"runtime"
	"os"
	"os/exec"
	"strings"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Println("无法获取当前路径:", err)
		return
	}
	
	log.Println("当前路径:", dir)
	arch := runtime.GOARCH
	EXEfile := ""
	link := os.Getenv("LINK")
	switch arch {
	case "amd64":
		log.Println("当前架构是:amd64")
		EXEfile="/exe/clash-speedtest_x86_64"
	case "386":
		log.Println("当前架构是:386")
		EXEfile="/exe/clash-speedtest_i386"
	case "arm64":
		log.Println("当前架构是:arm64")	
		EXEfile="/exe/clash-speedtest_arm64"
	default:
		log.Println("未知操作系统位数:", arch)
	}

	if EXEfile!=""{
		EXEfile = dir + EXEfile

		log.Println("EXEfile:", EXEfile)
		cmd := exec.Command(EXEfile, "-c", link, "--size", "2") // 在这里替换为你想要执行的命令
		output, err := cmd.Output()
		if err != nil {
			log.Println("命令执行失败:", err)
			return
		}
	
		result := strings.TrimSpace(string(output))
		log.Println("命令执行结果:")
		log.Println(result)

	}
}
