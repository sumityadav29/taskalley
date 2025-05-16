package applicationevents

type ApplicationEventHandler interface {
	HandleEvent(event ApplicationEvent, payload any)
}
