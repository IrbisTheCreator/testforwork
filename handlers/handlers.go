package handlers

import (
	"fmt"
	"irbistest/iternal/templates"
	"irbistest/models"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handlers) HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.notFound(w)
		return
	}
	h.render(w, r, "home.page.tmpl", &templates.TemplateData{})
}

func (h *Handlers) TakeTokens(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("GUID"))
	if err != nil || id < 1 {
		h.notFound(w)
		fmt.Print("Неверный GUID")
		return
	}

	var email string

	err = h.App.DB.QueryRow("SELECT email FROM users WHERE user_id = $1", id).Scan(&email)
	if err != nil {
		h.notFound(w)
		fmt.Print("пользователя с таким ID  не существует")
		return
	}

	ip := r.RemoteAddr
	tokenStringAcces, refresh, passHash, err := CreateT(ip, id)
	if err != nil {
		h.notFound(w)
		fmt.Print("Ошибка создания токенов -", err)
		return
	}
	expirationTime := time.Now().Add(55 * time.Minute)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenStringAcces,
		Expires: expirationTime,
		Path:    "/",
	})

	_, err = h.App.DB.Exec("UPDATE users SET refresh = $1 WHERE user_id = $2", passHash, id)
	if err != nil {
		fmt.Print("Ошибка при обновлении токена- ", err)
	}
	s := &models.Info{}
	s.Acces = tokenStringAcces
	s.Refresh = refresh
	s.Hash = passHash
	h.render(w, r, "acces.page.tmpl", &templates.TemplateData{
		Info: s,
	})

}

func (h *Handlers) RefreshToken(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("token")
	if err != nil {
		return
	}
	tokenStr := c.Value
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return
	}
	if !token.Valid {
		return
	}

	var hash, email string

	err = h.App.DB.QueryRow("SELECT email,refresh FROM users WHERE user_id = $1", claims.User_id).Scan(&email, &hash)
	if err != nil {
		h.notFound(w)
		fmt.Print("пользователя с таким ID  не существует")
		return
	}

	_, err = h.App.DB.Exec("UPDATE users SET refresh = '' WHERE user_id = $1", claims.User_id)
	if err != nil {
		fmt.Print("Ошибка при обновлении токена- ", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(claims.Refresh))
	if err != nil {
		h.render(w, r, "error.page.tmpl", &templates.TemplateData{})
		return
	}

	ip := r.RemoteAddr
	tokenStringAcces, refresh, passHash, err := CreateT(ip, claims.User_id)
	if err != nil {
		h.notFound(w)
		fmt.Print("Ошибка создания токенов -", err)
		return
	}
	expirationTime := time.Now().Add(55 * time.Minute)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenStringAcces,
		Expires: expirationTime,
		Path:    "/",
	})

	_, err = h.App.DB.Exec("UPDATE users SET refresh = $1 WHERE user_id = $2", passHash, claims.User_id)
	if err != nil {
		fmt.Print("Ошибка при обновлении токена- ", err)
	}
	s := &models.Info{}
	s.Acces = tokenStringAcces
	s.Refresh = refresh
	s.Hash = passHash

	var errorMessage string
	if ip != claims.Ip {
		errorMessage = "IP адрес с последнего получения токена изменился, отправляю сообщение на почту"
		h.Message(email)
	}

	h.render(w, r, "complete.page.tmpl", &templates.TemplateData{
		Info:         s,
		ErrorMessage: errorMessage,
	})

}
