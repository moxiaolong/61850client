package src

import (
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
	if typeSpec.Structure == nil {
		throw("decodeGetDataDefinitionResponse: Error decoding GetDataDefinitionResponsePdu")
	}
	structure := typeSpec.Structure.Components
	fcDataObjects := make([]*FcDataObject, 0)

	for _, fcComponent := range structure.SEQUENCE {
		if fcComponent.ComponentName == nil {
			throw("Error decoding GetDataDefinitionResponsePdu")
		}
		if fcComponent.ComponentType.TypeDescription == nil {
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
			fcComponent.ComponentType.TypeDescription.Structure.Components

		fcDataObjects = append(fcDataObjects, getFcDataObjectsFromSubStructure(lnRef, fc, subStructure)...)
	}

	ln := NewLogicalNode(lnRef, fcDataObjects)

	return ln
}

func getFcDataObjectsFromSubStructure(lnRef *ObjectReference, fc string, components *Components) []*FcDataObject {
	structComponents := components.SEQUENCE
	dataObjects := make([]*FcDataObject, len(structComponents))

	for _, doComp := range structComponents {
		if doComp.ComponentName == nil {
			throw("Error decoding GetDataDefinitionResponsePdu")
		}
		if doComp.ComponentType.TypeDescription == nil {
			throw("Error decoding GetDataDefinitionResponsePdu")
		}

		doRef := NewObjectReference(lnRef.toString() + "." + doComp.ComponentName.toString())
		children :=
			getDoSubModelNodesFromSubStructure(
				doRef,
				fc,
				doComp.ComponentType.TypeDescription.Structure.Components)
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

func getDoSubModelNodesFromSubStructure(ref *ObjectReference, fc string, components *Components) []*FcModelNode {
	return nil
}
