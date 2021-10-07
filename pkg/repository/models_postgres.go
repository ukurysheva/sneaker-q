package repository

import (
	"fmt"
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

func (r *ModelsPostgres) GetModelById(modelId int) (sneakerq.Model, error) {
	model := sneakerq.Model{}
	list, err := r.GetModelsByParams(sneakerq.SearchParams{Id: modelId})
	if err != nil || len(list) < 1 {
		return model, err
	}

	model = list[0]
	return model, nil
}

// GetShopModels finds models in db by shop name
func (r *ModelsPostgres) GetShopModels(shopname string) ([]sneakerq.Model, error) {
	var shopId int
	query := fmt.Sprintf("SELECT id FROM  %s WHERE s.class_name = $1", shopsTable)
	err := r.db.Select(&shopId, query, shopname)
	if shopId < 1 || err != nil {
		return []sneakerq.Model{}, err
	}

	return r.GetModelsByParams(sneakerq.SearchParams{ShopId: shopId})
}

// GetShopModels finds models in db by shop name
func (r *ModelsPostgres) GetModelsByParams(searchParams sneakerq.SearchParams) ([]sneakerq.Model, error) {
	models := []sneakerq.Model{}
	query := fmt.Sprintf("SELECT m.id, m.title, m.size, m.price, m.avail, m.page_url FROM  %s AS m  LEFT JOIN %s AS s ON m.shop_id = s.id WHERE 1=1", modelsTable, shopsTable)
	index := 1
	args := make([]interface{}, 0)

	if len(searchParams.Size) > 0 {
		for _, size := range searchParams.Size {
			args = append(args, size)
		}
		bindsQuery, lastindex := getMultiBinds(len(searchParams.Size), index)
		query += " AND m.size IN (" + bindsQuery + ")"
		index = lastindex
	}

	if searchParams.PriceFrom > 0 {
		query += " AND m.price >= $" + strconv.Itoa(index)
		args = append(args, searchParams.PriceFrom)
		index++
	}
	if searchParams.PriceTo > 0 {
		query += " AND m.price <= $" + strconv.Itoa(index)
		args = append(args, searchParams.PriceTo)
		index++
	}
	fmt.Printf(query)
	err := r.db.Select(&models, query, args...)
	if err != nil {

		fmt.Printf(err.Error())
		return models, err
	}
	return models, nil
}

func (r *ModelsPostgres) AddUpdateModelsList(models []*sneakerq.Model) error {
	for _, model := range models {
		updated, err := r.UpdateModel(model)
		if err != nil {
			fmt.Printf("Error while updating model %s: %s", model.PageUrl, err.Error())
			continue
		}

		if updated == 0 {
			err = r.AddModel(model)
			if err != nil {
				fmt.Printf("Error while adding new model %s: %s", model.PageUrl, err.Error())
			}
		}
	}

	return nil
}

func (r *ModelsPostgres) UpdateModel(model *sneakerq.Model) (int64, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if model.Price != 0 {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, model.Price)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE page_url = $%d AND title = $%d  AND shop_id = $%d", modelsTable, setQuery, argId, argId+1, argId+2)
	args = append(args, model.PageUrl, model.Title, model.ShopId)
	res, err := r.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return count, err
}

func (r *ModelsPostgres) AddModel(model *sneakerq.Model) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO %s (title, shop_id, price, page_url) VALUES ($1, $2, $3, $4)", modelsTable)
	_, err = tx.Exec(query, model.Title, model.ShopId, model.Price, model.PageUrl)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
