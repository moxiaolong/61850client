package src

type FullyEncodedData struct {
	seqOf []*PDVList
}

func (d *FullyEncodedData) getPDVList() []*PDVList {
	if d.seqOf == nil {
		d.seqOf = make([]*PDVList, 1)
	}
	return d.seqOf

}

func NewFullyEncodedData() *FullyEncodedData {
	return &FullyEncodedData{}
}
