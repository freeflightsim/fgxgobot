package mpserver


import (

)

// This is ported from an for research 
// http://gitorious.org/fgms/fgms-0-x/blobs/master/src/flightgear/MultiPlayer/mpmessages.hxx


// magic value for messages 
const MSG_MAGIC = 0x46474653  // "FGFS"


//   protocol version 
const PROTO_VER = 0x00010001;  // 1.1


const CHAT_MSG_ID = 1
const RESET_DATA_ID = 6
const POS_DATA_ID = 7
