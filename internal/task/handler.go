package task

import (
	"context"

	"github.com/lgu-elo/task/internal/deal/model"
	"github.com/lgu-elo/task/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type (
	taskHandler struct {
		service IService
		log     *logrus.Logger
		server  *grpc.Server
	}

	IHandler interface {
		GetAllTasks(c context.Context, _ *pb.Empty) (*pb.TasksList, error)
		GetTaskById(c context.Context, task *pb.TaskWithID) (*pb.Task, error)
		UpdateTask(c context.Context, task *pb.Task) (*pb.Task, error)
		DeleteTask(c context.Context, task *pb.TaskWithID) (*pb.Empty, error)
		CreateTask(c context.Context, task *pb.Task) (*pb.Empty, error)
	}
)

func NewHandler(service IService, log *logrus.Logger, server *grpc.Server) IHandler {
	return &taskHandler{service, log, server}
}

func (h *taskHandler) GetAllTasks(c context.Context, _ *pb.Empty) (*pb.TasksList, error) {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	var pbtasks pb.TasksList
	for _, task := range tasks {
		pbtasks.Tasks = append(pbtasks.Tasks, &pb.Task{
			Id:          int64(task.ID),
			Description: task.Description,
			Amount:      int64(task.Amount),
			ProjectId:   int64(task.Project_id),
			ClientName:  task.Client_name,
			UserId:      int64(task.User_id),
		})
	}

	return &pbtasks, nil
}

func (h *taskHandler) GetTaskById(c context.Context, request *pb.TaskWithID) (*pb.Task, error) {
	task, err := h.service.GetTaskById(int(request.Id))
	if err != nil {
		return nil, err
	}

	return &pb.Task{
		Id:          int64(task.ID),
		Description: task.Description,
		ProjectId:   int64(task.Project_id),
		Amount:      int64(task.Amount),
		ClientName:  task.Client_name,
		UserId:      int64(task.User_id),
	}, nil
}

func (h *taskHandler) UpdateTask(c context.Context, request *pb.Task) (*pb.Task, error) {
	task, err := h.service.UpdateTask(&model.Task{
		ID:          int(request.Id),
		Description: request.Description,
		Project_id:  int(request.ProjectId),
		Amount:      int(request.Amount),
		Client_name: request.ClientName,
		User_id:     int(request.UserId),
	})
	if err != nil {
		return nil, err
	}

	return &pb.Task{
		Id:          int64(task.ID),
		Description: task.Description,
		ProjectId:   int64(task.Project_id),
		Amount:      int64(task.Amount),
		ClientName:  task.Client_name,
		UserId:      int64(task.User_id),
	}, nil
}

func (h *taskHandler) DeleteTask(c context.Context, task *pb.TaskWithID) (*pb.Empty, error) {
	if err := h.service.DeleteTask(int(task.Id)); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (h *taskHandler) CreateTask(c context.Context, request *pb.Task) (*pb.Empty, error) {
	err := h.service.CreateTask(&model.Task{
		Description: request.Description,
		Project_id:  int(request.ProjectId),
		Amount:      int(request.Amount),
		Client_name: request.ClientName,
		User_id:     int(request.UserId),
	})

	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}
