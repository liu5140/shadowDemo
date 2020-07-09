package service

import (
	"errors"
	"fmt"
	"shadowDemo/model"
	"strconv"
	"time"

	"shadowDemo/zframework/credis"
	modelc "shadowDemo/zframework/model"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
)

type SessionService struct {
}

var sessionService *SessionService

func NewSessionService() *SessionService {
	if sessionService == nil {
		l.Lock()
		if sessionService == nil {
			sessionService = &SessionService{}
		}
		l.Unlock()
	}
	return sessionService
}

//DeleteSessionByUserID 根据用户ID删除session
func (service SessionService) DeleteSessionByUserID(userID int64) error {
	conn := credis.RedisManagerInstance(credis.REDIS_MANAGER).Client().Get()
	defer conn.Close()

	//删除session
	reply, err := redis.Strings(conn.Do("KEYS", fmt.Sprintf("session:%d:*", userID)))
	if err != nil {
		Log.Error(err)
		return err
	}
	if reply != nil && len(reply) > 0 {
		key := reply[0]
		conn.Do("DEL", key)
	}
	return nil
}

//DeleteSessionByToken 根据令牌删除token
func (service SessionService) DeleteSessionByToken(token string) error {
	conn := credis.RedisManagerInstance(credis.REDIS_MANAGER).Client().Get()
	defer conn.Close()

	reply, err := redis.Strings(conn.Do("KEYS", fmt.Sprintf("session:*:*:*:*:*:*:%s:", token)))
	if err != nil {
		Log.Error(err)
		return err
	}
	if reply != nil && len(reply) > 0 {
		key := reply[0]
		conn.Do("DEL", key)
	}
	return nil
}

//Createsession 创建session
func (service SessionService) CreateSession(userID int64, userType model.UserType, ip string, devInfo modelc.Device, loginURL string, token string) error {
	conn := credis.RedisManagerInstance(credis.REDIS_MANAGER).Client().Get()
	defer conn.Close()
	// geoipService := NewGeoipService()
	// addr := geoipService.SearchCityByip(ip).Address + geoipService.SearchCityByip(ip).NodeAddress
	key := fmt.Sprintf("session:%s:%d:%s:%d:%d:%s:%s:", fmt.Sprint(userID), userType, ip, 0, devInfo, loginURL, token)
	err := conn.Send("HMSET", key, "userID", fmt.Sprint(userID), "userType", userType, "IP", ip, "deviceInfo", devInfo, "loginURL", loginURL)
	if err != nil {
		Log.Errorln(err)
		return err
	}
	err = conn.Send("EXPIRE", key, 36000)
	if err != nil {
		Log.Errorln(err)
		return err
	}
	err = conn.Flush()
	if err != nil {
		Log.Errorln(err)
		return err
	}

	return nil
}

//UpdateSessionContent 更新session内容
func (service SessionService) UpdateSessionContent(userID string, hkey string, hvalue string) error {
	conn := credis.RedisManagerInstance(credis.REDIS_MANAGER).Client().Get()
	defer conn.Close()

	reply, err := redis.Strings(conn.Do("KEYS", fmt.Sprintf("session:%s:*:*:*:*:*:*:", userID)))
	if err != nil {
		Log.Error(err)
		return err
	}

	if reply != nil && len(reply) > 0 {
		key := reply[0]
		//更新session内容
		conn.Send("HSET", key, hkey, hvalue)
		conn.Send("EXPIRE", key, 36000)
		conn.Flush()
	}
	return nil
}

//GetSessionContent 获取session内容
func (service SessionService) GetSessionContent(userID string, hkey string) (value string, err error) {
	conn := credis.RedisManagerInstance(credis.REDIS_MANAGER).Client().Get()
	defer conn.Close()

	reply, err := redis.Strings(conn.Do("KEYS", fmt.Sprintf("session:%s:*:*:*:*:*:*:", userID)))
	if err != nil {
		Log.Error(err)
		return
	}

	if reply != nil && len(reply) > 0 {
		key := reply[0]
		//更新session内容
		value, err := redis.String(conn.Do("HGET", key, hkey))
		if err != nil {
			Log.Error(err)
			return value, err
		}
		return value, nil
	}
	return value, nil
}

//GetSessionContent 获取session内容
func (service SessionService) GetAllSessionContent(userID string) (hash map[string]string, err error) {
	conn := credis.RedisManagerInstance(credis.REDIS_MANAGER).Client().Get()
	defer conn.Close()

	reply, err := redis.Strings(conn.Do("KEYS", fmt.Sprintf("session:%s:*:*:*:*:*:*:", userID)))
	if err != nil {
		Log.Error(err)
		return
	}

	if reply != nil && len(reply) > 0 {
		key := reply[0]
		//更新session内容
		hash, err = redis.StringMap(conn.Do("HGETALL", key))
		if err != nil {
			Log.Error(err)
		}
	}
	return hash, nil
}

//GetSessionContent 获取session内容
func (service SessionService) GetAllSessionContentByUserid(userID string) (hash map[string]string, err error) {
	conn := credis.RedisManagerInstance(credis.REDIS_MANAGER).Client().Get()
	defer conn.Close()

	reply, err := redis.Strings(conn.Do("KEYS", fmt.Sprintf("session:*:%s:*:*:*:*:*:*:", userID)))
	if err != nil {
		Log.Error(err)
		return
	}

	if reply != nil && len(reply) > 0 {
		key := reply[0]
		//更新session内容
		hash, err = redis.StringMap(conn.Do("HGETALL", key))
		if err != nil {
			Log.Error(err)
		}
	}
	return hash, nil
}

//CreateJWT 创建登陆令牌
func (service SessionService) CreateJWT(userID string, account string, hmacSampleSecret string) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"PID":     userID,
		"Account": account,
		"Time":    time.Now(),
	})
	tokenString, err = token.SignedString([]byte(hmacSampleSecret))
	if err != nil {
		Log.Error(err)
	}
	return
}

func (service SessionService) ParseToken(tokenss string, hmacSampleSecret string) (user int64, err error) {
	token, err := jwt.Parse(tokenss, func(token *jwt.Token) (interface{}, error) {
		return []byte(hmacSampleSecret), nil
	})
	if err != nil {
		return
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("cannot convert claim to mapclaim")
		return
	}
	//验证token，如果token被修改过则为false
	if !token.Valid {
		err = errors.New("token is invalid")
		return
	}
	users := claim["PID"].(string)
	//user.Username = claim["username"].(string)
	user, err = strconv.ParseInt(users, 10, 64)
	if err != nil {
		return
	}
	return
}

//CreateSecureJWT 创建二级密码令牌, 此令牌有效期为5分钟
func (service SessionService) CreateSecureJWT(userID string, account string, hmacSampleSecret string) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"PID":     userID,
		"Account": account,
		"Time":    time.Now(),
		"exp":     time.Now().Add(5 * time.Minute).Unix(),
	})
	tokenString, err = token.SignedString([]byte(hmacSampleSecret))
	if err != nil {
		Log.Error(err)
	}
	return
}

//SetSessionExpireTime 延期session过期时间
func (service SessionService) SetSessionExpireTime(userID string, token string, maxLifeTime int64) error {
	conn := credis.RedisManagerInstance(credis.REDIS_MANAGER).Client().Get()
	defer conn.Close()
	key := fmt.Sprintf("session:%s:*:*:*:*:*:%s:", userID, token)
	reply, err := redis.Strings(conn.Do("KEYS", key))
	if err != nil {
		Log.Error(err)
		return err
	}

	if reply != nil && len(reply) > 0 {
		key := reply[0]
		//更新session过期时间
		conn.Do("EXPIRE", key, maxLifeTime)
		return nil
	}
	err = errors.New("token expired")
	Log.Error(err)
	return err
}

//IsSessionValid 延期session过期时间
func (service SessionService) IsSessionValid(userID int64) bool {
	conn := credis.RedisManagerInstance(credis.REDIS_MANAGER).Client().Get()
	defer conn.Close()

	reply, err := redis.Strings(conn.Do("KEYS", fmt.Sprintf("session:%d:*:*:*:*:*:*:", userID)))
	if err != nil {
		Log.Error(err)
		return false
	}

	if reply != nil && len(reply) > 0 {
		return true
	}
	return false
}
