package main

import (
	. "CSDN/pys"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	type Data struct {
		Scene      string `json:"scene"`
		Size       int    `json:"size"`
		SearchText string `json:"search_text"`
		Cursor     int    `json:"cursor"`
		Extra      S      `json:"extra"`
		Filters    A      `json:"filters"`
	}

	extra := S{
		"new_session_strategy": "1",
		"search_id":            "",
		"session_id":           "",
	}
	filters := A{}
	d := Data{
		Scene:      "PCSquareFeed",
		Size:       30,
		SearchText: "",
		Cursor:     0,
		Extra:      extra,
		Filters:    filters,
	}
	bs, _ := json.Marshal(d)
	data := bytes.NewBuffer(bs)

	// OK
	//v := `{"scene":"PCSquareFeed","size":30,"search_text":"","cursor":0,"extra":{"new_session_strategy":"1","search_id":"","session_id":""},"filters":{}}`
	//var data = strings.NewReader(v)

	req, err := http.NewRequest("POST", "https://buyin.jinritemai.com/pc/selection/common/material_list?verifyFp=verify_m3xx0xij_15642609_be7b_9c0f_f38a_5d6a04d96b9e&fp=verify_m3xx0xij_15642609_be7b_9c0f_f38a_5d6a04d96b9e&msToken=QOgVkJf6vSpqg-Cz-WUYrZSflIEzHdxgyPvduSrep1vuB5V1io44NTURr5gGYR5PkP1rEWi94XqmLYSCVkvWuFtAtHHBagNo-Hix30U-1ox-tpFh3BH3WnS3dZaTOZrViRJhXrAuqDlMKMANhb3oPvSfQd5OZUCshWneqEmq8qvQhYxY4nLSOKWK&a_bogus=DXUjDz7wx28RCplbmKnl94rlDKfArTuyl-iQbFZTSqzzTw0TQdpYDbAnnqO1BBmeLSp3we1HmnWuYdfa%2FdQYgHrkumpkussboT5nI08oMqw4Tt4mDHjMCw0zKwsxUcsqe%2F5tiIUI8UJ9gjnAkqdD%2FdA9yKoKQR8BF1xjkZYbP9s6ZzLAL1c-PQSDNhTCUHKf", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6,zh-TW;q=0.5")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", "store-region-src=uid; s_v_web_id=verify_m3xx0xij_15642609_be7b_9c0f_f38a_5d6a04d96b9e; MONITOR_WEB_ID=c5026176-cd58-461d-b7be-df8a494e4bdd; is_staff_user=false; store-region=cn-hb; passport_auth_status=60554ca32f48bebe297ac39f53c2a7b2%2C87587381daa87297ce57a13b52a9b798; passport_auth_status_ss=60554ca32f48bebe297ac39f53c2a7b2%2C87587381daa87297ce57a13b52a9b798; passport_csrf_token=649c0911da847dede5da37a59de2ee51; passport_csrf_token_default=649c0911da847dede5da37a59de2ee51; _tea_utm_cache_3813=undefined; scmVer=1.0.1.8429; qc_tt_tag=0; gfkadpd=2631,22740; msToken=7ebKd0VuWEVgc6H2s0jwTecNR5DFVWz9HAI-04T10CSDdHXtLSvmXQNVxWBdkkcqb9vazJ9a4vwhQffUz-j98gOD75HiJKRA1F1ehsElGOLAhZU2IyuvWSKungfojv1NtY5l; tt_scid=1..RkAF9TOcVdwn356TLejY47wh-ROCYIKWnTbUEw8BHtHI9ac8Hk31t7U6fTuT991dd; ttwid=1%7C2puL5DsL7IecvM3wIif_KG66FQBe6vrwuhPF32J2LTo%7C1736849473%7Ce807166c67d440daa75e16a6847fbd603499666f86d546e2766277f11c29feef; uid_tt=6af3fae5ef61aff4c77dc3791199a5c2; uid_tt_ss=6af3fae5ef61aff4c77dc3791199a5c2; sid_tt=74130090eac3fd4f91090bb56423c139; sessionid=74130090eac3fd4f91090bb56423c139; sessionid_ss=74130090eac3fd4f91090bb56423c139; odin_tt=44eca532c47b3413010ab617db88e807a07232bb3482673127c1e165f013c9adc057719022377718ea453e258dd98ad7848be5e1a282138c49bee23bc432acaa; ucas_c0_buyin=CkEKBTEuMC4wEICIh7TIh47DZxi9LyC865GgtMyjByiPETC8-uCgtMzXBEDF8Ji8BkjFpNW-BlCzvN-Si9SynWdYfhIUoJ6T1ELIprbfAvMHCok0LhJulqY; ucas_c0_ss_buyin=CkEKBTEuMC4wEICIh7TIh47DZxi9LyC865GgtMyjByiPETC8-uCgtMzXBEDF8Ji8BkjFpNW-BlCzvN-Si9SynWdYfhIUoJ6T1ELIprbfAvMHCok0LhJulqY; SASID=SID2_7459707794822103333; BUYIN_SASID=SID2_7459707794822103333; buyin_shop_type=24; buyin_account_child_type=1128; buyin_app_id=1128; buyin_shop_type_v2=24; buyin_account_child_type_v2=1128; buyin_app_id_v2=1128; sid_guard=74130090eac3fd4f91090bb56423c139%7C1736849540%7C5184000%7CSat%2C+15-Mar-2025+10%3A12%3A20+GMT; sid_ucp_v1=1.0.0-KGMzNTg3NGFiMjQyN2UzMWE0MDVkNmM0ZDRhYWFhZDZkMGQxYmNlNGYKGAi8-uCgtMzXBBCE8Zi8BhiPESAMOAhAJhoCbGYiIDc0MTMwMDkwZWFjM2ZkNGY5MTA5MGJiNTY0MjNjMTM5; ssid_ucp_v1=1.0.0-KGMzNTg3NGFiMjQyN2UzMWE0MDVkNmM0ZDRhYWFhZDZkMGQxYmNlNGYKGAi8-uCgtMzXBBCE8Zi8BhiPESAMOAhAJhoCbGYiIDc0MTMwMDkwZWFjM2ZkNGY5MTA5MGJiNTY0MjNjMTM5; COMPASS_LUOPAN_DT=session_7459708153870074131; csrf_session_id=a1b8ff09a40a3eb0e14fa2586bd492f2")
	req.Header.Set("referer", "https://buyin.jinritemai.com/dashboard/merch-picking-library?pre_universal_page_params_id=&universal_page_params_id=aea4cb59-0ce9-4852-8891-0a6ecd6b8068")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
