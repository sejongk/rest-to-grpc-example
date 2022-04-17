package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

//refer to https://golang.cafe/blog/golang-httptest-example.html
func TestGetMovieDetailHandler(t *testing.T) {
	expectedDetail := MovieDetail{
		ID:          0,
		Title:       "The Shawshank Redemption",
		Genre:       "Drama",
		ReleaseYear: 1994,
		RunningTime: 142,
		Director:    "Frank Darabont",
		Stars:       []string{"Tim Robbins", "Morgan Freeman", "Bob Gunton"},
	}

	svr := httptest.NewServer(NewHandler())
	defer svr.Close()

	res, err := http.Get(svr.URL + "/details/" + strconv.Itoa(expectedDetail.ID))
	if err != nil {
		t.Errorf("unable to complete Get request %v", err)
	}
	defer res.Body.Close()

	resBodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("unable to read response  %v", err)
	}

	var actualDetail MovieDetail
	err = json.Unmarshal(resBodyBytes, &actualDetail)
	if err != nil {
		t.Errorf("unable to unmarshal response %v", err)
	}

	assert.Equal(t, expectedDetail, actualDetail)

}

func TestCreateMovieDetailHandler(t *testing.T) {
	expectedDetail := MovieDetail{
		ID:          0,
		Title:       "The Shawshank Redemption",
		Genre:       "Drama",
		ReleaseYear: 1994,
		RunningTime: 142,
		Director:    "Frank Darabont",
		Stars:       []string{"Tim Robbins", "Morgan Freeman", "Bob Gunton"},
	}

	svr := httptest.NewServer(NewHandler())
	defer svr.Close()

	detailBytes, err := json.Marshal(expectedDetail)
	if err != nil {
		t.Errorf("unable to marshal detail struct to json %v", err)
	}
	reqBody := bytes.NewBuffer(detailBytes)
	res, err := http.Post(svr.URL+"/details", "application/json", reqBody)
	if err != nil {
		t.Errorf("unable to complete Post request %v", err)
	}
	defer res.Body.Close()

	resBodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("unable to read response  %v", err)
	}

	var actualDetail MovieDetail
	err = json.Unmarshal(resBodyBytes, &actualDetail)
	if err != nil {
		t.Errorf("unable to unmarshal response %v", err)
	}

	assert.Equal(t, expectedDetail, actualDetail)

}
