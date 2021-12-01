package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"path"
	"runtime"

	configs "authz-service/configs"

	authRPC "authz-service/modules/authz/handlers/rpc"
	"authz-service/package/casbinhelper"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	task = ""
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()\n", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})
	logrus.SetReportCaller(true)
	logrus.SetOutput(os.Stdout)

	if len(os.Args) > 1 {
		readFlags()
	}
}

func readFlags() {
	flag.StringVar(&task, "task", "", "Task to run: server")
}

func main() {
	// defer sentry.Flush(2 * time.Second)

	if len(os.Args) > 1 {
		executeCommand()
	} else {
		runServer()
	}
}

func executeCommand() {
	flag.Parse()

	switch task {
	case "server":
		runServer()
	default:
		fmt.Println("Unknow command:", task)
	}
}

func runServer() {
	authInstance, err := casbinhelper.InitAuthorization()
	if err != nil {
		log.Fatal().Msgf("Failed to Init Authorization: %v", err)
	}
	if configs.GrpcPort == "" {
		configs.GrpcPort = "50051"
	}
	lis, err := net.Listen("tcp", "0.0.0.0:"+configs.GrpcPort)
	if err != nil {
		log.Fatal().Msgf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	authRPC.RegisterServer(s, authInstance)

	log.Info().Msgf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal().Msgf("Failed to serve: %v", err)
	}
	r := gin.Default()

	if err := r.Run(); err != nil {
		log.Fatal().Msgf("Failed to serve http: %v", err)
	}
}
