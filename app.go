package main

import (
	"github.com/gin-gonic/gin"
)

var counter int = 20

/*func handle(w http.ResponseWriter, req *http.Request) {
	if _, err := fmt.Fprintf(w, "hello"); err != nil {
		fmt.Println("sending 'hej' error", err)
	}
}*/

func main() {
	/*fmt.Println("server: listening at :8090/")
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":8090", nil))*/
	r := gin.Default()
	r.POST("/user_tags", func(c *gin.Context) {
		c.JSON(204, gin.H{})
	})
	// POST /user_profiles/{cookie}?time_range=<time_range>?limit=<limit>
	r.POST("/user_profiles/:cookie", func(c *gin.Context) {

		//time_range := c.Query("time_range")
		//limit := c.Query("limit")
		c.JSON(200, c.Request.Body)

	})
	r.Run(":8090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
