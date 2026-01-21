package table

import "math"

type PaginationProp struct {
	CurrentPage  uint
	PageSize     uint
	TotalPages   uint
	TotalRecords uint
	SerialFrom   uint
	SerialTo     uint
}

type PageItem struct {
	Page     uint
	IsActive bool
	IsEllips bool
}

const PAGE_SIZE uint = 10
const WINDOW uint = 3

func PaginationFrom(currentPage, totalRecords uint) PaginationProp {
	pageSize := PAGE_SIZE
	serialFrom := (currentPage-1)*pageSize + 1
	serialTo := currentPage * pageSize
	if serialTo > totalRecords {
		serialTo = totalRecords
	}
	totalPages := (totalRecords + pageSize - 1) / pageSize

	return PaginationProp{
		CurrentPage:  currentPage,
		PageSize:     pageSize,
		TotalPages:   totalPages,
		TotalRecords: totalRecords,
		SerialFrom:   serialFrom,
		SerialTo:     serialTo,
	}
}

func BuildWindowedPagination(p PaginationProp) []PageItem {
	var items []PageItem
	window := WINDOW

	if p.TotalPages <= 1 {
		return []PageItem{{Page: 1, IsActive: true}}
	}

	// determine start/end pages
	start := int(p.CurrentPage) - int(window)
	end := int(p.CurrentPage) + int(window)

	if start < 1 {
		start = 1
		end = int(math.Min(float64(2*window+1), float64(p.TotalPages)))
	}

	if end > int(p.TotalPages) {
		end = int(p.TotalPages)
		start = int(math.Max(1, float64(end-int(2*window))))
	}

	// first page + ellipsis
	if start > 1 {
		items = append(items, PageItem{Page: 1})
		if start > 2 {
			items = append(items, PageItem{IsEllips: true})
		}
	}

	// main window
	for i := start; i <= end; i++ {
		items = append(items, PageItem{
			Page:     uint(i),
			IsActive: uint(i) == p.CurrentPage,
		})
	}

	// last page + ellipsis
	if end < int(p.TotalPages) {
		if end < int(p.TotalPages)-1 {
			items = append(items, PageItem{IsEllips: true})
		}
		items = append(items, PageItem{Page: p.TotalPages})
	}

	return items
}

func (p *PaginationProp) IsFirst() bool {
	return p.CurrentPage == 1
}

func (p *PaginationProp) IsLast() bool {
	return p.CurrentPage == p.TotalPages
}

func (p *PaginationProp) Iter() Iter {
	return Iter{p: p, Page: 1}
}

type Iter struct {
	p    *PaginationProp
	Page uint
}

func (it *Iter) Valid() bool {
	return it.Page <= it.p.TotalPages
}

func (it *Iter) Next() {
	it.Page += 1
}
