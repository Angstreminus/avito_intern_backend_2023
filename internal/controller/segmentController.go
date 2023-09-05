package controller

import (
	"net/http"
	"strconv"

	apperrors "github.com/Angstreminus/avito_intern_backend_2023/internal/AppErrors"
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
		respErr := apperrors.ResponseError{
			Message: err.Error(),
			Status:  http.StatusUnprocessableEntity,
		}
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, respErr)
	}

	resp, err := sc.segmentService.CreateSegment(&segment)
	if err != nil {
		respErr := apperrors.MatchError(err)
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
		respErr := apperrors.ResponseError{
			Message: err.Error(),
			Status:  http.StatusUnprocessableEntity,
		}
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, respErr)
	}

	err = sc.segmentService.DeleteSegment(segmentId)

	if err != nil {
		respErr := apperrors.MatchError(err)
		ctx.AbortWithStatusJSON(respErr.Status, respErr)
	}

	ctx.Status(http.StatusNoContent)
}
