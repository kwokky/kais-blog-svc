package http

import (
	"log"
	"strings"
)

// GetParam
func (c *Context) GetParam(key string) string {
	conType := strings.ToLower(c.GetHeader("Content-Type"))

	if strings.Contains(conType, "application/json") {
		return getJsonParam(c, key).(string)
	}
	if strings.Contains(conType, "multipart/form-data") || strings.Contains(conType, "application/x-www-form-urlencoded") {
		val, ok := c.GetPostForm(key)
		if ok {
			return val
		}
	}

	val, ok := c.GetQuery(key)
	if ok {
		return val
	}

	return ""
}

// GetParamArray
func (c *Context) GetParamArray(key string) []string {
	conType := strings.ToLower(c.GetHeader("Content-Type"))

	if strings.Contains(conType, "application/json") {
		return getJsonParam(c, key).([]string)
	}
	if strings.Contains(conType, "multipart/form-data") || strings.Contains(conType, "application/x-www-form-urlencoded") {
		arrayVal, ok := c.GetPostFormArray(key)
		if ok {
			return arrayVal
		}
	}

	val, ok := c.GetQueryArray(key)
	if ok {
		return val
	}

	return nil
}

// GetParamMap
func (c *Context) GetParamMap(key string) map[string]string {
	conType := strings.ToLower(c.GetHeader("Content-Type"))

	if strings.Contains(conType, "application/json") {
		return getJsonParam(c, key).(map[string]string)
	}
	if strings.Contains(conType, "multipart/form-data") || strings.Contains(conType, "application/x-www-form-urlencoded") {
		mapVal, ok := c.GetPostFormMap(key)
		if ok {
			return mapVal
		}
	}

	mapVal, ok := c.GetQueryMap(key)
	if ok {
		return mapVal
	}

	return nil
}

func getJsonParam(c *Context, key string) interface{} {
	json := make(map[string]interface{})
	err := c.BindJSON(&json)
	if err != nil {
		log.Printf("JSON参数解析失败 %s %s", json, err)
		return nil
	}

	if val, ok := json[key]; ok {
		return val
	}
	return nil
}
