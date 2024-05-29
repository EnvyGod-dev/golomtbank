/* trunk-ignore-all(golangci-lint/typecheck) */
package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	db "task/db/sqlc"
	"task/token"
	"task/utils"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type CreateAccountsRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Balance  int64  `json:"balance" binding:"required"`
	BankName string `json:"bankanme" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}

func (server *Server) createAccounts(ctx *gin.Context) {
	var req CreateAccountsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateAccountsParams{
		Id:       utils.GenerateAccountNumber(req.BankName),
		Owner:    authPayload.Username,
		Currency: req.Currency,
		BankName: req.BankName,
		Balance:  req.Balance,
	}

	account, err := server.store.CreateAccounts(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

type getAccount struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccounts(ctx *gin.Context) {
	var req getAccount

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Received Id:", req.Id)

	account, err := server.store.GetAccountById(ctx, req.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if account.Owner != authPayload.Username {
		err := errors.New("Хэрэглэгчийг баталгаажуулж чадсангүй")
		ctx.JSON(http.StatusUnauthorized, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type listAccountsRequest struct {
	PageId   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=5"`
}

func (server *Server) listAccounts(ctx *gin.Context) {
	var req listAccountsRequest
	if err := ctx.ShouldBind(&req); err != nil {
		fmt.Println("Error binding request:", err)
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.ListAccountsParams{
		Owner:  authPayload.Username,
		Limit:  req.PageSize,
		Offset: (req.PageId - 1) * req.PageSize,
	}

	fmt.Println("Request Parameters:", req)

	account, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		fmt.Println("Error fetching accounts:", err)
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	fmt.Println("Fetched accounts:", account)

	ctx.JSON(http.StatusOK, account)
}
