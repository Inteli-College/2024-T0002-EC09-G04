package aws

import (
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/domain/entity"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/google/uuid"
)

type OAuth2RepositoryCognito struct {
	client   *cognito.CognitoIdentityProvider
	id       string
	region   string
	pool_id  string
	password string
}

func NewOAuth2RepositoryCognito(id string, password string, region string, pool_id string) *OAuth2RepositoryCognito {
	config := &aws.Config{Region: aws.String("us-east-1")}
	sess, err := session.NewSession(config)
	
	if err != nil {
		panic(err)
	}
	client := cognito.New(sess)

	return &OAuth2RepositoryCognito{
		client:   client,
		id:       id,
		region:   region,
		pool_id:  pool_id,
		password: password,
	}
}

func (c *OAuth2RepositoryCognito) SignUp(user *entity.User) error {
	userCognito := &cognito.SignUpInput{
		ClientId: aws.String(c.id),
		Username: aws.String(user.Email),
		Password: aws.String(user.Password),
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String("name"),
				Value: aws.String(user.Name),
			},
			{
				Name:  aws.String("email"),
				Value: aws.String(user.Email),
			},
			{
				Name:  aws.String("custom:custom_id"),
				Value: aws.String(uuid.NewString()),
			},
		},
	}
	_, err := c.client.SignUp(userCognito)
	if err != nil {
		return err
	}
	return nil
}

func (c *OAuth2RepositoryCognito) ConfirmAccount(confirmation *entity.Confirmation) error {
	confirmationInput := &cognito.ConfirmSignUpInput{
		Username:         aws.String(confirmation.Email),
		ConfirmationCode: aws.String(confirmation.Code),
		ClientId:         aws.String(c.id),
	}
	_, err := c.client.ConfirmSignUp(confirmationInput)
	if err != nil {
		return err
	}
	return nil
}

func (c *OAuth2RepositoryCognito) SignIn(login *entity.Login) (string, error) {
	authInput := &cognito.InitiateAuthInput{
		AuthFlow: aws.String(c.password),
		AuthParameters: aws.StringMap(map[string]string{
			"USERNAME": login.Email,
			"PASSWORD": login.Password,
		}),
		ClientId: aws.String(c.id),
	}
	result, err := c.client.InitiateAuth(authInput)
	if err != nil {
		return "", err
	}
	return *result.AuthenticationResult.AccessToken, nil
}

func (c *OAuth2RepositoryCognito) GetUserByToken(token string) (*cognito.GetUserOutput, error) {
	input := &cognito.GetUserInput{
		AccessToken: aws.String(token),
	}
	result, err := c.client.GetUser(input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *OAuth2RepositoryCognito) UpdatePassword(login *entity.Login) error {
	input := &cognito.AdminSetUserPasswordInput{
		UserPoolId: aws.String(c.pool_id),
		Username:   aws.String(login.Email),
		Password:   aws.String(login.Password),
		Permanent:  aws.Bool(true),
	}
	_, err := c.client.AdminSetUserPassword(input)
	if err != nil {
		return err
	}
	return nil
}
