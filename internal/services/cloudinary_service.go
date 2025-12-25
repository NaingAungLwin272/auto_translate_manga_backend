package services

import (
	"auto_translate_manga_backend/internal/utils"
	"context"
	"fmt"
	"io"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryServcie struct {
	cloudinary *cloudinary.Cloudinary
}

func NewCloudinaryService(cloudinary *cloudinary.Cloudinary) *CloudinaryServcie {
	return &CloudinaryServcie{cloudinary: cloudinary}
}

func (cloudinaryService *CloudinaryServcie) UploadImage(ctx context.Context, reader io.Reader) (string, error) {
	cld := utils.GetCloudinaryClient()

	resp, err := cld.Upload.Upload(ctx, reader, uploader.UploadParams{
		Folder: "manga_covers", // optional folder name
	})
	if err != nil {
		return "", fmt.Errorf("upload failed: %v", err)
	}

	return resp.SecureURL, nil
}

func (cloudinarySerrvice *CloudinaryServcie) UploadPagesForChapter(ctx context.Context, reader io.Reader, mangaName string, chapterNumber int, pages int) (string, error) {
	cld := utils.GetCloudinaryClient()

	folderPath := fmt.Sprintf("%s/%d/pages", mangaName, chapterNumber)

	resp, err := cld.Upload.Upload(ctx, reader, uploader.UploadParams{
		Folder:   folderPath,
		PublicID: fmt.Sprintf("page_%d", pages),
	})

	if err != nil {
		return "", fmt.Errorf("upload failed: %v", err)
	}

	return resp.SecureURL, nil
}
