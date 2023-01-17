package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
	"net"
	"net/http"
)

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service": "GeoIP",
		})
	})
	r.GET("/ip", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ip": getClientIP(c),
		})
	})
	r.GET("/geoip/:area/:ip", func(c *gin.Context) {
		area := c.Param("area")
		if area == "country" {
			paramIP := c.Param("ip")
			if paramIP != "" {
				c.JSON(http.StatusOK, getGeoIPCountry(paramIP))
			}
		} else if area == "city" {
			paramIP := c.Param("ip")
			if paramIP != "" {
				c.JSON(http.StatusOK, getGeoIPCity(paramIP))
			}
		} else {
			c.JSON(http.StatusNotFound, JSONResponseError{"Invalid query"})
		}
	})
	err := r.Run(":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
}

type JSONResponse struct {
	IP         string  `json:"ip"`
	Country    string  `json:"country"`
	City       string  `json:"city"`
	PostalCode string  `json:"postal_code"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

type JSONResponseCountry struct {
	IP      string `json:"ip"`
	Country string `json:"country"`
}

type JSONResponseError struct {
	Error string `json:"error"`
}

func getClientIP(c *gin.Context) string {
	ip := c.Request.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = c.Request.Header.Get("X-Real-Ip")
	}
	if ip == "" {
		ip = c.Request.RemoteAddr
	}
	return ip
}

func getGeoIPCity(ip string) JSONResponse {
	db, err := geoip2.Open("geoip-data/GeoLite2-City.mmdb")
	if err != nil {
		panic(err)
	}
	defer func(db *geoip2.Reader) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
	parsedIP := net.ParseIP(ip)
	ipRecord, err := db.City(parsedIP)
	if err != nil {
		panic(err)
	}

	return JSONResponse{
		IP:         ip,
		Country:    ipRecord.Country.Names["en"],
		City:       ipRecord.City.Names["en"],
		PostalCode: ipRecord.Postal.Code,
		Latitude:   ipRecord.Location.Latitude,
		Longitude:  ipRecord.Location.Longitude,
	}
}

func getGeoIPCountry(ip string) JSONResponseCountry {
	db, err := geoip2.Open("geoip-data/GeoLite2-Country.mmdb")
	if err != nil {
		panic(err)
	}
	defer func(db *geoip2.Reader) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
	parsedIP := net.ParseIP(ip)
	ipRecord, err := db.Country(parsedIP)
	if err != nil {
		panic(err)
	}

	return JSONResponseCountry{
		IP:      ip,
		Country: ipRecord.Country.Names["en"],
	}
}
