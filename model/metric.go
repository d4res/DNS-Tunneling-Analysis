package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Metric struct {
	All   int64
	MX    int64
	TXT   int64
	CNAME int64
	Eval  int64
}

func GetOneMetric(col *mongo.Collection, key string, value any) (ret int64) {
	if key == "" {
		ret, _ = col.CountDocuments(context.Background(), bson.D{{}})
		return
	}
	ret, _ = col.CountDocuments(context.Background(), bson.D{{Key: key, Value: value}})

	return
}

func GetMetric(col *mongo.Collection) (metric Metric) {
	metric.All = GetOneMetric(col, "", "")
	metric.CNAME = GetOneMetric(col, "type", "CNAME")
	metric.MX = GetOneMetric(col, "type", "MX")
	metric.TXT = GetOneMetric(col, "type", "TXT")
	metric.Eval = GetOneMetric(col, "tag", true)
	return
}
