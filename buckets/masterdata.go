package buckets

import (
	"fmt"
	"log"
	"time"

	"github.com/coreos/bbolt"
	"github.com/timshannon/bolthold"

	"vaccine8/config"
)

const (
	Weeks  = "Weeks"
	Months = "Months"
	Years  = "Years"
)

//MasterDatas ...
type MasterDatas struct {
	ID       uint64
	Workflow string
	Createdate,
	Updatedate time.Time

	Age, AgeUOM, Vaccine,
	VaccineDetails, Diseases,
	DiseasesDetails string
}

func (masterdata MasterDatas) bucketName() string {
	return "MasterDatas"
}

//Create ...
func (masterdata MasterDatas) Create(bucketType *MasterDatas) (err error) {

	if err = config.Get().BoltHold.Bolt().Update(func(tx *bolt.Tx) error {

		if bucketType.Createdate.IsZero() {
			bucketType.Createdate = time.Now()
			bucketType.Updatedate = bucketType.Createdate
		}

		if bucketType.ID == 0 {
			bucket := tx.Bucket([]byte(masterdata.bucketName()))
			bucketType.ID, _ = bucket.NextSequence()
			bucketType.Createdate = time.Now()
		} else {
			bucketType.Updatedate = time.Now()
		}

		err = config.Get().BoltHold.TxUpsert(tx, bucketType.ID, bucketType)
		return err
	}); err != nil {
		log.Printf(err.Error())
	}
	return
}

//List ...
func (masterdata MasterDatas) List() (resultsALL []string) {
	var results []MasterDatas

	if err := config.Get().BoltHold.Bolt().View(func(tx *bolt.Tx) error {
		err := config.Get().BoltHold.Find(&results, bolthold.Where("ID").Gt(uint64(0)))
		return err
	}); err != nil {
		log.Printf(err.Error())
	} else {
		for _, record := range results {
			resultsALL = append(resultsALL, fmt.Sprintf("%+v", record))
		}
	}
	return
}

//GetFieldValue ...
func (masterdata MasterDatas) GetFieldValue(Field string, Value interface{}) (results []MasterDatas, err error) {

	if len(Field) > 0 {
		if err = config.Get().BoltHold.Bolt().View(func(tx *bolt.Tx) error {
			err = config.Get().BoltHold.Find(&results, bolthold.Where(Field).Eq(Value).SortBy("ID").Reverse())
			return err
		}); err != nil {
			log.Printf(err.Error())
		}
	}
	return
}

//Setup ....
func (masterdata MasterDatas) Setup() (err error) {
	var masterdataRecord MasterDatas
	masterdataRecord.Workflow = "enabled"
	masterdataRecord.Age = "0"
	masterdataRecord.AgeUOM = "Weeks"
	masterdataRecord.Vaccine = "HBV 1"
	masterdataRecord.Diseases = "Hepatitis B"
	masterdataRecord.VaccineDetails = "hepatitis B Vaccine"
	masterdataRecord.DiseasesDetails = ""
	masterdata.Create(&masterdataRecord)

	masterdataRecord.Age = "0"
	masterdataRecord.AgeUOM = "Weeks"
	masterdataRecord.Vaccine = "OPV 0"
	masterdataRecord.Diseases = "Polio"
	masterdataRecord.VaccineDetails = "Oral polio Vaccine"
	masterdataRecord.DiseasesDetails = ""
	masterdata.Create(&masterdataRecord)

	masterdataRecord.Age = "0"
	masterdataRecord.AgeUOM = "Weeks"
	masterdataRecord.Vaccine = "BCG"
	masterdataRecord.Diseases = "BCG"
	masterdataRecord.VaccineDetails = "Bacillus Calmette Guerin-Tuberculosis"
	masterdataRecord.DiseasesDetails = ""
	masterdata.Create(&masterdataRecord)

	return nil
}
