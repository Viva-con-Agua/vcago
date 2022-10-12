package vcago

import (
	"context"
	"errors"
	"log"

	"github.com/Viva-con-Agua/vcago/vmod"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type HydraClient struct {
	Oauth2Config oauth2.Config
	Verifier     *oidc.IDTokenVerifier
}

func NewHydraClient() (r *HydraClient) {
	if !Settings.Bool("OIDC_SKIP", "n", false) {
		ctx := context.Background()
		r = new(HydraClient)
		provider, err := oidc.NewProvider(ctx, Settings.String("OIDC_HOST", "w", "http://hydra.localhost/"))
		if err != nil {
			log.Print(err)
		}
		// Configure an OpenID Connect aware OAuth2 client.
		r.Oauth2Config = oauth2.Config{
			ClientID:     Settings.String("OIDC_CLIENT_ID", "w", "test"),
			ClientSecret: Settings.String("OIDC_CLIENT_SECRET", "w", "secret"),
			RedirectURL:  Settings.String("OIDC_REDIRECT_URL", "w", "http://localhost:8081/callback"),

			// Discovery returns the OAuth2 endpoints.
			Endpoint: provider.Endpoint(),

			// "openid" is a required scope for OpenID Connect flows.
			Scopes: []string{oidc.ScopeOpenID, "extra_vars"},
		}
		r.Oauth2Config.Endpoint.AuthStyle = oauth2.AuthStyleAutoDetect
		r.Verifier = provider.Verifier(&oidc.Config{ClientID: r.Oauth2Config.ClientID})
	}
	return
}

type Callback struct {
	Code string `json:"code"`
}

type UserClaims struct {
	User vmod.User `json:"user"`
}

func (i *HydraClient) Callback(ctx context.Context, callback *Callback) (r *vmod.User, err error) {
	oauth2Token := new(oauth2.Token)
	oauth2Token, err = i.Oauth2Config.Exchange(ctx, callback.Code)
	if err != nil {
		return
	}
	// Extract the ID Token from OAuth2 token.
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("hydra token is missing")
	}
	// Parse and verify ID Token payload.
	idToken := new(oidc.IDToken)
	idToken, err = i.Verifier.Verify(ctx, rawIDToken)
	if err != nil {
		return
	}
	resp := struct {
		OAuth2Token   *oauth2.Token
		IDTokenClaims *vmod.User // ID Token payload is just JSON.
	}{oauth2Token, new(vmod.User)}

	if err = idToken.Claims(&resp.IDTokenClaims); err != nil {
		return
	}
	r = resp.IDTokenClaims
	return
}
