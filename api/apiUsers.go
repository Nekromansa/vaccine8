package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/coreos/bbolt"
	"github.com/justinas/alice"
	"github.com/timshannon/bolthold"

	"vaccine8/buckets"
	"vaccine8/config"
	"vaccine8/utils"
)

type apiUsersStruct struct {
	ID       uint64
	Workflow string
	Createdate,
	Updatedate time.Time

	Fullname, Username,
	Email, Mobile,
	Password string
}

func apiHandlerUsers(middlewares alice.Chain, router *Router) {
	router.Get("/api/users", middlewares.ThenFunc(apiUserGet))
	router.Post("/api/users", middlewares.ThenFunc(apiUserPost))
	router.Get("/api/users/search", middlewares.ThenFunc(apiUsersSearch))
}

func apiUsersSearch(httpRes http.ResponseWriter, httpReq *http.Request) {

	httpRes.Header().Set("Content-Type", "application/json")

	var statusBody interface{}
	statusCode := http.StatusInternalServerError
	statusMessage := ""

	if claims := utils.VerifyJWT(httpRes, httpReq); claims == nil {
		statusBody = map[string]string{"Redirect": "/"}
	} else {

		searchResults := []buckets.Users{}
		formSearch := new(apiSearch)
		formSearch.Skip, _ = strconv.Atoi(httpReq.FormValue("skip"))
		formSearch.Limit, _ = strconv.Atoi(httpReq.FormValue("limit"))

		formSearch.Text = strings.TrimSpace(httpReq.FormValue("search"))
		formSearch.Field = strings.TrimSpace(httpReq.FormValue("field"))
		if formSearch.Text == "" {
			formSearch.Text = "."
		} else {
			formSearch.Text = regexp.QuoteMeta(formSearch.Text)
		}

		switch formSearch.Field {
		default:
			formSearch.Field = strings.Title(strings.ToLower(formSearch.Field))
		case "":
			formSearch.Field = "Username"
		}

		if err := config.Get().BoltHold.Bolt().View(func(tx *bolt.Tx) error {
			err := config.Get().BoltHold.Find(&searchResults,
				bolthold.Where(formSearch.Field).RegExp(
					regexp.MustCompile(`(?im)`+formSearch.Text)).SortBy("ID").Reverse().Limit(formSearch.Limit).Skip(formSearch.Skip),
			)
			return err
		}); err != nil {
			statusMessage = err.Error()
		} else {

			searchList := make([]apiSearchResult, len(searchResults))
			for pos, result := range searchResults {
				searchList[pos].ID = result.ID
				searchList[pos].Date = JSONTime(result.Updatedate)
				searchList[pos].Details = fmt.Sprintf("%v", result.Fullname)
			}

			statusCode = http.StatusOK
			statusBody = searchList
		}
	}

	json.NewEncoder(httpRes).Encode(Message{
		Code:    statusCode,
		Body:    statusBody,
		Message: statusMessage,
	})
}

func apiUserGet(httpRes http.ResponseWriter, httpReq *http.Request) {

	httpRes.Header().Set("Content-Type", "application/json")

	var statusBody interface{}
	statusCode := http.StatusOK
	statusMessage := ""

	if claims := utils.VerifyJWT(httpRes, httpReq); claims == nil {
		statusBody = map[string]string{"Redirect": "/"}
	} else {

		sUserID := strings.TrimSpace(httpReq.FormValue("id"))
		if sUserID == "" {
			statusCode = http.StatusInternalServerError
			statusMessage = "Error User ID is required to load form"
		} else {
			UserID, _ := strconv.ParseUint(sUserID, 0, 64)
			usersList, err := buckets.Users{}.GetFieldValue("ID", UserID)
			if err != nil {
				statusMessage = err.Error()
			} else {
				if len(usersList) > 0 {

					statusBody = apiUsersStruct{
						ID:         usersList[0].ID,
						Createdate: usersList[0].Createdate,
						Updatedate: usersList[0].Updatedate,
						Workflow:   usersList[0].Workflow,

						Fullname: usersList[0].Fullname,
						Username: usersList[0].Username,
						Email:    usersList[0].Email,
						Mobile:   usersList[0].Mobile,
					}
				}
			}
		}
	}

	json.NewEncoder(httpRes).Encode(Message{
		Code:    statusCode,
		Body:    statusBody,
		Message: statusMessage,
	})
}

func apiUserPost(httpRes http.ResponseWriter, httpReq *http.Request) {
	httpRes.Header().Set("Content-Type", "application/json")

	var statusBody interface{}
	statusCode := http.StatusInternalServerError
	statusMessage := ""

	if claims := utils.VerifyJWT(httpRes, httpReq); claims == nil {
		statusBody = map[string]string{"Redirect": "/"}
	} else {

		formStruct := buckets.Users{}
		err := json.NewDecoder(httpReq.Body).Decode(&formStruct)
		if err != nil {
			statusMessage = "Error Decoding Form Values: " + err.Error()
		} else {

			bucketUser := buckets.Users{}

			if formStruct.ID != 0 {
				bucketUserList, _ := buckets.Users{}.GetFieldValue("ID", formStruct.ID)
				if len(bucketUserList) != 1 {
					statusMessage = "Error Decoding Form Values: " + err.Error()
				} else {
					bucketUser = bucketUserList[0]
				}
			} else {
				formStruct.Workflow = Enabled
			}

			bucketUser.Fullname = formStruct.Fullname
			bucketUser.Username = formStruct.Username
			bucketUser.Email = formStruct.Email
			bucketUser.Mobile = formStruct.Mobile

			if statusMessage == "" {
				if bucketUser.Fullname == "" {
					statusMessage += "Fullname is Required \n"
				}

				if bucketUser.Mobile == "" {
					statusMessage += "Mobile is Required \n"
				}

				if bucketUser.Username == "" {
					statusMessage += "Username is Required \n"
				}

				if strings.HasSuffix(statusMessage, "\n") {
					statusMessage = statusMessage[:len(statusMessage)-2]
				}
			}

			if statusMessage == "" {
				err = bucketUser.Create(&bucketUser)
				if err != nil {
					statusMessage = "Error Saving Record: " + err.Error()
				} else {
					statusCode = http.StatusOK
					statusMessage = RecordSaved
					statusBody = bucketUser.ID
				}
			}
		}
	}

	json.NewEncoder(httpRes).Encode(Message{
		Code:    statusCode,
		Body:    statusBody,
		Message: statusMessage,
	})
}
