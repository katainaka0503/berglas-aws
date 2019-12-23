package resolution

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"strings"
)

func (c *Resolver) Resolve(url string) (string, error) {
	schemeAndArn := strings.SplitN(url, "://", 2)

	if len(schemeAndArn) != 2 || schemeAndArn[0] != "berglas-aws" {
		return "", fmt.Errorf("url format is invalid: %v", url)
	}

	arn := schemeAndArn[1]

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(arn),
	}

	output, err := c.secretsmanager.GetSecretValue(input)
	if err != nil {
		return "", fmt.Errorf("failed to get secret value: %w", err)
	}

	return *output.SecretString, nil
}

// returns whether value is in a format of below or not
// berglas-aws://arn:${Partition}:secretsmanager:${Region}:${Account}:secret:${SecretId}
func IsResolvable(url string) bool {
	schemeAndArn := strings.SplitN(url, "://", 2)

	if len(schemeAndArn) != 2 || schemeAndArn[0] != "berglas-aws" {
		return false
	}

	arn := schemeAndArn[1]

	partOfArn := strings.Split(arn, ":")
	return len(partOfArn) == 7 && partOfArn[0] == "arn" && partOfArn[2] == "secretsmanager" && partOfArn[5] == "secret"
}
