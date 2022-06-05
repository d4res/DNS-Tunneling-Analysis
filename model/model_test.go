package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestCon(t *testing.T) {
	Conn()
}

func TestFind(t *testing.T) {
	client, _ := Conn()
	col := client.Database("dns_pcap").Collection("request")
	cur, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		t.Error("error find", err)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var res Request
		err := cur.Decode(&res)
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("res: %v\n", res)
	}
}

func TestTime(t *testing.T) {
	client, _ := Conn()
	defer client.Disconnect(context.TODO())
	col := client.Database("dns_pcap").Collection("request")
	var res Request
	col.FindOne(context.TODO(), bson.D{}).Decode(&res)
	fmt.Println(res.Time.Local())
}

func TestInfo(t *testing.T) {
	client, err := Conn()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())
	infoCol := client.Database("dns_pcap").Collection("info")
	cur, err := infoCol.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.Background()) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		r, _ := bson.MarshalExtJSON(result, false, false)

		var json_data any
		json.Unmarshal(r, json_data)
		fmt.Printf("%T\n", json_data)
		break
	}
}

func TestCount(t *testing.T) {
	client, err := Conn()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	requestClient := client.Database("dns_pcap").Collection("request")
	ans, _ := requestClient.CountDocuments(context.Background(), bson.D{{Key: "type", Value: "MX"}})
	fmt.Println(ans)
}

func TestMetric(t *testing.T) {
	client, err := Conn()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	col := client.Database("dns_pcap").Collection("request")

	m := GetMetric(col)
	fmt.Println(m)
}
