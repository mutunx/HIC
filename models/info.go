package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

type IC struct {
	Name      string `json:"name"`
	Sex       string `json:"sex"`
	Nation    string `json:"nation"`
	Birthday  string `json:"birthday"`
	Hometown  string `json:"hometown"`
	Career    string `json:"career"`
	IDCard    string `json:"IDCard"`
	Unit      string `json:"unit"`
	Living    string `json:"living"`
	Father    string `json:"father"`
	Mother    string `json:"mother"`
	Education string `json:"education"`
}

func GetInfos() (infos []IC, total int) {
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		_ = fmt.Errorf("get Infos err: %s", err)
	}
	_ = cur.All(context.TODO(), &infos)
	total = len(infos)
	return
}

func GetInfo(queryMap map[string]string) (infos []IC) {
	cur, err := collection.Find(context.TODO(), queryMap)
	if err != nil {
		_ = fmt.Errorf("get info err: %s", err)
	}
	_ = cur.All(context.TODO(), &infos)
	return
}

func EditInfo(IDCard string, info *IC) (count int64, err error) {
	result, err := collection.UpdateOne(context.TODO(), bson.M{"idcard": IDCard}, bson.M{"$set": info})
	count = result.ModifiedCount
	return
}

func AddInfo(info *IC) (infoId interface{}, err error) {
	result, err := collection.InsertOne(context.TODO(), info)
	infoId = result.InsertedID
	return
}
