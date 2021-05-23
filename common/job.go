package common

import (
	"io"
	"time"

	"gopkg.in/yaml.v3"
)

type Job struct {
	Name        string            `yaml:"name" json:"name"`
	Image       string            `yaml:"image" json:"image"`
	DependsOn   []string          `yaml:"depends_on" json:"depends_on"`
	Steps       []*Step           `yaml:"steps" json:"steps"`
	Timeout     int               `yaml:"timeout" json:"timeout"`
	Environment map[string]string `yaml:"env" json:"env"`
	Files       []string          `yaml:"files" json:"files"`

	fileCollector FileCollector `yaml:"-" json:"-"`
}

func (j *Job) SetFileCollector(fc FileCollector) {
	j.fileCollector = fc
}

func (j *Job) GetFileCollector() FileCollector {
	return j.fileCollector
}

func LoadJob(content []byte) (*Job, error) {
	j := &Job{}
	err := yaml.Unmarshal(content, j)
	if err != nil {
		return nil, err
	}
	return j, nil
}

type Step struct {
	Name        string            `yaml:"name" json:"name"`
	Uses        string            `yaml:"uses" json:"uses"`
	With        map[string]string `yaml:"with" json:"with"`
	Run         []string          `yaml:"run" json:"run"`
	OnError     []string          `yaml:"on_error" json:"on_error"`
	Environment map[string]string `yaml:"env" json:"env"`
}

func LoadStep(content []byte) (*Step, error) {
	s := &Step{}
	err := yaml.Unmarshal(content, s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

type JobRunStatus string

var (
	JobRunStatusPending  JobRunStatus = "pending"
	JobRunStatusRunning  JobRunStatus = "running"
	JobRunStatusFinished JobRunStatus = "finished"
	JobRunStatusCanceled JobRunStatus = "canceled"
)

type Run struct {
	ID       string    `json:"id"`
	Phistage string    `json:"phistage"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
}

type JobRun struct {
	ID        string             `json:"id"`
	Phistage  string             `json:"phistage"`
	Job       string             `json:"job"`
	Status    JobRunStatus       `json:"status"`
	Start     time.Time          `json:"start"`
	End       time.Time          `json:"end"`
	LogTracer io.ReadWriteCloser `json:"-"`
}
