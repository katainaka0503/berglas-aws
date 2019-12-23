package client

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
)

type Client struct {
	secretsmanager secretsmanageriface.SecretsManagerAPI
}

func NewClientWithContext() (*Client, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, fmt.Errorf("failed to init session: %v", err)
	}

	return &Client{
		secretsmanager: secretsmanager.New(sess),
	}, nil
}
