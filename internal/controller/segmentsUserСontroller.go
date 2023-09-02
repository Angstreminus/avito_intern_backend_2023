package controller

import (
	"net/http"
	"strconv"

	"github.com/Angstreminus/avito_intern_backend_2023/internal/model"
	"github.com/Angstreminus/avito_intern_backend_2023/internal/service"
	"github.com/gin-gonic/gin"
)

type EditUserSegmentRequest struct {
	UserId         int      `json:"user_id"`
	AddSegments    []string `json:"add_segments,omitempty"`
	RemoveSegments []string `json:"remove_segments,omitempty"`
}

type SegmentsUserController struct {
	UserSegmentService *service.SegmentsUserService
}

func NewUserSegmentController(segmUsrService *service.SegmentsUserService) *SegmentsUserController {
	return &SegmentsUserController{
		UserSegmentService: segmUsrService,
	}
}

// @Summary CreateUserSegments
// @Tag UserSegments
// @Description CreateUserSegment
// @ID create-user-segments
// @Accept json
// @Produce json
// @Param input body model.Segments true "segment Info"
// @Success 200 {integer} integer 1
// @Failure 400, 404, 422 {object} model.ResponseError
// @Failure 500 {object} model.ResponseError
// @Router /user_segments [post]

func (sc SegmentsUserController) CreateSegmentsUsers(ctx *gin.Context) {
	var segmentsUsers model.SegmentsUsers

	err := ctx.BindJSON(segmentsUsers)
	if err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
	}

	resp, respErr := sc.UserSegmentService.CreateSegmentUser(&segmentsUsers)
	if respErr != nil {
		ctx.AbortWithStatusJSON(respErr.Status, respErr)
		return
	}
	ctx.JSON(http.StatusCreated, resp)
}

// @Summary EditUserSegments
// @Tag UserSegments
// @Description EditUserSegment
// @ID edit-user-segments
// @Accept json
// @Produce json
// @Param input body model.EditUserSegmentRequest true "edit segment Info"
// @Success 200 {integer} integer 1
// @Failure 400, 404, 422 {object} model.ResponseError
// @Failure 500 {object} model.ResponseError
// @Router /user_segments [put]

func (sc SegmentsUserController) EditUserSegment(ctx *gin.Context) {
	var req EditUserSegmentRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
	}

	respErr := sc.UserSegmentService.EditSegment(req.AddSegments, req.RemoveSegments, req.UserId)
	if respErr != nil {
		ctx.AbortWithStatusJSON(respErr.Status, respErr)
		return
	}
	ctx.Status(http.StatusNoContent)
}

// @Summary GetUserSegments
// @Tag UserSegments
// @Description GetUserSegment
// @ID get-user-segments
// @Accept json
// @Produce json
// @Success 200 {integer} integer 1
// @Failure 400, 404, 422 {object} model.ResponseError
// @Failure 500 {object} model.ResponseError
// @Router /user_segments [get]

func (sc SegmentsUserController) GetSegmentsNamesByUserID(ctx *gin.Context) {
	stringParamId := ctx.Param("id")
	userId, err := strconv.Atoi(stringParamId)
	if err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
	}

	userActiveSegments, respErr := sc.UserSegmentService.GetSegmentNamesByUserId(userId)
	if respErr != nil {
		ctx.AbortWithStatusJSON(respErr.Status, respErr)
		return
	}

	ctx.JSON(http.StatusOK, userActiveSegments)
}
