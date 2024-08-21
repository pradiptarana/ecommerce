package product

import (
	"database/sql"
	"errors"
	"fmt"

	model "github.com/pradiptarana/product/model/product"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (tr *ProductRepository) GetProducts(filter *model.GetProductFilter) ([]*model.Product, error) {
	whereQuery := []string{
		"is_active = 1 ",
	}
	params := []any{}
	if filter.Name != "" {
		params = append(params, filter.Name)
		whereQuery = append(whereQuery, "name = ? ")
	}
	if filter.Category != 0 {
		params = append(params, filter.Category)
		whereQuery = append(whereQuery, "category_id = ? ")
	}
	query := "SELECT * FROM product"
	for k, v := range whereQuery {
		if k == 0 {
			query = query + " WHERE "
		} else {
			query = query + "AND "
		}
		query = query + v
	}
	query = query + "order by id desc limit ? offset ?"
	// params = append(params, filter.SortBy)
	// params = append(params, filter.OrderType)
	params = append(params, filter.PageSize)
	params = append(params, filter.PageNum)
	fmt.Println(query)
	stmt, err := tr.db.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(params...)
	rows, err := stmt.Query(params...)
	fmt.Println(rows)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()
	var result []*model.Product
	for rows.Next() {
		var each = &model.Product{}
		var err = rows.Scan(&each.Id, &each.Name, &each.Description, &each.Category, &each.Price, &each.IsActive)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		result = append(result, each)
	}
	return result, nil
}

func (tr *ProductRepository) GetLatestProducts() ([]*model.Product, error) {
	query := `
		SELECT p.id, p.name, p.price, COALESCE(SUM(wp.stock), 0) as stock
		FROM products p
		LEFT JOIN warehouse_products wp ON wp.product_id = p.id
		LEFT JOIN warehouses w ON wp.warehouse_id = w.id
		WHERE w.status = 'active'
		GROUP BY p.id, p.name, p.price
	`

	rows, err := tr.db.Query(query)
	if err != nil {
		return nil, errors.New("unable to fetch products")
	}
	defer rows.Close()

	var products []*model.Product
	for rows.Next() {
		var product *model.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock)
		if err != nil {
			return nil, errors.New("unable to fetch products")
		}
		products = append(products, product)
	}

	// var us model.Product
	// if err := rows.Scan(&us.Id, &us.Name, &us.Description, &us.IsActive, &us.Price, &us.Category); err != nil {
	// 	return &us, fmt.Errorf("error when get user")
	// }
	return products, nil
}
