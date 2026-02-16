package geofence

import repo "github.com/engrsakib/erp-system/internal/repository/geofence"

type DeleteGeofenceService struct {
    Repo *repo.GeofenceRepository
}

func NewDeleteGeofenceService(r *repo.GeofenceRepository) *DeleteGeofenceService {
    return &DeleteGeofenceService{Repo: r}
}

func (s *DeleteGeofenceService) Execute(id int64) error {
    return s.Repo.Delete(id)
}
