package state

import "rsc.io/quote"

func GetQuote() string {
    return quote.Hello()
}

var state = "[chicken fox grain man ---\\ \\__/____________________/---]"


func Put() {
    state = "[chicken fox grain ---\\ \\_man_/____________________/---]"

}



func ViewState() string {
   return state
}

