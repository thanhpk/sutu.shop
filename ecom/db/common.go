package db

import (
	bson "gopkg.in/mgo.v2/bson"
)

func AttachId(obj interface{}, id bson.ObjectId) bson.M{
	doc := bson.D{}
	marshalledObj, _ := bson.Marshal(obj)
	bson.Unmarshal(marshalledObj, &doc)
	docMap := doc.Map()
	docMap["_id"] = id;
	delete(docMap, "Id")
	return docMap
}
