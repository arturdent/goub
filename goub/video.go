package goub

var (
	VideoTypes      = [...]string{"flw", "mp4"}
	VideoDementions = [...]string{"big", "med", "small"}
)

type CoubService struct {
	client *Client
}

type Web struct {
	Template *string
	// Types
	// Version
}

type WebChunks struct {
	Template *string
	// Types
	// Version
}

type HTML5 struct {
	Template *string
	Chunks   *WebChunks
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

type Coub struct {
	ID             *int
	TYPE           *string
	Permalink      *string
	Title          *string
	VisabilityType *string
	ChannelID      *int
	CreatedAt      *Timestamp
	UpdatedAt      *Timestamp
	IsDone         *bool
	Duration       *float64
	ViewsCount     *int
	COTD           *bool
	COTDAt         *Timestamp
	Recoub         *bool
	Like           *bool
	RecoubsCount   *int
	LikesCount     *int
	RecoubTo       *int
	Flag           *bool
	OriginalSound  *bool
	HasSound       *bool
	// FileVersions
}

func (c *CoubService) Get() {}
