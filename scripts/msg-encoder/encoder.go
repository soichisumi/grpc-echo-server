package main

import (
	"encoding/hex"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/soichisumi/go-util/logger"
	grpctesting "github.com/soichisumi/grpc-echo-server/pkg/proto"
	"go.uber.org/zap"
	"os"
)

func main() {
	req := grpctesting.EchoRequest{
		Message:              "yo, test msg",
	}
	b, err := proto.Marshal(&req)
	if err != nil {
		logger.Fatal(err.Error(), zap.Error(err))
	}
	msg := hex.EncodeToString(b)
	frame := "00" + fmt.Sprintf("%08x", len(msg)/2) + msg
	fmt.Printf("Frame: %s\n", frame)

	f, err := os.OpenFile("request.bin", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logger.Fatal(err.Error(), zap.Error(err))
	}
	defer f.Close()

	_, err = f.Write([]byte(frame))
	if err != nil {
		logger.Fatal(err.Error(), zap.Error(err))
	}
}