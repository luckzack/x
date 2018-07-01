package shell

import (
	"bufio"
	"io"
	"log"
	"os/exec"
	"strings"

	"github.com/fiam/gounidecode/unidecode"
)

//func main() {
//	command := "ping"
//	params := []string{"192.168.1.210"}
//	//执行cmd命令: ls -l
//	execCommand(command, params)
//}

func ExecCommandWithDetail(commandName string, params []string, output chan string) error {
	log.Println("ExecCommandWithDetail --->", commandName, params)
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	//fmt.Println(cmd.Args)

	ch_out := make(chan error)

	go func(ch chan error) {
		stdout, err := cmd.StdoutPipe()

		if err != nil {
			ch_out <- err
		}

		cmd.Start()

		reader := bufio.NewReader(stdout)

		for {
			line, err2 := reader.ReadString('\n')
			if err2 != nil && io.EOF != err2 {
				ch_out <- err2
			} else if io.EOF == err2 {
				//log.Println("##############################1")
				ch_out <- nil
			} else {
				output <- unidecode.Unidecode(line)
			}

		}
	}(ch_out)

	ch_err := make(chan error)
	go func() {
		stderr, err := cmd.StderrPipe()

		if err != nil {
			ch_err <- err
		}

		reader := bufio.NewReader(stderr)

		for {
			line, err2 := reader.ReadString('\n')

			if err2 != nil && io.EOF != err2 {
				ch_err <- err2
			} else if io.EOF == err2 {
				//log.Println("##############################2")
				ch_err <- nil
			} else {
				output <- unidecode.Unidecode(line)
			}

		}
	}()

	for {
		select {
		case err := <-ch_out:
			return err
		case err := <-ch_err:
			return err
		}

	}

	cmd.Wait()
	return nil
}

func ExecCommand(commandName string, params []string) error {
	log.Println("ExecCommand --->", commandName, params)
	cmd := exec.Command(commandName, params...)
	return cmd.Run()
}

func Chmod(file_path string, mode string) error {
	return ExecCommand("chmod", []string{mode, file_path})
}

func ChmodX(file_path string) error {
	return ExecCommand("chmod", []string{"+x", file_path})
}

///
func Exec(shell string) (string, error) {

	cmd := exec.Command("/bin/sh", "-c", shell)
	bytes, err := cmd.Output()

	return string(bytes), err
}

func ExecShellFile(file string) (string, error) {
	cmd := exec.Command("which", "sh")
	bytes, err := cmd.Output()
	if err != nil {
		return "", err
	}

	sh := strings.Replace(string(bytes), "\n", "", -1)
	cmd = exec.Command(sh, file)
	//cmd := exec.Command("/bin/sh", file)
	bytes, err = cmd.Output()

	return string(bytes), err
}

func ExecShellFileV2(file string) (string, error) {
	//cmd := exec.Command("which", "sh")
	//bytes, err := cmd.Output()
	//if err != nil {
	//	return "", err
	//}
	//sh := strings.Replace(string(bytes), "\n", "", -1)

	cmd := exec.Command("/bin/cat", file)
	bytes, err := cmd.Output()
	if err != nil {
		return "", err
	}
	content := strings.Replace(string(bytes), "\n", "", -1)
	content = strings.Replace(content, "\r", "", -1)
	//bytes, _ = json.Marshal(&content)
	//cmd = exec.Command(sh, "-c", string(bytes))
	//log.Println("#######", sh , string(bytes))
	//bytes, err = cmd.Output()

	//	return string(bytes),err
	return Exec(content)
}
