package geofence

import (
    "github.com/engrsakib/erp-system/internal/dto/geofence"
    repo "github.com/engrsakib/erp-system/internal/repository/geofence"
)

type GetAllGeofenceService struct {
    Repo *repo.GeofenceRepository
}

func NewGetAllGeofenceService(r *repo.GeofenceRepository) *GetAllGeofenceService {
    return &GetAllGeofenceService{Repo: r}
}

func (s *GetAllGeofenceService) Execute(q geofence.GeofenceSearchQuery) (interface{}, interface{}, error) {

    if q.Page == 0 {
        q.Page = 1
    }
    if q.Limit == 0 {
        q.Limit = 10
    }

    data, total, err := s.Repo.GetAll(q.Page, q.Limit, q.Search)
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
