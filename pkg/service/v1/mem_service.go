package v1

import (
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "github.com/cardenasrjl/emem/pkg/api/v1"
	"github.com/cardenasrjl/emem/pkg/storage/mysqldb"
)

// MemServiceServer is implementation of v1.ToDoServiceServer proto interface
type memService struct {
	database *mysqldb.DB
}

// NewMemService ...
func NewMemService(db *sql.DB) v1.MemServiceServer {
	st := mysqldb.NewDB(db)
	return &memService{database: st}
}

//GetVolumeMems
func (s *memService) GetVolumeMems(context.Context, *v1.GetVolumeMemsRequest) (*v1.GetVolumeMemsResponse, error) {
	panic("implement me")
}

//validateMem validates if the request to create a mem is valid
func validateMem(m *v1.NewMemRequest) error {
	if m.Title == "" || m.Description == "" {
		return status.Error(codes.Unknown, "missing title or description ")
	}
	return nil
}

// NewMem
func (s *memService) NewMem(ctx context.Context, req *v1.NewMemRequest) (*v1.NewMemResponse, error) {
	err := validateMem(req)
	if err != nil {
		return nil, err
	}

	id, err := s.database.CreateMem(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.NewMemResponse{
		Id: id,
	}, nil
}

func (s *memService) GetMem(ctx context.Context, req *v1.GetMemRequest) (resp *v1.GetMemResponse, err error) {
	m := &v1.Mem{}

	m, err = s.database.GetMem(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	resp = &v1.GetMemResponse{
		Mem: m,
	}
	return
}

func (s *memService) UpdateMem(ctx context.Context, req *v1.UpdateMemRequest) (*v1.UpdateMemResponse, error) {
	err := s.database.UpdateMem(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateMemResponse{}, nil
}

func (s *memService) DeleteMem(ctx context.Context, req *v1.DeleteMemRequest) (*v1.DeleteMemResponse, error) {
	c, err := s.database.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	err = s.database.DeleteMem(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteMemResponse{}, nil
}

func (s *memService) GetMems(ctx context.Context, req *v1.GetMemsRequest) (*v1.GetMemsResponse, error) {
	mems, err := s.database.GetMems(ctx)
	if err != nil {
		return nil, err
	}

	return &v1.GetMemsResponse{
		Mems: mems,
	}, nil
}
