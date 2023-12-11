package service

import (
	"context"
	"fmt"

	//model "github.com/ArquitecturaDeSistemas/taskmicroservice/dominio"
	model "github.com/ArquitecturaDeSistemas/taskmicroservice/src/dominio"
	repository "github.com/ArquitecturaDeSistemas/taskmicroservice/src/ports"
	pb "github.com/ArquitecturaDeSistemas/taskmicroservice/src/proto"
)

// este servicio implementa la interfaz UserServiceServer
// que se genera a partir del archivo proto
type TaskService struct {
	pb.UnimplementedTaskServiceServer
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {

	crearTareaInput := model.CrearTareaInput{
		Titulo:       req.GetTitulo(),
		Descripcion:  req.GetDescripcion(),
		FechaInicio:  req.GetFechaInicio(),
		FechaTermino: req.GetFechaTermino(),
		UserId:       req.GetUserId(),
	}
	u, err := s.repo.CrearTarea(crearTareaInput)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Tarea creado: %v", u)
	response := &pb.CreateTaskResponse{
		Id:           u.ID,
		Titulo:       u.Titulo,
		Descripcion:  u.Descripcion,
		FechaInicio:  u.FechaInicio,
		FechaTermino: u.FechaTermino,
		UserId:       u.UserId,
	}
	fmt.Printf("Tarea creado: %v", response)
	return response, nil
}

func (s *TaskService) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.GetTaskResponse, error) {
	u, err := s.repo.Tarea(req.GetId())
	if err != nil {
		return nil, err
	}
	response := &pb.GetTaskResponse{
		Id:           u.ID,
		Titulo:       u.Titulo,
		Descripcion:  u.Descripcion,
		FechaInicio:  u.FechaInicio,
		FechaTermino: u.FechaTermino,
		UserId:       u.UserId,
	}
	return response, nil
}

func (s *TaskService) ListTasks(ctx context.Context, req *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	tasks, err := s.repo.Tareas()
	if err != nil {
		return nil, err
	}
	var response []*pb.Task
	for _, u := range tasks {
		task := &pb.Task{
			Id:           u.ID,
			Titulo:       u.Titulo,
			Descripcion:  u.Descripcion,
			FechaInicio:  u.FechaInicio,
			FechaTermino: u.FechaTermino,
			UserId:       u.UserId,
		}
		response = append(response, task)
	}

	return &pb.ListTasksResponse{Tasks: response}, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	titulo := req.GetTitulo()
	descripcion := req.GetDescripcion()
	fechaInicio := req.GetFechaInicio()
	fechaTermino := req.GetFechaTermino()
	userId := req.GetUserId()
	fmt.Printf("Nombre: %v", titulo)
	actualizarTareaInput := &model.ActualizarTareaInput{
		Titulo:       &titulo,
		Descripcion:  &descripcion,
		FechaInicio:  &fechaInicio,
		FechaTermino: &fechaTermino,
		UserId:       &userId,
	}
	fmt.Printf("Tarea actualizada input: %v", actualizarTareaInput)
	u, err := s.repo.ActualizarTarea(req.GetId(), actualizarTareaInput)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Tarea actualizado: %v", u)
	response := &pb.UpdateTaskResponse{
		Id:           u.ID,
		Titulo:       u.Titulo,
		Descripcion:  u.Descripcion,
		FechaInicio:  u.FechaInicio,
		FechaTermino: u.FechaTermino,
		UserId:       u.UserId,
	}
	return response, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	respuesta, err := s.repo.EliminarTarea(req.GetId())
	if err != nil {
		return nil, err
	}
	response := &pb.DeleteTaskResponse{
		Mensaje: respuesta.Mensaje,
	}
	return response, nil
}
