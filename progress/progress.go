package progress

import (
	"math"
	"strings"
)

type Progress struct {
	TotalBytes     int64
	ProcessedBytes int64
}

func NewProgress() *Progress {
	return &Progress{}
}

func (p *Progress) GetProgressPercentage() float64 {
	if p.TotalBytes == 0 || p.ProcessedBytes == 0 {
		return 0
	}
	return float64(p.ProcessedBytes) / float64(p.TotalBytes) * 100
}

func (p *Progress) GetTotalBytes() int64 {
	return p.TotalBytes
}

func (p *Progress) GetProcessedBytes() int64 {
	return p.ProcessedBytes
}

func (p *Progress) SetTotalBytes(sizeInBytes int64) {
	p.TotalBytes = sizeInBytes
}

func (p *Progress) SetProcessedBytes(sizeInBytes int64) {
	p.ProcessedBytes = sizeInBytes
}

func (p *Progress) AddTotalBytes(sizeInBytes int64) {
	p.TotalBytes += sizeInBytes
}

func (p *Progress) AddProcessedBytes(sizeInBytes int64) {
	p.ProcessedBytes += sizeInBytes
}

func (p *Progress) GetProgressBarString(length int) string {
	color_green := "\x1b[38;2;0;102;204m"
	color_gray := "\x1b[38;2;64;64;64m"
	color_reset := "\x1b[0m"

	var proportion float64
	if p.TotalBytes <= 0 {
		proportion = 0
	} else {
		proportion = float64(p.ProcessedBytes) / float64(p.TotalBytes)
	}
	if math.IsNaN(proportion) || math.IsInf(proportion, 0) {
		proportion = 0
	}
	if proportion < 0 {
		proportion = 0
	}
	if proportion > 1 {
		proportion = 1
	}

	filled := int(math.Floor(proportion * float64(length)))
	if filled < 0 {
		filled = 0
	}
	if filled > length {
		filled = length
	}
	remaining := length - filled

	progressBar := color_green + strings.Repeat("━", filled) + color_reset
	if remaining > 0 {
		progressBar += color_gray + "╺" + strings.Repeat("━", remaining) + color_reset
	}

	return progressBar
}
