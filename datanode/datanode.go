package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"

	"google.golang.org/grpc"
)

/*
Definitivamente no compila! c:
No tenia el proto y no sabia que más hacer C: */

const ipDatanode0 = "localhost"
const ipDatanode1 = "localhost"
const ipDatanode2 = "localhost"

const ipNamenode = "localhost"

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Indice del presente DataNode (0, 1 o 2): ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)
	index, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}

	var ip, port string
	switch index {
	case 0:
		ip = ipDatanode0
		port = "9000"
	case 1:
		ip = ipDatanode1
		port = "9001"
	case 2:
		ip = ipDatanode2
		port = "9002"
	default:
		panic("No es un indice valido.")
	}

	//Iniciando proceso listen para datanode
	lis, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		fmt.Print("Fail listening on " + ip + ":" + port + ".")
		panic(err)
	}
	defer lis.Close()

	// Creando instancia del nodo
	srv := newDataNode()
	// creando instancia de servidor GRPC
	grpcServer := grpc.NewServer()

	// ProtoLogistic.RegisterProtoLogisticServiceServer(grpcServer, srv)

	// Montando servidor GRPC
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to mount GRPC server on port 9000: %v", err)
	}

}

type DataNode struct {

	// Para bloquear recursos entre hilos
	mutex sync.Mutex
}

//NewServer es el constructor del Server
func newDataNode() *Server {
	var srv Server

	return &srv
}
