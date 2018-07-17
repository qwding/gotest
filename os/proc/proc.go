package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*
* this file is kill all the process by name "gotty".
 */

func main() {
	KillProcByName("gotty")
}

func KillProcByName(proc string) {
	cmd1 := exec.Command("ps", "-ef")
	cmd2 := exec.Command("grep", proc)

	var err error
	cmd2.Stdin, err = cmd1.StdoutPipe()
	if err != nil {
		fmt.Println("error of pipe", err)
	}

	var b bytes.Buffer

	cmd2.Stdout = &b
	if err != nil {
		fmt.Println("stdoutpipe error", err)
	}
	cmd2.Start()
	cmd1.Run()
	cmd2.Wait()

	output, err := ioutil.ReadAll(&b)
	if err != nil {
		fmt.Println("error readall", err)
	}

	fmt.Println("Info ps -ef|grep gotty output:")
	fmt.Println(string(output))
	arr := strings.Split(string(output), "\n")
	for _, val := range arr {
		infos := strings.Split(val, " ")
		count := 0
		pidStr := ""
		for _, info := range infos {
			if info == " " || info == "" {
				continue
			}
			count++
			if count == 2 {
				pidStr = info
			}
		}
		if pidStr == "" {
			continue
		}
		fmt.Println("pid is:", pidStr)

		pid, err := strconv.Atoi(pidStr)
		if err != nil {
			fmt.Println("error", err)
			continue
		}
		process := os.Process{
			Pid: pid,
		}
		err = process.Kill()
		if err != nil {
			fmt.Println("error kill process:", err)
		}
	}
}
