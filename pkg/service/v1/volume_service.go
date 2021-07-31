package v1

import (
	"context"
	"database/sql"

	v1 "github.com/cardenasrjl/emem/pkg/api/v1"
	"github.com/cardenasrjl/emem/pkg/storage/mysqldb"
)

// VolumeService is implementation of v1.ToDoServiceServer proto interface
type VolumeService struct {
	database *mysqldb.DB
}

// NewMemServiceServer ...
func NeVolumeService(db *sql.DB) v1.VolumeServiceServer {
	st := mysqldb.NewDB(db)
	return &VolumeService{database: st}
}

func (v VolumeService) GetVolumes(context.Context, *v1.GetVolumesRequest) (*v1.GetVolumesResponse, error) {
	panic("implement me")
}

func (v VolumeService) GetVolume(context.Context, *v1.GetVolumeRequest) (*v1.GetVolumeResponse, error) {
	panic("implement me")
}

func (v VolumeService) DeleteVolume(context.Context, *v1.DeleteVolumeRequest) (*v1.DeleteResponseResponse, error) {
	panic("implement me")
}
