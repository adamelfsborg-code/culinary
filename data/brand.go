package data

import (
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

//lint:ignore U1000 Ignore unused function temporarily for debugging
type BrandDto struct {
	tableName struct{}  `pg:"core.brand,alias:b"`
	Id        uuid.UUID `json:"id" db:"id"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
	User      uuid.UUID `json:"user" db:"user"`
	Name      string    `json:"name" db:"name" validate:"min=3"`
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
type BrandFilterDto struct {
	Id   uuid.UUID `json:"id" db:"id"`
	Name string    `json:"name" db:"name"`
	Take uint16    `json:"take"`
	Skip uint16    `json:"skip"`
}

func NewBrandDto(user uuid.UUID, name string) (*BrandDto, error) {
	validate := validator.New()

	brand := &BrandDto{
		User: user,
		Name: name,
	}

	err := validate.Struct(brand)
	if err != nil {
		return nil, err
	}

	return brand, nil
}

func NewBrandFilterDto(id uuid.UUID, name string) (*BrandFilterDto, error) {
	validate := validator.New()

	filter := &BrandFilterDto{
		Id:   id,
		Name: name,
	}

	err := validate.Struct(filter)
	if err != nil {
		return nil, err
	}

	return filter, nil
}

func (d *DataConn) ListBrands(pageIndex, pageSize int) ([]BrandDto, error) {
	var brands []BrandDto

	err := d.DB.Model(&brands).Limit(pageSize).Offset(pageIndex * pageSize).Select()
	if err != nil {
		return nil, err
	}

	return brands, nil
}

func (d *DataConn) CountBrands() (int, error) {
	var brands []BrandDto

	count, err := d.DB.Model(&brands).Count()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (d *DataConn) GetBrandById(id uuid.UUID) (BrandDto, error) {
	var brand BrandDto

	err := d.DB.Model(&brand).Where("id = ?", id).Select()

	if err != nil {
		return brand, err
	}

	return brand, nil
}

func (d *DataConn) CreateBrand(dto BrandDto) error {
	_, err := d.DB.Model(&dto).Insert()

	pgErr, ok := err.(pg.Error)
	if ok && pgErr.IntegrityViolation() {
		return fmt.Errorf("name already exists")
	}

	return err
}

func (d *DataConn) EditBrand(id uuid.UUID, name string) error {
	var brand BrandDto
	_, err := d.DB.Model(&brand).Set("name = ?", name).Where("id = ?", id).Update()
	return err
}

func (d *DataConn) DeleteBrand(id uuid.UUID) error {
	var brand BrandDto
	_, err := d.DB.Model(&brand).Where("id = ?", id).Delete()
	return err
}
