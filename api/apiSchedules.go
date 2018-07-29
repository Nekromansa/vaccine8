package api

import (
	"encoding/json"
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

type apiScheduleStruct struct {
	ID, UserID, ProfileID,
	MasterID uint64

	Workflow string
	Createdate, Dob,
	Updatedate time.Time

	Vaccined bool
	VaccinedDate,
	VaccineDue time.Time
	VaccineLocation,

	User, Profile, Description,

	Age, AgeUOM, Vaccine,
	Diseases, VaccineDetails,
	DiseasesDetails string
}

func apiHandlerSchedules(middlewares alice.Chain, router *Router) {
	router.Get("/api/schedules", middlewares.ThenFunc(apiSchedulesGet))
	router.Post("/api/schedules", middlewares.ThenFunc(apiSchedulesPost))
	router.Get("/api/schedules/search", middlewares.ThenFunc(apiSchedulesSearch))
}

func apiSchedulesSearch(httpRes http.ResponseWriter, httpReq *http.Request) {

	httpRes.Header().Set("Content-Type", "application/json")

	var statusBody interface{}
	statusCode := http.StatusInternalServerError
	statusMessage := ""

	if claims := utils.VerifyJWT(httpRes, httpReq); claims == nil {
		statusBody = map[string]string{"Redirect": "/"}
	} else {

		searchResults := []buckets.Schedules{}
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

			searchList := make([]apiScheduleStruct, len(searchResults))
			for pos, result := range searchResults {
				User := ""
				if result.UserID != 0 {
					usersList, _ := buckets.Users{}.GetFieldValue("ID", result.UserID)
					if len(usersList) == 1 {
						User = usersList[0].Fullname
					}
				}

				Profile := ""
				if result.ProfileID != 0 {
					usersList, _ := buckets.Profiles{}.GetFieldValue("ID", result.ProfileID)
					if len(usersList) == 1 {
						Profile = usersList[0].Fullname
					}
				}

				var Age, AgeUOM, Vaccine,
					VaccineDetails, Diseases,
					DiseasesDetails string

				if result.MasterID != 0 {
					masterdataList, _ := buckets.MasterDatas{}.GetFieldValue("ID", result.MasterID)
					if len(masterdataList) == 1 {
						Age = masterdataList[0].Age
						AgeUOM = masterdataList[0].AgeUOM
						Vaccine = masterdataList[0].Vaccine
						VaccineDetails = masterdataList[0].VaccineDetails
						Diseases = masterdataList[0].Diseases
						DiseasesDetails = masterdataList[0].DiseasesDetails
					}
				}

				searchList[pos].Createdate = result.Createdate
				searchList[pos].Updatedate = result.Updatedate
				searchList[pos].Workflow = result.Workflow

				searchList[pos].ID = result.ID
				searchList[pos].User = User
				searchList[pos].UserID = result.UserID
				searchList[pos].Profile = Profile
				searchList[pos].ProfileID = result.ProfileID

				searchList[pos].Age = Age
				searchList[pos].AgeUOM = AgeUOM
				searchList[pos].Vaccine = Vaccine
				searchList[pos].Diseases = Diseases
				searchList[pos].VaccineDetails = VaccineDetails
				searchList[pos].DiseasesDetails = DiseasesDetails
				searchList[pos].AgeUOM = AgeUOM
				searchList[pos].MasterID = result.MasterID

				searchList[pos].Vaccined = result.Vaccined
				searchList[pos].VaccineDue = result.VaccineDue
				searchList[pos].VaccinedDate = result.VaccinedDate
				searchList[pos].VaccineLocation = result.VaccineLocation
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

func apiSchedulesGet(httpRes http.ResponseWriter, httpReq *http.Request) {

	httpRes.Header().Set("Content-Type", "application/json")

	var statusBody interface{}
	statusCode := http.StatusOK
	statusMessage := ""

	if claims := utils.VerifyJWT(httpRes, httpReq); claims == nil {
		statusBody = map[string]string{"Redirect": "/"}
	} else {

		schedulesList, err := buckets.Schedules{}.GetFieldValue("ID", uint64(claims["ID"].(float64)))
		if err != nil {
			statusMessage = err.Error()
		} else {
			if len(schedulesList) > 0 {

				User := ""
				if schedulesList[0].UserID != 0 {
					usersList, _ := buckets.Users{}.GetFieldValue("ID", schedulesList[0].UserID)
					if len(usersList) == 1 {
						User = usersList[0].Fullname
					}
				}

				Profile := ""
				if schedulesList[0].ProfileID != 0 {
					usersList, _ := buckets.Profiles{}.GetFieldValue("ID", schedulesList[0].ProfileID)
					if len(usersList) == 1 {
						Profile = usersList[0].Fullname
					}
				}

				var Age, AgeUOM, Vaccine,
					VaccineDetails, Diseases,
					DiseasesDetails string

				if schedulesList[0].MasterID != 0 {
					masterdataList, _ := buckets.MasterDatas{}.GetFieldValue("ID", schedulesList[0].MasterID)
					if len(masterdataList) == 1 {
						Age = masterdataList[0].Age
						AgeUOM = masterdataList[0].AgeUOM
						Vaccine = masterdataList[0].Vaccine
						VaccineDetails = masterdataList[0].VaccineDetails
						Diseases = masterdataList[0].Diseases
						DiseasesDetails = masterdataList[0].DiseasesDetails
					}
				}

				statusBody = apiScheduleStruct{
					Age:             Age,
					AgeUOM:          AgeUOM,
					Vaccine:         Vaccine,
					Diseases:        Diseases,
					VaccineDetails:  VaccineDetails,
					DiseasesDetails: DiseasesDetails,
					MasterID:        schedulesList[0].MasterID,

					User:   User,
					UserID: schedulesList[0].UserID,

					Profile:   Profile,
					ProfileID: schedulesList[0].ProfileID,

					ID:         schedulesList[0].ID,
					Createdate: schedulesList[0].Createdate,
					Updatedate: schedulesList[0].Updatedate,
					Workflow:   schedulesList[0].Workflow,

					Vaccined:        schedulesList[0].Vaccined,
					VaccinedDate:    schedulesList[0].VaccinedDate,
					VaccineDue:      schedulesList[0].VaccineDue,
					VaccineLocation: schedulesList[0].VaccineLocation,
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

func apiSchedulesPost(httpRes http.ResponseWriter, httpReq *http.Request) {
	httpRes.Header().Set("Content-Type", "application/json")

	var statusBody interface{}
	statusCode := http.StatusInternalServerError
	statusMessage := ""

	if claims := utils.VerifyJWT(httpRes, httpReq); claims == nil {
		statusBody = map[string]string{"Redirect": "/"}
	} else {

		formStruct := buckets.Schedules{}
		err := json.NewDecoder(httpReq.Body).Decode(&formStruct)
		if err != nil {
			statusMessage = "Error Decoding Form Values: " + err.Error()
		} else {

			bucketSchedule := buckets.Schedules{}

			bucketScheduleList, _ := buckets.Schedules{}.GetFieldValue("ID", uint64(claims["ID"].(float64)))
			if len(bucketScheduleList) != 1 {
				statusMessage = "Error Decoding Form Values: " + err.Error()
			} else {
				bucketSchedule = bucketScheduleList[0]
			}

			bucketSchedule.Vaccined = formStruct.Vaccined
			bucketSchedule.VaccineDue = formStruct.VaccineDue
			bucketSchedule.VaccinedDate = formStruct.VaccinedDate
			bucketSchedule.VaccineLocation = formStruct.VaccineLocation

			if statusMessage == "" {

				// if bucketSchedule.Title == "" {
				// 	statusMessage += "Title is Required \n"
				// }

				if strings.HasSuffix(statusMessage, "\n") {
					statusMessage = statusMessage[:len(statusMessage)-2]
				}
			}

			if statusMessage == "" {
				err = bucketSchedule.Create(&bucketSchedule)
				if err != nil {
					statusMessage = "Error Saving Record: " + err.Error()
				} else {
					statusCode = http.StatusOK
					statusMessage = RecordSaved
					statusBody = bucketSchedule.ID
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
