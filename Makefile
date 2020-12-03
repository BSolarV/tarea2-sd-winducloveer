
all: datanode clientDistributed clientCentralized namenode

buildDatanode: datanode/datanode.go
	export GOROOT=/usr/local/go 
	export GOPATH=$HOME/Projects/Proj1 
	export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
ifneq ("$(wildcard $(./bin/datanode))","")
	rm -r ./bin
endif
	go build -o ./bin/datanode datanode/datanode.go

buildClientDistributed: clientDistributed/clientDistributed.go
	export GOROOT=/usr/local/go 
	export GOPATH=$HOME/Projects/Proj1 
	export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
ifneq ("$(wildcard $(./bin/clientDistributed))","")
	rm -r ./bin
endif
	go build -o ./bin/clientDistributed clientDistributed/clientDistributed.go

buildClientCentralized: clientCentralized/clientCentralized.go
	export GOROOT=/usr/local/go 
	export GOPATH=$HOME/Projects/Proj1 
	export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
ifneq ("$(wildcard $(./bin/clientCentralized))","")
	rm -r ./bin
endif
	go build -o ./bin/clientCentralized clientCentralized/clientCentralized.go

buildNamenode: namenode/namenode.go
	export GOROOT=/usr/local/go 
	export GOPATH=$HOME/Projects/Proj1 
	export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
ifneq ("$(wildcard $(./bin/namenode))","")
	rm -r ./bin
endif
	go build -o ./bin/namenode namenode/namenode.go

clearRegisters:
	rm *.txt
	rm -r bookParts
	mkdir bookParts