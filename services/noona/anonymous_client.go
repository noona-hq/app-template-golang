package noona

import (
	"context"
	"net/http"

	noona "github.com/noona-hq/noona-sdk-go"
	"github.com/pkg/errors"
)

type AnonymousClient struct {
	cfg    Config
	Client *noona.ClientWithResponses
}

func (n AnonymousClient) clientCredentials(_ context.Context, req *http.Request) error {
	req.SetBasicAuth(n.cfg.ClientID, n.cfg.ClientSecret)

	query := req.URL.Query()
	query.Del("client_id")
	query.Del("client_secret")
	req.URL.RawQuery = query.Encode()

	return nil
}

func (n AnonymousClient) CodeTokenExchange(code string) (*noona.OAuthToken, error) {
	tokenResponse, err := n.Client.GetOAuthTokenWithResponse(context.Background(), &noona.GetOAuthTokenParams{}, noona.GetOAuthTokenJSONRequestBody{
		Code:      &code,
		GrantType: noona.AuthorizationCode,
	}, n.clientCredentials)
	if err != nil {
		return nil, errors.Wrap(err, "Error in code exchange")
	}

	if tokenResponse.StatusCode() != http.StatusOK {
		return nil, errors.New("Error in code exchange")
	}

	return tokenResponse.JSON200, nil
}

func (n AnonymousClient) RefreshTokenExchange(refreshToken string) (*noona.OAuthToken, error) {
	tokenResponse, err := n.Client.GetOAuthTokenWithResponse(context.Background(), &noona.GetOAuthTokenParams{}, noona.GetOAuthTokenJSONRequestBody{
		RefreshToken: &refreshToken,
		GrantType:    noona.RefreshToken,
	}, n.clientCredentials)
	if err != nil {
		return nil, errors.Wrap(err, "Error in refresh token exchange")
	}

	if tokenResponse.StatusCode() != http.StatusOK {
		return nil, errors.New("Error in refresh token exchange")
	}

	return tokenResponse.JSON200, nil
}
