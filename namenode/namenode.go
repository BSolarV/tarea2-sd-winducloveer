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
	"time"

	"github.com/BSolarV/tarea2-sd-winducloveer/protoName"
	protoNode "github.com/BSolarV/tarea2-sd-winducloveer/protoName"

	"google.golang.org/grpc"
)

/*
Definitivamente no compila! c:
No tenia el protoNode y no sabia que más hacer C: */

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
	srv := newNameNode()
	// creando instancia de servidor GRPC
	grpcServer := grpc.NewServer()

	// ProtoLogistic.RegisterProtoLogisticServiceServer(grpcServer, srv)
	protoName.RegisterProtoServiceServer(grpcServer, srv)

	// Montando servidor GRPC
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to mount GRPC server on port 9000: %v", err)
	}

}

//Book es la estructura de Libro
type Book struct {
	id            string
	bookname      string
	partsNum      string
	partsLocation []string
	timestamp     string
}

//NameNode es el Server
type NameNode struct {

	//el log es el registro de los libros
	log []Book

	// Para bloquear recursos entre hilos
	mutex sync.Mutex
}

//NewServer es el constructor del Server
func newNameNode() *NameNode {
	var srv NameNode

	return &srv
}

//ClientRequest es invocada por el cliente para leer el log
func (s *NameNode) ClientRequest(ctx context.Context, request *protoNode.ReadRequest) (*protoNode.LogData, error) {
	s.mutex.Lock()
	var paq *protoNode.LogData
	aux := ""
	msg := ""
	for _, book := range s.log {
		if book.bookname == request.GetBookname() {
			paq.Id = book.id
			paq.Timestamp = book.timestamp
			msg += book.bookname + ";" + book.partsNum + ";"
			for _, part := range book.partsLocation {
				aux += part + ","
			}
		}
	}
	s.mutex.Unlock()
	msg += aux

	paq.Message = msg[:len(msg)-1]

	return paq, nil
}

//WriteRequest es innecesaria al parecer
func (s *NameNode) WriteRequest(ctx context.Context, request *protoNode.WriteRequest) (*protoNode.Response, error) {
	return nil, nil
}

//WriteLog se usa cuando el nodo reune permisos y escribe con esta funcion
func (s *NameNode) WriteLog(ctx context.Context, packageToWrite *protoNode.LogData) (*protoNode.Empty, error) {
	s.mutex.Lock()
	var book Book

	aux := packageToWrite.GetMessage()
	aux1 := strings.Split(aux, ";")
	aux2 := strings.Split(aux1[2], ",")

	book.id = packageToWrite.GetId()
	book.bookname = aux1[0]
	book.partsNum = aux1[1]
	book.partsLocation = aux2
	book.timestamp = time.Now().Format("2006.01.02 15:04:05")

	s.log = append(s.log, book)

	s.mutex.Unlock()

	return &protoNode.Empty{}, nil

}

//CheckProposal chequea propuestas
func (s *NameNode) CheckProposal(ctx context.Context, proposal *protoNode.Proposal) (*protoNode.Response, error) {
	res := true

	return &protoNode.Response{Timestamp: time.Now().Unix(), Response: res}, nil
}
