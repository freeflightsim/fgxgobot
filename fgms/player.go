

package fgms

//+ From http://gitorious.org/fgms/fgms-0-x/blobs/master/src/server/fg_player.cxx

// http://gitorious.org/fgms/fgms-0-x/blobs/master/src/server/fg_player.hxx#line36
type FG_Player struct{
	Origin string
	NetAddress string // needs to be address
	Callsign string
	Passwd string // unused ?
	ModelName string
	JoinTime date.Time
	Timestamp int64
	Error string
	HasErrors bool
	ClientId int
	PktsSentTo int
	PktsForwarded  int
    PktsReceivedFrom  int
    LastRelayedToInactive int
}

// TODO
// http://gitorious.org/fgms/fgms-0-x/blobs/master/src/server/fg_player.cxx#line71
func (me FG_Player) assign() {

}


