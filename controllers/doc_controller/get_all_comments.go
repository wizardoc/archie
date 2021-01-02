package doc_controller

import (
	"archie/middlewares"
	"archie/models"
	"archie/utils"
	"archie/utils/helper"
	"archie/utils/jwt_utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetAllCommentsParams struct {
	utils.PageInfo
}

type ResComments struct {
	models.Comment
	User models.User `json:"user"`
}

func GetAllComments(ctx *gin.Context) {
	var res helper.Res
	var params GetAllCommentsParams

	if err := helper.BindWithValid(ctx, &params); err != nil {
		res.Status(http.StatusBadRequest).Error(err).Send(ctx)
		return
	}

	// parse JWT, the JWT may does not exist
	token, hasToken := middlewares.GetJWTFromHeader(ctx.Request)
	var claims jwt_utils.LoginClaims

	if hasToken {
		// invalidate token
		if err := middlewares.ParseToken2Claims(token, &claims); err != nil {
			res.Status(http.StatusUnauthorized).Error(err).Send(ctx)
			return
		}
	}

	params.ParsePageInfo()

	docID := ctx.Params.ByName("document_id")
	comment := models.Comment{
		DocumentID: docID,
	}

	var comments []models.Comment
	if err := comment.FindAll(params.Page, params.PageSize, &comments); err != nil {
		res.Status(http.StatusForbidden).Error(err).Send(ctx)
		return
	}

	// fill up and down
	for i, comment := range comments {
		downAndUpSum := len(comment.CommentStatus)
		var upStatuses []models.CommentStatus

		utils.ArrayFilter(comment.CommentStatus, func(item interface{}) bool {
			return item.(models.CommentStatus).IsUp
		}, &upStatuses)

		upStatusesCount := len(upStatuses)

		comment.Up = upStatusesCount
		comment.Down = downAndUpSum - upStatusesCount
		comment.Status = models.NONE

		if hasToken {
			var status models.CommentStatus

			if ok := utils.ArrayFind(comment.CommentStatus, func(item interface{}) bool {
				return item.(models.CommentStatus).UserID == claims.ID
			}, &status); ok {
				if status.IsUp {
					comment.Status = models.UP
				} else {
					comment.Status = models.DOWN
				}
			}
		}

		comments[i] = comment
	}

	res.Success(comment).Send(ctx)
}
