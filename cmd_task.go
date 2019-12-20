package omp

import (
	"encoding/xml"
)

// Task represent task entity in openvas,
// to use it as an argument in call to create task
// you should init its comment,name,and ID for config target, scanner
type Task struct {
	Hosts            []string
	Text             string    `xml:",chardata"`
	ID               string    `xml:"id,attr,omitempty"`
	Name             string    `xml:"name,omitempty"`
	Comment          string    `xml:"comment,omitempty"`
	CreationTime     string    `xml:"creation_time,omitempty"`
	ModificationTime string    `xml:"modification_time,omitempty"`
	Writable         string    `xml:"writable,omitempty"`
	InUse            string    `xml:"in_use,omitempty"`
	UserTags         *UserTags `xml:"user_tags,omitempty"`
	Owner            *struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,omitempty"`
	} `xml:"owner,omitempty"`
	Observers   string   `xml:"observers,omitempty"`
	Alterable   string   `xml:"alterable,omitempty"`
	Config      *Config  `xml:"config,omitempty"`
	Target      *Target  `xml:"target,omitempty"`
	Scanner     *Scanner `xml:"scanner,omitempty"`
	Status      string   `xml:"status,omitempty"`
	Progress    string   `xml:"progress,omitempty"`
	ReportCount *struct {
		Text     string `xml:",chardata"`
		Finished string `xml:"finished,omitempty"`
	} `xml:"report_count,omitempty"`
	Trend    string `xml:"trend,omitempty"`
	Schedule *struct {
		Text     string `xml:",chardata"`
		ID       string `xml:"id,attr,omitempty"`
		Name     string `xml:"name,omitempty"`
		NextTime string `xml:"next_time,omitempty"`
	} `xml:"schedule,omitempty"`
	FirstReport *struct {
		Text   string `xml:",chardata"`
		Report *struct {
			Text        string `xml:",chardata"`
			ID          string `xml:"id,attr,omitempty"`
			Timestamp   string `xml:"timestamp,omitempty"`
			ResultCount *struct {
				Text    string `xml:",chardata"`
				Debug   string `xml:"debug,omitempty"`
				Hole    string `xml:"hole,omitempty"`
				Info    string `xml:"info,omitempty"`
				Log     string `xml:"log,omitempty"`
				Warning string `xml:"warning,omitempty"`
			} `xml:"result_count,omitempty"`
			Severity string `xml:"severity,omitempty"`
		} `xml:"report,omitempty"`
	} `xml:"first_report,omitempty"`
	LastReport *struct {
		Text   string `xml:",chardata"`
		Report *struct {
			Text        string `xml:",chardata"`
			ID          string `xml:"id,attr,omitempty"`
			Timestamp   string `xml:"timestamp,omitempty"`
			ResultCount *struct {
				Text    string `xml:",chardata"`
				Debug   string `xml:"debug,omitempty"`
				Hole    string `xml:"hole,omitempty"`
				Info    string `xml:"info,omitempty"`
				Log     string `xml:"log,omitempty"`
				Warning string `xml:"warning,omitempty"`
			} `xml:"result_count,omitempty"`
			Severity string `xml:"severity,omitempty"`
		} `xml:"report,omitempty"`
	} `xml:"last_report,omitempty"`
	SecondLastReport *struct {
		Text   string `xml:",chardata"`
		Report *struct {
			Text        string `xml:",chardata"`
			ID          string `xml:"id,attr,omitempty"`
			Timestamp   string `xml:"timestamp,omitempty"`
			ResultCount *struct {
				Text    string `xml:",chardata"`
				Debug   string `xml:"debug,omitempty"`
				Hole    string `xml:"hole,omitempty"`
				Info    string `xml:"info,omitempty"`
				Log     string `xml:"log,omitempty"`
				Warning string `xml:"warning,omitempty"`
			} `xml:"result_count,omitempty"`
			Severity string `xml:"severity,omitempty"`
		} `xml:"report,omitempty"`
	} `xml:"second_last_report,omitempty"`
	Alert *struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr,omitempty"`
		Name string `xml:"name,omitempty"`
	} `xml:"alert,omitempty"`
	Reports *struct {
		Text   string `xml:",chardata"`
		Report []struct {
			Text          string `xml:",chardata"`
			ID            string `xml:"id,attr,omitempty"`
			Timestamp     string `xml:"timestamp,omitempty"`
			ScanRunStatus string `xml:"scan_run_status,omitempty"`
			ResultCount   *struct {
				Text    string `xml:",chardata"`
				Debug   string `xml:"debug,omitempty"`
				Hole    string `xml:"hole,omitempty"`
				Info    string `xml:"info,omitempty"`
				Log     string `xml:"log,omitempty"`
				Warning string `xml:"warning,omitempty"`
			} `xml:"result_count,omitempty"`
			Severity string `xml:"severity,omitempty"`
		} `xml:"report,omitempty"`
		Preferences *struct {
			Text       string `xml:",chardata"`
			Preference []struct {
				Text        string `xml:",chardata"`
				Name        string `xml:"name,omitempty"`
				ScannerName string `xml:"scanner_name,omitempty"`
				Value       string `xml:"value,omitempty"`
			} `xml:"preference,omitempty"`
		} `xml:"preferences,omitempty"`
	} `xml:"reports,omitempty"`
}

// CreateTask .
type CreateTask struct {
	XMLName xml.Name `xml:"create_task"`
	Text    string   `xml:",chardata"`
	Name    string   `xml:"name,omitempty"`
	Comment string   `xml:"comment,omitempty"`
	Config  *Config  `xml:"config,omitempty"`
	Target  *Target  `xml:"target,omitempty"`
	Scanner *Scanner `xml:"scanner,omitempty"`
}

// CreateTaskResponse .
type CreateTaskResponse struct {
	responseBase
	XMLName xml.Name `xml:"create_task_response"`
	Text    string   `xml:",chardata"`
	ID      string   `xml:"id,attr,omitempty"`
}

// DeleteTask .
type DeleteTask struct {
	XMLName xml.Name `xml:"delete_task"`
	Text    string   `xml:",chardata"`
	TaskID  string   `xml:"task_id,attr,omitempty"`
}

// DeleteTaskResponse .
type DeleteTaskResponse struct {
	XMLName xml.Name `xml:"delete_task_response"`
	Text    string   `xml:",chardata"`
	responseBase
}

// GetTasks .
type GetTasks struct {
	requestBase
	XMLName xml.Name `xml:"get_tasks"`
	Text    string   `xml:",chardata"`
	TaskID  string   `xml:"task_id,attr,omitempty"`
	Details string   `xml:"details,attr,omitempty"`
}

// MoveTask .
type MoveTask struct {
	responseBase
	XMLName xml.Name `xml:"move_task"`
	Text    string   `xml:",chardata"`
	TaskID  string   `xml:"task_id,attr,omitempty"`
	SlaveID string   `xml:"slave_id,attr,omitempty"`
}

// MoveTaskResponse .
type MoveTaskResponse struct {
	responseBase
	XMLName xml.Name `xml:"move_task_response"`
	Text    string   `xml:",chardata"`
}

// ResumeTask .
type ResumeTask struct {
	XMLName xml.Name `xml:"resume_task"`
	Text    string   `xml:",chardata"`
	TaskID  string   `xml:"task_id,attr,omitempty"`
}

// ResumeTaskResponse .
type ResumeTaskResponse struct {
	responseBase
	XMLName  xml.Name `xml:"resume_task_response"`
	Text     string   `xml:",chardata"`
	ReportID string   `xml:"report_id,omitempty"`
}

// StartTask .
type StartTask struct {
	requestBase
	XMLName xml.Name `xml:"start_task"`
	Text    string   `xml:",chardata"`
	TaskID  string   `xml:"task_id,attr,omitempty"`
}

// StartTaskResponse .
type StartTaskResponse struct {
	responseBase
	XMLName  xml.Name `xml:"start_task_response"`
	Text     string   `xml:",chardata"`
	ReportID string   `xml:"report_id,omitempty"`
}

// StopTask .
type StopTask struct {
	requestBase
	XMLName xml.Name `xml:"stop_task"`
	Text    string   `xml:",chardata"`
	TaskID  string   `xml:"task_id,attr,omitempty"`
}

// StopTaskResponse .
type StopTaskResponse struct {
	responseBase
	XMLName xml.Name `xml:"stop_task_response"`
	Text    string   `xml:",chardata"`
}

// ModifyTask .
type ModifyTask struct {
	requestBase
	XMLName xml.Name `xml:"modify_task"`
	Name    string   `xml:"name"`
	Scanner *Scanner `xml:"scanner,omitempty"`
	Text    string   `xml:",chardata"`
	TaskID  string   `xml:"task_id,attr,omitempty"`
	Comment string   `xml:"comment,omitempty"`
}

// ModifyTaskResponse .
type ModifyTaskResponse struct {
	responseBase
	XMLName xml.Name `xml:"modify_task_response"`
	Text    string   `xml:",chardata"`
}

// GetTasksResponse .
type GetTasksResponse struct {
	responseBase
	XMLName        xml.Name `xml:"get_tasks_response"`
	Text           string   `xml:",chardata"`
	ApplyOverrides string   `xml:"apply_overrides,omitempty"`
	Task           []Task   `xml:"task,omitempty"`
	Filters        *struct {
		Text string `xml:",chardata"`
		ID   string `xml:"id,attr,omitempty"`
		Term string `xml:"term,omitempty"`
	} `xml:"filters,omitempty"`
	Sort *struct {
		Text  string `xml:",chardata"`
		Field *struct {
			Text  string `xml:",chardata"`
			Order string `xml:"order,omitempty"`
		} `xml:"field,omitempty"`
	} `xml:"sort,omitempty"`
	TaskCount *struct {
		Text     string `xml:",chardata"`
		Filtered string `xml:"filtered,omitempty"`
		Page     string `xml:"page,omitempty"`
	} `xml:"task_count,omitempty"`
}

// CreateTask .
func (c *Connector) CreateTask(t *Task) (string, error) {
	req := CreateTask{
		Name:    t.Name,
		Comment: t.Comment,
		Config: &Config{
			ID: t.Config.ID,
		},
		Target: &Target{
			ID: t.Target.ID,
		},
		Scanner: &Scanner{
			ID: t.Scanner.ID,
		},
	}
	resp := &CreateTaskResponse{}
	err := c.doRequest(req, resp)
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}

// ModifyTask .
func (c *Connector) ModifyTask(t *Task) error {
	req := ModifyTask{
		TaskID:  t.ID,
		Scanner: t.Scanner,
		Name:    t.Name,
		Comment: t.Comment,
	}
	resp := &ModifyTaskResponse{}
	err := c.doRequest(req, resp)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTask .
func (c *Connector) DeleteTask(id string) error {
	req := DeleteTask{
		TaskID: id,
	}
	resp := &DeleteTaskResponse{}
	err := c.doRequest(req, resp)
	if err != nil {
		return err
	}
	return nil
}

// GetTasks .
func (c *Connector) GetTasks(id string) ([]Task, error) {
	res := []Task{}
	req := GetTasks{}
	if len(id) != 0 {
		req.TaskID = id
	}
	resp := &GetTasksResponse{}
	err := c.doRequest(req, resp)
	if err != nil {
		return res, err
	}
	for i := range resp.Task {
		res = append(res, resp.Task[i])
	}
	return res, nil
}

// StartTask .
func (c *Connector) StartTask(id string) (string, error) {
	req := StartTask{
		TaskID: id,
	}
	resp := &StartTaskResponse{}
	err := c.doRequest(req, resp)
	if err != nil {
		return "", err
	}
	return resp.ReportID, nil
}

// StopTask .
func (c *Connector) StopTask(t *Task) error {
	req := StopTask{
		TaskID: t.ID,
	}
	resp := &StopTaskResponse{}
	err := c.doRequest(req, resp)
	if err != nil {
		return err
	}
	return nil
}
