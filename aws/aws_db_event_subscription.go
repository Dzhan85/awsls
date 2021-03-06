// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func ListDbEventSubscription(client *Client) ([]Resource, error) {
	req := client.rdsconn.DescribeEventSubscriptionsRequest(&rds.DescribeEventSubscriptionsInput{})

	var result []Resource

	p := rds.NewDescribeEventSubscriptionsPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.EventSubscriptionsList {

			result = append(result, Resource{
				Type:   "aws_db_event_subscription",
				ID:     *r.CustSubscriptionId,
				Region: client.rdsconn.Config.Region,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
