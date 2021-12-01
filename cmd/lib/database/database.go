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

	char := strings.TrimRight(host, ".test")

	init := fmt.Sprintf("CREATE TABLE IF NOT EXISTS pcm.%s_report ("+
		"id INT AUTO_INCREMENT,"+
		"source_engagement_type VARCHAR(10),"+
		"source_site VARCHAR(30),"+
		"source_id INT,"+
		"attributed_on_site VARCHAR(30),"+
		"trigger_data INT,"+
		"version INT,"+
		"secret_token VARCHAR(100),"+
		"secret_token_signature VARCHAR(100),"+
		"refere VARCHAR(30),"+
		"host VARCHAR(30)"+
		"date DATETIME DEFAULT CURRENT_TIMESTAMP,"+
		"PRIMARY KEY (id)"+
		");", char)

	_, err = db.Exec(init)
	if err != nil {
		return fmt.Errorf("Failed init report table: %w", err)
	}

	insert := generateInsertReportQuery(r, referer, host, char)
	_, err = db.Exec(insert)
	if err != nil {
		return fmt.Errorf("Failed report data insert to DB: %w", err)
	}

	log.Println("Success report data insert to DB")
	return nil
}

func generateInsertReportQuery(r *model.Report, referer, host, char string) string {
	return fmt.Sprintf("INSERT INTO pcm.%s_report("+
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
		char,
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

	char := strings.TrimRight(host, ".test")

	init := fmt.Sprintf("CREATE TABLE IF NOT EXISTS pcm.%s_public_token ("+
		"id INT AUTO_INCREMENT,"+
		"token_public_key VARCHAR(100),"+
		"host VARCHAR(30)"+
		"date DATETIME DEFAULT CURRENT_TIMESTAMP,"+
		"PRIMARY KEY (id)"+
		");", char)

	_, err = db.Exec(init)
	if err != nil {
		return fmt.Errorf("Failed init public token table: %w", err)
	}

	insert := generateInsertPublicTokenQuery(pt, referer, host, char)
	_, err = db.Exec(insert)
	if err != nil {
		return fmt.Errorf("Failed public token insert to DB: %w", err)
	}

	log.Println("Success public token insert to DB")
	return nil
}

func generateInsertPublicTokenQuery(pt string, referer, host, char string) string {
	return fmt.Sprintf("INSERT INTO pcm.%s_public_token("+
		"token_public_key,"+
		"refere,"+
		"host"+
		")VALUES('%s','%s','%s')",
		char,
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

	char := strings.TrimRight(host, ".test")

	init := fmt.Sprintf("CREATE TABLE IF NOT EXISTS pcm.%s_unlinkable_token ("+
		"id INT AUTO_INCREMENT,"+
		"source_engagement_type VARCHAR(10),"+
		"source_nonce VARCHAR(30),"+
		"source_unlinkable_token VARCHAR(100),"+
		"version INT,"+
		"unlinkable_token VARCHAR(100),"+
		"refere VARCHAR(30),"+
		"host VARCHAR(30)"+
		"date DATETIME DEFAULT CURRENT_TIMESTAMP,"+
		"PRIMARY KEY (id)"+
		");", char)

	_, err = db.Exec(init)
	if err != nil {
		return fmt.Errorf("Failed init report table: %w", err)
	}

	insert := generateInsertUnlinkableTokenQuery(s, ut, referer, host, char)
	_, err = db.Exec(insert)
	if err != nil {
		return fmt.Errorf("Failed unlinkable token insert to DB: %w", err)
	}

	log.Println("Success unlinkable token insert to DB")
	return nil
}

func generateInsertUnlinkableTokenQuery(s *model.Source, ut, referer, host, char string) string {
	return fmt.Sprintf("INSERT INTO pcm.%s_unlinkable_token("+
		"source_engagement_type,"+
		"source_nonce,"+
		"source_unlinkable_token,"+
		"version,"+
		"unlinkable_token,"+
		"refere,"+
		"host"+
		")VALUES('%s','%s','%s',%d,'%s','%s','%s')",
		char,
		s.EngagementType,
		s.Nonce,
		s.SourceToken,
		s.Version,
		ut,
		referer,
		host)
}
