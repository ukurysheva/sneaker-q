package repository

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	sneakerq "github.com/ukurysheva/sneaker-q"
)

type ModelsPostgres struct {
	db *sqlx.DB
}

func NewModelsPostgres(db *sqlx.DB) *ModelsPostgres {
	return &ModelsPostgres{db: db}
}

func (r *ModelsPostgres) AddModelsList(models []*sneakerq.Model) error {
	tx, err := r.db.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}
	// query := fmt.Sprintf("INSERT INTO %s VALUES", modelsTable)
	var values []interface{}
	for _, model := range models {
		values = append(values, model.Title, model.Price, model.PageUrl, model.ShopId)
	}
	fmt.Println(values)
	query := "INSERT INTO %s(title, price, page_url, shop_id) VALUES"
	query = setupBindVars(query, 4, len(values))
	query += " ON CONFLICT (page_url) DO NOTHING"
	fmt.Println(query)
	stmt, err := r.db.Prepare(query)
	if err != nil {
		tx.Rollback()
		log.Fatalf("error while preparing sql: %s", err.Error())
		return err
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.
	if _, err := stmt.Exec(values...); err != nil {
		tx.Rollback()
		log.Fatalf("error while exec sql: %s", err.Error())
		return err
	}

	return tx.Commit()
}

func setupBindVars(stmt string, lenargs int, len int) string {
	fmt.Println("start binding")
	bindVars := ""
	i := 1
	for i <= len {
		bindVars += "("
		for j := 0; j < lenargs; j++ {
			bindVars += "$" + strconv.Itoa(i) + ","
			i++
		}
		bindVars = strings.TrimSuffix(bindVars, ",")
		bindVars += "),"
	}
	stmt += bindVars
	stmt = fmt.Sprintf(stmt+bindVars, modelsTable)
	return strings.TrimSuffix(stmt, ",")
}
