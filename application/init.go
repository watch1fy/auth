package application

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/supertokens/supertokens-golang/recipe/dashboard"
	"github.com/supertokens/supertokens-golang/recipe/dashboard/dashboardmodels"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty/tpmodels"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword/tpepmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func SuperTokensInit() {

	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	apiBasePath := os.Getenv("API_BASE_PATH")
	websiteBasePath := os.Getenv("WEBSITE_BASE_PATH")

	// Supertokens inititalization
	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: os.Getenv("SUPERTOKENS_URI"),
			APIKey:        os.Getenv("SUPERTOKENS_API_KEY"),
		},
		AppInfo: supertokens.AppInfo{
			AppName:         os.Getenv("APP_NAME"),
			APIDomain:       os.Getenv("API_DOMAIN") + os.Getenv("PORT") + "/",
			WebsiteDomain:   os.Getenv("WEBSITE_DOMAIN"),
			APIBasePath:     &apiBasePath,
			WebsiteBasePath: &websiteBasePath,
		},
		RecipeList: []supertokens.Recipe{
			thirdpartyemailpassword.Init(&tpepmodels.TypeInput{
				Providers: []tpmodels.ProviderInput{
					{
						Config: tpmodels.ProviderConfig{
							ThirdPartyId: "google",
							Clients: []tpmodels.ProviderClientConfig{
								{
									ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
									ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
								},
							},
						},
					},
					{
						Config: tpmodels.ProviderConfig{
							ThirdPartyId: "apple",
							Clients: []tpmodels.ProviderClientConfig{
								{
									ClientID: os.Getenv("APPLE_CLIENT_ID"),
									AdditionalConfig: map[string]interface{}{
										"keyId":      os.Getenv("APPLE_KEY_ID"),
										"privateKey": os.Getenv("APPLE_PRIVATE_KEY"),
										"teamId":     os.Getenv("APPLE_TEAM_ID"),
									},
								},
							},
						},
					},
				},
			}),
			session.Init(nil), // initializes session
			dashboard.Init(&dashboardmodels.TypeInput{
				Admins: &[]string{
					"anmolsudhir.2001@gmail.com",
				},
			}), // initializes dashboard
		},
	})

	if err != nil {
		panic(err.Error())
	}
}
