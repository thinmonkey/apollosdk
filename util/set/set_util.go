package set

/**
找到两个slice的交集元素slice
 */
func Intersection(pre []string, cur []string) []string {
	if pre == nil || cur == nil {
		return nil
	}
	preMap := covertSliceToMap(pre)
	curMap := covertSliceToMap(cur)
	return intersectionMapKeys(preMap, curMap)
}

/**
找到pre slice与cur slice不同的元素slice
 */
func Difference(pre []string, cur []string) []string {
	if pre == nil || cur == nil {
		return nil
	}
	preMap := covertSliceToMap(pre)
	curMap := covertSliceToMap(cur)
	return differenceMapKeys(preMap, curMap)
}

func differenceMapKeys(preProperties map[string]interface{}, curProperties map[string]interface{}) []string {
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

func intersectionMapKeys(preProperties map[string]interface{}, curProperties map[string]interface{}) []string {
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

func covertSliceToMap(pre []string) map[string]interface{} {
	if pre == nil {
		return nil
	}
	preMap := make(map[string]interface{}, len(pre))
	for index, value := range pre {
		preMap[value] = index
	}
	return preMap
}
