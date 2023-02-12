package src

import (
	"fmt"
	"net/http"
	"time"
)

func (Gu *Guilded) GetAuth() (Auth string, str string, err error) {
	s := time.Now()
	email, name := Gu.RandMail()
	pass := Gu.RandHex(5) + name

	data := map[string]string{
		"email":    email,
		"fullName": name + Gu.RandHex(5),
		"name":     name + Gu.RandHex(5),
		"password": pass,
	}
	payload, err := Gu.Load(data)
	Gu.Err(err)

	req, err := http.NewRequest("POST",
		"https://www.guilded.gg/api/users?type=email",
		payload,
	)
	Gu.Err(err)
	Gu.HeaderAuth(req, 64)

	resp, err := Client.Do(req)
	Gu.Err(err)
	e := time.Since(s)

	switch resp.StatusCode {
	default:
		fmt.Println("err: ", resp)
	case 200:
		defer resp.Body.Close()
		Gu.Err(err)
		for _, cookie := range resp.Cookies() {
			if cookie.Name == "hmac_signed_session" {
				hmac := cookie.Value
				fmt.Println(
					"[" + bg + "" + e.String()[:3] + "ms" + rb + "] [" + g + "✓" + r + "]	Hmac-Auth:" + hmac,
				)
				Auth = hmac

			}
		}

	case 429:
		fmt.Println("RateLimmit ", resp.StatusCode)
	}
	str = email + ":" + pass + ":"
	return Auth, str, err
}
