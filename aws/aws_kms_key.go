// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kms"
)

func ListKmsKey(client *Client) ([]Resource, error) {
	req := client.kmsconn.ListKeysRequest(&kms.ListKeysInput{})

	var result []Resource

	p := kms.NewListKeysPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Keys {

			result = append(result, Resource{
				Type:   "aws_kms_key",
				ID:     *r.KeyId,
				Region: client.kmsconn.Config.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
