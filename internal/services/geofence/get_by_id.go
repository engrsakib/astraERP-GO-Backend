package geofence

import repo "github.com/engrsakib/erp-system/internal/repository/geofence"

type GetGeofenceByIDService struct {
    Repo *repo.GeofenceRepository
}

func NewGetGeofenceByIDService(r *repo.GeofenceRepository) *GetGeofenceByIDService {
    return &GetGeofenceByIDService{Repo: r}
}

func (s *GetGeofenceByIDService) Execute(id int64) (interface{}, error) {
    return s.Repo.GetByID(id)
}
