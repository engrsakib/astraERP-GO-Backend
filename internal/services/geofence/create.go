package geofence

import (
    "github.com/engrsakib/erp-system/internal/dto/geofence"
    "github.com/engrsakib/erp-system/internal/models"
    repo "github.com/engrsakib/erp-system/internal/repository/geofence"
)

type CreateGeofenceService struct {
    Repo *repo.GeofenceRepository
}

func NewCreateGeofenceService(r *repo.GeofenceRepository) *CreateGeofenceService {
    return &CreateGeofenceService{Repo: r}
}

func (s *CreateGeofenceService) Execute(req geofence.GeofenceRequest) (*models.Geofence, error) {

    geo := &models.Geofence{
        Name:        req.Name,
        Description: req.Description,
    }

    err := s.Repo.Create(geo)
    return geo, err
}
