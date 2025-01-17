package spider

import (
	"fmt"
	"testing"
)

func TestResponse_JsonStringify(t *testing.T) {
	username := "MarkAdc"
	url := "https://blog.csdn.net/community/home-api/v1/get-business-list"
	head := map[string]string{
		"Cookie":     "uuid_tt_dd=10_4540732070-1731034040964-347962; fid=20_35293178015-1731034039789-794412; c_segment=10; Hm_lvt_6bcd52f51e9b3dce32bec4a3997715ac=1731034041,1731050076,1731050894,1731306786; HMACCOUNT=E5765DD71A5EB962; loginbox_strategy=%7B%22taskId%22%3A349%2C%22abCheckTime%22%3A1731306785662%2C%22version%22%3A%22exp11%22%2C%22blog-threeH-dialog-exp11tipShowTimes%22%3A3%2C%22blog-threeH-dialog-exp11%22%3A1731320671118%7D; UserName=MarkAdc; UserInfo=a18bead0e11948dd854ab5e684d079f4; UserToken=a18bead0e11948dd854ab5e684d079f4; UserNick=%E6%98%AF%E5%A4%A7%E5%98%9F%E5%98%9F%E5%91%80; AU=A02; UN=MarkAdc; BT=1731320893219; p_uid=U010000; csdn_newcert_MarkAdc=1; MarkAdccomment_new=1725238383974; SESSION=3d74e5c2-4a93-4a4a-8055-16ccb6675f73; redpack_close_time=1; dc_sid=41e9007de0d17e59fc6c963a9731a5db; Hm_ct_6bcd52f51e9b3dce32bec4a3997715ac=6525*1*10_4540732070-1731034040964-347962!5744*1*MarkAdc; FCNEC=%5B%5B%22AKsRol_4Oa7EufN2fxLQswluq8sAMT40XcsENawufa-B9C3pqX3arcJnPioYAwa9AmLnHXHNDzMjbyFtYC9YunsluKflaAnA1WCqQTK5JuKggrv3VXkIrkfEipMBNjiebYPTvFNsO8s0Mv0ZPnJAmvus7Z10H6JMcg%3D%3D%22%5D%5D; https_waf_cookie=c9b7065a-30d5-47538fae3ab7d727019404caa7ad1fa8fed9; _clck=1jlqcjx%7C2%7Cfr2%7C0%7C1776; _clsk=12hok43%7C1732177775782%7C3%7C0%7Cx.clarity.ms%2Fcollect; firstDie=1; __gads=ID=c2fcdfcc577969e6:T=1731034073:RT=1732257094:S=ALNI_MZw1UM5mSdwNoRvC7tvJ1zQqtHhtA; __gpi=UID=00000f7d48674d66:T=1731034073:RT=1732257094:S=ALNI_Mb5Nc_XyGh7U0NrPOMLgkDYPpdgEw; __eoi=ID=9ee5b5de4bb1dc00:T=1731034073:RT=1732257094:S=AA-AfjY0UJ4SEBo_hvnS-lURGX2n; dc_session_id=10_1732261395603.306341; c_pref=default; c_ref=default; c_first_ref=default; c_first_page=https%3A//blog.csdn.net/MarkAdc%3Fspm%3D1000.2115.3001.5343; creativeSetApiNew=%7B%22toolbarImg%22%3A%22https%3A//img-home.csdnimg.cn/images/20230921102607.png%22%2C%22publishSuccessImg%22%3A%22https%3A//img-home.csdnimg.cn/images/20240229024608.png%22%2C%22articleNum%22%3A122%2C%22type%22%3A2%2C%22oldUser%22%3Atrue%2C%22useSeven%22%3Afalse%2C%22oldFullVersion%22%3Atrue%2C%22userName%22%3A%22MarkAdc%22%7D; waf_captcha_marker=708354e3fae2e0cbf84e1de39d8c743c1710140ec675b56bc98ef3be83536a53; c_dsid=11_1732261470576.544448; c_page_id=default; log_Id_pv=2; creative_btn_mp=2; Hm_lpvt_6bcd52f51e9b3dce32bec4a3997715ac=1732261471; dc_tos=sncdkk; log_Id_view=36",
		"Referer":    fmt.Sprintf("https://blog.csdn.net/%v", username),
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0",
	}
	pms := map[string]string{
		"page":         "1",
		"size":         "20",
		"businessType": "lately",
		"noMore":       "false",
		"username":     username,
	}
	resp, err := Get(url, head, pms)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.JsonStringify())
}
