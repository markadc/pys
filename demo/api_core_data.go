package main

import (
	"CSDN/pys"
	"fmt"
)

func main() {
	url := "https://compass.jinritemai.com/compass_api/author/live/live_screen/core_data"
	params := pys.S{
		"room_id":        "7456706368301599507",
		"index_selected": "gpm,pay_ucnt,pay_combo_cnt,watch_pay_ucnt_ratio,product_click_pay_ucnt_ratio,online_user_cnt,live_show_watch_cnt_ratio,avg_watch_duration,watch_interact_ucnt_ratio,follow_anchor_ucnt",
		"a_bogus": "x",
	}

	headers := pys.S{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0",
		"Cookie":     "store-region-src=uid; s_v_web_id=verify_m3mpv4jf_SvZQsRfW_z1tG_4Qa4_BXpR_KAv7trxhisBQ; is_staff_user=false; store-region=cn-hb; passport_auth_status=60554ca32f48bebe297ac39f53c2a7b2%2C87587381daa87297ce57a13b52a9b798; passport_auth_status_ss=60554ca32f48bebe297ac39f53c2a7b2%2C87587381daa87297ce57a13b52a9b798; passport_csrf_token=649c0911da847dede5da37a59de2ee51; passport_csrf_token_default=649c0911da847dede5da37a59de2ee51; x-jupiter-uuid=17364034745062352; csrf_session_id=a1b8ff09a40a3eb0e14fa2586bd492f2; Hm_lvt_b6520b076191ab4b36812da4c90f7a5e=1736476875; HMACCOUNT=E5765DD71A5EB962; qc_tt_tag=0; tt_scid=s5phGOL9w9CXOPkqG.u2o9Op5qHt1v2pnNjzkiVJldd9PIr5RPGyEYPEF0i4unf1fd51; Hm_lpvt_b6520b076191ab4b36812da4c90f7a5e=1736842312; ttwid=1%7C2puL5DsL7IecvM3wIif_KG66FQBe6vrwuhPF32J2LTo%7C1736849473%7Ce807166c67d440daa75e16a6847fbd603499666f86d546e2766277f11c29feef; uid_tt=6af3fae5ef61aff4c77dc3791199a5c2; uid_tt_ss=6af3fae5ef61aff4c77dc3791199a5c2; sid_tt=74130090eac3fd4f91090bb56423c139; sessionid=74130090eac3fd4f91090bb56423c139; sessionid_ss=74130090eac3fd4f91090bb56423c139; odin_tt=44eca532c47b3413010ab617db88e807a07232bb3482673127c1e165f013c9adc057719022377718ea453e258dd98ad7848be5e1a282138c49bee23bc432acaa; BUYIN_SASID=SID2_7459707794822103333; ucas_c0_compass=CkEKBTEuMC4wEJqIiLjXj47DZxi9LyC865GgtMyjByiPETC8-uCgtMzXBECE8Zi8BkiEpdW-BlCzvN-Si9SynWdYfhIUcXLpeS42Pc9xigbZGQCXCJf8KLM; ucas_c0_ss_compass=CkEKBTEuMC4wEJqIiLjXj47DZxi9LyC865GgtMyjByiPETC8-uCgtMzXBECE8Zi8BkiEpdW-BlCzvN-Si9SynWdYfhIUcXLpeS42Pc9xigbZGQCXCJf8KLM; sid_guard=74130090eac3fd4f91090bb56423c139%7C1736849540%7C5184000%7CSat%2C+15-Mar-2025+10%3A12%3A20+GMT; sid_ucp_v1=1.0.0-KGMzNTg3NGFiMjQyN2UzMWE0MDVkNmM0ZDRhYWFhZDZkMGQxYmNlNGYKGAi8-uCgtMzXBBCE8Zi8BhiPESAMOAhAJhoCbGYiIDc0MTMwMDkwZWFjM2ZkNGY5MTA5MGJiNTY0MjNjMTM5; ssid_ucp_v1=1.0.0-KGMzNTg3NGFiMjQyN2UzMWE0MDVkNmM0ZDRhYWFhZDZkMGQxYmNlNGYKGAi8-uCgtMzXBBCE8Zi8BhiPESAMOAhAJhoCbGYiIDc0MTMwMDkwZWFjM2ZkNGY5MTA5MGJiNTY0MjNjMTM5; LUOPAN_DT=session_7459708153870074131; COMPASS_LUOPAN_DT=session_7459708153870074131",
	}

	res, err := pys.Get(url, headers, params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.JsonStringify())

}
