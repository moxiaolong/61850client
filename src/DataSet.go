package src

import "strings"

type DataSet struct {
	members          []*FcModelNode
	dataSetReference string
	deletable        bool
	// Map<Fc, Map<String, FcModelNode>> membersMap
	membersMap map[string]map[string]*FcModelNode
}

func (s *DataSet) getMembers() []*FcModelNode {
	return s.members
}

func NewDataSet() *DataSet {
	return &DataSet{}
}
func NewDataSetWithRef(dataSetReference string, members []*FcModelNode, deletable bool) *DataSet {
	d := &DataSet{}
	if !strings.HasPrefix(dataSetReference, "@") && strings.Index(dataSetReference, "/") == -1 {
		throw(
			"DataSet reference " + dataSetReference + " is invalid. Must either start with @ or contain a slash.")
	}
	d.members = make([]*FcModelNode, 0)
	d.dataSetReference = dataSetReference
	d.deletable = deletable
	d.membersMap = make(map[string]map[string]*FcModelNode)

	//TODO

	for _, member := range members {
		d.members = append(d.members, member)
		if d.membersMap[member.Fc] == nil {
			d.membersMap[member.Fc] = make(map[string]*FcModelNode)
		}
		d.membersMap[member.Fc][member.objectReference.toString()] = member
	}

	return d
}
