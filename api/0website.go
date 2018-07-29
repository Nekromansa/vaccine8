package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"fmt"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/justinas/alice"

	"vaccine8/buckets"
	"vaccine8/config"
	"vaccine8/utils"
)

func apiHandlerWebsite(middlewares alice.Chain, router *Router) {
	router.Post("/api/login", middlewares.ThenFunc(apiLogin))
	router.Post("/api/signup", middlewares.ThenFunc(apiSignup))
}

func apiLogin(httpRes http.ResponseWriter, httpReq *http.Request) {
	httpRes.Header().Set("Content-Type", "application/json")

	var formStruct struct {
		Username, Password string
	}

	statusBody := make(map[string]interface{})
	statusCode := http.StatusInternalServerError
	statusMessage := "Invalid Password"

	err := json.NewDecoder(httpReq.Body).Decode(&formStruct)
	if err == nil {
		users, _ := buckets.Users{}.GetFieldValue("Username", formStruct.Username)

		if len(users) == 1 {
			lValid := true
			User := users[0]

			if err = bcrypt.CompareHashAndPassword(User.Password, []byte(formStruct.Password)); err != nil {
				lValid = false
			}

			if User.Workflow != "enabled" && User.Workflow != "active" {
				lValid = false
			}

			if lValid {
				jwtClaims := jwt.MapClaims{}
				jwtClaims["ID"] = User.ID
				jwtClaims["Username"] = User.Username

				statusBody["Redirect"] = "/dashboard"

				cookieExpires := time.Now().Add(time.Hour * 24 * 14) // set the expire time
				jwtClaims["exp"] = cookieExpires.Unix()

				if jwtToken, errJwt := utils.GenerateJWT(jwtClaims); errJwt == nil {
					cookieMonster := &http.Cookie{
						Name: config.Get().COOKIE, Value: jwtToken, Expires: cookieExpires, Path: "/",
					}
					http.SetCookie(httpRes, cookieMonster)
					httpReq.AddCookie(cookieMonster)

					statusCode = http.StatusOK
					statusMessage = "Password Verified"
				}
			}

		}
	} else {
		println(err.Error())
	}

	json.NewEncoder(httpRes).Encode(Message{
		Code:    statusCode,
		Message: statusMessage,
		Body:    statusBody,
	})
}

func apiSignup(httpRes http.ResponseWriter, httpReq *http.Request) {

	httpRes.Header().Set("Content-Type", "application/json")

	statusMessage := ""
	statusBody := make(map[string]interface{})
	statusCode := http.StatusInternalServerError

	var formStruct struct {
		Password, Confirm,
		Fullname, Username,
		Email, Mobile string
	}

	err := json.NewDecoder(httpReq.Body).Decode(&formStruct)
	if err != nil {
		statusMessage = "Error Decoding Form Values " + err.Error()
	} else {
		users, err := buckets.Users{}.GetFieldValue("Username", formStruct.Username)
		if err != nil {
			statusMessage = fmt.Sprintf("Error Validating Username %s", err.Error())
		} else if len(users) > 0 {
			statusMessage = fmt.Sprintf("User exists")
		} else {

			//All Seems Clear, Create New User Now Now

			if formStruct.Password == "" {
				statusMessage += "Password " + IsRequired
			}

			if formStruct.Confirm == "" {
				statusMessage += "Confirm Password " + IsRequired
			}

			if statusMessage == "" {
				if formStruct.Password != formStruct.Confirm {
					statusMessage += "Passwords do not match "
				}
			}

			if strings.HasSuffix(statusMessage, "\n") {
				statusMessage = statusMessage[:len(statusMessage)-2]
			}

			if statusMessage == "" {
				bucketUser := buckets.Users{}
				bucketUser.Workflow = Enabled
				bucketUser.Username = formStruct.Username
				bucketUser.Fullname = formStruct.Fullname
				bucketUser.Mobile = formStruct.Mobile
				bucketUser.Email = formStruct.Email

				hash, _ := bcrypt.GenerateFromPassword([]byte(formStruct.Password), bcrypt.DefaultCost)
				bucketUser.Password = hash

				statusCode = http.StatusOK
				statusMessage = "Registered successfully, please login"
				bucketUser.Create(&bucketUser)

			}
			//All Seems Clear, Create New User Now Now

		}
	}

	json.NewEncoder(httpRes).Encode(Message{
		Code:    statusCode,
		Message: statusMessage,
		Body:    statusBody,
	})
	// //Send E-Mail
}
