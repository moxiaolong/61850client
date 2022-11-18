package src

import "strings"

type ObjectReference struct {
	nodeNames          []string
	objectReference    string
	arrayIndexPosition int
}

func (r *ObjectReference) toString() string {
	return ""
}

func (r *ObjectReference) getName() string {
	return ""
}

func (r *ObjectReference) get(i int) string {

	if r.nodeNames == nil {
		r.parseForNameList()
	}
	return r.nodeNames[i]

}

func (r *ObjectReference) parseForNameList() {
	r.nodeNames = make([]string, 0)

	lastDelim := -1
	nextDelim := strings.Index(r.objectReference, "/")

	if nextDelim == -1 {
		r.nodeNames = append(r.nodeNames, r.objectReference[lastDelim+1:])
		return
	}

	r.nodeNames = append(r.nodeNames, r.objectReference[lastDelim+1:nextDelim])

	dotIndex := -1
	openingbracketIndex := -1
	closingbracketIndex := -1
	for {
		lastDelim = nextDelim
		if dotIndex == -1 {
			//dotIndex = r.objectReference.indexOf('.', lastDelim+1)
			dotIndex = strings.Index(r.objectReference[lastDelim+1:], ".") + lastDelim + 1
			if dotIndex == -1 {
				dotIndex = len(r.objectReference)
			}
		}
		if openingbracketIndex == -1 {
			openingbracketIndex = strings.Index(r.objectReference[lastDelim+1:], "(") + lastDelim + 1
			if openingbracketIndex == -1 {
				openingbracketIndex = len(r.objectReference)
			}
		}
		if closingbracketIndex == -1 {
			closingbracketIndex = strings.Index(r.objectReference[lastDelim+1:], ")") + lastDelim + 1
			if closingbracketIndex == -1 {
				closingbracketIndex = len(r.objectReference)
			}
		}

		if dotIndex == openingbracketIndex && dotIndex == closingbracketIndex {
			r.nodeNames = append(r.nodeNames, r.objectReference[lastDelim+1:])
			return
		}

		if dotIndex < openingbracketIndex && dotIndex < closingbracketIndex {
			nextDelim = dotIndex
			dotIndex = -1
		} else if openingbracketIndex < dotIndex && openingbracketIndex < closingbracketIndex {
			nextDelim = openingbracketIndex
			openingbracketIndex = -1
			r.arrayIndexPosition = len(r.nodeNames) + 1
		} else if closingbracketIndex < dotIndex && closingbracketIndex < openingbracketIndex {
			if closingbracketIndex == (len(r.objectReference) - 1) {
				r.nodeNames = append(r.nodeNames, r.objectReference[lastDelim+1:closingbracketIndex])
				return
			}
			nextDelim = closingbracketIndex + 1
			closingbracketIndex = -1
			dotIndex = -1
			r.nodeNames = append(r.nodeNames, r.objectReference[lastDelim+1:nextDelim-1])
			continue
		}
		r.nodeNames = append(r.nodeNames, r.objectReference[lastDelim+1:nextDelim])
	}
}

func NewObjectReference(string2 string) *ObjectReference {
	return &ObjectReference{arrayIndexPosition: -1}
}
