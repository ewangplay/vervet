package main

import (
	_ "encoding/json"
	"fmt"

	"github.com/ewangplay/vervet"
)

type DemoAnimalHandler struct {
	*vervet.BaseHandler
}

func (this *DemoAnimalHandler) ProcessFunc(version int, resource string, method string, params map[string]string, body []byte, result map[string]interface{}) error {

	id, has_id := params["id"]

	switch method {
	case "GET":
		return this.GetDemoAnimalInfo(result)

	case "POST":
		if !has_id {
			return fmt.Errorf("missing parameter id")
		}
		return this.AddDemoAnimalInfo(id, body, result)

	case "PUT":
		if !has_id {
			return fmt.Errorf("missing parameter id")
		}
		return this.UpdateDemoAnimalInfo(id, body, result)

	case "DELETE":
		if !has_id {
			return fmt.Errorf("missing parameter id")
		}
		return this.DeleteDemoAnimalInfo(id, result)

	default:
		return fmt.Errorf("unsupported http method: %v", method)
	}

	return nil
}

func (this *DemoAnimalHandler) GetDemoAnimalInfo(result map[string]interface{}) error {
    result["demo"] = "http get: call GetDemoAnimalInfo"

	return nil
}

func (this *DemoAnimalHandler) AddDemoAnimalInfo(id string, body []byte, result map[string]interface{}) error {

	return nil
}

func (this *DemoAnimalHandler) UpdateDemoAnimalInfo(id string, body []byte, result map[string]interface{}) error {
	return nil
}

func (this *DemoAnimalHandler) DeleteDemoAnimalInfo(id string, result map[string]interface{}) error {
	return nil
}
