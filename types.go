package main

// Resp is response
//
// Example:
// {
//    "code":0,
//    "data":{
//       "country":"\u4e2d\u56fd",
//       "country_id":"CN",
//       "area":"\u534e\u4e1c",
//       "area_id":"300000",
//       "region":"\u6c5f\u82cf\u7701",
//       "region_id":"320000",
//       "city":"\u5357\u4eac\u5e02",
//       "city_id":"320100",
//       "county":"",
//       "county_id":"-1",
//       "isp":"\u7535\u4fe1",
//       "isp_id":"100017",
//       "ip":"117.12.3.72"
//    }
// }
type Loc struct {
	Code int
	Data Data `json:"data"`
}

type Data struct {
	IP        string `json:"ip"`
	Country   string `json:"country"`
	CountryID string `json:"country_id"`
	Area      string `json:"area"`
	AreaID    string `json:"area_id"`
	Region    string `json:"region"`
	RegionID  string `json:"region_id"`
	City      string `json:"city"`
	CityID    string `json:"city_id"`
	County    string `json:"county"`
	CountyID  string `json:"county_id"`
	ISP       string `json:"isp"`
	ISPID     string `json:"isp_id"`
}
