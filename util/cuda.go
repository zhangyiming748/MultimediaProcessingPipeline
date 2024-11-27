package util

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// CheckCUDA checks if CUDA is supported by running the nvidia-smi command.
func CheckCUDA() bool {

	cmd := exec.Command("nvidia-smi")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		fmt.Println("CUDA is not supported on this machine.")
		return false
	}

	output := out.String()
	if strings.Contains(output, "CUDA Version") {
		fmt.Println("CUDA is supported on this machine.")
		return true
	}

	fmt.Println("CUDA is not supported on this machine.")
	return false
}
