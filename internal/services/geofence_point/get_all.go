package geofence_point

import (
    "github.com/engrsakib/erp-system/internal/dto/geofence_point"
    repo "github.com/engrsakib/erp-system/internal/repository/geofence_point"
)

type GetAllGeofencePointService struct {
    Repo repo.GeofencePointDAO
}

func NewGetAllGeofencePointService(r repo.GeofencePointDAO) *GetAllGeofencePointService {
    return &GetAllGeofencePointService{Repo: r}
}

func (s *GetAllGeofencePointService) Execute(q geofence_point.GeofencePointSearchQuery) (interface{}, interface{}, error) {

    if q.Page == 0 {
        q.Page = 1
    }
    if q.Limit == 0 {
        q.Limit = 10
    }

    data, total, err := s.Repo.GetAll(q.Page, q.Limit, q.GeofenceID)
    if err != nil {
        return nil, nil, err
    }

    meta := map[string]interface{}{
        "page":  q.Page,
        "limit": q.Limit,
        "total": total,
    }

    return data, meta, nil
}
