
package mpserver

import (
	"fmt"
	//"time"
)




// Top level domain of mpnetwork
const TLD = "flightgear.org"

// Returns subdomain from a no - zero padded eg "mpserver01"
func GetSubDomainName(no int) string {
	return fmt.Sprintf("mpserver%02d", no) 
}

// Returns the FQDN from a no eg "mpserver09.flightgear.org"
func GetServerName(no int) string {
	return fmt.Sprintf( "%s.%s", GetSubDomainName(no), TLD)
}


//-------------------------------------------------------------


// Unknown is at startup
const STATUS_UNKNOWN string = "Unknown" 

// DNS entry was found
const STATUS_DNS string = "Dns Found"

// The IP Address exists - TODO for now its gonna be a telnet query
const STATUS_IP string = "IP Exists"

// The IP address exists and so does telnet
const STATUS_TELNET string = "Telnet Ok"

// Assume the server is up
const STATUS_UP string = "Up"


//-------------------------------------------------------------

// MpServer is the object pointer  in the mpserver.MpServersStore
type MpServer struct {
	Status string  `json:"status"`
	No int `json:"no"`
	Domain string  `json:"domain"`
	Ip string  `json:"ip"`
	LastTelnet string `json:"last_telnet"`
	TelnetReply string `json:"telnet_reply"`
	//LastSeen time.Time //TODO we need to do this after telnet  
}


//--------------------------------------------------------------------

// A Struct for spooling out the mpserver payload in Ajax
type AjaxMpServersPayload struct {
	Success bool `json:"success"`
	MpServers []*MpServer `json:"mpservers"`
}


// Constucts and returns a new MpServer instance
func NewMpServer() *MpServer {
	return &MpServer{}
}



