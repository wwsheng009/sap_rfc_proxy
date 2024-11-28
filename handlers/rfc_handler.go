package handlers

import (
	"fmt"
	"net/http"
	"sap_rfc_proxy/config"
	"sap_rfc_proxy/utils"
	"strconv"
	"strings"

	"sap_rfc_proxy/gorfc"

	"github.com/gin-gonic/gin"
)

func RFCCall(c *gin.Context) {
	funcName := c.Query("fname")

	var payload map[string]interface{}
	if err := c.BindJSON(&payload); err != nil {
		utils.Logger.Printf("Invalid request payload: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	conn, err := gorfc.ConnectionFromParams(config.LoadConfig())
	if err != nil {
		utils.Logger.Printf("Connection error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "SAP connection failed"})
		return
	}
	defer conn.Close()
	funDesc, err := conn.GetFunctionDescription(funcName)
	if err != nil {
		utils.Logger.Printf("RFC call error for function %s: %v", funcName, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var payload2 interface{} = payload

	payload2, err = convertMapObject(payload2, funDesc)
	if err != nil {
		utils.Logger.Printf("Failed to convert interface{} to map[string]interface{}")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	payload, _ = payload2.(map[string]interface{})

	result, err := conn.Call(funcName, payload)
	if err != nil {
		utils.Logger.Printf("RFC call error for function %s: %v", funcName, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.Logger.Printf("Function %s called successfully", funcName)
	c.JSON(http.StatusOK, map[string]interface{}{"data": result})
}

func RFCmeta(c *gin.Context) {
	funcName := c.Query("fname")
	conn, err := gorfc.ConnectionFromParams(config.LoadConfig())
	if err != nil {
		utils.Logger.Printf("Connection error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "SAP connection failed"})
		return
	}
	defer conn.Close()
	funDesc, err := conn.GetFunctionDescription(funcName)
	if err != nil {
		utils.Logger.Printf("RFC call error for function %s: %v", funcName, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.Logger.Printf("Function %s called successfully", funcName)
	c.JSON(http.StatusOK, map[string]interface{}{"meta": funDesc})
}

// try to convert the input map to the correct type
// the input is interface{},it's converted from the json object
func convertMapObject(input interface{}, funDesc gorfc.FunctionDescription) (output interface{}, err error) {

	switch in := input.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{})
		in = upperCaseKeys(in)

		for _, param := range funDesc.Parameters {
			if param.Direction == "RFC_EXPORT" {
				continue
			}
			value, ok := in[strings.ToUpper(param.Name)]
			if !ok || value == nil {
				if param.Optional || param.ParameterType == "RFCTYPE_TABLE" {
					continue
				}
				return nil, fmt.Errorf("missing required field: %s", param.Name)
			}

			switch param.ParameterType {
			case "RFCTYPE_TABLE":
				if param.TypeDesc.Fields == nil {
					return nil, fmt.Errorf("invalid table parameter type description: %s", param.Name)
				}
				array, ok := value.([]interface{})
				if !ok {
					return nil, fmt.Errorf("expected array for parameter: %s", param.Name)
				}
				convertedArray := []map[string]interface{}{}
				for _, elem := range array {
					elemMap, ok := elem.(map[string]interface{})
					if !ok {
						return nil, fmt.Errorf("invalid element in table: %s", param.Name)
					}
					convElem, err := convertMapObject(elemMap, gorfc.FunctionDescription{
						Parameters: createFieldParams(param.TypeDesc.Fields),
					})
					if err != nil {
						return nil, err
					}
					convertedArray = append(convertedArray, convElem.(map[string]interface{}))
				}
				result[param.Name] = convertedArray

			case "RFCTYPE_STRUCTURE":
				if param.TypeDesc.Fields == nil {
					return nil, fmt.Errorf("invalid structure parameter type description: %s", param.Name)
				}
				valueMap, ok := value.(map[string]interface{})
				if !ok {
					return nil, fmt.Errorf("expected map for structure: %s", param.Name)
				}
				convValue, err := convertMapObject(valueMap, gorfc.FunctionDescription{
					Parameters: createFieldParams(param.TypeDesc.Fields),
				})
				if err != nil {
					return nil, err
				}
				result[param.Name] = convValue

			case "RFCTYPE_STRING", "RFCTYPE_CHAR":
				strValue, ok := value.(string)
				if !ok {
					strValue = fmt.Sprintf("%v", value)
					value = strValue
					// return nil, fmt.Errorf("expected string for parameter: %s", param.Name)
				}
				result[param.Name] = value
			case "RFCTYPE_INT1", "RFCTYPE_INT2", "RFCTYPE_INT", "RFCTYPE_INT8":
				switch v := value.(type) {
				case string:
					i, err := strconv.ParseInt(v, 10, 64)
					if err != nil {
						return nil, fmt.Errorf("expected int for field: %s", param.Name)
					}
					value = i
				case float64:
					u := int(v)
					value = u
				}
				result[param.Name] = value
			case "RFCTYPE_FLOAT", "RFCTYPE_BCD", "RFCTYPE_DECF16", "RFCTYPE_DECF34":
				switch v := value.(type) {
				case string:
					f, err := strconv.ParseFloat(v, 64)
					if err != nil {
						return nil, fmt.Errorf("expected float for field: %s", param.Name)
					}
					value = f
				}
				result[param.Name] = value
			// Add cases for other simple types if necessary
			default:
				result[param.Name] = value
			}
		}
		return result, nil

	default:
		return nil, fmt.Errorf("unsupported input type: %T", input)
	}

}
func upperCaseKeys(input map[string]interface{}) map[string]interface{} {
	output := make(map[string]interface{})
	for key, value := range input {
		lowerKey := strings.ToUpper(key)
		switch v := value.(type) {
		case map[string]interface{}:
			// Recursively convert nested maps
			output[lowerKey] = upperCaseKeys(v)
		default:
			output[lowerKey] = value
		}
	}
	return output
}

// Helper function to create ParameterDescription from Fields
func createFieldParams(fields []gorfc.FieldDescription) []gorfc.ParameterDescription {
	params := make([]gorfc.ParameterDescription, len(fields))
	for i, field := range fields {
		params[i] = gorfc.ParameterDescription{
			Name:          field.Name,
			ParameterType: field.FieldType,
			NucLength:     field.NucLength,
			UcLength:      field.UcLength,
			Decimals:      field.Decimals,
			TypeDesc:      field.TypeDesc,
			Optional:      true,
		}
	}
	return params
}
