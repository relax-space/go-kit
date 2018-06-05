package base

import (
	"encoding/json"
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

func Test_JoinMapJsonMessage(t *testing.T) {
	body := `{
		"CustomerCode": "jimu",
		"Data": {
			"ClosingDt": "20180320",
			"BrandCode": "SF",
			"DeliveryList": [
				{
					"TradeID": "T12",
					"ConfirmTime": "2018-03-23T08:08:05Z",
					"DeliveryType": "OS",
					"PackingSKUs": [
						{
							"SkuID": "SAAC54NA0720FRE",
							"Qty": "2"
						}
					],
					"ShippingCompanyName": "申通",
					"ShippingNo": "1209874532",
					"TmallOrderID": "R123456789",
					"TpPlantCode": ""
				}
			]
		}
	}`
	param := make(map[string]*json.RawMessage, 0)
	err := json.Unmarshal([]byte(body), &param)
	test.Ok(t, err)
	newParam := JoinMapJsonRawMessage(param)
	fmt.Println(newParam)
	test.Equals(t, `CustomerCode=jimu&Data={
			"ClosingDt": "20180320",
			"BrandCode": "SF",
			"DeliveryList": [
				{
					"TradeID": "T12",
					"ConfirmTime": "2018-03-23T08:08:05Z",
					"DeliveryType": "OS",
					"PackingSKUs": [
						{
							"SkuID": "SAAC54NA0720FRE",
							"Qty": "2"
						}
					],
					"ShippingCompanyName": "申通",
					"ShippingNo": "1209874532",
					"TmallOrderID": "R123456789",
					"TpPlantCode": ""
				}
			]
		}`, newParam)
}

func Test_JoinMapJsonMessageEncode(t *testing.T) {
	body := `{
		"CustomerCode": "jimu",
		"Data": {
			"ClosingDt": "20180320",
			"BrandCode": "SF",
			"DeliveryList": [
				{
					"TradeID": "T12",
					"ConfirmTime": "2018-03-23T08:08:05Z",
					"DeliveryType": "OS",
					"PackingSKUs": [
						{
							"SkuID": "SAAC54NA0720FRE",
							"Qty": "2"
						}
					],
					"ShippingCompanyName": "申通",
					"ShippingNo": "1209874532",
					"TmallOrderID": "R123456789",
					"TpPlantCode": ""
				}
			]
		}
	}`
	param := make(map[string]*json.RawMessage, 0)
	err := json.Unmarshal([]byte(body), &param)
	test.Ok(t, err)
	newParam := JoinMapJsonRawMessageEncode(param)
	fmt.Println(newParam) //tools http://tool.oschina.net/encode?type=4
	test.Equals(t, `CustomerCode=jimu%26Data=%7B%0A%09%09%09%22ClosingDt%22%3A%20%2220180320%22%2C%0A%09%09%09%22BrandCode%22%3A%20%22SF%22%2C%0A%09%09%09%22DeliveryList%22%3A%20%5B%0A%09%09%09%09%7B%0A%09%09%09%09%09%22TradeID%22%3A%20%22T12%22%2C%0A%09%09%09%09%09%22ConfirmTime%22%3A%20%222018-03-23T08%3A08%3A05Z%22%2C%0A%09%09%09%09%09%22DeliveryType%22%3A%20%22OS%22%2C%0A%09%09%09%09%09%22PackingSKUs%22%3A%20%5B%0A%09%09%09%09%09%09%7B%0A%09%09%09%09%09%09%09%22SkuID%22%3A%20%22SAAC54NA0720FRE%22%2C%0A%09%09%09%09%09%09%09%22Qty%22%3A%20%222%22%0A%09%09%09%09%09%09%7D%0A%09%09%09%09%09%5D%2C%0A%09%09%09%09%09%22ShippingCompanyName%22%3A%20%22%E7%94%B3%E9%80%9A%22%2C%0A%09%09%09%09%09%22ShippingNo%22%3A%20%221209874532%22%2C%0A%09%09%09%09%09%22TmallOrderID%22%3A%20%22R123456789%22%2C%0A%09%09%09%09%09%22TpPlantCode%22%3A%20%22%22%0A%09%09%09%09%7D%0A%09%09%09%5D%0A%09%09%7D`, newParam)
}
