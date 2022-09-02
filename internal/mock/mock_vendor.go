package mock

import (
	"math"
	"time"

	"github.com/sergera/marketplace-vendor-edge-api/internal/domain"
	"github.com/sergera/marketplace-vendor-edge-api/internal/service"
	"github.com/sergera/marketplace-vendor-edge-api/pkg/random"
)

type MockVendor struct {
	service *service.WorkerService
}

func NewMockVendor() *MockVendor {
	return &MockVendor{service.NewWorkerService()}
}

func (mv MockVendor) MockOrderStatusUpdates(o domain.OrderModel) {
	currentStatus := domain.Unconfirmed
	for int(currentStatus) <= int(domain.Delivered) {
		var minSeconds int
		var maxSeconds int
		if currentStatus == domain.Unconfirmed {
			minSeconds = 10
			maxSeconds = 30
		} else {
			minSeconds = int(currentStatus) * 10
			maxSeconds = int(math.Pow(float64(currentStatus), 5))
		}
		seconds := random.IntInRange(minSeconds, maxSeconds)
		time.Sleep(time.Duration(seconds) * time.Second)
		currentStatus++
		o.Status = currentStatus.String()
		mv.service.UpdateStatus(o)
	}
}
