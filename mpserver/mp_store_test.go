package mpserver

import (
	"testing"
	"fmt"
)

func TestMachine( t *testing.T ){
	
	fmt.Println("Testing MpServer", "---------")
	
	
	stateMach.MpServers = mpserver.NewMpServersStore()
	//stateMach.MpServerChan = make(chan *mpnet.MpServer) ??
	stateMach.MpServers.StartDnsTimer()
	
}
