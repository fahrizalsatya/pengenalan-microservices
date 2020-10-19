package database

import "gorm.io/gorm"

//A Menu represents [ID, MenuName, Price]
type Menu struct {
	ID       int    `json:"-" gorm:"primary_key"`
	MenuName string `json:"menu_name"`
	Price    int    `json:"price"`
}

//Insert db as database in gorm.DB, based on Menu type as menu's column
func (menu *Menu) Insert(db *gorm.DB) error {
	result := db.Create(menu)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

//GetAll db database, to display db content based on Menu type as menu's column
func (menu *Menu) GetAll(db *gorm.DB) ([]Menu, error) {
	var menus []Menu
	result := db.Find(&menus)
	if result.Error != nil {
		return nil, result.Error
	}
	return menus, nil
}
