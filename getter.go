// Jack Schefer, began 9/3/16
// Purpose: to print out the schedule without scraping it from the website
//
package main
//
import(
   "fmt"
   "net/http"
   "encoding/json"
   "io/ioutil"
)
//
func main() {
   data := map[string]interface{}{}
   //
   res, err := http.Get("https://ion.tjhsst.edu/api/schedule?format=json")
   check(err)
   defer res.Body.Close()
   //
   body, err := ioutil.ReadAll(res.Body)
   check(err)
   json.Unmarshal(body, &data)
   //
   fmt.Println(data)
}
//
///////////////////////////////////////////////////////////
//
func check(e error) {
   if e != nil {
      panic(e.Error())
   }
}
//
// End of file.
