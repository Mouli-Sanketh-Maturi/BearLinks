package service

import (
	"bearLinks/datastore"
	"context"
	"encoding/json"
	"github.com/dchest/uniuri"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"io"
	"log"
	"net/http"
	"time"
)

var ctx = context.Background()
var db *pgxpool.Pool

func Redirect(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	shortLink := vars["shortLink"]

	longLink, err := datastore.Rdb.Get(ctx, shortLink).Result()

	if err == redis.Nil {
		err = getDb().QueryRow(ctx, "SELECT longlink FROM bear_links WHERE shortlink = $1", shortLink).Scan(&longLink)

		if err != nil {
			http.Error(w, "Short Link does not exist", http.StatusNotFound)
			return
		}

		datastore.Rdb.Set(ctx, shortLink, longLink, time.Hour)
	}

	http.Redirect(w, r, longLink, http.StatusSeeOther)

}

func saveLinkToDB(link Link) {

	getDb().Exec(ctx, "INSERT INTO bear_links (shortlink, longlink, enabled) VALUES ($1, $2, $3)", link.ShortLink, link.LongLink, link.Enabled)

}

func GenerateShortLink(w http.ResponseWriter, r *http.Request)  {
	reqBody, _ := io.ReadAll(r.Body)

	var link Link

	json.Unmarshal(reqBody, &link)

	if link.ShortLink != "" {

		if checkIfShortLinkExists(link.ShortLink) {
			http.Error(w, "Short Link already exists", http.StatusBadRequest)
			return
		}

	} else {

		link.ShortLink = generateRandomShortUrl(5)

	}

	link.Enabled = true
	saveLinkToDB(link)
	datastore.Rdb.Set(ctx, link.ShortLink, link.LongLink, time.Hour)
	json.NewEncoder(w).Encode(link)

}

func generateRandomShortUrl(i int) string {
	for true {
		randomString := uniuri.NewLen(i)
		if !checkIfShortLinkExists(randomString) {
			return randomString
		}
	}
	return ""
}

func DeleteShortLink(w http.ResponseWriter, r *http.Request) {
	shortLink := r.URL.Query().Get("shortLink")
	exists := checkIfShortLinkExists(shortLink)
	if exists {
		getDb().Exec(ctx, "DELETE FROM bear_links WHERE shortlink = $1", shortLink)
		datastore.Rdb.Del(ctx, shortLink)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Short Link does not exist", http.StatusNotFound)
	}
}

func checkIfShortLinkExists(link string) bool {

	var shortLink string
	err := getDb().QueryRow(ctx, "SELECT shortlink FROM bear_links WHERE shortlink = $1", link).Scan(&shortLink)

	if err != nil {
		return false
	}

	return true
}

func getDb() *pgxpool.Pool {
	if(db == nil) {
		var err error
		db, err = pgxpool.New(ctx, "postgresql://moulisanketh:password@localhost/bearLinks")

		if err != nil {
			log.Fatal(err)
		}
	}
	return db
}