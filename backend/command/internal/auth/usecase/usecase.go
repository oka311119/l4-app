package usecase

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"

	"github.com/oka311119/l4-app/backend/command/internal/auth"
	"github.com/oka311119/l4-app/backend/command/internal/domain/entity"
	"github.com/oka311119/l4-app/backend/command/internal/helpers/saltgen"
	"github.com/oka311119/l4-app/backend/command/internal/helpers/uuidgen"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *entity.User `json:"user"`
}

type AuthUseCase struct {
	userRepo auth.Repository
	pepper string
	signingKey []byte
	expireDuration time.Duration
	uuidgen uuidgen.UUIDGenerator
	saltgen saltgen.SaltGenerator
}

func NewAuthUseCase(
	userRepo auth.Repository,
	pepper string,
	signingKey []byte,
	tokenTTL time.Duration,
	uuidgen uuidgen.UUIDGenerator,
	saltgen saltgen.SaltGenerator) *AuthUseCase {
	return &AuthUseCase{
		userRepo: userRepo,
		pepper: pepper,
		signingKey: signingKey,
		expireDuration: time.Second * tokenTTL,
		uuidgen: uuidgen,
		saltgen: saltgen,
	}
}

func (a *AuthUseCase) SignUp(ctx context.Context, username, password string) error {
	salt, err := a.saltgen.Generate()
	if err != nil {
		return auth.ErrFailedSaltGeneration
	}

	id := a.uuidgen.V4()
	
	pwd := sha256.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(salt))
	pwd.Write([]byte(a.pepper))
	
	user := entity.NewUser(
		id,
		username,
		fmt.Sprintf("%x", pwd.Sum(nil)),
		salt,
	)

	return a.userRepo.CreateUser(ctx, user)
}

func (a *AuthUseCase) SignIn(ctx context.Context, username, password string) (string, error) {
	user, err := a.userRepo.GetUser(ctx, username)
	if err != nil {
		return "", auth.ErrUserNotFound
	}

	// パスワード検証
	pwd := sha256.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(user.Salt))
	pwd.Write([]byte(a.pepper))

	if user.Password != fmt.Sprintf("%x", pwd.Sum(nil)) {
		return "", auth.ErrInvalidAccessToken
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.signingKey)
}

func (a *AuthUseCase) ParseToken(ctx context.Context, accessToken string) (*entity.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return a.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, auth.ErrInvalidAccessToken
}
