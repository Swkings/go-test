package test

import (
	"fmt"
	"time"

	"testing"

	"github.com/go-resty/resty/v2"
)

type TaskType interface {
	Order | BackOrder
}

type Order struct {
	OrderId               string  `json:"orderId"`
	MapId                 string  `json:"mapId"`
	Capacity              float64 `json:"capacity"`
	CargoTagId            string  `json:"cargoTagId"`
	Vin                   string  `json:"vin"`
	EstimatedStartTime    int64   `json:"estimatedStartTime"`
	EstimatedCompleteTime int64   `json:"estimatedCompleteTime"`
	Status                int64   `json:"status"`
}

type BackOrder struct {
	BackOrderId       string    `json:"backOrderId"`
	VehicleExternalId string    `json:"vehicleExternalId"`
	MapId             string    `json:"mapId"`
	Reason            int       `json:"reason"`
	State             int       `json:"state"`
	DepotStopId       int32     `json:"depotStopId"`
	IsManualCompleted bool      `json:"isManualCompleted"`
	ErrorCode         int       `json:"errorCode"`
	ErrorMessage      string    `json:"errorMessage"`
	CompletedAt       time.Time `json:"completedAt"`
	AbortedAt         time.Time `json:"abortedAt"`
	CreatedAt         time.Time `json:"createdAt"`
	IsBlockError      bool      `json:"isBlockError"`
}

func (o *BackOrder) GetVin() string {
	return o.VehicleExternalId
}

type Orders[OrderType TaskType] struct {
	Code    int64       `json:"code"`
	Title   string      `json:"title"`
	Message string      `json:"message"`
	Page    int         `json:"page"`
	Size    int         `json:"size"`
	Total   int         `json:"total"`
	List    []OrderType `json:"list"`
}

type GetOrdersBody struct {
	MapId string `json:"mapId"`
	State []int  `json:"state"`
}

func TestGetOrders(t *testing.T) {
	UrlGetOrders := "/order/huadian/internal/sub/:list4multidispatch"
	client := resty.New().SetBaseURL("http://127.0.0.1:28024").SetTimeout(5 * time.Minute)
	resp := Orders[Order]{}
	body := GetOrdersBody{
		MapId: "179",
		State: []int{1, 2, 6},
	}
	_, err := client.R().
		SetBody(body).
		SetResult(&resp).SetError(&resp).
		Post(UrlGetOrders)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(PrettyMapStruct(resp, true))
}
