package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

var wg sync.WaitGroup

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
		EXEfile = "/exe/clash-speedtest_x86_64"
	case "386":
		log.Println("当前架构是:386")
		EXEfile = "/exe/clash-speedtest_i386"
	case "arm64":
		log.Println("当前架构是:arm64")
		EXEfile = "/exe/clash-speedtest_arm64"
	default:
		log.Println("未知操作系统位数:", arch)
	}

	if EXEfile != "" {
		EXEfile = dir + EXEfile
		log.Println("EXEfile:", EXEfile)
		profiles := strings.Split(link, ",")

		for i, profile := range profiles {
			wg.Add(1)
			if profile != "" {
				go cmdStart(EXEfile, profile, i)
			}
		}
		wg.Wait()
	}
}

func cmdStart(EXEfile string, speedUrl string, num int) {
	defer wg.Done()
	cmd := exec.Command(EXEfile, "-c", speedUrl, "--size", "2") // 在这里替换为你想要执行的命令
	output, err := cmd.Output()
	result := strings.TrimSpace(string(output))
	if err != nil {
		log.Println(num, "命令执行失败:", err)
	}
	log.Println(num, "命令执行结果:", result)
}
