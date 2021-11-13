package controller

import "github.com/gin-gonic/gin"

func ParseJson(c *gin.Context) {
	str := `{
   "name":"computers",
   "description":"List of computer products",
   "vendor":{
      "name":"Star Trek",
      "email":"info@example.com",
      "website":"www.example.com",
      "items":[
         {"id":1, "name":"MacBook Pro 13 inch retina","price":1350},
         {"id":2, "name":"MacBook Pro 15 inch retina", "price":1700},
         {"id":3, "name":"Sony VAIO", "price":1200},
         {"id":4, "name":"Fujitsu", "price":850},
         {"id":5, "name":"HP core i5", "price":850, "key": 2300},
         {"id":6, "name":"HP core i7", "price":950},
         {"id":null, "name":"HP core i3 SSD", "price":850}
      ],
      "prices":[
         2400,
         2100,
         1200,
         400.87,
         89.90,
         150.10
     ]
   }
	`
	println("str=", str)

	return
}
