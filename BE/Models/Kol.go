package Models

import (
	"time"
	"wan-api-kol-event/Const"
)

type Kol struct {
	KolID                int64     `json:"KolID"  gorm:"primaryKey;  column:KolID"`
	UserProfileID        int64     `json:"UserProfileID" gorm:"column:UserProfileID"`
	Language             string    `json:"Language" gorm:"column:Language"`
	Education            string    `json:"Education" gorm:"column:Education"`
	ExpectedSalary       int64     `json:"ExpectedSalary" gorm:"column:ExpectedSalary"`
	ExpectedSalaryEnable bool      `json:"ExpectedSalaryEnable" gorm:"column:ExpectedSalaryEnable"`
	ChannelSettingTypeID int64     `json:"ChannelSettingTypeID" gorm:"column:ChannelSettingTypeID"`
	IDFrontURL           string    `json:"IDFrontURL" gorm:"column:IDFrontURL"`
	IDBackURL            string    `json:"IDBackURL" gorm:"column:IDBackURL"`
	PortraitURL          string    `json:"PortraitURL" gorm:"column:PortraitURL"`
	RewardID             int64     `json:"RewardID" gorm:"column:RewardID"`
	PaymentMethodID      int64     `json:"PaymentMethodID" gorm:"column:PaymentMethodID"`
	TestimonialsID       int64     `json:"TestimonialsID" gorm:"column:TestimonialsID"`
	VerificationStatus   bool      `json:"VerificationStatus" gorm:"column:VerificationStatus"`
	Enabled              bool      `json:"Enabled" gorm:"column:Enabled"`
	ActiveDate           time.Time `json:"ActiveDate" gorm:"column:ActiveDate"`
	Active               bool      `json:"Active" gorm:"column:Active"`
	CreatedBy            string    `json:"CreatedBy" gorm:"column:CreatedBy"`
	CreatedDate          time.Time `json:"CreatedDate" gorm:"column:CreatedDate"`
	ModifiedBy           string    `json:"ModifiedBy" gorm:"column:ModifiedBy"`
	ModifiedDate         time.Time `json:"ModifiedDate" gorm:"column:ModifiedDate"`
	IsRemove             bool      `json:"IsRemove" gorm:"column:IsRemove"`
	IsOnBoarding         bool      `json:"IsOnBoarding" gorm:"column:IsOnBoarding"`
	Code                 string    `json:"Code" gorm:"column:Code"`
	PortraitRightURL     string    `json:"PortraitRightURL" gorm:"column:PortraitRightURL"`
	PortraitLeftURL      string    `json:"PortraitLeftURL" gorm:"column:PortraitLeftURL"`
	LivenessStatus       bool      `json:"LivenessStatus" gorm:"column:LivenessStatus"`
}

func (Kol) TableName() string {
	return Const.TABLE_KOL
}
