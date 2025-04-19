package helper

import (
	"errors"
	"os"
	"sarana-dafa-ai-service/model/collection"
	"sarana-dafa-ai-service/storage/env"
	"time"

	elastic "github.com/elastic/go-elasticsearch/v8"
	"github.com/geronimo794/go-mongolog"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/go-extras/elogrus.v8"
	"gorm.io/gorm"
)

var errorLog = logrus.New()
var accessLog = logrus.New()
var generateLog = logrus.New()
var jobVacancyLog = logrus.New()
var couponRedeemLog = logrus.New()
var emailLog = logrus.New()

const logDir = "files/logs/"

func InitLogger(db *mongo.Database) {

	// Text formatter logrus
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = time.RFC1123Z
	customFormatter.FullTimestamp = true
	errorLog.SetFormatter(customFormatter)
	accessLog.SetFormatter(customFormatter)
	couponRedeemLog.SetFormatter(customFormatter)
	emailLog.SetFormatter(customFormatter)

	accessLog.SetLevel(logrus.InfoLevel)
	errorLog.SetLevel(logrus.WarnLevel)
	couponRedeemLog.SetLevel(logrus.InfoLevel)
	emailLog.SetLevel(logrus.InfoLevel)

	// Temporary disable mongodb hook
	// setUpAccessLog(db)
	// setUpErrorLog(db)
	// setUpCouponRedeemLog(db)
	// setUpEmailLog(db)
}

func InitLoggerElastic(cl *elastic.Client) {

	// Text formatter logrus
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = time.RFC1123Z
	customFormatter.FullTimestamp = true
	errorLog.SetFormatter(customFormatter)
	accessLog.SetFormatter(customFormatter)
	generateLog.SetFormatter(customFormatter)
	jobVacancyLog.SetFormatter(customFormatter)

	setUpAccessElasticLog(cl)
	setUpErrorElasticLog(cl)
}

func setUpAccessElasticLog(cl *elastic.Client) {
	hookAccessLog, err := elogrus.NewAsyncElasticHook(cl, os.Getenv(env.LOG_ELASTIC_ORIGIN_HOST), logrus.InfoLevel, os.Getenv(env.LOG_ELASTIC_ACCESS_LOG_INDEX))
	if err != nil {
		accessLog.Panic(err)
	}

	// Set error log level and hook
	accessLog.SetLevel(logrus.InfoLevel)
	accessLog.Hooks.Add(hookAccessLog)
}

func setUpErrorElasticLog(cl *elastic.Client) {
	hookErrorLog, err := elogrus.NewAsyncElasticHook(cl, os.Getenv(env.LOG_ELASTIC_ORIGIN_HOST), logrus.WarnLevel, os.Getenv(env.LOG_ELASTIC_ERROR_LOG_INDEX))
	if err != nil {
		errorLog.Panic(err)
	}

	// Set error log level and hook
	errorLog.SetLevel(logrus.WarnLevel)
	errorLog.Hooks.Add(hookErrorLog)
}

func setUpAccessLog(db *mongo.Database) {
	// Define mongolog access log
	hookAccessLog, err := mongolog.NewHookDatabase(db, collection.LogAccess{}.CollectionName())
	if err != nil {
		errorLog.Panic(err)
	}

	err = hookAccessLog.SetFailoverFilePath(logDir + "mongolog-failover-access.log")
	if err != nil {
		errorLog.Panic(err)
	}
	hookAccessLog.SetWriteTimeout(5 * time.Second)
	hookAccessLog.SetIsAsync(true)

	// Set acces log level and hook
	accessLog.Hooks.Add(hookAccessLog)
}
func setUpErrorLog(db *mongo.Database) {
	// Define mongolog error
	hookErrorLog, err := mongolog.NewHookDatabase(db, collection.LogError{}.CollectionName())
	if err != nil {
		errorLog.Panic(err)
	}

	err = hookErrorLog.SetFailoverFilePath(logDir + "mongolog-failover-error.log")
	if err != nil {
		errorLog.Panic(err)
	}

	hookErrorLog.SetWriteTimeout(5 * time.Second)
	hookErrorLog.SetIsAsync(true)

	// Set error log level and hook
	errorLog.Hooks.Add(hookErrorLog)
}
func setUpCouponRedeemLog(db *mongo.Database) {
	// Define mongolog error
	hookCouponRedeemLog, err := mongolog.NewHookDatabase(db, collection.LogCouponRedeem{}.CollectionName())
	if err != nil {
		couponRedeemLog.Panic(err)
	}

	err = hookCouponRedeemLog.SetFailoverFilePath(logDir + "mongolog-failover-coupon-redeem.log")
	if err != nil {
		couponRedeemLog.Panic(err)
	}

	hookCouponRedeemLog.SetWriteTimeout(5 * time.Second)
	hookCouponRedeemLog.SetIsAsync(true)

	// Set error log level and hook.
	couponRedeemLog.Hooks.Add(hookCouponRedeemLog)
}
func setUpEmailLog(db *mongo.Database) {
	// Define mongolog error
	hookEmailLog, err := mongolog.NewHookDatabase(db, collection.LogEmail{}.CollectionName())
	if err != nil {
		emailLog.Panic(err)
	}

	err = hookEmailLog.SetFailoverFilePath(logDir + "mongolog-failover-email.log")
	if err != nil {
		emailLog.Panic(err)
	}

	hookEmailLog.SetWriteTimeout(5 * time.Second)
	hookEmailLog.SetIsAsync(true)

	// Set error log level and hook
	emailLog.Hooks.Add(hookEmailLog)
}

func PanicIfError(err error) {
	if err != nil {
		// If the error is error not found, no need to be panic
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			panic(err)
		}
	}
}
func ErrorLog() *logrus.Logger {
	return errorLog
}
func AccessLog() *logrus.Logger {
	return accessLog
}
func GenerateLog() *logrus.Logger {
	return generateLog
}
func JobVacancyLog() *logrus.Logger {
	return jobVacancyLog
}
func CouponRedeemLog() *logrus.Logger {
	return couponRedeemLog
}
func EmailLog() *logrus.Logger {
	return emailLog
}
