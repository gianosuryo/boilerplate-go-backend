package collection

import "go.mongodb.org/mongo-driver/bson/primitive"

type LogCouponRedeem struct {
	ID primitive.ObjectID `bson:"_id"`
	// Connection
	At int64 `bson:"at"`

	SubCompanyId int64  `bson:"sub_company_id"`
	ExternalId   string `bson:"external_id"`
	CouponCode   string `bson:"coupon_code"`
	UserData     string `bson:"user_data"`
	CouponData   string `bson:"coupon_data"`
}

func (LogCouponRedeem) CollectionName() string {
	return "log-coupon-redeem"
}
