package src

import (
	"math"
	"strconv"
	"unsafe"
)

type DataDefinitionResParser struct {
}

func NewDataDefinitionResParser() *DataDefinitionResParser {
	return &DataDefinitionResParser{}
}

func parseGetDataDefinitionResponse(confirmedServiceResponse *ConfirmedServiceResponse, lnRef *ObjectReference) *LogicalNode {
	if confirmedServiceResponse.getVariableAccessAttributes == nil {
		throw("decodeGetDataDefinitionResponse: Error decoding GetDataDefinitionResponsePdu")
	}
	varAccAttrs :=
		confirmedServiceResponse.getVariableAccessAttributes
	typeSpec := varAccAttrs.typeDescription
	if typeSpec.structure == nil {
		throw("decodeGetDataDefinitionResponse: Error decoding GetDataDefinitionResponsePdu")
	}
	structure := typeSpec.structure.components
	fcDataObjects := make([]*FcDataObject, 0)

	for _, fcComponent := range structure.seqOf {
		if fcComponent.ComponentName == nil {
			throw("Error decoding GetDataDefinitionResponsePdu")
		}
		if fcComponent.ComponentType.typeDescription == nil {
			throw(
				"Error decoding GetDataDefinitionResponsePdu")
		}
		fcString := fcComponent.ComponentName.toString()
		if fcString == ("LG") || fcString == ("GO") || fcString == ("GS") || fcString == ("MS") || fcString == ("US") {
			continue
		}
		//fc
		fc := fcComponent.ComponentName.toString()
		subStructure :=
			fcComponent.ComponentType.typeDescription.structure.components

		fcDataObjects = append(fcDataObjects, getFcDataObjectsFromSubStructure(lnRef, fc, subStructure)...)
	}

	ln := NewLogicalNode(lnRef, fcDataObjects)

	return ln
}

func getFcDataObjectsFromSubStructure(lnRef *ObjectReference, fc string, components *Components) []*FcDataObject {
	structComponents := components.seqOf
	dataObjects := make([]*FcDataObject, 0)

	for _, doComp := range structComponents {
		if doComp.ComponentName == nil {
			throw("Error decoding GetDataDefinitionResponsePdu")
		}
		if doComp.ComponentType.typeDescription == nil {
			throw("Error decoding GetDataDefinitionResponsePdu")
		}

		doRef := NewObjectReference(lnRef.toString() + "." + doComp.ComponentName.toString())
		children :=
			getDoSubModelNodesFromSubStructure(
				doRef,
				fc,
				doComp.ComponentType.typeDescription.structure.components)
		if fc == RP {
			pointer := unsafe.Pointer(NewUrcb(doRef, children))
			dataObjects = append(dataObjects, (*FcDataObject)(pointer))
		} else if fc == BR {
			pointer := unsafe.Pointer(NewBrcb(doRef, children))
			dataObjects = append(dataObjects, (*FcDataObject)(pointer))
		} else {
			pointer := unsafe.Pointer(NewFcDataObject(doRef, fc, children))
			dataObjects = append(dataObjects, (*FcDataObject)(pointer))
		}
	}

	return dataObjects
}

func getDoSubModelNodesFromSubStructure(parentRef *ObjectReference, fc string, structure *Components) []*FcModelNode {
	structComponents := structure.getSEQUENCE()
	dataObjects := make([]*FcModelNode, 0)

	for _, component := range structComponents {
		if component.ComponentName == nil {
			throw(
				"PARAMETER_VALUE_INAPPROPRIATE Error decoding GetDataDefinitionResponsePdu")
		}

		childName := component.ComponentName.toString()
		dataObjects = append(dataObjects, getModelNodesFromTypeSpecification(NewObjectReference(parentRef.toString()+"."+childName), fc, component.ComponentType))

	}

	return dataObjects
}

func getModelNodesFromTypeSpecification(ref *ObjectReference, fc string, mmsTypeSpec *TypeSpecification) *FcModelNode {
	if mmsTypeSpec.typeDescription.array != nil {

		numArrayElements :=
			mmsTypeSpec.typeDescription.array.numberOfElements.intValue()
		arrayChildren := make([]*FcModelNode, 0)

		for i := 0; i < numArrayElements; i++ {
			arrayChildren = append(arrayChildren, getModelNodesFromTypeSpecification(
				NewObjectReference(ref.toString()+"("+strconv.Itoa(i)+")"),
				fc,
				mmsTypeSpec.typeDescription.array.elementType))

		}

		array := NewArray(ref, fc, arrayChildren)
		return (*FcModelNode)(unsafe.Pointer(array))
	}

	if mmsTypeSpec.typeDescription.structure != nil {
		children :=
			getDoSubModelNodesFromSubStructure(
				ref, fc, mmsTypeSpec.typeDescription.structure.components)
		attribute := NewConstructedDataAttribute(ref, fc, children)
		return (*FcModelNode)(unsafe.Pointer(attribute))
	}

	// it is a single element

	bt := convertMmsBasicTypeSpec(ref, fc, mmsTypeSpec.typeDescription)
	if bt == nil {
		throw(
			"PARAMETER_VALUE_INAPPROPRIATE decodeGetDataDefinitionResponse: Unknown data type received " + ref.toString())
	}
	return (*FcModelNode)(unsafe.Pointer(bt))
}

func convertMmsBasicTypeSpec(ref *ObjectReference, fc string, mmsTypeSpec *TypeDescription) *BasicDataAttribute {
	if mmsTypeSpec.bool != nil {
		boolean := NewBdaBoolean(ref, fc, "", false, false)
		return (*BasicDataAttribute)(unsafe.Pointer(boolean))
	}
	if mmsTypeSpec.bitString != nil {
		bitStringMaxLength := math.Abs(float64(mmsTypeSpec.bitString.intValue()))

		if bitStringMaxLength == 13 {
			return (*BasicDataAttribute)(unsafe.Pointer(NewBdaQuality(ref, fc, "", false)))
		} else if bitStringMaxLength == 10 {
			return (*BasicDataAttribute)(unsafe.Pointer(NewBdaOptFlds(ref, fc)))
		} else if bitStringMaxLength == 6 {
			return (*BasicDataAttribute)(unsafe.Pointer(NewBdaTriggerConditions(ref, fc)))
		} else if bitStringMaxLength == 2 {
			if fc == CO {
				// if name == ctlVal
				if ref.getName()[1] == 't' {
					return (*BasicDataAttribute)(unsafe.Pointer(NewBdaTapCommand(ref, fc, "", false, false)))
				} else {
					// name == Check
					return (*BasicDataAttribute)(unsafe.Pointer(NewBdaCheck(ref)))
				}
			} else {
				return (*BasicDataAttribute)(unsafe.Pointer(NewBdaDoubleBitPos(ref, fc, "", false, false)))
			}
		}
		return nil
	} else if mmsTypeSpec.integer != nil {
		switch mmsTypeSpec.integer.intValue() {

		case 8:
			return (*BasicDataAttribute)(unsafe.Pointer(NewBdaInt8(ref, fc, "", false, false)))
		case 16:
			return (*BasicDataAttribute)(unsafe.Pointer(NewBdaInt16(ref, fc, "", false, false)))
		case 32:
			return (*BasicDataAttribute)(unsafe.Pointer(NewBdaInt32(ref, fc, "", false, false)))
		case 64:
			return (*BasicDataAttribute)(unsafe.Pointer(NewBdaInt64(ref, fc, "", false, false)))
		case 128:
			return (*BasicDataAttribute)(unsafe.Pointer(NewBdaInt128(ref, fc, "", false, false)))
		}
	} else if mmsTypeSpec.unsigned != nil {
		switch mmsTypeSpec.unsigned.intValue() {
		case 8:
			return (*BasicDataAttribute)(unsafe.Pointer(NewBdaInt8U(ref, fc, "", false, false)))
		case 16:
			return (*BasicDataAttribute)(unsafe.Pointer(NewBdaInt16U(ref, fc, "", false, false)))
		case 32:
			return (*BasicDataAttribute)(unsafe.Pointer(NewBdaInt32U(ref, fc, "", false, false)))
		}
	} else if mmsTypeSpec.floatingPoint != nil {

		floatSize := mmsTypeSpec.floatingPoint.formatWidth.intValue()
		if floatSize == 32 {
			return (*BasicDataAttribute)(unsafe.Pointer(NewBdaFloat32(ref, fc, "", false, false)))

		} else if floatSize == 64 {
			return (*BasicDataAttribute)(unsafe.Pointer(NewBdaFloat64(ref, fc, "", false, false)))

		}
		throw(
			"PARAMETER_VALUE_INAPPROPRIATE, FLOAT of size: " + strconv.Itoa(floatSize) + " is not supported.")
	} else if mmsTypeSpec.octetString != nil {

		stringSize := mmsTypeSpec.octetString.intValue()
		if stringSize > 255 || stringSize < -255 {
			throw(
				"PARAMETER_VALUE_INAPPROPRIATE OCTET_STRING of size: " + strconv.Itoa(stringSize) + " is not supported.")
		}
		return (*BasicDataAttribute)(unsafe.Pointer(NewBdaOctetString(ref, fc, "", int(math.Abs(float64(stringSize))), false, false)))

	} else if mmsTypeSpec.visibleString != nil {
		stringSize := mmsTypeSpec.visibleString.intValue()
		if stringSize > 255 || stringSize < -255 {
			throw(
				"PARAMETER_VALUE_INAPPROPRIATE VISIBLE_STRING of size: " + strconv.Itoa(stringSize) + " is not supported.")
		}
		return (*BasicDataAttribute)(unsafe.Pointer(NewBdaVisibleString(ref, fc, "", int(math.Abs(float64(stringSize))), false, false)))

	} else if mmsTypeSpec.mMSString != nil {
		stringSize := mmsTypeSpec.mMSString.intValue()
		if stringSize > 255 || stringSize < -255 {
			throw(
				"PARAMETER_VALUE_INAPPROPRIATE UNICODE_STRING of size: " + strconv.Itoa(stringSize) + " is not supported.")
		}
		return (*BasicDataAttribute)(unsafe.Pointer(NewBdaUnicodeString(ref, fc, "", int(math.Abs(float64(stringSize))), false, false)))

	} else if mmsTypeSpec.utcTime != nil {
		return (*BasicDataAttribute)(unsafe.Pointer(NewBdaTimestamp(ref, fc, "", false, false)))

	} else if mmsTypeSpec.binaryTime != nil {
		return (*BasicDataAttribute)(unsafe.Pointer(NewBdaEntryTime(ref, fc, "", false, false)))

	}
	return nil
}
