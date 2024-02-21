package data

import (
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

//lint:ignore U1000 Ignore unused function temporarily for debugging
type FoodType struct {
	tableName struct{}  `pg:"core.food_type,alias:b"`
	Id        uuid.UUID `json:"id" db:"id"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
	User      uuid.UUID `json:"user" db:"user"`
	Category  uuid.UUID `json:"category" db:"category"`
	Name      string    `json:"name" db:"name" validate:"max=20,min=3"`
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
type FoodTypeFilterDto struct {
	Id       uuid.UUID `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Category uuid.UUID `json:"category" db:"category"`
	Take     uint16    `json:"take"`
	Skip     uint16    `json:"skip"`
}

func NewFoodType(user uuid.UUID, name string, category uuid.UUID) (*FoodType, error) {
	validate := validator.New()

	foodType := &FoodType{
		User:     user,
		Name:     name,
		Category: category,
	}

	err := validate.Struct(foodType)
	if err != nil {
		return nil, err
	}

	return foodType, nil
}

func NewFoodTypeFilterDto(id uuid.UUID, name string, category uuid.UUID) (*FoodTypeFilterDto, error) {
	validate := validator.New()

	filter := &FoodTypeFilterDto{
		Id:       id,
		Name:     name,
		Category: category,
	}

	err := validate.Struct(filter)
	if err != nil {
		return nil, err
	}

	return filter, nil
}

func (d *DataConn) ListFoodTypes() ([]FoodType, error) {
	var foodTypes []FoodType

	err := d.DB.Model(&foodTypes).Select()
	if err != nil {
		return nil, err
	}

	return foodTypes, nil
}

func (d *DataConn) GetFoodTypeById(id uuid.UUID) (FoodType, error) {
	var foodType FoodType

	err := d.DB.Model(&foodType).Where("id = ?", id).Select()

	if err != nil {
		return foodType, err
	}

	return foodType, nil
}

func (d *DataConn) CreateFoodType(dto FoodType) error {
	_, err := d.DB.Model(&dto).Insert()

	pgErr, ok := err.(pg.Error)
	if ok && pgErr.IntegrityViolation() {
		return fmt.Errorf("name already exists")
	}

	return err
}

func (d *DataConn) EditFoodType(id uuid.UUID, name string, category uuid.UUID) error {
	var foodType FoodType
	_, err := d.DB.Model(&foodType).Set("name = ?", name).Set("category = ?", category).Where("id = ?", id).Update()
	return err
}

func (d *DataConn) DeleteFoodType(id uuid.UUID) error {
	var foodType FoodType
	_, err := d.DB.Model(&foodType).Where("id = ?", id).Delete()
	return err
}