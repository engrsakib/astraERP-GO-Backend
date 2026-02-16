package geofence_point

import "github.com/engrsakib/erp-system/internal/models"

type GeofencePointDAO interface {
    Create(point *models.GeofencePoint) error
    GetAll(page, limit int, geofenceID int64) ([]models.GeofencePoint, int64, error)
    GetByID(id int64) (*models.GeofencePoint, error)
    Update(point *models.GeofencePoint) error
    Delete(id int64) error
}
