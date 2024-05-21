package noona

type Config struct {
	BaseURL      string `default:"https://api.noona.is"`
	AppStoreURL  string `default:"https://hq.noona.app/week#settings-apps"`
	ClientID     string `default:""` // TODO: Fill in with your client ID
	ClientSecret string `default:""` // TODO: Fill in with your client secret
	// Our app will rely on webhooks to implement the Blacklist.
	// We need to know where the webhook should be sent.
	AppBaseURL string `default:"http://localhost:8080"`
	// The Bearer token that all webhooks will be sent with
	AppWebhookToken string `default:"very-secure-token-secret"`
}
