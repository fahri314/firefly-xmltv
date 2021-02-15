package models

/*ChannelIndex Digiturk API channel index scheme */
type DigiturkChannelIndex struct {
	Init struct {
		Channels []struct {
			Title             string `json:"Title"`
			ChannelNo         int    `json:"ChannelNo"`
			BroadcastDuration int    `json:"BroadcastDuration"`
			WebSiteURL        string `json:"WebSiteUrl"`
			FileName          string `json:"FileName"`
			PageID            int    `json:"PageID"`
			Description       string `json:"Description"`
			ChannelTypeID     int    `json:"ChannelTypeID"`
			EPGCode           string `json:"EPG_Code"`
			ServiceID         string `json:"ServiceID"`
			ThemeID           string `json:"ThemeId"`
			ID                int    `json:"Id"`
		} `json:"channels"`
	} `json:"init"`
}

/*Digiturk API return scheme */
type DigiturkEPG struct {
	Channel struct {
		EPGList []struct {
			ProgramID           interface{} `json:"ProgramId"`
			ChannelID           int         `json:"ChannelId"`
			ProgramName         string      `json:"ProgramName"`
			OrginalName         string      `json:"OrginalName"`
			BroadcastStart      string      `json:"BroadcastStart"`
			BroadcastEnd        string      `json:"BroadcastEnd"`
			BroadcastDuration   int         `json:"BroadcastDuration"`
			PartNo              interface{} `json:"PartNo"`
			HasSubtitle         bool        `json:"HasSubtitle"`
			ProgramLanguage     interface{} `json:"ProgramLanguage"`
			ScreenRatio         string      `json:"ScreenRatio"`
			IsLive              bool        `json:"IsLive"`
			Synopsis            string      `json:"Synopsis"`
			CreatedBy           string      `json:"CreatedBy"`
			CreatedDate         string      `json:"CreatedDate"`
			UpdatedBy           string      `json:"UpdatedBy"`
			LastModifyDate      string      `json:"LastModifyDate"`
			LastIP              interface{} `json:"LastIP"`
			Genre               string      `json:"Genre"`
			Rating              string      `json:"Rating"`
			Year                string      `json:"Year"`
			Actors              string      `json:"Actors"`
			SeriesID            interface{} `json:"SeriesId"`
			SeasonID            interface{} `json:"SeasonId"`
			LongDescription     string      `json:"LongDescription"`
			ServiceRef          string      `json:"ServiceRef"`
			ContentRef          string      `json:"ContentRef"`
			EventID             string      `json:"EventId"`
			ScreenSize          interface{} `json:"ScreenSize"`
			AudioType           interface{} `json:"AudioType"`
			BroadcastTimeStamp  int         `json:"BroadcastTimeStamp"`
			Directors           string      `json:"Directors"`
			ProductionCountries string      `json:"ProductionCountries"`
			MasterProductionID  string      `json:"MasterProductionID"`
			EPGBroadcastID      string      `json:"EPGBroadcastID"`
			ID                  int         `json:"Id"`
		} `json:"ID"`
	} `json:"listings"`
}
