package warehouse

import (
	"database/sql"
	"errors"
)

type WarehouseRepository struct {
	db *sql.DB
}

func NewWarehouseRepository(db *sql.DB) *WarehouseRepository {
	return &WarehouseRepository{db}
}

func (tr *WarehouseRepository) TransferStock(fromWarehouseID int, toWarehouseID int, productID int, quantity int) error {

	tx, err := tr.db.Begin()
	if err != nil {
		return errors.New("Transaction begin failed")
	}

	_, err = tx.Exec("UPDATE warehouse_products SET stock = stock - ? WHERE warehouse_id = ? AND product_id = ? AND stock >= ?", quantity, fromWarehouseID, productID, quantity)
	if err != nil {
		tx.Rollback()
		return errors.New("Stock deduction failed")
	}

	_, err = tx.Exec("UPDATE warehouse_products SET stock = stock + ? WHERE warehouse_id = ? AND product_id = ?", quantity, toWarehouseID, productID)
	if err != nil {
		tx.Rollback()
		return errors.New("Stock addotion failed")
	}

	err = tx.Commit()
	if err != nil {
		return errors.New("Transaction commit failed")
	}

	return nil
}

func (tr *WarehouseRepository) ActivateWarehouse(warehouseID int) error {
	_, err := tr.db.Exec("UPDATE warehouses SET active=TRUE WHERE id=$1", warehouseID)
	if err != nil {
		return errors.New("Error activating warehouse")
	}

	return nil
}

func (tr *WarehouseRepository) DeactivateWarehouse(warehouseID int) error {
	_, err := tr.db.Exec("UPDATE warehouses SET active=FALSE WHERE id=$1", warehouseID)
	if err != nil {
		return errors.New("Error deactivating warehouse")
	}

	return nil
}
