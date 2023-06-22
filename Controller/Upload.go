package Controller

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"liamelior-api/Model"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type PhotoLandingPageInput struct {
	Photo   *multipart.FileHeader `form:"photo" json:"photo" binding:"required"`
	Context string                `form:"context" json:"context" binding:"required"`
}

func ContextPhoto(context *gin.Context) {
	var input PhotoLandingPageInput

	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uploadedPhoto, err := UploadPhoto(context, "photo")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload photo", "message": err.Error()})
		return
	}

	photo := Model.Photo{
		Photo:   uploadedPhoto,
		Context: input.Context,
	}

	_, err = photo.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to save photo", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Caraousel uploaded successfully!", "photo": uploadedPhoto})
}

func UploadPhoto(context *gin.Context, fieldName string) (string, error) {
	file, err := context.FormFile(fieldName)
	if err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	tempFile, err := ioutil.TempFile("", "upload-*"+file.Filename)
	if err != nil {
		return "", err
	}
	defer os.Remove(tempFile.Name())

	dst, err := os.Create(tempFile.Name())
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	webpPath := tempFile.Name() + ".webp"
	err = convertToWebP(tempFile.Name(), webpPath)
	if err != nil {
		return "", err
	}
	defer os.Remove(webpPath)

	fileBytes, err := ioutil.ReadFile(webpPath)
	if err != nil {
		return "", err
	}

	// Encode the file bytes to base64
	encodedFile := base64.StdEncoding.EncodeToString(fileBytes)

	imgbbKey := os.Getenv("IMGBB_TOKEN")
	client := &http.Client{}
	formData := url.Values{}
	formData.Set("key", imgbbKey)
	formData.Set("image", encodedFile)

	url := os.Getenv("IMGBB_URL_UPLOAD")

	res, err := client.PostForm(url+"key="+imgbbKey, formData)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var response struct {
		Data struct {
			Image struct {
				URL string `json:"url"`
			} `json:"image"`
		} `json:"data"`
		Success bool `json:"success"`
		Status  int  `json:"status"`
	}

	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return "", err
	}

	if response.Success && response.Status == 200 {
		return response.Data.Image.URL, nil
	}

	return "", fmt.Errorf("failed to upload photo" + string(responseData))

}

func convertToWebP(inputPath, outputPath string) error {
	cmd := exec.Command("cwebp", inputPath, "-o", outputPath)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to convert image to WebP: %v", err)
	}
	return nil
}
