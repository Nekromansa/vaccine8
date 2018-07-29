package api

import (
	"encoding/base64"
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

type apiProfileStruct struct {
	ID, UserID uint64
	Workflow   string
	Createdate, Dob,
	Updatedate time.Time

	IsDependent bool

	Code, Email, Mobile, Fullname,
	Title, Firstname, Lastname,
	Othername, Gender, MaritalStatus,
	Street, City, State, Country,
	StateOfOrigin, LocalGovtArea,

	Image, Relationship,
	Description, User string
}

func apiHandlerProfiles(middlewares alice.Chain, router *Router) {
	router.Get("/api/profiles", middlewares.ThenFunc(apiProfilesGet))
	router.Post("/api/profiles", middlewares.ThenFunc(apiProfilesPost))
	router.Get("/api/profiles/search", middlewares.ThenFunc(apiProfilesSearch))
}

func apiProfilesSearch(httpRes http.ResponseWriter, httpReq *http.Request) {

	httpRes.Header().Set("Content-Type", "application/json")

	var statusBody interface{}
	statusCode := http.StatusInternalServerError
	statusMessage := ""

	if claims := utils.VerifyJWT(httpRes, httpReq); claims == nil {
		statusBody = map[string]string{"Redirect": "/"}
	} else {

		searchResults := []buckets.Profiles{}
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
			formSearch.Field = "Fullname"
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

func apiProfilesGet(httpRes http.ResponseWriter, httpReq *http.Request) {

	httpRes.Header().Set("Content-Type", "application/json")

	var statusBody interface{}
	statusCode := http.StatusOK
	statusMessage := ""

	if claims := utils.VerifyJWT(httpRes, httpReq); claims == nil {
		statusBody = map[string]string{"Redirect": "/"}
	} else {

		profilesList, err := buckets.Profiles{}.GetFieldValue("ID", uint64(claims["ID"].(float64)))
		if err != nil {
			statusMessage = err.Error()
		} else {
			if len(profilesList) > 0 {
				if len(profilesList[0].Image) > 3 {
					profilesList[0].Image += "?" + strings.ToLower(utils.RandomString(3))
				}

				User := ""
				if profilesList[0].UserID != 0 {
					usersList, _ := buckets.Users{}.GetFieldValue("ID", profilesList[0].UserID)
					if len(usersList) == 1 {
						User = usersList[0].Fullname
					}
				}

				statusBody = apiProfileStruct{
					User:   User,
					UserID: profilesList[0].UserID,

					ID:         profilesList[0].ID,
					Createdate: profilesList[0].Createdate,
					Updatedate: profilesList[0].Updatedate,
					Workflow:   profilesList[0].Workflow,

					IsDependent: profilesList[0].IsDependent,

					Code:   profilesList[0].Code,
					Email:  profilesList[0].Email,
					Mobile: profilesList[0].Mobile,

					Title:     profilesList[0].Title,
					Fullname:  profilesList[0].Fullname,
					Firstname: profilesList[0].Firstname,
					Lastname:  profilesList[0].Lastname,
					Othername: profilesList[0].Othername,

					Dob:    profilesList[0].Dob,
					Gender: profilesList[0].Gender,

					MaritalStatus: profilesList[0].MaritalStatus,
					Street:        profilesList[0].Street,
					City:          profilesList[0].City,
					State:         profilesList[0].State,
					Country:       profilesList[0].Country,
					StateOfOrigin: profilesList[0].StateOfOrigin,
					LocalGovtArea: profilesList[0].LocalGovtArea,

					Image:        profilesList[0].Image,
					Relationship: profilesList[0].Relationship,
					Description:  profilesList[0].Description,
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

func apiProfilesPost(httpRes http.ResponseWriter, httpReq *http.Request) {
	httpRes.Header().Set("Content-Type", "application/json")

	var statusBody interface{}
	statusCode := http.StatusInternalServerError
	statusMessage := ""

	if claims := utils.VerifyJWT(httpRes, httpReq); claims == nil {
		statusBody = map[string]string{"Redirect": "/"}
	} else {

		formStruct := buckets.Profiles{}
		err := json.NewDecoder(httpReq.Body).Decode(&formStruct)
		if err != nil {
			statusMessage = "Error Decoding Form Values: " + err.Error()
		} else {

			bucketProfile := buckets.Profiles{}

			bucketProfileList, _ := buckets.Profiles{}.GetFieldValue("ID", uint64(claims["ID"].(float64)))
			if len(bucketProfileList) != 1 {
				statusMessage = "Error Decoding Form Values: " + err.Error()
			} else {
				bucketProfile = bucketProfileList[0]
			}

			bucketProfile.Title = formStruct.Title
			bucketProfile.Firstname = formStruct.Firstname
			bucketProfile.Lastname = formStruct.Lastname
			bucketProfile.Othername = formStruct.Othername
			bucketProfile.Fullname = fmt.Sprintf("%s %s %s %s",
				bucketProfile.Title, bucketProfile.Firstname,
				bucketProfile.Lastname, bucketProfile.Othername)

			bucketProfile.Email = formStruct.Email
			bucketProfile.Mobile = formStruct.Mobile
			bucketProfile.Gender = formStruct.Gender
			bucketProfile.Dob = formStruct.Dob

			bucketProfile.IsDependent = formStruct.IsDependent
			bucketProfile.MaritalStatus = formStruct.MaritalStatus

			bucketProfile.Street = formStruct.Street
			bucketProfile.City = formStruct.City
			bucketProfile.State = formStruct.State
			bucketProfile.Country = formStruct.Country

			bucketProfile.StateOfOrigin = formStruct.StateOfOrigin
			bucketProfile.LocalGovtArea = formStruct.LocalGovtArea

			bucketProfile.Relationship = formStruct.Relationship
			bucketProfile.Description = formStruct.Description

			if statusMessage == "" {

				if bucketProfile.Title == "" {
					statusMessage += "Title is Required \n"
				}

				if bucketProfile.Firstname == "" {
					statusMessage += "Firstname is Required \n"
				}

				if bucketProfile.Lastname == "" {
					statusMessage += "Lastname is Required \n"
				}

				if strings.HasSuffix(statusMessage, "\n") {
					statusMessage = statusMessage[:len(statusMessage)-2]
				}
			}

			if statusMessage == "" {

				if !strings.HasPrefix(formStruct.Image, "data:image/") {
					formStruct.Image = ""
				} else {
					base64Bytes, errNew := base64.StdEncoding.DecodeString(
						strings.Split(formStruct.Image, "base64,")[1])

					if base64Bytes != nil && errNew == nil {
						fileExt, fileType := utils.GetFileExt(formStruct.Image[:20])

						if fileExt != "" {
							fileName := fmt.Sprintf("dp-%s%s", utils.RandomString(12), fileExt)
							formStruct.Image = utils.SaveFile(fileName, fileType, base64Bytes)
						}
					}
				}

				if formStruct.Image != "" {
					bucketProfile.Image = formStruct.Image
				}

				err = bucketProfile.Create(&bucketProfile)
				if err != nil {
					statusMessage = "Error Saving Record: " + err.Error()
				} else {
					statusCode = http.StatusOK
					statusMessage = RecordSaved
					statusBody = bucketProfile.ID
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
