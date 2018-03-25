package library

import (
	"fmt"
	"errors"
)
// 音乐实体
type MusicEntry struct {
	Id string
	Name string
	Artist string
	Source string
	Type string
}
// 数据组切片,存储歌曲
type MusicManager struct {
	musics [] MusicEntry
}

// 新建歌曲管理器
func NewMusicManager() *MusicManager {
	return &MusicManager{ make([] MusicEntry, 0) }
}
/**
 * Method 
 * func (_ receiver_type) methodName(parameter_list) (return_value_list) { ... }
 **/
// 统计歌曲数量
func (m *MusicManager) Len() int {
	return len(m.musics)
}
// 根据索引获取指定歌曲信息
func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range")
	}
	return &m.musics[index], nil
}
// 根据歌名查找歌曲
func (m *MusicManager) Find(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}
	for _, m := range m.musics {
		if m.Name == name {
			return &m
		}
	}
	return nil
}
// 添加一首歌曲
func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}
// 删除一首歌曲
func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(m.musics) {
		return nil
	}
	removeMusic := &m.musics[index]
	m.musics = append(m.musics[:index], m.musics[index+1:]...)

	fmt.Println("the remove music Nmae: ", removeMusic.Name)
	return removeMusic
}

func (m *MusicManager) RemoveByName(name string) *MusicEntry {
	for index, ele := range m.musics {
		if ele.Name == name {
			m.musics = append(m.musics[:index], m.musics[index+1:]...)
			return &ele
		}
	}
	fmt.Print("no found music %", name)
	return nil
}