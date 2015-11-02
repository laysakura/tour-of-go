package main
import (
	"net/http"
	"fmt"
)

type String string
func (h String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "RESPONSE: ", h)
}

type Struct struct {
	Greeting, Punct, Who string
}
func (h *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "RESPONSE: ", h.Greeting, h.Punct, h.Who)
}

func main()  {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	http.ListenAndServe("localhost:4000", nil)
}
