package mock

import (
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
		minSeconds := int(currentStatus)
		maxSeconds := int(currentStatus) * 5
		seconds := random.IntInRange(minSeconds, maxSeconds)
		time.Sleep(time.Duration(seconds) * time.Second)
		currentStatus++
		o.Status = currentStatus.String()
		mv.service.UpdateStatus(o)
	}
}
