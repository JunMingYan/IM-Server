package service

import (
	"errors"
	"fmt"
	"github.com/kayon/iploc"
	"golang.org/x/crypto/bcrypt"
	"server/constant"
	"server/internal/models"
	"server/internal/mysql"
	"server/pkg/sonyflake"
	"time"
)

const secter = "JULIA"

// 1.检查参数(昵称)是否重复
// 2.检查密码是否相同
// 3.对密码进行加密
// 4.注册用户
func Register(user models.UserRegisterForm) error {
	//list, err := mongodb.UserList(bson.M{"userName": user.UserName})
	result, err := mysql.CheckExistByUserName(user.UserName)
	if result != (models.User{}) {
		return errors.New(constant.UserNameError)
	}

	if user.Password != user.Confirm_Password {
		return errors.New(constant.PasswordError)
	}
	salt, err := encryptPassword([]byte(user.Password))
	if err != nil {
		return errors.New(constant.SaltError)
	}

	userID, _ := sonyflake.GetUserID()
	from := models.GetUser()
	from.UserName = user.UserName
	from.UserID = userID
	from.Salt = string(salt)
	from.Avatar = user.Avatar

	return mysql.SaveUser(from)
}

// 1.检查用户是否存在(根据用户名查询)
// 2.检查密码是否正确
func Login(user models.UserLoginForm) (*models.User, map[string][]uint64, map[uint64]string, error) {
	//list, err := mongodb.UserList(bson.M{"userName": user.UserName})
	result, err := mysql.CheckExistByUserName(user.UserName)
	if err != nil {
		return nil, nil, nil, err
	}
	if result == (models.User{}) {
		return nil, nil, nil, errors.New(constant.NotUserNameError)
	}

	if !comparePasswords([]byte(result.Salt), []byte(user.Password)) {
		return nil, nil, nil, errors.New(constant.PasswordError)
	}

	list1, err := findFriendGroup(result.UserID)
	if err != nil {
		return nil, nil, nil, err
	}

	list2, err := findFriendNote(result.UserID)
	if err != nil {
		return nil, nil, nil, err
	}

	return &result, list1, list2, nil
}

func GetUserInfo(userID uint64) (models.User, error) {
	return mysql.UsrInfo(userID)
}

//
func encryptPassword(data []byte) (salt []byte, err error) {
	return bcrypt.GenerateFromPassword(data, bcrypt.DefaultCost)
}

//
func comparePasswords(pwd1, pwd2 []byte) bool {
	err := bcrypt.CompareHashAndPassword(pwd1, pwd2)
	return err == nil
}

//
func GetDate(now time.Time) time.Time {
	layout := "2006-01-02"
	t := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())
	timeStr := t.Format(layout)
	T, _ := time.Parse(layout, timeStr)
	return T
}

func SearchUser(user models.SearchUser) ([]models.User, error) {
	return mysql.SearchUser(user)
}

func UpdatePwd(user models.UpdateUserPwd) error {
	// 1.根据用户ID查询用户
	result, err := mysql.UsrInfo(user.UserID)
	if err != nil {
		return err
	}
	if result == (models.User{}) {
		return errors.New("用户不存在")
	}
	// 2.验证旧密码是否正确
	if !comparePasswords([]byte(result.Salt), []byte(user.OldPwd)) {
		return errors.New(constant.PasswordError)
	}
	// 3.检查新密码和旧密码是否相同
	if comparePasswords([]byte(result.Salt), []byte(user.NewPwd)) {
		return errors.New(constant.UpdatePasswordError)
	}
	// 4.检查新密码和确认密码是否相同
	if user.NewPwd != user.ReNewPwd {
		return errors.New("新密码和确认密码不相同,请重新输入")
	}
	// 5.更新密码
	salt, err := encryptPassword([]byte(user.NewPwd))
	if err != nil {
		return err
	}
	err = mysql.UpdatePwd(user.UserID, string(salt))
	if err != nil {
		return err
	}

	return nil
}

func UpdateUserInfo(user models.UpdateUserInfo) error {
	// 1.根据用户ID查询用户
	result, err := mysql.UsrInfo(user.UserID)
	if err != nil {
		return err
	}
	if result == (models.User{}) {
		return errors.New("用户不存在")
	}

	return mysql.UpdateUserInfo(user)
}

func UpdateUserConfigure(user models.UpdateUserConfigure) error {
	// 1.根据用户ID查询用户
	result, err := mysql.UsrInfo(user.UserID)
	if err != nil {
		return err
	}
	if result == (models.User{}) {
		return errors.New("用户不存在")
	}

	return mysql.UpdateUserConfigure(user)
}

func BrowserSetting() error {
	ip := "120.230.94.77"
	loc, err := iploc.Open("D:\\Software\\go\\workspace\\chat\\server\\resources\\data.dat")
	if err != nil {
		return err
	}
	detail := loc.Find(ip)
	fmt.Println(detail.Country, detail.Province, detail.City, detail.County)

	return nil
}

func findFriendGroup(userID uint64) (map[string][]uint64, error) {
	groups, err := mysql.FriendGroup(userID)
	if err != nil {
		return nil, err
	}

	names, err := mysql.FriendGroupName(userID)
	if err != nil {
		return nil, err
	}
	results := make(map[string][]uint64)
	for _, v1 := range names {
		var ids []uint64
		for _, v2 := range groups {
			if v1.GroupName != v2.GroupName {
				continue
			}
			ids = append(ids, v2.FriendID)
		}
		results[v1.GroupName] = ids
	}

	return results, nil
}

func findFriendNote(userID uint64) (map[uint64]string, error) {
	notes, err := mysql.FriendNotes(userID)
	if err != nil {
		return nil, err
	}
	results := make(map[uint64]string)
	for _, v := range notes {
		results[v.FriendID] = v.Notes
	}

	return results, nil
}
