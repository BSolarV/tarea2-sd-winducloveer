package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/BSolarV/tarea2-sd-winducloveer/protoName"
	"github.com/BSolarV/tarea2-sd-winducloveer/protoNode"

	"google.golang.org/grpc"
)

/*
Definitivamente no compila! c:
No tenia el protoName y no sabia que más hacer C: */

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

	//Iniciando proceso listen para Namenode
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
	partsLocation map[int64]string // index -> node
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
func (s *NameNode) ClientRequest(ctx context.Context, request *protoName.ReadRequest) (*protoName.LogData, error) {
	s.mutex.Lock()
	paq := protoName.LogData{}
	for _, book := range s.log {

		if book.bookname == request.GetBookname() {
			paq.BookName = book.bookname
			paq.NumParts = book.partsNum
			paq.PartsLocation = make([]*protoName.Part, 0)
			for indx, node := range book.partsLocation {
				fmt.Printf("node: %s\n", node)
				paq.PartsLocation = append(paq.PartsLocation, &protoName.Part{Index: indx, IpPuertoDatanode: node})
			}
			for _, elemento := range paq.PartsLocation {
				fmt.Printf("indice: %d; ipPuerto: %s\n", elemento.GetIndex(), elemento.GetIpPuertoDatanode())
			}
		}

	}
	s.mutex.Unlock()

	return &paq, nil
}

//WriteRequest es innecesaria al parecer
func (s *NameNode) WriteRequest(ctx context.Context, request *protoName.WriteRequest) (*protoName.Response, error) {
	return nil, nil
}

//WriteLog se usa cuando el nodo reune permisos y escribe con esta funcion
func (s *NameNode) WriteLog(ctx context.Context, packageToWrite *protoName.LogData) (*protoName.Empty, error) {
	s.mutex.Lock()
	var book Book
	fmt.Println("Writing Log")

	book.bookname = packageToWrite.GetBookName()
	book.partsNum = packageToWrite.GetNumParts()
	book.partsLocation = make(map[int64]string)

	for _, chunk := range packageToWrite.GetPartsLocation() {
		book.partsLocation[chunk.GetIndex()] = chunk.GetIpPuertoDatanode()
	}

	s.log = append(s.log, book)
	fmt.Println("LOG")
	for _, i := range s.log {
		fmt.Printf("%s: ", i.bookname)
		fmt.Printf("   %d:  ", i.partsNum)
		fmt.Printf("   %x: ", i.partsLocation)
		fmt.Println("---------------------------")
	}
	fmt.Println("")

	//Codigo para escribir el log
	file, err := os.OpenFile("Log.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		file, err = os.Create("Log.txt")
		if err != nil {
			panic(err)
		}
	}
	defer file.Close()
	stringToWrite := ""
	stringToWrite += book.bookname + " " + strconv.Itoa(int(book.partsNum)) + "\n"
	for key, value := range book.partsLocation {
		stringToWrite += "Parte: " + strconv.Itoa(int(key)) + "IpMaquina: " + value + "\n"
	}

	_, err = file.WriteString(stringToWrite)

	s.mutex.Unlock()

	return &protoName.Empty{}, nil

}

//DistributeProposal reparte la propuesta del datanode correspondiente
func (s *NameNode) DistributeProposal(ctx context.Context, proposal *protoName.ProposalToNameNode) (*protoName.Response, error) {
	res := true

	//Se conecta a todos los nodos y pregunta
	var validIndexes []int
	var connections []*grpc.ClientConn
	var service protoNode.ProtoServiceClient

	// Estableciendo conexiones para proposal y enviar chunks
	var conn *grpc.ClientConn
	var err error
	for i := 0; i < 3; i++ {
		conn, err = grpc.Dial(IPDIRECTIONS[int64(i)]+":"+PORTS[int64(i)], grpc.WithInsecure())
		if err != nil {
			// No se añade su indice como uno valido
			connections = append(connections, nil)
			fmt.Printf("ERROR! %s\n", err)
			continue
		}

		service = protoNode.NewProtoServiceClient(conn)
		_, err = service.HeartBeat(context.Background(), &protoNode.Empty{})
		if err != nil {
			// No se añade su indice como uno valido
			connections = append(connections, nil)
			fmt.Printf("ERROR! %s\n", err)
			continue
		}

		fmt.Printf("Indice: %d, estado: %s", i, conn.GetState())

		validIndexes = append(validIndexes, i)
		connections = append(connections, conn)
	}
	//Enviar a cada nodo la propuesta

	//finaliza

	//luego de recibir respuestas evalúa, y luego reenvía repuestao crea una nueva propuesta
	return &protoName.Response{Timestamp: time.Now().Unix(), Response: res}, nil
}

func (s *NameNode) GetBooks(ctx context.Context, empty *protoName.Empty) (*protoName.EveryBook, error) {
	var books []string
	for _, book := range s.log {
		books = append(books, book.bookname)
	}
	return &protoName.EveryBook{Books: books}, nil
}
