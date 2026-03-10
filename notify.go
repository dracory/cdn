package cdn

func Notify_3_0_0() string {
	return cdnBase("https://cdn.jsdelivr.net/npm/") + "notifyjs@3.0.0/dist/notify.min.js"
}

func Notify_0_4_2() string {
	return cdnBase("https://cdnjs.cloudflare.com/ajax/libs/") + "notify/0.4.2/notify.min.js"
}
