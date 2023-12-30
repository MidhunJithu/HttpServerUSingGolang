package models

type Event struct {
	Event          string `json:"ev"`
	Trigger        string `json:"et"`
	Id             string `json:"id"`
	UserId         string `json:"uid"`
	MessageId      string `json:"mid"`
	Title          string `json:"t"`
	Page           string `json:"p"`
	Language       string `json:"l"`
	ScreenSize     string `json:"sc"`
	AttrKey1       string `json:"atrk1"`
	AttrValue1     string `json:"atrv1"`
	AttrType1      string `json:"atrt1"`
	AttrKey2       string `json:"atrk2"`
	AttrValue2     string `json:"atrv2"`
	AttrType2      string `json:"atrt2"`
	UserTraitKey1  string `json:"uatrk1"`
	UserTraitVal1  string `json:"uatrv1"`
	UserTraitTYpe1 string `json:"uatrt1"`
	UserTraitKey2  string `json:"uatrk2"`
	UserTraitVal2  string `json:"uatrv2"`
	UserTraitTYpe2 string `json:"uatrt2"`
	UserTraitKey3  string `json:"uatrk3"`
	UserTraitVal3  string `json:"uatrv3"`
	UserTraitTYpe3 string `json:"uatrt3"`
}

type StrcturedEvent struct {
	Event      string                 `json:"event"`
	Trigger    string                 `json:"event_type"`
	Id         string                 `json:"app_id"`
	UserId     string                 `json:"user_id"`
	MessageId  string                 `json:"message_id"`
	Title      string                 `json:"page_title"`
	Page       string                 `json:"page_url"`
	Language   string                 `json:"browser_language"`
	ScreenSize string                 `json:"screen_size"`
	Attributes map[string]interface{} `json:"attributes"`
	Traits     map[string]interface{} `json:"traits"`
}
