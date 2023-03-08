package api

import (
	"github.com/go-playground/validator/v10"
	"gofiber-idp/server/config"
	"gofiber-idp/server/dto"
	"gofiber-idp/server/errors"
	"gofiber-idp/server/logger"
	"gofiber-idp/server/service"
	validate2 "gofiber-idp/server/util/validate"
	"gorm.io/gorm"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserApi struct {
	service *service.UserService
}

func InitUserApi(app *fiber.App, db *gorm.DB) {
	_userApi = &UserApi{
		service: service.NewUserService(db),
	}

	routes := app.Group(config.UserPath)
	routes.Get("/:id", _userApi.getUser)
	routes.Post("", _userApi.createUser)
	routes.Post("/:id", _userApi.updateUser)
	routes.Delete("/:id", _userApi.deleteUser)
}

func (c *UserApi) getUser(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.LogError(err.Error())
		return ctx.JSON(errors.ErrorResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid Input.",
		})
	}

	user, err := c.service.GetUser(id)
	if err != nil {
		logger.LogError(err.Error())
		return ctx.JSON(errors.ErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    "User Not Found.",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&user)
}
func (c *UserApi) createUser(ctx *fiber.Ctx) error {
	body := dto.CreateUserRequest{}

	if err := ctx.BodyParser(&body); err != nil {
		logger.LogError(err.Error())
		return ctx.JSON(errors.ErrorResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid User Create Request Input.",
		})
	}

	validate := validator.New()
	validate.RegisterValidation("phone", validate2.PhoneValidation)

	//validate error
	if err := validate.Struct(body); err != nil {
		logger.LogError(err.Error())
		return ctx.JSON(errors.ErrorResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid User Create Request Input.",
		})
	}

	user, err := c.service.CreateUser(body)
	if err != nil {
		logger.LogError(err.Error())
		return ctx.JSON(errors.ErrorResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "User Create Failed",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&user)
}

func (c *UserApi) updateUser(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.LogError(err.Error())
		return ctx.JSON(errors.ErrorResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid Input.",
		})
	}

	body := dto.UpdateUserRequest{}

	if err := ctx.BodyParser(&body); err != nil {
		logger.LogError(err.Error())
		return ctx.JSON(errors.ErrorResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid User Update Request Input.",
		})
	}

	validate := validator.New()
	validate.RegisterValidation("phone", validate2.PhoneValidation)

	//validate error
	if err := validate.Struct(body); err != nil {
		logger.LogError(err.Error())
		return ctx.JSON(errors.ErrorResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid User Update Request Input.",
		})
	}

	user, err := c.service.UpdateUser(id, body)
	if err != nil {
		logger.LogError(err.Error())
		return ctx.JSON(errors.ErrorResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "User Update Failed",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&user)
}

func (c *UserApi) deleteUser(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.LogError(err.Error())
		return ctx.JSON(errors.ErrorResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid Input.",
		})
	}

	err = c.service.DeleteUser(id)
	if err != nil {
		logger.LogError(err.Error())
		return ctx.JSON(errors.ErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    "User Not Found.",
		})
	}

	return ctx.Status(fiber.StatusOK).SendString("")
}

//func (c *UserApi) join(ctx *fiber.Ctx) error {
//	req := config.JoinRequest{}
//
//	if err := ctx.BodyParser(&req); err != nil {
//		logger.LogError(err.Error())
//		return ctx.JSON(errors.ErrorResponse{
//			ResultCode: config.FailBodyParser,
//		})
//	}
//
//	lastInsertId, err := c.service.CreateAccount(req.Nickname)
//	if err != nil {
//		logger.LogError(err.Error())
//		return ctx.JSON(errors.ErrorResponse{
//			ResultCode: config.DBFail,
//		})
//	}
//
//	return ctx.JSON(errors.ErrorResponse{
//		Response: config.JoinReponse{
//			Id: lastInsertId,
//		},
//	})
//}

var _userApi *UserApi
