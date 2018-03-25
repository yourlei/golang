package library

import ("testing")

func TestOpts(t *testing.T) {
	// 01 测试新建音乐管理器功能
	mm := NewMusicManager()
	if mm == nil {
		t.Error("newMusicManager failed")
	}
	if mm.Len() != 0 {
		t.Error("NewMusicMaanger failed, not empty.")
	}
	// 02 测试Add, Find
	m0 := &MusicEntry{
		"1",
		"My Heart Will Go On", 
		"Celion Dion",
		"http://qbox.md/24501234",
		"MP3",
	}
	mm.Add(m0)
	if mm.Len() != 1 {
		t.Error("MusicManager.Add() failed.")
	}
	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("MusicManager.Find() failed.")
	}
	// 注意:换行时操作符在前一行行尾
	if m.Id != m0.Id || 
		 m.Artist != m0.Artist || 
		 m.Name != m0.Name || 
		 m.Source != m0.Source || 
		 m.Type != m0.Type {
		t.Error("MusicManager.Find() failed. Found item mismatch.")
	}
	// 03 测试Get, Remove
	m, err := mm.Get(0)
	if m == nil {
		t.Error("MusicManager.Get() failed.", err)
	}
	// m = mm.Remove(0)
	// if m == nil || mm.Len() != 0 {
	// 	t.Error("MusicManager.Remove() failed.", err)
	// }
	m = mm.RemoveByName("My Heart Will Go On")
	if m == nil || mm.Len() != 0 {
		t.Error("MusicManager.Remove() failed.", err)
	}
}