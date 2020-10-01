package main

import (
	"errors"
	"github.com/imroc/req"
)

func globAuthenticate(userID, password string) (string, error) {
	res, err := req.Post("https://applerussiasr.b2xcare.com/crmServiceNodeProd/api/globAuthenticate",
		req.BodyJSON(&globAuthenticateRequest{
			UserID:   userID,
			Password: password,
		}))
	if err != nil {
		return "", err
	}
	var auth globAuthenticateResponse
	err = res.ToJSON(&auth)
	if err != nil {
		return "", err
	}

	if auth.Status != 200 || auth.SessionID == "" {
		return "", errors.New("can't auth")
	}
	return auth.SessionID, nil
}

func repairSummaryLookup(authToken, sessionID, jobNumber string) (string, error) {
	res, err := req.Post("https://applerussiasr.b2xcare.com/crmServiceNodeProd/api/repairSummaryLookup",
		req.Header{
			"authtoken": authToken,
			"sessionid": sessionID,
		},
		req.BodyJSON(&repairSummaryLookupRequest{JobNumber: jobNumber}))
	if err != nil {
		return "", err
	}
	var status repairSummaryLookupResponse
	err = res.ToJSON(&status)
	if err != nil {
		return "", err
	}

	if status.Status != 200 || status.JobDetails.ActionStatus == "" {
		return "", errors.New("can't get status")
	}
	return status.JobDetails.ActionStatus, nil
}

type globAuthenticateRequest struct {
	UserID   string `json:"userId"`
	Password string `json:"password"`
}

type globAuthenticateResponse struct {
	Status    int    `json:"status"`
	SessionID string `json:"sessionId"`
}

type repairSummaryLookupRequest struct {
	SerialNumber string `json:"serialNumber"`
	IMEINumber string `json:"imeiNumber"`
	JobNumber string `json:"jobNumber"`
}

type repairSummaryLookupResponse struct {
	Status     int `json:"status"`
	JobDetails struct {
		ActionStatus string `json:"actionStatus"`
	} `json:"jobDetails"`
}
