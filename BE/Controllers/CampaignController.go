package Controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"wan-api-kol-event/Const"
	"wan-api-kol-event/Initializers"
	"wan-api-kol-event/Logic"
	"wan-api-kol-event/Models"
	"wan-api-kol-event/ViewModels"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetKolsController(context *gin.Context) {
	var KolsVM ViewModels.KolViewModel
	var guid = uuid.New().String()

	// * Get Kols from the database based on the range of pageIndex and pageSize
	// * TODO: Implement the logic to get parameters from the request
	// ? If parameter passed in the request is not valid, return the response with HTTP Status Bad Request (400)
	// @params: pageIndex
	pageIndex := context.DefaultQuery("pageIndex", "1")
	// @params: pageSize
	pageSize := context.DefaultQuery("pageSize", "10")


	// * Perform Logic Here

	// Convert to int
	pageIndexInt, err := strconv.Atoi(pageIndex)
	if err != nil || pageIndexInt <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pageIndex"})
		return
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pageSize"})
		return
	}

	// ! Pass the parameters to the Logic Layer
	kols, error := Logic.GetKolLogic(pageIndexInt, pageSizeInt)
	if error != nil {
		KolsVM.Result = Const.UnSuccess
		KolsVM.ErrorMessage = error.Error()
		KolsVM.PageIndex = int64(pageIndexInt) // * change this to the actual page index from the request
		KolsVM.PageSize = int64(pageSizeInt) // * change this to the actual page size from the request
		KolsVM.Guid = guid
		context.JSON(http.StatusInternalServerError, KolsVM)
		return
	}

	// * Return the response after the logic is executed
	// ? If the logic is successful, return the response with HTTP Status OK (200)
	KolsVM.Result = Const.Success
	KolsVM.ErrorMessage = ""
	KolsVM.PageIndex = int64(pageIndexInt) // * change this to the actual page index from the request
	KolsVM.PageSize = int64(pageSizeInt) // * change this to the actual page size from the request
	KolsVM.Guid = guid
	KolsVM.KOL = kols
	KolsVM.TotalCount = int64(len(kols))
	context.JSON(http.StatusOK, KolsVM)
}

func GenerateDummyData(context *gin.Context) {
	// Get the count from the request body (as a query parameter or in the body)
	count := context.DefaultQuery("count", "10") // Default to 10 if not provided

	// Convert the count to integer
	countInt, err := strconv.Atoi(count)
	if err != nil || countInt <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid count"})
		return
	}

	// Generate dummy data based on the count
	dummyData := Logic.GenerateDummyData(countInt)

	// Prepare a slice for the Kol model to be inserted into the database
	var kolModels []Models.Kol
	for _, kol := range dummyData {
		kolModel := Models.Kol{
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
		// Append kolModel to the slice
		kolModels = append(kolModels, kolModel)
	}

	// Insert the data into the database in a single batch (bulk insert)
	if err := Initializers.DB.Create(&kolModels).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to insert data into the database: %v", err)})
		return
	}

	// Return success response
	context.JSON(http.StatusOK, gin.H{
		"result":  "Success",
		"message": fmt.Sprintf("%d dummy data inserted successfully", countInt),
		"data":    dummyData,
	})
}
