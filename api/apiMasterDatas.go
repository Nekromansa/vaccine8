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

type apiMasterDatasStruct struct {
	ID       uint64
	Workflow string
	Createdate,
	Updatedate time.Time

	Age, AgeUOM, Vaccine,
	VaccineDetails, Diseases,
	DiseasesDetails string
}

func apiHandlerMasterDatas(middlewares alice.Chain, router *Router) {
	router.Get("/api/masterdatas", middlewares.ThenFunc(apiMasterDataGet))
	router.Post("/api/masterdatas", middlewares.ThenFunc(apiMasterDataPost))
	router.Get("/api/masterdatas/search", middlewares.ThenFunc(apiMasterDatasSearch))
}

func apiMasterDatasSearch(httpRes http.ResponseWriter, httpReq *http.Request) {

	httpRes.Header().Set("Content-Type", "application/json")

	var statusBody interface{}
	statusCode := http.StatusInternalServerError
	statusMessage := ""

	if claims := utils.VerifyJWT(httpRes, httpReq); claims == nil {
		statusBody = map[string]string{"Redirect": "/"}
	} else {

		searchResults := []buckets.MasterDatas{}
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
			formSearch.Field = "Disease"
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

			searchList := make([]apiMasterDatasStruct, len(searchResults))
			for pos, result := range searchResults {
				searchList[pos].Createdate = result.Createdate
				searchList[pos].Updatedate = result.Updatedate
				searchList[pos].Workflow = result.Workflow

				searchList[pos].ID = result.ID
				searchList[pos].Age = result.Age
				searchList[pos].AgeUOM = result.AgeUOM
				searchList[pos].Vaccine = result.Vaccine
				searchList[pos].VaccineDetails = result.VaccineDetails
				searchList[pos].DiseasesDetails = result.DiseasesDetails
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

func apiMasterDataGet(httpRes http.ResponseWriter, httpReq *http.Request) {

	httpRes.Header().Set("Content-Type", "application/json")

	var statusBody interface{}
	statusCode := http.StatusOK
	statusMessage := ""

	if claims := utils.VerifyJWT(httpRes, httpReq); claims == nil {
		statusBody = map[string]string{"Redirect": "/"}
	} else {

		sMasterDataID := strings.TrimSpace(httpReq.FormValue("id"))
		if sMasterDataID == "" {
			statusCode = http.StatusInternalServerError
			statusMessage = "Error MasterData ID is required to load form"
		} else {
			MasterDataID, _ := strconv.ParseUint(sMasterDataID, 0, 64)
			masterdatasList, err := buckets.MasterDatas{}.GetFieldValue("ID", MasterDataID)
			if err != nil {
				statusMessage = err.Error()
			} else {
				if len(masterdatasList) > 0 {

					statusBody = apiMasterDatasStruct{
						ID:         masterdatasList[0].ID,
						Createdate: masterdatasList[0].Createdate,
						Updatedate: masterdatasList[0].Updatedate,
						Workflow:   masterdatasList[0].Workflow,

						Age:             masterdatasList[0].Age,
						AgeUOM:          masterdatasList[0].AgeUOM,
						Vaccine:         masterdatasList[0].Vaccine,
						Diseases:        masterdatasList[0].Diseases,
						VaccineDetails:  masterdatasList[0].VaccineDetails,
						DiseasesDetails: masterdatasList[0].DiseasesDetails,
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

func apiMasterDataPost(httpRes http.ResponseWriter, httpReq *http.Request) {
	httpRes.Header().Set("Content-Type", "application/json")

	var statusBody interface{}
	statusCode := http.StatusInternalServerError
	statusMessage := ""

	if claims := utils.VerifyJWT(httpRes, httpReq); claims == nil {
		statusBody = map[string]string{"Redirect": "/"}
	} else {

		formStruct := buckets.MasterDatas{}
		err := json.NewDecoder(httpReq.Body).Decode(&formStruct)
		if err != nil {
			statusMessage = "Error Decoding Form Values: " + err.Error()
		} else {

			bucketMasterData := buckets.MasterDatas{}

			if formStruct.ID != 0 {
				bucketMasterDataList, _ := buckets.MasterDatas{}.GetFieldValue("ID", formStruct.ID)
				if len(bucketMasterDataList) != 1 {
					statusMessage = "Error Decoding Form Values: " + err.Error()
				} else {
					bucketMasterData = bucketMasterDataList[0]
				}
			} else {
				formStruct.Workflow = Enabled
			}

			bucketMasterData.Age = formStruct.Age
			bucketMasterData.AgeUOM = formStruct.AgeUOM
			bucketMasterData.Vaccine = formStruct.Vaccine
			bucketMasterData.Diseases = formStruct.Diseases
			bucketMasterData.VaccineDetails = formStruct.VaccineDetails
			bucketMasterData.DiseasesDetails = formStruct.DiseasesDetails

			if statusMessage == "" {
				if bucketMasterData.Age == "" {
					statusMessage += "Age is Required \n"
				}

				if bucketMasterData.AgeUOM == "" {
					statusMessage += "AgeUOM is Required \n"
				}

				if bucketMasterData.Vaccine == "" {
					statusMessage += "Vaccine is Required \n"
				}

				if bucketMasterData.Diseases == "" {
					statusMessage += "Diseases is Required \n"
				}

				if strings.HasSuffix(statusMessage, "\n") {
					statusMessage = statusMessage[:len(statusMessage)-2]
				}
			}

			if statusMessage == "" {
				err = bucketMasterData.Create(&bucketMasterData)
				if err != nil {
					statusMessage = "Error Saving Record: " + err.Error()
				} else {
					statusCode = http.StatusOK
					statusMessage = RecordSaved
					statusBody = bucketMasterData.ID
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
