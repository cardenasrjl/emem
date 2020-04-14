package v1

import (
	"context"
	"database/sql"

	redis2 "github.com/go-redis/redis/v7"

	"github.com/cardenasrjl/emem/pkg/storage/redis"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "github.com/cardenasrjl/emem/pkg/api/v1"
	"github.com/cardenasrjl/emem/pkg/storage/mysqldb"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// toDoServiceServer is implementation of v1.ToDoServiceServer proto interface
type toDoServiceServer struct {
	database *mysqldb.DB
	redis    *redis.RdsClient
}

// NewToDoServiceServer creates ToDo service
func NewToDoServiceServer(db *sql.DB, rd *redis2.Client) v1.ToDoServiceServer {
	st := mysqldb.NewDB(db)
	rdsClient := redis.NewClient(rd)
	return &toDoServiceServer{database: st, redis: rdsClient}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *toDoServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// Create new todo task
func (s *toDoServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	tdObj := req.ToDo
	if tdObj == nil {
		return nil, status.Errorf(codes.Unimplemented,
			"request nil ", apiVersion)
	}

	id, err := s.database.CreateTodo(ctx, tdObj)
	if err != nil {
		return nil, err
	}

	return &v1.CreateResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

// Read todo task
func (s *toDoServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (resp *v1.ReadResponse, err error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	td := &v1.ToDo{}
	found := s.redis.Get(s.redis.Key(req.Id), td)
	if !found {

		td, err = s.database.ReadTodo(ctx, req.Id)
		if err != nil {
			return
		}

		//save for next time
		err = s.redis.Set(s.redis.Key(req.Id), td)
		if err != nil {
			return
		}

	}
	//err = s.redis.GetSet(fmt.Sprintf("todo:%d", req.Id), &td, func() (interface{}, error) {
	//	return s.database.ReadTodo(ctx, req.Id)
	//})
	//
	//if err != nil {
	//	return
	//}
	//create the response
	resp = &v1.ReadResponse{
		Api:  apiVersion,
		ToDo: td,
	}

	return

}

// Update todo task
func (s *toDoServiceServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	//update the resourse
	rows, err := s.database.UpdateTodo(ctx, req.GetToDo())
	if err != nil {
		return nil, err
	}

	//remove from cache
	err = s.redis.Del(s.redis.Key(req.ToDo.Id))
	if err != nil {
		return nil, err
	}

	//create the response
	return &v1.UpdateResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

// Delete todo task
func (s *toDoServiceServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	// get SQL connection from pool
	c, err := s.database.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	rows, err := s.database.DeleteTodo(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	//remove from cache
	err = s.redis.Del(s.redis.Key(req.Id))
	if err != nil {
		return nil, err
	}

	return &v1.DeleteResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}

// Read all todo tasks
func (s *toDoServiceServer) ReadAll(ctx context.Context, req *v1.ReadAllRequest) (*v1.ReadAllResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	//get data from database
	list, err := s.database.ReadAllTodo(ctx)
	if err != nil {
		return nil, err
	}

	//return the response
	return &v1.ReadAllResponse{
		Api:   apiVersion,
		ToDos: list,
	}, nil
}
