package sql

import (
	"fmt"
	"math/rand"
	"time"

	"angel_clothes.make_friends/m/v2/constants"
	"angel_clothes.make_friends/m/v2/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "root:123456@tcp(127.0.0.1:3306)/" + constants.DATABASE_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"

var (
	db *gorm.DB
)

func OpenDbConnection() *gorm.DB {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败！" + err.Error())
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  // 空闲连接数
	sqlDB.SetMaxOpenConns(100) // 最大连接数
	sqlDB.SetConnMaxLifetime(time.Minute)
	return db
}

// 创建用户
func Create(user *models.User) {
	if user == nil {
		panic("用户不存在！")
	} else {
		result := db.Create(user)
		if result.Error != nil && result.RowsAffected == 0 {
			panic("添加失败：" + result.Error.Error())
		}
	}
}

// 创建用户信息
func CreatePerson(person *models.Person) {
	if person == nil {
		panic("用户不存在！")
	} else {
		result := db.Create(person)
		if result.Error != nil && result.RowsAffected == 0 {
			panic("添加失败：" + result.Error.Error())
		}
	}
}

// 更新用户
// func Update(user *models.User) bool {
// 	if user == nil {
// 		panic("用户不存在！")
// 	} else {
// 		// 查最新的用户
// 		var existUser models.User
// 		result := db.Where("account = ?", *&user.Account).First(&existUser)
// 		if result.RowsAffected > 0 {
// 			existUser.Account = user.Account
// 			existUser.Password = user.Password
// 			db.Save(&existUser)
// 		} else {
// 			panic("用户不存在！")
// 		}
// 	}
// 	return true
// }

// 更新头像
func UpdatePhotoUrl(photoUrl string, id uint) bool {
	// 查最新的用户
	var existPerson models.Person
	result := db.Where("id = ?", id).First(&existPerson)
	if result.RowsAffected > 0 {
		existPerson.AvatarUrl = photoUrl
		db.Save(&existPerson)
		return true
	} else {
		fmt.Println("user does not exist")
		return false
	}
}

// 更新个人图片列表，最多6张
func UpdatePhotoUrlToList(photoUrl string, id uint) bool {
	// 查询用户
	var existPerson models.Person
	result := db.Where("id = ?", id).First(&existPerson)
	if result.RowsAffected > 0 {
		existPerson.Photos = append(existPerson.Photos, photoUrl)
		db.Save(&existPerson)
		return true
	} else {
		fmt.Println("user does not exist")
		return false
	}
}

// 从图片列表中删除一张图片链接
func DeletePhotoUrlFromList(photoUrl string, id uint) bool {
	// 查询用户
	var existPerson models.Person
	result := db.Where("id = ?", id).First(&existPerson)
	if result.RowsAffected > 0 {
		if len(existPerson.Photos) == 0 {
			fmt.Println("photo list is empty")
			return false
		} else {
			var delIdx int = -1
			for index, value := range existPerson.Photos {
				if value == photoUrl {
					delIdx = index
					break
				}
			}
			if delIdx >= 0 {
				len := len(existPerson.Photos)
				if delIdx == 0 {
					existPerson.Photos = existPerson.Photos[1:]
				} else if delIdx == len-1 {
					existPerson.Photos = existPerson.Photos[:len]
				} else {
					existPerson.Photos = append(existPerson.Photos[:delIdx], existPerson.Photos[delIdx+1:]...)
				}
				db.Save(&existPerson)
				return true
			} else {
				return false
			}
		}
	} else {
		fmt.Println("user does not exist")
		return false
	}
}

// 更新个人信息
func UpdatePersonalInfo(person *models.Person) bool {
	// 查最新的用户
	db.Save(person)
	return true
}

func FindUserByAccount(account string) *models.User {
	var user models.User
	if err := db.Where("account = ?", account).First(&user).Error; err != nil {
		return nil
	} else {
		return &user
	}
}

func FindUserByAccountAndPassword(account string, password string) *models.User {
	var user models.User
	if err := db.Where("account = ? AND password = ?", account, password).First(&user).Error; err != nil {
		return nil
	} else {
		return &user
	}
}

func FindUserByOpenid(openid string) *models.User {
	var user models.User
	if err := db.Where("openid = ?", openid).First(&user).Error; err != nil {
		return nil
	} else {
		return &user
	}
}

func FindPersonById(id uint) *models.Person {
	var person models.Person
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		return nil
	} else {
		return &person
	}
}

// deprecated. Use RandomFindPerson2 instead.
func RandomFindPerson(lookingForGirl bool, id uint) *models.Person {
	var randomPerson models.Person
	if err := db.Where("id != ? AND is_boy = ?", id, !lookingForGirl).Take(&randomPerson).Error; err != nil {
		return nil
	} else {
		return &randomPerson
	}
}

func RandomFindPerson2(lookingForGirl bool, id uint, previousId uint) *models.Person {
	var matchedPeople []models.Person
	var err error
	if previousId == (^uint(0)) {
		err = db.Where("id != ? AND is_boy = ?", id, !lookingForGirl).Limit(1000).Find(&matchedPeople).Error
	} else {
		err = db.Where("id != ? AND id != ? AND is_boy = ?", id, previousId, !lookingForGirl).Limit(1000).Find(&matchedPeople).Error
	}
	if err != nil {
		return nil
	} else {
		var len int = len(matchedPeople)
		if len == 1 {
			return &matchedPeople[0]
		} else {
			index := rand.Intn(len)
			return &matchedPeople[index]
		}
	}
}

// 随机获取指定数目的异性友人
func RandomFindPeople(startId uint, lookingForGirl bool, count int) []models.Person {
	var matchedPeople []models.Person
	var err error
	// db.Debug()打印sql语句
	// err = db.Debug().Where("id > ? AND is_boy = ? AND (avatar_url != '' OR photos != 0x5B5D)", startId, !lookingForGirl).Limit(count).Find(&matchedPeople).Error
	err = db.Debug().Where("id > ? AND is_boy = ? AND (avatar_url != '' OR photos != 0x5B5D)", startId, !lookingForGirl).Limit(count).Find(&matchedPeople).Error
	if err != nil {
		return nil
	} else {
		var len int = len(matchedPeople)
		if len == count {
			return matchedPeople
		} else {
			var complementaryPeople []models.Person
			// db.Debug()打印sql语句
			// err = db.Debug().Where("id > ? AND is_boy = ? AND (avatar_url != '' OR photos != 0x5B5D)", 0, !lookingForGirl).Limit(count - len).Find(&complementaryPeople).Error
			err = db.Debug().Where("id > ? AND is_boy = ? AND (avatar_url != '' OR photos != 0x5B5D)", 0, !lookingForGirl).Limit(count - len).Find(&complementaryPeople).Error
			if err != nil {
				return matchedPeople
			} else {
				return append(matchedPeople, complementaryPeople...)
			}
		}
	}
}

// 获取用户
func Get(id uint) *models.User {
	var user models.User
	db.First(&user, id)
	return &user
}

// 删除用户
func Delete(id uint) bool {
	db.Delete(&models.User{}, id)
	return true
}

// 获取所有用户
func GetUserAll() *[]models.User {
	var users []models.User
	result := db.Select([]string{}).Find(&users)
	if result.RowsAffected > 0 {
		return &users
	}
	return nil
}
