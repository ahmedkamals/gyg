package app

import (
	loggerLib "gyg/middleware/logger"
	"runtime"
	"io/ioutil"
	"gyg/middleware/config"
	"encoding/json"
	"net/http"
	"gyg/middleware/communication"
	httpProtocol "gyg/middleware/communication/protocols/http"
	"sync"
	"strconv"
	"math"
)

type (
	Dispatcher struct {
		logger *loggerLib.Logger
	}

	Tour struct {
		Id      string  `json:"id"`
		Title	string	`json:"title"`
		Price	float64	`json:"price"`
		Currency	string	`json:"currency"`
		IsSpecialOffer	bool	`json:"isSpecialOffer"`
		Rating 	        float64  `json:"rating"`
	}
)

var (
	output chan map[string]map[string]*Tour
)

func NewDispatcher(logger *loggerLib.Logger) *Dispatcher {
	return &Dispatcher{
		logger: logger,
	}
}

func (d *Dispatcher) Run() *Dispatcher {
	runtime.GOMAXPROCS(runtime.NumCPU())

	output = make(chan map[string]map[string]*Tour, 2000)

	return d
}

func (d *Dispatcher) dispatch(searchTerm string) {
	var tours *config.Tours
	var ratings *config.Ratings
	var err error
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		config.Configuration.Backend["search"].Params["query"] = searchTerm
		tours, err = getTours(config.Configuration.Backend["search"])

		if nil != err {
			d.logger.Error("Dispatcher fetch:" + err.Error())
		}

		wg.Done()
	}()

	go func() {
		ratings, err = loadRating();

		if nil != err {
			d.logger.Error("Dispatcher parse:" + err.Error())
		}

		wg.Done()
	}()

	wg.Wait()

	mergedData := map[string]map[string]*Tour{
		"tours": map[string]*Tour{},
	}

	if tours != nil && ratings != nil {
		mergedData, err = mergeData(tours, ratings)

		if err != nil {
			d.logger.Error("Dispatcher merge:" + err.Error())
		}

	}

	output <- mergedData
}

func getTours(endPoint config.APIEndPoint) (*config.Tours, error) {
	request := communication.NewRequest([]byte{}, map[string]interface{}{
		httpProtocol.METHOD:   endPoint.Method,
		httpProtocol.PROTOCOL: endPoint.Protocol,
		httpProtocol.HOST:     endPoint.Host,
		httpProtocol.PORT:     endPoint.Port,
		httpProtocol.PATH:     endPoint.Path,
		httpProtocol.PARAMS:   endPoint.Params,
	})

	response, err := httpProtocol.NewProtocol(&http.Transport{}).Send(request)

	if err != nil || !response.IsSuccessful() {
		return nil, err
	}

	var toursData config.Tours
	err = json.Unmarshal([]byte(response.GetPayload()), &toursData)

	if nil != err {
		return nil, err
	}

	return &toursData, nil
}

func loadRating() (*config.Ratings, error) {
	data, err := ioutil.ReadFile(config.Configuration.RatingFile)

	if nil != err {
		return nil, err
	}

	parsedRating := config.Ratings{}
	err = json.Unmarshal(data, &parsedRating)

	if nil != err {
		return nil, err
	}

	return &parsedRating, nil
}

func mergeData(tours *config.Tours, ratings *config.Ratings) (map[string]map[string]*Tour, error) {
	data := map[string]map[string]*Tour{
		"tours": map[string]*Tour{},
	}

	toursData := data["tours"]
	ratingTrack := map[string]int{}

	for _, tour := range tours.Items {
		id := strconv.Itoa(tour.Id)

		toursData[id] = &Tour{
			Id: id,
			Title: tour.Title,
			Price: tour.Price,
			Currency: tour.Currency,
			IsSpecialOffer: tour.IsSpecialOffer,
		}
	}

	for _, ratingItem := range *ratings {
		ratingValue, err := strconv.ParseFloat(ratingItem.Rating, 2)
		id := strconv.Itoa(ratingItem.Id)

		if err != nil {
			return nil, err
		}

		if _, ok := toursData[id]; ok {
			ratingTrack[id]++
			(*toursData[id]).Rating += ratingValue
		}
	}

	for id, count := range ratingTrack {
		ratingValue := toursData[id].Rating
		(*toursData[id]).Rating = math.Floor(ratingValue / float64(count))
	}

	return data, nil
}
