package model

import (
	"crypto/md5"
	"encoding/hex"
	"math"
	"time"
)

type Paginate struct {
	PageCount int
	Total     int
	NextPage  int
	PrePage   int
	Todo      []*Todo
}

func SetMd5(pass string) string {
	h := md5.New()
	h.Write([]byte(pass + "Salt"))
	return hex.EncodeToString(h.Sum(nil))
}
func TimeFormat(data time.Time) string {
	return data.Format("2006-01-02 15:04:05")
}
func Paginator(post []*Todo, pagenow, pagesize int) *Paginate {
	pageCount := len(post)
	total := math.Ceil(float64(pageCount) / float64(pagesize))
	switch {
	case pagenow <= 0:
		pagenow = 0
		if pageCount > pagesize {
			post = post[:pagesize]
		} else {
			post = post[:]
		}
	case pagenow > int(total):
		pagenow = int(total)
		post = post[(pagenow-1)*pagesize:]
	default:
		if post = post[(pagenow-1)*pagesize:]; len(post) > pagesize {
			post = post[:pagesize]
		}
	}
	var prepage, nextpage int
	if prepage = pagenow - 1; pagenow <= 1 {
		prepage = -1
	}
	if nextpage = pagenow + 1; pagenow >= int(total) {
		nextpage = -1
	}
	return &Paginate{
		PageCount: pageCount,
		Total:     int(total),
		NextPage:  nextpage,
		PrePage:   prepage,
		Todo:      post,
	}
}
