package qbittorrent_api

import (
	"fmt"
	"testing"
)

const (
	host = "http://localhost:8080"
	name = "admin"
	pwd  = "xxx"
)

func TestNewClient(t *testing.T) {
	c := NewClient(host)
	//c.Login()
	if ts, err := c.Torrents(&Filter{
		//StatusFilter: "downloading",
		//Category:      "",
		//Sort: FilterSortName,
		//Reverse: true,
		//Limit:  1,
		//Offset: -4,
		//TorrentHashes: nil,
		//Tag: "tag2",
	}); err != nil {
		t.Error(err)
	} else {
		for _, torrent := range ts {
			t.Log(torrent)
		}
	}
}

func TestDownloadFromLink(t *testing.T) {
	c := NewClient(host)
	//c.Login()
	if err := c.DownloadFromLink(MagnetDLConfig{
		Urls: []string{"magnet:?xt=urn:btih:076e0288a6147459d261ee1a55c18f1f91a241c2"},
		DownloadBaseConfig: DownloadBaseConfig{
			SavePath: "/bt",
			Tags:     []string{"tag1", "tag2"},
			Category: "Category",
			DlLimit:  512 * 1024,
		},
	}); err != nil {
		t.Error(err)
	}
}

func TestPauseAll(t *testing.T) {
	c := NewClient(host)
	//if err := c.Login(); err != nil {
	//	t.Error(err)
	//}
	if err := c.PauseAll(); err != nil {
		t.Error(err)
	}
}

func TestResumeAll(t *testing.T) {
	c := NewClient(host)
	//if err := c.Login(); err != nil {
	//	t.Error(err)
	//}
	if err := c.ResumeAll(); err != nil {
		t.Error(err)
	}
}

func TestPriorityByHashes(t *testing.T) {
	c := NewClient(host)
	//if err := c.TopPriorityByHashes([]string{"e40127b663555092b5ac7b1f621cb2a7364adbe1"}); err != nil {
	//	t.Error(err)
	//}

	//if err := c.IncreasePriorityByHashes([]string{"e40127b663555092b5ac7b1f621cb2a7364adbe1"}); err != nil {
	//	t.Error(err)
	//}

	if err := c.DecreasePriorityByHashes([]string{"e40127b663555092b5ac7b1f621cb2a7364adbe1"}); err != nil {
		t.Error(err)
	}

	//if err := c.BottomPriorityByHashes([]string{"e40127b663555092b5ac7b1f621cb2a7364adbe1"}); err != nil {
	//	t.Error(err)
	//}
}

func TestToggle(t *testing.T) {
	c := NewClient(host)
	if err := c.ToggleFirstLastPiecePriorityByHashes([]string{"e40127b663555092b5ac7b1f621cb2a7364adbe1"}); err != nil {
		t.Error(err)
	}

	if err := c.ToggleSequentialDownloadByHashes([]string{"e40127b663555092b5ac7b1f621cb2a7364adbe1"}); err != nil {
		t.Error(err)
	}
}

func TestCategories(t *testing.T) {
	c := NewClient(host)
	//if err := c.CreateCategory(Category{
	//	Name:         "test2",
	//	SavePath:     "/bt",
	//	DownloadPath: "/test",
	//}); err != nil {
	//	t.Error(err)
	//}

	//if err := c.EditCategory(Category{
	//	Name:         "test1",
	//	SavePath:     "/bt",
	//	DownloadPath: "/test",
	//}); err != nil {
	//	t.Error(err)
	//}

	if cs, err := c.GetAllCategories(); err != nil {
		t.Error(err)
	} else {
		for _, category := range cs {
			t.Log(fmt.Sprintf("%+v", category))
		}
	}
}

func TestTags(t *testing.T) {
	c := NewClient(host)
	if err := c.DeleteTags("tag2"); err != nil {
		t.Error(err)
	}

	if err := c.AddTagsByHashes([]string{"tag2"}, []string{"e40127b663555092b5ac7b1f621cb2a7364adbe1"}); err != nil {
		t.Error(err)
	}

	if tags, err := c.GetAllTags(); err != nil {
		t.Error(err)
	} else {
		t.Log(fmt.Sprintf("%+v", tags))
	}
}

func TestTorrentFiles(t *testing.T) {
	c := NewClient(host)
	if tfs, err := c.GetAllTorrentFilesByHash("6ec865593c0b4ab75c6264c30c60dc1c59ace7c0"); err != nil {
		t.Error(err)
	} else {
		for _, tf := range tfs {
			t.Log(fmt.Sprintf("%+v", tf))
		}
	}

}
