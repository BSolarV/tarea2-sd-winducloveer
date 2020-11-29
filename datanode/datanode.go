package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"

	protoNode "github.com/BSolarV/tarea2-sd-winducloveer/protoNode"

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
	protoNode.RegisterProtoServiceServer(grpcServer, srv)

	// Montando servidor GRPC
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to mount GRPC server on port 9000: %v", err)
	}

}

// DataNode debe implementar:
//	sendChunk(ChunkPackage) returns (Empty) {}
//	writePermisions(WriteRequest) returns (Response) {}  // Ricart & Agrawala
//	updateLog(LogData) returns (Empty) {}
//	checkProporsal(Proporsal) returns (Response) {}

type DataNode struct {
	proposalQueue      []string
	writeProposalQueue []string

	// Para bloquear recursos entre hilos
	mutex sync.Mutex
}

//newDataNode es el constructor del Server
func newDataNode() *DataNode {
	var srv DataNode

	return &srv
}

//Usaremos bully algorithm de https://moodle.inf.utfsm.cl/pluginfile.php/104700/mod_resource/content/0/Chapter%204_clase3.pdf
func (*DataNode) BuildProposal() (*protoNode.Proporsal, error) {
	return &protoNode.Proporsal{}, nil
}

//checkProporsal: Coordinación entre DataNodes y entre NameNode con DataNode para verificar la propuesta
func (*DataNode) CheckProporsal(ctx context.Context, proporsal *protoNode.Proporsal) (*protoNode.Response, error) {
	response := &protoNode.Response{}
	return response, nil
}

//updateLog: Escritura de DataNode en NameNode
func (*DataNode) UpdateLog(ctx context.Context, logData *protoNode.LogData) (*protoNode.Empty, error) {
	return &protoNode.Empty{}, nil
}

//aquí se aplica ricart y agrawala
//writePermisions: Coordinación entre DataNodes para escribir en el log de NameNode
func (*DataNode) WritePermisions(ctx context.Context, writeRequest *protoNode.WriteRequest) (*protoNode.Response, error) {
	response := &protoNode.Response{}
	return response, nil
}

//sendChunk: Envia el chunk a guardar
func (*DataNode) RecieveChunk(ctx context.Context, chunkPackage *protoNode.ChunkPackage) (*protoNode.Empty, error) {
	return &protoNode.Empty{}, nil
}
