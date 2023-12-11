package service

import (
	"context"
	"fmt"

	model "github.com/ArquitecturaDeSistemas/taskmicroservice/dominio"
	repository "github.com/ArquitecturaDeSistemas/taskmicroservice/ports"
	pb "github.com/ArquitecturaDeSistemas/taskmicroservice/proto"
)

// este servicio implementa la interfaz UserServiceServer
// que se genera a partir del archivo proto
type TaskService struct {
	pb.UnimplementedUserServiceServer
	repo repository.UserRepository
}

func NewTaskService(repo repository.UserRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {

	crearTareaInput := model.CrearTareaInput{
		titulo:       req.GetTitulo(),
		descripcion:  req.GetDescripcion(),
		fechaInicio:  req.GetfechaInicio(),
		fechaTermino: req.GetfechaTermino(),
		userId:       req.GetuserId(),
	}
	u, err := s.repo.CrearUsuario(crearTareaInput)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Tarea creado: %v", u)
	response := &pb.CreateTaskResponse{
		Id:           u.ID,
		titulo:       u.Titulo,
		descripcion:  u.Descripcion,
		fechaInicio:  u.fechaInicio,
		fechaTermino: u.fechaTermino,
		userId:       u.userId,
	}
	fmt.Printf("Tarea creado: %v", response)
	return response, nil
}

func (s *TaskService) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.GetTaskResponse, error) {
	u, err := s.repo.Usuario(req.GetId())
	if err != nil {
		return nil, err
	}
	response := &pb.GetTaskResponse{
		Id:           u.ID,
		titulo:       u.Titulo,
		descripcion:  u.Descripcion,
		fechaInicio:  u.fechaInicio,
		fechaTermino: u.fechaTermino,
		userId:       u.userId,
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
			titulo:       u.Titulo,
			descripcion:  u.Descripcion,
			fechaInicio:  u.fechaInicio,
			fechaTermino: u.fechaTermino,
			userId:       u.userId,
		}
		response = append(response, task)
	}

	return &pb.ListTasksResponse{Tasks: response}, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	titulo := req.GetTitulo()
	descripcion := req.GetDescripcion()
	fechaInicio := req.GetfechaInicio()
	fechaTermino := req.GetfechaTermino()
	userId := req.GetuserId()
	fmt.Printf("Nombre: %v", titulo)
	actualizarTareaInput := &model.ActualizarTareaInput{
		Titulo:       &titulo,
		descripcion:  &descripcion,
		fechaInicio:  &fechaInicio,
		fechaTermino: &fechaTermino,
		userId:       &userId,
	}
	fmt.Printf("Tarea actualizada input: %v", actualizarTareaInput)
	u, err := s.repo.ActualizarTarea(req.GetId(), actualizarTareaInput)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Tarea actualizado: %v", u)
	response := &pb.UpdateTaskResponse{
		Id:           u.ID,
		titulo:       u.Titulo,
		descripcion:  u.Descripcion,
		fechaInicio:  u.fechaInicio,
		fechaTermino: u.fechaTermino,
		userId:       u.userId,
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
