package handlers 

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/maulidihsan/go-gin-blog/database"		

)
func AddBlog(w http.ResponseWriter, r *http.Request) {
	var blog model.Blog 
	if err :=json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, "invalid body request", http.StatusBadRequest)
		return ; 
}
collection :=database.GetMongoClient().Database("college").Collection("blogs")

_,err :=collection.InsertOne(context.Background(),blog)
if err !=nil {
	log.Fatal(err)
}
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(map[string]string{"status": "created"})
}

func GetBlogByID(w http.ResponseWriter, r *http.Request) {
	blogID :=chi.URLParam(r,"ID")
	objID, err :=primitive.ObjectIDFromHex(blogID); 
	if err !=nil {
		log.Fatal(err)
		http.Error(w, "invalid blog id", http.StatusBadRequest)
	return 
	}
	collection :=database.GetMongoClient().Database("college").Collection("blogs")
	filter :=bson.D{{"ID",objID}}
	blog,err :=database.FindOneById(collection,filter)
	if err !=nil {
		http.Error(w,"Failed to fetch blog",http.StatusInternalServerError)
	return 
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(blog)		

	}
	func DeleteBlogByID(w http.ResponseWriter, r *http.Request) {
		blogID :=chi.URLParam(r,"ID")
		objID,err :=primitive.ObjectIDFromHex(blogID)
		if err !=nil {
			http.Error(w."invalid blog id",http.StatusBadRequest)
		return 
		}
		collection :=database.GetMongoClient().Database("college").Collection("blogs")
        filter :=bson.D{{"ID",objID}}; 
		err = database.DeleteOne(collection,filter)
		if err !=nil {
			http.Error(w,"Failed to delete blog",http.StatusInternalServerError)
		return 
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
		}
	func UpdateBlog(w http.ResponseWriter, r *http.Request) {
		blogID :=chi.URLParam(r,"ID")
		objID,err :=primitive.ObjectIDFromHex(blogID); 
		if err !=nil {
			http.Error(w,"invalid blog id",http.StatusBadRequest)
			return 
		}
		var updatedBlog model.Blog 
		if err :=json.NewDecoder(r.Body).Decode(&updatedBlog); err !=nil {
			http.Error(w,"invalid body request",http.StatusBadRequest)

		}
	}
	func GetAllBlogs(w http.ResponseWriter , r * http.Request){
		var blogs []model.Blog

		collection : =database.GetMongoClient().Database("college").Collection("blogs")
		blogs,err :=database.FindAll(collection); 
		if err !=nil {
			http.Error(w,"Failed to fetch blog",http.StatusInternalServerError)
			return 

		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(blogs); 

	}
	
	