package geofence_point

import (
    "github.com/engrsakib/erp-system/internal/dto/geofence_point"
    "github.com/engrsakib/erp-system/internal/models"
    repo "github.com/engrsakib/erp-system/internal/repository/geofence_point"
)

type CreateGeofencePointService struct {
    Repo repo.GeofencePointDAO
}

func NewCreateGeofencePointService(r repo.GeofencePointDAO) *CreateGeofencePointService {
    return &CreateGeofencePointService{Repo: r}
}

func (s *CreateGeofencePointService) Execute(req geofence_point.GeofencePointRequest) (*models.GeofencePoint, error) {

    point := &models.GeofencePoint{
        GeofenceID: req.GeofenceID,
        Latitude:   req.Latitude,
        Longitude:  req.Longitude,
        SeqOrder:   req.SeqOrder,
    }

    err := s.Repo.Create(point)
    return point, err
}
