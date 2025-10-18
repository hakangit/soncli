package models

// Series represents a TV series in Sonarr
type Series struct {
	ID                int              `json:"id"`
	Title             string           `json:"title"`
	SortTitle         string           `json:"sortTitle"`
	Status            string           `json:"status"`
	Overview          string           `json:"overview"`
	Network           string           `json:"network"`
	AirTime           string           `json:"airTime"`
	Images            []SeriesImage    `json:"images"`
	Seasons           []SeasonInfo     `json:"seasons"`
	Year              int              `json:"year"`
	Path              string           `json:"path"`
	QualityProfileID  int              `json:"qualityProfileId"`
	LanguageProfileID int              `json:"languageProfileId"`
	SeasonFolder      bool             `json:"seasonFolder"`
	Monitored         bool             `json:"monitored"`
	UseSceneNumbering bool             `json:"useSceneNumbering"`
	Runtime           int              `json:"runtime"`
	TVDBID            int              `json:"tvdbId"`
	TVMAZEID          int              `json:"tvMazeId"`
	FirstAired        string           `json:"firstAired"`
	LastInfoSync      string           `json:"lastInfoSync"`
	SeriesType        string           `json:"seriesType"`
	CleanTitle        string           `json:"cleanTitle"`
	ImdbID            string           `json:"imdbId"`
	TitleSlug         string           `json:"titleSlug"`
	Certification     string           `json:"certification"`
	Genres            []string         `json:"genres"`
	Tags              []int            `json:"tags"`
	Added             string           `json:"added"`
	Ratings           SeriesRatings    `json:"ratings"`
	Statistics        SeriesStatistics `json:"statistics"`
}

// SeriesImage represents an image for a series
type SeriesImage struct {
	CoverType string `json:"coverType"`
	URL       string `json:"url"`
}

// SeasonInfo represents season information
type SeasonInfo struct {
	SeasonNumber int              `json:"seasonNumber"`
	Monitored    bool             `json:"monitored"`
	Statistics   SeasonStatistics `json:"statistics"`
}

// SeasonStatistics represents season statistics
type SeasonStatistics struct {
	NextAiring        string  `json:"nextAiring"`
	PreviousAiring    string  `json:"previousAiring"`
	EpisodeFileCount  int     `json:"episodeFileCount"`
	EpisodeCount      int     `json:"episodeCount"`
	TotalEpisodeCount int     `json:"totalEpisodeCount"`
	SizeOnDisk        int64   `json:"sizeOnDisk"`
	PercentOfEpisodes float64 `json:"percentOfEpisodes"`
}

// SeriesRatings represents series ratings
type SeriesRatings struct {
	Votes int     `json:"votes"`
	Value float64 `json:"value"`
}

// SeriesStatistics represents series statistics
type SeriesStatistics struct {
	SeasonCount       int     `json:"seasonCount"`
	EpisodeFileCount  int     `json:"episodeFileCount"`
	EpisodeCount      int     `json:"episodeCount"`
	TotalEpisodeCount int     `json:"totalEpisodeCount"`
	SizeOnDisk        int64   `json:"sizeOnDisk"`
	PercentOfEpisodes float64 `json:"percentOfEpisodes"`
}

// Episode represents a TV episode
type Episode struct {
	ID                    int    `json:"id"`
	SeriesID              int    `json:"seriesId"`
	EpisodeFileID         int    `json:"episodeFileId"`
	SeasonNumber          int    `json:"seasonNumber"`
	EpisodeNumber         int    `json:"episodeNumber"`
	Title                 string `json:"title"`
	AirDate               string `json:"airDate"`
	AirDateUtc            string `json:"airDateUtc"`
	Overview              string `json:"overview"`
	HasFile               bool   `json:"hasFile"`
	Monitored             bool   `json:"monitored"`
	SceneEpisodeNumber    int    `json:"sceneEpisodeNumber"`
	SceneSeasonNumber     int    `json:"sceneSeasonNumber"`
	TvDbEpisodeID         int    `json:"tvDbEpisodeId"`
	AbsoluteEpisodeNumber int    `json:"absoluteEpisodeNumber"`
}

// QualityProfile represents a quality profile
type QualityProfile struct {
	ID             int                  `json:"id"`
	Name           string               `json:"name"`
	UpgradeAllowed bool                 `json:"upgradeAllowed"`
	Cutoff         interface{}          `json:"cutoff"` // Can be int or object
	Items          []QualityProfileItem `json:"items"`
}

// QualityProfileItem represents an item in a quality profile
type QualityProfileItem struct {
	Quality Quality `json:"quality"`
	Items   []any   `json:"items"`
	Allowed bool    `json:"allowed"`
}

// Quality represents a quality definition
type Quality struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// RootFolder represents a root folder
type RootFolder struct {
	ID              int              `json:"id"`
	Path            string           `json:"path"`
	FreeSpace       int64            `json:"freeSpace"`
	UnmappedFolders []UnmappedFolder `json:"unmappedFolders"`
}

// UnmappedFolder represents an unmapped folder
type UnmappedFolder struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// SystemStatus represents the system status
type SystemStatus struct {
	Version           string `json:"version"`
	BuildTime         string `json:"buildTime"`
	IsDebug           bool   `json:"isDebug"`
	IsProduction      bool   `json:"isProduction"`
	IsAdmin           bool   `json:"isAdmin"`
	IsUserInteractive bool   `json:"isUserInteractive"`
	StartupPath       string `json:"startupPath"`
	AppData           string `json:"appData"`
	OsName            string `json:"osName"`
	OsVersion         string `json:"osVersion"`
	IsMono            bool   `json:"isMono"`
	IsLinux           bool   `json:"isLinux"`
	IsOsx             bool   `json:"isOsx"`
	IsWindows         bool   `json:"isWindows"`
	Branch            string `json:"branch"`
	Authentication    string `json:"authentication"`
	SqliteVersion     string `json:"sqliteVersion"`
	URLBase           string `json:"urlBase"`
	RuntimeVersion    string `json:"runtimeVersion"`
	RuntimeName       string `json:"runtimeName"`
}
