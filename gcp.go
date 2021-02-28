package printer

import (
	"net/http"
)

var assets Assets

// GCP Cloud Functions requires that the function being deployed be on a file on the root directory
// The directory in which the code resides inside the Cloud Function is `serverless_function_source_code`
// That's the reason why in this case we need to append that folder as prefix to our assets
func init() {
	assets = Assets{
		BgImgPath: "serverless_function_source_code/assets/00-instagram-background.png",
		FontPath:  "serverless_function_source_code/assets/FiraSans-Light.ttf",
		FontSize:  60,
	}
}

func Printer(w http.ResponseWriter, r *http.Request) {
	assets.Serve(w, r)
}
