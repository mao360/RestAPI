package handler

import "github.com/labstack/echo/v4"

func (h *Handler) CreateList(c echo.Context) error { //возможно в этих методах не нужно возвращать ошибку
	return nil // пока отпраляем nil
}

func (h *Handler) GetAllLists(c echo.Context) error {
	return nil
}

func (h *Handler) GetListById(c echo.Context) error {
	return nil
}

func (h *Handler) UpdateList(c echo.Context) error {
	return nil
}

func (h *Handler) DeleteList(c echo.Context) error {
	return nil
}
