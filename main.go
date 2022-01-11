package main

import (
   "fmt"
   "github.com/youngmemo/rivercrossing/state"
)

func main() {
   // fmt.Println(state.GetQuote());
      fmt.Println(state.ViewState());
      state.Put();
      fmt.Println(state.ViewState());
}
