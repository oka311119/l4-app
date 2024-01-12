package usecase

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"

	"github.com/oka311119/l4-app/backend/command/internal/area"
	"github.com/oka311119/l4-app/backend/command/internal/auth"
	"github.com/oka311119/l4-app/backend/command/internal/domain/entity"
	"github.com/oka311119/l4-app/backend/command/internal/helpers/saltgen"
	"github.com/oka311119/l4-app/backend/command/internal/helpers/uuidgen"
)

type AuthClaims struct {
	jwt.StandardClaims
	UserID string `json:"userid"`
}

type AuthUseCase struct {
	userRepo       auth.Repository
	areaRepo       area.Repository
	pepper         string
	signingKey     []byte
	expireDuration time.Duration
	uuidgen        uuidgen.UUIDGenerator
	saltgen        saltgen.SaltGenerator
}

func NewAuthUseCase(
	userRepo auth.Repository,
	areaRepo area.Repository,
	pepper string,
	signingKey []byte,
	tokenTTL time.Duration,
	uuidgen uuidgen.UUIDGenerator,
	saltgen saltgen.SaltGenerator) *AuthUseCase {
	return &AuthUseCase{
		userRepo:       userRepo,
		areaRepo:       areaRepo,
		pepper:         pepper,
		signingKey:     signingKey,
		expireDuration: time.Second * tokenTTL,
		uuidgen:        uuidgen,
		saltgen:        saltgen,
	}
}

func (a *AuthUseCase) SignUp(ctx context.Context, username, password string) error {
	// Generate new password
	salt, err := a.saltgen.Generate()
	if err != nil {
		return auth.ErrFailedSaltGeneration
	}

	id := a.uuidgen.V4()
	pwd := sha256.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(salt))
	pwd.Write([]byte(a.pepper))

	// Create User
	user := entity.NewUser(
		id,
		username,
		fmt.Sprintf("%x", pwd.Sum(nil)),
		salt,
	)

	err = a.userRepo.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	// Create Default Area
	area := entity.NewArea(
		a.uuidgen.V4(),
		user.ID,
		entity.DefaultAreaName,
	)

	return a.areaRepo.CreateArea(ctx, area)
}

func (a *AuthUseCase) SignIn(ctx context.Context, username, password string) (string, error) {
	user, err := a.userRepo.GetUser(ctx, username)
	if err != nil {
		return "", auth.ErrUserNotFound
	}

	// Password validation
	pwd := sha256.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(user.Salt))
	pwd.Write([]byte(a.pepper))

	if user.Password != fmt.Sprintf("%x", pwd.Sum(nil)) {
		return "", auth.ErrInvalidAccessToken
	}

	claims := AuthClaims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.signingKey)
}

func (a *AuthUseCase) ParseToken(ctx context.Context, accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return a.signingKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.UserID, nil
	}

	return "", auth.ErrInvalidAccessToken
}
