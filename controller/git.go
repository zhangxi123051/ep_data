package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

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

type ConfigStruct struct {
	Host              string   `json:"host"`
	Port              int      `json:"port"`
	AnalyticsFile     string   `json:"analytics_file"`
	StaticFileVersion int      `json:"static_file_version"`
	StaticDir         string   `json:"static_dir"`
	TemplatesDir      string   `json:"templates_dir"`
	SerTcpSocketHost  string   `json:"serTcpSocketHost"`
	SerTcpSocketPort  int      `json:"serTcpSocketPort"`
	Fruits            []string `json:"fruits"`
}

type Other struct {
	SerTcpSocketHost string   `json:"serTcpSocketHost"`
	SerTcpSocketPort int      `json:"serTcpSocketPort"`
	Fruits           []string `json:"fruits"`
}

func TestJson(c *gin.Context) {
	jsonStr := `{"host": "http://localhost:9090","port": 9090,"analytics_file": "",
					"static_file_version": 1,"static_dir": "E:/Project/goTest/src/",
                    "templates_dir": "E:/Project/goTest/src/templates/",
					"serTcpSocketHost": ":12340","serTcpSocketPort": 12340,
					"fruits": ["apple", "peach"]}`

	jsonStr2 := `{"accessToken":"507b5e08ee444dck887b66bd08672905",
				"clientToken":"64e3a5415bfe405d9485f1jf2ea5c68e",
				"selectedProfile":{"id":"selID","name":"Bluek404"},
				"availableProfiles":[{"id":"测试ava","name":"Bluek404"}]}`

	//json str 转map
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr2), &dat); err == nil {
		fmt.Println("==============json str 转map=======================")
		fmt.Println(dat)

		mapTmp := dat["selectedProfile"].(map[string]interface{})
		fmt.Println(mapTmp["id"])
		/*
		   var dat2 map[string]interface{}
		   if err := json.Unmarshal([]byte(jsonStr), &dat2); err == nil {
		       fmt.Println( dat2["firstName"])
		   }
		*/

		mapTmp2 := (dat["availableProfiles"].([]interface{}))[0].(map[string]interface{})
		//mapTmp3 := mapTmp2[0].(map[string]interface {})
		fmt.Println(mapTmp2["id"])
	}

	//json str 转struct
	var config ConfigStruct
	if err := json.Unmarshal([]byte(jsonStr), &config); err == nil {
		fmt.Println("================json str 转struct==")
		fmt.Println(config)
		fmt.Println(config.Host)
	}

	//json str 转struct(部份字段)
	var part Other
	if err := json.Unmarshal([]byte(jsonStr), &part); err == nil {
		fmt.Println("================json str 转struct==")
		fmt.Println(part)
		fmt.Println(part.SerTcpSocketPort)
	}

	//struct 到json str
	if b, err := json.Marshal(config); err == nil {
		fmt.Println("================struct 到json str==")
		fmt.Println(string(b))
	}

	//map 到json str
	fmt.Println("================map 到json str=====================")
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(dat)

	//array 到 json str
	arr := []string{"hello", "apple", "python", "golang", "base", "peach", "pear"}
	lang, err := json.Marshal(arr)
	if err == nil {
		fmt.Println("================array 到 json str==")
		fmt.Println(string(lang))
	}

	//json 到 []string
	var wo []string
	if err := json.Unmarshal(lang, &wo); err == nil {
		fmt.Println("================json 到 []string==")
		fmt.Println(wo)
	}
}

func StrToJson(c *gin.Context) {
	type Response struct {
		RequestID     string                   `json:"RequestId"`
		SendStatusSet []map[string]interface{} `json:"SendStatusSet"`
	}
	type r struct {
		Response Response `json:"Response"`
	}
	txt := `{
        "Response": {
            "SendStatusSet": [{
                    "SerialNo": "5000:1045710669157053657849499619",
                    "PhoneNumber": "+8618511122233",
                    "Fee": 1,
                    "SessionContext": "test",
                    "Code": "Ok",
                    "Message": "send success",
                    "IsoCode": "CN"
                },
                {
                    "SerialNo": "5000:104571066915705365784949619",
                    "PhoneNumber": "+8618511122266",
                    "Fee": 1,
                    "SessionContext": "test",
                    "Code": "Ok",
                    "Message": "send success",
                    "IsoCode": "CN"
                }
            ],
            "RequestId": "a0aabda6-cf91-4f3e-a81f-9198114a2279"
        }
    }`
	// fmt.Println(txt)
	p := &r{}
	err := json.Unmarshal([]byte(txt), p)
	fmt.Println(err)
	fmt.Println(*p)

	type simple struct {
		Response map[string]interface{}
	}
	zzz := new(simple)
	err = json.Unmarshal([]byte(txt), zzz)
	fmt.Println(err)
	fmt.Println("--------------")
	fmt.Println(*zzz)

	simpleJSON := `{"Name":"Xiao mi 6","ProductID":1,"Number":10000,"Price":2499,"IsOnSale":true}`
	type k struct {
		Name string
	}
	kk := &k{}
	err = json.Unmarshal([]byte(simpleJSON), kk)
	fmt.Println(err)
	fmt.Println(*kk)

}
