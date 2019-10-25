package vod

import (
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/BeAce/vcloud-sdk-golang/base"
)

const (
	UPDATE_INTERVAL = 10
)

type Vod struct {
	*base.Client
	DomainCache map[string]map[string]int
	Lock        sync.RWMutex
}

var Instance *Vod
var once sync.Once

func NewInstance() *Vod {
	once.Do(func() {
		Instance = &Vod{
			DomainCache: make(map[string]map[string]int),
			Client:      base.NewClient(ServiceInfoMap[base.RegionCnNorth1], ApiInfoList),
		}
	})
	return Instance
}

func NewInstanceWithRegion(region string) *Vod {
	var serviceInfo *base.ServiceInfo
	var ok bool
	if serviceInfo, ok = ServiceInfoMap[region]; !ok {
		panic("Cant find the region, please check it carefully")
	}

	once.Do(func() {
		Instance = &Vod{
			DomainCache: make(map[string]map[string]int),
			Client:      base.NewClient(serviceInfo, ApiInfoList),
		}
	})
	return Instance
}

var (
	ServiceInfoMap = map[string]*base.ServiceInfo{
		base.RegionCnNorth1: {
			Timeout: 5 * time.Second,
			Host:    "vod.bytedanceapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "vod"},
		},
		base.RegionApSingapore: {
			Timeout: 5 * time.Second,
			Host:    "vod.ap-singapore-1.bytedanceapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionApSingapore, Service: "vod"},
		},
		base.RegionUsEast1: {
			Timeout: 5 * time.Second,
			Host:    "vod.us-east-1.bytedanceapi.com",
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{Region: base.RegionUsEast1, Service: "vod"},
		},
	}

	ServiceInfo = &base.ServiceInfo{
		Timeout: 5 * time.Second,
		Host:    "vod.bytedanceapi.com",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
		Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "vod"},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		"GetPlayInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetPlayInfo"},
				"Version": []string{"2019-03-15"},
			},
		},
		"RedirectPlay": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"RedirectPlay"},
				"Version": []string{"2018-01-01"},
			},
		},
		"GetOriginVideoPlayInfo": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetOriginVideoPlayInfo"},
				"Version": []string{"2018-01-01"},
			},
		},
		"StartTranscode": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"StartTranscode"},
				"Version": []string{"2018-01-01"},
			},
		},
		"UploadMediaByUrl": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"UploadMediaByUrl"},
				"Version": []string{"2018-01-01"},
			},
		},
		"ApplyUpload": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ApplyUpload"},
				"Version": []string{"2018-01-01"},
			},
		},
		"CommitUpload": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CommitUpload"},
				"Version": []string{"2018-01-01"},
			},
		},
		"SetVideoPublishStatus": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SetVideoPublishStatus"},
				"Version": []string{"2018-01-01"},
			},
		},
		"GetCdnDomainWeights": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetCdnDomainWeights"},
				"Version": []string{"2019-07-01"},
			},
		},
		"ModifyVideoInfo": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ModifyVideoInfo"},
				"Version": []string{"2018-01-01"},
			},
		},
	}
)
