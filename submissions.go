package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"io/ioutil"
	"text/template"
	"sort"
)

type Member struct {
	handle string
}

type Problem struct {
	ContestId int64
	Index,Name,Type string
	Points float64
	Tags []string
}

type Party struct {
	TeamId, Room, StartTimeSeconds, ContestId int64
	Members []Member
	ParticipantType, TeamName string
	Ghost bool
}

type Submission struct {
	Id, ContestId, CreationTimeSeconds,RelativeTimeSeconds, PassedTestCount, TimeConsumedMillis, MemoryConsumedBytes int64
	ProgrammingLanguage, Verdict,Testset string
	Problem Problem
	Author Party
}

type SubmissionsResponse struct {
	Status string
	Result []Submission
}

type SubmissionStats struct {
	Handle string
	FailedCount, SuccessCount int64
}

type SubmissionsTemplateData struct {
	Stats []SubmissionStats
	StatsLatestIndex int64
}

type BySuccessCount []SubmissionStats

func (a BySuccessCount) Len() int           { return len(a) }
func (a BySuccessCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySuccessCount) Less(i, j int) bool { return a[i].SuccessCount > a[j].SuccessCount }

func (sr *SubmissionsResponse) Get(handle string) {
	var i int

	for sr.Status != "OK" {

		resp, err := http.Get(fmt.Sprintf("http://codeforces.com/api/user.status?handle=%s", handle))
		if err != nil {
			log.Println("get submissions error: ", err)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		err = json.Unmarshal(body, sr)
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

func submissionsPage(w http.ResponseWriter, r *http.Request) {
	template_name := "submissions.tpl"
	var ss []SubmissionsResponse

	for i := 0; i < len(config.Handles); i++ {
		var sr SubmissionsResponse
		sr.Get(config.Handles[i])
		ss = append(ss, sr)
		log.Println(config.Handles[i])
	}

	sss := make([]SubmissionStats, len(config.Handles))
	var c int64

	for i:=0; i<len(config.Handles); i++ {
		c = 0
		for j:=0; j<len(ss[i].Result); j++ {
			if ss[i].Result[j].Verdict == "OK" {
				c++
			}
		}

		sss[i].Handle = config.Handles[i]
		sss[i].FailedCount = int64(len(ss[i].Result))-c
		sss[i].SuccessCount = c

	}

	sort.Sort(BySuccessCount(sss))

	var data SubmissionsTemplateData

	data.Stats = sss
	data.StatsLatestIndex = int64(len(sss))-1

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
