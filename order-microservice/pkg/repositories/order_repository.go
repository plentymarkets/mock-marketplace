package repositories

import (
	"gorm.io/gorm"
	"order-microservice/pkg/models"
)

type OrderRepository struct {
	Database *gorm.DB
}

func NewOrderRepository(databaseConnection *gorm.DB) OrderRepository {
	repository := OrderRepository{}
	repository.Database = databaseConnection
	return repository
}

func (OrderRepository OrderRepository) FindAll(offset *int, limit *int) (*[]models.Order, error) {
	OrderRepository.Database.Preload("OrderItems")

	if offset != nil {
		OrderRepository.Database = OrderRepository.Database.Offset(*offset)

		if OrderRepository.Database.Error != nil {
			return nil, OrderRepository.Database.Error
		}
	}

	if limit != nil {
		OrderRepository.Database = OrderRepository.Database.Limit(*limit)

		if OrderRepository.Database.Error != nil {
			return nil, OrderRepository.Database.Error
		}
	}

	var orders []models.Order
	OrderRepository.Database.Find(&orders)

	if OrderRepository.Database.Error != nil {
		return nil, OrderRepository.Database.Error
	}

	return &orders, nil
}

func (OrderRepository OrderRepository) FindById(id int, offset *int, limit *int) (*models.Order, error) {
	var order models.Order
	OrderRepository.Database.Preload("OrderItems")

	if offset != nil {
		OrderRepository.Database = OrderRepository.Database.Offset(*offset)

		if OrderRepository.Database.Error != nil {
			return nil, OrderRepository.Database.Error
		}
	}

	if limit != nil {
		OrderRepository.Database = OrderRepository.Database.Limit(*limit)

		if OrderRepository.Database.Error != nil {
			return nil, OrderRepository.Database.Error
		}
	}

	OrderRepository.Database.Find(&order, id)

	return &order, nil
}

func (OrderRepository OrderRepository) FindOneById(id int) (*models.Order, error) {
	OrderRepository.Database.Preload("OrderItems")

	var order models.Order
	OrderRepository.Database.First(&order, id)

	if OrderRepository.Database.Error != nil {
		return nil, OrderRepository.Database.Error
	}

	return &order, nil
}

func (OrderRepository OrderRepository) FindOneByField(field string, value string) (*models.Order, error) {
	var order models.Order
	OrderRepository.Database.Preload("OrderItems")

	OrderRepository.Database.Where(field+" = ?", value).First(&order)

	if OrderRepository.Database.Error != nil {
		return nil, OrderRepository.Database.Error
	}

	if OrderRepository.Database.RowsAffected == 0 {
		return nil, nil
	}

	return &order, nil
}

func (OrderRepository OrderRepository) FindOneByFields(fields map[string]string) (*models.Order, error) {
	var order models.Order
	OrderRepository.Database = OrderRepository.Database.Preload("OrderItems")

	for key, val := range fields {
		OrderRepository.Database = OrderRepository.Database.Where(key+" = ?", val)

		if OrderRepository.Database.Error != nil {
			return nil, OrderRepository.Database.Error
		}
	}

	OrderRepository.Database.First(&order)

	if OrderRepository.Database.Error != nil {
		return nil, OrderRepository.Database.Error
	}

	return &order, nil
}

func (OrderRepository OrderRepository) FindByField(field string, value string, offset *int, limit *int) (*[]models.Order, error) {
	var orders []models.Order
	OrderRepository.Database = OrderRepository.Database.Preload("OrderItems")

	if offset != nil {
		OrderRepository.Database = OrderRepository.Database.Offset(*offset)

		if OrderRepository.Database.Error != nil {
			return nil, OrderRepository.Database.Error
		}
	}

	if limit != nil {
		OrderRepository.Database = OrderRepository.Database.Limit(*limit)

		if OrderRepository.Database.Error != nil {
			return nil, OrderRepository.Database.Error
		}
	}

	OrderRepository.Database.Where(field+" = ?", value).Find(&orders)

	if OrderRepository.Database.Error != nil {
		return nil, OrderRepository.Database.Error
	}

	if OrderRepository.Database.RowsAffected == 0 {
		return nil, nil
	}

	return &orders, nil
}

func (OrderRepository OrderRepository) FindByFields(fields map[string]string, offset *int, limit *int) (*[]models.Order, error) {
	var orders []models.Order
	OrderRepository.Database.Preload("OrderItems")

	if offset != nil {
		OrderRepository.Database = OrderRepository.Database.Offset(*offset)

		if OrderRepository.Database.Error != nil {
			return nil, OrderRepository.Database.Error
		}
	}

	if limit != nil {
		OrderRepository.Database = OrderRepository.Database.Limit(*limit)

		if OrderRepository.Database.Error != nil {
			return nil, OrderRepository.Database.Error
		}
	}

	for key, val := range fields {
		OrderRepository.Database = OrderRepository.Database.Where(key+" = ?", val)

		if OrderRepository.Database.Error != nil {
			return nil, OrderRepository.Database.Error
		}
	}

	OrderRepository.Database.Find(&orders)

	if OrderRepository.Database.Error != nil {
		return nil, OrderRepository.Database.Error
	}

	return &orders, nil
}

func (OrderRepository OrderRepository) Create(order *models.Order) error {
	OrderRepository.Database.Create(&order)

	if OrderRepository.Database.Error != nil {
		return OrderRepository.Database.Error
	}

	return nil
}

func (OrderRepository OrderRepository) Update(order *models.Order) error {
	OrderRepository.Database.Save(&order)

	if OrderRepository.Database.Error != nil {
		return OrderRepository.Database.Error
	}

	return nil
}

func (OrderRepository OrderRepository) FindByTimeRange(field string, start string, end string, offset *int, limit *int) (*[]models.Order, error) {
	var orders []models.Order
	OrderRepository.Database.Preload("OrderItems")

	if offset != nil {
		OrderRepository.Database = OrderRepository.Database.Offset(*offset)

		if OrderRepository.Database.Error != nil {
			return nil, OrderRepository.Database.Error
		}
	}

	if limit != nil {
		OrderRepository.Database = OrderRepository.Database.Limit(*limit)

		if OrderRepository.Database.Error != nil {
			return nil, OrderRepository.Database.Error
		}
	}

	OrderRepository.Database.Where(field+"BETWEEN ? AND ?", start, end).Find(&orders)

	if OrderRepository.Database.Error != nil {
		return nil, OrderRepository.Database.Error
	}

	return &orders, nil
}
