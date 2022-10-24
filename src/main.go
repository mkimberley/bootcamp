package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var startTime = time.Now()
var ignoreState = false

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/call_layers", callLayers)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.GET("/health_check", getHealthCheck)
	router.GET("/set_ignore", setIgnore)

	router.Run("localhost:8080")
}

func callLayers(c *gin.Context) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	c.IndentedJSON(http.StatusOK, gin.H{"IP address": localAddr.IP, "message": "OK"})

}
func setIgnore(c *gin.Context) {
	ignoreState = true
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Health check state set to ignore"})
}

func uptime() time.Duration {
	return (time.Duration(time.Since(startTime).Seconds()))
}

func getHealthCheck(c *gin.Context) {

	if !ignoreState {
		c.IndentedJSON(http.StatusOK, gin.H{"uptime": uptime(), "message": "OK"})
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
	}
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
