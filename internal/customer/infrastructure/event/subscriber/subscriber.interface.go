package subscriber

import "context"

type Subscriber interface {
	Run(ctx context.Context)
}
