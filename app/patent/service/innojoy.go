package service

import (
	"encoding/json"
	"fmt"
	"go-admin/app/patent/my_config"
	"go-admin/app/patent/service/dto"
	"io/ioutil"
	"strings"
	"sync"
)

const (
	authUrl   = "http://www.innojoy.com/accountAuth.aspx"
	searchUrl = "http://www.innojoy.com/service/patentSearch.aspx"
)

var CurrentInnojoy *InnojoyClient
var innojoyCreateOnce sync.Once

func GetCurrentInnojoy() *InnojoyClient {
	innojoyCreateOnce.Do(func() {
		CurrentInnojoy = newInnojoyClient()
	})
	return CurrentInnojoy
}

type callback func() error

type InnojoyClient struct {
	email    string
	password string

	hc    *httpClient
	token string
}

func newInnojoyClient() *InnojoyClient {
	return &InnojoyClient{
		email:    my_config.CurrentPatentConfig.InnojoyUser,
		password: my_config.CurrentPatentConfig.InnojoyPassword,
		hc:       newHttpClient(),
	}
}

func (ic *InnojoyClient) autoLogin() error {
	req := &loginReq{UserConfig: UserConfig{
		EMail:    ic.email,
		Password: ic.password,
	}}

	resp, err := ic.hc.Post(authUrl, req, nil)
	if err != nil {
		return err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	loginRes := loginResp{}
	if err = json.Unmarshal(buf, &loginRes); err != nil {
		return err
	} else if loginRes.ReturnValue != 0 {
		return fmt.Errorf("login to innojoy failed: %s", loginRes.ErrorInfo)
	}

	ic.token = strings.Split(strings.Split(loginRes.Option, ":\"")[1], "\"")[0]

	return nil
}

func (ic *InnojoyClient) SimpleSearch(query string, db string) (result []*dto.PatentDetail, err error) {
	req := ic.parseSimpleSearchQuery(query, db)
	return ic.search(req, ic.autoLogin)
}

func (ic *InnojoyClient) parseSimpleSearchQuery(query string, db string) *SearchReq {
	queryFormat := fmt.Sprintf("TI='%s'", query)
	return &SearchReq{
		Token: ic.token,
		PatentSearchConfig: &PatentSearchConfig{
			GUID:      "",
			Action:    "Search",
			Query:     queryFormat,
			Database:  db,
			Page:      "1",
			PageSize:  "100",
			Sortby:    "-公开（公告）日,公开（公告）号",
			FieldList: "TI,AN,AD,PNM,PD,PA,PINN,CL",
		},
	}
}

func (ic *InnojoyClient) search(sr *SearchReq, cb callback) (result []*dto.PatentDetail, err error) {
	var retried bool
	for {
		resp, err := ic.hc.Post(searchUrl, sr, nil)
		if err != nil {
			return nil, err
		}

		buf, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		searchRes := dto.InnojoySearchResult{}
		if err = json.Unmarshal(buf, &searchRes); err != nil {
			return nil, err
		}

		if searchRes.ReturnValue != 0 {
			if retried {
				return nil, fmt.Errorf("patent search failed: %s", searchRes.ErrorInfo)
			}
			if err = cb(); err != nil {
				return nil, fmt.Errorf("seatch call callback error: %w", err)
			}
			sr.Token = ic.token
			retried = true
		} else {
			return searchRes.Option.PatentList, nil
		}
	}
}

type SearchReq struct {
	Token              string              `json:"token"`
	PatentSearchConfig *PatentSearchConfig `json:"patentSearchConfig"`
}

type PatentSearchConfig struct {
	GUID      string `json:"GUID"`
	Action    string `json:"Action"`
	Query     string `json:"Query"`
	Database  string `json:"Database"`
	Page      string `json:"Page"`
	PageSize  string `json:"PageSize"`
	Sortby    string `json:"Sortby"`
	FieldList string `json:"FieldList"`
}

type loginReq struct {
	UserConfig UserConfig `json:"userConfig"`
}

type UserConfig struct {
	EMail    string `json:"EMail"`
	Password string `json:"Password"`
}

type loginResp struct {
	ReturnValue int    `json:"ReturnValue"`
	Option      string `json:"Option"`
	ErrorInfo   string `json:"ErrorInfo"`
}
