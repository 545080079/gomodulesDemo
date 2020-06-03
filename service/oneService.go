package service

import (
	"crypto/md5"
	"demo/controller"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Session struct {
	Client controller.Client
}

func (s *Session) Set(id int64) (token string, err error) {

	key := strconv.FormatInt(time.Now().Unix()+id, 10)
	data := []byte(key)
	has := md5.Sum(data)
	token = fmt.Sprintf("%x", has)

	err = s.Client.Set(token, strconv.FormatInt(id, 10))
	if err != nil {
		return token, err
	}
	return
}

func (s *Session) Get(token string) (id int64, err error) {
	str, err := s.Client.Get(token)
	if err != nil {
		return 0, err
	}
	id, err = strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return
}

func RunServiceOne(tp int) (err error) {
	fmt.Println("RunServiceOne: ", tp)
	if tp > 9 || tp < 0 {
		return errors.New("RunServiceOneError[1]")
	}
	return
}
