package milb

import "errors"

type MusicEntry struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

type MusicManager struct {
	musics []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("下标越界")
	}
	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) *MusicEntry {
	for _, v := range m.musics {
		if v.Name == name {
			return &v
		}
	}
	return nil
}

func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) (music *MusicEntry) {
	if index < 0 || index >= len(m.musics) {
		return nil
	}
	re := &m.musics[index]
	m.musics = append(m.musics[:index], m.musics[index+1:]...)
	return re
}
