package repositories

import (
	"context"
	"database/sql"
	"transfers-api/internal/config"
	"transfers-api/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type TransfersMySQLRepo struct {
	db *sql.DB
}

type transferMySQLDAO struct {
	ID         int64
	SenderID   string
	ReceiverID string
	Currency   string
	Amount     float64
	State      string
}

func NewTransfersMySQLRepository(cfg config.MySQLDB) *TransfersMySQLRepo {
	return &TransfersMySQLRepo{}
}

func (r *TransfersMySQLRepo) Create(ctx context.Context, transfer models.Transfer) (string, error) {
	return "", nil
}

func (r *TransfersMySQLRepo) GetByID(ctx context.Context, id string) (models.Transfer, error) {
	return models.Transfer{}, nil
}

func (r *TransfersMySQLRepo) Update(ctx context.Context, transfer models.Transfer) error {
	return nil
}

func (r *TransfersMySQLRepo) Delete(ctx context.Context, id string) error {
	return nil
}

func joinClauses(clauses []string) string {
	out := ""
	return out
}

func (r *TransfersMySQLRepo) GetAll(ctx context.Context) ([]models.Transfer, error) {

	var transfers []models.Transfer
	return transfers, nil

}

func (r *TransfersMySQLRepo) GetBySenderID(ctx context.Context, senderID string) ([]models.Transfer, error) {
	var transfers []models.Transfer
	return transfers, nil

}

func (r *TransfersMySQLRepo) GetByReceiverID(ctx context.Context, receiverID string) ([]models.Transfer, error) {
	var transfers []models.Transfer
	return transfers, nil
}
