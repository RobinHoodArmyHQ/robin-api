package auth

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/RobinHoodArmyHQ/robin-api/internal/env"
	u "github.com/RobinHoodArmyHQ/robin-api/internal/repositories/user"
	"github.com/RobinHoodArmyHQ/robin-api/internal/util"
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AuthHandler(c *gin.Context) {

	// validate country isd code
	countryCode, err := strconv.ParseUint(c.PostForm("country_code"), 10, 8)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.AuthResponse{
			Status: models.StatusFailed(fmt.Sprintf("invalid country code %d", countryCode)),
		})
		return
	}

	// validate mobile number
	mobileNumber, err := strconv.ParseUint(c.PostForm("mobile_number"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.AuthResponse{
			Status: models.StatusFailed(fmt.Sprintf("invalid mobile number %d", mobileNumber)),
		})
		return
	}

	// generate request id and send response
	requestId := uuid.Must(uuid.NewRandom())
	c.JSON(http.StatusOK, models.AuthResponse{
		Status:    models.StatusSuccess(),
		RequestID: requestId,
	})
}

func RegisterUser(c *gin.Context) {
	var request u.RegisterUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, u.RegisterUserResponse{
			Status: *models.StatusFailed(fmt.Sprintln("Invalid inputs")),
		})
		return
	}

	registerUser := &u.CheckIfUserExistsRequest{
		EmailId: request.EmailId,
	}

	userRepo := env.FromContext(c).UserRepository
	user, err := userRepo.CheckIfUserExists(registerUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, u.RegisterUserResponse{
			Status: *models.StatusSomethingWentWrong(),
		})
		return
	}

	if user.IsExisting {
		c.JSON(http.StatusConflict, u.RegisterUserResponse{
			Status:    *models.StatusSuccess(fmt.Sprintf("User already exist with the email: %s", request.EmailId)),
			IsNewUser: user.IsExisting,
		})
		return
	}

	hashedPassword, err := util.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, u.RegisterUserResponse{
			Status: *models.StatusSomethingWentWrong(),
		})
		return
	}

	firstName, lastName := util.GetFirstNameAndLastName(request.FullName)
	newUserData := &models.User{
		FirstName:    firstName,
		LastName:     lastName,
		EmailId:      request.EmailId,
		PasswordHash: hashedPassword,
	}

	newUser, err := userRepo.CreateUser(&u.CreateUserRequest{
		User: newUserData,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, u.RegisterUserResponse{
			Status: *models.StatusSomethingWentWrong(),
		})
		return
	}

	util.SendEmailVerificationCode(newUser)

	c.JSON(http.StatusCreated, u.RegisterUserResponse{
		IsNewUser: !user.IsExisting,
		Status:    *models.StatusSuccess(fmt.Sprintln("User created successfully")),
	})
}

func LoginUser(c *gin.Context) {
	var request u.LoginUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, u.LoginUserResponse{
			Status: *models.StatusFailed(fmt.Sprintln("Invalid creadentitials")),
		})
		return
	}

	loginUser := &u.GetUserByEmailIdRequest{
		EmailId: request.EmailId,
	}
	userRepo := env.FromContext(c).UserRepository

	user, err := userRepo.GetUserByEmailId(loginUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, u.LoginUserResponse{
			Status: *models.StatusSomethingWentWrong(),
		})
		return
	}

	if err == nil && user == nil {
		c.JSON(http.StatusNoContent, u.LoginUserResponse{
			Status: *models.StatusFailed(fmt.Sprintf("No user found with given email_id: %s", request.EmailId)),
		})
		return
	}

	// verify user password
	ok := util.CheckPasswordHash(request.Password, user.User.PasswordHash)

	if !ok {
		c.JSON(http.StatusBadRequest, u.LoginUserResponse{
			Status: *models.StatusFailed(fmt.Sprintln("Incorrect email or password")),
		})
		return
	}

	//	now create a JWT token and return the token in request
	jwtToken, err := util.GenerateJwt(request.EmailId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, u.LoginUserResponse{
			Status: *models.StatusSomethingWentWrong(),
		})
		return
	}

	c.JSON(http.StatusOK, u.LoginUserResponse{
		Status: *models.StatusSuccess("User logged-in successfully"),
		Token:  jwtToken,
	})
}
