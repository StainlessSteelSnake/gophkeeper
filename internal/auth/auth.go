package auth

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
	"sync"
)

const secretKey = "TheSecretKey" // Ключ шифрования для создания подписи токена авторизованного пользователя.

// user хранит данные об авторизованном пользователе.
type user struct {
	login        string // Имя авторизованного пользователя.
	passwordHash string // Хэш пароля авторизованного пользователя.
	tokens       []string
}

// UserAdderGetter предоставляет функции добавления нового пользователя и получения данных о существующем пользователе.
type UserAdderGetter interface {
	AddUser(context.Context, string, string) error
	GetUser(context.Context, string) (string, int, error)
}

// authentication хранит список текущих авторизованных пользователей и ссылку на хранилище пользователей
type authentication struct {
	users          sync.Map        //         map[string]*user // Список авторизованных за время работы сервиса пользователей.
	userController UserAdderGetter // Ссылка на контроллер хранилища для получения и добавления пользователей.
}

// Authenticator позволяет зарегистрировать новых пользователей, авторизовать существующих пользователей
// и идентифицировать активную сессию ранее авторизованных пользователей.
type Authenticator interface {
	Register(context.Context, string, string) (string, error)
	Login(context.Context, string, string) (string, error)
	Authenticate(context.Context, string) (string, string, error)
	Logout(context.Context, string) error
}

// NewAuthentication создаёт контроллер авторизаций пользователей используя ссылку на хранилище пользователей.
func NewAuthentication(userController UserAdderGetter) (Authenticator, error) {
	if userController == nil {
		return nil, errors.New("не задан контроллер хранилища данных")
	}

	a := authentication{userController: userController}
	return &a, nil
}

// getHash создаёт хэш по алгоритму SHA256 для переданной строки и возвращает его в шестнадцатиричном представлении.
func getHash(s string) (string, error) {
	hasher := sha256.New()

	_, err := hasher.Write([]byte(s))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

// getRandom создаёт последовательность заданной длины из случайных байт
// и возвращает её в виде строки в шестнадцатиричном представлении.
func getRandom(size int) (string, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

// getSign создаёт подпись переданной строки с использованием секретного ключа.
func getSign(s string) (string, error) {
	h := hmac.New(sha256.New, []byte(secretKey))

	_, err := h.Write([]byte(s))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

// createToken создаёт токен на основании хэша логина пользователя и случайной последовательности символов.
func (a *authentication) createToken(loginHash string) (string, error) {
	randomKey, err := getRandom(6)
	if err != nil {
		return "", err
	}

	token := loginHash + ":" + randomKey

	tokenSign, err := getSign(token)
	if err != nil {
		return "", err
	}

	token = token + tokenSign

	u, ok := a.users.Load(loginHash)
	if !ok {
		return "", errors.New("пользователь не найден")
	}

	u.(*user).tokens = append(u.(*user).tokens, token)
	return token, nil
}

// Register создаёт нового пользователя с указанным логином и паролем.
func (a *authentication) Register(ctx context.Context, login, password string) (string, error) {
	passwordHash, err := getHash(password)
	if err != nil {
		return "", err
	}

	err = a.userController.AddUser(ctx, login, passwordHash)
	if err != nil {
		return "", err
	}

	loginHash, err := getHash(login)
	if err != nil {
		return "", err
	}

	u := &user{
		login:        login,
		passwordHash: passwordHash,
		tokens:       make([]string, 0),
	}

	a.users.Store(loginHash, u)

	token, err := a.createToken(loginHash)
	if err != nil {
		return "", err
	}

	return token, err
}

// checkPassword проверяет соответствие хэша переданного пароля для указанного пользователя,
// сохранённому для него хэшу пароля.
func (a *authentication) checkPassword(ctx context.Context, login, password string) (string, string, error) {
	loginHash, err := getHash(login)
	if err != nil {
		return "", "", err
	}

	savedPasswordHash := ""
	u, ok := a.users.Load(loginHash)
	if ok {
		savedPasswordHash = u.(*user).passwordHash
	}

	if !ok {
		savedPasswordHash, _, err = a.userController.GetUser(ctx, login)
	}
	if err != nil {
		return "", "", err
	}

	passwordHash, err := getHash(password)
	if err != nil {
		return "", "", err
	}

	if savedPasswordHash != passwordHash {
		return "", "", errors.New("переданный и сохранённый пароли не совпадают")
	}

	return loginHash, passwordHash, err
}

// Login создаёт токен для переданного пользователя, если он существует и верно указан пароль.
func (a *authentication) Login(ctx context.Context, login, password string) (string, error) {
	loginHash, passwordHash, err := a.checkPassword(ctx, login, password)
	if err != nil {
		return "", err
	}

	if _, ok := a.users.Load(loginHash); !ok {
		u := &user{
			login:        login,
			passwordHash: passwordHash,
			tokens:       make([]string, 0),
		}
		a.users.Store(loginHash, u)
	}

	token, err := a.createToken(loginHash)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Authenticate проверяет переданный токен и возвращает логин пользователя, если токен для него найден.
func (a *authentication) Authenticate(ctx context.Context, t string) (string, string, error) {
	tokenParts := strings.Split(t, ":")
	if len(tokenParts) != 2 {
		return "", "", errors.New("токен авторизации передан в неправильном формате")
	}

	loginHash := tokenParts[0]
	u, ok := a.users.Load(loginHash)
	if !ok {
		return "", "", errors.New("указанный пользователь не авторизован")
	}

	var tokenIsValid bool
	for _, v := range u.(*user).tokens {
		if v == t {
			tokenIsValid = true
			break
		}
	}

	if !tokenIsValid {
		return "", "", errors.New("подпись токена авторизации пользователя не соответствует сохранённой")
	}

	return u.(*user).login, loginHash, nil
}

func (a *authentication) Logout(ctx context.Context, t string) error {
	login, loginHash, err := a.Authenticate(ctx, t)
	if err != nil {
		return err
	}

	u, ok := a.users.Load(loginHash)
	if !ok {
		return errors.New("пользователь " + login + " не авторизован")
	}

	for i, value := range u.(*user).tokens {
		if value == t {
			if i == 0 {
				u.(*user).tokens = u.(*user).tokens[1:]
			} else if i == len(u.(*user).tokens)-1 {
				u.(*user).tokens = u.(*user).tokens[:i]
			} else {
				u.(*user).tokens = append(u.(*user).tokens[:i], u.(*user).tokens[i+1:]...)
			}

			break
		}
	}

	return nil
}
