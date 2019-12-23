package resolution

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
)

type Resolver struct {
	secretsmanager secretsmanageriface.SecretsManagerAPI
}

func NewResolverWithContext() (*Resolver, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, fmt.Errorf("failed to init session: %v", err)
	}

	return &Resolver{
		secretsmanager: secretsmanager.New(sess),
	}, nil
}
