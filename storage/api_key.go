package storage

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"github.com/jhoancamilorayomejia/TGacueducto/models"
)

const (
	tgaISS = "tga"

	apiKEYMaxDuration = 438000 * time.Hour
)

var (
	ErrMissingCompanyID     = errors.New("missing company id")
	ErrMissingUsername      = errors.New("missing username")
	ErrInvalidGrant         = errors.New("invalid grant")
	ErrNoHMACSigningMethod  = errors.New("no signing method HMAC")
	ErrClaimsNotFound       = errors.New("claims not found")
	ErrMissingTokenID       = errors.New("missing token id")
	ErrAPIKeyExpired        = errors.New("the api key has expired")
	ErrUnauthorizedEndpoint = errors.New("unauthorized endpoint")
)

type APIKey struct {
	ID         string `json:"id"`
	CompanyID  int    `json:"company_id"`
	UserID     string `json:"user_id"`
	Username   string `json:"username"`
	PrivateKey string `json:"private_key"`
	TimeToLive int64  `json:"time_to_live,omitempty"`
}

type APIKeyResponse struct {
	AccessToken string `json:"access_token"`
}

func NewAPIKey(admin *models.Admin, userID string, username string) (*APIKey, error) {
	if admin.IDadmin == 0 {
		return nil, ErrMissingCompanyID
	}

	if username == "" {
		return nil, ErrMissingUsername
	}

	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	apiKey := &APIKey{
		ID:         id.String(),
		CompanyID:  admin.IDadmin,
		UserID:     userID,
		Username:   username,
		PrivateKey: admin.SecretKey,
	}

	return apiKey, nil
}

// NewCompanyAPIKey is the APIKey constructor for companies
func NewCompanyAPIKey(company *models.Company, userID string, username string) (*APIKey, error) {
	if company.IDcompany == 0 {
		return nil, ErrMissingCompanyID
	}

	if username == "" {
		return nil, ErrMissingUsername
	}

	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	apiKey := &APIKey{
		ID:         id.String(),
		CompanyID:  company.IDcompany,
		UserID:     userID,
		Username:   username,
		PrivateKey: company.SecretKey,
	}

	return apiKey, nil
}

func (apiKey *APIKey) GenerateAPIKey(ctx context.Context) (*APIKeyResponse, error) {
	atClaims := apiKey.getClaims()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	accessToken, err := at.SignedString([]byte(apiKey.PrivateKey))
	if err != nil {
		return nil, err
	}

	return &APIKeyResponse{
		AccessToken: accessToken,
	}, nil
}

func (apiKey *APIKey) getClaims() jwt.MapClaims {
	now := time.Now()

	return jwt.MapClaims{
		"id":         apiKey.ID,
		"company_id": strconv.Itoa(apiKey.CompanyID),
		"user_id":    apiKey.UserID,
		"username":   apiKey.Username,
		"iat":        now.Unix(),
		"exp":        now.Add(apiKEYMaxDuration).Unix(),
		"iss":        tgaISS,
	}
}

func VerifyAPIKey(ctx context.Context, tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", ErrNoHMACSigningMethod
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return "", ErrClaimsNotFound
		}

		id := ClaimsToString(claims, "id")
		if id == "" {
			return "", ErrMissingTokenID
		}

		companyID := ClaimsToString(claims, "company_id")
		if companyID == "" {
			return "", ErrMissingCompanyID
		}

		clientID, err := strconv.Atoi(companyID)
		if err != nil {
			return "", err
		}

		username := ClaimsToString(claims, "username")
		if username == "" {
			return "", ErrMissingUsername
		}

		admin, err := GetAdminByID(clientID, username)
		if err != nil {
			company, err := GetCompanyByID(clientID, username)
			if err != nil {
				return "", err
			}
			return []byte(company.SecretKey), nil
		}

		return []byte(admin.SecretKey), nil
	})

	err = validateAPIKey(err)

	return token, err
}

func ClaimsToString(claims jwt.MapClaims, key string) string {
	value, ok := claims[key].(string)
	if !ok || value == "" {
		return ""
	}

	return value
}

func ClaimsToInt(claims jwt.MapClaims, key string) int {
	value, ok := claims[key].(int)
	if !ok || value == 0 {
		return 0
	}

	return value
}

func validateAPIKey(err error) error {
	prevErr := err
	err = validateTokenExpiration(err)
	if err != nil {
		return err
	}

	if prevErr != nil {
		return errors.New(prevErr.Error())
	}

	return nil
}

func validateTokenExpiration(err error) error {
	var jwtErr *jwt.ValidationError
	if ok := errors.As(err, &jwtErr); ok && jwtErr.Errors == jwt.ValidationErrorExpired {
		return ErrAPIKeyExpired
	}

	return nil
}
