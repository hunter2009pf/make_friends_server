package handlers

import (
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"angel_clothes.make_friends/m/v2/constants"
	"angel_clothes.make_friends/m/v2/models"
	"angel_clothes.make_friends/m/v2/sql"

	"github.com/chai2010/webp"
	"github.com/gin-gonic/gin"
)

// 上传个人照片
func UploadPhoto(c *gin.Context) {
	idStr := c.PostForm("id")
	id, _ := strconv.Atoi(idStr)
	id2 := uint(id)
	fmt.Printf("id is %s\n", idStr)
	file, header, _ := c.Request.FormFile("photo")
	fileName := header.Filename
	var lastDotIdx int = strings.LastIndex(fileName, ".")

	// 如果是webp格式，直接保存
	var ext string = fileName[lastDotIdx+1:]
	if strings.ToLower(ext) == "webp" {
		filePath := path.Join(constants.PHOTO_SAVED_DIR, fileName)
		if err := c.SaveUploadedFile(header, filePath); err != nil {
			fmt.Println((err.Error()))
			c.JSON(http.StatusOK, gin.H{
				"code":    constants.ERROR_COPY_FILE_FAILED,
				"message": "copy file failed",
			})
		} else {
			// 数据库更新图片访问地址
			var photoUrl string = "/images/" + fileName
			sql.UpdatePhotoUrl(photoUrl, id2)
			c.JSON(http.StatusOK, gin.H{
				"code":    constants.SUCCESS,
				"message": "upload photo successfully",
				"data":    photoUrl,
			})
		}
		return
	}

	fileName = fileName[0:lastDotIdx] + ".webp"
	fmt.Println("file name is " + fileName)

	// decode normal image
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    constants.ERROR_DECODE_FILE_FAILED,
			"message": "decode file failed",
		})
		return
	}

	// Encode lossless webp
	result, err := webp.EncodeLosslessRGB(img)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    constants.ERROR_ENCODE_FILE_FAILED,
			"message": "encode file failed",
		})
		return
	}

	filePath := path.Join(constants.PHOTO_SAVED_DIR, fileName)
	if err = os.WriteFile(filePath, result, 0666); err != nil {
		fmt.Println((err.Error()))
		c.JSON(http.StatusOK, gin.H{
			"code":    constants.ERROR_COPY_FILE_FAILED,
			"message": "copy file failed",
		})
	} else {
		// 数据库更新图片访问地址
		var photoUrl string = "/images/" + fileName
		sql.UpdatePhotoUrl(photoUrl, id2)
		c.JSON(http.StatusOK, gin.H{
			"code":    constants.SUCCESS,
			"message": "upload photo successfully",
			"data":    photoUrl,
		})
	}
}

func UploadPhotoToList(c *gin.Context) {
	idStr := c.PostForm("id")
	id, _ := strconv.Atoi(idStr)
	id2 := uint(id)
	fmt.Printf("id is %s\n", idStr)
	file, header, _ := c.Request.FormFile("photo")
	fileName := header.Filename
	var lastDotIdx int = strings.LastIndex(fileName, ".")

	// 如果是webp格式，直接保存
	var ext string = fileName[lastDotIdx+1:]
	if strings.ToLower(ext) == "webp" {
		filePath := path.Join(constants.PHOTO_SAVED_DIR, fileName)
		if err := c.SaveUploadedFile(header, filePath); err != nil {
			fmt.Println((err.Error()))
			c.JSON(http.StatusOK, gin.H{
				"code":    constants.ERROR_COPY_FILE_FAILED,
				"message": "copy file failed",
			})
		} else {
			// 数据库更新图片访问地址
			var photoUrl string = "/images/" + fileName
			sql.UpdatePhotoUrlToList(photoUrl, id2)
			c.JSON(http.StatusOK, gin.H{
				"code":    constants.SUCCESS,
				"message": "upload photo successfully",
				"data":    photoUrl,
			})
		}
		return
	}

	fileName = fileName[0:lastDotIdx] + ".webp"
	fmt.Println("file name is " + fileName)

	// decode normal image
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    constants.ERROR_DECODE_FILE_FAILED,
			"message": "decode file failed",
		})
		return
	}

	// Encode lossless webp
	result, err := webp.EncodeLosslessRGB(img)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    constants.ERROR_ENCODE_FILE_FAILED,
			"message": "encode file failed",
		})
		return
	}

	filePath := path.Join(constants.PHOTO_SAVED_DIR, fileName)
	if err = os.WriteFile(filePath, result, 0666); err != nil {
		fmt.Println((err.Error()))
		c.JSON(http.StatusOK, gin.H{
			"code":    constants.ERROR_COPY_FILE_FAILED,
			"message": "copy file failed",
		})
	} else {
		// 数据库更新图片访问地址
		var photoUrl string = "/images/" + fileName
		sql.UpdatePhotoUrlToList(photoUrl, id2)
		c.JSON(http.StatusOK, gin.H{
			"code":    constants.SUCCESS,
			"message": "upload photo successfully",
			"data":    photoUrl,
		})
	}
}

// 删除个人生活照
func DeletePhotoFromList(c *gin.Context) {
	idStr := c.Query("id")
	id, _ := strconv.Atoi(idStr)
	id2 := uint(id)
	photoUrl := c.Query("photoUrl")
	var isOk bool = sql.DeletePhotoUrlFromList(photoUrl, id2)
	if isOk {
		c.JSON(http.StatusOK, gin.H{
			"code":    constants.SUCCESS,
			"message": "delete photo successfully",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    constants.ERROR_DELETE_PHOTO_FAILED,
			"message": "fail to delete photo",
		})
	}
}

// 上传个人信息
func UploadInfo(c *gin.Context) {
	info := c.PostForm("info")
	if info == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    constants.ERROR_PERSONAL_INFO_EMPTY,
			"message": "personal info should not be empty",
		})
	} else {
		var person models.Person
		json.Unmarshal([]byte(info), &person)
		if person.Id == 0 {
			c.JSON(http.StatusOK, gin.H{
				"code":    constants.ERROR_PERSONAL_INFO_EMPTY,
				"message": "personal info should not be empty",
			})
			return
		}
		var isOk bool = sql.UpdatePersonalInfo(&person)
		if isOk {
			c.JSON(http.StatusOK, gin.H{
				"code":    constants.SUCCESS,
				"message": "update personal info successfully",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    constants.ERROR_UPDATE_INFO_FAILED,
				"message": "personal info update failed",
			})
		}
	}
}

// 随机抽取1个人的数据返回
func RandomGetInfo(c *gin.Context) {
	var lookingForGirlStr string = c.Query("lookingForGirl")
	fmt.Printf("lookingForGirl is %s\n", lookingForGirlStr)
	var lookingForGirl bool = (lookingForGirlStr == "true")

	idStr := c.Query("id")
	id, _ := strconv.Atoi(idStr)
	id2 := uint(id)

	previousIdStr := c.Query("previousId")
	var previousId2 uint = ^uint(0)
	if previousIdStr != "" {
		previousId, _ := strconv.Atoi(previousIdStr)
		previousId2 = uint(previousId)
	}

	var randomPerson *models.Person = sql.RandomFindPerson2(lookingForGirl, id2, previousId2)
	if randomPerson == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    constants.ERROR_RANDOM_FIND_FAILED,
			"message": "random find person failed",
		})
	} else {
		jsonPerson, _ := json.Marshal(randomPerson)
		c.JSON(http.StatusOK, gin.H{
			"code":    constants.SUCCESS,
			"message": "ok",
			"data":    string(jsonPerson),
		})
	}
}

// 获取微信小程序登录信息
func OnLoginHandler(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    constants.ERROR_LOGIN_CODE_EMPTY,
			"message": "login code should not be empty",
		})
	} else {
		log.Println("code is " + code)
		appID := "wxef8d8c6fbf9db177"
		secret := "b8d357ce3494bbc7bec1a66020e43711"
		jsCode := code
		grantType := "authorization_code"

		url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=%s", appID, secret, jsCode, grantType)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error making GET request:", err)
			return
		}
		defer resp.Body.Close()

		// Process response here
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}
		fmt.Printf("body is %s\n", string(body))

		// 从数据库查询这个账号
		var jsonMap map[string]interface{}
		json.Unmarshal(body, &jsonMap)
		openid := fmt.Sprint(jsonMap["openid"])
		var userPtr *models.User = sql.FindUserByOpenid(openid)
		if userPtr == nil {
			sessionKey := fmt.Sprint(jsonMap["session_key"])
			// idolList := []uint{1, 2, 3}
			// sql.Create(&models.User{SessionKey: sessionKey, Openid: openid, IdolList: idolList})
			sql.Create(&models.User{SessionKey: sessionKey, Openid: openid})
			var user *models.User = sql.FindUserByOpenid(openid)
			randomGeneratedName := constants.DEFAULT_NAME_PREFIX[int(user.Id)%len(constants.DEFAULT_NAME_PREFIX)] + fmt.Sprint(user.Id)
			sql.CreatePerson(&models.Person{Id: user.Id, Name: randomGeneratedName, Photos: []string{}})
			var person *models.Person = sql.FindPersonById(user.Id)
			var dataMap map[string]interface{}
			dataMap = make(map[string]interface{})
			dataMap["user"] = *user
			dataMap["person"] = *person
			dataJson, _ := json.Marshal(dataMap)
			c.JSON(http.StatusOK, gin.H{
				"code":    constants.SUCCESS,
				"message": "success",
				"data":    string(dataJson),
			})
		} else {
			// 用户已存在
			var person *models.Person = sql.FindPersonById(userPtr.Id)
			var dataMap map[string]interface{}
			dataMap = make(map[string]interface{})
			dataMap["user"] = *userPtr
			dataMap["person"] = *person
			dataJson, _ := json.Marshal(dataMap)
			c.JSON(http.StatusOK, gin.H{
				"code":    constants.SUCCESS,
				"message": "success",
				"data":    string(dataJson),
			})
		}
	}
}

// 根据带过来的id值，向增大方向查询10个异性有照片的，如果查出来不足10，返回id为1的起始点再查差额有照片的人，如果仍然凑不够，直接返还
func RandomGetPhotos(c *gin.Context) {
	idStr := c.Query("start_id")
	id, _ := strconv.Atoi(idStr)
	startId := uint(id)
	isBoyStr := c.Query("is_boy")
	var isBoy bool
	if isBoyStr == "true" {
		fmt.Print("look for girls!")
		isBoy = true
	} else {
		fmt.Print("look for boys!")
		isBoy = false
	}
	cntStr := c.Query("count")
	cnt, _ := strconv.Atoi(cntStr)
	var peopleList []models.Person = sql.RandomFindPeople(startId, isBoy, cnt)
	if peopleList == nil || len(peopleList) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    constants.ERROR_RANDOM_FIND_FAILED,
			"message": "random find people failed",
		})
	} else {
		jsonPeople, _ := json.Marshal(peopleList)
		c.JSON(http.StatusOK, gin.H{
			"code":    constants.SUCCESS,
			"message": "ok",
			"data":    string(jsonPeople),
		})
	}
}
