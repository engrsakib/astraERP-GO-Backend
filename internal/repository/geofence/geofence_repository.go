package geofence

import (
    "github.com/engrsakib/erp-system/internal/models"
    "gorm.io/gorm"
)

type GeofenceRepository struct {
    DB *gorm.DB
}

func NewGeofenceRepository(db *gorm.DB) *GeofenceRepository {
    return &GeofenceRepository{DB: db}
}

// Create
func (r *GeofenceRepository) Create(geo *models.Geofence) error {
    return r.DB.Create(geo).Error
}

// Get All + Search + Pagination
func (r *GeofenceRepository) GetAll(page, limit int, search string) ([]models.Geofence, int64, error) {
    var geofences []models.Geofence
    var total int64

    query := r.DB.Model(&models.Geofence{})

    if search != "" {
        query = query.Where("name LIKE ?", "%"+search+"%")
    }

    query.Count(&total)

    offset := (page - 1) * limit
    err := query.Offset(offset).Limit(limit).Find(&geofences).Error

    return geofences, total, err
}

// Get By ID
func (r *GeofenceRepository) GetByID(id int64) (*models.Geofence, error) {
    var geo models.Geofence
    err := r.DB.First(&geo, id).Error
    return &geo, err
}

// Update
func (r *GeofenceRepository) Update(geo *models.Geofence) error {
    return r.DB.Save(geo).Error
}

// Soft Delete
func (r *GeofenceRepository) Delete(id int64) error {
    return r.DB.Delete(&models.Geofence{}, id).Error
}
