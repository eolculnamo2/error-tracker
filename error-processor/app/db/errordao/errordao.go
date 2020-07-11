package errordao

import (
	"github.com/eolculnamo2/error-tracker/error-processor/app/structs"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/eolculnamo2/error-tracker/error-processor/app/db"
)

func SaveNewWebError(newError structs.WebError) {
	errorTrackerDb := db.MongoClient.Database("error-tracker")
	webErrorsConnection := errorTrackerDb.Collection("web-errors")
	_, err := webErrorsConnection.InsertOne(db.MongoCtx, bson.D{
			{Key: "msg", Value: newError.Msg},
			{Key: "url", Value: newError.Url},
			{Key: "lineNo", Value: newError.LineNo},
			{Key: "columnNo", Value: newError.ColumnNo},
			{Key: "err", Value: newError.Err},
	})

	if err != nil {
		panic(err)
	}
}