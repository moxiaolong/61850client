package src

type ObjectReference struct {
}

func (r *ObjectReference) toString() string {
	return ""
}

func (r *ObjectReference) getName() string {
	return ""
}

func NewObjectReference(string2 string) *ObjectReference {
	return &ObjectReference{}
}
