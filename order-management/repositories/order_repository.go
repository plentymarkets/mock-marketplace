package repositories

import (
	"gorm.io/gorm"
	"order-management/models"
)

type OrderRepository struct {
	DatabaseConnection *gorm.DB
}

func NewOrderRepository(databaseConnection *gorm.DB) OrderRepository {
	repository := OrderRepository{}
	repository.DatabaseConnection = databaseConnection
	return repository
}

func (OrderRepository OrderRepository) FindAll(offset *int, limit *int) (*[]models.Order, error) {
	OrderRepository.DatabaseConnection.Preload("OrderItems")

	if offset != nil {
		OrderRepository.DatabaseConnection = OrderRepository.DatabaseConnection.Offset(*offset)

		if OrderRepository.DatabaseConnection.Error != nil {
			return nil, OrderRepository.DatabaseConnection.Error
		}
	}

	if limit != nil {
		OrderRepository.DatabaseConnection = OrderRepository.DatabaseConnection.Limit(*limit)

		if OrderRepository.DatabaseConnection.Error != nil {
			return nil, OrderRepository.DatabaseConnection.Error
		}
	}

	var orders []models.Order
	OrderRepository.DatabaseConnection.Find(&orders)

	if OrderRepository.DatabaseConnection.Error != nil {
		return nil, OrderRepository.DatabaseConnection.Error
	}

	return &orders, nil
}

func (OrderRepository OrderRepository) FindById(id int, offset *int, limit *int) (*models.Order, error) {
	var order models.Order
	OrderRepository.DatabaseConnection.Preload("OrderItems")

	if offset != nil {
		OrderRepository.DatabaseConnection = OrderRepository.DatabaseConnection.Offset(*offset)

		if OrderRepository.DatabaseConnection.Error != nil {
			return nil, OrderRepository.DatabaseConnection.Error
		}
	}

	if limit != nil {
		OrderRepository.DatabaseConnection = OrderRepository.DatabaseConnection.Limit(*limit)

		if OrderRepository.DatabaseConnection.Error != nil {
			return nil, OrderRepository.DatabaseConnection.Error
		}
	}

	OrderRepository.DatabaseConnection.Find(&order, id)

	return &order, nil
}

func (OrderRepository OrderRepository) FindOneById(id int) (*models.Order, error) {
	OrderRepository.DatabaseConnection.Preload("OrderItems")

	var order models.Order
	OrderRepository.DatabaseConnection.First(&order, id)

	if OrderRepository.DatabaseConnection.Error != nil {
		return nil, OrderRepository.DatabaseConnection.Error
	}

	return &order, nil
}

func (OrderRepository OrderRepository) FindOneByField(field string, value string) (*models.Order, error) {
	var order models.Order
	OrderRepository.DatabaseConnection.Preload("OrderItems")

	OrderRepository.DatabaseConnection.Where(field+" = ?", value).First(&order)

	if OrderRepository.DatabaseConnection.Error != nil {
		return nil, OrderRepository.DatabaseConnection.Error
	}

	return &order, nil
}

func (OrderRepository OrderRepository) FindOneByFields(fields map[string]string) (*models.Order, error) {
	var order models.Order
	OrderRepository.DatabaseConnection = OrderRepository.DatabaseConnection.Preload("OrderItems")

	for key, val := range fields {
		OrderRepository.DatabaseConnection = OrderRepository.DatabaseConnection.Where(key+" = ?", val)

		if OrderRepository.DatabaseConnection.Error != nil {
			return nil, OrderRepository.DatabaseConnection.Error
		}
	}

	OrderRepository.DatabaseConnection.First(&order)

	if OrderRepository.DatabaseConnection.Error != nil {
		return nil, OrderRepository.DatabaseConnection.Error
	}

	return &order, nil
}

func (OrderRepository OrderRepository) FindByField(field string, value string, offset *int, limit *int) (*[]models.Order, error) {
	var orders []models.Order
	OrderRepository.DatabaseConnection = OrderRepository.DatabaseConnection.Preload("OrderItems")

	if offset != nil {
		OrderRepository.DatabaseConnection = OrderRepository.DatabaseConnection.Offset(*offset)

		if OrderRepository.DatabaseConnection.Error != nil {
			return nil, OrderRepository.DatabaseConnection.Error
		}
	}

	if limit != nil {
		OrderRepository.DatabaseConnection = OrderRepository.DatabaseConnection.Limit(*limit)

		if OrderRepository.DatabaseConnection.Error != nil {
			return nil, OrderRepository.DatabaseConnection.Error
		}
	}

	OrderRepository.DatabaseConnection.Where(field+" = ?", value).Find(&orders)

	if OrderRepository.DatabaseConnection.Error != nil {
		return nil, OrderRepository.DatabaseConnection.Error
	}

	return &orders, nil
}

func (OrderRepository OrderRepository) FindByFields(fields map[string]string, offset *int, limit *int) (*[]models.Order, error) {
	var orders []models.Order
	OrderRepository.DatabaseConnection.Preload("OrderItems")

	if offset != nil {
		OrderRepository.DatabaseConnection = OrderRepository.DatabaseConnection.Offset(*offset)

		if OrderRepository.DatabaseConnection.Error != nil {
			return nil, OrderRepository.DatabaseConnection.Error
		}
	}

	if limit != nil {
		OrderRepository.DatabaseConnection = OrderRepository.DatabaseConnection.Limit(*limit)

		if OrderRepository.DatabaseConnection.Error != nil {
			return nil, OrderRepository.DatabaseConnection.Error
		}
	}

	for key, val := range fields {
		OrderRepository.DatabaseConnection = OrderRepository.DatabaseConnection.Where(key+" = ?", val)

		if OrderRepository.DatabaseConnection.Error != nil {
			return nil, OrderRepository.DatabaseConnection.Error
		}
	}

	OrderRepository.DatabaseConnection.Find(&orders)

	if OrderRepository.DatabaseConnection.Error != nil {
		return nil, OrderRepository.DatabaseConnection.Error
	}

	return &orders, nil
}
