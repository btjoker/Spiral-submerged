package library

import "errors"

var (
	errIndex = errors.New("Index out of range")
)

// MusicManager .
type MusicManager struct {
	musics []*MusicEntry
}

// NewMusicManager .
func NewMusicManager() *MusicManager {
	return &MusicManager{make([]*MusicEntry, 0)}
}

// Len .
func (m *MusicManager) Len() int {
	return len(m.musics)
}

// Get .
func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= m.Len() {
		return nil, errIndex
	}
	return m.musics[index], nil
}

// Find .
func (m *MusicManager) Find(name string) *MusicEntry {
	if m.Len() != 0 {
		for _, m := range m.musics {
			if m.Name == name {
				return m
			}
		}
	}
	return nil
}

// Add .
func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, music)
}

// Remove .
func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= m.Len() {
		return nil
	}

	removedMusic := m.musics[index]

	if index < m.Len()-1 {
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	} else if index == 0 {
		m.musics = make([]*MusicEntry, 0)
	} else {
		m.musics = m.musics[:index-1]
	}
	return removedMusic
}

// RemoveByName .
func (m *MusicManager) RemoveByName(name string) *MusicEntry {
	var index int
	for k, v := range m.musics {
		if v.Name == name {
			index = k
			break
		}
	}
	removedMusic := m.musics[index]

	if index == 0 && m.musics[index].Name == name {
		m.musics = m.musics[1:]
	} else if index == 0 {
		return nil
	} else {
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	}
	return removedMusic
}
