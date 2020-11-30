package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

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
	srv := newDataNode(index)
	// creando instancia de servidor GRPC
	grpcServer := grpc.NewServer()

	// ProtoLogistic.RegisterProtoLogisticServiceServer(grpcServer, srv)
	protoNode.RegisterProtoServiceServer(grpcServer, srv)

	// Montando servidor GRPC
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to mount GRPC server on port 9000: %v", err)
	}

}

// DataNode debe modificar los metodos que utilice:
//	sendChunk(ChunkPackage) returns (Empty) {}
//	writePermisions(WriteRequest) returns (Response) {}  // Ricart & Agrawala
//	updateLog(LogData) returns (Empty) {}
//	checkProporsal(Proporsal) returns (Response) {}

type DataNode struct {
	index          int
	availableSpace int

	proposalQueue        []string
	writePermisionsQueue []string

	// Para bloquear recursos entre hilos
	mutex sync.Mutex
}

//newDataNode es el constructor del Server
func newDataNode(actualIndex int) *DataNode {
	var srv DataNode
	srv.index = actualIndex
	srv.availableSpace = 1024 // espacio para chunks, 256 MegaBytes aproximadamente
	return &srv
}

//checkProporsal: Coordinación entre DataNodes y entre NameNode con DataNode para verificar la propuesta
func (srv *DataNode) CheckProposal(ctx context.Context, proposal *protoNode.Proposal) (*protoNode.Response, error) {
	res := true
	srv.mutex.Lock()
	if srv.availableSpace-int(proposal.GetNumChunks()) < 0 {
		res = false
	}
	srv.mutex.Unlock()
	return &protoNode.Response{Timestamp: time.Now().Unix(), Response: res}, nil
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

//RecieveChunks: Envia el chunk a guardar
func (*DataNode) RecieveChunks(ctx context.Context, chunksPackage *protoNode.ChunksPackage) (*protoNode.Empty, error) {
	for _, oneChunk := range chunksPackage.Chunks {
		fmt.Printf("Recieved part %d of book %s.\n", oneChunk.GetNumChunkActual(), chunksPackage.GetBookName())
	}
	return &protoNode.Empty{}, nil
}

type Proposal struct {
	dict map[int][]int // mapa de indice DataNode a indices de Chunks a enviar a dicho dataNode
}

// Usaremos bully algorithm de https://moodle.inf.utfsm.cl/pluginfile.php/104700/mod_resource/content/0/Chapter%204_clase3.pdf
// Creo que no hace falta el bully, mejor random, porque no se puede pasar el mando, la propuesta y envio de chunks debe hacerlo el mismo datanode que los recibio
func (srv *DataNode) BuildProposal(validIndexes []int, connections []*grpc.ClientConn, numOfChunks int) (Proposal, error) {
	proposal := Proposal{dict: make(map[int][]int)}
	var index int
	var availableDataNodes []int
	var dataNodeProposal *protoNode.Proposal
	var response *protoNode.Response
	var err error

	flag := false
	for flag == false {

		flag = true

		availableDataNodes = validIndexes
		for i := 0; i < numOfChunks; i++ {
			if len(availableDataNodes) == 0 {
				availableDataNodes = validIndexes
			}
			index = rand.Intn(len(availableDataNodes))
			fmt.Printf("Picked %d\n", index)
			availableDataNodes[index] = availableDataNodes[len(availableDataNodes)-1]
			availableDataNodes[len(availableDataNodes)-1] = 0
			availableDataNodes = availableDataNodes[:len(availableDataNodes)-1]

			fmt.Printf("Pickeeed %d\n", index)
			proposal.dict[index] = append(proposal.dict[index], i)
		}
		fmt.Println("++++++++++++++++++++++++++++")
		for key, value := range proposal.dict {
			fmt.Printf("	datanode %d\n", key)
			for i := 0; i < len(value); i++ {
				fmt.Printf("		parte %d\n", value[i])
			}
		}
		fmt.Println("++++++++++++++++++++++++++++")
		var dataNodeService protoNode.ProtoServiceClient
		for key, value := range proposal.dict {
			dataNodeProposal = &protoNode.Proposal{Node: int64(srv.index), NumChunks: int64(len(value)), Timestamp: time.Now().Unix()}
			if key == srv.index {
				response, err = srv.CheckProposal(context.Background(), dataNodeProposal)
			} else {
				dataNodeService = protoNode.NewProtoServiceClient(connections[key])
				response, err = dataNodeService.CheckProposal(context.Background(), dataNodeProposal)
			}
			if err != nil || !response.Response {
				validIndexes[key] = validIndexes[len(validIndexes)-1]
				validIndexes[len(validIndexes)-1] = 0
				validIndexes = validIndexes[:len(validIndexes)-1]
				flag = false
				continue
			}
		}
	}
	return proposal, nil
}

// El plan era establecer conexion y mantenerla hasta enviar los paquetes, pero la forma para guardar los servicios no esta funcionando.
// Usar punteros a interfaces por algun motivo es pesima idea segun stackoverflow
// El otro plan sería guardar los servicios como un atributo del server.
// El otro plan sería guardar solamente las conexiones, montar los servicios a medida que se requieran.
func (srv *DataNode) UploadFile(ctx context.Context, splittedFile *protoNode.SplittedFile) (*protoNode.Empty, error) {
	fmt.Printf("%s: %x\n", splittedFile.Name, len(splittedFile.Chunks))

	var validIndexes []int
	var connections []*grpc.ClientConn

	// Estableciendo conexiones para proposal y enviar chunks
	var ip, port string
	var conn *grpc.ClientConn
	var err error
	for i := 0; i < 3; i++ {
		if i == srv.index {
			validIndexes = append(validIndexes, i)
			connections = append(connections, nil)
			continue
		}
		switch i {
		case 0:
			ip = ipDatanode0
			port = "9000"
		case 1:
			ip = ipDatanode1
			port = "9001"
		case 2:
			ip = ipDatanode2
			port = "9002"
		}

		conn, err = grpc.Dial(ip+":"+port, grpc.WithInsecure())
		if err != nil {
			// No se añade su indice como uno valido
			connections = append(connections, nil)
			continue
		}

		validIndexes = append(validIndexes, i)
		connections = append(connections, conn)
	}

	var dataNodeService protoNode.ProtoServiceClient
	var proposal Proposal

	var chunks []*protoNode.Chunk
	var chunksToSend *protoNode.ChunksPackage

	flag := false
	for flag == false {

		flag = true

		proposal, err = srv.BuildProposal(validIndexes, connections, len(splittedFile.Chunks))
		if err != nil {
			flag = false
			fmt.Printf("ERROR! %s\n", err)
			continue
		}

		for key, value := range proposal.dict {

			fmt.Printf("Seding to %d\n", key)
			chunks = chunks[:0]

			for i := 0; i < len(value); i++ {
				chunks = append(chunks, &protoNode.Chunk{FileName: splittedFile.Name, NumChunkActual: int64(value[i]), Chunk: splittedFile.Chunks[value[i]]})
				fmt.Printf("	added chunk n: %d\n", chunks[i].NumChunkActual)
			}
			chunksToSend = &protoNode.ChunksPackage{BookName: splittedFile.Name, Chunks: chunks}

			if key == srv.index {
				_, err = srv.RecieveChunks(context.Background(), chunksToSend)
			} else {
				dataNodeService = protoNode.NewProtoServiceClient(connections[key])
				fmt.Println("-----------------------")
				asdf, fdsa := dataNodeService.PrintIndex(context.Background(), &protoNode.Empty{})
				if fdsa != nil {
					panic(fdsa)
				}
				fmt.Printf("sended to: %s\n", asdf.Id)
				fmt.Println("-----------------------")
				_, err = dataNodeService.RecieveChunks(context.Background(), chunksToSend)
			}

			if err != nil {
				flag = false
				break
			}
			fmt.Printf("DONE WITH DATANODE: %d!!!!!!!!\n", key)
		}
	}
	for index, connection := range connections {
		fmt.Printf("Closing %d\n", index)
		if connection == nil {
			continue
		}
		fmt.Printf("Esta así: %d\n", connection.GetState())
		connection.Close()
	}
	return &protoNode.Empty{}, nil
}

func (srv *DataNode) PrintIndex(ctx context.Context, _ *protoNode.Empty) (*protoNode.Response, error) {
	fmt.Printf("Soy el indice %d\n", srv.index)
	return &protoNode.Response{Id: strconv.Itoa(srv.index)}, nil
}
