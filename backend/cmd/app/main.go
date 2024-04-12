package main

import (
	"context"
	"fmt"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/infra/aws"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/infra/repository"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/infra/web"
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/usecase"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

func main() {
	options := options.Client().ApplyURI(
		fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&appName=%s",
			os.Getenv("MONGODB_ATLAS_USERNAME"),
			os.Getenv("MONGODB_ATLAS_PASSWORD"),
			os.Getenv("MONGODB_ATLAS_CLUSTER_HOSTNAME"),
			os.Getenv("MONGODB_ATLAS_APP_NAME")))
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	OAuth2Repository := aws.NewOAuth2RepositoryCognito(
		os.Getenv("AWS_COGNITO_CLIENT_ID"),
		os.Getenv("AWS_COGNITO_PASSWORD"),
		os.Getenv("AWS_COGNITO_REGION"),
		os.Getenv("AWS_COGNITO_POOL_ID"),
	)
	createUserUsecase := usecase.NewCreateUserUsecase(OAuth2Repository)
	getUserUsecase := usecase.NewGetUserUsecase(OAuth2Repository)
	userConfirmation := usecase.NewUserConfirmationUsecase(OAuth2Repository)
	userSignInUsecase := usecase.NewUserSignInUsecase(OAuth2Repository)
	userHandlers := web.NewUserHandlers(createUserUsecase, getUserUsecase, userConfirmation, userSignInUsecase)

	sensorsRepository := repository.NewSensorRepositoryMongo(client, "mongodb", "sensors")
	createSensorUseCase := usecase.NewCreateSensorUseCase(sensorsRepository)
	sensorHandlers := web.NewSensorHandlers(createSensorUseCase)

	alertRepository := repository.NewAlertRepositoryMongo(client, "mongodb", "alerts")
	createAlertUseCase := usecase.NewCreateAlertUseCase(alertRepository)
	findAllAlertsUseCase := usecase.NewFindAllAlertsUseCase(alertRepository)
	alertHandlers := web.NewAlertHandlers(createAlertUseCase, findAllAlertsUseCase)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", userHandlers.ValidateHandler)
	mux.HandleFunc("POST /users/signup", userHandlers.CreateUserHandler)
	mux.HandleFunc("POST /users/confirmation", userHandlers.UserConfirmationHandler)
	mux.HandleFunc("POST /users/signin", userHandlers.UserSignInHandler)
	mux.HandleFunc("GET /sensors", sensorHandlers.CreateSensorHandler)
	mux.HandleFunc("GET /alerts", alertHandlers.CreateAlertHandler)
	mux.HandleFunc("POST /alerts", alertHandlers.CreateAlertHandler)
	mux.HandleFunc("POST /sensors", sensorHandlers.CreateSensorHandler)
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)
}
