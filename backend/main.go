package main

import (
	"DNSpcap/model"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	client, err := model.Conn()
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.TODO())
	requestCol := client.Database("dns_pcap").Collection("request")
	infoCol := client.Database("dns_pcap").Collection("info")
	//responseCol := client.Database("dns_pcap").Collection("response")

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "DELETE"},
		AllowHeaders: []string{"*"},
		MaxAge:       12 * time.Hour,
	}))

	router.GET("/data", func(ctx *gin.Context) {

		cur, err := requestCol.Find(context.TODO(), bson.D{})
		if err != nil {
			panic(err)
		}
		defer cur.Close(context.Background())

		resList := make([]model.LabeledDomain, 0)

		for cur.Next(context.Background()) {
			var rawData model.Request
			var res model.LabeledDomain
			err := cur.Decode(&rawData)
			if err != nil {
				panic(err)
			}
			rawData.Time = rawData.Time.Local()
			res.Payload = rawData.Payload
			res.Tag = rawData.Tag
			res.Time = rawData.Time.Format("2006-01-02 15:04:05")

			resList = append(resList, res)
		}

		ctx.JSON(http.StatusAccepted, resList)
	})

	router.GET("/info", func(ctx *gin.Context) {
		cur, err := infoCol.Find(context.TODO(), bson.D{})
		if err != nil {
			panic(err)
		}
		defer cur.Close(context.Background())

		resList := make([]string, 0)

		for cur.Next(context.Background()) {
			var rawData bson.D
			err := cur.Decode(&rawData)

			if err != nil {
				panic(err)
			}
			s, _ := bson.MarshalExtJSON(rawData, false, false)
			resList = append(resList, string(s))
		}

		res := strings.Join(resList, ",")
		res = `[` + res + `]`
		ctx.Data(http.StatusAccepted, "application/json", []byte(res))
	})

	router.GET("/metric", func(ctx *gin.Context) {
		m := model.GetMetric(requestCol)
		ctx.JSON(http.StatusOK, m)
	})
	router.Run("localhost:8888")
}
