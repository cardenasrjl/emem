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
type memServiceServer struct {
	database *mysqldb.DB
}

// NewMemServiceServer ...
func NewMemServiceServer(db *sql.DB) v1.MemServiceServer {
	st := mysqldb.NewDB(db)
	return &memServiceServer{database: st}
}

//validateMem validates if the request to create a mem is valid
func validateMem(m *v1.NewMemRequest) error {
	if m.Title == "" || m.Description == "" {
		return status.Error(codes.Unknown, "missing title or description ")
	}
	return nil
}

// CreateMem
func (s *memServiceServer) NewMem(ctx context.Context, req *v1.NewMemRequest) (*v1.NewMemResponse, error) {
	//validates the entry
	err := validateMem(req)
	if err != nil {
		return nil, err
	}
	//create the mem
	id, err := s.database.CreateMem(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.NewMemResponse{
		Id: id,
	}, nil
}

// GetMem
func (s *memServiceServer) GetMem(ctx context.Context, req *v1.GetMemRequest) (resp *v1.GetMemResponse, err error) {
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

// UpdateMem
func (s *memServiceServer) UpdateMem(ctx context.Context, req *v1.UpdateMemRequest) (*v1.UpdateMemResponse, error) {
	//update the resourse
	err := s.database.UpdateMem(ctx, req)
	if err != nil {
		return nil, err
	}

	//create the response
	return &v1.UpdateMemResponse{}, nil
}

//DeleteMem Deletes a mem entry
func (s *memServiceServer) DeleteMem(ctx context.Context, req *v1.DeleteMemRequest) (*v1.DeleteMemResponse, error) {
	// get SQL connection from pool
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

// GetMems gets all mem
func (s *memServiceServer) GetMems(ctx context.Context, req *v1.GetMemsRequest) (*v1.GetMemsResponse, error) {

	//get data from database
	mems, err := s.database.GetMems(ctx)
	if err != nil {
		return nil, err
	}

	//return the response
	return &v1.GetMemsResponse{
		Mems: mems,
	}, nil
}
