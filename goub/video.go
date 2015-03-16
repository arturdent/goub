package goub

type CoubService struct {
	client *Client
}

type Web struct {
	Template *string
	Types    []string
	Versions []string
}

type WebChunks struct {
	Template *string
	Types    []string
	Versions []string
	Chuncks  []int
}

type HTML5 struct {
	Template *string
	Chunks   []int
}

type IPhone struct {
	URL *string
}

type Mobile struct {
	AudioURL    *string
	Base64URl   *string
	Base64Files *string
	FramesCount *int
}

type AudioVersionsChuncks struct {
	Template *string
	Versions []string
	Chuncks  []int
}

type AudioVersions struct {
	Template *string
	Version  []string
	Chunks   *AudioVersionsChuncks
}

type FLVAudioVersionsFile struct {
	Template *string
	Version  []string
}

type FLVAudioVersionsChuncks struct {
	Template *string
	Version  []string
	Chuncks  []int
}

type FLVAudioVersions struct {
	File    *FLVAudioVersionsFile
	Chuncks *FLVAudioVersionsChuncks
}

type ImageVersions struct {
	Template *string
	Version  []string
}

type FirstFrameVersions struct {
	Template *string
	Version  []string
}

type Dimenstions struct {
	Big   [][2]int
	Med   [][2]int
	Small [][2]int
}

type ExternalDownload struct {
	Type *string
	URL  *string
}

type AvatarVersions struct {
	Type     *string
	Versions []string
}

type Channel struct {
	ID             *int
	Permalink      *string
	Description    *string
	Title          *string
	IFollowHim     *bool
	FollowersCount *int
	FollowingCount *int
}

type TrackMeta struct {
	Year   *string
	Album  *string
	Title  *string
	Artist *string
}

type AudioTrack struct {
	ID            *int
	Title         *string
	URL           *string
	Image         *string
	ImageRetina   *string
	ItunesURL     *string
	AmazonURL     *string
	BadcampURL    *string
	SoundCloudURL *string
	TrackName     *string
	TrackArtist   *string
	TrackAlbum    *string
	Meta          *TrackMeta
}

type ExternalVideo struct {
	ID            *int
	Title         *string
	URL           *string
	Image         *string
	ImageRetina   *string
	ItunesURL     *string
	AmazonURL     *string
	BadcampURL    *string
	SoundCloudURL *string
}

type MediaBlocks struct {
	CoubsAudioTrack    *AudioTrack
	CoubsExternalVideo *ExternalVideo
}

type FileVersions struct {
	Web        *Web
	WebChuncks *WebChunks
	HTML5      *HTML5
	IPhone     *IPhone
	Mobile     *Mobile
}

type Tag struct {
	ID    *int
	Type  *string
	Value *string
}

type Coub struct {
	ID                  *int
	TYPE                *string
	Permalink           *string
	Title               *string
	VisabilityType      *string
	ChannelID           *int
	CreatedAt           *Timestamp
	UpdatedAt           *Timestamp
	IsDone              *bool
	Duration            *float64
	ViewsCount          *int
	COTD                *bool
	COTDAt              *Timestamp
	Recoub              *bool
	Like                *bool
	RecoubsCount        *int
	LikesCount          *int
	RecoubTo            *int
	Flag                *bool
	OriginalSound       *bool
	HasSound            *bool
	FileVersions        *FileVersions
	AudioVersions       *AudioVersions
	FLVAudioVersions    *FLVAudioVersions
	ImageVersions       *ImageVersions
	FirstFrameVersions  *FirstFrameVersions
	Dimenstions         *Dimenstions
	AgeResticted        *bool
	AgeRestictedByAdmin *bool
	AllowReuse          *bool
	Banned              *bool
	ExternalDownload    *ExternalDownload
	Channel             *Channel
	PersentDone         *int
	Tags                []Tag
	RawVideoID          *int
	MediaBlocks         *MediaBlocks
	RawVideoThumbNail   *string
	RawVideoTitle       *string
}

func (c *CoubService) Get() {}
