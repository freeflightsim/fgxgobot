package radio

import (
	"testing"
	"fmt"
)

func TestRadioAlpha( t *testing.T ){
	
	fmt.Println("TestRadioAlpha", "---------")
	
	fmt.Println("Spool Alpabet>", GetAjaxAlphabet() )
	
	fmt.Println("Test Callsigns", "-------------" )
	callsigns := []string{"unix09", "win98", "go380", "test123", "bug- 23.34", " -_$%ba"}
	for idx, callsign := range callsigns{
		fmt.Println(">", idx, "\t", callsign, "\t", Callsign2Words(callsign))
	}
	
	
}
