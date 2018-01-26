package geo

type Property interface{}
type PropertyCollection []interface{}

type Properties map[string]Property
type PropertyCollections map[string]PropertyCollection

func (pcs PropertyCollections) AppendProperties(ps Properties) {
	for key, property := range ps {
		if _, ok := pcs[key]; !ok {
			pcs[key] = PropertyCollection{property}
		} else {
			pcs[key] = append(pcs[key], property)
		}
		logger.WithField("key", key).Debug("Appending property key is following.")
		logger.WithField("property", pcs[key]).Debug("Appending property value is following.")
	}
}

func ParseProperties(headers []string, record []string, ex ...string) Properties {
	properties := make(Properties)
	for i, value := range record {
		header := headers[i]
		if contains(header, ex) {
			continue
		}
		properties[header] = value
	}
	return properties
}

func contains(value string, array []string) bool {
	for _, elm := range array {
		if elm == value {
			return true
		}
	}
	return false
}
