package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
	"net"
	"net/http"
)

func main() {
	db, err := geoip2.Open("geoip-data/GeoLite2-City.mmdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	ip := net.ParseIP("4.10.15.22")
	ipRecord, err := db.City(ip)
	if err != nil {
		panic(err)
	}
	fmt.Println(ipRecord.Country.Names["en"])
	fmt.Println(ipRecord.City.Names["en"])
	fmt.Println(ipRecord.Postal.Code)
	fmt.Println(ipRecord.Location.Latitude, ipRecord.Location.Longitude)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service": "GeoIP",
		})
	})
	r.GET("/ip", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ip": "",
		})
	})
	err = r.Run(":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
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
