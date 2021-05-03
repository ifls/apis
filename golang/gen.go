//go:generate go run gen.go
package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	if strings.HasSuffix(wd, "golang") {
		err = os.Chdir(strings.TrimSuffix(wd, "/golang"))
		if err != nil {
			log.Fatalln(err)
		}
	}

	wd, err = os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	if !strings.HasSuffix(wd, "apis") {
		log.Fatalln("not at apis")
	}

	err = os.Chdir(path.Join(wd, "pb"))
	if err != nil {
		log.Fatalln(err)
	}

	wd, err = os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	if !strings.HasSuffix(wd, "pb") {
		log.Fatalln("not at pb")
	}

	err = filepath.Walk(wd, func(path string, info fs.FileInfo, err error) error {
		tmp := strings.TrimPrefix(path, wd+"/")
		if info.IsDir() {
			return nil
		}

		if strings.Contains(tmp, "google") {
			return nil
		}

		//log.Println(tmp)
		return walkProto(tmp)
	})
	if err != nil {
		log.Fatalln(fmt.Errorf("wall dir err:%w", err))
	}
}

func makeGoCmd(relativePath string) *exec.Cmd {
	return exec.Command("protoc", "--proto_path=.", "--go_out=../golang", "--go_opt", "paths=source_relative", relativePath)
}

func makeGrpcCmd(relativePath string) *exec.Cmd {
	return exec.Command("protoc", "--proto_path=.", "--go-grpc_out=../golang", "--go-grpc_opt", "paths=source_relative", relativePath)
}

func makeGrpcGatewayCmd(relativePath string) *exec.Cmd {
	return exec.Command("protoc", "--proto_path=.", "--grpc-gateway_out=../golang", "--grpc-gateway_opt", "paths=source_relative", relativePath)
}

func makePwdCmd() *exec.Cmd {
	return exec.Command("pwd")
}

func walkProto(relativePath string) error {
	err := runCmd(makePwdCmd())
	if err != nil {
		return err
	}

	err = runCmd(makeGoCmd(relativePath))
	if err != nil {
		return err
	}

	err = runCmd(makeGrpcCmd(relativePath))
	if err != nil {
		return err
	}

	err = runCmd(makeGrpcGatewayCmd(relativePath))
	if err != nil {
		return err
	}

	return nil
}

func runCmd(cmd *exec.Cmd) error {
	//fmt.Printf("%s\n", cmd.String())

	body, err := cmd.CombinedOutput()

	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("%s -> %s\n", cmd.String(), body))
	}

	return nil
}
