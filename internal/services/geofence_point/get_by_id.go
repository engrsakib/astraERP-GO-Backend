package geofence_point

import repo "github.com/engrsakib/erp-system/internal/repository/geofence_point"

type GetGeofencePointByIDService struct {
    Repo repo.GeofencePointDAO
}

func NewGetGeofencePointByIDService(r repo.GeofencePointDAO) *GetGeofencePointByIDService {
    return &GetGeofencePointByIDService{Repo: r}
}

func (s *GetGeofencePointByIDService) Execute(id int64) (interface{}, error) {
    return s.Repo.GetByID(id)
}
