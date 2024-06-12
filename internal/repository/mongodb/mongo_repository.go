package mongodb 

import (
	"context"

)
type MongoDBRepository[T any] struct {
	collection *mongo.Collection 
}
func NewMongoDBRepository[T any](dbURI,dbName,collectionName string)(*MongoDBRepository[T],error){
     clientOptions :=options.Client().ApplyURI(dbURI)
	client, err :=mongo.Connect(context.TODO(),clientOptions)
	if err !=nil {
		return nil, err 
	}
	collection :=client.Database(dbName).Collection(collectionName); 
	return &MongoDBRepository[T]{collection},nil; 


}
func (m *MongoDBRepository[T]) GetAll(ctx context.Context)([]T,error){
var results []T 
cursor, err :=m.collection.Find*(ctx, bson.M{})
if err !=nil {
	return nil , err
}
if err :=cursor.All(ctx, &results); err !=nil {
	return nil, err
}
return results, nil

}
func (m *MongoDBRepository[T]) Create(ctx context.Context ,entity T) error {
	_ ,err := m.collection.InsertOne(ctx,entity)
	return err; 

}
func (m *MongoDBRepository[T])GetByID(ctx context.Context , id string )(T, error ){
	var result T 
	err := m.collection.FindOne(ctx,bson.M{"id": id}).Decode(&result); 
	if err ==mongo.ErrNoDocuments {
		return result, repository.ErrNotFound
	}
	return result, err
}

func (m *MongoDBRepository[T])Update (ctx context.Context ,filter bson.M, update bson.M) error {

	_, err := m.collection.UpdateOne(ctx, filter, update)
	return err; 

}
func (m *MongoDBRepository[T])Delete(ctx context.Context, filter bson.M) error {
	_, err := m.collection.DeleteOne(ctx, filter)
	return err
}
