package DTO

import "time"

type KolDTO struct {
	KolID                int64     `json:"KolID"`
	UserProfileID        int64     `json:"UserProfileID"`
	Language             string    `json:"Language"`
	Education            string    `json:"Education"`
	ExpectedSalary       int64     `json:"ExpectedSalary"`
	ExpectedSalaryEnable bool      `json:"ExpectedSalaryEnable"`
	ChannelSettingTypeID int64     `json:"ChannelSettingTypeID"`
	IDFrontURL           string    `json:"IDFrontURL"`
	IDBackURL            string    `json:"IDBackURL"`
	PortraitURL          string    `json:"PortraitURL"`
	RewardID             int64     `json:"RewardID"`
	PaymentMethodID      int64     `json:"PaymentMethodID"`
	TestimonialsID       int64     `json:"TestimonialsID"`
	VerificationStatus   bool      `json:"VerificationStatus"`
	Enabled              bool      `json:"Enabled"`
	ActiveDate           time.Time `json:"ActiveDate"`
	Active               bool      `json:"Active"`
	CreatedBy            string    `json:"CreatedBy"`
	CreatedDate          time.Time `json:"CreatedDate"`
	ModifiedBy           string    `json:"ModifiedBy"`
	ModifiedDate         time.Time `json:"ModifiedDate"`
	IsRemove             bool      `json:"IsRemove"`
	IsOnBoarding         bool      `json:"IsOnBoarding"`
	Code                 string    `json:"Code"`
	PortraitRightURL     string    `json:"PortraitRightURL"`
	PortraitLeftURL      string    `json:"PortraitLeftURL"`
	LivenessStatus       bool      `json:"LivenessStatus"`
}
