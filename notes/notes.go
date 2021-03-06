package notes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"nextnotes/utils"

	"google.golang.org/api/iterator"
)

type Note struct {
	Name   string `json:"Name"`
	Colour string `json:"Colour"`
	Text   string `json:"Text"`
	UserId string `json:"UserId"`
}

func Notes(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client := utils.GetClientFirestore(ctx)
	utils.EnableCors(&w)

	switch r.Method {
	case http.MethodGet:
		id := r.URL.Query().Get("id")
		userId := r.URL.Query().Get("userId")

		var documents []map[string]interface{}

		if id != "" && userId != "" {
			iter := client.Collection("notes").Where(
				"UserId", "==", userId,
			).Where(
				"__name__", "==", id,
			).Documents(ctx)
			for {
				doc, err := iter.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					log.Fatalf("Fatality: %v", err)
				}
				var m map[string]interface{}
				doc.DataTo(&m)
				m["ID"] = doc.Ref.ID
				documents = append(documents, m)
			}
		} else if userId != "" {
			iter := client.Collection("notes").Where(
				"UserId", "==", userId,
			).Documents(ctx)
			for {
				doc, err := iter.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					log.Fatalf("Fatality: %v", err)
				}
				var m map[string]interface{}
				doc.DataTo(&m)
				m["ID"] = doc.Ref.ID
				documents = append(documents, m)
			}
		}
		json.NewEncoder(w).Encode(documents)
	case http.MethodPost:
		note := Note{}
		err := json.NewDecoder(r.Body).Decode(&note)
		if err != nil {
			log.Fatalf("Fatality: %v", err)
			return
		}
		ref, result, err := client.Collection("notes").Add(ctx, note)
		_ = ref
		_ = result
		if err != nil {
			log.Fatalf("Fatality: %v", err)
			return
		}
		json.NewEncoder(w).Encode(result)
	case http.MethodPut:
		note := Note{}
		err := json.NewDecoder(r.Body).Decode(&note)
		if err != nil {
			log.Fatalf("Fatality: %v", err)
			return
		}

		id := r.URL.Query().Get("id")
		if id != "" {
			ref := client.Collection("notes").Doc(id)
			result, err := ref.Set(ctx, note)
			json.NewEncoder(w).Encode(result)
			if err != nil {
				log.Fatalf("Fatality: %v", err)
				return
			}
			return
		} else {
			http.Error(w, "404 - Not found", http.StatusNotFound)
		}
	case http.MethodDelete:
		id := r.URL.Query().Get("id")

		if id != "" {
			result, err := client.Collection("notes").Doc(id).Delete(ctx)
			if err != nil {
				log.Fatalf("Fatality: %v", err)
				return
			}
			json.NewEncoder(w).Encode(result)
			return
		} else {
			http.Error(w, "404 - Not found", http.StatusNotFound)
		}
	case http.MethodOptions:
		return
	default:
		// http.Error(w, "403 - Forbidden", http.StatusForbidden)
		http.Error(w, "405 - Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
