// Jack Schefer, began 9/3/16
// Purpose: to print out the schedule without scraping it from the website
//
package main
//
import(
   "fmt"
   "time"
   "net/http"
   "encoding/json"
   "io/ioutil"
)
//
func main() {
   data := map[string]interface{}{}
   //
   // 1. Get and print out the current date and time.
   //
   fmt.Println()
   now := time.Now()
   var tsuffix string
   var zero string = ""
   fmt.Printf("%v, %v %v\n", now.Weekday(), now.Month(), now.Day())
   if now.Hour() / 12 == 0 {
      tsuffix = "am"
   } else {
      tsuffix = "pm"
   }
   if now.Minute() < 10 {
      zero = "0"
   }
   hr := now.Hour() % 12
   if hr == 0 {
      hr = 12
   }
   fmt.Printf("Time: %v:%s%v %s\n", hr, zero, now.Minute(), tsuffix)
   fmt.Println()
   //
   // 2. Get the schedule data from Ion
   res, err := http.Get("https://ion.tjhsst.edu/api/schedule?format=json")
   check(err)
   defer res.Body.Close()
   body, err := ioutil.ReadAll(res.Body)
   check(err)
   json.Unmarshal(body, &data)
   //
   // 3. check if there's school today
   schedule_date := data["results"].([]interface{})[0].(map[string]interface{})["date"]
   today := false
   if now.Format("2006-01-02") == schedule_date {
      today = true
   }
   //
   // 4. if so, print it out
   if today {
      fmt.Println("there's school today")
   }
   title := data["results"].([]interface{})[0].(map[string]interface{})["day_type"].(map[string]interface{})["name"]
   fmt.Println(title)
   //
   blocks := data["results"].([]interface{})[0].(map[string]interface{})["day_type"].(map[string]interface{})["blocks"].([]interface{})
   for _, b := range blocks {
      name  := b.(map[string]interface{})["name"]
      start := b.(map[string]interface{})["start"]
      end   := b.(map[string]interface{})["end"]
      fmt.Printf("%s: %s - %s\n", name, start, end)
   }
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
