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

type UserSegmentController struct {
	UserSegmentService *service.UserSegmentService
}

func NewUserSegmentController(usrSegmService *service.UserSegmentService) *UserSegmentController {
	return &UserSegmentController{
		UserSegmentService: usrSegmService,
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

func (uc UserSegmentController) CreateUserSegment(ctx *gin.Context) {
	var userSegment model.UserSegments

	err := ctx.BindJSON(userSegment)
	if err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
	}

	resp, respErr := uc.UserSegmentService.CreateUserSegment(&userSegment)
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

func (uc UserSegmentController) EditUserSegment(ctx *gin.Context) {
	var req EditUserSegmentRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
	}

	respErr := uc.UserSegmentService.EditUserSegment(req.AddSegments, req.RemoveSegments, req.UserId)
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

func (uc UserSegmentController) GetSegmentsNamesByUserID(ctx *gin.Context) {
	stringParamId := ctx.Param("id")
	userId, err := strconv.Atoi(stringParamId)
	if err != nil {
		ctx.AbortWithError(http.StatusUnprocessableEntity, err)
	}

	userActiveSegments, respErr := uc.UserSegmentService.GetSegmentNamesByUserId(userId)
	if respErr != nil {
		ctx.AbortWithStatusJSON(respErr.Status, respErr)
		return
	}

	ctx.JSON(http.StatusOK, userActiveSegments)
}
