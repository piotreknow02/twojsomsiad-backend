package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"twojsomsiad/config"
	"twojsomsiad/model"

	postman "github.com/rbretecher/go-postman-collection"
)

var (
	host     string
	port     string
	protocol string
)

func GenerateRequestCollections() {
	c := postman.CreateCollection("twojsomsiad-backend", "Endpoints of Twój somsiad backend")
	port = config.Conf.Port
	host = config.Conf.Host
	protocol = "http"

	// --- authentication ---
	authentication := postman.CreateAuth(postman.Bearer, postman.CreateAuthParam("bearer", "<access_token>"))

	// --- auth ---
	auth := c.AddItemGroup("auth")
	// login
	loginData, err := json.Marshal(&model.AuthLoginDTO{
		Email:    "john.doe@example.com",
		Password: "pass12345678",
	})
	if err != nil {
		panic(err)
	}
	auth.AddItem(&postman.Items{
		Name: "log in",
		Request: &postman.Request{
			URL:         getUrl("/auth/login"),
			Method:      postman.Post,
			Description: "Log in",
			Body: &postman.Body{
				Mode: "raw",
				Raw:  string(loginData),
				Options: &postman.BodyOptions{
					Raw: postman.BodyOptionsRaw{
						Language: postman.JSON,
					},
				},
			},
		},
	})
	// register
	registerData, err := json.Marshal(&model.AuthRegisterDTO{
		Username: "johndoe",
		Name:     "John",
		Surname:  "Doe",
		Email:    "john.doe@example.com",
		Password: "pass12345678",
	})
	if err != nil {
		panic(err)
	}
	auth.AddItem(&postman.Items{
		Name: "register",
		Request: &postman.Request{
			URL:         getUrl("/auth/register"),
			Method:      postman.Post,
			Description: "register",
			Body: &postman.Body{
				Mode: "raw",
				Raw:  string(registerData),
				Options: &postman.BodyOptions{
					Raw: postman.BodyOptionsRaw{
						Language: postman.JSON,
					},
				},
			},
		},
	})
	// refresh
	auth.AddItem(&postman.Items{
		Name: "refresh",
		Request: &postman.Request{
			URL:         getUrl("/auth/refresh"),
			Method:      postman.Get,
			Description: "Refresh exipired token",
			Auth:        authentication,
		},
	})

	// --- user ---
	user := c.AddItemGroup("user")
	// get user
	user.AddItem(&postman.Items{
		Name: "get user",
		Request: &postman.Request{
			URL:         getUrl("/user/1"),
			Method:      postman.Get,
			Description: "Get user information by id",
		},
	})
	// get my user
	user.AddItem(&postman.Items{
		Name: "get my user",
		Request: &postman.Request{
			URL:         getUrl("/user"),
			Method:      postman.Get,
			Description: "Get current user info from JWT",
			Auth:        authentication,
		},
	})
	// update user
	updateUserData, err := json.Marshal(model.UserUpdateDTO{
		Username: "johndoe",
		Name:     "John",
		Surname:  "Doe",
		Password: "pass12345678",
	})
	if err != nil {
		panic(err)
	}
	user.AddItem(&postman.Items{
		Name: "update user",
		Request: &postman.Request{
			URL:         getUrl("/user/"),
			Method:      postman.Post,
			Description: "Update current user",
			Auth:        authentication,
			Body: &postman.Body{
				Mode: "raw",
				Raw:  string(updateUserData),
				Options: &postman.BodyOptions{
					Raw: postman.BodyOptionsRaw{
						Language: postman.JSON,
					},
				},
			},
		},
	})
	// get adverts
	user.AddItem(&postman.Items{
		Name: "get adverts",
		Request: &postman.Request{
			URL:         getUrl("/user/adverts"),
			Method:      postman.Get,
			Description: "Get adverts for user",
			Auth:        authentication,
		},
	})

	// --- advert ---
	advert := c.AddItemGroup("advert")
	// get adverts
	advert.AddItem(&postman.Items{
		Name: "get adverts",
		Request: &postman.Request{
			URL:         getUrl("/advert/"),
			Method:      postman.Get,
			Description: "Get adverts",
		},
	})
	// remove advert
	advert.AddItem(&postman.Items{
		Name: "remove advert",
		Request: &postman.Request{
			URL:         getUrl("/advert/1"),
			Method:      postman.Delete,
			Description: "Remove advert by id",
		},
	})
	// create advert
	advertCreateData, err := json.Marshal(model.CreateAdvertDTO{
		Title:       "Test",
		Description: "test description",
		City:        "Kraków",
		Date:        time.Now(),
	})
	if err != nil {
		panic(err)
	}
	advert.AddItem(&postman.Items{
		Name: "create advert",
		Request: &postman.Request{
			URL:         getUrl("advert/"),
			Method:      postman.Post,
			Description: "Create advert",
			Body: &postman.Body{
				Mode: "raw",
				Raw:  string(advertCreateData),
				Options: &postman.BodyOptions{
					Raw: postman.BodyOptionsRaw{
						Language: postman.JSON,
					},
				},
			},
		},
	})

	//apply to advert
	advert.AddItem(&postman.Items{
		Name: "apply to advert",
		Request: &postman.Request{
			URL:         getUrl("advert/1/apply"),
			Method:      postman.Get,
			Description: "Apply to advert",
			Auth:        authentication,
		},
	})

	// application
	advert.AddItem(
		&postman.Items{
			Name: "application",
			Request: &postman.Request{
				URL:         getUrl("advert/1/application"),
				Method:      postman.Get,
				Description: "Get application for advert",
				Auth:        authentication,
			},
		},
	)

	// verify application
	advert.AddItem(
		&postman.Items{
			Name: "verify application",
			Request: &postman.Request{
				URL:         getUrl("advert/1/application/1/verify"),
				Method:      postman.Get,
				Description: "Verify application for advert",
				Auth:        authentication,
			},
		},
	)

	// --- swagger ---
	c.AddItem(&postman.Items{
		Name: "swagger",
		Request: &postman.Request{
			URL:         getUrl("/swagger"),
			Method:      postman.Get,
			Description: "Swagger UI",
		},
	})

	// Save collection to file
	file, err := os.Create("collection.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = c.Write(file, postman.V210)
	if err != nil {
		panic(err)
	}
}

func getUrl(endpoint string) *postman.URL {
	return &postman.URL{
		Raw:      fmt.Sprintf("%s://%s:%s%s", protocol, host, port, endpoint),
		Protocol: protocol,
		Host:     strings.Split(host, "."),
		Port:     port,
		Path:     strings.Split(endpoint, "/"),
	}
}
