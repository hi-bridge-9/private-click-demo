package database

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
)

var (
	user     = os.Getenv("MYSQL_USER")
	pass     = os.Getenv("MYSQL_PASSWORD")
	ip       = os.Getenv("MYSQL_IP")
	port     = os.Getenv("MYSQL_PORT")
	protocol = os.Getenv("MYSQL_PROTOCOL")
	name     = os.Getenv("MYSQL_DATABASE")
)

func Connect() (*sql.DB, error) {
	conf := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8mb4",
		user,
		pass,
		protocol,
		ip,
		port,
		name)

	db, err := sql.Open("mysql", conf)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GenerateInsertReportQuery(r *http.Request) string {
	return fmt.Sprintf("NSERT INTO publisher_report("+
		"source_engagement_type,"+
		"source_site,"+
		"source_id,"+
		"attributed_on_site,"+
		"trigger_data,"+
		"version,"+
		"secret_token,"+
		"secret_token_signature,"+
		"refere,"+
		"host"+
		")VALUES('%s','%s',%s,'%s',%s,%s,'%s','%s','%s','%s')",
		r.FormValue("source_engagement_type"),
		r.FormValue("source_site"),
		r.FormValue("source_id"),
		r.FormValue("attributed_on_site"),
		r.FormValue("trigger_data"),
		r.FormValue("version"),
		r.FormValue("secret_token"),
		r.FormValue("secret_token_signature"),
		r.Referer(),
		r.Host)
}

func GenerateInsertPublicTokenQuery(pt string, r *http.Request) string {
	return fmt.Sprintf("NSERT INTO publisher_public_token("+
		"token_public_key,"+
		"refere,"+
		"host"+
		")VALUES('%s','%s','%s')",
		pt,
		r.Referer(),
		r.Host)
}

func GenerateInsertUnlinkableTokenQuery(ut string, r *http.Request) string {
	return fmt.Sprintf("NSERT INTO publisher_unlinkable_token("+
		"source_engagement_type,"+
		"source_nonce,"+
		"source_unlinkable_token,"+
		"version,"+
		"unlinkable_token,"+
		"refere,"+
		"host"+
		")VALUES('%s','%s','%s',%s,'%s','%s','%s')",
		r.FormValue("source_engagement_type"),
		r.FormValue("source_nonce"),
		r.FormValue("source_unlinkable_token"),
		r.FormValue("version"),
		ut,
		r.Referer(),
		r.Host)
}
