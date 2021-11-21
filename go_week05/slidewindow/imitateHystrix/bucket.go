package imitateHystrix

import (
	"sync"
	"time"
)

//存储请求结果的桶
type Bucket struct {
	sync.Mutex
	Total     int //请求总数
	Failed    int //失败总数
	Timestamp time.Time
}

func InitBucket() *Bucket {
	return &Bucket{
		Timestamp: time.Now(),
	}
}

// save request result
func (b *Bucket) Record(result bool) {
	b.Lock()
	defer b.Unlock()
	if !result {
		b.Failed++
	}
	b.Total++
}
