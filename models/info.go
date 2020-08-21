package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

type IC struct {
	_id       string
	Name      string
	Sex       string
	Nation    string
	Birthday  string
	Hometown  string
	Career    string
	IDCard    string
	Unit      string
	Living    string
	Father    string
	Mother    string
	Brother   string
	Education string
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
