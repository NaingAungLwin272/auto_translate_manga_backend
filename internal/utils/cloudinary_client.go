package utils

import (
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

func GetCloudinaryClient() *cloudinary.Cloudinary {
	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		log.Fatalf("failed to init Cloudinary: %v", err)
	}
	return cld
}
