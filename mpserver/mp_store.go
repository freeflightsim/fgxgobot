
package mpserver

import (
	"fmt"
	"time"
	"net"
	"encoding/json"
)


// The max dns to lookup 
const MAX_DNS_SERVER = 30


//--------------------------------------------------------------------

// NewMpServersStore constructor
func NewMpServersStore() *MpServersStore {
	ob := new(MpServersStore)
	ob.MpServers =  make(map[int]*MpServer)
	return ob
}


//= MpServersStore is a memory database
type MpServersStore struct {
	MpServers map[int]*MpServer
}

//= Starts the timer to lookup the DNS periodically
func (me *MpServersStore) StartDnsTimer() {

	// We need to check every so often..  for test every 60 secs
	ticker := time.NewTicker(time.Millisecond * 60000)
    go func() {
    	me.DoDnsScan()
		for _ = range ticker.C {	
			me.DoDnsScan()
		}
	}()
}


// DoDnsScan - Scans for Mp Servers with DNS range 1 - n
// BUG(pete): these need to be fired off not all at once maybe intervals of 2 seconds
func (me *MpServersStore) DoDnsScan() {

	fmt.Println(">> DoDnsScan")	
	for i := 0; i < MAX_DNS_SERVER; i++ {
		go me.DnsLookupServer(  i  )
	}
}

//= DnsLookupServer -  return the ip address or error
// This Creates/Updates an MpServer object in the Store
//
// ? Help ? is this the right way to do it in go ?
func (me *MpServersStore) DnsLookupServer(no int) {
	
	fqdn := GetServerName(no)
	//fmt.Println("Start>> :", fqdn)
	
	addrs, err := net.LookupHost(fqdn)
	if err != nil {
		//fmt.Println(" <<Lookup ERR: ", fqdn)	
		//panic(err)
		fmt.Println(" << Dns NO: ", fqdn)
		return	
	}
	
	//= Check if MpServer exists in Store map
	Mp, ok := me.MpServers[no]
	if !ok {
		// No entry for this server no so create one
		Mp = new(MpServer)
		Mp.Status = STATUS_DNS
		me.MpServers[no] = Mp
	}
	Mp.No = no
	Mp.SubDomain = GetSubDomainName(no)
	
	//= TODO ring bells if changed
	Mp.Ip = addrs[0]
	
	fmt.Println(" << Dns Ok: ", fqdn)	
}


//= GetAjaxPayload - spools out the /mpservers as json string 
// - this is send to client whether ajax request or websocket
func (me *MpServersStore) GetAjaxPayload() string {

    // Create new payload  MpServers as Array instead of Map
    var pay = new(AjaxMpServersPayload)
    pay.Success = true
    pay.MpServers = make([]*MpServer,0)
    
    for _, ele := range me.MpServers {
    	pay.MpServers = append(pay.MpServers, ele)
   		//fmt.Println("IODX()",  idx, ele)
    }
    
    s, _ := json.Marshal(pay)
    
    return string(s)
}
