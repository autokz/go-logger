package logger_for_services

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/autokz/go-logger/proto"
	"google.golang.org/grpc"
)

type serviceLogger struct {
	name              string
	uuid              string
	serverAddress     string
	serverPort        string
	connectionTimeout int
}

func (sLog *serviceLogger) Send(msgType, text string) {
	conn := createGrpcConnect(sLog.serverAddress, sLog.serverPort)
	client := pb.NewSendLogsClient(conn)
	ctx := context.Background()
	currentTime := time.Now().Local().String()

	_, err := client.Send(ctx, &pb.Logs{
		Name: sLog.name,
		Uuid: sLog.uuid,
		Time: currentTime,
		Type: msgType,
		Text: text,
	})
	if err != nil {
		log.Print(err)
	}
}

func getIpAddress(domain string) string {
	ips, _ := net.LookupIP(domain)
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			return ipv4.String()
		}
	}
	return ""
}

func createGrpcConnect(address, port string) *grpc.ClientConn {
	ip := getIpAddress(address)
	conn, err := grpc.Dial(ip+":"+port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return conn
}

var privateServiceLogger *serviceLogger

func Init(name, uuid, addr, port string, t int) {
	privateServiceLogger = &serviceLogger{
		name:              name,
		uuid:              uuid,
		serverAddress:     addr,
		serverPort:        port,
		connectionTimeout: t,
	}
}

func SendLog(msgType, text string) {
	timeout := time.Duration(privateServiceLogger.connectionTimeout) * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	go func(ctx context.Context, cancel context.CancelFunc) {
		defer cancel()
		// privateServiceLogger.Send(msgType, text)
		return
	}(ctx, cancel)
}
