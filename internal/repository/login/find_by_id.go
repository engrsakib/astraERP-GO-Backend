package login

import (
	"github.com/engrsakib/erp-system/internal/models"
)


func (r *LoginRepository) FindByID(id int64) (*models.User, error) {
	var user models.User
	
	err := r.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}