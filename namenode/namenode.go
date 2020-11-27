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
	srv := newNameNode()
	// creando instancia de servidor GRPC
	grpcServer := grpc.NewServer()

	// ProtoLogistic.RegisterProtoLogisticServiceServer(grpcServer, srv)

	// Montando servidor GRPC
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to mount GRPC server on port 9000: %v", err)
	}

}

type Book struct {
	id            string
	bookname      string
	partsNum      string
	partsLocation []string
	timestamp     string
}

type NameNode struct {

	//el log es el registro de los libros
	log []Book
	// Para bloquear recursos entre hilos
	mutex sync.Mutex
}

//NewServer es el constructor del Server
func newNameNode() *Server {
	var srv Server

	return &srv
}

//ReadRequest es invocada por el cliente para leer el log
func (s *Server) ReadRequest(ctx context.Context, request *proto.ReadRequest) (*proto.LogData, error) {
	mutex.lock()
	var paq *proto.LogData
	for _, book := range log {
		aux := ""
		if book.bookname == request.bookname {
			paq.id = book.id
			paq.timestamp = book.timestamp
			msg := book.bookname + ";" + book.partsNum + ";"
			for _, part := range book.partsLocation {
				aux += part + ","
			}
		}
	}
	mutex.unlock()
	msg += aux

	paq.message = msg[:-1]

	return paq, nil
}

func (s *Server) WriteRequest(ctx context.Context, request *proto.EditRequest) (*proto.Response, error) {

}

//cuando el nodo reune permisos escribe con esta funcion
func (s *Server) WriteLog(ctx context.Context, packageToWrite *proto.LogData) (*proto.Empty, error) {
	mutex.lock()
	var book Book;
	book.id = packageToWrite.Getid()
	
	book.bookname = packageToWrite.Getbookname()
	book.partsNum = packageToWrite.

	log = append(log, book)
	
}
