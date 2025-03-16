package curl

import (
	"io"
	"os/exec"
)

func Call(url string, flags []string, values []string) (string, error) {
	if len(flags) != len(values) {
		return "", io.ErrUnexpectedEOF // TODO: change to custom error (ValueError-like)
	}

	n := len(flags)
	args := make([]string, 0, 2*n+1)
	args = append(args, "curl")
	for i := 0; i < n; i++ {
		args = append(args, flags[i], values[i])
	}
	args = append(args, url)

	cmd := exec.Command(args[0], args[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(out), nil
}