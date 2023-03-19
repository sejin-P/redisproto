package proto

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GenerateGoCode(protoFile string) (string, error) {
	goOutDir := "path/to/generated/files"
	err := os.MkdirAll(goOutDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	generatedFileName := strings.TrimSuffix(filepath.Base(protoFile), ".proto") + ".pb.go"
	generatedFilePath := filepath.Join(goOutDir, generatedFileName)

	cmd := exec.Command("protoc", "--go_out="+goOutDir, "--go_opt=paths=source_relative", protoFile)
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return generatedFilePath, nil
}
