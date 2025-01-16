package pys

import (
	"fmt"
	"testing"
)

func TestPOST(t *testing.T) {
	cookie := "weblogger_did=web_5163939243CCBF9; kwpsecproductname=PCLive; did=web_4fae163cdf0a1bdc3f260a323beec71c; merchantId=undefined; userBid=1000389752609; merchantRoleId=0; merchantRoleType=0; kwfv1=PnGU+9+Y8008S+nH0U+0mjPf8fP08f+98f+nLlwnrIP9+Sw/ZFGfzY+eGlGf+f+e4SGfbYP0QfGnLFwBLU80mYGAYYG/PI8/4fwnLlG9+f+nL98eDFG/H9weHF+Ab08ePFGAr7GAHF+nP7+AqI80b0+Bzj8eGUGALA+ArE8/+0wBr=; userType=3; bUserId=1000332207233; userId=4457855415; kuaishou.lbs.merchant_st=ChhrdWFpc2hvdS5sYnMubWVyY2hhbnQuc3QSsAF6s6ZksL1rjWemHH8x8bTJQxSNgDlCC0XUp4RNQW7FiWZgcoFKfm-lizkGw42xx6iUD2H-6X-f_hvP-cpP1-cbJK54cUhAJk5dHSauxF9VcSTKadI7biK7-Fkve6O1OyaNSZSvAVgGZvSAhIF-u73bxT1cV-sMh8hfqeP7kTi7y6rHiZP3FhqnKxOl7jl405NWt5dhLCwZWz9DcOTI_7-3bmzjPGUHWbx0gEwOFGobwRoSGBjONuzjtXmcOCGUu_lfqFsGIiCeiQDlRhIuaLDrrsAHJ6m7gK4zMRaoMKay8_XtVpj_zCgFMAE; kuaishou.lbs.merchant_ph=86fd9e576b73db6739806a265640247c9a24"
	ua := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0"
	headers := S{"Cookie": cookie, "User-Agent": ua}
	api := "https://lbs.kuaishou.com/rest/locallife/live/pc/b/liveInfo"
	response, err := POST(api, headers, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(response.JsonStringify())
}

func TestWauoAPI(t *testing.T) {
	api := "http://120.27.156.188:9001/calculate_dy_wss_url"
	payload := A{"room_id": "7443740741525883682"}
	resp, _ := POST(api, nil, nil, payload)
	fmt.Println(resp.JsonStringify())
}