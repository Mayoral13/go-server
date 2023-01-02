package main
//CREATING A WEBSERVER IS GO
import(
	"fmt"
	"log"
	"net/http"
)
//FUNCTION TO HANDLE FORM.html
func formHandler(w http.ResponseWriter,r *http.Request){
	if err:= r.ParseForm(); err!= nil{
		fmt.Fprintf(w,"ParseForm() err: %v",err)
		return
	}
	fmt.Fprintf(w,"POST REQUEST SUCCESSFUL")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n",name);
	fmt.Fprintf(w, "Address = %s\n",address);
}
//FUNCTION TO HANDLE INDEX.html
func helloHandler(w http.ResponseWriter,r *http.Request){
if(r.URL.Path != "/hello"){
	http.Error(w,"404 not Found",http.StatusNotFound)
	return
}
if(r.Method != "GET"){
	http.Error(w,"method is not supported",http.StatusNotFound);
	return
}
fmt.Fprintf(w,"hello")
} 

//MAIN SERVER CODE
func main(){
	fileServer := http.FileServer(http.Dir("./static"));
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler);
	http.HandleFunc("/hello",helloHandler);
	fmt.Println("STARTING SERVER @ PORT 8080");
	if err := http.ListenAndServe(":8080",nil); err!=nil {
		 log.Fatal((err));
	}
	
} 