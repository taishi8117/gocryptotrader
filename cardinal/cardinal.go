package cardinal

var (
	tickerBus chan *TickerData
)

type TickerData struct {
	ExchangeName string
	Data         interface{}
}

func SetupTickerBus() {
	if tickerBus == nil {
		tickerBus = make(chan *TickerData)
	}
}

func TickerEnabled() bool {
	return tickerBus != nil
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
