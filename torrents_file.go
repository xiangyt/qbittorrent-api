package qbittorrent_api

import (
	"encoding/json"
	"strconv"
	"strings"
)

const (
	actionGetFiles        torrentsAction = "files"
	actionSetFilePriority torrentsAction = "filePrio"
)

type FilePriority = int

const (
	FilePriorityDoNotDL FilePriority = 0 // Do not download
	FilePriorityNormal  FilePriority = 1 // Normal priority
	FilePriorityHigh    FilePriority = 6 // High priority
	FilePriorityMaximal FilePriority = 7 // Maximal priority
)

type TorrentFile struct {
	Index        int          `json:"index"`        // File index
	Name         string       `json:"name"`         // File name (including relative path)
	Size         int64        `json:"size"`         // File size (bytes)
	Progress     float64      `json:"progress"`     // File progress (percentage/100)
	Priority     FilePriority `json:"priority"`     // File priority. See possible values here below
	IsSeed       bool         `json:"is_seed"`      // True if file is seeding/complete
	PieceRange   []int        `json:"piece_range"`  // The first number is the starting piece index and the second number is the ending piece index (inclusive)
	Availability float64      `json:"availability"` // Percentage of file pieces currently available (percentage/100)
}

type TorrentFiles []*TorrentFile

func (tfs TorrentFiles) Len() int {
	return len(tfs)
}

func (tfs TorrentFiles) Less(i, j int) bool {
	return tfs[i].Size < tfs[j].Size
}

func (tfs TorrentFiles) Swap(i, j int) {
	tfs[i], tfs[j] = tfs[j], tfs[i]
}

func (tfs TorrentFiles) Filter(fn func(tf *TorrentFile) bool) TorrentFiles {
	var Filtered TorrentFiles
	for _, tf := range tfs {
		if fn(tf) {
			Filtered = append(Filtered, tf)
		}
	}
	return Filtered
}

// GetAllTorrentFiles
// Get torrent contents.
func (t *torrentsApi) GetAllTorrentFiles(torrent *Torrent) ([]*TorrentFile, error) {
	return t.GetAllTorrentFilesByHash(torrent.Hash)
}

// GetAllTorrentFilesByHash
// Get torrent contents.
func (t *torrentsApi) GetAllTorrentFilesByHash(hash string) ([]*TorrentFile, error) {
	_, body, err := t.getForTorrent(actionGetFiles, map[string]string{
		"hash": hash,
	})
	if err != nil {
		return nil, err
	}

	var tfs []*TorrentFile
	if err = json.Unmarshal(body, &tfs); err != nil {
		//logrus.Error(err)
		return nil, err
	}
	return tfs, nil
}

func (t *torrentsApi) SetFilePriority(hash string, tfs []*TorrentFile, priority FilePriority) error {
	if len(tfs) == 0 {
		return nil
	}
	var ids []string
	for _, tf := range tfs {
		ids = append(ids, strconv.Itoa(tf.Index))
	}

	return t.actionByParams(newAction(actionSetFilePriority).
		setParam("hash", hash).
		setParam("id", strings.Join(ids, "|")).
		setIntParam("priority", priority))
}
