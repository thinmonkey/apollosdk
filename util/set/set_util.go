package set

func DifferenceMapKeys(preProperties map[string]interface{}, curProperties map[string]interface{}) []string {
	diffKeys := make([]string, 10)
	for key := range curProperties {
		if _, ok := preProperties[key]; !ok {
			diffKeys = append(diffKeys, key)
		}
	}
	for key := range preProperties {
		if _, ok := curProperties[key]; !ok {
			diffKeys = append(diffKeys, key)
		}
	}
	return diffKeys
}

func IntersectionMapKeys(preProperties map[string]interface{}, curProperties map[string]interface{}) []string {
	commKeys := make([]string, 0)
	for key := range preProperties {
		for key2 := range curProperties {
			if key == key2 {
				commKeys = append(commKeys, key)
				break
			}
		}
	}
	return commKeys
}

func Intersection(pre []string, cur []string) []string {
	if pre == nil || cur == nil {
		return nil
	}
	preMap := make(map[string]interface{}, len(pre))
	curMap := make(map[string]interface{}, len(cur))
	for index, value := range pre {
		preMap[value] = index
	}
	for index, value := range cur {
		curMap[value] = index
	}
	return IntersectionMapKeys(preMap, curMap)
}

func Difference(pre []string, cur []string) []string {
	if pre == nil || cur == nil {
		return nil
	}
	preMap := make(map[string]interface{}, len(pre))
	curMap := make(map[string]interface{}, len(cur))
	for index, value := range pre {
		preMap[value] = index
	}
	for index, value := range cur {
		curMap[value] = index
	}
	return DifferenceMapKeys(preMap, curMap)
}
