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
	// ErrMissingCustomerID error when missing customer id
	ErrMissingCustomerID = errors.New("missing customer id")
	// ErrMissingUsername error when missing username
	ErrMissingUsername = errors.New("missing username")
	// ErrInvalidGrant error when there is an invalid grant
	ErrInvalidGrant = errors.New("invalid grant")
	// ErrNoHMACSigningMethod error when no signing method HMAC
	ErrNoHMACSigningMethod = errors.New("no signing method HMAC")
	// ErrClaimsNotFound error when claims not found
	ErrClaimsNotFound = errors.New("claims not found")
	// ErrMissingTokenID error when missing token id
	ErrMissingTokenID = errors.New("missing token id")
	// ErrAPIKeyExpired error when the api key has expired
	ErrAPIKeyExpired = errors.New("the api key has expired")
	// ErrUnauthorizedEndpoint error when an unauthorized endpoint is called
	ErrUnauthorizedEndpoint = errors.New("unauthorized endpoint")
)

// APIKey is a structure that contains api key information
type APIKey struct {
	ID         string `json:"id"`
	CustomerID int    `json:"customer_id"`
	UserID     string `json:"user_id"`
	Username   string `json:"username"`
	PrivateKey string `json:"private_key"`
	TimeToLive int64  `json:"time_to_live,omitempty"`
}

// APIKeyResponse is a structure containing the api key
type APIKeyResponse struct {
	AccessToken string `json:"access_token"`
}

// NewAPIKey is the APIKey constructor
func NewAPIKey(admin *models.Admin, userID string, username string) (*APIKey, error) {
	if admin.IDadmin == 0 {
		return nil, ErrMissingCustomerID
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
		CustomerID: admin.IDadmin,
		UserID:     userID,
		Username:   username,
		PrivateKey: admin.SecretKey,
	}

	return apiKey, nil
}

// GenerateAPIKey is a function to generate an api key
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
		"id":          apiKey.ID,
		"customer_id": strconv.Itoa(apiKey.CustomerID),
		"user_id":     apiKey.UserID,
		"username":    apiKey.Username,
		"iat":         now.Unix(),
		"exp":         now.Add(apiKEYMaxDuration).Unix(),
		"iss":         tgaISS,
	}
}

// VerifyAPIKey is a function to validate an api key
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

		customerID := ClaimsToString(claims, "customer_id")
		if customerID == "" {
			return "", ErrMissingCustomerID
		}

		clientID, err := strconv.Atoi(customerID)
		if err != nil {
			return "", err
		}

		username := ClaimsToString(claims, "username")
		if username == "" {
			return "", ErrMissingUsername
		}

		admin, err := GetAdminByID(clientID, username)
		if err != nil {
			return "", err
		}

		return []byte(admin.SecretKey), nil
	})

	err = validateAPIKey(err)

	return token, err
}

// ClaimsToString returns the value in the jwt claims
func ClaimsToString(claims jwt.MapClaims, key string) string {
	value, ok := claims[key].(string)
	if !ok || value == "" {
		return ""
	}

	return value
}

// ClaimsToInt returns the value in the jwt claims
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
