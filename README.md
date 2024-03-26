# Adyen Device Fingerprint Encryptor in GO
# This is just rewrite of Volt#9540 repo https://github.com/benastahl/adyen_df_python

```go
package main

import (
	"fmt"

	adyen "github.com/vimbing/adyen_df_go"
)

func main() {
	f := &adyen.Fingerprinter{
		Plugins:          10,
		NrOfPlugins:      3,
		Fonts:            10,
		NrOfFonts:        3,
		TimeZone:         10,
		Video:            10,
		SuperCookies:     10,
		UserAgent:        10,
		MimeTypes:        10,
		NrOfMimeTypes:    3,
		Canvas:           10,
		CpuClass:         5,
		Platform:         5,
		DoNotTrack:       5,
		WebGlFp:          10,
		JsFonts:          10,
		UserAgentString:  "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36 Edge/12.0",
		PluginsString:    "Plugin 0: Chrome PDF Viewer; Portable Document Format; internal-pdf-viewer; (Portable Document Format; application/pdf; pdf) (Portable Document Format; text/pdf; pdf). Plugin 1: Chromium PDF Viewer; Portable Document Format; internal-pdf-viewer; (Portable Document Format; application/pdf; pdf) (Portable Document Format; text/pdf; pdf). Plugin 2: Microsoft Edge PDF Viewer; Portable Document Format; internal-pdf-viewer; (Portable Document Format; application/pdf; pdf) (Portable Document Format; text/pdf; pdf). Plugin 3: PDF Viewer; Portable Document Format; internal-pdf-viewer; (Portable Document Format; application/pdf; pdf) (Portable Document Format; text/pdf; pdf). Plugin 4: WebKit built-in PDF; Portable Document Format; internal-pdf-viewer; (Portable Document Format; application/pdf; pdf) (Portable Document Format; text/pdf; pdf). ",
		PluginCount:      5,
		ScreenWidth:      1440,
		ScreenHeight:     900,
		ScreenColorDepth: 30,
		DeviceStroage:    "DOM-LS: Yes, DOM-SS: Yes",
		OXMLStorage:      ", IE-UD: No",
		MimeTypesString:  "Portable Document Formatapplication/pdfpdfPortable Document Formattext/pdfpdf",
		MimetTypesLength: 2,
		PlatformString:   "MacIntel",
		DoNotTrackString: "1",
		Entropy:          "40",
	}

	fmt.Printf("fingerprint: %s\n", f.GenerateFingerPrint())
}
```

### Output:
```sh
fingerprint: DpqwU4zEdN0050000000000000CK1aUgqatB0039372870cVB94iKzBGRaIEwiHpQzBix7RX3az8002rKkEK1Ha8P00000YVxErxMpCOfKkhnraRhXiZCqnI4lsk:40```
