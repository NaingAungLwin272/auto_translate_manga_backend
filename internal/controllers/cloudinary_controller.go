package controllers

import (
	dtos "auto_translate_manga_backend/internal/dto"
	"auto_translate_manga_backend/internal/models"
	"auto_translate_manga_backend/internal/services"
	"auto_translate_manga_backend/internal/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CloudinaryController struct {
	cloudinaryService *services.CloudinaryServcie
}

func NewCloudinaryController(cloudinaryService *services.CloudinaryServcie) *CloudinaryController {
	return &CloudinaryController{cloudinaryService: cloudinaryService}
}

func (cloudinaryController *CloudinaryController) UploadCoverImageToCloudinary(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   http.StatusText(http.StatusBadRequest),
			Message: "file must be require",
		})
	}

	src, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   http.StatusText(http.StatusInternalServerError),
			Message: "failed to open file",
		})
		return
	}
	defer src.Close()

	url, err := cloudinaryController.cloudinaryService.UploadImage(ctx, src)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(201, utils.SuccessResponse[dtos.UploadImageResponse]{
		Status:  201,
		Message: "cover_image uploaded successfully",
		Data: dtos.UploadImageResponse{
			File_Url: url,
		},
	})
}

func (cloudinaryController *CloudinaryController) UploadChapterPages(ctx *gin.Context) {
	mangaTitle := ctx.PostForm("manga_title")
	chapterNumberStr := ctx.PostForm("chapter_number")

	chapterNumber, _ := strconv.Atoi(chapterNumberStr)

	form, err := ctx.MultipartForm()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Error:   http.StatusText(http.StatusInternalServerError),
			Message: "invalid form",
		})
		return
	}

	files := form.File["pages"]
	fmt.Println(files, "this is files...")
	if files == nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse{
			Status:  http.StatusBadRequest,
			Error:   http.StatusText(http.StatusBadRequest),
			Message: "pages require",
		})
		return
	}

	var uploadedData []models.ChapterDetail

	for i, file := range files {
		pageNumber := i + 1

		f, _ := file.Open()
		defer f.Close()

		url, err := cloudinaryController.cloudinaryService.UploadPagesForChapter(
			ctx,
			f,
			mangaTitle,
			chapterNumber,
			pageNumber,
		)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Error:   http.StatusText(http.StatusInternalServerError),
				Message: err.Error(),
			})
			return
		}

		uploadedData = append(uploadedData, models.ChapterDetail{
			PageNumber: int64(pageNumber),
			ImageUrl:   url,
		})
	}

	ctx.JSON(201, utils.SuccessResponse[dtos.UploadPagesResponse]{
		Status:  201,
		Message: "chapter uploaded successfully",
		Data: dtos.UploadPagesResponse{
			Pages: uploadedData,
		},
	})
}
