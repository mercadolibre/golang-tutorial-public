package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Fullstack conf location
const CENTER_LAT = -34.5596416
const CENTER_LNG = -58.5076697
const CENTER_TITLE = "Estas acá"

type marker struct {
	Title string
	Lat   float64
	Lng   float64
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func search(c *gin.Context) {
	query, err := validateQuery(c)
	if err != nil {
		return
	}

	result,err:=searchItems(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError,  apiErr{err.Error(),http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, result)
}

func render(c *gin.Context) {
	query, err := validateQuery(c)
	if err != nil {
		return
	}

	result, err := searchItems(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, apiErr{err.Error(), http.StatusInternalServerError})
		return
	}

	markers := make([]marker, 0, len(result)+1)

	//adding 'you are here' position
	markers = append(markers, marker{CENTER_TITLE, CENTER_LAT, CENTER_LNG})

	//adding all the items positions and info
	for _, itemInfo := range result {
		//skip items without geolocation info
		if itemInfo.Location.Latitude != 0 && itemInfo.Location.Longitude != 0 {
			markers = append(markers, marker{fmt.Sprintf("%s - $%v", itemInfo.Title, itemInfo.Price), itemInfo.Location.Latitude, itemInfo.Location.Longitude})
		}
	}

	c.HTML(http.StatusOK, "map.tmpl", gin.H{"title": fmt.Sprintf("Resulados para: %s", query), "centerLat": CENTER_LAT, "centerLng": CENTER_LNG, "markers": markers})
}

func validateQuery(c *gin.Context) (string, error) {
	query := c.Query("query")

	if len(query) == 0 {
		err := apiErr{"'query' param should be send", http.StatusBadRequest}

		c.JSON(http.StatusBadRequest, err)

		return "", err
	}

	return query, nil
}

//Custom api error
type apiErr struct {
	ErrorMessage string `json:"message"`
	ErrorStatus  int    `json:"status"`
}

func (e apiErr) Error() string {
	return fmt.Sprintf("Message: %s; Status: %d", e.ErrorMessage, e.ErrorStatus)
}
