
package main

import (
  "log"
  "fmt"
  "net/http"
  "os"
)

func determineListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
}

func hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Hello World")
}

func main() {
  addr, err := determineListenAddress()
  if err != nil {
    log.Fatal(err)
  }
  http.HandleFunc("/", hello)
  log.Printf("Listening on %s...\n", addr)
  if err := http.ListenAndServe(addr, nil); err != nil {
    panic(err)
  }
}

//package main

//import (
//	"log"
//	"net/http"
//	"os"

//	"github.com/gin-gonic/gin"
//)

//func main() {
//	port := os.Getenv("PORT")

//	if port == "" {
//		log.Fatal("$PORT must be set")
//	}

//	router := gin.New()
//	router.Use(gin.Logger())
//	router.LoadHTMLGlob("templates/*.tmpl.html")
//	router.Static("/static", "static")

//	router.GET("/", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "index.tmpl.html", nil)
//	})

//	router.Run(":" + port)
//}
