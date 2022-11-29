package src

import "strconv"

type FcModelNodeI interface {
	ModelNodeI
	getFc() string
	getMmsVariableDef() *VariableDefsSEQUENCE
}

type FcModelNode struct {
	ModelNode
	variableDef *VariableDefsSEQUENCE
	Fc          string
}

func (n *FcModelNode) getFc() string {
	return n.Fc
}

func (n *FcModelNode) getMmsVariableDef() *VariableDefsSEQUENCE {
	if n.variableDef != nil {
		return n.variableDef
	}

	var alternateAccess *AlternateAccess = nil
	preArrayIndexItemId := ""
	preArrayIndexItemId = n.ObjectReference.get(1)
	preArrayIndexItemId += "$"
	preArrayIndexItemId += n.Fc

	arrayIndexPosition := n.ObjectReference.getArrayIndexPosition()
	if arrayIndexPosition != -1 {

		for i := 2; i < arrayIndexPosition; i++ {
			preArrayIndexItemId += "$"
			preArrayIndexItemId += n.ObjectReference.get(i)
		}

		alternateAccess = NewAlternateAccess()
		atoi, err := strconv.Atoi(n.ObjectReference.get(arrayIndexPosition))
		if err != nil {
			panic(err)
		}
		indexBerInteger := NewUnsigned32(atoi)

		if arrayIndexPosition < (n.ObjectReference.size() - 1) {
			// this reference points to a sub-node of an array element

			postArrayIndexItemId := n.ObjectReference.get(arrayIndexPosition + 1)

			for i := arrayIndexPosition + 2; i < n.ObjectReference.size(); i++ {
				postArrayIndexItemId += "$"
				postArrayIndexItemId += n.ObjectReference.get(i)
			}

			subIndexReference := NewBasicIdentifier([]byte(postArrayIndexItemId))

			subIndexReferenceSelectAccess := NewSelectAccess()

			component := NewSelectAccessComponent()
			component.basic = subIndexReference
			subIndexReferenceSelectAccess.component = component

			subIndexReferenceAlternateAccessSelection := NewAlternateAccessSelection()
			subIndexReferenceAlternateAccessSelection.selectAccess = subIndexReferenceSelectAccess

			subIndexReferenceAlternateAccessSubChoice := NewAlternateAccessCHOICE()
			subIndexReferenceAlternateAccessSubChoice.unnamed = subIndexReferenceAlternateAccessSelection

			subIndexReferenceAlternateAccess := NewAlternateAccess()

			subIndexReferenceAlternateAccess.seqOf = append(subIndexReferenceAlternateAccess.seqOf, subIndexReferenceAlternateAccessSubChoice)

			// set array index:

			indexAccessSelectionChoice := NewAccessSelection()
			indexAccessSelectionChoice.index = indexBerInteger

			indexAndLowerReferenceSelectAlternateAccess := NewSelectAlternateAccess()
			indexAndLowerReferenceSelectAlternateAccess.accessSelection = indexAccessSelectionChoice
			indexAndLowerReferenceSelectAlternateAccess.alternateAccess = subIndexReferenceAlternateAccess
			indexAndLowerReferenceAlternateAccessSelection := NewAlternateAccessSelection()
			indexAndLowerReferenceAlternateAccessSelection.selectAlternateAccess = indexAndLowerReferenceSelectAlternateAccess

			indexAndLowerReferenceAlternateAccessChoice := NewAlternateAccessCHOICE()
			indexAndLowerReferenceAlternateAccessChoice.unnamed = indexAndLowerReferenceAlternateAccessSelection

			alternateAccess.seqOf = append(alternateAccess.seqOf, indexAndLowerReferenceAlternateAccessChoice)

		} else {
			selectAccess := NewSelectAccess()
			selectAccess.index = indexBerInteger
			alternateAccessSelection := NewAlternateAccessSelection()
			alternateAccessSelection.selectAccess = selectAccess
			alternateAccessChoice := NewAlternateAccessCHOICE()
			alternateAccessChoice.unnamed = alternateAccessSelection
			alternateAccess.seqOf = append(alternateAccess.seqOf, alternateAccessChoice)
		}

	} else {
		for i := 2; i < n.ObjectReference.size(); i++ {
			preArrayIndexItemId += "$"
			preArrayIndexItemId += n.ObjectReference.get(i)
		}

	}

	domainSpecificObjectName := NewDomainSpecific()
	domainSpecificObjectName.domainID = NewIdentifier([]byte(n.ObjectReference.get(0)))
	domainSpecificObjectName.itemID = NewIdentifier([]byte(preArrayIndexItemId))

	objectName := NewObjectName()
	objectName.domainSpecific = domainSpecificObjectName

	varSpec := NewVariableSpecification()
	varSpec.name = objectName

	variableDef := NewVariableDefsSEQUENCE()
	variableDef.alternateAccess = alternateAccess
	variableDef.variableSpecification = varSpec

	return variableDef
}

func NewFcModelNode() *FcModelNode {

	return &FcModelNode{}
}
