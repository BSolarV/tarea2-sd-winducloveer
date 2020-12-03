
all: datanode clientDistributed clientCentralized namenode

datanode: datanode/datanode.go
	export GOROOT=/usr/local/go 
	export GOPATH=$HOME/Projects/Proj1 
	export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
ifneq ("$(wildcard $(./bin/datanode))","")
	rm -r ./bin
endif
	go build -o ./bin/datanode datanode/datanode.go

clientDistributed: clientDistributed/clientDistributed.go
	export GOROOT=/usr/local/go 
	export GOPATH=$HOME/Projects/Proj1 
	export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
ifneq ("$(wildcard $(./bin/clientDistributed))","")
	rm -r ./bin
endif
	go build -o ./bin/clientDistributed clientDistributed/clientDistributed.go

clientCentralized: clientCentralized/clientCentralized.go
	export GOROOT=/usr/local/go 
	export GOPATH=$HOME/Projects/Proj1 
	export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
ifneq ("$(wildcard $(./bin/clientCentralized))","")
	rm -r ./bin
endif
	go build -o ./bin/clientCentralized clientCentralized/clientCentralized.go

namenode: namenode/namenode.go
	export GOROOT=/usr/local/go 
	export GOPATH=$HOME/Projects/Proj1 
	export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
ifneq ("$(wildcard $(./bin/namenode))","")
	rm -r ./bin
endif
	go build -o ./bin/namenode namenode/namenode.go

clearRegisters:
	rm *.csv