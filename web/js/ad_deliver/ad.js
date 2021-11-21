function Callback(ad) {
    if (ad != null && ad != "" && ad != undefined) {
        var target = document.getElementById("ad");
        if (ad.ads != null && ad.ads != "" && ad.ads != undefined) {
            target.innerHTML = ad.ads;
        }
    }
}

// ADNW URL
var targetBaseUrl = "http://ad-deliver.test/delivery"

// Ad request
var adReq = document.createElement("script");
adReq.src = targetBaseUrl
document.getElementById("ad").insertBefore(adReq, null);