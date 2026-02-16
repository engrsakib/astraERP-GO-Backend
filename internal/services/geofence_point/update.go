package geofence_point

import (
    "github.com/engrsakib/erp-system/internal/dto/geofence_point"
    repo "github.com/engrsakib/erp-system/internal/repository/geofence_point"
)

type UpdateGeofencePointService struct {
    Repo repo.GeofencePointDAO
}

func NewUpdateGeofencePointService(r repo.GeofencePointDAO) *UpdateGeofencePointService {
    return &UpdateGeofencePointService{Repo: r}
}

func (s *UpdateGeofencePointService) Execute(id int64, req geofence_point.GeofencePointRequest) error {

    point, err := s.Repo.GetByID(id)
    if err != nil {
        return err
    }

    point.GeofenceID = req.GeofenceID
    point.Latitude = req.Latitude
    point.Longitude = req.Longitude
    point.SeqOrder = req.SeqOrder

    return s.Repo.Update(point)
}
