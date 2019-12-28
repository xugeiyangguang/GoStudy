package milb

import "testing"

func TestOps(t *testing.T) {
	a := NewMusicManager()
	if a == nil {
		t.Errorf("初始化错误")
	}
	music := &MusicEntry{
		Id:     "1",
		Name:   "爱在西元前",
		Artist: "jay",
		Source: "http://qbox.me",
		Type:   "Mp3",
	}
	a.Add(music)
	if a.Len() != 1 {
		t.Errorf("假如元素错误！")
	}
	b := a.Find(music.Name)
	if b == nil {
		t.Errorf("根据名字查找错误！")
	}
	if b.Name != music.Name || b.Artist != music.Artist || b.Id != music.Id || b.Source != music.Source ||
		b.Type != music.Type {
		t.Errorf("根据名字查找的结果错误！")
	}

	c, err := a.Get(0)
	if c == nil {
		t.Error("根据下标读取错误，报错：", err)
	}

	m := a.Remove(0)
	if m == nil || a.Len() != 0 {
		t.Error("根据下表删除元素错误！")
	}

}
