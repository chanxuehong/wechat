// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     magicshui(shuiyuzhe@gmail.com)

package datacube

// https://api.weixin.qq.com/datacube/getusersummary?access_token=ACCESS_TOKEN
func dataCubeGetUserSummaryUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/datacube/getusersummary?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/datacube/getusercumulate?access_token=ACCESS_TOKEN
func dataCubeGetUserCumulateUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/datacube/getusercumulate?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/datacube/getarticlesummary?access_token=ACCESS_TOKEN
func dataCubeGetArticleSummaryUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/datacube/getarticlesummary?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/datacube/getarticletotal?access_token=ACCESS_TOKEN
func dataCubeGetArticleTotalUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/datacube/getarticletotal?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/datacube/getuserread?access_token=ACCESS_TOKEN
func dataCubeGetUserReadUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/datacube/getuserread?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/datacube/getuserreadhour?access_token=ACCESS_TOKEN
func dataCubeGetUserReadHourUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/datacube/getuserreadhour?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/datacube/getusershare?access_token=ACCESS_TOKEN
func dataCubeGetUserShareUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/datacube/getusershare?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/datacube/getusersharehour?access_token=ACCESS_TOKEN
func dataCubeGetUserShareHourUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/datacube/getusersharehour?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/datacube/getupstreammsg?access_token=ACCESS_TOKEN
func dataCubeGetUpstreamMsgUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/datacube/getupstreammsg?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/datacube/getupstreammsghour?access_token=ACCESS_TOKEN
func dataCubeGetUpstreamMsgHourUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/datacube/getupstreammsghour?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/datacube/getupstreammsgweek?access_token=ACCESS_TOKEN
func dataCubeGetUpstreamMsgWeekUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/datacube/getupstreammsgweek?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/datacube/getupstreammsgmonth?access_token=ACCESS_TOKEN
func dataCubeGetUpstreamMsgMonthUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/datacube/getupstreammsgmonth?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/datacube/getupstreammsgdist?access_token=ACCESS_TOKEN
func dataCubeGetUpstreamMsgDistUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/datacube/getupstreammsgdist?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/datacube/getupstreammsgdistweek?access_token=ACCESS_TOKEN
func dataCubeGetUpstreamMsgDistWeekUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/datacube/getupstreammsgdistweek?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/datacube/getupstreammsgdistmonth?access_token=ACCESS_TOKEN
func dataCubeGetUpstreamMsgDistMonthUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/datacube/getupstreammsgdistmonth?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/datacube/getinterfacesummary?access_token=ACCESS_TOKEN
func dataCubeGetInterfaceSummaryUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/datacube/getinterfacesummary?access_token=" +
		accesstoken
}

// https://api.weixin.qq.com/datacube/getinterfacesummaryhour?access_token=ACCESS_TOKEN
func dataCubeGetInterfaceSummaryHourUrl(accesstoken string) string {
	return "https://api.weixin.qq.com/datacube/getinterfacesummaryhour?access_token=" +
		accesstoken
}
