package handler

import (
	"bioskop/dto"
	bioskopservice "bioskop/service/bioskop_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BioskopHandler struct {
	svc bioskopservice.BioskopService
}

type APIResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
	Data    any    `json:"data"`
}

func NewBioskopHandler(r *gin.RouterGroup, svc bioskopservice.BioskopService) {
	h := &BioskopHandler{svc: svc}
	r.POST("/bioskop", h.Create)
	r.GET("/bioskop", h.Get)
	r.GET("/bioskop/:id", h.GetById)
	r.PUT("/bioskop/:id", h.Update)
	r.DELETE("/bioskop/:id", h.Delete)
}

func (h *BioskopHandler) Create(ctx *gin.Context) {
	var payload dto.CreateRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, APIResponse{
			Message: "Invalid request body",
			Status:  false,
			Data:    err.Error(),
		})
		return
	}

	response, err := h.svc.Create(&payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, APIResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusCreated, APIResponse{
		Message: "Berhasil membuat bioskop",
		Status:  true,
		Data:    response,
	})
}

func (h *BioskopHandler) Get(ctx *gin.Context) {
	response, err := h.svc.Get()
	if err != nil {
		ctx.JSON(http.StatusNotFound, APIResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, APIResponse{
		Message: "Berhasil mendapatkan data bioskop",
		Status:  true,
		Data:    response,
	})
}

func (h *BioskopHandler) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, APIResponse{
			Message: "ID tidak valid",
			Status:  false,
			Data:    nil,
		})
		return
	}

	response, errResponse := h.svc.GetById(uint(id))
	if errResponse != nil {
		ctx.JSON(http.StatusNotFound, APIResponse{
			Message: errResponse.Error(),
			Status:  false,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, APIResponse{
		Message: "Berhasil mendapatkan data bioskop",
		Status:  true,
		Data:    response,
	})
}

func (h *BioskopHandler) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, APIResponse{
			Message: "ID tidak valid",
			Status:  false,
			Data:    nil,
		})
		return
	}

	var payload dto.UpdateRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, APIResponse{
			Message: "Invalid request body",
			Status:  false,
			Data:    err.Error(),
		})
		return
	}

	response, err := h.svc.Update(uint(id), &payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, APIResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, APIResponse{
		Message: "Berhasil memperbarui bioskop",
		Status:  true,
		Data:    response,
	})
}

func (h *BioskopHandler) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, APIResponse{
			Message: "ID tidak valid",
			Status:  false,
			Data:    nil,
		})
		return
	}

	if err := h.svc.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, APIResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, APIResponse{
		Message: "Bioskop berhasil dihapus",
		Status:  true,
		Data:    nil,
	})
}
