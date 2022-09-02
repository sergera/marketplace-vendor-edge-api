package api

import (
	"encoding/json"
	"net/http"

	"github.com/sergera/marketplace-vendor-edge-api/internal/domain"
	"github.com/sergera/marketplace-vendor-edge-api/internal/mock"
)

type VendorAPI struct {
	mock *mock.MockVendor
}

func NewVendorAPI() *VendorAPI {
	return &VendorAPI{mock.NewMockVendor()}
}

func (v *VendorAPI) SendOrder(w http.ResponseWriter, r *http.Request) {
	var m domain.OrderModel

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := m.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orderInBytes, err := json.Marshal(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(orderInBytes)

	go v.mock.MockOrderStatusUpdates(m)
}
