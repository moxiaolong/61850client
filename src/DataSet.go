package src

import "strings"

type DataSet struct {
	Members          []FcModelNodeI
	DataSetReference string
	deletable        bool
	// Map<Fc, Map<String, FcModelNode>> MembersMap
	MembersMap    map[string]map[string]FcModelNodeI
	mmsObjectName *ObjectName
}

func (s *DataSet) getMembers() []FcModelNodeI {
	return s.Members
}

func (s *DataSet) getMmsObjectName() *ObjectName {
	if s.mmsObjectName != nil {
		return s.mmsObjectName
	}

	if s.DataSetReference[0] == '@' {
		s.mmsObjectName = NewObjectName()
		s.mmsObjectName.aaSpecific = NewIdentifier([]byte(s.DataSetReference))
		return s.mmsObjectName
	}

	slash := strings.Index(s.DataSetReference, "/")
	domainID := s.DataSetReference[0:slash]
	itemID := strings.ReplaceAll(s.DataSetReference[slash+1:], ".", "$")

	domainSpecificObjectName := NewDomainSpecific()
	domainSpecificObjectName.domainID = NewIdentifier([]byte(domainID))
	domainSpecificObjectName.itemID = NewIdentifier([]byte(itemID))

	s.mmsObjectName = NewObjectName()
	s.mmsObjectName.domainSpecific = domainSpecificObjectName

	return s.mmsObjectName
}

func NewDataSet() *DataSet {
	return &DataSet{}
}
func NewDataSetWithRef(dataSetReference string, members []FcModelNodeI, deletable bool) *DataSet {
	d := &DataSet{}
	if !strings.HasPrefix(dataSetReference, "@") && strings.Index(dataSetReference, "/") == -1 {
		throw(
			"DataSet reference " + dataSetReference + " is invalid. Must either start with @ or contain a slash.")
	}
	d.Members = make([]FcModelNodeI, 0)
	d.DataSetReference = dataSetReference
	d.deletable = deletable
	d.MembersMap = make(map[string]map[string]FcModelNodeI)

	//TODO

	for _, member := range members {
		d.Members = append(d.Members, member)
		if d.MembersMap[member.getFc()] == nil {
			d.MembersMap[member.getFc()] = make(map[string]FcModelNodeI)
		}
		d.MembersMap[member.getFc()][member.getObjectReference().toString()] = member
	}

	return d
}
