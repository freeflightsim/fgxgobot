
package mpserver

import (
	"fmt"
	"time"
	"net"
	"encoding/json"
)


// The max dns no to lookup 
const MAX_DNS_SERVER = 25


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


// DoDnsScan() -  Starts a scans for Mp Servers with dns range 1 - MAX_MP_SERVERS
// 
// BUG(pete): these need to be fired off not all at once maybe intervals of a few seconds
func (me *MpServersStore) DoDnsScan() {

	fmt.Println(">> DoDnsScan")	
	for i := 1; i < MAX_DNS_SERVER; i++ {
		go me.DnsLookupServer(  i  )
	}
}


// DnsLookupServer - lookup the ip address, and if it exists
// t then creates or updates an MpServer object in the mpserver.MpServersStore
//
// BUG(pete): help -  is this the right way to do it in go ?
func (me *MpServersStore) DnsLookupServer(no int) {
	
	fqdn := GetServerName(no)
	
    //fmt.Println("Start>> :", fqdn)
    //= Check if MpServer exists in Store map
    Mp, ok := me.MpServers[no]
    if !ok {
        // No entry for this server no so create one
        Mp = new(MpServer)
        
        me.MpServers[no] = Mp
    }
    Mp.No = no
    Mp.Domain = fqdn
    
	addrs, err := net.LookupHost(fqdn)
	if err != nil {
		//fmt.Println(" <<Lookup ERR: ", fqdn)	
		//panic(err)
        Mp.LastErrMsg = "DNS Lookup Failed"
		fmt.Println(" << No DNS: ", fqdn)
		return	
	}
	

	
	//= TODO ring bells if changed
	Mp.Ip = addrs[0]
	
	
	//fmt.Println(" << Dns Ok: ", fqdn)	
    
    
    tcp_server := fqdn + ":5001"
    tcpAddr, err := net.ResolveTCPAddr("tcp", tcp_server)
    
    //=============================
    // Make Telnet Connection
    telnet_start := time.Now().UTC().UnixNano()
    telnetConn, err := net.DialTCP("tcp", nil, tcpAddr)
    defer telnetConn.Close()
    
    if err != nil {
        println("Telnet failed:", err.Error())
        Mp.LastErrMsg = "Telnet Failed"
        
    }else{
        reply := make([]byte, 2048)
        
        _, err = telnetConn.Read(reply)
        
        //Mp.TelnetReply = string(reply)
        
        telnet_end := time.Now().UTC()
        Mp.LastTelnet = telnet_end.Format(time.RFC3339) 
        //println("reply from server=", string(reply))
        //println("timers=", telnet_start)
        
        //println("timers=", telnet_end)
        
        diff_nano := telnet_end.UnixNano() - telnet_start
        diff_ms := diff_nano / 1000000 
        Mp.TelnetLag = diff_ms
    
        //println("diff=", diff_ms)
    }
    
    
    //=================================================
    //  Make UDP connection
    udp_server := fqdn + ":5004"
    udp_addr, err := net.ResolveUDPAddr("udp", udp_server)
    if err != nil {
        //conn.Active = false
        Mp.LastErrMsg = "Could not resolve UDP address"
        println("Could not resolve UDP address")
        //log.Println("\tFAIL: Crossfeed to resolve UDP address:  ", udp_server, err)
        return
    }
    
    //= open socket and listen
    udp_start := time.Now().UTC().UnixNano()
    udpConn, err_listen := net.DialUDP("udp", nil,  udp_addr)
    defer udpConn.Close()
    if err_listen != nil {
        //conn.Active = false
        //conn.LastError = "Couldnt open UDP port"
        Mp.LastErrMsg = "Could not open UDP port"
        println("Could not open UDP port", err_listen.Error())
        //log.Println("\tFAIL: Crossfeed FAIL to Open:  ", conn.Url, udp_addr, err_listen)
        return
    }else{
        udp_end := time.Now().UTC()
        Mp.LastUDP = udp_end.Format(time.RFC3339) 
        Mp.UDPLag = (udp_end.UnixNano() - udp_start) / 1000
    }
    //println("Yes Udp")
        
    
    
}


// GetAjaxPayload returns the MpServers data as a json string.
// This can then be sent to client whether ajax request or websocket whatever
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


