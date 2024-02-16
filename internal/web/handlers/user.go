package handlers

import (
	"github.com/ghuyng/gohtmx/internal/web/models"
	"github.com/ghuyng/gohtmx/internal/web/view/user"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type UserHandler struct {
	logger *zap.SugaredLogger
}

func NewUserHandler(logger *zap.SugaredLogger) *UserHandler {
	return &UserHandler{
		logger: logger,
	}
}

func (h *UserHandler) GetUsers(ctx echo.Context) error {
	u := models.User{
		Email: "a@gg.com",
	}
	h.logger.Info("get users")
	return render(ctx, user.Show(u))
}
