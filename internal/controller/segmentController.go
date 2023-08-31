package controller

import (
	"net/http"
	"strconv"

	"github.com/Angstreminus/avito_intern_backend_2023/internal/model"
	"github.com/Angstreminus/avito_intern_backend_2023/internal/service"
	"github.com/gin-gonic/gin"
)

type SegmentController struct {
	segmentService *service.SegmentService
}

func NewSegmentController(segmentService *service.SegmentService) *SegmentController {
	return &SegmentController{
		segmentService: segmentService,
	}
}

// @Summary CreateSegments
// @Tag Segments
// @Description CreateSegment
// @ID create-segments
// @Accept json
// @Produce json
// @Param input body model.Segments true "segment Info"
// @Success 200 {integer} integer 1
// @Failure 400, 404, 422 {object} model.ResponseError
// @Failure 500 {object} model.ResponseError
// @Router /segments [post]

func (sc SegmentController) CreateSegment(ctx *gin.Context) {
	var segment model.Segments

	err := ctx.BindJSON(segment)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, err)
	}

	resp, respErr := sc.segmentService.CreateSegment(&segment)
	if respErr != nil {
		ctx.AbortWithStatusJSON(respErr.Status, respErr)
		return
	}
	ctx.JSON(http.StatusCreated, resp)
}

// @Summary DeleteSegments
// @Tag Segments
// @Description DeleteSegment
// @ID delete-segments
// @Accept json
// @Produce json
// @Param input body model.Segments true "segment Info"
// @Success 200 {integer} integer 1
// @Failure 400, 404, 422 {object} model.ResponseError
// @Failure 500 {object} model.ResponseError
// @Router /segments [delete]

func (sc SegmentController) DeleteSegment(ctx *gin.Context) {
	queryId := ctx.Param("id")

	segmentId, err := strconv.Atoi(queryId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	respErr := sc.segmentService.DeleteSegment(segmentId)

	if respErr != nil {
		ctx.AbortWithStatusJSON(respErr.Status, respErr)
	}

	ctx.Status(http.StatusNoContent)
}
