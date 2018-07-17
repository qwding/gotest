package main

import (
	"flag"
	"fmt"
	// "io/ioutil"
	// "os"
	"os/exec"
	"time"
)

var container *string = flag.String("con", "", "input the container id.")

type Gotty struct {
	Cmd *exec.Cmd
}

func main() {
	flag.Parse()
	if *container == "" {
		fmt.Println("please give container id.")
		return
	}
	fmt.Println("this is begging")
	time.Sleep(time.Second * 3)

	gotty := Gotty{}

	err := gotty.RunGotty(*container)
	if err != nil {
		fmt.Println("error!", err)
	}
	fmt.Println("run. end!")
	var input string
	fmt.Scanln(&input)
	if err := gotty.Cmd.Process.Kill(); err != nil {
		fmt.Println("failed to kill: ", err)
	} else {
		fmt.Println("success killed.")
	}
}

func (g *Gotty) RunGotty(container string) error {
	// command := "nohup"
	// commands := []string{"gotty", "-w", "-port=8002", "docker", "exec", "-it", container, "/bin/bash"}
	command := "gotty"
	commands := []string{"-w", "-port=8002", "docker", "exec", "-it", container, "/bin/bash"}
	fmt.Println("cmd:", command, commands)
	return g.ExecRunFunc(command, commands)
}

func (g *Gotty) ExecRunFunc(command string, commands []string) error {
	g.Cmd = exec.Command(command, commands...)
	var err error
	fmt.Println("debug: ExecFunc", "before Run.")
	go func() {
		err = g.Cmd.Run()
		if err != nil {
			fmt.Println("debug ExecRunFunc", err)
			return
		}
		fmt.Println("debug ExecFunc", command, commands)
	}()

	return nil
}

//run exec command
func ExecFunc(command string, commands []string) error {
	cmd := exec.Command(command, commands...)
	var err error
	//cmd error
	/*stderr, errPipe := cmd.StderrPipe()
	if errPipe != nil {
		return errPipe
	}*/
	fmt.Println("debug: ExecFunc", "before Start.")
	if errStart := cmd.Start(); errStart != nil {
		return errStart
	}
	/*fmt.Println("debug: ExecFunc", "before errStart.")
	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		fmt.Println("debug: ExecFunc", err)
		return err
	}*/
	/*fmt.Println("debug: ExecFunc", "before len(bytesErr).")
	if len(bytesErr) != 0 {
		return fmt.Errorf(string(bytesErr))
	}*/
	/*fmt.Println("debug: ExecFunc", "before wait.")
	if err := cmd.Wait(); err != nil {
		return err
	}*/
	fmt.Println("debug ExecFunc", command, commands)
	return err
}
