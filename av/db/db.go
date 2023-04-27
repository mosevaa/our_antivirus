package db

//здесь планируется доставать из бд все, что связано с сигнатурами и тд

import (
	"database/sql"
	"fmt"
	"log"
	"our_antivirus/av/signature"
	"strconv"

	_ "github.com/lib/pq"
)

var db *sql.DB

var signatures []signature.Signature

func Signatures() []signature.Signature {
	return signatures
}

func GetConnection() *sql.DB {
	return db
}

func Connect_db() {

	db, err := sql.Open("postgres", "postgres://root:root@localhost/av_signatures?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query the table
	rows, err := db.Query("SELECT * FROM signatures")

	loadedSignatures := []signature.Signature{}
	for rows.Next() {

		signature := signature.Signature{}
		var offsetBegin, offsetEnd string

		if err := rows.Scan(
			&signature.Id,
			&signature.Sign,
			&signature.Sha,
			&offsetBegin,
			&offsetEnd,
			&signature.Dtype,
		); err != nil {
			return
		}

		if offBegin, err := parseOffsetString(offsetBegin); err != nil {
			return
		} else {
			signature.OffsetBegin = offBegin
		}
		if offEnd, err := parseOffsetString(offsetEnd); err != nil {
			return
		} else {
			signature.OffsetEnd = offEnd
		}

		loadedSignatures = append(loadedSignatures, signature)

		fmt.Printf("Сигнатура №%d\n", signature.Id)
		fmt.Printf("Байт: %s\n", signature.Sign)
		fmt.Printf("SHA-256: %s\n", signature.Sha)
		fmt.Printf("offsetBegin: %d\n", signature.OffsetBegin)
		fmt.Printf("offsetEnd: %d\n", signature.OffsetEnd)
		fmt.Printf("Тип файла: %s\n", signature.Dtype)
		fmt.Printf("--------------------------------------------\n")
	}

}
func parseOffsetString(offset string) (int64, error) {
	n, err := strconv.ParseUint(offset, 16, 64)
	if err != nil {
		return 0, err
	}

	return int64(n), nil
}
