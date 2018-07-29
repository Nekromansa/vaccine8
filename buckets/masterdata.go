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
	masterdataRecord.DiseasesDetails = "Bacillus Calmette Guerin-Tuberculosis"
	masterdata.Create(&masterdataRecord)

	masterdataRecord.Age = "10"
	masterdataRecord.AgeUOM = "Weeks"
	masterdataRecord.Vaccine = "OPV 2"
	masterdataRecord.Diseases = "Polio"
	masterdataRecord.VaccineDetails = ""
	masterdataRecord.DiseasesDetails = ""
	masterdata.Create(&masterdataRecord)

	masterdataRecord.Age = "10"
	masterdataRecord.AgeUOM = "Weeks"
	masterdataRecord.Vaccine = "Pentavalent 2"
	masterdataRecord.Diseases = "DPT,HIB,Hepatitis B"
	masterdataRecord.VaccineDetails = ""
	masterdataRecord.DiseasesDetails = "( Diphtheria, Pertussis or Whooping Cough, Tetanus),(Haemophilus influenzae type B)"
	masterdata.Create(&masterdataRecord)

	masterdataRecord.Age = "14"
	masterdataRecord.AgeUOM = "Weeks"
	masterdataRecord.Vaccine = "OPV 3"
	masterdataRecord.Diseases = "Polio"
	masterdataRecord.VaccineDetails = ""
	masterdataRecord.DiseasesDetails = ""
	masterdata.Create(&masterdataRecord)

	masterdataRecord.Age = "14"
	masterdataRecord.AgeUOM = "Weeks"
	masterdataRecord.Vaccine = "Pentavalent 3"
	masterdataRecord.Diseases = "DPT,HIB,Hepatitis B"
	masterdataRecord.VaccineDetails = ""
	masterdataRecord.DiseasesDetails = "( Diphtheria, Pertussis or Whooping Cough, Tetanus),(Haemophilus influenzae type B)"
	masterdata.Create(&masterdataRecord)

	masterdataRecord.Age = "14"
	masterdataRecord.AgeUOM = "Weeks"
	masterdataRecord.Vaccine = "PCV 3"
	masterdataRecord.Diseases = "Pneumonia & Otitismedia"
	masterdataRecord.VaccineDetails = ""
	masterdataRecord.DiseasesDetails = ""
	masterdata.Create(&masterdataRecord)

	masterdataRecord.Age = "6"
	masterdataRecord.AgeUOM = "Months"
	masterdataRecord.Vaccine = "Vitamin A 1st Dose"
	masterdataRecord.Diseases = "Vitamin A deficiency"
	masterdataRecord.VaccineDetails = ""
	masterdataRecord.DiseasesDetails = "VAD"
	masterdata.Create(&masterdataRecord)

	masterdataRecord.Age = "9"
	masterdataRecord.AgeUOM = "Months"
	masterdataRecord.Vaccine = "Measles Vaccine"
	masterdataRecord.Diseases = "Measles "
	masterdataRecord.VaccineDetails = ""
	masterdataRecord.DiseasesDetails = ""
	masterdata.Create(&masterdataRecord)

	masterdataRecord.Age = "9"
	masterdataRecord.AgeUOM = "Months"
	masterdataRecord.Vaccine = "Yellow fever vaccine"
	masterdataRecord.Diseases = "Yellow fever"
	masterdataRecord.VaccineDetails = ""
	masterdataRecord.DiseasesDetails = ""
	masterdata.Create(&masterdataRecord)

	masterdataRecord.Age = "12-24"
	masterdataRecord.AgeUOM = "Months"
	masterdataRecord.Vaccine = "Meningitis vaccine"
	masterdataRecord.Diseases = "Meningitis & Septicaemia"
	masterdataRecord.VaccineDetails = "Nimenrix"
	masterdataRecord.DiseasesDetails = ""
	masterdata.Create(&masterdataRecord)

	masterdataRecord.Age = "15-18"
	masterdataRecord.AgeUOM = "Months"
	masterdataRecord.Vaccine = "MMR (Priorix)"
	masterdataRecord.Diseases = "Measles, Mumps, Rubella"
	masterdataRecord.VaccineDetails = ""
	masterdataRecord.DiseasesDetails = "German "
	masterdata.Create(&masterdataRecord)

	masterdataRecord.Age = "15-18"
	masterdataRecord.AgeUOM = "Months"
	masterdataRecord.Vaccine = "Chicken pox vaccine"
	masterdataRecord.Diseases = "Measles, Mumps, Rubella"
	masterdataRecord.VaccineDetails = ""
	masterdataRecord.DiseasesDetails = "German, Chicken pox, Varicella"
	masterdata.Create(&masterdataRecord)

	masterdataRecord.Age = "24"
	masterdataRecord.AgeUOM = "Months"
	masterdataRecord.Vaccine = "Typhoid Vaccine"
	masterdataRecord.Diseases = "Typhoid fever"
	masterdataRecord.VaccineDetails = "Typherix"
	masterdataRecord.DiseasesDetails = ""
	masterdata.Create(&masterdataRecord)





	return nil
}
