package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/username/GitRepoName/internal/models"
)

type UserController struct {
	UserRepository models.UserRepository
}

func (sc *UserController) GetProfile(c *gin.Context) {
	userID := c.GetUint("userID")

	profile, err := sc.UserRepository.GetProfile(c, int(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_GET_USER_PROFILE",
					Message: "Can't get profile from db",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Result: profile})
}
