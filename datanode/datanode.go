package main

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/BSolarV/tarea2-sd-winducloveer/protoName"
	"github.com/BSolarV/tarea2-sd-winducloveer/protoNode"

	"google.golang.org/grpc"
)

var IPDIRECTIONS = map[int64]string{
	0: "localhost",
	1: "localhost",
	2: "localhost",
	3: "localhost",
}
var PORTS = map[int64]string{
	0: "9000",
	1: "9001",
	2: "9002",
	3: "9003",
}

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

	//Iniciando proceso listen para datanode
	lis, err := net.Listen("tcp", IPDIRECTIONS[int64(index)]+":"+PORTS[int64(index)])
	if err != nil {
		fmt.Print("Fail listening on " + IPDIRECTIONS[int64(index)] + ":" + PORTS[int64(index)] + ".")
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
	iWantToWrite   bool
	sinceWhenIWant int64

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
	srv.iWantToWrite = false
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
func (srv *DataNode) WritePermisions(ctx context.Context, writeRequest *protoNode.WriteRequest) (*protoNode.Response, error) {
	time.Sleep(5 * time.Second)
	standBy := false
	if srv.iWantToWrite && srv.sinceWhenIWant < writeRequest.Timestamp {
		standBy = true
	}
	if standBy {
		readyToRespond := make(chan bool)
		go func() {
			for {
				if srv.iWantToWrite == false {
					readyToRespond <- true
					break
				}
			}
		}()
		<-readyToRespond
	}
	response := &protoNode.Response{Response: true}
	return response, nil
}

//RecieveChunks: Envia el chunk a guardar
func (*DataNode) RecieveChunks(ctx context.Context, chunksPackage *protoNode.ChunksPackage) (*protoNode.Empty, error) {
	for _, oneChunk := range chunksPackage.Chunks {
		fmt.Printf("Recieved part %d of book %s.\n", oneChunk.GetNumChunkActual(), chunksPackage.GetBookName())
		fileName := "bookParts/" + chunksPackage.GetBookName() + "_" + strconv.FormatUint(uint64(oneChunk.GetNumChunkActual()), 10)
		_, err := os.Create(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// write/save buffer to disk
		ioutil.WriteFile(fileName, oneChunk.Chunk, os.ModeAppend)
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
			availableDataNodes[index] = availableDataNodes[len(availableDataNodes)-1]
			availableDataNodes[len(availableDataNodes)-1] = 0
			availableDataNodes = availableDataNodes[:len(availableDataNodes)-1]
			proposal.dict[index] = append(proposal.dict[index], i)
		}
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
	var service protoNode.ProtoServiceClient

	// Estableciendo conexiones para proposal y enviar chunks
	var conn *grpc.ClientConn
	var err error
	for i := 0; i < 3; i++ {
		if i == srv.index {
			validIndexes = append(validIndexes, i)
			connections = append(connections, nil)
			continue
		}

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

	fmt.Printf("%x\n", validIndexes)
	var dataNodeService protoNode.ProtoServiceClient
	var proposal Proposal

	var chunks []*protoNode.Chunk
	var chunksToSend *protoNode.ChunksPackage

	tempValidIndexes := make([]int, len(validIndexes))
	flag := false
	for flag == false {

		flag = true

		_ = copy(tempValidIndexes, validIndexes)
		fmt.Printf("%x\n", tempValidIndexes)
		proposal, err = srv.BuildProposal(tempValidIndexes, connections, len(splittedFile.Chunks))
		if err != nil {
			flag = false
			fmt.Printf("ERROR! %s\n", err)
			continue
		}

		for key, value := range proposal.dict {

			chunks = chunks[:0]

			for i := 0; i < len(value); i++ {
				chunks = append(chunks, &protoNode.Chunk{FileName: splittedFile.Name, NumChunkActual: int64(value[i]), Chunk: splittedFile.Chunks[value[i]]})
			}
			chunksToSend = &protoNode.ChunksPackage{BookName: splittedFile.Name, Chunks: chunks}

			if key == srv.index {
				_, err = srv.RecieveChunks(context.Background(), chunksToSend)
			} else {
				dataNodeService = protoNode.NewProtoServiceClient(connections[key])
				_, err = dataNodeService.RecieveChunks(context.Background(), chunksToSend)
			}

			if err != nil {
				flag = false
				break
			}
		}
	}

	// RICART & AGRAWALA
	fmt.Println("Beginning Ricart")
	srv.iWantToWrite = true
	srv.sinceWhenIWant = time.Now().Unix()

	done := make(chan bool)
	counter := 0
	for _, index := range validIndexes {
		if index != srv.index {
			go func(index int) {
				conn := connections[index]
				dataNodeService := protoNode.NewProtoServiceClient(conn)
				_, err = dataNodeService.WritePermisions(context.Background(), &protoNode.WriteRequest{Node: int64(srv.index), Timestamp: srv.sinceWhenIWant})
				if err != nil {
					fmt.Printf("ERROR! %s\n", err)
				}
				fmt.Printf("Recieved Response from %d\n", index)
				done <- true
			}(index)
		}
	}
	for range done {
		counter++
		if counter == len(tempValidIndexes)-1 {
			fmt.Println("Yes")
			break
		}
	}
	con, err := grpc.Dial(IPDIRECTIONS[3]+":"+PORTS[3], grpc.WithInsecure())
	if err != nil {
		fmt.Printf("ERROR! %s\n", err)
	}

	nameNodeService := protoName.NewProtoNameServiceClient(con)
	fmt.Println("sendToWrite")
	var sendToWrite protoName.LogData

	sendToWrite.BookName = splittedFile.Name
	sendToWrite.NumParts = int64(len(splittedFile.Chunks))
	sendToWrite.PartsLocation = make([]*protoName.Part, sendToWrite.NumParts)
	fmt.Println("PartsLocation")
	var ipPort string
	for nod, chnklist := range proposal.dict {
		ipPort = IPDIRECTIONS[int64(nod)] + ":" + PORTS[int64(nod)]
		fmt.Printf("ippuerto: %s", ipPort)
		for _, chnk := range chnklist {
			sendToWrite.PartsLocation = append(sendToWrite.GetPartsLocation(), &protoName.Part{Index: int64(chnk), IpPuertoDatanode: ipPort})
		}
	}
	fmt.Println("WriteLog")
	_, err = nameNodeService.WriteLog(context.Background(), &sendToWrite)

	srv.iWantToWrite = false

	for _, connection := range connections {
		if connection == nil {
			continue
		}
		connection.Close()
	}
	return &protoNode.Empty{}, nil
}

func (srv *DataNode) PrintIndex(ctx context.Context, _ *protoNode.Empty) (*protoNode.Response, error) {
	fmt.Printf("Soy el indice %d\n", srv.index)
	return &protoNode.Response{Id: strconv.Itoa(srv.index)}, nil
}

func (*DataNode) HeartBeat(ctx context.Context, Empty *protoNode.Empty) (*protoNode.Empty, error) {
	return &protoNode.Empty{}, nil
}

func (*DataNode) GetChunk(ctx context.Context, chunk *protoNode.Chunk) (*protoNode.Chunk, error) {
	fmt.Println("In GetChunk")
	fileName := "bookParts/" + chunk.FileName + "_" + strconv.FormatUint(uint64(chunk.NumChunkActual), 10)
	file, err := os.Open(fileName)
	fmt.Println("Oppened")
	if err != nil {
		fmt.Printf("Somthing went worng: \n")
		panic(err)
	}
	defer file.Close()
	partBuffer := make([]byte, 256000)
	file.Read(partBuffer)
	fmt.Println("Readed")
	chunk.Chunk = partBuffer
	fmt.Println("returning Chunk")
	return chunk, nil
}
func (srv *DataNode) CentralizedUploadFile(ctx context.Context, splittedFile *protoNode.SplittedFile) (*protoNode.Empty, error) {
	var validIndexes []int
	var connections []*grpc.ClientConn
	var service protoNode.ProtoServiceClient

	// Estableciendo conexiones para proposal y enviar chunks
	var conn *grpc.ClientConn
	var err error
	for i := 0; i < 3; i++ {
		if i == srv.index {
			validIndexes = append(validIndexes, i)
			connections = append(connections, nil)
			continue
		}

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

	fmt.Printf("%x\n", validIndexes)
	var dataNodeService protoNode.ProtoServiceClient
	var proposal Proposal

	var chunks []*protoNode.Chunk
	var chunksToSend *protoNode.ChunksPackage

	tempValidIndexes := make([]int, len(validIndexes))

	_ = copy(tempValidIndexes, validIndexes)
	fmt.Printf("%x\n", tempValidIndexes)
	proposal, err = srv.BuildProposal(tempValidIndexes, connections, len(splittedFile.Chunks))

	//Hay que enviar la propuesta al namenode
	var pToNameNode *protoName.ProposalToNameNode
	for key, value := range proposal.dict {
		pToNameNode.NumChunks = int64(len(splittedFile.Chunks))
		if key == 0 {
			for _, i := range value {
				pToNameNode.ChunksNode1 = append(pToNameNode.ChunksNode1, int64(i))
			}
		}
		if key == 1 {
			for _, i := range value {
				pToNameNode.ChunksNode2 = append(pToNameNode.ChunksNode2, int64(i))
			}
		}
		if key == 2 {
			for _, i := range value {
				pToNameNode.ChunksNode3 = append(pToNameNode.ChunksNode3, int64(i))
			}
		}
		//Enviar proposal al namenode
		//abir conexion al namenode
		var conn *grpc.ClientConn
		conn, err = grpc.Dial(IPDIRECTIONS[int64(3)]+":"+PORTS[int64(3)], grpc.WithInsecure())
		NameService := protoName.NewProtoNameServiceClient(conn)
		response, err := NameService.DistributeProposal(context.Background(), pToNameNode)
		if response.Response == false || err != nil {
			//algo anda mal
			fmt.Println("Algo anda mal")
		}

		//hacer el upload a cada nodo
		for key, value := range proposal.dict {

			chunks = chunks[:0]

			for i := 0; i < len(value); i++ {
				chunks = append(chunks, &protoNode.Chunk{FileName: splittedFile.Name, NumChunkActual: int64(value[i]), Chunk: splittedFile.Chunks[value[i]]})
			}
			chunksToSend = &protoNode.ChunksPackage{BookName: splittedFile.Name, Chunks: chunks}

			if key == srv.index {
				_, err = srv.RecieveChunks(context.Background(), chunksToSend)
			} else {
				dataNodeService = protoNode.NewProtoServiceClient(connections[key])
				_, err = dataNodeService.RecieveChunks(context.Background(), chunksToSend)
			}

			if err != nil {
				//algo anda mal
				fmt.Println("Algo anda mal")
				break
			}
		}
		//Nuevo###############################################
		fmt.Println("sendToWrite")
		var sendToWrite protoName.LogData

		sendToWrite.BookName = splittedFile.Name
		sendToWrite.NumParts = int64(len(splittedFile.Chunks))
		sendToWrite.PartsLocation = make([]*protoName.Part, sendToWrite.NumParts)
		fmt.Println("PartsLocation")
		var ipPort string
		for nod, chnklist := range proposal.dict {
			ipPort = IPDIRECTIONS[int64(nod)] + ":" + PORTS[int64(nod)]
			fmt.Printf("ippuerto: %s", ipPort)
			for _, chnk := range chnklist {
				sendToWrite.PartsLocation = append(sendToWrite.GetPartsLocation(), &protoName.Part{Index: int64(chnk), IpPuertoDatanode: ipPort})
			}
		}
		fmt.Println("WriteLog")
		_, err = NameService.WriteLog(context.Background(), &sendToWrite)

		srv.iWantToWrite = false

		for _, connection := range connections {
			if connection == nil {
				continue
			}
			connection.Close()
		}
		conn.Close()

	}

	return &protoNode.Empty{}, nil
}

//CentralizedBuildProposal es para Construir la unica propuesta que necesita para enviarsela al NameNode
func (srv *DataNode) CentralizedBuildProposal(validIndexes []int, connections []*grpc.ClientConn, numOfChunks int) (Proposal, error) {
	proposal := Proposal{dict: make(map[int][]int)}
	var index int
	var availableDataNodes []int
	var dataNodeProposal *protoNode.Proposal
	var response *protoNode.Response
	var err error

	availableDataNodes = validIndexes
	for i := 0; i < numOfChunks; i++ {
		if len(availableDataNodes) == 0 {
			availableDataNodes = validIndexes
		}
		index = rand.Intn(len(availableDataNodes))
		availableDataNodes[index] = availableDataNodes[len(availableDataNodes)-1]
		availableDataNodes[len(availableDataNodes)-1] = 0
		availableDataNodes = availableDataNodes[:len(availableDataNodes)-1]
		proposal.dict[index] = append(proposal.dict[index], i)
	}
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
			continue
		}
	}
	return proposal, nil
}
