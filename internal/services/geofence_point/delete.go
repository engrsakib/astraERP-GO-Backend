package geofence_point

import repo "github.com/engrsakib/erp-system/internal/repository/geofence_point"

type DeleteGeofencePointService struct {
    Repo repo.GeofencePointDAO
}

func NewDeleteGeofencePointService(r repo.GeofencePointDAO) *DeleteGeofencePointService {
    return &DeleteGeofencePointService{Repo: r}
}

func (s *DeleteGeofencePointService) Execute(id int64) error {
    return s.Repo.Delete(id)
}
