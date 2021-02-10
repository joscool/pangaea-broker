package store

type Broker struct {
	ds map[string]Set
}

func NewBroker() Broker {
	return Broker{
		ds: make(map[string]Set),
	}
}

func (b *Broker) RegisterSubscriber(topic, url string) (err error) {

	existingSet, ok := b.ds[topic]
	if ok {
		existingSet.Add(url)
	} else {
		newSet := NewSet()
		newSet.Add(url)
		b.ds[topic] = *newSet
	}
	return nil
}

func (b *Broker) GetTopicSubscribers(topic string) (data Set) {
	data = b.ds[topic]
	return
}
