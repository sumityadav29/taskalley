package applicationevents

type ApplicationEventHandler interface {
	HandleEvent(event string, payload any)
}
