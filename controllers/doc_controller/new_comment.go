package doc_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/utils/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NewCommentParams struct {
	DocumentID string `json:"documentID" form:"documentID" validate:"required"`
	Content    string `json:"content" form:"content" validate:"required"`
	Reply      string `json:"reply" form:"reply"`
}

func NewComment(ctx *gin.Context) {
	res := helper.Res{}

	var params NewCommentParams
	if err := helper.BindWithValid(ctx, &params); err != nil {
		res.Status(http.StatusBadRequest).Error(err).Send(ctx)
		return
	}

	claims, err := middlewares.GetClaims(ctx)
	if err != nil {
		res.Status(http.StatusUnauthorized).Error(err).Send(ctx)
		return
	}

	comment := models.Comment{
		DocumentID: params.DocumentID,
		Content:    params.Content,
		Reply:      params.Reply,
		UserID:     claims.ID,
	}

	if err := comment.New(); err != nil {
		res.Status(http.StatusForbidden).Error(err).Send(ctx)
		return
	}

	res.Success(comment).Send(ctx)
}
