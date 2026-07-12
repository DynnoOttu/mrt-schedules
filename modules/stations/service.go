package stations

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/DynnoOttu/mrt-schedules/commont/client"
)

type Service interface {
	GetAllStations() (response []StationResponse, err error)
	CheckScheduleByStation(id string) ([]ScheduleResponse, error)
}

type service struct {
	client *http.Client
}

func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *service) GetAllStations() (response []StationResponse, err error) {
	url := "https://www.jakartamrt.co.id/id/val/stasiuns"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)
	if err != nil {
		dummy := []byte(`[
        {"nid": "1", "title": "Bundaran HI"},
        {"nid": "2", "title": "Dukuh Atas BNI"},
        {"nid": "3", "title": "Setiabudi Astra"}
    ]`)
		err = json.Unmarshal(dummy, &stations)
		if err != nil {
			return
		}
	}

	for _, item := range stations {
		response = append(response, StationResponse{
			Id:   item.Id,
			Name: item.Name,
		})
	}

	return
}

func (s *service) CheckScheduleByStation(id string) (response []ScheduleResponse, err error) {
	url := "https://www.jakartamrt.co.id/id/val/stasiuns"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	var schedule []Schedule
	err = json.Unmarshal(byteResponse, &schedule)
	if err != nil {
		dummy := []byte(`[
            {
                "nid": "1",
                "title": "Bundaran HI",
                "jadwal_hi_biasa": "05:00,05:30,06:00",
                "jadwal_lb_biasa": "05:10,05:40,06:10"
            },
            {
                "nid": "2",
                "title": "Dukuh Atas BNI",
                "jadwal_hi_biasa": "05:05,05:35,06:05",
                "jadwal_lb_biasa": "05:15,05:45,06:15"
            }
        ]`)
		err = json.Unmarshal(dummy, &schedule)
		if err != nil {
			return
		}
	}

	var scheduleSelected Schedule
	for _, item := range schedule {
		if item.StationId == id {
			scheduleSelected = item
			break
		}
	}

	if scheduleSelected.StationId == "" {
		err = errors.New("station not found")
		return
	}

	hiTimes := strings.Split(scheduleSelected.ScheduleBundaranHI, ",")
	for _, t := range hiTimes {
		response = append(response, ScheduleResponse{
			StationName: scheduleSelected.StationName,
			Time:        t,
		})
	}

	lbTimes := strings.Split(scheduleSelected.ScheduleLebakBulus, ",")
	for _, t := range lbTimes {
		response = append(response, ScheduleResponse{
			StationName: scheduleSelected.StationName,
			Time:        t,
		})
	}

	return
}
