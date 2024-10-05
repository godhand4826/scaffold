package ent

import (
	"context"
)

func (c *Client) Ping(ctx context.Context) error {
	return c.driver.Exec(ctx, "SELECT 1", []interface{}{}, nil)
}
