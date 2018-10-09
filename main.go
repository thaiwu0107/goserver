package main

import (
	"bytes"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	Init "git.com.ggttoo44/src/init"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	// r.GET("/user/:name", func(c *gin.Context) {
	// 	user := c.Params.ByName("name")
	// 	value, ok := db[user]
	// 	if ok {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	// 	} else {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	// 	}
	// })
	r.GET("/user/:name", func(c *gin.Context) {
		var in = [7]int{144, 131, 121, 122, 91, 24, 132}
		var cardMap = [5][15]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		}
		var loop5 = [5]int{0, 1, 2, 3, 4}
		var loop7 = [7]int{0, 1, 2, 3, 4, 5, 6}
		var loop13 = [13]int{14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
		var loopIndex int
		var pointNumber int
		var suitNuber int
		var suitInt int
		suitMap := make(map[int][]int, 7)
		loopIndex = 0
		for _, loopIndex = range loop7 {
			suitInt = in[loopIndex]
			suitNuber = suitInt % 10
			pointNumber = (suitInt - suitNuber) / 10
			cardMap[suitNuber][pointNumber]++
			cardMap[suitNuber][0]++
			cardMap[0][pointNumber]++

			if suitMap[pointNumber] != nil {
				suitMap[pointNumber] = append(suitMap[pointNumber], suitNuber)
			} else {
				suitMap[pointNumber] = []int{suitNuber}
			}
		}
		isFlush := false
		var selectSuit int
		loopIndex = 0
		for _, loopIndex = range loop5 {
			if cardMap[loopIndex][0] >= 5 {
				isFlush = true
				selectSuit = loopIndex
				break
			}
		}
		var buffer bytes.Buffer
		loopIndex = 14
		for _, loopIndex = range loop13 {
			buffer.WriteString(strconv.Itoa(cardMap[selectSuit][loopIndex]))
		}
		keyOfRank := buffer.String()
		var rankInfo *Init.Rank7
		selectCards := make([]string, 5)
		if isFlush {
			rankInfo = Init.RankTable7CF[keyOfRank]
			loopIndex = 0
			for _, loopIndex = range loop5 {
				var point1 int
				point1 = rankInfo.CardPoint[loopIndex]
				selectCards[loopIndex] = strconv.Itoa(point1*10 + selectSuit)
			}
		} else {
			rankInfo = Init.RankTable7CNF[keyOfRank]
			loopIndex = 0
			for _, loopIndex = range loop5 {
				var point2 int
				var suit int
				point2 = rankInfo.CardPoint[loopIndex]
				suit, suitMap[point2] = suitMap[point2][len(suitMap[point2])-1], suitMap[point2][:len(suitMap[point2])-1]
				selectCards[loopIndex] = strconv.Itoa(point2*10 + suit)
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"CardPoint": selectCards,
			"Type7":     rankInfo.Type7,
			"Rank":      int32(rankInfo.Rank),
			"Type5Ch":   rankInfo.Type5Ch,
			"Type5En":   rankInfo.Type5En,
		})
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	Init.Start()
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
