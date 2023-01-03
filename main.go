//FIRST LINE
package main
//IMPORT ALL THE PACKAGES
import(
	"fmt"//BASICALLY A CONSOLE.LOG
	"log"//TO LOG OUT ERRORS AND STUFF
	"net/http"//HTTP METHOD
)
//FUNCTION TO HANDLE THE INDEX.html FILE
func indexHandler(w http.ResponseWriter, r *http.Request){//BASICALLY A (RES,REQ) => LIKE IN NODEJS
	if(r.URL.Path != "/hello"){//CONDITIONAL FOR PATH
     http.Error(w,"NOT FOUND",http.StatusNotFound);//ERROR CODE
	}
	if(r.Method != "GET"){//CONDITIONAL FOR THE METHOD
		http.Error(w,"NOT FOUND",http.StatusNotFound);
	}
	fmt.Fprintf(w,"HELLO")//PRINTS OUT HELLO ON THE PAGE
	
}
func formHandler(w http.ResponseWriter, r *http.Request){
	if err:= r.ParseForm();err != nil{//TO PARSE OUT DATA FROM THE FORM
		fmt.Fprintf(w,"ERROR AT FORM  %v",err)
	}
	fmt.Printf("POST REQUEST SUCCESSFUL\n");
	name := r.FormValue("name");//READING OUT NAME VALUE FROM THE FORM
	address := r.FormValue("address");//READING OUT ADDRESS VALUE FROM THE FORM
	fmt.Fprintf(w,"NAME IS %v\n",name);//PRINTS OUT THE NAME VALUE
	fmt.Fprintf(w,"ADDRESS IS %v",address);//PRINTS OUT THE ADDRESS VALUE

}
func main(){
fileServer := http.FileServer(http.Dir("./static"));//INITIALIZE A VARIABLE FILESERVER WITH THE STATIC FILES TO BE SERVE
http.Handle("/",fileServer);//USES THE PATH + FILES IE(APP.USE("/",API)) --- EXPRESS JS
http.HandleFunc("/hello",indexHandler);//HANDLES THE PATH /HELLO 
http.HandleFunc("/form",formHandler);//HANDLES THE PATH FORM
fmt.Printf("SERVER IS LIVE @ PORT 3000");
if err:= http.ListenAndServe(":3000",nil);err!=nil{//USES A PORT 
	log.Fatal((err))
}


}