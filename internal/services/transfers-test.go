package services

import (
	"context"
	"errors"
	"testing"
	"transfers-api/internal/config"
	"transfers-api/internal/enums"
	"transfers-api/internal/known_errors"
	"transfers-api/internal/models"
	"transfers-api/internal/services/mocks"

	"github.com/stretchr/testify/assert"
)

func TestTransferService_GetById(t *testing.T) {
	var (
		context = context.Background()
		cfg     = config.BusinessConfig{
			TransferMinAmount: 1,
		}
		transferRepo   = mocks.NewTransferServiceMock(t)
		transferCache  = mocks.NewTransferServiceMock(t)
		eventPublisher = mocks.NewEventPublisherMock(t)
		transferId     = "test-1"
		transferDto    = models.Transfer{
			ID:         transferId,
			SenderID:   "139877-123",
			ReceiverID: "11121455-334",
			Amount:     1293.34,
			Currency:   enums.CurrencyUSD,
		}
	)
	transferCache.On("GetByID", context, transferId).Return(transferDto, nil)

	service := NewTransfersService(cfg, transferRepo, transferCache, eventPublisher)

	tran, err := service.GetByID(context, transferId)

	assert.Nil(t, err)

	assert.Equal(t, transferId, tran.ID)
}
func TestTransferService_Update(t *testing.T) {
	var (
		context = context.Background()
		cfg     = config.BusinessConfig{
			TransferMinAmount: 1,
		}
		transferRepo   = mocks.NewTransferServiceMock(t)
		transferCache  = mocks.NewTransferServiceMock(t)
		eventPublisher = mocks.NewEventPublisherMock(t)
		transferDto    = models.Transfer{
			ID:         "mal-id",
			SenderID:   "139877-123",
			ReceiverID: "11121455-334",
			Amount:     1293.34,
			Currency:   enums.CurrencyUSD,
		}
	)

	transferRepo.On("Update", context, transferDto).Return(known_errors.ErrBadRequest)
	service := NewTransfersService(cfg, transferRepo, transferCache, eventPublisher)

	err := service.Update(context, transferDto)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, known_errors.ErrBadRequest) {
		t.Fatalf("expected ErrBadRequest, got %v", err)
	}
}
