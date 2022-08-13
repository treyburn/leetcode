package destination

func destCity(paths [][]string) string {
	destinationsWithNoArrival := make(map[string]struct{})
	arrivals := make(map[string]struct{})
	for _, pair := range paths {
		arrivals[pair[0]] = struct{}{}
		if _, ok := destinationsWithNoArrival[pair[0]]; ok {
			delete(destinationsWithNoArrival, pair[0])
		}
		if _, ok := arrivals[pair[1]]; !ok {
			destinationsWithNoArrival[pair[1]] = struct{}{}
		}
	}

	for k := range destinationsWithNoArrival {
		return k
	}

	return ""
}
