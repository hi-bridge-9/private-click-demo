package validation

import (
	"strconv"
)

func IsValidTriggerData(triggerData string) bool {
	tdInt, err := strconv.Atoi(triggerData)
	if err != nil || tdInt > 16 {
		return false
	}
	return true
}

func IsValidTriggerDataAndPriority(triggerData, priority string) bool {
	if !IsValidTriggerData(triggerData) {
		return false
	}
	pInt, err := strconv.Atoi(priority)
	if err != nil || pInt > 63 {
		return false
	}
	return true
}
