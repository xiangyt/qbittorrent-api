package qbittorrent_api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type torrentsApi struct {
	client *Client
}

type Torrent struct {
	AddedOn                int          `json:"added_on"`
	AmountLeft             int64        `json:"amount_left"`
	AutoTmm                bool         `json:"auto_tmm"`
	Availability           float64      `json:"availability"`
	Category               string       `json:"category"`
	Completed              int64        `json:"completed"`
	CompletionOn           int          `json:"completion_on"`
	ContentPath            string       `json:"content_path"`
	DlLimit                int          `json:"dl_limit"`
	DlSpeed                int          `json:"dlspeed"`
	DownloadPath           string       `json:"download_path"`
	Downloaded             int64        `json:"downloaded"`
	DownloadedSession      int          `json:"downloaded_session"`
	Eta                    int          `json:"eta"`
	FirstLastPiecePriority bool         `json:"f_l_piece_prio"`
	ForceStart             bool         `json:"force_start"`
	Hash                   string       `json:"hash"`
	InfoHashV1             string       `json:"infohash_v1"`
	InfoHashV2             string       `json:"infohash_v2"`
	LastActivity           int          `json:"last_activity"`
	MagnetUri              string       `json:"magnet_uri"`
	MaxRatio               int          `json:"max_ratio"`
	MaxSeedingTime         int          `json:"max_seeding_time"`
	Name                   string       `json:"name"`
	NumComplete            int          `json:"num_complete"`
	NumIncomplete          int          `json:"num_incomplete"`
	NumLeechs              int          `json:"num_leechs"`
	NumSeeds               int          `json:"num_seeds"`
	Priority               int          `json:"priority"`
	Progress               float64      `json:"progress"`
	Ratio                  float64      `json:"ratio"`
	RatioLimit             int          `json:"ratio_limit"`
	SavePath               string       `json:"save_path"`
	SeedingTime            int          `json:"seeding_time"`
	SeedingTimeLimit       int          `json:"seeding_time_limit"`
	SeenComplete           int          `json:"seen_complete"`
	SequentialDownload     bool         `json:"seq_dl"`
	Size                   int64        `json:"size"`
	State                  TorrentState `json:"state"`
	SuperSeeding           bool         `json:"super_seeding"`
	Tags                   string       `json:"tags"`
	TimeActive             int          `json:"time_active"`
	TotalSize              int64        `json:"total_size"`
	Tracker                string       `json:"tracker"`
	TrackersCount          int          `json:"trackers_count"`
	UpLimit                int          `json:"up_limit"`
	Uploaded               int64        `json:"uploaded"`
	UploadedSession        int          `json:"uploaded_session"`
	UpSpeed                int          `json:"upspeed"`
}

type TorrentState = string

const (
	StateError                  TorrentState = "error"
	StateMissingFiles           TorrentState = "missingFiles"
	StateUploading              TorrentState = "uploading"
	StatePausedUpload           TorrentState = "pausedUP"
	StateQueuedUpload           TorrentState = "queuedUP"
	StateStalledUpload          TorrentState = "stalledUP"
	StateCheckingUpload         TorrentState = "checkingUP"
	StateForcedUpload           TorrentState = "forcedUP"
	StateAllocating             TorrentState = "allocating"
	StateDownloading            TorrentState = "downloading"
	StateMetadataDownload       TorrentState = "metaDL"
	StateForcedMetadataDownload TorrentState = "forcedMetaDL"
	StatePausedDownload         TorrentState = "pausedDL"
	StateQueuedDownload         TorrentState = "queuedDL"
	StateForcedDownload         TorrentState = "forcedDL"
	StateStalledDownload        TorrentState = "stalledDL"
	StateCheckingDownload       TorrentState = "checkingDL"
	StateCheckingResumeData     TorrentState = "checkingResumeData"
	StateMoving                 TorrentState = "moving"
	StateUnknown                TorrentState = "unknown"
)

func (t Torrent) IsDownloading() bool {
	switch t.State {
	case StateDownloading,
		StateMetadataDownload,
		StateForcedMetadataDownload,
		StateStalledDownload,
		StateCheckingDownload,
		StatePausedDownload,
		StateQueuedDownload,
		StateForcedDownload:
		return true
	}
	return false
}

func (t Torrent) IsUploading() bool {
	switch t.State {
	case StateUploading,
		StateStalledUpload,
		StateCheckingUpload,
		StateQueuedUpload,
		StateForcedUpload:
		return true
	}
	return false
}

func (t Torrent) IsComplete() bool {
	switch t.State {
	case StateUploading,
		StateStalledUpload,
		StateCheckingUpload,
		StatePausedUpload,
		StateQueuedUpload,
		StateForcedUpload:
		return true
	}
	return false
}

func (t Torrent) IsChecking() bool {
	return t.State == StateCheckingUpload || t.State == StateCheckingDownload || t.State == StateCheckingResumeData

}

func (t Torrent) IsErrored() bool {
	return t.State == StateMissingFiles || t.State == StateError

}

func (t Torrent) IsPaused() bool {
	return t.State == StatePausedUpload || t.State == StatePausedDownload
}

// Filter
// param status_filter: Filter list by all, downloading, completed, paused, active, inactive, resumed
//
//	stalled, stalled_uploading and stalled_downloading added in Web API 2.4.1
//
// param category: Filter list by category
// param sort: Sort list by any property returned
// param reverse: Reverse sorting
// param limit: Limit length of list
// param offset: Start of list (if < 0, offset from end of list)
// param torrent_hashes: Filter list by hash (separate multiple hashes with a '|')
// param tag: Filter list by tag (empty string means "untagged"; no "tag" param means "any tag"; added in Web API 2.8.3)
type Filter struct {
	StatusFilter TorrentState `json:"filter"`
	Category     string       `json:"category"`
	Sort         FilterSort   `json:"sort"`
	Reverse      bool         `json:"reverse"`
	Limit        int          `json:"limit"`
	Offset       int          `json:"offset"`
	Hashes       []string     `json:"torrent_hashes"`
	Tag          string       `json:"tag"`
}

type FilterSort = string

const (
	FilterSortPriority = "priority"
	FilterSortName     = "name"
	FilterSortProgress = "progress"
)

func (f *Filter) toMap() map[string]string {
	var data = map[string]string{
		"filter": f.StatusFilter,
		//"category": f.Category,
		"sort":    f.Sort,
		"reverse": "false",
		//"limit": limit,
		//"offset": offset,
		"hashes": strings.Join(f.Hashes, "|"),
		//"tag":    f.Tag,
	}
	if f.Reverse && f.Sort != "" {
		data["reverse"] = "true"
	}

	if f.Sort == "" {
		data["sort"] = FilterSortPriority
	}
	if f.Limit > 0 {
		data["limit"] = strconv.Itoa(f.Limit)
	}
	if f.Offset != 0 {
		data["offset"] = strconv.Itoa(f.Offset)
	}
	if f.Category != "" {
		data["category"] = f.Category
	}
	if f.Tag != "" {
		data["tag"] = f.Tag
	}

	return data
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
func (t *torrentsApi) Torrents(params *Filter) ([]*Torrent, error) {
	_, body, err := t.postForTorrent("info", params.toMap())
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

type Category struct {
	Name         string `json:"name"`
	SavePath     string `json:"savePath"`
	DownloadPath string `json:"download_path"`
}

// GetAllCategories Retrieve all category definitions.
func (t *torrentsApi) GetAllCategories() ([]*Category, error) {
	_, body, err := t.getForTorrent(actionGetAllCategories, nil)
	if err != nil {
		return nil, err
	}

	var categories = map[string]*Category{}
	if err = json.Unmarshal(body, &categories); err != nil {
		//logrus.Error(err)
		return nil, err
	}

	var res []*Category
	for _, category := range categories {
		res = append(res, category)
	}
	return res, nil
}

// GetAllTags Retrieve all category definitions.
func (t *torrentsApi) GetAllTags() ([]string, error) {
	_, body, err := t.getForTorrent(actionGetAllTags, nil)
	if err != nil {
		return nil, err
	}

	var tags []string
	if err = json.Unmarshal(body, &tags); err != nil {
		//logrus.Error(err)
		return nil, err
	}
	return tags, nil
}

// DownloadBaseConfig
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
type DownloadBaseConfig struct {
	SavePath           string        `json:"savepath"`           // 保存文件到：
	Category           string        `json:"category"`           // 分类：
	Tags               []string      `json:"tags"`               // 标签
	SkipChecking       bool          `json:"skip_checking"`      // 跳过哈希校验
	Paused             bool          `json:"paused"`             // 开始 Torrent
	ContentLayout      ContentLayout `json:"contentLayout"`      // 内容布局：
	Rename             string        `json:"rename"`             // 重命名 torrent
	UpLimit            int           `json:"upLimit"`            // 限制上传速率 bytes/second
	DlLimit            int           `json:"dlLimit"`            // 限制下载速率 bytes/second
	AutoTMM            bool          `json:"autoTMM"`            // 自动 Torrent 管理
	SequentialDownload bool          `json:"sequentialDownload"` // 按顺序下载
	FirstLastPiecePrio bool          `json:"firstLastPiecePrio"` // 先下载首尾文件块
	//RootFolder         bool        `json:"root_folder"`
	//RatioLimit         float64        `json:"ratioLimit"`
	//SeedingTimeLimit   string        `json:"seedingTimeLimit"`
	//DownloadPath       string        `json:"downloadPath"`
	//UseDownloadPath    bool          `json:"useDownloadPath"`
}

func (cfg DownloadBaseConfig) toMap() map[string]string {
	data := map[string]string{
		"autoTMM": "false",
		//"savepath": cfg.SavePath,
		"rename":   cfg.Rename,
		"category": cfg.Category,
		"tags":     strings.Join(cfg.Tags, ","),
		"paused":   "false",
		//"sequentialDownload": "false",
		//"firstLastPiecePrio": "false",
		"contentLayout": cfg.ContentLayout,
		//"dlLimit":            "NaN",
		//"upLimit":            "NaN",
	}
	if cfg.AutoTMM {
		data["autoTMM"] = "true"
	} else {
		data["savepath"] = cfg.SavePath
	}
	if cfg.Paused {
		data["paused"] = "true"
	}
	if cfg.ContentLayout == "" {
		data["contentLayout"] = ContentLayoutOriginal
	}

	if cfg.DlLimit > 0 {
		data["dlLimit"] = strconv.Itoa(cfg.DlLimit)
	}
	if cfg.UpLimit > 0 {
		data["upLimit"] = strconv.Itoa(cfg.UpLimit)
	}
	if cfg.SkipChecking {
		data["skip_checking"] = "true"
	}
	if cfg.SequentialDownload {
		data["sequentialDownload"] = "true"
	}
	if cfg.FirstLastPiecePrio {
		data["firstLastPiecePrio"] = "true"
	}
	return data
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
	Urls []string `json:"urls"`
	DownloadBaseConfig
}

func (cfg MagnetDLConfig) toMap() map[string]string {
	data := cfg.DownloadBaseConfig.toMap()
	data["urls"] = strings.Join(cfg.Urls, "\n")
	return data
}

// DownloadFromLink
// Add one or more torrents by URLs and/or torrent files.
//
// :raises UnsupportedMediaType415Error: if file is not a valid torrent file
// :raises TorrentFileNotFoundError: if a torrent file doesn't exist
// :raises TorrentFilePermissionError: if read permission is denied to torrent file
//
// :return: “Ok.“ for success and “Fails.“ for failure
func (t *torrentsApi) DownloadFromLink(cfg MagnetDLConfig) error {
	resp, err := t.client.postMultipartData(apiNameTorrents, "add", cfg.toMap())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	//fmt.Println(string(body))
	//fmt.Println(resp.StatusCode)
	if resp.StatusCode == http.StatusOK {
		if string(body) == "Ok." {
			return nil
		} else {
			return errors.New("DownloadFromLink Fails.")
		}
	}

	return handleResponsesErr(resp.StatusCode)
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
	File string `json:"file"`
	DownloadBaseConfig
}

func (t *torrentsApi) DownloadFromFile(cfg TorrentDLConfig) error {
	resp, err := t.client.postMultipartFile(apiNameTorrents, "add", cfg.File, cfg.toMap())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	//fmt.Println(string(body))
	//fmt.Println(resp.StatusCode)
	if resp.StatusCode == http.StatusOK {
		if string(body) == "Ok." {
			return nil
		} else {
			return errors.New("DownloadFromLink Fails.")
		}
	}

	return handleResponsesErr(resp.StatusCode)
}

func (t *torrentsApi) getForTorrent(path string, data map[string]string) (int, []byte, error) {
	resp, err := t.client.request.get(apiNameTorrents, path, data)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	//fmt.Println(string(body))
	//fmt.Println(resp.StatusCode)
	if err := handleResponsesErr(resp.StatusCode); err != nil {
		return 0, nil, err
	}
	return resp.StatusCode, body, nil
}

func (t *torrentsApi) postForTorrent(path string, data map[string]string) (int, []byte, error) {
	resp, err := t.client.request.post(apiNameTorrents, path, data)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	//fmt.Println(string(body))
	//fmt.Println(resp.StatusCode)
	return resp.StatusCode, body, nil
}

type torrentsAction = string

const (
	actionPause            torrentsAction = "pause"
	actionResume           torrentsAction = "resume"
	actionDelete           torrentsAction = "delete"
	actionRecheck          torrentsAction = "recheck"
	actionReAnnounce       torrentsAction = "reannounce"
	actionIncreasePriority torrentsAction = "increasePrio"
	actionDecreasePriority torrentsAction = "decreasePrio"
	actionTopPriority      torrentsAction = "topPrio"
	actionBottomPriority   torrentsAction = "bottomPrio"
	actionGetAllCategories torrentsAction = "categories"
	actionSetCategory      torrentsAction = "setCategory"
	actionCreateCategory   torrentsAction = "createCategory"
	actionEditCategory     torrentsAction = "editCategory"
	actionRemoveCategories torrentsAction = "removeCategories"
	actionGetAllTags       torrentsAction = "tags"
	actionAddTags          torrentsAction = "addTags"
	actionRemoveTags       torrentsAction = "removeTags"
	actionCreateTags       torrentsAction = "createTags"
	actionDeleteTags       torrentsAction = "deleteTags"
)
const (
	actionToggleSequentialDownload     torrentsAction = "toggleSequentialDownload"
	actionToggleFirstLastPiecePriority torrentsAction = "toggleFirstLastPiecePrio"
)

var actionForAll = []string{"all"}

type action struct {
	param  map[string]string
	method torrentsAction
}

func newAction(method torrentsAction) *action {
	return &action{param: map[string]string{}, method: method}
}

func (a *action) withHashes(ts []*Torrent) *action {
	var hashes []string
	for _, torrent := range ts {
		hashes = append(hashes, torrent.Hash)
	}
	a.param["hashes"] = strings.Join(hashes, "|")
	return a
}

func (a *action) withTags(tags []string) *action {
	return a.setParam("tags", strings.Join(tags, ","))
}

func (a *action) setParam(key, val string) *action {
	a.param[key] = val
	return a
}

func (a *action) setBoolParam(key string, val bool) *action {
	a.param[key] = "false"
	if val {
		a.setParam(key, "true")
	}
	return a
}

func (a *action) setIntParam(key string, val int) *action {
	a.param[key] = strconv.Itoa(val)
	return a
}

func (t *torrentsApi) actionByHashes(hashes []string, act *action) error {
	return t.actionByParams(act.setParam("hashes", strings.Join(hashes, "|")))
}

func (t *torrentsApi) actionByParams(act *action) error {
	statusCode, _, err := t.client.postForTorrent(act.method, act.param)
	if err != nil {
		return err
	}

	if statusCode == http.StatusOK {
		return nil
	}
	return handleResponsesErr(statusCode)
}

// Pause pause one or more torrents in qBittorrent.
func (t *torrentsApi) Pause(ts ...*Torrent) error {
	if len(ts) == 0 {
		return nil
	}

	return t.actionByParams(newAction(actionPause).withHashes(ts))
}

// PauseByHashes pause torrents in qBittorrent by hashes.
func (t *torrentsApi) PauseByHashes(hashes []string) error {
	return t.actionByHashes(hashes, newAction(actionPause))
}

// PauseAll pause all torrents in qBittorrent.
func (t *torrentsApi) PauseAll() error {
	return t.PauseByHashes(actionForAll)
}

// Resume resume one or more torrents in qBittorrent.
func (t *torrentsApi) Resume(ts ...*Torrent) error {
	if len(ts) == 0 {
		return nil
	}

	return t.actionByParams(newAction(actionResume).withHashes(ts))
}

// ResumeByHashes resume torrents in qBittorrent by hashes.
func (t *torrentsApi) ResumeByHashes(hashes []string) error {
	return t.actionByHashes(hashes, newAction(actionResume))
}

// ResumeAll resume all torrents in qBittorrent.
func (t *torrentsApi) ResumeAll() error {
	return t.ResumeByHashes(actionForAll)
}

// Delete
// Remove a torrent from qBittorrent and optionally delete its files.
// param delete_files: True to delete the torrent's files
// param torrent_hashes: single torrent hash or list of torrent hashes. Or “all“ for all torrents.
// return: None
func (t *torrentsApi) Delete(deleteFiles bool, ts ...*Torrent) error {
	if len(ts) == 0 {
		return nil
	}

	return t.actionByParams(newAction(actionDelete).withHashes(ts).setBoolParam("deleteFiles", deleteFiles))
}

func (t *torrentsApi) DeleteByHashes(deleteFiles bool, hashes []string) error {
	return t.actionByHashes(hashes, newAction(actionDelete).setBoolParam("deleteFiles", deleteFiles))
}

func (t *torrentsApi) DeleteAll(deleteFiles bool) error {
	return t.DeleteByHashes(deleteFiles, actionForAll)
}

// Recheck a torrent in qBittorrent.
func (t *torrentsApi) Recheck(ts ...*Torrent) error {
	if len(ts) == 0 {
		return nil
	}

	return t.actionByParams(newAction(actionRecheck).withHashes(ts))
}

func (t *torrentsApi) RecheckByHashes(hashes []string) error {
	return t.actionByHashes(hashes, newAction(actionRecheck))
}

func (t *torrentsApi) RecheckAll() error {
	return t.RecheckByHashes(actionForAll)
}

// ReAnnounce a torrent in qBittorrent.
func (t *torrentsApi) ReAnnounce(ts ...*Torrent) error {
	if len(ts) == 0 {
		return nil
	}

	return t.actionByParams(newAction(actionReAnnounce).withHashes(ts))
}

func (t *torrentsApi) ReAnnounceByHashes(hashes []string) error {
	return t.actionByHashes(hashes, newAction(actionReAnnounce))
}

func (t *torrentsApi) ReAnnounceAll() error {
	return t.ReAnnounceByHashes(actionForAll)
}

// IncreasePriority
// Increase the priority of a torrent. Torrent Queuing must be enabled.
// 向上移动队列
func (t *torrentsApi) IncreasePriority(ts ...*Torrent) error {
	if len(ts) == 0 {
		return nil
	}

	return t.actionByParams(newAction(actionIncreasePriority).withHashes(ts))
}

func (t *torrentsApi) IncreasePriorityByHashes(hashes []string) error {
	return t.actionByHashes(hashes, newAction(actionIncreasePriority))
}

//func (t *torrentsApi) IncreasePriorityAll() error {
//	return t.IncreasePriorityByHashes(actionForAll)
//}

// DecreasePriority
// Decrease the priority of a torrent. Torrent Queuing must be enabled.
// 向下移动队列
func (t *torrentsApi) DecreasePriority(ts ...*Torrent) error {
	if len(ts) == 0 {
		return nil
	}

	return t.actionByParams(newAction(actionDecreasePriority).withHashes(ts))
}

func (t *torrentsApi) DecreasePriorityByHashes(hashes []string) error {
	return t.actionByHashes(hashes, newAction(actionDecreasePriority))
}

//func (t *torrentsApi) DecreasePriorityAll() error {
//	return t.DecreasePriorityByHashes(actionForAll)
//}

// TopPriority
// Set torrent as highest priority. Torrent Queuing must be enabled.
// 移动到队列顶部
func (t *torrentsApi) TopPriority(ts ...*Torrent) error {
	if len(ts) == 0 {
		return nil
	}

	return t.actionByParams(newAction(actionTopPriority).withHashes(ts))
}

func (t *torrentsApi) TopPriorityByHashes(hashes []string) error {
	return t.actionByHashes(hashes, newAction(actionTopPriority))
}

//func (t *torrentsApi) TopPriorityAll() error {
//	return t.TopPriorityByHashes(actionForAll)
//}

// BottomPriority
// Set torrent as highest priority. Torrent Queuing must be enabled.
// 移动到队列底部
func (t *torrentsApi) BottomPriority(ts ...*Torrent) error {
	if len(ts) == 0 {
		return nil
	}

	return t.actionByParams(newAction(actionBottomPriority).withHashes(ts))
}

func (t *torrentsApi) BottomPriorityByHashes(hashes []string) error {
	return t.actionByHashes(hashes, newAction(actionBottomPriority))
}

//func (t *torrentsApi) BottomPriorityAll() error {
//	return t.BottomPriorityByHashes(actionForAll)
//}

// SetCategory  Set a category for one or more torrents.
func (t *torrentsApi) SetCategory(category string, ts ...*Torrent) error {
	if len(ts) == 0 {
		return nil
	}

	return t.actionByParams(newAction(actionSetCategory).withHashes(ts).setParam("category", category))
}

func (t *torrentsApi) SetCategoryByHashes(category string, hashes []string) error {
	return t.actionByHashes(hashes, newAction(actionSetCategory).setParam("category", category))
}

func (t *torrentsApi) SetCategoryForAll(category string) error {
	return t.SetCategoryByHashes(category, actionForAll)
}

// CreateCategory Create a new torrent category.
// param name: name for new category
// param save_path: location to save torrents for this category (added in Web API 2.1.0)
// param download_path: download location for torrents with this category
// param enable_download_path: True or False to enable or disable download path
func (t *torrentsApi) CreateCategory(category Category) error {
	act := newAction(actionCreateCategory).
		setParam("category", category.Name).
		setParam("savePath", category.SavePath).
		setParam("downloadPath", category.DownloadPath).
		setBoolParam("downloadPathEnabled", category.DownloadPath != "")
	return t.actionByParams(act)
}

// EditCategory Edit an existing category.
func (t *torrentsApi) EditCategory(category Category) error {
	act := newAction(actionEditCategory).
		setParam("category", category.Name).
		setParam("savePath", category.SavePath).
		setParam("downloadPath", category.DownloadPath).
		setBoolParam("downloadPathEnabled", category.DownloadPath != "")
	return t.actionByParams(act)
}

// RemoveCategories Delete one or more categories.
func (t *torrentsApi) RemoveCategories(categories []string) error {
	return t.actionByParams(newAction(actionRemoveCategories).setParam("categories", strings.Join(categories, "\n")))
}

// CreateTags Create one or more tags.
func (t *torrentsApi) CreateTags(tags ...string) error {
	if len(tags) == 0 {
		return nil
	}
	return t.actionByParams(newAction(actionCreateTags).withTags(tags))
}

// DeleteTags Delete one or more tags.
func (t *torrentsApi) DeleteTags(tags ...string) error {
	if len(tags) == 0 {
		return nil
	}
	return t.actionByParams(newAction(actionDeleteTags).withTags(tags))
}

// AddTags Add one or more tags to one or more torrents.
// Note: Tags that do not exist will be created on-the-fly.
//
// param tags: tag name or list of tags
// param torrent_hashes: single torrent hash or list of torrent hashes. Or “all“ for all torrents.
func (t *torrentsApi) AddTags(tags []string, ts ...*Torrent) error {
	return t.actionByParams(newAction(actionAddTags).withHashes(ts).withTags(tags))
}

func (t *torrentsApi) AddTagsByHashes(tags, hashes []string) error {
	return t.actionByHashes(hashes, newAction(actionAddTags).withTags(tags))
}

func (t *torrentsApi) AddTagsForAll(tags []string) error {
	return t.AddTagsByHashes(tags, actionForAll)
}

// RemoveTags Remove one or more tags to one or more torrents.
func (t *torrentsApi) RemoveTags(tags []string) error {
	return t.actionByParams(newAction(actionRemoveTags).withTags(tags))
}

func (t *torrentsApi) RemoveTagsByHashes(tags, hashes []string) error {
	return t.actionByHashes(hashes, newAction(actionRemoveTags).withTags(tags))
}

func (t *torrentsApi) RemoveTagsForAll(tags []string) error {
	return t.RemoveTagsByHashes(tags, actionForAll)
}

// ToggleFirstLastPiecePriority
// Increase the priority of a torrent. Torrent Queuing must be enabled.
// 选中/取消 先下载首尾文件块
func (t *torrentsApi) ToggleFirstLastPiecePriority(ts ...*Torrent) error {
	if len(ts) == 0 {
		return nil
	}

	return t.actionByParams(newAction(actionToggleFirstLastPiecePriority).withHashes(ts))
}

func (t *torrentsApi) ToggleFirstLastPiecePriorityByHashes(hashes []string) error {
	return t.actionByHashes(hashes, newAction(actionToggleFirstLastPiecePriority))
}

func (t *torrentsApi) ToggleFirstLastPiecePriorityForAll() error {
	return t.ToggleFirstLastPiecePriorityByHashes(actionForAll)
}

// ToggleSequentialDownload
// Increase the priority of a torrent. Torrent Queuing must be enabled.
// 选中/取消 按顺序下载
func (t *torrentsApi) ToggleSequentialDownload(ts ...*Torrent) error {
	if len(ts) == 0 {
		return nil
	}

	return t.actionByParams(newAction(actionToggleSequentialDownload).withHashes(ts))
}

func (t *torrentsApi) ToggleSequentialDownloadByHashes(hashes []string) error {
	return t.actionByHashes(hashes, newAction(actionToggleSequentialDownload))
}

func (t *torrentsApi) ToggleSequentialDownloadForAll() error {
	return t.ToggleSequentialDownloadByHashes(actionForAll)
}
