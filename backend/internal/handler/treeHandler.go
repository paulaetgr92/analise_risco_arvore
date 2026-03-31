package handler

import (
	"arvore/internal/model"
	"arvore/internal/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TreeHandler struct {
	service *service.TreeService
}

func NewTreeHandler(s *service.TreeService) *TreeHandler {
	return &TreeHandler{
		service: s,
	}
}

// ✅ CREATE TREE
func (h *TreeHandler) Create(c echo.Context) error {
	var tree model.Tree

	// Bind do JSON
	if err := c.Bind(&tree); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"erro": "dados inválidos",
		})
	}

	// ✅ CORREÇÃO AQUI: passar contexto + variável correta
	result, err := h.service.CreateTreeService(c.Request().Context(), tree)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"erro": "erro ao salvar",
		})
	}

	return c.JSON(http.StatusCreated, result)
}

// ✅ LISTAR TODAS
func (h *TreeHandler) ListTree(c echo.Context) error {
	trees, err := h.service.ListTreesService(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"erro": "erro ao listar",
		})
	}

	return c.JSON(http.StatusOK, trees)
}
func (h *TreeHandler) ListTreesByBoundingBox(c echo.Context) error {
	lat1, _ := strconv.ParseFloat(c.QueryParam("lat1"), 64)
	lat2, _ := strconv.ParseFloat(c.QueryParam("lat2"), 64)
	lng1, _ := strconv.ParseFloat(c.QueryParam("lng1"), 64)
	lng2, _ := strconv.ParseFloat(c.QueryParam("lng2"), 64)

	result, err := h.service.LisTreesByBoundingBoxService(
		c.Request().Context(),
		model.ListTreesByBoundingBoxRequest{
			Latitude:    lat1,
			Latitude_2:  lat2,
			Longitude:   lng1,
			Longitude_2: lng2,
		},
	)

	if err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(200, result)
}
func (h *TreeHandler) GetTreeByID(c echo.Context) error {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "id inválido",
		})
	}

	tree, err := h.service.GetTreeByIdService(
		c.Request().Context(),
		id,
	)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, tree)
}

func (h *TreeHandler) UpdateTree(c echo.Context) error {
	var req model.UpdateTreeRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "dados inválidos",
		})
	}

	if req.ID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "id é obrigatório",
		})
	}

	tree, err := h.service.UpdateTreeService(
		c.Request().Context(),
		req,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, tree)
}

func (h *TreeHandler) DeleteTree(c echo.Context) error {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "id inválido",
		})
	}

	err = h.service.DeleteTreeService(
		c.Request().Context(),
		id,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "árvore removida com sucesso",
	})
}

func (h *TreeHandler) ListPotentialRiskTree(c echo.Context) error {
	tree, err := h.service.ListPotencialRiskTree(
		c.Request().Context(),
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, tree)
}
