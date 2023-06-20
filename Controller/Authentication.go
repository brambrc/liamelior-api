package Controller

import (
	"liamelior-api/Helper"
	"liamelior-api/Model"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"io/ioutil"
	"os"
	"io"
	"net/url"
	"encoding/json"
	"fmt"
	"encoding/base64"
)

func Register(context *gin.Context) {
	var input Model.AuthenticationInput
    if err := context.ShouldBind(&input); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	
	avatarFileName, err := UploadPhoto(context, "avatar")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload avatar"})
		return
	}


	user := Model.User{
		Username: input.Username,
		Password: input.Password,
		Email: input.Email,
		Name: input.Name,
		Role: input.Role,
		Avatar: avatarFileName,
	}

	_, err = user.Save()

	if err != nil {
		log.Fatal("Error saving user to database", err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func UploadPhoto(context *gin.Context, fieldName string) (string, error){
	file, err := context.FormFile(fieldName)
	if err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	tempFile, err := ioutil.TempFile("", "upload-*.jpg")
	if err != nil {
		return "", err
	}
	defer os.Remove(tempFile.Name())

	dst, err := os.Create(tempFile.Name())
	if err != nil {
		return "", err
	}

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	fileBytes, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		return "", err
	}

	// Encode the file bytes to base64
	encodedFile := base64.StdEncoding.EncodeToString(fileBytes)

	imgbbKey := os.Getenv("IMGBB_KEY")
	client := &http.Client{}
	formData := url.Values{}
	formData.Set("key", imgbbKey)
	formData.Set("image", encodedFile)

	url := os.Getenv("IMGBB_URL_UPLOAD")

	res, err := client.PostForm(url, formData)
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

	return "", fmt.Errorf("failed to upload photo")

}

func Login(context *gin.Context) {
    var input Model.AuthenticationInput

    if err := context.ShouldBindJSON(&input); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := Model.FindUserByUsername(input.Username)

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err = user.ValidatePassword(input.Password)

    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    jwt, err := helper.GenerateJWT(user)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}


