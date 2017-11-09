# go-kit
Provide some common methods

    - auth
    - base
    - config
    - data
    - httpreq
    - log
    - mapstruct
    - model
    - random
    - sign
    - test

## auth

NewToken:generate jwt token
```go

	token, err := auth.NewToken(map[string]interface{}{"app_id": "wx1234", "role_id": "1"})
```

## base

JoinMapObject: map => string with "&" 
```go

    param := make(map[string]string, 0)
    param["app_id"] = `wx1"234`
    param["name"] = "query"
    
    newParam := base.JoinMapString(param)//output:app_id=wx12"34&name=query    
```

GetZoneTime: you can get zone time,like utc,china
```go

    now := time.Now()//2017-11-09 13:25:00
    utcTime := base.GetZoneTime(UTCZONE, now)//2017-11-09 13:25:00
    chinaTime := base.GetZoneTime(CHINAZONE, now)//2017-11-09 21:25:00
```

ToString: convert interface{} to string
```go

	var a int64 = 12
	var b int32 = 13
	var c bool = false
	var d float32 = 14
	var e float64 = 15

	fmt.Printf("\nint64:%T-%v", base.ToString(a), a)//int64:string-12
	fmt.Printf("\nint32:%T-%v", base.ToString(b), b)//int32:string-13
	fmt.Printf("\nbool:%T-%v", base.ToString(c), c)//bool:string-false
	fmt.Printf("\nfloat32:%T-%v", base.ToString(d), d)//float32:string-14
	fmt.Printf("\nfloat64:%T-%v", base.ToString(e), e)//float64:string-15
```


XmlToMap:  xml => map
```go
case 1:
	xmlData := `<xml>
	<app_id>wx1111</app_id>
	<name>apple</name>
	</xml>`

    data, err := base.XmlToMap(xmlData)
    fmt.Printf("\n%+v", data["xml"])//map[name:apple app_id:wx1111]

case 2:
    xmlData := `<apple>
	<color>red</color>
	<price>18</price>
	</apple>`
    data, err := base.XmlToMap(xmlData)
    fmt.Printf("\n%+v", data["apple"])//map[name:apple app_id:wx1111]
    
```

XmlMap: map => xml
```go
	xmlMap := base.XmlMap{}
	xmlMap["color"] = "red"
	xmlMap["price"] = 18

    s, err := xml.MarshalIndent(xmlMap, "", " ")
    // <xml>
    //  <color>red</color>
    //  <price>18</price>
    // </xml>
```
## mapstruct

Decode: map=> struct
```go
    var apple struct {
		Color string `json:"color"`
		Price int    `json:"price"`
	}

	appleMap := map[string]interface{}{
		"color": "red",
		"price": 14,
    }
    mapstruct.Decode(appleMap, &apple)//{Color:red Price:14}
```

## sign

GetMD5Hash:
```go
    s, err := sign.GetMD5Hash("secret jwt")
    s1, err := sign.GetMD5Hash("secret jwt", true)
    fmt.Println(s, s1)
    //s:  26F69B1B2FA0F9E8F958337C36F3045D 
    //s1: 26f69b1b2fa0f9e8f958337c36f3045d
```
## random

Uuid: generate uuid
```go

    uuId := random.Uuid("")//9031646797813138653
```





