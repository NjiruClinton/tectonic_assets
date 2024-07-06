package timetool

import (
	"encoding/json"
	"time"
)

type Timetool struct {
	StartTime time.Time
	EndTime   time.Time
}

func NewTime() *Timetool {
	return &Timetool{}
}

func (p *Timetool) Start() {
	p.StartTime = time.Now()
}

func (p *Timetool) Stop() {
	p.EndTime = time.Now()
}

func (p *Timetool) GetData() map[string]interface{} {
	return map[string]interface{}{
		"startTime": p.StartTime,
		"endTime":   p.EndTime,
		"duration":  p.EndTime.Sub(p.StartTime).String(),
		// Add more data as necessary
	}
}

func (p *Timetool) ToJSON() (string, error) {
	data := p.GetData()
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func (p *Timetool) Reset() {
	*p = Timetool{}
}
