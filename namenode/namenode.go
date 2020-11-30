package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/BSolarV/tarea2-sd-winducloveer/protoName"
	protoNode "github.com/BSolarV/tarea2-sd-winducloveer/protoName"

	"google.golang.org/grpc"
)

/*
Definitivamente no compila! c:
No tenia el protoNode y no sabia que más hacer C: */

//IPDIRECTIONS son las direcciones Ip's
var IPDIRECTIONS = map[int64]string{
	0: "localhost",
	1: "localhost",
	2: "localhost",
	3: "localhost",
}

//PORTS son los puertos en los que esas direcciones ip's escuchan
var PORTS = map[int64]string{
	0: "9000",
	1: "9001",
	2: "9002",
	3: "9003",
}

func main() {

	//reader := bufio.NewReader(os.Stdin)

	fmt.Print("Indice del presente DataNode (0, 1 o 2): ")

	//Iniciando proceso listen para datanode
	lis, err := net.Listen("tcp", IPDIRECTIONS[3]+":"+PORTS[3])
	if err != nil {
		fmt.Print("Fail listening on " + IPDIRECTIONS[3] + ":" + PORTS[3] + ".")
		panic(err)
	}
	defer lis.Close()

	// Creando instancia del nodo
	srv := newNameNode()
	// creando instancia de servidor GRPC
	grpcServer := grpc.NewServer()

	// ProtoLogistic.RegisterProtoLogisticServiceServer(grpcServer, srv)
	protoName.RegisterProtoNameServiceServer(grpcServer, srv)

	// Montando servidor GRPC
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to mount GRPC server on port 9000: %v", err)
	}

}

//Book es la estructura de Libro
type Book struct {
	id            string
	bookname      string
	partsNum      int64
	partsLocation map[int64]int64 // index -> node
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

	//var aux [len(request.partsLocation]string
	for _, book := range s.log {
		if book.bookname == request.GetBookname() {
			paq.BookName = book.bookname
			paq.NumParts = book.partsNum

			for indx, node := range book.partsLocation {
				paq.PartsLocation = append(paq.GetPartsLocation(), &protoName.Chunk{Index: indx, Datanode: node})
			}

		}
	}
	s.mutex.Unlock()
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

	book.bookname = packageToWrite.GetBookName()
	book.partsNum = packageToWrite.GetNumParts()

	for _, chunk := range packageToWrite.GetPartsLocation() {
		book.partsLocation[chunk.GetIndex()] = chunk.GetDatanode()
	}

	s.log = append(s.log, book)

	s.mutex.Unlock()

	return &protoNode.Empty{}, nil

}

//CheckProposal chequea propuestas
func (s *NameNode) CheckProposal(ctx context.Context, proposal *protoNode.Proposal) (*protoNode.Response, error) {
	res := true

	return &protoNode.Response{Timestamp: time.Now().Unix(), Response: res}, nil
}
