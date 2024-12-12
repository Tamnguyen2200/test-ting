package Logic

import (
	"fmt"
	"time"
	"wan-api-kol-event/DTO"
	"wan-api-kol-event/Initializers" // Import Initializers package for DB
	"wan-api-kol-event/Models"

	"golang.org/x/exp/rand"
)

// * Get Kols from the database based on the range of pageIndex and pageSize
// ! USE GORM TO QUERY THE DATABASE
// ? There are some support function that can be access in Utils folder (/BE/Utils)
// --------------------------------------------------------------------------------
// @params: pageIndex
// @params: pageSize
// @return: List of KOLs and error message
func GetKolLogic(pageIndex int, pageSize int) ([]*DTO.KolDTO, error) {

	var kols []Models.Kol
	offset := (pageIndex - 1) * pageSize

	// Get data from database
	if err := Initializers.DB.Offset(offset).Limit(pageSize).Find(&kols).Error; err != nil {
		return nil, err
	}

	// If no KOLs are found, return an empty list
	if len(kols) == 0 {
		return []*DTO.KolDTO{}, nil
	}

	// Map the result to DTO
	var kolDTOs []*DTO.KolDTO
	for _, kol := range kols {
		kolDTO := &DTO.KolDTO{
			KolID:                kol.KolID,
			UserProfileID:        kol.UserProfileID,
			Language:             kol.Language,
			Education:            kol.Education,
			ExpectedSalary:       kol.ExpectedSalary,
			ExpectedSalaryEnable: kol.ExpectedSalaryEnable,
			ChannelSettingTypeID: kol.ChannelSettingTypeID,
			IDFrontURL:           kol.IDFrontURL,
			IDBackURL:            kol.IDBackURL,
			PortraitURL:          kol.PortraitURL,
			RewardID:             kol.RewardID,
			PaymentMethodID:      kol.PaymentMethodID,
			TestimonialsID:       kol.TestimonialsID,
			VerificationStatus:   kol.VerificationStatus,
			Enabled:              kol.Enabled,
			ActiveDate:           kol.ActiveDate,
			Active:               kol.Active,
			CreatedBy:            kol.CreatedBy,
			CreatedDate:          kol.CreatedDate,
			ModifiedBy:           kol.ModifiedBy,
			ModifiedDate:         kol.ModifiedDate,
			IsRemove:             kol.IsRemove,
			IsOnBoarding:         kol.IsOnBoarding,
			Code:                 kol.Code,
			PortraitRightURL:     kol.PortraitRightURL,
			PortraitLeftURL:      kol.PortraitLeftURL,
			LivenessStatus:       kol.LivenessStatus,
		}
		kolDTOs = append(kolDTOs, kolDTO)
	}

	// Return the mapped DTOs
	return kolDTOs, nil
}

// GenerateRandomKolDTO generates a random KolDTO for testing
func GenerateRandomKolDTO(index int64) DTO.KolDTO {
	rand.Seed(uint64(time.Now().UnixNano() + index))

	// Random values for Language and Education
    languages := []string{"English", "Spanish", "French", "German", "Chinese"}
    educations := []string{"Bachelor's Degree", "Master's Degree", "PhD", "High School", "Associate Degree"}

	return DTO.KolDTO{
		// KolID:                int64(index),
		UserProfileID:        int64(rand.Intn(10000) + 1),
		Language:             languages[rand.Intn(len(languages))], 
		Education:            educations[rand.Intn(len(educations))],
		ExpectedSalary:       rand.Int63n(10000) + 1000, // Expected salary between 1000 and 11000
		ExpectedSalaryEnable: rand.Intn(2) == 1,
		ChannelSettingTypeID: int64(rand.Intn(10000) + 1),
		IDFrontURL:           fmt.Sprintf("https://example.com/front/%d", index),
		IDBackURL:            fmt.Sprintf("https://example.com/back/%d", index),
		PortraitURL:          fmt.Sprintf("https://example.com/portrait/%d.jpg", index),
		RewardID:             int64(rand.Intn(10000) + 1),
		PaymentMethodID:      int64(rand.Intn(10000) + 1),
		TestimonialsID:       int64(rand.Intn(10000) + 1),
		VerificationStatus:   rand.Intn(2) == 1, // Random true/false
		Enabled:              rand.Intn(2) == 1,
		ActiveDate:           time.Now(),
		Active:               rand.Intn(2) == 1,
		CreatedBy:            "Admin",
		CreatedDate:          time.Now().Add(time.Duration(rand.Intn(365)) * time.Hour * 24),
		ModifiedBy:           "Admin",
		ModifiedDate:         time.Now(),
		IsRemove:             rand.Intn(2) == 1,
		IsOnBoarding:         rand.Intn(2) == 1,
		Code:                 fmt.Sprintf("KOL%d", index),
		PortraitRightURL:     fmt.Sprintf("https://example.com/right/%d.jpg", index),
		PortraitLeftURL:      fmt.Sprintf("https://example.com/left/%d.jpg", index),
		LivenessStatus:       rand.Intn(2) == 1,
	}
}

// GenerateDummyData generates a list of dummy KolDTO objects
func GenerateDummyData(count int) []*DTO.KolDTO {
	var kols []*DTO.KolDTO
	for i := int64(1); i <= int64(count); i++ {
		kol := GenerateRandomKolDTO(i)
		kols = append(kols, &kol)  // Append pointer to KolDTO
	}
	return kols
}