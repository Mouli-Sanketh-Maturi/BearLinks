package main

import (
    "context"
    "encoding/json"
    "github.com/dchest/uniuri"
    "github.com/gorilla/mux"
    "github.com/jackc/pgx/v5/pgxpool"
    "io"
    "log"
    "net/http"
)


func main() {
    httpRouter := mux.NewRouter().StrictSlash(true)
    httpRouter.HandleFunc("/generateShortLink", generateShortLink).Methods("POST")
    httpRouter.HandleFunc("/{shortLink}", redirect).Methods("GET")
    http.ListenAndServe(":80", httpRouter)
}

func redirect(w http.ResponseWriter, r *http.Request)  {
    vars := mux.Vars(r)
    shortLink := vars["shortLink"]

    pool, err := pgxpool.New(context.Background(), "postgresql://moulisanketh:password@localhost/bearLinks")

    if err != nil {
        log.Fatal(err)
    }

    var longLink string
    err = pool.QueryRow(context.Background(), "SELECT longlink FROM bear_links WHERE shortlink = $1", shortLink).Scan(&longLink)

    if err != nil {
        http.Error(w, "Short Link does not exist", http.StatusNotFound)
        return
    }

    http.Redirect(w, r, longLink, http.StatusSeeOther)

}

func saveLinkToDB(link Link) {

    pool, err := pgxpool.New(context.Background(), "postgresql://moulisanketh:password@localhost/bearLinks")

    if err != nil {
        log.Fatal(err)
    }

    pool.Exec(context.Background(), "INSERT INTO bear_links (shortlink, longlink, enabled) VALUES ($1, $2, $3)", link.ShortLink, link.LongLink, link.Enabled)

}

func generateShortLink(w http.ResponseWriter, r *http.Request)  {
    reqBody, _ := io.ReadAll(r.Body)

    var link Link

    json.Unmarshal(reqBody, &link)

    if link.ShortLink != "" {

        if !checkIfShortLinkExists(link.ShortLink) {
            link.Enabled = true
            saveLinkToDB(link)
            json.NewEncoder(w).Encode(link)
            return
        } else {
            http.Error(w, "Short Link already exists", http.StatusBadRequest)
            return
        }
    }

    link.ShortLink = generateRandomShortUrl(5)
    link.Enabled = true
    saveLinkToDB(link)
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

func checkIfShortLinkExists(randomString string) bool {
    pool, err := pgxpool.New(context.Background(), "postgresql://moulisanketh:password@localhost/bearLinks")

    if err != nil {
        log.Fatal(err)
    }

    var shortLink string
    err = pool.QueryRow(context.Background(), "SELECT shortlink FROM bear_links WHERE shortlink = $1", randomString).Scan(&shortLink)

    if err != nil {
        return false
    }

    return true
}