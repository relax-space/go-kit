package base

import (
	"fmt"
	"testing"

	"github.com/relax-space/go-kit/test"
)

func Test_JoinMapString(t *testing.T) {

	param := make(map[string]string, 0)
	param["app_id"] = `wx1"234`
	param["name"] = "query"

	newParam := JoinMapString(param)
	test.Equals(t, "app_id=wx1\"234&name=query", newParam)
}

func Test_JoinMapStringEncode(t *testing.T) {

	param := make(map[string]string, 0)
	param["app_id"] = `wx1"234`
	param["name"] = "query"

	newParam := JoinMapStringEncode(param)
	test.Equals(t, "app_id=wx1%22234&name=query", newParam)
}

func Test_JoinMapObject(t *testing.T) {

	param := make(map[string]interface{}, 0)
	param["name"] = "query"
	param["age"] = 18
	param["app_id"] = `wx1"234`

	newParam := JoinMapObject(param)
	test.Equals(t, "age=18&app_id=wx1\"234&name=query", newParam)
}

func Test_JoinMapObjectEncode(t *testing.T) {

	param := make(map[string]interface{}, 0)
	param["name"] = "query"
	param["age"] = 18
	param["app_id"] = `wx1"234`

	newParam := JoinMapObjectEncode(param)
	test.Equals(t, "age=18&app_id=wx1%22234&name=query", newParam)
}

func Test_ParseMapObjectEncode(t *testing.T) {
	raw := `gmt_create=2017-12-07+11%3A15%3A39&amp;charset=UTF-8&amp;seller_email=eland_pay%40elandsystems.cn&amp;subject=xiaomiao+test+ali&amp;sign=QF4Im5T30zkgCTH029kvyyuPiw51Y9LH%2F4yzDZSIXkPgmVj7vf3TSbScH%2Fv2DW8ad53XusjNvSKHcnxnAQIsJoxqcfg52voJ3aGk2TNKP66%2FL4VRak9dVYRwgWhA0AMYPeqyw%2FzHQGF%2FBOsoRu%2Bkh4JflT5WODPbS3QOBpXqcd0%3D&amp;buyer_id=2088702305824122&amp;invoice_amount=0.01&amp;notify_id=50b1bbc78907f7d891e14f3209ecde0gxe&amp;fund_bill_list=%5B%7B%22amount%22%3A%220.01%22%2C%22fundChannel%22%3A%22PCREDIT%22%7D%5D&amp;notify_type=trade_status_sync&amp;trade_status=TRADE_SUCCESS&amp;receipt_amount=0.01&amp;buyer_pay_amount=0.01&amp;app_id=2015081700219294&amp;sign_type=RSA&amp;seller_id=2088312582701209&amp;gmt_payment=2017-12-07+11%3A15%3A43&amp;notify_time=2017-12-07+11%3A15%3A52&amp;version=1.0&amp;out_trade_no=131712072074192779392994717&amp;total_amount=0.01&amp;trade_no=2017120721001004120213770419&amp;auth_app_id=2015081700219294&amp;buyer_logon_id=xia***%40163.com&amp;point_amount=0.00`
	fmt.Printf("\n%v", ParseMapObjectEncode(raw, "&amp;", "="))
}

func Test_ParseMapObject(t *testing.T) {
	raw := `gmt_create=2017-12-07+11%3A15%3A39&charset=UTF-8&seller_email=eland_pay%40elandsystems.cn&subject=xiaomiao+test+ali&sign=QF4Im5T30zkgCTH029kvyyuPiw51Y9LH%2F4yzDZSIXkPgmVj7vf3TSbScH%2Fv2DW8ad53XusjNvSKHcnxnAQIsJoxqcfg52voJ3aGk2TNKP66%2FL4VRak9dVYRwgWhA0AMYPeqyw%2FzHQGF%2FBOsoRu%2Bkh4JflT5WODPbS3QOBpXqcd0%3D&buyer_id=2088702305824122&invoice_amount=0.01&notify_id=50b1bbc78907f7d891e14f3209ecde0gxe&fund_bill_list=%5B%7B%22amount%22%3A%220.01%22%2C%22fundChannel%22%3A%22PCREDIT%22%7D%5D&notify_type=trade_status_sync&trade_status=TRADE_SUCCESS&receipt_amount=0.01&buyer_pay_amount=0.01&app_id=2015081700219294&sign_type=RSA&seller_id=2088312582701209&gmt_payment=2017-12-07+11%3A15%3A43&notify_time=2017-12-07+11%3A15%3A52&version=1.0&out_trade_no=131712072074192779392994717&total_amount=0.01&trade_no=2017120721001004120213770419&auth_app_id=2015081700219294&buyer_logon_id=xia***%40163.com&point_amount=0.00`
	fmt.Printf("\n%v", ParseMapObject(raw, "&", "="))
}

func Test_10_JoinMap(t *testing.T) {

	param := make(map[string]interface{}, 0)
	param["name"] = "query"
	param["age"] = 18
	param["app_id"] = `wx1"234`

	newParam := JoinMap(param, "", "")
	fmt.Println(newParam)
	test.Equals(t, "age18app_idwx1\"234namequery", newParam)
}

func Test_11_JoinMapEncode(t *testing.T) {

	param := make(map[string]interface{}, 0)
	param["name"] = "query"
	param["age"] = 18
	param["app_id"] = `wx1"234`

	newParam := JoinMapEncode(param, "", "")
	fmt.Println(newParam)

	test.Equals(t, "age18app_idwx1%22234namequery", newParam)
}
