package radio

// Package radio contains all the stuff to do with radio telephony in sim
// Maybe this could be a service on appengine ?
//
//   - http://en.wikipedia.org/wiki/NATO_phonetic_alphabet
//   - http://en.wikipedia.org/wiki/International_Phonetic_Alphabet#History

import (
	"fmt"
	"sort"
	"strings"
	"encoding/json"
)


//================================================
// Definition
//================================================
type Alpha struct{
	Char string `json:"char"`
	Word string `json:"word"`
	Phonetic string `json:"phonetic"`
	Morse string `json:"morse"`
}

//= Create the map
// BUG(perdo) - this needs to be icao

//= maybe this could be in and the byte no of "character" is index == restricted ;-)
var AlphabetICAO = map[string]*Alpha {

	"-": {"-", "dash", "?", ""},
	".": {".", "dot", "DOT", ""},
	
	
	"0": {"0", "zero", "ZEE-RO", ""},
	"1": {"1", "one", "WUN", ""},
	"2": {"2", "two", "TOO", ""},
	"3": {"3", "three", "TREE", ""},
	"4": {"4", "four", "FOW-ER", ""},
	"5": {"5", "five", "FIFE", ""},
	"6": {"6", "six", "SIX", ""},
	"7": {"7", "seven", "SEV-EN", ""},
	"8": {"8", "eight", "AIT", ""},
	"9": {"9", "niner", "NIN-ER", ""},
	
	
	"A": {"A", "alfa", "AL-FA", ".-"},
	"B": {"B", "bravo", "BRAH-VOH", "-..."},
	"C": {"C", "charlie", "SHAR-LEE", "-.-."},
	"D": {"F", "delta", "DELL-TAH", ""},
	"E": {"E", "echo", "ECK-OH", ""},
	
	"F": {"F", "foxtrot", "", ""},	
	"G": {"G", "golf", "", ""},
	"H": {"H", "hotel", "", ""},
	"I": {"I", "india", "", ""},
	"J": {"J", "juliet", "", ""},
	
	"K": {"K", "kilo", "", ""},
	"L": {"L", "lima", "", ""},
	"M": {"M", "mike", "", ""},
	"N": {"N", "november", "", ""},
	"O": {"O", "oscar", "", ""},
	
	"P": {"P", "papa", "", ""},
	"Q": {"Q", "quebec", "", ""},
	"R": {"R", "romeo", "", ""},
	"S": {"S", "sierra", "", ""},
	"T": {"T", "tango", "", ""},
	
	"U": {"U", "uniform", "OO-NEE-FORM", "..-"},
	"V": {"V", "victor", "", ""},
	"W": {"W", "whiskey", "", ""},
	"X": {"X", "xray", "", ""},
	"Y": {"Y", "yankee", "", ""},
	"Z": {"Z", "zulu", "", ""},
}



//---------------------------------------------------------------
// Ajax
//--------------------------------------------------------------
type AjaxRadioAlphabetPayload struct{
	Success bool `json:"success"`
	Alphabet[] *Alpha `json:"alphabet"`
}

//= Spool out the ajax
// ? help ? theres got be a way to make a "constant" variable global This feels wrong !!
//           in this mod? there must be a better way.. AND sorted keys BTW
func GetAjaxAlphabet() string{
	payload := new(AjaxRadioAlphabetPayload)
	payload.Success = true
	payload.Alphabet = make([]*Alpha,0 )
	
	//+ now we have to sort
	sort_idx := make([]string, 0, len(AlphabetICAO))
    for i := range AlphabetICAO {
        sort_idx = append(sort_idx, i)
    }
    sort.Strings(sort_idx)
    for _, i := range sort_idx {
    	//var ajF = NewAjaxFlight(me.Flights[i])
    	payload.Alphabet = append(payload.Alphabet, AlphabetICAO[i])
    }

 	s, _ := json.MarshalIndent(payload, "" , "  ")

    return string(s)
}


//= Will return the words from a callsing eg BA19 returns "bravo alfa zero niner"
func Callsign2Words  (callsign string) string {
	

	// need to split up words
	words := make([]string, 0)
	for _, c := range callsign{
		C := strings.ToUpper(string(c))
		alpha, ok := AlphabetICAO[C]
		if ok  {
			//fmt.Println("c=", idx, string(c), alpha.Word)
			words = append(words, alpha.Word)
		}else{
			fmt.Println("NOT FOUND",  C)
			
		}
	}	
	final_str := strings.Join(words, " ")	
	return final_str
}