package main
//CREATING A WEBSERVER IS GO
import(
	"fmt"
	"log"
	"net/http"
)
func main(){
	fileServer := http.FileServer(http.Dir("./static"))
}