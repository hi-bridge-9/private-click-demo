package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kyu-takahahsi/private-click-demo/cmd/lib/database/model"
)

var (
	user     = os.Getenv("MYSQL_USER")
	pass     = os.Getenv("MYSQL_PASSWORD")
	ip       = os.Getenv("MYSQL_IP")
	port     = os.Getenv("MYSQL_PORT")
	protocol = os.Getenv("MYSQL_PROTOCOL")
	name     = os.Getenv("MYSQL_DATABASE")
)

func connect() (*sql.DB, error) {
	conf := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8mb4",
		user,
		pass,
		protocol,
		ip,
		port,
		name)

	db, err := sql.Open("mysql", conf)
	if err != nil {
		return nil, fmt.Errorf("Invalid database config infomation: %w", err)
	}
	return db, nil
}

func InsertReport(r *model.Report, referer, host string) error {
	db, err := connect()
	if err != nil {
		return fmt.Errorf("Failed connect to DB: %w", err)
	}
	defer db.Close()

	insert := generateInsertReportQuery(r, referer, host, strings.TrimRight(host, ".test"))
	_, err = db.Exec(insert)
	if err != nil {
		return fmt.Errorf("Failed report data insert to DB: %w", err)
	}

	log.Println("Success report data insert to DB")
	return nil
}

func generateInsertReportQuery(r *model.Report, referer, host, member string) string {
	return fmt.Sprintf("INSERT INTO %s_report("+
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
		")VALUES('%s','%s',%d,'%s',%d,%d,'%s','%s','%s','%s')",
		member,
		r.EngagementType,
		r.SourceSite,
		r.SourceId,
		r.AttributedOnSite,
		r.TriggerData,
		r.Version,
		r.SecretToken,
		r.SecretTokenSignature,
		referer,
		host)
}

func InsertPublicToken(pt string, referer, host string) error {
	db, err := connect()
	if err != nil {
		return fmt.Errorf("Failed connect to DB: %w", err)
	}
	defer db.Close()

	insert := generateInsertPublicTokenQuery(pt, referer, host, strings.TrimRight(host, ".test"))
	_, err = db.Exec(insert)
	if err != nil {
		return fmt.Errorf("Failed public token insert to DB: %w", err)
	}

	log.Println("Success public token insert to DB")
	return nil
}

func generateInsertPublicTokenQuery(pt string, referer, host, member string) string {
	return fmt.Sprintf("INSERT INTO %s_public_token("+
		"token_public_key,"+
		"refere,"+
		"host"+
		")VALUES('%s','%s','%s')",
		member,
		pt,
		referer,
		host)
}

func InsertUnlinkableToken(s *model.Source, ut, referer, host string) error {
	db, err := connect()
	if err != nil {
		return fmt.Errorf("Failed connect to DB: %w", err)
	}
	defer db.Close()

	insert := generateInsertUnlinkableTokenQuery(s, ut, referer, host, strings.TrimRight(host, ".test"))
	_, err = db.Exec(insert)
	if err != nil {
		return fmt.Errorf("Failed unlinkable token insert to DB: %w", err)
	}

	log.Println("Success unlinkable token insert to DB")
	return nil
}

func generateInsertUnlinkableTokenQuery(s *model.Source, ut, referer, host, member string) string {
	return fmt.Sprintf("INSERT INTO %s_unlinkable_token("+
		"source_engagement_type,"+
		"source_nonce,"+
		"source_unlinkable_token,"+
		"version,"+
		"unlinkable_token,"+
		"refere,"+
		"host"+
		")VALUES('%s','%s','%s',%d,'%s','%s','%s')",
		member,
		s.EngagementType,
		s.Nonce,
		s.SourceToken,
		s.Version,
		ut,
		referer,
		host)
}
