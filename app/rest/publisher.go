package rest

import (
	"base-gin/app/domain/dto"
	"base-gin/app/service"
	"base-gin/exception"
	"base-gin/server"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PublisherHandler struct {
	hr      *server.Handler
	service *service.PublisherService
}

func newPublisherHandler(hr *server.Handler, publisherService *service.PublisherService,
) *PublisherHandler {
	return &PublisherHandler{hr: hr, service: publisherService}
}

func (h *PublisherHandler) Route(app *gin.Engine) {
	grp := app.Group(server.RootPublisher)
	grp.POST("", h.hr.AuthAccess(), h.create)
	grp.GET("/:id", h.getByID)
	grp.PUT("/:id", h.hr.AuthAccess(), h.update)
	grp.GET("", h.getList)          // Route for getting the list of publishers
	grp.DELETE("/:id", h.hr.AuthAccess(), h.delete)  
}

// getList godoc
//
//	@Summary Get a list of publishers
//	@Description Get a list of publishers.
//	@Produce json
//	@Success 200 {object} dto.SuccessResponse[[]dto.PublisherDetailResp]
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /publishers [get]
func (h *PublisherHandler) getList(c *gin.Context) {
	data, err := h.service.GetList()
	if err != nil {
		h.hr.ErrorInternalServer(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[[]dto.PublisherDetailResp]{
		Success: true,
		Message: "List Publisher",
		Data:    data,
	})
}

// getByID godoc
//
//	@Summary Get a publisher's detail
//	@Description Get a publisher's detail.
//	@Produce json
//	@Param id path int true "Publisher's ID"
//	@Success 200 {object} dto.SuccessResponse[dto.PublisherDetailResp]
//	@Failure 400 {object} dto.ErrorResponse
//	@Failure 404 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /publishers/{id} [get]
func (h *PublisherHandler) getByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("ID tidak valid"))
		return
	}

	data, err := h.service.GetByID(uint(id))
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrUserNotFound):
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse(exception.ErrDataNotFound.Error()))
		default:
			h.hr.ErrorInternalServer(c, err)
		}

		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[dto.PublisherDetailResp]{
		Success: true,
		Message: "Detail Publisher",
		Data:    data,
	})
}

// create godoc
//
//	@Summary Create a new publisher
//	@Description Create a new publisher with name and city details.
//	@Accept json
//	@Produce json
//	@Security BearerAuth
//	@Param data body dto.PublisherCreateReq true "Publisher's details"
//	@Success 201 {object} dto.SuccessResponse[dto.PublisherCreateResp]
//	@Failure 400 {object} dto.ErrorResponse
//	@Failure 401 {object} dto.ErrorResponse
//	@Failure 403 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /publishers [post]
func (h *PublisherHandler) create(c *gin.Context) {
	var req dto.PublisherCreateReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.hr.BindingError(err)
		return
	}

	data, err := h.service.Create(&req)
	if err != nil {
		h.hr.ErrorInternalServer(c, err)
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessResponse[*dto.PublisherCreateResp]{
		Success: true,
		Message: "Data penerbit berhasil disimpan",
		Data:    data,
	})
}

// update godoc
//
//	@Summary Update a publisher's detail
//	@Description Update a publisher's detail.
//	@Accept json
//	@Produce json
//	@Security BearerAuth
//	@Param id path int true "Publisher's ID"
//	@Param detail body dto.PublisherUpdateReq true "Publisher's detail"
//	@Success 200 {object} dto.SuccessResponse[any]
//	@Failure 400 {object} dto.ErrorResponse
//	@Failure 401 {object} dto.ErrorResponse
//	@Failure 403 {object} dto.ErrorResponse
//	@Failure 404 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /publishers/{id} [put]
func (h *PublisherHandler) update(c *gin.Context) {
	// Mendapatkan ID dari parameter URL
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("ID tidak valid"))
		return
	}


	var req dto.PublisherCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.hr.BindingError(err)
		return
	}

	
	data, err := h.service.Update(uint(id), &req)
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrUserNotFound):
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse("Publisher tidak ditemukan"))
		default:
			h.hr.ErrorInternalServer(c, err)
		}
		return
	}


	c.JSON(http.StatusOK, dto.SuccessResponse[*dto.PublisherDetailResp]{
		Success: true,
		Message: "Data penerbit berhasil diperbarui",
		Data:    data,
	})
}

// delete godoc
//
//	@Summary Delete a publisher
//	@Description Delete a publisher by ID.
//	@Security BearerAuth
//	@Param id path int true "Publisher's ID"
//	@Success 200 {object} dto.SuccessResponse[any]
//	@Failure 400 {object} dto.ErrorResponse
//	@Failure 401 {object} dto.ErrorResponse
//	@Failure 403 {object} dto.ErrorResponse
//	@Failure 404 {object} dto.ErrorResponse
//	@Failure 500 {object} dto.ErrorResponse
//	@Router /publishers/{id} [delete]
func (h *PublisherHandler) delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, h.hr.ErrorResponse("ID tidak valid"))
		return
	}

	err = h.service.Delete(uint(id))
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrUserNotFound):
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse("Publisher tidak ditemukan"))
		default:
			h.hr.ErrorInternalServer(c, err)
		}
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[any]{
		Success: true,
		Message: "Data penerbit berhasil dihapus",
	})
}
