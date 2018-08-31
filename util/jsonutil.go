package util

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type JsonTime time.Time

var TimeLayout = "2006-01-02T15:04:05"

//MarshalJSON 实现它的json.MarshalJSONer序列化方法
func (jsontime JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(jsontime).Format(TimeLayout))
	return []byte(stamp), nil
}

//String 实现Stringer接口
func (jsontime JsonTime) String() string {
	return time.Time(jsontime).Format(TimeLayout)
}

//Time 转化为time.Time
func (jsontime JsonTime) Time() time.Time {
	return time.Time(jsontime)
}

//Format 实现format接口
func (jsontime JsonTime) Format(layout string) string {
	return time.Time(jsontime).Format(layout)
}

//UnmarshalJSON 实现它的json.Unmarshaler 序列化方法
func (jsontime *JsonTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	str := strings.Trim(string(data), "\"")
	date, err := time.Parse(TimeLayout, str)
	if err == nil {
		*jsontime = JsonTime(date)
	}
	return err
}

//JsonEncode json序列化
func JsonEncode(obj interface{}) (string, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

//JsonDecode Json反序列化
func JsonDecode(data string, obj interface{}) error {
	err := json.Unmarshal([]byte(data), obj)
	return err
}
