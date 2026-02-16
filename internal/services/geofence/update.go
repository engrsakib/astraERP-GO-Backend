package geofence

import (
    "github.com/engrsakib/erp-system/internal/dto/geofence"
    repo "github.com/engrsakib/erp-system/internal/repository/geofence"
)

type UpdateGeofenceService struct {
    Repo *repo.GeofenceRepository
}

func NewUpdateGeofenceService(r *repo.GeofenceRepository) *UpdateGeofenceService {
    return &UpdateGeofenceService{Repo: r}
}

func (s *UpdateGeofenceService) Execute(id int64, req geofence.GeofenceRequest) error {

    geo, err := s.Repo.GetByID(id)
    if err != nil {
        return err
    }

    geo.Name = req.Name
    geo.Description = req.Description

    return s.Repo.Update(geo)
}
