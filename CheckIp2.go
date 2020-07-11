 package main //entry point of our executable program

 import (
  "fmt" //implements formatted I/O with functions analogous to C's printf and scanf
  "net/http" //Provides client and server implementations
  "time" //provides functionality for measuring and displaying time
  "encoding/json" // implements encoding and decoding of JSON
  "io/ioutil" ////implements some I/O utility functions
 )
type GeoIP struct {
        // The right side is the name of the JSON variable
  Ip          string  `json:"ip"`
  CountryName string  `json:"country_name""`
  RegionName  string  `json:"region_name"`
  City        string  `json:"city"`
}

var ( //creating and defining variables
  address  string
  err      error
  geo      GeoIP
  response *http.Response
  body     []byte
)


 func main() {
  http.HandleFunc("/", IndexPage)
  http.ListenAndServe(":8080", nil)
}

 func IndexPage(w http.ResponseWriter, r *http.Request) {
// Use api.apistack.com  to get a JSON response
// In place of my ip address(71.234.148.143), Please put the Ip address you got after running CheckIp.go 
response, err := http.Get("http://api.ipstack.com/71.234.148.143?access_key=6e7f24bda9ad57dbc8cdb432b095a0b3&output=json&legacy=1" + address) 
 if err != nil {
 fmt.Println(err)
  }
   

  // response.Body() is a reader type. We have to use ioutil.ReadAll() to read the data in to a byte slice(string)
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    fmt.Println(err)
  }

  // Unmarshal the JSON byte slice to a GeoIP struct
  err = json.Unmarshal(body, &geo)
  if err != nil { //if the error is not equal to zero
    fmt.Println(err) //Print error
  }


  // Everything accessible in struct now
  fmt.Fprintf(w, "Current IP address:\t" + geo.Ip + "\n") //Print the IP
  fmt.Fprintf(w, "Country Name:\t"+ geo.CountryName + "\n") //Print the Country Name
  fmt.Fprintf(w, "Region Name:\t"+ geo.RegionName + "\n") // Print the Region Name
  fmt.Fprintf(w, "City:\t"+ geo.City + "\n") //Print the City
  
 t := time.Now() // Calling the Time api
 fmt.Fprintf(w,"\n" + "Date and Time:" + t.Format(time.ANSIC)) // Printing the time and date in a set format
}
 



  
 
