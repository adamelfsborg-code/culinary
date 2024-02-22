package data

import (
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

//lint:ignore U1000 Ignore unused function temporarily for debugging
type FoodDto struct {
	tableName   struct{}  `pg:"core.food,alias:f"`
	Id          uuid.UUID `json:"id" db:"id"`
	Timestamp   time.Time `json:"timestamp" db:"timestamp"`
	User        uuid.UUID `json:"user" db:"user"`
	FoodType    uuid.UUID `json:"foodtype" db:"food_type"`
	Brand       uuid.UUID `json:"brand" db:"brand"`
	Name        string    `json:"name" db:"name" validate:"max=20,min=3"`
	KCAL        uint8     `json:"kcal" db:"kcal"`
	Protein     uint8     `json:"protein" db:"protein"`
	Carbs       uint8     `json:"carbs" db:"carbs"`
	Fat         uint8     `json:"fat" db:"fat"`
	Saturated   uint8     `json:"saturated" db:"saturated"`
	Unsaturated uint8     `json:"unsaturated" db:"unsaturated"`
	Fiber       uint8     `json:"fiber" db:"fiber"`
	Sugars      uint8     `json:"sugars" db:"sugars"`
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
type FoodFilterDto struct {
	Id       uuid.UUID `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Category uuid.UUID `json:"category" db:"category"`
	Take     uint16    `json:"take"`
	Skip     uint16    `json:"skip"`
}

func NewFood(name string, kcal uint8, protein uint8, carbs uint8, fat uint8, saturated uint8, unstaturated uint8, fiber uint8, sugars uint8, user, foodType, brand uuid.UUID) (*FoodDto, error) {
	validate := validator.New()

	food := &FoodDto{
		User:        user,
		Name:        name,
		FoodType:    foodType,
		Brand:       brand,
		KCAL:        kcal,
		Protein:     protein,
		Carbs:       carbs,
		Fat:         fat,
		Saturated:   saturated,
		Unsaturated: unstaturated,
		Fiber:       fiber,
		Sugars:      sugars,
	}

	err := validate.Struct(food)
	if err != nil {
		return nil, err
	}

	return food, nil
}

func NewFoodFilterDto(id uuid.UUID, name string) (*FoodFilterDto, error) {
	validate := validator.New()

	filter := &FoodFilterDto{
		Id:   id,
		Name: name,
	}

	err := validate.Struct(filter)
	if err != nil {
		return nil, err
	}

	return filter, nil
}

func (d *DataConn) ListFoods() ([]FoodDto, error) {
	var foods []FoodDto

	err := d.DB.Model(&foods).Select()
	if err != nil {
		return nil, err
	}

	return foods, nil
}

func (d *DataConn) GetFoodById(id uuid.UUID) (FoodDto, error) {
	var food FoodDto

	err := d.DB.Model(&food).Where("id = ?", id).Select()

	if err != nil {
		return food, err
	}

	return food, nil
}

func (d *DataConn) CreateFood(dto FoodDto) error {
	_, err := d.DB.Model(&dto).Insert()

	pgErr, ok := err.(pg.Error)
	if ok && pgErr.IntegrityViolation() {
		return fmt.Errorf("name already exists")
	}

	return err
}

func (d *DataConn) EditFood(name string, kcal uint8, protein uint8, carbs uint8, fat uint8, saturated uint8, unstaturated uint8, fiber uint8, sugars uint8, brand, foodtype, id uuid.UUID) error {
	var food FoodDto
	_, err := d.DB.Model(&food).
		Set("name = ?", name).
		Set("kcal = ?", kcal).
		Set("protein = ?", protein).
		Set("carbs = ?", carbs).
		Set("fat = ?", fat).
		Set("saturated = ?", saturated).
		Set("unsaturated = ?", unstaturated).
		Set("fiber = ?", fiber).
		Set("sugars = ?", sugars).
		Set("brand = ?", brand).
		Set("food_type = ?", foodtype).
		Where("id = ?", id).
		Update()
	return err
}

func (d *DataConn) DeleteFood(id uuid.UUID) error {
	var food FoodDto
	_, err := d.DB.Model(&food).Where("id = ?", id).Delete()
	return err
}
