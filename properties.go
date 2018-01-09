package labstack

type (
	Properties map[string]interface{}
)

func (p Properties) Bool(key string) bool {
	if val, ok := p[key]; ok {
		return val.(bool)
	}
	return false
}

func (p Properties) String(key string) string {
	if val, ok := p[key]; ok {
		return val.(string)
	}
	return ""
}

func (p Properties) Float64(key string) float64 {
	if val, ok := p[key]; ok {
		return val.(float64)
	}
	return 0
}
