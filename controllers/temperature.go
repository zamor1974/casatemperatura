package controllers

import (
	"casatemperatura/lang"
	"casatemperatura/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

// BaseHandler will hold everything that controller needs
type BaseHandlerSqlx struct {
	db *sqlx.DB
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandlerSqlx(db *sqlx.DB) *BaseHandlerSqlx {
	return &BaseHandlerSqlx{
		db: db,
	}
}

// swagger:model CommonError
type CommonError struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the error
	// in: string
	Message string `json:"message"`
}

// swagger:model CommonSuccess
type CommonSuccess struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the error
	// in: string
	Message string `json:"message"`
}

// swagger:model GetTemperatures
type GetTemperatures struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the response
	// in: string
	Message string               `json:"message"`
	Data    *models.Temperatures `json:"data"`
}

// swagger:model GetTemperature
type GetTemperature struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the response
	// in: string
	Message string `json:"message"`
	// Umidity value
	Data *models.Temperature `json:"data"`
}

// ErrHandler returns error message response
func ErrHandler(errmessage string) *CommonError {
	errresponse := CommonError{}
	errresponse.Status = 0
	errresponse.Message = errmessage
	return &errresponse
}

// swagger:route GET /temperatures listTemperature
// Get Temperature list
//
// security:
// - apiKey: []
// responses:
//  401: CommonError
//  200: GetTemperatures
func (h *BaseHandlerSqlx) GetTemperaturesSqlx(w http.ResponseWriter, r *http.Request) {
	response := GetTemperatures{}

	temperatures := models.GetTemperatureSqlx(h.db.DB)

	response.Status = 1
	response.Message = lang.Get("success")
	response.Data = temperatures

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// swagger:route GET /lasthour lastHour
// Get list of last hour of temperature values .... or the last value inserted
//
// security:
// - apiKey: []
// responses:
//  401: CommonError
//  200: GetTemperatures
func (h *BaseHandlerSqlx) GetLastHourSqlx(w http.ResponseWriter, r *http.Request) {
	response := GetTemperatures{}

	temperatures := models.GetLastHourSqlx(h.db.DB)

	response.Status = 1
	response.Message = lang.Get("success")
	response.Data = temperatures

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// swagger:route POST /temperature addTemperature
// Create a new temperature value
//
// security:
// - apiKey: []
// responses:
//  401: CommonError
//  200: GetTemperature
func (h *BaseHandlerSqlx) PostTemperatureSqlx(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	response := GetTemperature{}

	decoder := json.NewDecoder(r.Body)
	var reqTemperature *models.ReqAddTemperature
	err := decoder.Decode(&reqTemperature)
	fmt.Println(err)

	if err != nil {
		json.NewEncoder(w).Encode(ErrHandler(lang.Get("invalid_request")))
		return
	}

	rain, errmessage := models.PostTemperatureSqlx(h.db.DB, reqTemperature)
	if errmessage != "" {
		json.NewEncoder(w).Encode(ErrHandler(errmessage))
		return
	}

	response.Status = 1
	response.Message = lang.Get("insert_success")
	response.Data = rain
	json.NewEncoder(w).Encode(response)
}
