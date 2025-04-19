package settings

import "time"

const (
	TimeoutForTimeoutMiddleware = 20 * time.Second

	WebsocketReadBufferSize  = 1024
	WebsocketWriteBufferSize = 1024
)
