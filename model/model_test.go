package model

import (
	"context"
	"fmt"
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
