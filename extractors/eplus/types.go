package eplus

type eplusData struct {
	AppId                 string `json:"app_id"`
	AppName               string `json:"app_name"`
	VenueName             string `json:"venue_name"`
	StreamType            string `json:"stream_type"`
	EmbeddedType          string `json:"embedded_type"`
	PlayerMode            string `json:"player_mode"`
	DeliveryStatus        string `json:"delivery_status"`
	StreamStatus          string `json:"stream_status"`
	NumberOfChannel       string `json:"number_of_channel"`
	EventDatetime         string `json:"event_datetime"`
	EventDatetimeText     string `json:"event_datetime_text"`
	BuyTicketUrl          string `json:"buy_ticket_url"`
	Description           string `json:"description"`
	PageDisplayPeriod     string `json:"page_display_period"`
	PageDisplayPeriodText string `json:"page_display_period_text"`
	IsLiveStream          bool   `json:"is_live_stream"`
	ChatMode              string `json:"chat_mode"`
	DrmMode               string `json:"drm_mode"`
	DrmStrictMode         string `json:"drm_strict_mode"`
	VirtualLiveMode       string `json:"virtual_live_mode"`
	ChatType              string `json:"chat_type"`
	ChatOpenMode          string `json:"chat_open_mode"`
	StreamEndAt           string `json:"stream_end_at"`
	ChatEmbedLink         string `json:"chat_embed_link"`
	ArchiveMode           string `json:"archive_mode"` // I didn't see values other than "ON"
	ArchivedStatus        string `json:"archived_status"`
	ImageUrl              string `json:"image_url"`
	EndEventImageUrl      string `json:"end_event_image_url"`
	AllInOneThumbnail     string `json:"all_in_one_thumbnail"`
	Content               string `json:"content"`
	AllowNumberOfViewable string `json:"allow_number_of_viewable"`
	EventType             string `json:"event_type"`
	TermEventFrom         string `json:"term_event_from"`
	TermEventTo           string `json:"term_event_to"`
	IsPassTicket          string `json:"is_pass_ticket"`
	DeleteFlag            string `json:"delete_flag"`
	CopyrightCode         string `json:"copyright_code"`
	ShootingAngle         string `json:"shooting_angle"`
	Extra                 string `json:"extra"`
	ExtraArchive          string `json:"extra_archive"`
	IsDvr                 string `json:"is_dvr"`
	ArchiveTime           string `json:"archive_time"`
	ArchiveKubun          string `json:"archive_kubun"`
	ZoomDateText          struct {
		Year          string `json:"year"`
		MonthDayFrom  string `json:"month_day_from"`
		MonthDayTo    string `json:"month_day_to"`
		DayOfWeekFrom string `json:"day_of_week_from"`
		DayOfWeekTo   string `json:"day_of_week_to"`
		Time          string `json:"time"`
	} `json:"zoom_date_text"`
	// The following 4 fields will disappear if the archive becomes unavailable.
	EmergencyMessage    string `json:"emergency_message"`
	EmergencyMode       string `json:"emergency_mode"`
	EmergencyImage      string `json:"emergency_image"`
	OriginChatEmbedLink string `json:"origin_chat_embed_link"`
}
