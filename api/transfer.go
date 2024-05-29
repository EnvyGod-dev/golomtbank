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
	BankName      string `json:"Bank_name" binding:"required"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	fromAccount, valid := server.validAccount(ctx, req.FromAccountId, req.Currency)
	if !valid {
		return
	}

	authorizationPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if fromAccount.Owner != authorizationPayload.Username {
		err := errors.New("from account doesnt belong authicated user")
		ctx.JSON(http.StatusUnauthorized, errResponse(err))
		return
	}

	_, valid = server.validAccount(ctx, req.ToAccountId, req.Currency)
	if !valid {
		return
	}

	var fee int64
	if req.BankName == "Голомт Банк" {
		fee = 0
	} else {
		fee = 200
	}

	adjustedAmount := req.Amount + fee

	if adjustedAmount < 0 {
		err := errors.New("Үлдэгдэл хүрэлцэхгүй байна")
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
		err := fmt.Errorf("account [%d] currency missmatch %v vs %v", account.Id, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return account, false
	}
	return account, true
}
