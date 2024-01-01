package std

import (
	"context"

	"github.com/mehdi-shokohi/mongoHelper"
)

func GetDefaultDb[T any](ctx context.Context,colName string, model T)(mongoHelper.MongoContainer[T]){
	dbId:="dafault"
	uri:="mongodb://inflowdb:27017/inflowenger?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false"
	inflowDatabase:="default"
	db:= mongoHelper.NewMongoById[T](ctx,dbId,uri,inflowDatabase,colName,model)
	return db
}