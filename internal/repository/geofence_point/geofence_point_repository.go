package geofence_point

import (
    "github.com/engrsakib/erp-system/internal/models"
    "gorm.io/gorm"
)

type geofencePointRepository struct {
    DB *gorm.DB
}

func NewGeofencePointRepository(db *gorm.DB) GeofencePointDAO {
    return &geofencePointRepository{DB: db}
}

func (r *geofencePointRepository) Create(point *models.GeofencePoint) error {
    return r.DB.Create(point).Error
}

func (r *geofencePointRepository) GetAll(page, limit int, geofenceID int64) ([]models.GeofencePoint, int64, error) {
    var points []models.GeofencePoint
    var total int64

    query := r.DB.Model(&models.GeofencePoint{})

    if geofenceID != 0 {
        query = query.Where("geofence_id = ?", geofenceID)
    }

    query.Count(&total)

    offset := (page - 1) * limit
    err := query.Offset(offset).Limit(limit).Find(&points).Error

    return points, total, err
}

func (r *geofencePointRepository) GetByID(id int64) (*models.GeofencePoint, error) {
    var point models.GeofencePoint
    err := r.DB.First(&point, id).Error
    return &point, err
}

func (r *geofencePointRepository) Update(point *models.GeofencePoint) error {
    return r.DB.Save(point).Error
}

func (r *geofencePointRepository) Delete(id int64) error {
    return r.DB.Delete(&models.GeofencePoint{}, id).Error
}
