package tutils

// copy from github.com/henrylee2cn/goutil
import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

var cmdBash = "/bin/bash"
var cmdArg = "-c"

// RunCmd exec cmd and catch the result.
// Waits for the given command to finish with a timeout.
// If the command times out, it attempts to kill the process.
func RunCmd(cmdLine string, envs map[string]string, timeout ...time.Duration) *Result {
	cmd := exec.Command(cmdBash, cmdArg, cmdLine)
	var ret = new(Result)
	cmd.Stdout = &ret.buf
	cmd.Stderr = &ret.buf
	cmd.Env = os.Environ()
	for k, v := range envs {
		env := fmt.Sprintf("%s=%s", k, v)
		cmd.Env = append(cmd.Env, env)
	}
	ret.err = cmd.Start()
	if ret.err != nil {
		return ret
	}
	if len(timeout) == 0 || timeout[0] <= 0 {
		ret.err = cmd.Wait()
		return ret
	}
	timer := time.NewTimer(timeout[0])
	defer timer.Stop()
	done := make(chan error)
	go func() { done <- cmd.Wait() }()
	select {
	case ret.err = <-done:
	case <-timer.C:
		if err := cmd.Process.Kill(); err != nil {
			ret.err = fmt.Errorf("command timed out and killing process fail: %s", err.Error())
		} else {
			// wait for the command to return after killing it
			<-done
			ret.err = errors.New("command timed out")
		}
	}
	return ret
}

// Result cmd exec result
type Result struct {
	buf bytes.Buffer
	err error
	str string
}

// Err returns the error log.
func (r *Result) Error() error {
	if r.err == nil {
		return nil
	}
	r.err = errors.New(r.String())
	return r.err
}

// String returns the exec log.
func (r *Result) String() string {
	b := bytes.TrimSpace(r.buf.Bytes())
	if r.err != nil {
		b = append(b, ' ', '(')
		b = append(b, r.err.Error()...)
		b = append(b, ')')
	}
	r.str = string(b)
	return r.str
}
