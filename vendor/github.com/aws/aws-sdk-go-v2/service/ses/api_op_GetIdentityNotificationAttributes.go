// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to return the notification attributes for a list of
// identities you verified with Amazon SES. For information about Amazon SES
// notifications, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/notifications.html).
type GetIdentityNotificationAttributesInput struct {
	_ struct{} `type:"structure"`

	// A list of one or more identities. You can specify an identity by using its
	// name or by using its Amazon Resource Name (ARN). Examples: user@example.com,
	// example.com, arn:aws:ses:us-east-1:123456789012:identity/example.com.
	//
	// Identities is a required field
	Identities []string `type:"list" required:"true"`
}

// String returns the string representation
func (s GetIdentityNotificationAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetIdentityNotificationAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetIdentityNotificationAttributesInput"}

	if s.Identities == nil {
		invalidParams.Add(aws.NewErrParamRequired("Identities"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the notification attributes for a list of identities.
type GetIdentityNotificationAttributesOutput struct {
	_ struct{} `type:"structure"`

	// A map of Identity to IdentityNotificationAttributes.
	//
	// NotificationAttributes is a required field
	NotificationAttributes map[string]IdentityNotificationAttributes `type:"map" required:"true"`
}

// String returns the string representation
func (s GetIdentityNotificationAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetIdentityNotificationAttributes = "GetIdentityNotificationAttributes"

// GetIdentityNotificationAttributesRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Given a list of verified identities (email addresses and/or domains), returns
// a structure describing identity notification attributes.
//
// This operation is throttled at one request per second and can only get notification
// attributes for up to 100 identities at a time.
//
// For more information about using notifications with Amazon SES, see the Amazon
// SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/notifications.html).
//
//    // Example sending a request using GetIdentityNotificationAttributesRequest.
//    req := client.GetIdentityNotificationAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/GetIdentityNotificationAttributes
func (c *Client) GetIdentityNotificationAttributesRequest(input *GetIdentityNotificationAttributesInput) GetIdentityNotificationAttributesRequest {
	op := &aws.Operation{
		Name:       opGetIdentityNotificationAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetIdentityNotificationAttributesInput{}
	}

	req := c.newRequest(op, input, &GetIdentityNotificationAttributesOutput{})
	return GetIdentityNotificationAttributesRequest{Request: req, Input: input, Copy: c.GetIdentityNotificationAttributesRequest}
}

// GetIdentityNotificationAttributesRequest is the request type for the
// GetIdentityNotificationAttributes API operation.
type GetIdentityNotificationAttributesRequest struct {
	*aws.Request
	Input *GetIdentityNotificationAttributesInput
	Copy  func(*GetIdentityNotificationAttributesInput) GetIdentityNotificationAttributesRequest
}

// Send marshals and sends the GetIdentityNotificationAttributes API request.
func (r GetIdentityNotificationAttributesRequest) Send(ctx context.Context) (*GetIdentityNotificationAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetIdentityNotificationAttributesResponse{
		GetIdentityNotificationAttributesOutput: r.Request.Data.(*GetIdentityNotificationAttributesOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetIdentityNotificationAttributesResponse is the response type for the
// GetIdentityNotificationAttributes API operation.
type GetIdentityNotificationAttributesResponse struct {
	*GetIdentityNotificationAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetIdentityNotificationAttributes request.
func (r *GetIdentityNotificationAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}