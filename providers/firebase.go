package providers

import (
	"context"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/samber/do"
	"google.golang.org/api/option"
)

func ProvideFirebase(injector *do.Injector) {
    do.ProvideNamed(injector, constants.FCMClient, func(i *do.Injector) (*messaging.Client, error) {
        ctx := context.Background()
        credPath := os.Getenv("FIREBASE_CREDENTIALS")
        var app *firebase.App
        var err error
        if credPath != "" {
            app, err = firebase.NewApp(ctx, nil, option.WithCredentialsFile(credPath))
        } else {
            app, err = firebase.NewApp(ctx, nil)
        }
        if err != nil {
            return nil, err
        }
        client, err := app.Messaging(ctx)
        if err != nil {
            return nil, err
        }
        return client, nil
    })
}
