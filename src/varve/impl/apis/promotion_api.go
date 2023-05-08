package apis

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"golangTask/src/varve/impl/config"
	"golangTask/src/varve/impl/entities"
	"golangTask/src/varve/impl/service"
	"net/http"
)

func GetPromotionByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/promotions/"):]
	val, err := service.GetById(id, config.ConnToRedis())
	if err != nil {
		if err == redis.Nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var promotion entities.Promotion
	err = json.Unmarshal([]byte(val), &promotion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(promotion)

}
