package auth

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/RobinHoodArmyHQ/robin-api/internal/env"
	userrepo "github.com/RobinHoodArmyHQ/robin-api/internal/repositories/user"
	"github.com/RobinHoodArmyHQ/robin-api/internal/util"
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/nanoid"
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
	var request RegisterUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, RegisterUserResponse{
			Status: models.StatusFailed(fmt.Sprintln("Invalid inputs")),
		})
		return
	}

	registerUser := &userrepo.GetUserByEmailRequest{
		EmailId: request.EmailId,
	}

	userRepo := env.FromContext(c).UserRepository
	user, err := userRepo.GetUserByEmail(registerUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, RegisterUserResponse{
			Status: models.StatusSomethingWentWrong(),
		})
		return
	}

	if user != nil {
		log.Printf("existing user")
		c.JSON(http.StatusOK, RegisterUserResponse{
			Status:    models.StatusSuccess(),
			IsNewUser: 0,
		})
		return
	}

	// creating hashed password
	hashedPassword, err := util.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, RegisterUserResponse{
			Status: models.StatusSomethingWentWrong(),
		})
		return
	}

	// generate 6 digit OTP
	otp, err := util.GenerateOtp(6)
	if err != nil {
		c.JSON(http.StatusInternalServerError, RegisterUserResponse{
			Status: models.StatusSomethingWentWrong(),
		})
		return
	}

	// convert otp string to uint64
	uiOtp, err := strconv.ParseUint(otp, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, RegisterUserResponse{
			Status: models.StatusSomethingWentWrong(),
		})
		return
	}

	// create a new user in user_verificatons table
	extraDetails := map[string]interface{}{
		"first_name":    request.FirstName,
		"last_name":     request.LastName,
		"password_hash": hashedPassword,
	}

	newUserData := &models.UserVerification{
		EmailId:        request.EmailId,
		Otp:            uiOtp,
		OtpGeneratedAt: time.Now(),
		OtpExpiresAt:   time.Now().Add(10 * time.Minute),
		ExtraDetails:   extraDetails,
	}

	newUser, err := userRepo.CreateUnverifiedUser(&userrepo.CreateUnverifiedUserRequest{
		User: newUserData,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, RegisterUserResponse{
			Status: models.StatusSomethingWentWrong(),
		})
		return
	}

	// TO-DO: send verification otp via aws-ses

	c.JSON(http.StatusCreated, RegisterUserResponse{
		UserID: newUser.UserID.String(),
		Status: models.StatusSuccess(),
	})
}

func LoginUser(c *gin.Context) {
	var request LoginUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, LoginUserResponse{
			Status: models.StatusFailed(fmt.Sprintln("Invalid credentials")),
		})
		return
	}

	userRepo := env.FromContext(c).UserRepository

	loginUser := &userrepo.GetUserByEmailRequest{
		EmailId: request.EmailId,
	}
	user, err := userRepo.GetUserByEmail(loginUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, LoginUserResponse{
			Status: models.StatusSomethingWentWrong(),
		})
		return
	}

	if user == nil {
		c.JSON(http.StatusBadRequest, LoginUserResponse{
			Status: models.StatusFailed(fmt.Sprintln("Incorrect email or password")),
		})
		return
	}

	// verify user password
	ok := util.CheckPasswordHash(request.Password, user.User.PasswordHash)

	if !ok {
		c.JSON(http.StatusBadRequest, LoginUserResponse{
			Status: models.StatusFailed(fmt.Sprintln("Incorrect email or password")),
		})
		return
	}

	// create a JWT token
	jwtToken, err := util.GenerateJwt(user.User.UserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, LoginUserResponse{
			Status: models.StatusSomethingWentWrong(),
		})
		return
	}

	c.JSON(http.StatusOK, LoginUserResponse{
		Status: models.StatusSuccess(),
		Token:  jwtToken,
	})
}

func VerifyOtp(c *gin.Context) {
	var request VerifyOtpRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, VerifyOtpResponse{
			Status: models.StatusFailed(fmt.Sprintln("")),
		})
		return
	}

	userRepo := env.FromContext(c).UserRepository

	// get user by user_id
	user, err := userRepo.GetUnverifiedUserByUserID(&userrepo.GetUnverifiedUserByUserIdRequest{UserID: nanoid.NanoID(request.UserID)})

	if err != nil {
		c.JSON(http.StatusInternalServerError, VerifyOtpResponse{
			Status: models.StatusSomethingWentWrong(),
		})
		return
	}

	if user == nil {
		c.JSON(http.StatusBadRequest, VerifyOtpResponse{
			Status: models.StatusFailed(fmt.Sprintf("No user found with given user_id: %s", request.UserID)),
		})
		return
	}

	// check if we have already created a user with this users email_id
	existingUser, err := userRepo.GetUserByEmail(&userrepo.GetUserByEmailRequest{EmailId: user.User.EmailId})

	if err != nil {
		c.JSON(http.StatusInternalServerError, VerifyOtpResponse{
			Status: models.StatusSomethingWentWrong(),
		})
		return
	}

	if existingUser != nil {
		c.JSON(http.StatusBadRequest, VerifyOtpResponse{
			Status: models.StatusFailed("User already verified, please login to continue"),
		})
		return
	}

	// check if otp has expired
	currTime := time.Now()
	if currTime.After(user.User.OtpExpiresAt) {
		c.JSON(http.StatusOK, VerifyOtpResponse{
			Status: models.StatusFailed("Otp Expired"),
		})
		return
	}

	// match the otp
	if request.Otp != user.User.Otp {
		c.JSON(http.StatusBadRequest, VerifyOtpResponse{
			Status: models.StatusFailed("Wrong Otp"),
		})
		return
	}

	newUser := &models.User{
		FirstName:    user.User.ExtraDetails["first_name"].(string),
		LastName:     user.User.ExtraDetails["last_name"].(string),
		EmailId:      user.User.EmailId,
		PasswordHash: user.User.ExtraDetails["password_hash"].(string),
		UserID:       user.User.UserID,
	}

	// now create a new entry in users table
	createdUser, err := userRepo.CreateUser(&userrepo.CreateUserRequest{
		User: newUser,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, VerifyOtpResponse{
			Status: models.StatusSomethingWentWrong(),
		})
		return
	}

	// set user verified in user_verifications table
	updateUser := &userrepo.UpdateUnverifiedUserRequest{
		UserID: user.User.UserID,
		Values: map[string]interface{}{
			"is_verified": 1,
		},
	}

	if _, err := userRepo.UpdateUnverifiedUser(updateUser); err != nil {
		c.JSON(http.StatusInternalServerError, VerifyOtpResponse{
			Status: models.StatusSomethingWentWrong(),
		})
		return
	}

	// create a new jwt-token
	token, err := util.GenerateJwt(createdUser.UserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, VerifyOtpResponse{
			Status: models.StatusSomethingWentWrong(),
		})
		return
	}

	c.JSON(http.StatusOK, VerifyOtpResponse{
		Status: models.StatusSuccess(),
		Token:  token,
	})
}

func ResendOtp(c *gin.Context) {
	var request ResendOtpRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ResendOtpResponse{
			Status: models.StatusFailed("Missing Params"),
		})
		return
	}

	userRepo := env.FromContext(c).UserRepository

	user, err := userRepo.GetUnverifiedUserByUserID(&userrepo.GetUnverifiedUserByUserIdRequest{
		UserID: nanoid.NanoID(request.UserID),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, ResendOtpResponse{
			Status: models.StatusSomethingWentWrong(),
		})
		return
	}

	// update otp_retry_count
	updateUser := &userrepo.UpdateUnverifiedUserRequest{
		UserID: nanoid.NanoID(request.UserID),
		Values: map[string]interface{}{
			"otp_expires_at":  time.Now().Add(10 * time.Minute),
			"otp_retry_count": user.User.OtpRetryCount + 1,
		},
	}

	if _, err := userRepo.UpdateUnverifiedUser(updateUser); err != nil {
		c.JSON(http.StatusInternalServerError, ResendOtpResponse{
			Status: models.StatusSomethingWentWrong(),
		})
		return
	}

	// TO-DO resend verification code

	c.JSON(http.StatusOK, ResendOtpResponse{
		Status: models.StatusSuccess(),
	})
}
