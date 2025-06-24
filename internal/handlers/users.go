package handlers

import (
	"context"

	"pet-project/internal/usersService"
	"pet-project/internal/web/users"
	"time"
)

type UserHandler struct {
	Service usersService.UserService
}

func NewUserHandler(s usersService.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	id := request.Id
	updatedUser := request.Body
	user, err := h.Service.GetByID(id)
	if err != nil {
		return nil, err
	}

	if updatedUser.Email != "" {
		user.Email = updatedUser.Email
	}

	if updatedUser.Password != "" {
		user.Password = updatedUser.Password
	}

	user.UpdatedAt = time.Now()
	editedTa, err := h.Service.Update(id, usersService.UserUpdateRequest{Email: user.Email, Password: user.Password, UpdatedAt: user.UpdatedAt})
	if err != nil {
		return nil, err
	}

	response := users.PatchUsersId200JSONResponse{
		Id:    &editedTa.ID,
		Email: &editedTa.Email,
	}
	return response, nil
}

func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	// Получение всех задач из сервиса
	allUsers, err := h.Service.GetUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, tsk := range allUsers {
		user := users.User{
			Id:    &tsk.ID,
			Email: &tsk.Email,
		}
		response = append(response, user)
	}

	return response, nil
}

func (h *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	userRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	userToCreate := usersService.UserCreateRequest{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	createdUser, err := h.Service.Create(userToCreate)

	if err != nil {
		return nil, err
	}

	// создаем структуру респонс
	response := users.PostUsers201JSONResponse{
		Id:    &createdUser.ID,
		Email: &createdUser.Email,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *UserHandler) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	id := request.Id
	// Обращаемся к сервису и удаляем задачу
	err := h.Service.Delete(id)
	if err != nil {
		return nil, err
	}

	response := users.DeleteUsersId204Response{}

	return response, nil
}

func (h *UserHandler) GetUsersId(_ context.Context, request users.GetUsersIdRequestObject) (users.GetUsersIdResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	id := request.Id
	// Обращаемся к сервису и получаем задачу
	task, err := h.Service.GetByID(id)
	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := users.GetUsersId200JSONResponse{
		Id:    &task.ID,
		Email: &task.Email,
	}
	// Просто возвращаем респонс!
	return response, nil
}
