package helper

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

type SessionToken int

type ProcessHandler interface {

	HandleEcho(token SessionToken, echo string)

	HandleError(token SessionToken, exitCode int, state string)

	HandleSuccess(token SessionToken)
}

func processOut(reader io.ReadCloser) chan string {
	out := make(chan string)
	go func() {
		buf := make([]byte, 1024)
		for {
			bCount, err := reader.Read(buf)
			if err != nil {
				break
			}
			if bCount > 0 {
				out <- string(buf[0:bCount])
			}
		}
		close(out)
	}()
	return out
}

func ExecShellAdmin(s string, handler ProcessHandler, token SessionToken) {
	script := fmt.Sprintf("osascript -e \"do shell script \\\"%s\\\" with administrator privileges\"", s)
	ExecShellAsync(script, handler, token)
}

func ExecShellAsync(s string, handler ProcessHandler, token SessionToken) {
	cmd := exec.Command("/bin/bash", "-c", s+" 2>&1")
	out, _ := cmd.StdoutPipe()
	ch := processOut(out)
	cmd.Start()
	for echo := range ch {
		handler.HandleEcho(token, echo)
	}
	cmd.Wait()
	state := cmd.ProcessState
	if state.Success() {
		handler.HandleSuccess(token)
	} else {
		handler.HandleError(token, state.ExitCode(), state.String())
	}
}

func ExecShell(s string) (*os.ProcessState, string) {
	cmd := exec.Command("/bin/bash", "-c", s+" 2>&1")
	out, _ := cmd.Output()
	cmd.Run()
	return cmd.ProcessState, string(out)
}