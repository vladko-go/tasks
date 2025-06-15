package handlers

import (
	"context"
	"fmt"
	"pet-project/internal/tasksService"
	"pet-project/internal/web/tasks"
)

type Handler struct {
	Service tasksService.TaskService
}

func NewHandler(s tasksService.TaskService) *Handler {
	return &Handler{Service: s}
}

func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	id := request.Id
	task, err := h.Service.GetByID(id)
	if err != nil {
		return nil, err
	}
	task.IsDone = !task.IsDone
	editedTa, err := h.Service.Update(id, tasksService.TaskUpdateRequest{IsDone: task.IsDone})
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &editedTa.ID,
		Task:   &editedTa.Task,
		IsDone: &editedTa.IsDone,
	}
	return response, nil
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получение всех задач из сервиса
	allTasks, err := h.Service.GetTasks()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := tasks.GetTasks200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	fmt.Println(request.Body)
	// Обращаемся к сервису и создаем задачу
	taskToCreate := tasksService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.Create(taskToCreate)

	if err != nil {
		return nil, err
	}

	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *Handler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	id := request.Id
	// Обращаемся к сервису и удаляем задачу
	err := h.Service.Delete(id)
	if err != nil {
		return nil, err
	}

	response := tasks.DeleteTasksId204Response{}

	return response, nil
}

func (h *Handler) GetTasksId(_ context.Context, request tasks.GetTasksIdRequestObject) (tasks.GetTasksIdResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	id := request.Id
	// Обращаемся к сервису и получаем задачу
	task, err := h.Service.GetByID(id)
	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.GetTasksId200JSONResponse{
		Id:     &task.ID,
		Task:   &task.Task,
		IsDone: &task.IsDone,
	}
	// Просто возвращаем респонс!
	return response, nil
}
