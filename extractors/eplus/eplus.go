package eplus

import (
	"encoding/json"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/iawia002/lux/config"
	"github.com/iawia002/lux/extractors"
	"github.com/iawia002/lux/request"
	"github.com/iawia002/lux/utils"
)

func init() {
	extractors.Register("eplus", New())
}

const (
	downloadclass = ".dloaddivcol"
)

type src struct {
	url     string
	quality string
	sizestr string
	size    int64
}

func getSrcMeta(text string) *src {
	sti := strings.Index(text, "(")
	ste := strings.Index(text, ")")
	itext := text[sti+1 : ste]
	strs := strings.Split(itext, ",")
	s := &src{}

	if len(strs) == 2 {
		s.quality = strings.Trim(strs[0], " ")
		s.sizestr = strings.Trim(strs[1], " ")
	}

	if s.sizestr == "" {
		s.size = 0
		return s
	}

	valunit := strings.Split(s.sizestr, " ")
	val, err := strconv.ParseFloat(valunit[0], 64)
	if err != nil {
		s.size = 0
		return s
	}
	unit := valunit[1]
	switch unit {
	case "KB":
		s.size = int64(val * 1024)
	case "MB":
		s.size = int64(val * 1024 * 1024)
	case "GB":
		s.size = int64(val * 1024 * 1024 * 1024)
	default:
		s.size = int64(val)
	}
	return s
}

func getEplsData(url string) {

}

type extractor struct{}

// New returns a eporner extractor.
func New() extractors.Extractor {
	return &extractor{}
}

// Extract is the main function to extract the data.
func (e *extractor) Extract(u string, option extractors.Options) ([]*extractors.Data, error) {
	html, err := request.Get(u, u, nil)

	{
		res, err := request.Request(http.MethodGet, u, nil, nil)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		defer res.Body.Close() // nolint

		cookiesArr := make([]string, 0)
		cookies := res.Cookies()

		for _, c := range cookies {
			cookiesArr = append(cookiesArr, c.Name+"="+c.Value)
		}

		os.Stdout.WriteString(strings.Join(cookiesArr, "; "))

		config.FakeHeaders["Cookie"] = strings.Join(cookiesArr, "; ")
	}

	re2 := regexp.MustCompile(`<script>\s*var app = (.+);\n`)

	match2 := re2.FindStringSubmatch(html)

	if match2 == nil {
		return nil, errors.New("the video data not found.")
	}

	os.Stdout.WriteString(match2[1])

	bgData := &eplusData{}

	json.Unmarshal([]byte(match2[1]), bgData)

	switch bgData.DeliveryStatus {
	case "PREPARING":
		return nil, errors.New("This event has not started yet.")
	case "STARTED":
		const errMsg = "Downloading live streaming is not supported. "
		if bgData.ArchiveMode == "ON" {
			return nil, errors.New(errMsg + "Downloading archive is supported and the archive will be available in " + bgData.ArchiveTime + " (Japan Standard Time).")
		} else {
			return nil, errors.New(errMsg)
		}
	case "STOPPED":
		const errMsg = "This event has been ended. "
		if bgData.ArchiveMode == "ON" {
			return nil, errors.New(errMsg + "The archive is not available, yet.")
		} else {
			return nil, errors.New(errMsg + "No archive for this event.")
		}
	case "WAIT_CONFIRM_ARCHIVED":
		// "bgData.ArchivedStatus" should be "SUCCEEDED"
		return nil, errors.New("The archive will be available shortly.")
	case "CONFIRMED_ARCHIVE":
		// The archive may be available.
	default:
		return nil, errors.New("Unknown delivery_status = " + bgData.DeliveryStatus + ".")
	}

	re := regexp.MustCompile(`var listChannels = \["(.+)"\]`)

	match := re.FindStringSubmatch(html)

	if match == nil {
		return nil, errors.New("can not find the playlist URL, the streaming channel may have ended.")
	}

	channel_url := strings.Replace(match[1], `\/`, "/", -1)

	urls, err := utils.M3u8URLs(channel_url)

	for _, u := range urls {
		os.Stdout.WriteString(u)
	}

	// How to pass the cookies to the downlaoder?

	// if err != nil {
	// 	return nil, errors.WithStack(err)
	// }

	// uu, err := url.Parse(u)
	// if err != nil {
	// 	return nil, errors.WithStack(err)
	// }
	// srcs := list.New() //getSrc(html)
	// streams := make(map[string]*extractors.Stream, len(srcs))
	// for _, src := range srcs {
	// 	srcurl := uu.Scheme + "://" + uu.Host + src.url
	// 	// skipping an extra HEAD request to the URL.
	// 	// size, err := request.Size(srcurl, u)
	// 	if err != nil {
	// 		return nil, errors.WithStack(err)
	// 	}
	// 	urlData := &extractors.Part{
	// 		URL:  srcurl,
	// 		Size: src.size,
	// 		Ext:  "mp4",
	// 	}
	// 	streams[src.quality] = &extractors.Stream{
	// 		Parts:   []*extractors.Part{urlData},
	// 		Size:    src.size,
	// 		Quality: src.quality,
	// 	}
	// }
	// return []*extractors.Data{
	// 	{
	// 		Site:    "EPORNER eporner.com",
	// 		Title:   title,
	// 		Type:    extractors.DataTypeVideo,
	// 		Streams: streams,
	// 		URL:     u,
	// 	},
	// }, nil

	return nil, err
}
