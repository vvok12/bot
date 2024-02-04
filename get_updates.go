package bot

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-telegram/bot/models"
)

const (
	maxTimeoutAfterError = time.Second * 5
)

type configurableGetUpdateParams struct {
	Limit          int      `json:"limit,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

type getUpdatesParams struct {
	configurableGetUpdateParams `json:"-,flatten"`
	Offset                      int64 `json:"offset,omitempty"`
	Timeout                     int   `json:"timeout,omitempty"`
}

// GetUpdates https://core.telegram.org/bots/api#getupdates
func (b *Bot) getUpdates(ctx context.Context, wg *sync.WaitGroup, params ...*configurableGetUpdateParams) {
	defer wg.Done()

	var timeoutAfterError time.Duration

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		if timeoutAfterError > 0 {
			if b.isDebug {
				b.debugHandler("wait after error, %v", timeoutAfterError)
			}
			select {
			case <-ctx.Done():
				return
			case <-time.After(timeoutAfterError):
			}
		}

		var updates []*models.Update

		getUpdatesParams := b.buildGetUpdatesParams(params...)
		errRequest := b.rawRequest(ctx, "getUpdates", getUpdatesParams, &updates)
		if errRequest != nil {
			if errors.Is(errRequest, context.Canceled) {
				return
			}
			b.error("error get updates, %v", errRequest)
			timeoutAfterError = incErrTimeout(timeoutAfterError)
			continue
		}

		timeoutAfterError = 0

		for _, upd := range updates {
			atomic.StoreInt64(&b.lastUpdateID, upd.ID)
			select {
			case b.updates <- upd:
			default:
				b.error("error send update to processing, channel is full")
			}
		}
	}
}

func (b *Bot) buildGetUpdatesParams(params ...*StartParams) *getUpdatesParams {
	result := &getUpdatesParams{
		Timeout: int((b.pollTimeout - time.Second).Seconds()),
		Offset:  atomic.LoadInt64(&b.lastUpdateID) + 1,
	}

	if len(params) > 0 {
		p := params[0]

		result.AllowedUpdates = p.AllowedUpdates
		result.Limit = p.Limit
	}

	return result
}

func incErrTimeout(timeout time.Duration) time.Duration {
	if timeout == 0 {
		return time.Millisecond * 100 // first timeout
	}
	timeout *= 2
	if timeout > maxTimeoutAfterError {
		return maxTimeoutAfterError
	}
	return timeout
}
