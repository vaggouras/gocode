// My microservice!
package main
import ("github.com/ant0ine/go-json-rest/rest"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"os")
func main() {
    // try making database connection and query
    db, err := sql.Open("mysql", "user:pass@tcp(ip:port)/databasename")
    defer db.Close()

    if err != nil {
        fmt.Println("Failed to connect", err)
        return
    }
	
	var xCoordinate int
	var yCoordinate int
	err = db.QueryRow("SELECT * from Sample").Scan(&xCoordinate, &yCoordinate)
	switch {
	case err == sql.ErrNoRows:
	case err != nil:
		fmt.Println("Failed to get results" ,err)
	default:
        fmt.Println(xCoordinate, yCoordinate)
	}
	
	// try making http request against sample site
	response, err := http.Get("http://ip.jsontest.com/")
	if err != nil {
		fmt.Println("Error making GET call")
	} else {
	    defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
		  fmt.Println("Error parsing GET call")
		} else {
		  fmt.Printf("%s\n", string(contents))
		}
		
	}
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"Body": "Figher is running!"})
	}))

	fmt.Println(http.ListenAndServe(":"+os.Getenv("PORT"), api.MakeHandler()))
}
