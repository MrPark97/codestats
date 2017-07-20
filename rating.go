package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"io/ioutil"
	"text/template"
)

type RatingChange struct {
	ContestId, Rank, RatingUpdateTimeSeconds, OldRating, NewRating int64
	ContestName, Handle                                            string
}

type Rating []RatingChange

type RatingResponse struct {
	Status string
	Result Rating
}

type RatingChartPoints map[int64][]int64

type RatingTemplateData struct {
	Handles []string
	Points RatingChartPoints
	HandlesLatestIndex,PointsLatestKey int64
}

func (rr *RatingResponse) Get(handle string) {
	var i int

	for rr.Status != "OK" {

		resp, err := http.Get(fmt.Sprintf("http://codeforces.com/api/user.rating?handle=%s", handle))
		if err != nil {
			log.Println("get rating error: ", err)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		err = json.Unmarshal(body, rr)
		if err != nil {
			log.Println("response unmarshalling error: ", err)
		}

		if i>0 {
			time.Sleep(5 * time.Second)
		}

		i++
	}

	return
}

func ratingPage(w http.ResponseWriter, r *http.Request) {
	template_name := "rating.tpl"
	var rs []RatingResponse
	var latest_key int64

	for i := 0; i < len(config.Handles); i++ {
		var rr RatingResponse
		rr.Get(config.Handles[i])
		rs = append(rs, rr)
	}

	rcp := make(map[int64][]int64);

	for i:=0; i<len(config.Handles); i++ {
		for j:=0; j<len(rs[i].Result); j++ {
			key := rs[i].Result[j].ContestId;
			value := rs[i].Result[j].NewRating;

			if key > latest_key {
				latest_key = key
			}

			_, has := rcp[key];

			if !has {
				rcp[key] = make([]int64, len(config.Handles));
			}

			rcp[key][i] = value;
		}
	}

	var data RatingTemplateData

	data.Handles = config.Handles
	data.Points = rcp
	data.HandlesLatestIndex = int64(len(config.Handles))-1
	data.PointsLatestKey = latest_key

	header(w)

	t, err := template.ParseFiles(fmt.Sprintf("./templates/%s", template_name))
	if err != nil {
	    log.Println("template error", err)
	}

	err = t.Execute(w, data)
	if err != nil {
	    log.Println("template print error", err)
	}

	footer(w)
}
