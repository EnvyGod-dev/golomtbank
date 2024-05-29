package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	db "task/db/sqlc"
	"task/token"

	"github.com/gin-gonic/gin"
)

type transferRequest struct {
	FromAccountId int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountId   int64  `json:"to_account_id" binding:"required,min=1"`
	BankName      string `json:"bank_name" binding:"required"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	// Validate BankName before proceeding
	if !isValidBankName(req.BankName) {
		err := fmt.Errorf("invalid BankName: %s", req.BankName)
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	fromAccount, valid := server.validAccount(ctx, req.FromAccountId, req.Currency)
	if !valid {
		return
	}

	authorizationPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if fromAccount.Owner != authorizationPayload.Username {
		err := errors.New("from account doesn't belong to authenticated user")
		ctx.JSON(http.StatusUnauthorized, errResponse(err))
		return
	}

	_, valid = server.validAccount(ctx, req.ToAccountId, req.Currency)
	if !valid {
		return
	}

	var fee int64
	if req.BankName != "Голомт Банк" {
		fee = 100
	}

	adjustedAmount := req.Amount + fee

	if adjustedAmount <= 0 {
		err := errors.New("balance is 0")
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.TransferTxParams{
		FromAccountId: req.FromAccountId,
		ToAccountId:   req.ToAccountId,
		BankName:      req.BankName,
		Amount:        adjustedAmount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func isValidBankName(bankName string) bool {
	// Define the valid bank names as per the check constraint
	validBankNames := map[string]bool{
		"Голомт Банк": true,
		"Хаан банк":   true,
		"Mbank":       true,
		"Төрийн банк": true,
		"Худалдаа хөгжлийн банк": true,
		"Богд Банк":              true,
	}

	return validBankNames[bankName]
}

func (server *Server) validAccount(ctx *gin.Context, accountId int64, currency string) (db.Account, bool) {
	account, err := server.store.GetAccountById(ctx, accountId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return account, false
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return account, false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch %v vs %v", account.Id, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return account, false
	}
	return account, true
}

type getTransfersRequest struct {
	ID int64 `form:"id" binding:"required"`
}

func (server *Server) getTransfer(ctx *gin.Context) {
	var req getTransfersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	transfer, err := server.store.ListTransfers(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, transfer)
}
