package requestx

import "testing"

func TestUrlDecode(t *testing.T) {
}

func TestUrlEncode(t *testing.T) {
	t.Log(UrlEncode("wd=世界&sa=fyb_n_homepage&rsv_dl=fyb_n_homepage&from=super&cl=3&tn=baidutop10&fr=top1000&rsv_idx=2&hisfilter=1"))
}
