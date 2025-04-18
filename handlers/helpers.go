package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"html/template"
	"irbistest/iternal/app"
	"irbistest/iternal/templates"
	"time"

	//"irbistest/postgres"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	User_id int    `json:"user_id"`
	Ip      string `json:"ip"`
	Refresh string `json:"refresh"`
	jwt.RegisteredClaims
}

type Handlers struct {
	App           *app.Application
	templateCache map[string]*template.Template
}

func NewHandlers(app *app.Application) *Handlers {
	templateCache, err := templates.NewTemplateCache("./ui/html/")
	if err != nil {
		fmt.Errorf("Ошибка", err)
	}
	return &Handlers{
		App:           app,
		templateCache: templateCache,
	}
}

func (app *Handlers) render(w http.ResponseWriter, r *http.Request, name string, td *templates.TemplateData) {

	ts, ok := app.templateCache[name]
	if !ok {
		fmt.Errorf("Шаблон %s не существует!", name)
		return
	}

	err := ts.Execute(w, td)
	if err != nil {
		fmt.Errorf("Ошибка", err)
	}
}

func (app *Handlers) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Handlers) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

var (
	signingKey = []byte("testtask")
)

func CreateT(ip string, id int) (string, string, string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		fmt.Print("Ошибка генерации рефреш токена - ", err)
		return "", "", "", err
	}
	refresh := base64.URLEncoding.EncodeToString(b)
	expirationTime := time.Now().Add(55 * time.Minute)
	claims := &Claims{
		User_id: id,
		Ip:      ip,
		Refresh: refresh,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStringAcces, err := token.SignedString(signingKey)
	if err != nil {
		fmt.Print("Ошибка при создании аксес токена- ", err)
		return "", "", "", err
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(refresh), bcrypt.DefaultCost)
	if err != nil {
		fmt.Print("Ошибка при хэшировании пароля- ", err)
		return "", "", "", err
	}

	return tokenStringAcces, refresh, string(passHash), nil
}

func (h Handlers) Message(to string) {
	fmt.Printf("Отправляю емайл на почту клиента %s с сообщением warning", to)
}
