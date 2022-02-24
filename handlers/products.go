package handlers
import "log"
import "net/http"


type Products struct {
	l *log.Logger
}

func NewProducts(l*log.Logger) *Products{
	return &Products{l}
}


func (p*Products) ServeHTTP (rw http.ResponseWriter, h *http.Request) {

}