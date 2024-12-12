package Controllers

import (
	"net/http"
	"strconv"
	"wan-api-kol-event/Const"
	"wan-api-kol-event/Logic"
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
