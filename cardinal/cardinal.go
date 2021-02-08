package cardinal

var (
	tickerBus  chan *TickerData
	watcherBus chan *WatcherData
)

type TickerData struct {
	ExchangeName string
	Data         interface{}
}

type WatcherData struct {
	ExchangeName string
	Data         interface{}
}

func SetupTickerBus() {
	if tickerBus == nil {
		tickerBus = make(chan *TickerData)
	}
}

func SetupWatcherBus() {
	if watcherBus == nil {
		watcherBus = make(chan *WatcherData)
	}
}

func TickerEnabled() bool {
	return tickerBus != nil
}

func WatcherEnabled() bool {
	return watcherBus != nil
}

func PushTrade(exch string, d interface{}) {
	tickerBus <- &TickerData{
		ExchangeName: exch,
		Data:         d,
	}
}

func ReadTicker() <-chan *TickerData {
	return tickerBus
}

func ReadWatcher() <-chan *WatcherData {
	return watcherBus
}
