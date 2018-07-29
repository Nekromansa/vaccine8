package buckets

import (
	"strings"

	"github.com/coreos/bbolt"

	"vaccine8/config"
)

//Init ...
func Init() {
	Empty("/all")
	MasterDatas{}.Setup()
}

var allowedBuckets = map[string]bool{
	"MasterData": true,
	"Profiles":   true,
	"Schedules":  true,
	"Users":      true,
}

//BucketList ...
func BucketList(bucketName string) (bucketList []string) {
	if len(bucketName) > 0 {
		bucketName = bucketName[1:]
	}
	switch bucketName {
	default:
		bucketList = append(bucketList, "Please Specify Bucket --> Bucket "+bucketName+" Invalid!!")

	case "MasterData":
		bucketList = append(bucketList, strings.Join(new(MasterDatas).List(), "\n"))

	case "Profiles":
		bucketList = append(bucketList, strings.Join(new(Profiles).List(), "\n"))

	case "Schedules":
		bucketList = append(bucketList, strings.Join(new(Schedules).List(), "\n"))

	case "Users":
		bucketList = append(bucketList, strings.Join(new(Users).List(), "\n"))
	}
	return
}

//Empty ...
func Empty(bucketName string) (Message []string) {

	switch bucketName {
	default:
		bucketName = bucketName[1:]
		if allowedBuckets[bucketName] {
			Message = append(Message, empty(bucketName))
		} else {
			Message = append(Message, "Please Specify Bucket")
		}

	case "/all":
		for bucket := range allowedBuckets {
			bucket = strings.Title(strings.ToLower(bucket))
			Message = append(Message, empty(bucket))
		}
		//Setup MasterData
		MasterDatas{}.Setup()
	}
	return Message
}

func empty(bucketName string) string {
	if allowedBuckets[bucketName] {
		if err := config.Get().BoltHold.Bolt().Update(func(tx *bolt.Tx) (err error) {
			tx.DeleteBucket([]byte(bucketName))
			_, err = tx.CreateBucket([]byte(bucketName))
			return
		}); err != nil {
			return bucketName + " Bucket --> " + err.Error()
		}
		return bucketName + " Bucket -->  Emptied "
	}
	return bucketName + " Bucket -->  Does not Exist"
}
