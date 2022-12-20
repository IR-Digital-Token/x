package gochannel

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/IR-Digital-Token/x/messages"
	"github.com/IR-Digital-Token/x/queue"
)

// GoChannel is the simplest Pub/Sub implementation.
// It is based on Golang's channels which are sent within the process.
//
// GoChannel has no global state,
// that means that you need to use the same instance for Publishing and Subscribing!
type GoChannel struct {
	outputChannelBuffer    int
	subscribersWg          sync.WaitGroup
	subscribers            map[string][]*subscriber
	subscribersLock        sync.RWMutex
	subscribersByTopicLock sync.Map // map of *sync.Mutex

	closed     bool
	closedLock sync.Mutex
	closing    chan struct{}
}

func NewGoChannel(outputChannelBuffer int) queue.Q {
	return &GoChannel{
		outputChannelBuffer:    outputChannelBuffer,
		subscribers:            make(map[string][]*subscriber),
		subscribersByTopicLock: sync.Map{},
		closing:                make(chan struct{}),
	}
}

func (g *GoChannel) Connect() error {
	return nil
}

func (g *GoChannel) Publish(ctx context.Context, topic string, message *messages.Message) error {
	if g.isClosed() {
		return errors.New("Pub/Sub closed")
	}
	g.subscribersLock.RLock()
	defer g.subscribersLock.RUnlock()

	subLock, _ := g.subscribersByTopicLock.LoadOrStore(topic, &sync.Mutex{})
	subLock.(*sync.Mutex).Lock()
	defer subLock.(*sync.Mutex).Unlock()

	return g.sendMessage(topic, message)
}

func (g *GoChannel) sendMessage(topic string, message *messages.Message) error {
	subscribers := g.topicSubscribers(topic)

	if len(subscribers) == 0 {
		return fmt.Errorf("not subscribers found for topic %s", topic)
	}

	go func(subscribers []*subscriber) {
		wg := &sync.WaitGroup{}

		for i := range subscribers {
			subscriber := subscribers[i]

			wg.Add(1)
			go func() {
				subscriber.sendMessageToSubscriber(message)
				wg.Done()
			}()
		}

		wg.Wait()
	}(subscribers)

	return nil
}

func (g *GoChannel) Subscribe(ctx context.Context, topic string, callback queue.CallbackFn) error {
	g.closedLock.Lock()

	if g.closed {
		return errors.New("queue closed")
	}

	g.subscribersWg.Add(1)
	g.closedLock.Unlock()
	g.subscribersLock.Lock()

	subLock, _ := g.subscribersByTopicLock.LoadOrStore(topic, &sync.Mutex{})
	subLock.(*sync.Mutex).Lock()

	s := &subscriber{
		outputChannel: make(chan *messages.Message, g.outputChannelBuffer),
	}

	go func(s *subscriber, g *GoChannel) {
		<-g.closing

		s.Close()

		g.subscribersLock.Lock()
		defer g.subscribersLock.Unlock()

		subLock, _ := g.subscribersByTopicLock.Load(topic)
		subLock.(*sync.Mutex).Lock()
		defer subLock.(*sync.Mutex).Unlock()

		g.removeSubscriber(topic, s)
		g.subscribersWg.Done()
	}(s, g)

	defer g.subscribersLock.Unlock()
	defer subLock.(*sync.Mutex).Unlock()

	g.addSubscriber(topic, s)

	go func() {
		for message := range s.outputChannel {
			err := callback(message)
			if err != nil {
				log.Println("error handling message.", err)
			}
		}
	}()
	return nil
}

func (g *GoChannel) addSubscriber(topic string, s *subscriber) {
	if _, ok := g.subscribers[topic]; !ok {
		g.subscribers[topic] = make([]*subscriber, 0)
	}
	g.subscribers[topic] = append(g.subscribers[topic], s)
}

func (g *GoChannel) removeSubscriber(topic string, toRemove *subscriber) {
	for i, sub := range g.subscribers[topic] {
		if sub == toRemove {
			g.subscribers[topic] = append(g.subscribers[topic][:i], g.subscribers[topic][i+1:]...)
			break
		}
	}
}

func (g *GoChannel) topicSubscribers(topic string) []*subscriber {
	subscribers, ok := g.subscribers[topic]
	if !ok {
		return nil
	}

	// let's do a copy to avoid race conditions and deadlocks due to lock
	subscribersCopy := make([]*subscriber, len(subscribers))
	for i, s := range subscribers {
		subscribersCopy[i] = s
	}

	return subscribersCopy
}

// Close closes the GoChannel Pub/Sub.
func (g *GoChannel) Close() error {
	g.closedLock.Lock()
	defer g.closedLock.Unlock()

	if g.closed {
		return nil
	}

	g.closed = true
	close(g.closing)
	return nil
}

func (g *GoChannel) isClosed() bool {
	g.closedLock.Lock()
	defer g.closedLock.Unlock()

	return g.closed
}

type subscriber struct {
	sending       sync.Mutex
	outputChannel chan *messages.Message

	closed bool
}

func (s *subscriber) Close() {
	if s.closed {
		return
	}

	s.sending.Lock()
	defer s.sending.Unlock()

	s.closed = true
	close(s.outputChannel)
}

func (s *subscriber) sendMessageToSubscriber(msg *messages.Message) {
	s.sending.Lock()
	defer s.sending.Unlock()

	if s.closed {
		return
	}

	s.outputChannel <- msg
}
