package crossfeed

import (
	"testing"
	"fmt"
)

func TestCrossFeedClient( t *testing.T ){
	
	fmt.Println("Testing CrossFeedClient", "---------")
	
	//var f = new(CrossFeedClient)
	
	c := make(chan CF_Reply)
	
	go Make_server_request(c)
	
	res := <-c 
	
	fmt.Println(res)
	
}

