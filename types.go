package tg

import ()

// https://core.telegram.org/bots/api#update
type Update struct {
	UpdateId           int                `json:"update_id"`
	Message            Message            `json:"message"`
	EditedMessage      Message            `json:"edited_message"`
	ChannelPost        Message            `json:"channel_post"`
	EditedChannelPost  Message            `json:"edited_channel_post"`
	InlineQuery        InlineQuery        `json:"inline_query"`
	ChosenInlineResult ChosenInlineResult `json:"chosen_inline_result"`
	CallbackQuery      CallbackQuery      `json:"callback_query"`
	ShippingQuery      ShippingQuery      `json:"shipping_query"`
	PreCheckoutQuery   PreCheckoutQuery   `json:"pre_checkout_query"`
}

// https://core.telegram.org/bots/api#user
type User struct {
	Id           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

// https://core.telegram.org/bots/api#chat
type Chat struct {
	Id                          int    `json:"id"`
	Type                        string `json:"type"`
	Title                       string `json:"title"`
	Username                    string `json:"username"`
	FirstName                   string `json:"first_name"`
	LasteName                   string `json:"last_name"`
	AllMembersAreAdministrators bool   `json:"all_members_are_administrators"`
}

// https://core.telegram.org/bots/api#message
type Message struct {
	MessageId             int               `json:"message_id"`
	From                  User              `json:"from"`
	Data                  int               `json:"data"`
	Chat                  Chat              `json:"chat"`
	ForwardFrom           User              `json:"forward_from"`
	ForwardFromChat       Chat              `json:"forward_from_chat"`
	ForwardFromMessageId  int               `json:"forward_from_message_id"`
	ForwardDate           int               `json:"forward_date"`
	ReplyToMessage        *Message          `json:"reply_to_message"`
	EditDate              int               `json:"edit_date"`
	Text                  string            `json:"text"`
	Entities              []MessageEntity   `json:"entities"`
	Audio                 Audio             `json:"audio"`
	Document              Document          `json:"document"`
	Game                  Game              `json:"game"`
	Photo                 []PhotoSize       `json:"photo"`
	Sticker               Sticker           `json:"sticker"`
	Video                 Video             `json:"video"`
	Voice                 Voice             `json:"voice"`
	VideoNote             VideoNote         `json:"video_note"`
	NewChatMembers        []User            `json:"new_chat_members"`
	Caption               string            `json:"caption"`
	Contact               Contact           `json:"contact"`
	Location              Location          `json:"location"`
	Venu                  Venu              `json:"venu"`
	NewChatMember         User              `json:"new_chat_member"`
	LeftChatMember        User              `json:"left_chat_member"`
	NewChatTitle          string            `json:"new_chat_title"`
	NewChatPhoto          []PhotoSize       `json:"new_chat_photo"`
	DeleteChatPhoto       bool              `json:"delete_chat_photo"`
	GroupChatCreated      bool              `json:"group_chat_creates"`
	SupergroupChatCreated bool              `json:"supergroup_chat_created"`
	ChannelChatCreated    bool              `json:"channel_chat_created"`
	MigrateToChatId       int               `json:"migrate_to_chat_id"`
	MigrateFromChatId     int               `json:"migrate_from_chat_id"`
	PinnedMessage         *Message          `json:"pinned_message"`
	Invoice               Invoice           `json:"invoice"`
	SuccessfulPayment     SuccessfulPayment `json:"successful_payment"`
}

// https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	Url    string `json:"url"`
	User   User   `json:"user"`
}

// https://core.telegram.org/bots/api#photosize
type PhotoSize struct {
	FileId   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size"`
}

// https://core.telegram.org/bots/api#audio
type Audio struct {
	FileId    string `json:"file_id"`
	Duration  int    `json:"duration"`
	Performer string `json:"performer"`
	Title     string `json:"title"`
	MimeType  string `json:"mime_type"`
	FileSize  int    `json:"file_size"`
}

// https://core.telegram.org/bots/api#document
type Document struct {
	FileId   string    `json:"file_id"`
	Thumb    PhotoSize `json:"thumb"`
	FileName string    `json:"file_name"`
	MimeType string    `json:"mime_type"`
	FileSize int       `json:"file_size"`
}

// https://core.telegram.org/bots/api#sticker
type Sticker struct {
	FileId   string    `json:"file_id"`
	Width    int       `json:"width"`
	Height   int       `json:"height"`
	Thumb    PhotoSize `json:"thumb"`
	Emoji    string    `json:"emoji"`
	FileSize int       `json:"file_size"`
}

// https://core.telegram.org/bots/api#video
type Video struct {
	FileId   string    `json:"file_id"`
	Width    int       `json:"width"`
	Height   int       `json:"height"`
	Duration int       `json:"duration"`
	Thumb    PhotoSize `json:"thumb"`
	MimeType string    `json:"mime_type"`
	FileSize int       `json:"file_size"`
}

// https://core.telegram.org/bots/api#voice
type Voice struct {
	FileId   string    `json:"file_id"`
	Duration int       `json:"duration"`
	Thumb    PhotoSize `json:"thumb"`
	MimeType string    `json:"mime_type"`
	FileSize int       `json:"file_size"`
}

// https://core.telegram.org/bots/api#videonote
type VideoNote struct {
	FileId   string    `json:"file_id"`
	Length   int       `json:"length"`
	Duration int       `json:"duration"`
	Thumb    PhotoSize `json:"thumb"`
	FileSize int       `json:"file_size"`
}

// https://core.telegram.org/bots/api#contact
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserId      int    `json:"user_id"`
}

// https://core.telegram.org/bots/api#location
type Location struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

// https://core.telegram.org/bots/api#venue
type Venu struct {
	Location     Location `json:"location"`
	Title        string   `json:"title"`
	Address      string   `json:"address"`
	FoursquareId string   `json:"foursquare_id"`
}

// https://core.telegram.org/bots/api#userprofilephotos
type UserProfilePhotos struct {
	TotalCount int         `json:"total_count"`
	Photos     []PhotoSize `json:"photos"`
}

// https://core.telegram.org/bots/api#file
type File struct {
	FileId   string `json:"file_id"`
	FileSize int    `json:"file_size"`
	FilePath string `json:"file_path"`
}

// https://core.telegram.org/bots/api#replykeyboardmarkup
type ReplyKeyboardMarkup struct {
	Keyboard        [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool               `json:"resize_keyboard"`
	OneTimeKeyboard bool               `json:"one_time_keyboard"`
	Selective       bool               `json:"selective"`
}

// https://core.telegram.org/bots/api#keyboardbutton
type KeyboardButton struct {
	Text            string `json:"text"`
	RequestContact  string `json:"request_contact"`
	RequestLocation string `json:"request_location"`
}

// https://core.telegram.org/bots/api#replykeyboardremove
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"` // Always True
	Selective      bool `json:"selective"`
}

// https://core.telegram.org/bots/api#inlinekeyboardmarkup
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// https://core.telegram.org/bots/api#inlinekeyboardbutton
type InlineKeyboardButton struct {
	Text                         string `json:"text"`
	Url                          string `json:"url"`
	CallbackData                 string `json:"callback_data"`
	SwitchInlineQuery            string `json:"switch_inline_query"`
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat"`
	CallbackGame                 string `json:"callback_game"`
	Pay                          bool   `json:"pay"`
}

// https://core.telegram.org/bots/api#callbackquery
type CallbackQuery struct {
	Id              string  `json:"id"`
	From            User    `json:"from"`
	Message         Message `json:"message"`
	InlineMessageId string  `json:"inline_message_id"`
	ChatInstance    string  `json:"chat_instance"`
	Data            string  `json:"data"`
	GameShortName   string  `json:"get_short_name"`
}

// https://core.telegram.org/bots/api#forcereply
type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective"`
}

// https://core.telegram.org/bots/api#chatmember
type ChatMember struct {
	User   User   `json:"user"`
	Status string `json:"status"`
}

// https://core.telegram.org/bots/api#responseparameters
type ResponseParameters struct {
	MigrateToChatId int `json:"migrate_to_chat_id"`
	RetryAfter      int `json:"retry_after"`
}

//

// https://core.telegram.org/bots/api#labeledprice
type LabeledPrice struct {
	Label  string `json:"label"`
	Amount int    `json:"amount"`
}

// https://core.telegram.org/bots/api#invoice
type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
}

// https://core.telegram.org/bots/api#shippingaddress
type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

// https://core.telegram.org/bots/api#orderinfo
type OrderInfo struct {
	Name            string          `json:"name"`
	PhoneNumber     string          `json:"phone_number"`
	Email           string          `json:"email"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

// https://core.telegram.org/bots/api#shippingoption
type ShippingOption struct {
	Id    string         `json:"id"`
	Title string         `json:"title"`
	Price []LabeledPrice `json:"price"`
}

// https://core.telegram.org/bots/api#successfulpayment
type SuccessfulPayment struct {
	Currency                string    `json:"currency"`
	TotalAmount             int       `json:"total_amount"`
	InvoicePayload          string    `json:"invoice_payload"`
	ShippingOptionId        string    `json:"shipping_option_id"`
	OrderInfo               OrderInfo `json:"order_info"`
	TelegramPaymentChargeId string    `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeId string    `json:"provider_payment_charge_id"`
}

// https://core.telegram.org/bots/api#shippingquery
type ShippingQuery struct {
	Id              string          `json:"id"`
	From            User            `json:"from"`
	InvoicePayload  string          `json:"invoice_payload"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

// https://core.telegram.org/bots/api#precheckoutquery
type PreCheckoutQuery struct {
	Id               string    `json:"id"`
	From             User      `json:"from"`
	Currency         string    `json:"currency"`
	TotalAmount      int       `json:"total_amount"`
	InvoicePayload   string    `json:"invoice_payload"`
	ShippingOptionId string    `json:"shipping_option_id"`
	OrderInfo        OrderInfo `json:"order_info"`
}

// https://core.telegram.org/bots/api#game
type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         string          `json:"text"`
	TextEntities []MessageEntity `json:"text_entities"`
	Animation    Animation       `json:"animation"`
}

// https://core.telegram.org/bots/api#animation
type Animation struct {
	FileId   string    `json:"file_id"`
	Thumb    PhotoSize `json:"thumb"`
	FileName string    `json:"file_name"`
	MimeType string    `json:"mime_type"`
	FileSize int       `json:"file_size"`
}

// https://core.telegram.org/bots/api#gamehighscore
type GameHighScore struct {
	Position int  `json:"position"`
	User     User `json:"user"`
	Score    int  `json:"score"`
}

// https://core.telegram.org/bots/api#inlinequery
type InlineQuery struct {
	Id       string   `json:"id"`
	From     User     `json:"from"`
	Location Location `json:"location"`
	Query    string   `json:"query"`
	Offset   string   `json:"offset"`
}

// https://core.telegram.org/bots/api#inlinequeryresultarticle
type InlineQueryResultArticle struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	Title               string               `json:"title"`
	//TODO InputMessageContent InputMessageContent  `json:"input_message_content"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	Url                 string               `json:"url"`
	HideUrl             bool                 `json:"hide_url"`
	Description         string               `json:"description"`
	ThumbUrl            string               `json:"thumb_url"`
	ThumbWidth          int                  `json:"thumb_width"`
	ThumbHeight         int                  `json:"thumb_height"`
}

// https://core.telegram.org/bots/api#inlinequeryresultphoto
type InlineQueryResultPhoto struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	PhotoUrl            string               `json:"photo_url"`
	ThumbUrl            string               `json:"thumb_url"`
	PhotoWidth          int                  `json:"photo_width"`
	PhotoHeight         int                  `json:"photo_height"`
	Title               string               `json:"title"`
	Description         string               `json:"description"`
	Caption             string               `json:"caption"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	//TODO InputMessageContent InputMessageContent  `json:"input_message_content"`
}

// https://core.telegram.org/bots/api#inlinequeryresultgif
type InlineQueryResultGif struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	GifUrl              string               `json:"gif_url"`
	GifWidth            int                  `json:"gif_width"`
	GifHeight           int                  `json:"gif_height"`
	GifDuration         int                  `json:"gif_duration"`
	ThumbUrl            string               `json:"thumb_url"`
	Title               string               `json:"title"`
	Caption             string               `json:"caption"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	//TODO InputMessageContent InputMessageContent  `json:"input_message_content"`
}

// https://core.telegram.org/bots/api#inlinequeryresultmpeg4gif
type InlineQueryResultMpeg4Gif struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	Mpeg4Url            string               `json:"mpeg4_url"`
	Mpeg4Width          int                  `json:"mpeg4_width"`
	Mpeg4Height         int                  `json:"mpeg4_height"`
	Mpeg4Duration       int                  `json:"mpeg4_duration"`
	ThumbUrl            string               `json:"thumb_url"`
	Title               string               `json:"title"`
	Caption             string               `json:"caption"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	//TODO InputMessageContent InputMessageContent  `json:"input_message_content"`
}

// https://core.telegram.org/bots/api#inlinequeryresultvideo
type InlineQueryResultVideo struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	VideoUrl            string               `json:"video_url"`
	MimeType            string               `json:"mime_type"`
	ThumbUrl            string               `json:"thumb_url"`
	Title               string               `json:"title"`
	Caption             string               `json:"caption"`
	VideoWidth          int                  `json:"video_width"`
	VideoHeight         int                  `json:"video_height"`
	VideoDuration       int                  `json:"video_duration"`
	Description         string               `json:"description"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	//TODO InputMessageContent InputMessageContent  `json:"input_message_content"`
}

// https://core.telegram.org/bots/api#inlinequeryresultaudio
type InlineQueryResultAudio struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	AudioUrl            string               `json:"audio_url"`
	Title               string               `json:"title"`
	Caption             string               `json:"caption"`
	Performer           string               `json:"performer"`
	AudioDuration       int                  `json:"audio_duration"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	//TODO InputMessageContent InputMessageContent  `json:"input_message_content"`
}

// https://core.telegram.org/bots/api#inlinequeryresultvoice
type InlineQueryResultVoice struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	VoiceUrl            string               `json:"voice_url"`
	Title               string               `json:"title"`
	Caption             string               `json:"caption"`
	VoiceDuration       int                  `json:"voice_duration"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	//TODO InputMessageContent InputMessageContent  `json:"input_message_content"`
}

// https://core.telegram.org/bots/api#inlinequeryresultdocument
type InlineQueryResultDocument struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	Title               string               `json:"title"`
	Caption             string               `json:"caption"`
	DocumentUrl         string               `json:"document_url"`
	MimeType            string               `json:"mime_type"`
	Description         string               `json:"description"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	//TODO InputMessageContent InputMessageContent  `json:"input_message_content"`
	ThumbUrl            string               `json:"thumb_url"`
	ThumbWidth          int                  `json:"thumb_width"`
	ThumbHeight         int                  `json:"thumb_height"`
}

// https://core.telegram.org/bots/api#inlinequeryresultlocation
type InlineQueryResultLocation struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	Latitude            float32              `json:"latitude"`
	Longitude           float32              `json:"longitude"`
	Title               string               `json:"title"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	//TODO InputMessageContent InputMessageContent  `json:"input_message_content"`
	ThumbUrl            string               `json:"thumb_url"`
	ThumbWidth          int                  `json:"thumb_width"`
	ThumbHeight         int                  `json:"thumb_height"`
}

// https://core.telegram.org/bots/api#inlinequeryresultvenue
type InlineQueryResultVenue struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	Latitude            float32              `json:"latitude"`
	Longitude           float32              `json:"longitude"`
	Title               string               `json:"title"`
	Address             string               `json:"address"`
	FoursquareId        string               `json:"foursquare_id"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	//TODO InputMessageContent InputMessageContent  `json:"input_message_content"`
	ThumbUrl            string               `json:"thumb_url"`
	ThumbWidth          int                  `json:"thumb_width"`
	ThumbHeight         int                  `json:"thumb_height"`
}

// https://core.telegram.org/bots/api#inlinequeryresultcontact
type InlineQueryResultContact struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	PhoneNumber         string               `json:"phone_number"`
	FirstName           string               `json:"first_name"`
	LastName            string               `json:"last_name"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	//TODO InputMessageContent InputMessageContent  `json:"input_message_content"`
	ThumbUrl            string               `json:"thumb_url"`
	ThumbWidth          int                  `json:"thumb_width"`
	ThumbHeight         int                  `json:"thumb_height"`
}

// https://core.telegram.org/bots/api#inlinequeryresultgame
type InlineQueryResultGame struct {
	Type          string               `json:"type"`
	Id            string               `json:"id"`
	GameShortName string               `json:"game_short_name"`
	ReplyMarkup   InlineKeyboardMarkup `json:"reply_markup"`
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedphoto
type InlineQueryResultCachedPhoto struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	PhotoFileId         string               `json:"photo_file_id"`
	Title               string               `json:"title"`
	Description         string               `json:"description"`
	Caption             string               `json:"caption"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	//TODO InputMessageContent InputMessageContent  `json:"input_message_content"`
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedgif
type InlineQueryResultCachedGif struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	GifFileId           string               `json:"gif_file_id"`
	Title               string               `json:"title"`
	Caption             string               `json:"caption"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	//TODO InputMessageContent InputMessageContent  `json:"input_message_content"`
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedmpeg4gif
type InlineQueryResultCachedMpeg4Gif struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	Mpeg4FileId         string               `json:"mpeg4_file_id"`
	Title               string               `json:"title"`
	Caption             string               `json:"caption"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	//TODO InputMessageContent InputMessageContent  `json:"input_message_content"`
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedsticker
type InlineQueryResultCachedSticker struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	StickerFileId       string               `json:"sticker_file_id"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	//TODO InputMessageContent InputMessageContent  `json:"input_message_content"`
}

// https://core.telegram.org/bots/api#inlinequeryresultcacheddocument
type InlineQueryResultCachedDocument struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	Title               string               `json:"title"`
	DocumentFileId      string               `json:"document_file_id"`
	Description         string               `json:"description"`
	Caption             string               `json:"caption"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	//TODO InputMessageContent InputMessageContent  `json:"input_message_content"`
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedvideo
type InlineQueryResultCachedVideo struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	VideoFileId         string               `json:"video_file_id"`
	Title               string               `json:"title"`
	Description         string               `json:"description"`
	Caption             string               `json:"caption"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	//TODO InputMessageContent InputMessageContent  `json:"input_message_content"`
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedvoice
type InlineQueryResultCachedVoice struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	VoiceFileId         string               `json:"voice_file_id"`
	Title               string               `json:"title"`
	Caption             string               `json:"caption"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	//TODO InputMessageContent InputMessageContent  `json:"input_message_content"`
}

// https://core.telegram.org/bots/api#inlinequeryresultcachedaudio
type InlineQueryResultCachedAudio struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	AudioFileId         string               `json:"audio_file_id"`
	Caption             string               `json:"caption"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	//TODO InputMessageContent InputMessageContent  `json:"input_message_content"`
}

// https://core.telegram.org/bots/api#inputtextmessagecontent
type InputTextMessageContent struct {
	MessageText           string `json:"message_text"`
	ParseMode             string `json:"parse_mode"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview"`
}

// https://core.telegram.org/bots/api#inputlocationmessagecontent
type InputLocationMessageContent struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

// https://core.telegram.org/bots/api#inputvenuemessagecontent
type InputVenueMessageContent struct {
	Latitude     float32 `json:"latitude"`
	Longitude    float32 `json:"longitude"`
	Title        string  `json:"title"`
	Address      string  `json:"address"`
	FoursquareId string  `json:"foursquare_id"`
}

// https://core.telegram.org/bots/api#inputcontactmessagecontent
type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}

// https://core.telegram.org/bots/api#choseninlineresult
type ChosenInlineResult struct {
	ResultId        string   `json:"result_id"`
	From            User     `json:"from"`
	Location        Location `json:"location"`
	InlineMessageId string   `json:"inline_message_id"`
	Query           string   `json:"query"`
}

