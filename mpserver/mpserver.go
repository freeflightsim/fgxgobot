
package mpserver

import (
	"fmt"
	//"time"
)




// Top level domain of mpnetwork
const TLD = "flightgear.org"

// Subdomain from no - zero padded eg "mpserver01"
func GetSubDomainName(no int) string {
	return fmt.Sprintf("mpserver%02d", no) 
}

// Full server name from no eg "mpserver09.flightgear.org"
func GetServerName(no int) string {
	return fmt.Sprintf( "%s.%s", GetSubDomainName(no), TLD)
}


//-------------------------------------------------------------
// The statuses of an MpServer
const STATUS_UNKNOWN string = "Unknown" 
const STATUS_DNS string = "Dns Found"
const STATUS_IP string = "IP Exists"
const STATUS_TELNET string = "Telnet Ok"
const STATUS_UP string = "Up"


//-------------------------------------------------------------


// MpServer is the record stores in the MpServersStore
type MpServer struct {
	Status string  `json:"status"`
	No int `json:"no"`
	SubDomain string  `json:"subdomain"`
	Ip string  `json:"ip"`
	//LastSeen time.Time //TODO we need to do this after telnet  
}


//===========================================
//== A Struct for spooling out with ajax
type AjaxMpServersPayload struct {
	Success bool `json:"success"`
	MpServers []*MpServer `json:"mpservers"`
}


// NewMpServer created an instance
func NewMpServer() *MpServer {
	return &MpServer{}
}



