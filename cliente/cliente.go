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

	"github.com/BSolarV/tarea2-sd-winducloveer/protoName"
	"github.com/BSolarV/tarea2-sd-winducloveer/protoNode"
	"google.golang.org/grpc"
)

var IPDIRECTIONS = map[int64]string{
	0: "localhost",
	1: "localhost",
	2: "localhost",
}
var PORTS = map[int64]string{
	0: "9000",
	1: "9001",
	2: "9002",
}

func main() {
	keep := make(chan bool)
	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("Escoja una opción:")
			fmt.Println("	1 - Subir un Libro")
			fmt.Println("	2 - Ver Lista de Libros")
			fmt.Println("	3 - Bajar un libro")
			fmt.Println("	0 - Salir")

			fmt.Print("Elección (1, 2, 3 o 4): ")
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			text = strings.Replace(text, "\r", "", -1)

			switch text {
			case "1":
				fmt.Print("Ingrese el nombre del libro: ")
				text, _ := reader.ReadString('\n')
				text = strings.Replace(text, "\n", "", -1)
				text = strings.Replace(text, "\r", "", -1)
				name := text
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
				sendFile(name, fileToBeChunked, dataNode)
			case "2":

			case "3":

			case "4":
				keep <- true
			default:
				fmt.Println("La opcion ingresada no es válida.")
			}

		}
	}()
	<-keep
}

func sendFile(name string, fileToBeChunked string, dataNode int) {
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

	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

	var listOfChunks [][]byte

	for i := uint64(0); i < totalPartsNum; i++ {

		partSize := int(math.Min(FILECHUNK, float64(fileSize-int64(i*FILECHUNK))))
		partBuffer := make([]byte, partSize)

		file.Read(partBuffer)

		fmt.Printf("	%d piece: %x\n", i, partBuffer[:5])
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
	_, err = dataNodeService.UploadFile(context.Background(), splittedFile)
	if err != nil {
		fmt.Println("Something Went Wrong: ")
		panic(err)
	}
	fmt.Printf("Sumbitted!\n")
}

func getBookList() []string {

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

func rebuildFile(totalPartsNum uint64) {
	// just for fun, let's recombine back the chunked files in a new file

	newFileName := "NEWbigfile.zip"
	_, er := os.Create(newFileName)

	if er != nil {
		fmt.Println(er)
		os.Exit(1)
	}

	//set the newFileName file to APPEND MODE!!
	// open files r and w

	file, err := os.OpenFile(newFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// IMPORTANT! do not defer a file.Close when opening a file for APPEND mode!
	// defer file.Close()

	// just information on which part of the new file we are appending
	var writePosition int64 = 0

	for j := uint64(0); j < totalPartsNum; j++ {

		//read a chunk
		currentChunkFileName := "bigfile_" + strconv.FormatUint(j, 10)

		newFileChunk, err := os.Open(currentChunkFileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer newFileChunk.Close()

		chunkInfo, err := newFileChunk.Stat()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// calculate the bytes size of each chunk
		// we are not going to rely on previous data and constant

		var chunkSize int64 = chunkInfo.Size()
		chunkBufferBytes := make([]byte, chunkSize)

		fmt.Println("Appending at position : [", writePosition, "] bytes")
		writePosition = writePosition + chunkSize

		// read into chunkBufferBytes
		reader := bufio.NewReader(newFileChunk)
		_, err = reader.Read(chunkBufferBytes)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// DON't USE ioutil.WriteFile -- it will overwrite the previous bytes!
		// write/save buffer to disk
		//ioutil.WriteFile(newFileName, chunkBufferBytes, os.ModeAppend)

		n, err := file.Write(chunkBufferBytes)

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

		fmt.Println("Written ", n, " bytes")

		fmt.Println("Recombining part [", j, "] into : ", newFileName)
	}

	// now, we close the newFileName
	file.Close()
}
