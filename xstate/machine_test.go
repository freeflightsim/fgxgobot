package xstate

import (
	"testing"
	"fmt"
)

func TestMachine( t *testing.T ){
	
	fmt.Println("Testing Machine", "---------")
	
	
	xStateMachine := GetStateMachine()
	
	//xStateMachine.init()
	xStateMachine.Start()
	fmt.Println("==", xStateMachine)
	
	xStateMachine2 := GetStateMachine()
	fmt.Println("==", xStateMachine2)
	
	fmt.Println("==", xStateMachine == xStateMachine2)
	
}
