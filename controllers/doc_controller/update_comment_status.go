package doc_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateCommentStatusParams struct {
	CommentID string `validate:"required" form:"commentID" json:"commentID"`
	Operator  int    `form:"operator" json:"operator"`
}

func UpdateCommentStatus(ctx *gin.Context) {
	var res helper.Res
	var params UpdateCommentStatusParams
	if err := helper.BindWithValid(ctx, &params); err != nil {
		res.Status(http.StatusBadRequest).Error(ctx, err)
		return
	}

	claims, err := middlewares.GetClaims(ctx)
	if err != nil {
		res.Status(http.StatusUnauthorized).Error(ctx, err)
		return
	}

	cs := models.CommentStatus{
		CommentID: params.CommentID,
		UserID:    claims.ID,
	}

	// 无赞踩
	if params.Operator == models.NONE {
		if err := cs.Delete(); err != nil {
			res.Status(http.StatusForbidden).Error(ctx, err)
			return
		}

		res.Send(ctx, cs)
		return
	}

	cs.IsUp = params.Operator == models.UP

	if err := cs.New(); err != nil {
		res.Status(http.StatusForbidden).Error(ctx, err)
		return
	}

	res.Send(ctx, cs)
}
