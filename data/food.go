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
	Name        string    `json:"name" db:"name" validate:"min=3"`
	KCAL        float32   `json:"kcal" db:"kcal"`
	Protein     float32   `json:"protein" db:"protein"`
	Carbs       float32   `json:"carbs" db:"carbs"`
	Fat         float32   `json:"fat" db:"fat"`
	Saturated   float32   `json:"saturated" db:"saturated"`
	Unsaturated float32   `json:"unsaturated" db:"unsaturated"`
	Fiber       float32   `json:"fiber" db:"fiber"`
	Sugars      float32   `json:"sugars" db:"sugars"`
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
type FoodTableDto struct {
	tableName   struct{}     `pg:"core.food,alias:f"`
	Id          uuid.UUID    `json:"id" pg:"id"`
	Timestamp   time.Time    `json:"timestamp" pg:"timestamp"`
	UserId      uuid.UUID    `json:"-" pg:"user"`
	FoodTypeId  uuid.UUID    `json:"-" pg:"food_type"`
	BrandId     uuid.UUID    `json:"-" pg:"brand"`
	User        *AuthDto     `json:"user" pg:"fk:user,rel:has-one"`
	FoodType    *FoodTypeDto `json:"foodtype" pg:"fk:food_type,rel:has-one"`
	Brand       *BrandDto    `json:"brand" pg:"fk:brand,rel:has-one"`
	Name        string       `json:"name" pg:"name" validate:"min=3"`
	KCAL        float32      `json:"kcal" pg:"kcal"`
	Protein     float32      `json:"protein" pg:"protein"`
	Carbs       float32      `json:"carbs" pg:"carbs"`
	Fat         float32      `json:"fat" pg:"fat"`
	Saturated   float32      `json:"saturated" pg:"saturated"`
	Unsaturated float32      `json:"unsaturated" pg:"unsaturated"`
	Fiber       float32      `json:"fiber" pg:"fiber"`
	Sugars      float32      `json:"sugars" pg:"sugars"`
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
type FoodFilterDto struct {
	Id       uuid.UUID `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Category uuid.UUID `json:"category" db:"category"`
	Take     uint16    `json:"take"`
	Skip     uint16    `json:"skip"`
}

func NewFood(name string, kcal float32, protein float32, carbs float32, fat float32, saturated float32, unstaturated float32, fiber float32, sugars float32, user, foodType, brand uuid.UUID) (*FoodDto, error) {
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

func (d *DataConn) ListFoods(pageIndex, pageSize int) ([]FoodTableDto, error) {
	var foods []FoodTableDto

	err := d.DB.Model(&foods).
		Relation("User").
		Relation("FoodType").
		Relation("Brand").
		Limit(pageSize).
		Offset(pageSize * pageIndex).
		Select()
	if err != nil {
		return nil, err
	}

	return foods, nil
}

func (d *DataConn) CountFoods() (int, error) {
	var foods []FoodTableDto

	count, err := d.DB.Model(&foods).
		Relation("User").
		Relation("FoodType").
		Relation("Brand").
		Count()

	if err != nil {
		return 0, err
	}

	return count, nil
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

func (d *DataConn) EditFood(name string, kcal float32, protein float32, carbs float32, fat float32, saturated float32, unstaturated float32, fiber float32, sugars float32, brand, foodtype, id uuid.UUID) error {
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
