package main

import (
	_ "encoding/json"
	"fmt"
)

type DemoHandler struct {
	*BaseHandler
}

func (this *DemoHandler) ProcessFunc(version int, resource string, method string, params map[string]string, body []byte, result map[string]interface{}) error {

	id, has_id := params["id"]

	switch resource {
	case "demo":
		switch method {
		case "GET":
			return this.GetDemoInfo(result)

		case "POST":
			if !has_id {
				return fmt.Errorf("missing parameter id")
			}
			return this.AddDemoInfo(id, body, result)

		case "PUT":
			if !has_id {
				return fmt.Errorf("missing parameter id")
			}
			return this.UpdateDemoInfo(id, body, result)

		case "DELETE":
			if !has_id {
				return fmt.Errorf("missing parameter id")
			}
			return this.DeleteDemoInfo(id, result)

		default:
			return fmt.Errorf("unsupported http method: %v", method)
		}

	default:
		return fmt.Errorf("unsupported request resource: %v", resource)
	}

	return nil
}

func (this *DemoHandler) GetDemoInfo(result map[string]interface{}) error {
    result["demo"] = "http get"

    return nil
}

func (this *DemoHandler) AddDemoInfo(id string, body []byte, result map[string]interface{}) error {

    return nil
}

func (this *DemoHandler) UpdateDemoInfo(id string, body []byte, result map[string]interface{}) error {
    return nil
}

func (this *DemoHandler) DeleteDemoInfo(id string, result map[string]interface{}) error {
    return nil
}
