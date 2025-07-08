package main

import (
    "net/http"
    "libraries/gin-1.10.1"
    "libraries/cors-master"
)

type album struct {
    ID string `json:"id"`
    User string `json:"user"`
    FirstName string `json:"firstName"`
}

var albums = []album{
    {ID: "1", User: "Tcraig01", FirstName: "Tucker"},
    {ID: "2", User: "Bcraig71", FirstName: "Butch"},
    {ID: "3", User: "Dcraig73", FirstName: "David"},

}

func main() {

    router := gin.Default()

    router.Use(cors.Default())

    router.GET("/albums", getAlbums)

    router.Run("localhost:8080")
    
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
     c.IndentedJSON(http.StatusOK, albums)
}   

