package qbittorrent_api

import (
	"encoding/json"
	"github.com/xiangyt/qbittorrent-api/definition"
	"io"
)

type TorrentsApi struct {
	client *Client
}

type Torrent struct {
	AddedOn           int           `json:"added_on"`
	AmountLeft        int64         `json:"amount_left"`
	AutoTmm           bool          `json:"auto_tmm"`
	Availability      float64       `json:"availability"`
	Category          string        `json:"category"`
	Completed         int64         `json:"completed"`
	CompletionOn      int           `json:"completion_on"`
	ContentPath       string        `json:"content_path"`
	DlLimit           int           `json:"dl_limit"`
	Dlspeed           int           `json:"dlspeed"`
	DownloadPath      string        `json:"download_path"`
	Downloaded        int64         `json:"downloaded"`
	DownloadedSession int           `json:"downloaded_session"`
	Eta               int           `json:"eta"`
	FLPiecePrio       bool          `json:"f_l_piece_prio"`
	ForceStart        bool          `json:"force_start"`
	Hash              string        `json:"hash"`
	InfohashV1        string        `json:"infohash_v1"`
	InfohashV2        string        `json:"infohash_v2"`
	LastActivity      int           `json:"last_activity"`
	MagnetUri         string        `json:"magnet_uri"`
	MaxRatio          int           `json:"max_ratio"`
	MaxSeedingTime    int           `json:"max_seeding_time"`
	Name              string        `json:"name"`
	NumComplete       int           `json:"num_complete"`
	NumIncomplete     int           `json:"num_incomplete"`
	NumLeechs         int           `json:"num_leechs"`
	NumSeeds          int           `json:"num_seeds"`
	Priority          int           `json:"priority"`
	Progress          float64       `json:"progress"`
	Ratio             float64       `json:"ratio"`
	RatioLimit        int           `json:"ratio_limit"`
	SavePath          string        `json:"save_path"`
	SeedingTime       int           `json:"seeding_time"`
	SeedingTimeLimit  int           `json:"seeding_time_limit"`
	SeenComplete      int           `json:"seen_complete"`
	SeqDl             bool          `json:"seq_dl"`
	Size              int64         `json:"size"`
	State             TorrentStates `json:"state"`
	SuperSeeding      bool          `json:"super_seeding"`
	Tags              string        `json:"tags"`
	TimeActive        int           `json:"time_active"`
	TotalSize         int64         `json:"total_size"`
	Tracker           string        `json:"tracker"`
	TrackersCount     int           `json:"trackers_count"`
	UpLimit           int           `json:"up_limit"`
	Uploaded          int64         `json:"uploaded"`
	UploadedSession   int           `json:"uploaded_session"`
	Upspeed           int           `json:"upspeed"`
}

type TorrentInfoList []*Torrent
type TorrentStates string

const (
	ERROR                    TorrentStates = "error"
	MISSING_FILES                          = "missingFiles"
	UPLOADING                              = "uploading"
	PAUSED_UPLOAD                          = "pausedUP"
	QUEUED_UPLOAD                          = "queuedUP"
	STALLED_UPLOAD                         = "stalledUP"
	CHECKING_UPLOAD                        = "checkingUP"
	FORCED_UPLOAD                          = "forcedUP"
	ALLOCATING                             = "allocating"
	DOWNLOADING                            = "downloading"
	METADATA_DOWNLOAD                      = "metaDL"
	FORCED_METADATA_DOWNLOAD               = "forcedMetaDL"
	PAUSED_DOWNLOAD                        = "pausedDL"
	QUEUED_DOWNLOAD                        = "queuedDL"
	FORCED_DOWNLOAD                        = "forcedDL"
	STALLED_DOWNLOAD                       = "stalledDL"
	CHECKING_DOWNLOAD                      = "checkingDL"
	CHECKING_RESUME_DATA                   = "checkingResumeData"
	MOVING                                 = "moving"
	UNKNOWN                                = "unknown"
)

func (t Torrent) IsDownloading() bool {
	switch t.State {
	case DOWNLOADING,
		METADATA_DOWNLOAD,
		FORCED_METADATA_DOWNLOAD,
		STALLED_DOWNLOAD,
		CHECKING_DOWNLOAD,
		PAUSED_DOWNLOAD,
		QUEUED_DOWNLOAD,
		FORCED_DOWNLOAD:
		return true
	}
	return false
}

func (t Torrent) IsUploading() bool {
	switch t.State {
	case UPLOADING,
		STALLED_UPLOAD,
		CHECKING_UPLOAD,
		QUEUED_UPLOAD,
		FORCED_UPLOAD:
		return true
	}
	return false
}

func (t Torrent) IsComplete() bool {
	switch t.State {
	case UPLOADING,
		STALLED_UPLOAD,
		CHECKING_UPLOAD,
		PAUSED_UPLOAD,
		QUEUED_UPLOAD,
		FORCED_UPLOAD:
		return true
	}
	return false
}

func (t Torrent) IsChecking() bool {
	return t.State == CHECKING_UPLOAD || t.State == CHECKING_DOWNLOAD || t.State == CHECKING_RESUME_DATA

}

func (t Torrent) IsErrored() bool {
	return t.State == MISSING_FILES || t.State == ERROR

}

func (t Torrent) IsPaused() bool {
	return t.State == PAUSED_UPLOAD || t.State == PAUSED_DOWNLOAD
}

// Torrents
// Retrieves list of info for torrents.
// Note: “hashes“ introduced in Web API 2.0.1
//
// @param status_filter: Filter list by all, downloading, completed, paused, active, inactive, resumed
//
//	stalled, stalled_uploading and stalled_downloading added in Web API 2.4.1
//
// @param category: Filter list by category
// @param sort: Sort list by any property returned
// @param reverse: Reverse sorting
// @param limit: Limit length of list
// @param offset: Start of list (if < 0, offset from end of list)
// @param torrent_hashes: Filter list by hash (separate multiple hashes with a '|')
// @param tag: Filter list by tag (empty string means "untagged"; no "tag" param means "any tag"; added in Web API 2.8.3)
// @return: :class:`TorrentInfoList` - `<https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-torrent-list>`_
// noqa: E501
func (t *TorrentsApi) Torrents(data map[string]string) ([]*Torrent, error) {
	resp, err := t.client.post(definition.Torrents, "info", data)
	if err != nil {
		//logrus.Error(err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		//logrus.Error(err)
		return nil, err
	}
	var torrents []*Torrent
	if err = json.Unmarshal(body, &torrents); err != nil {
		//logrus.Error(err)
		return nil, err
	}
	return torrents, nil
}

// addTorrents
// """
// Add one or more torrents by URLs and/or torrent files.
//
// :raises UnsupportedMediaType415Error: if file is not a valid torrent file
// :raises TorrentFileNotFoundError: if a torrent file doesn't exist
// :raises TorrentFilePermissionError: if read permission is denied to torrent file
//
// :return: “Ok.“ for success and “Fails.“ for failure
// """  # noqa: E501
func (t *TorrentsApi) addTorrents(data map[string]string) ([]*Torrent, error) {
	return nil, nil
}

// downloadBase
// :param save_path: location to save the torrent data
// :param cookie: cookie to retrieve torrents by URL
// :param category: category to assign to torrent(s)
// :param is_skip_checking: skip hash checking
// :param is_paused: True to start torrent(s) paused
// :param is_root_folder: True or False to create root folder (superseded by content_layout with v4.3.2)
// :param rename: new name for torrent(s)
// :param upload_limit: upload limit in bytes/second
// :param download_limit: download limit in bytes/second
// :param use_auto_torrent_management: True or False to use automatic torrent management
// :param is_sequential_download: True or False for sequential download
// :param is_first_last_piece_priority: True or False for first and last piece download priority
// :param tags: tag(s) to assign to torrent(s) (added in Web API 2.6.2)
// :param content_layout: Original, Subfolder, or NoSubfolder to control filesystem structure for content (added in Web API 2.7)
// :param ratio_limit: share limit as ratio of upload amt over download amt; e.g. 0.5 or 2.0 (added in Web API 2.8.1)
// :param seeding_time_limit: number of minutes to seed torrent (added in Web API 2.8.1)
// :param download_path: location to download torrent content before moving to save_path (added in Web API 2.8.4)
// :param use_download_path: whether the download_path should be used...defaults to True if download_path is specified (added in Web API 2.8.4)
type downloadBase struct {
	SavePath           string        `json:"savepath"` // 保存文件到：
	Category           string        `json:"category"` // 分类：
	Tags               string        `json:"tags"`
	SkipChecking       bool          `json:"skip_checking"` // 跳过哈希校验
	Paused             bool          `json:"paused"`        // 开始 Torrent
	RootFolder         string        `json:"root_folder"`
	ContentLayout      ContentLayout `json:"contentLayout"` // 内容布局：
	Rename             string        `json:"rename"`        // 重命名 torrent
	UpLimit            string        `json:"upLimit"`       // 限制上传速率
	DlLimit            string        `json:"dlLimit"`       // 限制下载速率
	RatioLimit         string        `json:"ratioLimit"`
	SeedingTimeLimit   string        `json:"seedingTimeLimit"`
	AutoTMM            string        `json:"autoTMM"`
	SequentialDownload bool          `json:"sequentialDownload"` // 按顺序下载
	FirstLastPiecePrio bool          `json:"firstLastPiecePrio"` // 先下载首尾文件块
	DownloadPath       string        `json:"downloadPath"`
	UseDownloadPath    bool          `json:"useDownloadPath"`
}

type ContentLayout = string

const (
	ContentLayoutOriginal    = "Original"    // 原始
	ContentLayoutSubFolder   = "Subfolder"   // 创建子文件夹
	ContentLayoutNoSubFolder = "NoSubfolder" // 不创建子文件夹
)

// MagnetDLConfig
// :param urls: single instance or an iterable of URLs (http://, https://, magnet: and bc://bt/)
type MagnetDLConfig struct {
	Url string `json:"urls"`
	downloadBase
}

// TorrentDLConfig
// :param torrent_files: several options are available to send torrent files to qBittorrent:
//   - single instance of bytes: useful if torrent file already read from disk or downloaded from internet.
//   - single instance of file handle to torrent file: use open(<filepath>, 'rb') to open the torrent file.
//   - single instance of a filepath to torrent file: e.g. '/home/user/torrent_filename.torrent'
//   - an iterable of the single instances above to send more than one torrent file
//   - dictionary with key/value pairs of torrent name and single instance of above object
//     Note: The torrent name in a dictionary is useful to identify which torrent file
//     errored. qBittorrent provides back that name in the error text. If a torrent
//     name is not provided, then the name of the file will be used. And in the case of
//     bytes (or if filename cannot be determined), the value 'torrent__n' will be used
type TorrentDLConfig struct {
	downloadBase
}

func (t *TorrentsApi) DownloadFromLink(cfg MagnetDLConfig) ([]*Torrent, error) {
	return nil, nil
}
