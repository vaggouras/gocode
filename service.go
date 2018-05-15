package main
import ("github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"os")
func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"Body": "Figher is running!"})
	}))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), api.MakeHandler()))
	
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
}