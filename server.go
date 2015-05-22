package main

import (
    "io/ioutil"
    "fmt"
	_ "net/http"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
    "github.com/sriniprash/pb/store"
    "github.com/sriniprash/pb/utils"

)

func main() {
	e := echo.New()
	e.Use(mw.Logger)

    // Initialize the store
    store := store.FileStore{RootDir: "pastes"}
    err := store.Init()
    if err != nil {
        fmt.Println(err)
        return
    }
    // Create a new Post.
    e.Post("/", func(c *echo.Context) {
        data, err := ioutil.ReadAll(c.Request.Body)
        if err != nil {
            c.JSON(500, err)
        } else {
            pasteID := utils.RandomString()
            fmt.Println("pasteID: ", pasteID)
            err := store.Create(pasteID, data)
            if err != nil {
                c.JSON(500, err)
            } else {
                c.JSON(200, pasteID)
            }
        }
    })

    // Fetch an existing post
	e.Get("/:pasteID", func(c *echo.Context) {
        pasteID := c.P(0)
        data, err := store.Get(pasteID)
        if err != nil {
            c.JSON(500, err)
        } else {
            c.String(200, string(data))
        }
    })

    // Modify an existing post
    e.Put("/:pasteID", func(c *echo.Context) {
        data, err := ioutil.ReadAll(c.Request.Body)
        if err != nil {
            c.JSON(500, err)
        } else {
            pasteID := c.P(0)
            err := store.Update(pasteID, data)
            if err != nil {
                c.JSON(500, err)
            } else {
                c.JSON(200, "Existing paste modified")
            }
            
        }
        
    })

    // Start the server
    e.Run(":8000")
}
