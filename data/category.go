package data

import (
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

//lint:ignore U1000 Ignore unused function temporarily for debugging
type CategoryDto struct {
	tableName struct{}  `pg:"core.category,alias:c"`
	Id        uuid.UUID `json:"id" db:"id"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
	User      uuid.UUID `json:"user" db:"user"`
	Name      string    `json:"name" db:"name" validate:"min=3"`
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
type CategoryFilterDto struct {
	Id   uuid.UUID `json:"id" db:"id"`
	Name string    `json:"name" db:"name"`
	Take uint16    `json:"take"`
	Skip uint16    `json:"skip"`
}

func NewCategoryDto(user uuid.UUID, name string) (*CategoryDto, error) {
	validate := validator.New()

	category := &CategoryDto{
		User: user,
		Name: name,
	}

	err := validate.Struct(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func NewCategoryFilterDto(id uuid.UUID, name string) (*CategoryFilterDto, error) {
	validate := validator.New()

	filter := &CategoryFilterDto{
		Id:   id,
		Name: name,
	}

	err := validate.Struct(filter)
	if err != nil {
		return nil, err
	}

	return filter, nil
}

func (d *DataConn) ListCategories(pageIndex, pageSize int) ([]CategoryDto, error) {
	var categories []CategoryDto

	err := d.DB.Model(&categories).Limit(pageSize).Offset(pageIndex * pageSize).Select()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (d *DataConn) CountCategories() (int, error) {
	var categories []CategoryDto

	count, err := d.DB.Model(&categories).Count()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (d *DataConn) GetCategoryById(id uuid.UUID) (CategoryDto, error) {
	var category CategoryDto

	err := d.DB.Model(&category).Where("id = ?", id).Select()

	if err != nil {
		return category, err
	}

	return category, nil
}

func (d *DataConn) CreateCategory(dto CategoryDto) error {
	_, err := d.DB.Model(&dto).Insert()

	pgErr, ok := err.(pg.Error)
	if ok && pgErr.IntegrityViolation() {
		return fmt.Errorf("name already exists")
	}

	return err
}

func (d *DataConn) EditCategory(id uuid.UUID, name string) error {
	var category CategoryDto
	_, err := d.DB.Model(&category).Set("name = ?", name).Where("id = ?", id).Update()
	return err
}

func (d *DataConn) DeleteCategory(id uuid.UUID) error {
	var category CategoryDto
	_, err := d.DB.Model(&category).Where("id = ?", id).Delete()
	return err
}
