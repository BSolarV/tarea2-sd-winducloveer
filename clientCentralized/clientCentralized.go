package main

//"https://www.socketloop.com/tutorials/golang-recombine-chunked-files-example"

import (
	"bufio"
	"context"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/BSolarV/tarea2-sd-winducloveer/protoName"
	"github.com/BSolarV/tarea2-sd-winducloveer/protoNode"
	"google.golang.org/grpc"
)

//IPDIRECTIONS son las direcciones Ip's
var IPDIRECTIONS = map[int64]string{
	0: "10.10.28.63",
	1: "10.10.28.64",
	2: "10.10.28.65",
	3: "10.10.28.66",
}
var PORTS = map[int64]string{
	0: "9000",
	1: "9001",
	2: "9002",
	3: "9003",
}

var DEBUG = false

func main() {
	var name, text string

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Iniciar en modo debug? ")
	text, _ = reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)
	if text == "yes" {
		DEBUG = true
	}

	keep := make(chan bool)
	go func() {

		for {

			fmt.Println("Escoja una opción:")
			fmt.Println("	1 - Subir un Libro")
			fmt.Println("	2 - Ver Lista de Libros")
			fmt.Println("	3 - Bajar un libro")
			fmt.Println("	0 - Salir")

			fmt.Print("Elección (1, 2, 3 o 4): ")
			text, _ = reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			text = strings.Replace(text, "\r", "", -1)

			switch text {
			case "1":
				fmt.Print("Ingrese el nombre del libro: ")
				text, _ = reader.ReadString('\n')
				text = strings.Replace(text, "\n", "", -1)
				text = strings.Replace(text, "\r", "", -1)
				name = text
				fmt.Print("Ingrese el nombre del archivo a subir: ")
				text, _ = reader.ReadString('\n')
				text = strings.Replace(text, "\n", "", -1)
				text = strings.Replace(text, "\r", "", -1)
				fileToBeChunked := text
				fmt.Print("Ingrese a que DataNode subir (0, 1 o 2): ")
				text, _ = reader.ReadString('\n')
				text = strings.Replace(text, "\n", "", -1)
				text = strings.Replace(text, "\r", "", -1)
				dataNode, err := strconv.Atoi(text)
				if err != nil {
					panic(err)
				}
				if DEBUG {
					fmt.Printf("Enviando libro a: %d", time.Now().UnixNano()/int64(time.Millisecond))
				}
				sendFileCentralized(name, fileToBeChunked, dataNode)

			case "2":
				bookList := getBookListCentralized()
				for _, oneBook := range bookList {
					fmt.Printf("	%s\n", oneBook)
				}

			case "3":
				fmt.Print("Ingrese el nombre del libro a bajar: ")
				text, _ = reader.ReadString('\n')
				text = strings.Replace(text, "\n", "", -1)
				text = strings.Replace(text, "\r", "", -1)
				name = text
				rebuildFileCentralized(name)

			case "4":
				keep <- true

			default:
				fmt.Println("La opcion ingresada no es válida.")

			}
			fmt.Println()
			fmt.Println()
		}
	}()
	<-keep
}

func sendFileCentralized(name string, fileToBeChunked string, dataNode int) {
	const FILECHUNK = 256000 // 250 KB = 250 * 1024 B

	file, err := os.Open(fileToBeChunked)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fileInfo, _ := file.Stat()

	var fileSize int64 = fileInfo.Size()

	// calculate total number of parts the file will be chunked into

	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(FILECHUNK)))

	if DEBUG {
		fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)
	}

	var listOfChunks [][]byte

	for i := uint64(0); i < totalPartsNum; i++ {

		partSize := int(math.Min(FILECHUNK, float64(fileSize-int64(i*FILECHUNK))))
		partBuffer := make([]byte, partSize)

		file.Read(partBuffer)
		if DEBUG {
			fmt.Printf("	%d piece: %x\n", i, partBuffer[:5])
		}
		listOfChunks = append(listOfChunks, partBuffer)
	}

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(IPDIRECTIONS[int64(dataNode)]+":"+PORTS[int64(dataNode)], grpc.WithInsecure())
	if err != nil {
		fmt.Print("Couldn't connect:")
		panic(err)
	}
	defer conn.Close()
	dataNodeService := protoNode.NewProtoServiceClient(conn)
	splittedFile := &protoNode.SplittedFile{Name: name, Chunks: listOfChunks}
	_, err = dataNodeService.CentralizedUploadFile(context.Background(), splittedFile)
	if err != nil {
		fmt.Println("Something Went Wrong: ")
		panic(err)
	}
	fmt.Printf("Libro enviado satisfactoriamente.\n")
}

func getBookListCentralized() []string {

	conn, err := grpc.Dial(IPDIRECTIONS[3]+":"+PORTS[3], grpc.WithInsecure())
	if err != nil {
		fmt.Print("Couldn't connect: ")
		panic(err)
	}
	defer conn.Close()

	nameNodeService := protoName.NewProtoNameServiceClient(conn)
	bookList, err := nameNodeService.GetBooks(context.Background(), &protoName.Empty{})
	if err != nil {
		fmt.Println("Something Went Wrong: ")
		panic(err)
	}
	return bookList.Books
}

func rebuildFileCentralized(name string) {

	conn, err := grpc.Dial(IPDIRECTIONS[3]+":"+PORTS[3], grpc.WithInsecure())
	if err != nil {
		fmt.Print("Couldn't connect: ")
		panic(err)
	}
	if DEBUG {
		fmt.Printf("Conecting to Namenode\n")
	}
	nameNodeService := protoName.NewProtoNameServiceClient(conn)
	bookData, err := nameNodeService.ClientRequest(context.Background(), &protoName.ReadRequest{Bookname: name})
	if err != nil {
		fmt.Print("Something wetn wrong: ")
		panic(err)
	}
	conn.Close()

	newFileName := name + ".pdf"
	_, err = os.Create(newFileName)
	if err != nil {
		fmt.Print("Something went wrong: ")
		panic(err)
	}

	file, err := os.OpenFile(newFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Print("Something wetn wrong: ")
		panic(err)
	}

	// IMPORTANT! do not defer a file.Close when opening a file for APPEND mode!
	// defer file.Close()

	// just information on which part of the new file we are appending
	var writePosition int64 = 0

	var chunk *protoNode.Chunk
	var chunkBufferBytes []byte
	var socket string
	if DEBUG {
		fmt.Printf("Recorriendo BookData\n")
	}
	for j := uint64(0); j < uint64(bookData.GetNumParts()); j++ {

		for _, part := range bookData.PartsLocation {
			if part.Index == int64(j) {
				socket = part.IpPuertoDatanode
				break
			}
		}

		var conn *grpc.ClientConn
		conn, err = grpc.Dial(socket, grpc.WithInsecure())
		if err != nil {
			fmt.Print("Couldn't connect:")
			panic(err)
		}
		defer conn.Close()
		dataNodeService := protoNode.NewProtoServiceClient(conn)
		chunk, err = dataNodeService.GetChunk(context.Background(), &protoNode.Chunk{FileName: name, NumChunkActual: int64(j)})
		if err != nil {
			fmt.Println("Error!")
			panic(err)
		}

		writePosition = writePosition + 256000

		// DON't USE ioutil.WriteFile -- it will overwrite the previous bytes!
		// write/save buffer to disk
		//ioutil.WriteFile(newFileName, chunkBufferBytes, os.ModeAppend)
		chunkBufferBytes = chunk.Chunk
		_, err := file.Write(chunkBufferBytes)
		if DEBUG {
			fmt.Println("archivo Escrito")
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		file.Sync() //flush to disk

		// free up the buffer for next cycle
		// should not be a problem if the chunk size is small, but
		// can be resource hogging if the chunk size is huge.
		// also a good practice to clean up your own plate after eating

		chunkBufferBytes = nil // reset or empty our buffer
	}

	// now, we close the newFileName
	file.Close()
}
