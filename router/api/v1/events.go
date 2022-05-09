package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Event struct {
	EventId          int       `json:"event_id"`
	EventName        string    `json:"event_name" example:"helloworld"`
	EventDescription string    `json:"event_description"`
	StartTime        time.Time `json:"start_time"`
	TimeDuration     int       `json:"time_duration"`
}

type EventsResponse struct {
	Code   string `json:"code"`
	Total  int    `json:"total"`
	Events []Event
}

// @Description  all events
// @Tags         event
// @Accept       json
// @Produce      json
// @Success      200  {object} v1.EventsResponse  "GET/api/v1/events"
// @Failure      400  {object}  utils.Error
// @Failure      500  {object}  utils.Error
// @Router       /events [GET]
func AllEvents(ctx *gin.Context) {
	var resp EventsResponse
	resp.Code = "SUCCESS"
	for i := 0; i < 10; i++ {
		resp.Total += 1
		eventTmp := Event{
			EventId:          i,
			EventName:        "name" + strconv.Itoa(i),
			EventDescription: "I am name" + strconv.Itoa(i),
			StartTime:        time.Now(),
			TimeDuration:     1000,
		}
		resp.Events = append(resp.Events, eventTmp)
	}
	ctx.JSON(http.StatusOK, resp)
}

type EventResponse struct {
	Code  string
	Event Event
}

// @Description  single event
// @Tags         event
// @param 		 event-id   path   string    true    "event id"
// @Accept       json
// @Produce      json
// @Success      200  {object}  EventResponse "GET/api/v1/events/balala"
// @Failure      400  {object}  utils.Error
// @Failure      500  {object}  utils.Error
// @Router       /events/{event-id} [GET]
func SingleEvent(ctx *gin.Context) {
	var resp EventResponse
	resp.Code = "SUCCESS"
	reqId := ctx.Param("event-id")
	event_id, _ := strconv.Atoi(reqId)
	resp.Event.EventId = event_id
	resp.Event.EventName = "name" + reqId
	resp.Event.EventDescription = "I am event " + reqId
	resp.Event.StartTime = time.Now()
	resp.Event.TimeDuration = 10000
	ctx.JSON(http.StatusOK, resp)
}

type EventItem struct {
	ItemID         int    `json:"item_id" example:"123455"`
	ItemCollection int    `json:"collection_id" example:"5"`
	OwnerId        string `josn:"owner_id" example:"mazhengwang-ust-hk"`
	Image          string `json:"image" exmaple:"http://www.iamge.com/123455"`
	LocalFavorites int    `josn:"local_favorites" example:"1"`
}

type EventItemsResponse struct {
	Code  string `json:"code" example:"SUCCESS"`
	Total int    `json:"totel" example:"1"`
	Items []EventItem
}

// @Description  items in event
// @Tags         event
// @param 		 event-id   path   string    true    "event id"
// @Accept       json
// @Produce      json
// @Success      200  {object}  EventItemsResponse "GET/api/v1/events/balala/items"
// @Failure      400  {object}  utils.Error
// @Failure      500  {object}  utils.Error
// @Router       /events/{event-id}/items [GET]
func EventItems(ctx *gin.Context) {
	var resp EventItemsResponse
	resp.Code = "SUCCESS"
	for i := 0; i < 10; i++ {
		resp.Total += 1
		EventItemTmp := EventItem{
			ItemID:         i,
			ItemCollection: 1,
			OwnerId:        "mazhengwang-ust-hk",
			Image:          "www.image.com/" + strconv.Itoa(i),
			LocalFavorites: i,
		}
		resp.Items = append(resp.Items, EventItemTmp)
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Description  item ranks in event
// @Tags         event
// @param 		 event-id   path   string    true    "event id"
// @Accept       json
// @Produce      json
// @Success      200  {object}  EventItemsResponse "GET/api/v1/events/balala/ranks"
// @Failure      400  {object}  utils.Error
// @Failure      500  {object}  utils.Error
// @Router       /events/{event-id}/ranks [GET]
func EventItemsRank(ctx *gin.Context) {
	var resp EventItemsResponse
	resp.Code = "SUCCESS"
	for i := 0; i < 10; i++ {
		resp.Total += 1
		ItemTmp := EventItem{
			ItemID:         i,
			ItemCollection: 1,
			OwnerId:        "mazhengwang-ust-hk",
			Image:          "www.image.com/" + strconv.Itoa(i),
			LocalFavorites: 10 - i,
		}
		resp.Items = append(resp.Items, ItemTmp)
	}
	ctx.JSON(http.StatusOK, resp)
}

type EventLikesResponse struct {
	Code string `json:"code" example:"SUCCESS“ example:USER_LIKED,JOIN_NEED,EMPTY_TICKET"`
}

// @Description  user join event
// @Tags         event
// @param 		 event-id   path   int    true    "event id"
// @param        item-id    query  int    true    "item id"
// @Accept       json
// @Produce      json
// @Success      200  {object}  JoinEventResponse "POST/api/v1/events/balala/likes"
// @Failure      400  {object}  utils.Error
// @Failure      500  {object}  utils.Error
// @Router       /events/{event-id}/likes [POST]
func EventLikes(ctx *gin.Context) {
	var resp EventLikesResponse
	resp.Code = "SUCCESS"
	ctx.JSON(http.StatusOK, resp)
}

type JoinEventResponse struct {
	Code string `json:"code" example:"SUCCESS,USER_JOINED"`
}

// @Description  user join event
// @Tags         event
// @param 		 event-id   path   string    true    "event id"
// @Accept       json
// @Produce      json
// @Success      200  {object}  JoinEventResponse "POST/api/v1/events/balala/join"
// @Failure      400  {object}  utils.Error
// @Failure      500  {object}  utils.Error
// @Router       /events/{event-id}/join [POST]
func JoinEvent(ctx *gin.Context) {
	var resp JoinEventResponse
	resp.Code = "SUCCESS"
	ctx.JSON(http.StatusOK, resp)
}

type SubmitItemResponse struct {
	Code string `json:"code" example:"SUCCESS,NO_AUTH,NO_ITEM"`
}

// @Description  user submit item
// @Tags         event
// @param 		 event-id   path   string    true    "event id"
// @param        item-id    query  string    true    "item id"
// @Accept       json
// @Produce      json
// @Success      200  {string}  string "POST/api/v1/events/balala/submit-item"
// @Failure      400  {object}  utils.Error
// @Failure      500  {object}  utils.Error
// @Router       /events/{event-id}/submit-item [POST]
func SubmitItem(ctx *gin.Context) {
	var resp SubmitItemResponse
	resp.Code = "SUCCESS"
	ctx.JSON(http.StatusOK, resp)
}
