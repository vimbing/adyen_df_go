package adyendfgo

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"math"
	"strings"
)

func (f *Fingerprinter) calculateMd5(str string, num int) string {
	hash := md5.Sum([]byte(str))
	hashedString := base64.StdEncoding.EncodeToString(hash[:])
	return f.padString(hashedString, num)
}

func (f *Fingerprinter) rJust(str string, lenght int, char rune) string {
	var text string

	for i := 0; i < lenght-len(str); i++ {
		text += string(char)
	}

	return fmt.Sprintf("%s%s", text, str)
}

func (f *Fingerprinter) padString(str string, num int) string {
	paddedString := f.rJust(str, num, '0')

	if len(paddedString) > num {
		return paddedString[:num]
	}

	return paddedString
}

func (f *Fingerprinter) getSuperCookies() string {
	superCookiesPadding := math.Floor(f.SuperCookies / 2)
	deviceStorageValue := f.calculateMd5(f.DeviceStroage, int(superCookiesPadding))
	IEUDValue := f.calculateMd5(f.OXMLStorage, int(superCookiesPadding))

	superCookies := deviceStorageValue + IEUDValue

	return superCookies
}

func (f *Fingerprinter) getEntropy() string {
	// Idk what's that is supposed to do exactly, in this place useragent is already hashed, so it doesn't really check anything.

	// mobile := []string{"iPad", "iPhone", "iPod"}
	// if slices.Contains(mobile, f.userAgent) {
	// return "20"
	// }

	return "40"
}

func (f *Fingerprinter) GenerateFingerPrint() string {
	plugins := f.calculateMd5(f.PluginsString, f.Plugins)
	nrOfPlugins := f.padString(fmt.Sprint(f.PluginCount), f.NrOfPlugins)
	fonts := f.padString("", f.Fonts)
	nrOfFonts := f.padString("", f.NrOfFonts)
	timeZone := "CK1aUgqatB"
	video := f.padString(fmt.Sprint((f.ScreenWidth+7)*(f.ScreenHeight+7)*f.ScreenColorDepth), f.Video)
	superCookies := f.getSuperCookies()
	userAgent := f.calculateMd5(f.UserAgentString, f.UserAgent)
	mimeTypes := f.calculateMd5(f.MimeTypesString, f.MimeTypes)
	nrOfMimetypes := f.padString(fmt.Sprint(f.MimetTypesLength), f.NrOfMimeTypes)
	canvas := "rKkEK1Ha8P"
	cpuClass := f.padString("", f.CpuClass)
	platform := f.calculateMd5(f.PlatformString, f.Platform)
	doNotTrack := f.calculateMd5(f.DoNotTrackString, f.DoNotTrack)
	jsFonts := "iZCqnI4lsk"
	webglFp := "fKkhnraRhX"
	entropy := f.getEntropy()

	if f.Verbose {
		fmt.Printf("plugins: %v\n", plugins)
		fmt.Printf("nrOfPlugins: %v\n", nrOfPlugins)
		fmt.Printf("fonts: %v\n", fonts)
		fmt.Printf("nrOfFonts: %v\n", nrOfFonts)
		fmt.Printf("timeZone: %v\n", timeZone)
		fmt.Printf("video: %v\n", video)
		fmt.Printf("superCookies: %v\n", superCookies)
		fmt.Printf("userAgent: %v\n", userAgent)
		fmt.Printf("mimeTypes: %v\n", mimeTypes)
		fmt.Printf("nrOfMimetypes: %v\n", nrOfMimetypes)
		fmt.Printf("canvas: %v\n", canvas)
		fmt.Printf("cpuClass: %v\n", cpuClass)
		fmt.Printf("platform: %v\n", platform)
		fmt.Printf("doNotTrack: %v\n", doNotTrack)
		fmt.Printf("jsFonts: %v\n", jsFonts)
		fmt.Printf("webglFp: %v\n", webglFp)
		fmt.Printf("entropy: %v\n", entropy)

	}

	adyenFingerprint := fmt.Sprintf(
		"%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s:%s",
		plugins,
		nrOfPlugins,
		fonts,
		nrOfFonts,
		timeZone,
		video,
		superCookies,
		userAgent,
		mimeTypes,
		nrOfMimetypes,
		canvas,
		cpuClass,
		platform,
		doNotTrack,
		webglFp,
		jsFonts,
		entropy,
	)

	adyenFingerprint = strings.ReplaceAll(adyenFingerprint, "+", "G")
	adyenFingerprint = strings.ReplaceAll(adyenFingerprint, "/", "D")

	return adyenFingerprint
}
