package main
import ("github.com/ant0ine/go-json-rest/rest"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"os")
func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"Body": "Figher is running!"})
	}))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), api.MakeHandler()))
	
	response, err := http.Get("http://ip.jsontest.com/")
    if err == nil {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err == nil {
            fmt.Printf("%s\n", string(contents))
        }
    }
}