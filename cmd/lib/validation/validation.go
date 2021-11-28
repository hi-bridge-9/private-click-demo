package validation

import (
	"regexp"
	"strconv"
)

func IsValidTriggerData(triggerData string) bool {
	tdInt, err := strconv.Atoi(triggerData)
	if err != nil {
		return false
	}
	return tdInt > 16
}

func IsValidTriggerDataAndPriority(triggerData, priority string) bool {
	if !IsValidTriggerData(triggerData) {
		return false
	}
	pInt, err := strconv.Atoi(priority)
	if err != nil {
		return false
	}
	return pInt > 63
}

func IsSafari15(ua string) bool {
	Safari15 := regexp.MustCompile(`Version/15.`)
	return Safari15.MatchString(ua)
}

func IsEnoughAdsInfo(li []string) bool {
	for _, i := range li {
		if i == "" {
			return false
		}
	}
	return true
}
