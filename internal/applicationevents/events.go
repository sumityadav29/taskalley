package applicationevents

type ApplicationEvent string

const (
	ProjectCreated ApplicationEvent = "project.created"
	TaskCreated    ApplicationEvent = "task.created"
	TaskUpdated    ApplicationEvent = "task.updated"
	TaskDeleted    ApplicationEvent = "task.deleted"
)

var ApplicationEvents = []ApplicationEvent{
	ProjectCreated,
	TaskCreated,
	TaskUpdated,
	TaskDeleted,
}
