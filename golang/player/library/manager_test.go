package library

import "testing"

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Errorf("NewMusicManager failed.")
	}
	if mm.Len() != 0 {
		t.Errorf("NewMusicManager failed, not empty.")
	}

	m0 := &MusicEntry{
		ID:     "1",
		Name:   "My Heart Will Go On",
		Artist: "Celion Dion",
		Genre:  "Pop",
		Source: "http://qbox.me/24501234",
		Type:   "MP3",
	}

	mm.Add(m0)

	if mm.Len() != 1 {
		t.Errorf("NewMusicManager.Add() failed.")
	}

	m := mm.Find(m0.Name)
	if m == nil {
		t.Errorf("NewMusicManager.Find() failed.")
	}

	if m.ID != m0.ID || m.Artist != m0.Artist ||
		m.Name != m0.Name || m.Genre != m0.Genre ||
		m.Source != m0.Source || m.Type != m0.Type {
		t.Errorf("NewMusicManager.Find() failed. Found item mismatch.")
	}

	m, err := mm.Get(0)
	if m == nil {
		t.Errorf("NewMusicManager.Get() failed. %v", err)
	}

	m = mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Errorf("NewMusicManager.Remove() failed.")
	}
}
