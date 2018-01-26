package geo

import (
	"fmt"
)

func NamedCRS(srid int) map[string]interface{} {
	ps := make(map[string]interface{})
	ps["type"] = "name"
	ps["properties"] = map[string]interface{}{
		"name": fmt.Sprintf("EPSG:%d", srid),
	}
	return ps
}
