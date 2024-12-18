// Program is always in package main
package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {

/*Use Context.BindJSON to bind the request body to newAlbum.
Append the album struct initialized from the JSON to the albums slice.
Add a 201 status code to the response, along with JSON representing the album you added.*/

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

/*Use Context.Param to retrieve the id path parameter from the URL. When you map this handler to a path, you’ll include a placeholder for the parameter in the path.

Loop over the album structs in the slice, looking for one whose ID field value matches the id parameter value. If it’s found, you serialize that album struct to JSON and return it as a response with a 200 OK HTTP code.

As mentioned above, a real-world service would likely use a database query to perform this lookup.

Return an HTTP 404 error with http.StatusNotFound if the album isn’t found.*/

    id := c.Param("id")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

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
	router := gin.Default()	// Initialize a Gin router using default
	router.GET("/albums", getAlbums)	// Using the getAlbums() method handler for /albums path
	router.GET("/albums/:id", getAlbumByID)	// In Gin, the : that prcedes an item specifies that item is a path parameter
	router.POST("/albums", postAlbums)
	
	router.Run("localhost:8080")	// Use the Run function to attach the router to an http.Server and start the server.
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {

/*Write a getAlbums function that takes a gin.Context parameter. Note that you could have given this function any name – neither Gin nor Go require a particular function name format.

gin.Context is the most important part of Gin. It carries request details, validates and serializes JSON, and more. (Despite the similar name, this is different from Go’s built-in context package.)

Call Context.IndentedJSON to serialize the struct into JSON and add it to the response.

The function’s first argument is the HTTP status code you want to send to the client. Here, you’re passing the StatusOK constant from the net/http package to indicate 200 OK.

Note that you can replace Context.IndentedJSON with a call to Context.JSON to send more compact JSON. In practice, the indented form is much easier to work with when debugging and the size difference is usually small.
*/

    c.IndentedJSON(http.StatusOK, albums)
}
