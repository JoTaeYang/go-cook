package cache

import (
	"context"
	"fmt"
	"log"
	"strings"
)

const prefix = "Bearer "

/*
Authorization 정보에서 세션 키 생성
@params auth Authorization 데이터
*/
func MakeSessionToken(auth string) (token string) {
	if strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}

	return
}

/*
세션 키로 UID 조회
@params sessionKey 세션키 정보
*/
func GetUid(sessionKey string) (string, error) {
	key := fmt.Sprintf("%s:%s:session", appName, sessionKey)

	uid, err := WriteRedisClient.Get(context.Background(), key).Result()
	if err != nil {
		log.Println(err, " ", sessionKey)
		return "", err
	}

	return uid, nil
}
