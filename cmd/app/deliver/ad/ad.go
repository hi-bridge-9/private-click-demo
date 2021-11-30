package ad

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/kyu-takahahsi/private-click-demo/cmd/lib/validation"
)

var (
	imageLink   = os.Getenv("")
	width       = os.Getenv("450")
	height      = os.Getenv("450")
	clickServer = os.Getenv("CLICK_SERVER_DOMAIN")
	lp          = os.Getenv("ADVERTISER_LP_PAGE")
	href        = fmt.Sprintf("%s/click?lp=%s", clickServer, lp)
	cvLocation  = os.Getenv("CV_LOCATION_DOMAIN")
	sourceId    = RandDigit(256) // 0〜255
	nonce       = RandStr(16)    // 16byte
	info        = []string{imageLink, width, height, href, lp, cvLocation, string(sourceId), nonce}
)

func Generate(ua string) (string, error) {
	if !validation.IsEnoughAdsInfo(info) {
		return "", errors.New("Not enough ad information")
	}

	imgTag := fmt.Sprintf("<img src=\"%s\" width=\"%s\" height=\"%s\">", imageLink, width, height)
	if validation.IsSafari15(ua) {
		return fmt.Sprintf("<a href=\"%s\" attributiondestination=\"%s\""+
			"attributionsourceid=%d attributionsourcenonce=\"%s\">%s</a>", href, cvLocation, sourceId, nonce, imgTag), nil
	}
	return fmt.Sprintf("<a href=\"%s\" attributeon=\"%s\" attributionsourceid=\"%d\">%s</a>",
		href, cvLocation, sourceId, imgTag), nil
}

func RandDigit(n int) int {
	return rand.Intn(n)
}

func RandStr(n int) string {
	rand.Seed(time.Now().UnixNano())
	var Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = Letters[rand.Intn(len(Letters))]
	}
	return string(b)
}
