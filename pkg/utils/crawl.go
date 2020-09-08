package utils


func FindDuplicateDomain(url string, visitedDomains []string) bool {
	for _, elem := range visitedDomains {
		if elem == url {
			return true
		}
	}
	return false
}
