package src

type ClientEventListener struct {
}

func (l *ClientEventListener) associationClosed(err any) {

}

func (l *ClientEventListener) newReport(report *Report) {

}

func NewClientEventListener() *ClientEventListener {
	return &ClientEventListener{}
}
