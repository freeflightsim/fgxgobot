package xstate

import (
    "time"
)
import (
    "github.com/fgx/fgxgobot/mpserver"
    "github.com/fgx/fgxgobot/flights"
)

// The "StateMachine" is the in memory database instance 
//
// FlightsStore() and MpServers() as the "database"
// There is a single instance "Singleton" return via GetStateMachine() 
// HELP: ? is this the right way ?
//

//= Instance created in GetStateMachine() as a Singleton
// HELP: ? Is this the right way
var GStateMachine *StateMachine

//var _init_ctx sync.Once // Flag to make sure above only runs once


// StateMachine Object
type StateMachine struct {
	
	Debug bool
	DateStarted time.Time
	
	MpServers *mpserver.MpServersStore 
	MpServerChan chan *mpserver.MpServer	// NOT USED yet ?? help
	
	Flights *flights.FlightsStore
			
	
	
}

// Start() will start to DNS background timer and the CrossFeed timers WIP
func (me *StateMachine) Start() {
	me.MpServers.StartDnsTimer()
	me.Flights.StartCrossfeedTimer()
}


//= First attempt at creating a new instance - failed
func InitializeStateMachine()  *StateMachine{
	 
	GStateMachine = new(StateMachine)
	GStateMachine.DateStarted = time.Now()
	
	GStateMachine.MpServers = mpserver.NewMpServersStore()
	GStateMachine.MpServerChan = make(chan *mpserver.MpServer)
	
	
	GStateMachine.Flights = flights.NewFlightsStore()
	
	return GStateMachine
}


//===============================================================
//= Returns and created a new instance as a singleton
// HELP ? is this the right way
/*
func GetStateMachine() *StateMachine {
	_init_ctx.Do( func () { 
		GStateMachine = new(StateMachine)
		
		// Tthese below are now in main.go as they caused import circular
		//GStateMachine.DateStarted = time.Now()
		
		//GStateMachine.MpServers = mpnet.NewMpServersStore()
		//GStateMachine.MpServerChan = make(chan *mpnet.MpServer)		
		//GStateMachine.Flights = flights.NewFlightsStore()
		
		//GStateMachine.CfReplyChan = make(chan crossfeed.CF_Reply)
		//GStateMachine.CF_BackBot = crossfeed.NewBackgroundBot()
		
		
	})
    return GStateMachine
}
*/


/*
//= First attempt at creating a new instance - failed
func NewStateMachine() *StateMachine {
	 
	GStateMachine = new(StateMachine)
	GStateMachine.DateStarted = time.Now()
	
	GStateMachine.MpServers = mpnet.NewMpServersStore()
	GStateMachine.MpServerChan = make(chan *mpnet.MpServer)
	
	GStateMachine.Flights = flights.NewFlightsStore()
	
	GStateMachine.CfReplyChan = make(chan crossfeed.CF_Reply)
	GStateMachine.CF_BackBot = crossfeed.NewBackgroundBot()
	
    return GStateMachine
}
*/
