package tes_server

import (
	"fmt"
	"github.com/boltdb/bolt"
	proto "github.com/golang/protobuf/proto"
	uuid "github.com/nu7hatch/gouuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"log"
	"strings"
	"tes/ga4gh"
	"tes/server/proto"
)

// TaskBucket documentation
// TODO: documentation
var TaskBucket = []byte("tasks")

// TaskAuthBucket documentation
// TODO: documentation
var TaskAuthBucket = []byte("tasks-auth")

// JobsQueued documentation
// TODO: documentation
var JobsQueued = []byte("jobs-queued")

// JobsActive documentation
// TODO: documentation
var JobsActive = []byte("jobs-active")

// JobsComplete documentation
// TODO: documentation
var JobsComplete = []byte("jobs-complete")

// JobsLog documentation
// TODO: documentation
var JobsLog = []byte("jobs-log")

var WorkerJobs = []byte("worker-jobs")
var JobWorker = []byte("job-worker")

// TaskBolt documentation
// TODO: documentation
type TaskBolt struct {
	db           *bolt.DB
	serverConfig ga4gh_task_ref.ServerConfig
}

// NewTaskBolt documentation
// TODO: documentation
func NewTaskBolt(path string, config ga4gh_task_ref.ServerConfig) *TaskBolt {
	db, _ := bolt.Open(path, 0600, nil)
	//Check to make sure all the required buckets have been created
	db.Update(func(tx *bolt.Tx) error {
		if tx.Bucket(TaskBucket) == nil {
			tx.CreateBucket(TaskBucket)
		}
		if tx.Bucket(TaskAuthBucket) == nil {
			tx.CreateBucket(TaskAuthBucket)
		}
		if tx.Bucket(JobsQueued) == nil {
			tx.CreateBucket(JobsQueued)
		}
		if tx.Bucket(JobsActive) == nil {
			tx.CreateBucket(JobsActive)
		}
		if tx.Bucket(JobsComplete) == nil {
			tx.CreateBucket(JobsComplete)
		}
		if tx.Bucket(JobsLog) == nil {
			tx.CreateBucket(JobsLog)
		}
		if tx.Bucket(WorkerJobs) == nil {
			tx.CreateBucket(WorkerJobs)
		}
		if tx.Bucket(JobWorker) == nil {
			tx.CreateBucket(JobWorker)
		}
		return nil
	})
	return &TaskBolt{db: db, serverConfig: config}
}

// TODO this is duplicating ListJobs? Refactor this before merging.
func (taskBolt *TaskBolt) ReadQueue(n int) []*ga4gh_task_exec.Task {
	tasks := make([]*ga4gh_task_exec.Task, 0)
	taskBolt.db.View(func(tx *bolt.Tx) error {

		taskBkt := tx.Bucket(TaskBucket)

		// Iterate over the JobsQueued bucket, reading the first `n` tasks
		c := tx.Bucket(JobsQueued).Cursor()
		for k, _ := c.First(); k != nil && len(tasks) < n; k, _ = c.Next() {
			// The values in "JobsQueued" are actually Tasks.
			task := &ga4gh_task_exec.Task{}
			proto.Unmarshal(taskBkt.Get(k), task)
			tasks = append(tasks, task)
		}
		return nil
	})
	return tasks
}

// getJWT
// This function extracts the JWT token from the rpc header and returns the string
func getJWT(ctx context.Context) string {
	jwt := ""
	v, _ := metadata.FromContext(ctx)
	auth, ok := v["authorization"]
	if !ok {
		return jwt
	}
	for _, i := range auth {
		if strings.HasPrefix(i, "JWT ") {
			jwt = strings.TrimPrefix(i, "JWT ")
		}
	}
	return jwt
}

// RunTask documentation
// TODO: documentation
func (taskBolt *TaskBolt) RunTask(ctx context.Context, task *ga4gh_task_exec.Task) (*ga4gh_task_exec.JobID, error) {
	log.Println("Receiving Task for Queue", task)

	jobID, _ := uuid.NewV4()

	if len(task.Docker) == 0 {
		return nil, fmt.Errorf("No docker commands found")
	}

	// Check inputs of the task
	for _, input := range task.GetInputs() {
		diskFound := false
		for _, res := range task.Resources.Volumes {
			if strings.HasPrefix(input.Path, res.MountPoint) {
				diskFound = true
			}
		}
		if !diskFound {
			return nil, fmt.Errorf("Required volume '%s' not found in resources", input.Path)
		}
		//Fixing blank value to File by default... Is this too much hand holding?
		if input.Class == "" {
			input.Class = "File"
		}
	}

	for _, output := range task.GetOutputs() {
		if output.Class == "" {
			output.Class = "File"
		}
	}

	jwt := getJWT(ctx)
	log.Printf("JWT: %s", jwt)

	ch := make(chan *ga4gh_task_exec.JobID, 1)
	err := taskBolt.db.Update(func(tx *bolt.Tx) error {

		taskopB := tx.Bucket(TaskBucket)
		v, _ := proto.Marshal(task)
		taskopB.Put([]byte(jobID.String()), v)

		taskopA := tx.Bucket(TaskAuthBucket)
		taskopA.Put([]byte(jobID.String()), []byte(jwt))

		queueB := tx.Bucket(JobsQueued)
		queueB.Put([]byte(jobID.String()), []byte(ga4gh_task_exec.State_Queued.String()))
		ch <- &ga4gh_task_exec.JobID{Value: jobID.String()}
		return nil
	})
	if err != nil {
		return nil, err
	}
	a := <-ch
	return a, err
}

func (taskBolt *TaskBolt) getJobState(jobID string) (ga4gh_task_exec.State, error) {

	ch := make(chan ga4gh_task_exec.State, 1)
	err := taskBolt.db.View(func(tx *bolt.Tx) error {
		bQ := tx.Bucket(JobsQueued)
		bA := tx.Bucket(JobsActive)
		bC := tx.Bucket(JobsComplete)

		if v := bQ.Get([]byte(jobID)); v != nil {
			//if its queued
			ch <- ga4gh_task_exec.State(ga4gh_task_exec.State_value[string(v)])
		} else if v := bA.Get([]byte(jobID)); v != nil {
			//if its active
			ch <- ga4gh_task_exec.State(ga4gh_task_exec.State_value[string(v)])
		} else if v := bC.Get([]byte(jobID)); v != nil {
			//if its complete
			ch <- ga4gh_task_exec.State(ga4gh_task_exec.State_value[string(v)])
		} else {
			ch <- ga4gh_task_exec.State_Unknown
		}
		return nil
	})
	a := <-ch
	return a, err
}

func (taskBolt *TaskBolt) getJob(jobID string) (*ga4gh_task_exec.Job, error) {
	ch := make(chan *ga4gh_task_exec.Job, 1)
	taskBolt.db.View(func(tx *bolt.Tx) error {

		bT := tx.Bucket(TaskBucket)
		v := bT.Get([]byte(jobID))
		task := &ga4gh_task_exec.Task{}
		proto.Unmarshal(v, task)

		job := ga4gh_task_exec.Job{}
		job.JobID = jobID
		job.Task = task
		job.State, _ = taskBolt.getJobState(jobID)

		//if there is logging info
		bL := tx.Bucket(JobsLog)
		out := make([]*ga4gh_task_exec.JobLog, len(job.Task.Docker), len(job.Task.Docker))
		for i := range job.Task.Docker {
			o := bL.Get([]byte(fmt.Sprint(jobID, i)))
			if o != nil {
				var log ga4gh_task_exec.JobLog
				proto.Unmarshal(o, &log)
				out[i] = &log
			} else {
				out[i] = &ga4gh_task_exec.JobLog{}
			}
		}
		job.Logs = out
		ch <- &job
		return nil
	})
	a := <-ch
	return a, nil
}

// GetJob documentation
// TODO: documentation
// Get info about a running task
func (taskBolt *TaskBolt) GetJob(ctx context.Context, job *ga4gh_task_exec.JobID) (*ga4gh_task_exec.Job, error) {
	log.Printf("Getting Task Info")
	ch := make(chan *ga4gh_task_exec.Task, 1)
	taskBolt.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(TaskBucket)
		v := b.Get([]byte(job.Value))
		out := ga4gh_task_exec.Task{}
		proto.Unmarshal(v, &out)
		ch <- &out
		return nil
	})
	a := <-ch
	if a == nil {
		return nil, fmt.Errorf("Job Not Found")
	}
	b, err := taskBolt.getJob(job.Value)
	return b, err
}

// ListJobs returns a list of jobIDs
func (taskBolt *TaskBolt) ListJobs(ctx context.Context, in *ga4gh_task_exec.JobListRequest) (*ga4gh_task_exec.JobListResponse, error) {
	log.Printf("Getting Task List")
	ch := make(chan ga4gh_task_exec.JobDesc, 1)
	go taskBolt.db.View(func(tx *bolt.Tx) error {
		taskopB := tx.Bucket(TaskBucket)
		c := taskopB.Cursor()
		log.Println("Scanning")
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			jobID := string(k)
			jobState, _ := taskBolt.getJobState(jobID)
			ch <- ga4gh_task_exec.JobDesc{JobID: jobID, State: jobState}
		}
		close(ch)
		return nil
	})

	jobDescArray := make([]*ga4gh_task_exec.JobDesc, 0, 10)
	for t := range ch {
		jobDescArray = append(jobDescArray, &t)
	}

	out := ga4gh_task_exec.JobListResponse{
		Jobs: jobDescArray,
	}

	log.Println("Returning", out)
	return &out, nil
}

// CancelJob documentation
// TODO: documentation
// Cancel a running task
func (taskBolt *TaskBolt) CancelJob(ctx context.Context, taskop *ga4gh_task_exec.JobID) (*ga4gh_task_exec.JobID, error) {
	taskBolt.db.Update(func(tx *bolt.Tx) error {
		bQ := tx.Bucket(JobsQueued)
		bQ.Delete([]byte(taskop.Value))
		bjw := tx.Bucket(JobWorker)

		bA := tx.Bucket(JobsActive)
		bA.Delete([]byte(taskop.Value))

		workerID := bjw.Get([]byte(taskop.Value))
		bjw.Delete([]byte(taskop.Value))

		bW := tx.Bucket(WorkerJobs)
		bW.Delete([]byte(workerID))

		bC := tx.Bucket(JobsComplete)
		bC.Put([]byte(taskop.Value), []byte(ga4gh_task_exec.State_Canceled.String()))
		return nil
	})
	return taskop, nil
}

// GetServiceInfo documentation
// TODO: documentation
func (taskBolt *TaskBolt) GetServiceInfo(ctx context.Context, info *ga4gh_task_exec.ServiceInfoRequest) (*ga4gh_task_exec.ServiceInfo, error) {
	//BUG: this isn't the best translation, probably lossy. Maybe ServiceInfo data structure schema needs to be refactored
	out := map[string]string{}
	for _, i := range taskBolt.serverConfig.Storage {
		for j, k := range i.Config {
			out[fmt.Sprintf("%s.%s", i.Protocol, j)] = k
		}
	}
	return &ga4gh_task_exec.ServiceInfo{StorageConfig: out}, nil
}
