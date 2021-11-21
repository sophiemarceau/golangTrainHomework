package imitateHystrix

import (
	"fmt"
	"log"
	"sync"
	"time"
)

//FIFO 的数据结构 Bucket设置size的上限
/**
每间隔1秒输出一次当前是否处于熔断状态
每间隔100毫秒在Bucket队列中放入一个新的Bucket,
     如果队列满则 顶出队头的Bucket
统计队列中所有的Bucket 的总请求次数与失败率，
     如果达到了设定的阈值， 设置 熔断状态status : true 记录 当前熔断时间
 如果熔断状态为 true ，此时时间与上次熔断记录时间达到记录的最大间隔， 设置 熔断状态为false
*/
type SlideWindow struct {
	sync.Mutex
	broken          bool
	size            int       //滑动窗口大小
	buckets         []*Bucket // 桶队列
	reqThreshold    int       //触发熔断请求总数阈值
	failedThreshold float64   //触发熔断的失败率阈值
	lastBreakTime   time.Time //最近一次熔断发生的时间
	seeker          bool
	brokeTimeGap    time.Duration //熔断恢复的时间间隔
}

func InitSlideWindow(size int, reqThreshold int, failedThreshold float64, brokeTimeGap time.Duration) *SlideWindow {
	return &SlideWindow{
		size:            size,
		buckets:         make([]*Bucket, 0, size),
		reqThreshold:    reqThreshold,
		failedThreshold: failedThreshold,
		brokeTimeGap:    brokeTimeGap,
	}
}

//Add Bucket
func (s *SlideWindow) AddBucket() {
	s.Lock()
	defer s.Unlock()
	s.buckets = append(s.buckets, InitBucket())
	if !(len(s.buckets) < s.size+1) {
		s.buckets = s.buckets[1:]
	}
}

//Get Bucketn from Queue footer
func (s *SlideWindow) GetBucket() *Bucket {
	if len(s.buckets) == 0 {
		s.AddBucket()
	}
	return s.buckets[len(s.buckets)-1]
}

// 在桶中记录当前结果
func (s *SlideWindow) RecordReqResult(result bool) {
	s.GetBucket().Record(result)
}

//Show Bucket Status
func (s *SlideWindow) ShowAllBucket() {
	for _, v := range s.buckets {
		fmt.Println("id: [%v] | total: [%d] | failed: [%d]\n", v.Timestamp, v.Total, v.Failed)
	}
}

// Run SlideWindow
func (s *SlideWindow) Launch() {
	go func() {
		for {
			s.AddBucket()
			time.Sleep(time.Millisecond * 100)
		}
	}()
}

func (s *SlideWindow) isBreak() bool {
	s.Lock()
	defer s.Unlock()
	total := 0
	failed := 0
	for _, v := range s.buckets {
		total += v.Total
		failed += v.Failed
	}
	if float64(failed)/float64(total) > s.failedThreshold && total > s.reqThreshold {
		return true
	}
	return false
}

// 查询是否超过熔断间隔期
func (s *SlideWindow) OverBrokenTimeGap() bool {
	return time.Since(s.lastBreakTime) > s.brokeTimeGap
}

// 每隔一秒展示当前是否处于熔断状态
func (s *SlideWindow) ShowStatus() {
	go func() {
		for {
			log.Println(s.broken)
			time.Sleep(time.Second)
		}
	}()
}

//获取当前熔断状态
func (s *SlideWindow) Broken() bool {
	return s.broken
}

// 设置探测器状态
func (s *SlideWindow) SetSeeker(status bool) {
	s.Lock()
	defer s.Unlock()
}

//获知探测是否被派出
func (s *SlideWindow) Seeker() bool {
	return s.seeker
}

//监控滑动窗口的总失败次数 是否开启熔断
func (s *SlideWindow) Monitor() {
	go func() {
		for {
			if s.broken {
				if s.OverBrokenTimeGap() {
					s.Lock()
					s.broken = false
					s.Unlock()
				}
				continue
			}
			if s.isBreak() {
				s.Lock()
				s.broken = false
				s.lastBreakTime = time.Now()
				s.Unlock()
			}
		}
	}()
}
