package main //entry point of our executable program
import (
    "flag"// Implements command line flag parsing
    "io"  //Provides basic interface to I/O primitives
    "net/http" //Provides client and server implementations 
    "os" //provides a platform-independent interface to operating system functionality
    "fmt" // implements formatted I/O with functions analogous to C's printf and scanf
    "io/ioutil" //implements some I/O utility functions
)
    
func main() {
  http.HandleFunc("/", IndexPage) //Calling the Function Index Page
  http.ListenAndServe(":8080", nil) //Listen at port 8080
}

 func IndexPage(w http.ResponseWriter, r *http.Request) { //Writes the response and reads the request from handler
    flag.Parse() //Parsing flag
     resp, err := http.Get("http://myexternalip.com/raw") //Taking the Response from the link
     htmlData, err := ioutil.ReadAll(resp.Body) //Reading the response and converting it to string form
     fmt.Fprintf(w, "Current IP Address:\t"+ string(htmlData) + "\n") //Printing the Ip address on localhost
    if err != nil { //error 
        os.Stderr.WriteString(err.Error())
        os.Stderr.WriteString("\n")
        os.Exit(1)
    }
    io.Copy(os.Stdout, resp.Body)// Copies from the body to stdout
}



