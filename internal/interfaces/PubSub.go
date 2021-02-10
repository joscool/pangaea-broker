package interfaces

// PubSub ...
type PubSub interface {
	Subscribe(topic, url string) error
	Publish(topic, message string) error
}
