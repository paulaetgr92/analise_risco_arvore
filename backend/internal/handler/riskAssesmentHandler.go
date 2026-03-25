package handler

import (
	"net/http"
	"strconv"

	"arvore/internal/service"

	"github.com/labstack/echo/v4"
)

type RiskAssessmentHandler struct {
	service service.RiskAssesmentServiceInterface
}

func NewRiskAssessmentHandler(
	service service.RiskAssesmentServiceInterface,
) *RiskAssessmentHandler {
	return &RiskAssessmentHandler{service: service}
}

func (h *RiskAssessmentHandler) CreateRiskAssessment(c echo.Context) error {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "id inválido",
		})
	}

	risk, err := h.service.GetRiskByID(
		c.Request().Context(),
		id,
	)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, risk)
}

func (h *RiskAssessmentHandler) ListRiskByTree(c echo.Context) error {
	treeIDParam := c.Param("treeId")

	treeID, err := strconv.ParseInt(treeIDParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "tree_id inválido",
		})
	}

	list, err := h.service.ListRiskByTree(
		c.Request().Context(),
		treeID,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, list)
}

func (h *RiskAssessmentHandler) GetLatestRiskByTree(c echo.Context) error {
	treeIDParam := c.Param("treeId")

	treeID, err := strconv.ParseInt(treeIDParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "tree_id inválido",
		})
	}

	risk, err := h.service.GetLatestRiskByTree(
		c.Request().Context(),
		treeID,
	)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, risk)
}
