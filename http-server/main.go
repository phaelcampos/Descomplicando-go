package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type GithubWebhookRequestBody struct {
	Action     string
	Number     int64
	Repository string
	Sender     string
}

// Handler
func echoHandler(c echo.Context) error {
	req := c.Request()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	obj := new(GithubWebhookRequestBody)
	err = json.Unmarshal(body, obj)
	if err != nil {
		panic(err)
	}
	c.Logger().Info("Action", obj.Action)
	c.Logger().Info("Number", obj.Number)
	c.Logger().Info("Sender", obj.Sender)
	c.Logger().Info("Repository", obj.Repository)
	return c.String(http.StatusOK, "Ok")
}

func main() {
	server := echo.New()
	server.POST("/", echoHandler)
	server.Logger.SetLevel(log.DEBUG)
	server.Logger.SetOutput(os.Stdout)
	server.Logger.Fatal(server.Start(":8080"))
}
