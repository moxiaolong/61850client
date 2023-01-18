package src

import (
	"errors"
	"fmt"
	"github.com/beevik/etree"
	"strconv"
	"strings"
)

type TypeDefinitions struct {
	lnodeTypes []*LnType
	doTypes    []*DoType
	daTypes    []*DaType
	enumTypes  []*EnumType
}

func (d *TypeDefinitions) putLNodeType(lnType *LnType) {
	d.lnodeTypes = append(d.lnodeTypes, lnType)
}

func (d *TypeDefinitions) putDOType(doType *DoType) {
	d.doTypes = append(d.doTypes, doType)

}

func (d *TypeDefinitions) putDAType(daType *DaType) {
	d.daTypes = append(d.daTypes, daType)

}

func (d *TypeDefinitions) putEnumType(enumType *EnumType) {
	d.enumTypes = append(d.enumTypes, enumType)

}

func (d *TypeDefinitions) getLNodeType(lnType string) *LnType {
	for _, ntype := range d.lnodeTypes {
		if ntype.Id == (lnType) {
			return ntype
		}
	}

	return nil
}

func (d *TypeDefinitions) getDOType(doType string) *DoType {
	for _, dotype := range d.doTypes {
		if dotype.Id == (doType) {
			return dotype
		}
	}

	return nil
}

func (d *TypeDefinitions) getDaType(daType string) *DaType {
	for _, datype := range d.daTypes {
		if datype.Id == daType {
			return datype
		}
	}
	return nil
}

func (d *TypeDefinitions) getEnumType(enumTypeRef string) *EnumType {
	for _, enumType := range d.enumTypes {
		if enumType.Id == enumTypeRef {
			return enumType
		}
	}
	return nil
}

func NewSclParser() *SclParser {
	return &SclParser{
		ServerModel: nil,
		doc:         nil,
		typeDefinitions: &TypeDefinitions{
			lnodeTypes: nil,
			doTypes:    nil,
			daTypes:    nil,
			enumTypes:  nil,
		},
		iedName:              "",
		useResvTmsAttributes: false,
		dataSetDefs:          nil,
		dataSetsMap:          nil,
	}
}

type SclParser struct {
	ServerModel          *ServerModel
	doc                  *etree.Document
	typeDefinitions      *TypeDefinitions
	iedName              string
	useResvTmsAttributes bool
	dataSetDefs          []*LnSubDef
	dataSetsMap          map[string]*DataSet
}

func (p *SclParser) ParseStream(content []byte) error {
	p.doc = etree.NewDocument()
	err := p.doc.ReadFromBytes(content)
	if err != nil {
		return err
	}
	rootNode := p.doc.Root()

	if "SCL" != rootNode.Tag {
		return errors.New("Root node in SCL file is not of type \"SCL\"")

	}

	err = p.readTypeDefinitions()
	if err != nil {
		return err
	}

	iedList := p.doc.Root().SelectElements("IED")
	if len(iedList) == 0 {
		return errors.New("No IED section found!")

	}

	for _, iedNode := range iedList {
		p.useResvTmsAttributes = false
		p.iedName = iedNode.SelectAttr("name").Value

		if p.iedName == "" {
			return errors.New("IED must have a name!")
		}

		iedElements := iedNode.ChildElements()
		for _, element := range iedElements {
			nodeName := element.Tag
			if "AccessPoint" == (nodeName) {
				sap, err := p.createAccessPoint(element)
				if err != nil {
					return err
				}
				serverSap := sap
				if serverSap != nil {
					p.ServerModel = serverSap.serverModel
				}
			} else if "Services" == (nodeName) {
				servicesElements := element.ChildElements()
				for _, servicesElement := range servicesElements {
					if "ReportSettings" == (servicesElement.Tag) {
						resvTmsAttribute := servicesElement.SelectAttr("resvTms")
						if resvTmsAttribute != nil {
							p.useResvTmsAttributes = strings.ToLower(resvTmsAttribute.Value) == ("true")
						}
					}
				}
			}
		}
	}

	return nil
}

func (p *SclParser) readTypeDefinitions() error {
	dttSections := p.doc.Root().SelectElements("DataTypeTemplates")
	if len(dttSections) != 1 {
		return errors.New("Only one DataTypeSection allowed")
	}
	dtt := dttSections[0]

	dataTypes := dtt.ChildElements()

	for _, element := range dataTypes {
		nodeName := element.Tag

		if nodeName == ("LNodeType") {
			p.typeDefinitions.putLNodeType(NewLnType(element))
		} else if nodeName == ("DOType") {
			p.typeDefinitions.putDOType(NewDoType(element))
		} else if nodeName == ("DAType") {
			p.typeDefinitions.putDAType(NewDaType(element))
		} else if nodeName == ("EnumType") {
			p.typeDefinitions.putEnumType(NewEnumType(element))
		}
	}
	return nil

}

func (p *SclParser) createAccessPoint(iedServer *etree.Element) (serverSap *ServerSap, err error) {

	elements := iedServer.ChildElements()
	for _, element := range elements {
		if element.Tag == ("Server") {

			s, err := p.createServerModel(element)
			if err != nil {
				return nil, err
			}
			server := s

			namedItem := iedServer.SelectAttr("name")
			if namedItem == nil {
				err = errors.New("AccessPoint has no name attribute!")
				return nil, err
			}
			// TODO save this name?
			serverSap = NewServerSap(102, 0, server)

			break
		}
	}

	return
}

func (p *SclParser) createServerModel(serverXMLNode *etree.Element) (s *ServerModel, err error) {
	elements := serverXMLNode.ChildElements()
	logicalDevices := make([]*LogicalDevice, 0)
	for _, element := range elements {
		if element.Tag == ("LDevice") {
			ld, err := p.createNewLDevice(element)
			if err != nil {
				return nil, err
			}
			logicalDevices = append(logicalDevices, ld)
		}
	}

	serverModel := NewServerModel(logicalDevices, nil)

	p.dataSetsMap = make(map[string]*DataSet)
	for _, dataSetDef := range p.dataSetDefs {

		dataSet, err := p.createDataSet(serverModel, dataSetDef.logicalNode, dataSetDef.defXmlNode)
		if err != nil {
			return nil, err
		}
		p.dataSetsMap[dataSet.DataSetReference] = dataSet
	}
	dataSets := make([]*DataSet, 0)
	for _, value := range p.dataSetsMap {
		dataSets = append(dataSets, value)
	}
	serverModel.addDataSets(dataSets)

	p.dataSetDefs = make([]*LnSubDef, 0)

	return serverModel, nil
}

func (p *SclParser) createNewLDevice(ldXmlNode *etree.Element) (ld *LogicalDevice, err error) {
	inst := ""
	ldName := ""

	attributes := ldXmlNode.Attr
	for _, node := range attributes {
		nodeName := node.Key
		if nodeName == ("inst") {
			inst = node.Value
		} else if nodeName == ("ldName") {
			ldName = node.Value
		}
	}

	if inst == "" {
		err = errors.New("Required attribute \"inst\" in logical device not found!")
		return
	}

	elements := ldXmlNode.ChildElements()
	logicalNodes := make([]ModelNodeI, 0)

	ref := ""
	if ldName != "" {
		ref = ldName
	} else {
		ref = p.iedName + inst
	}

	for _, element := range elements {
		if element.Tag == ("LN") || element.Tag == ("LN0") {
			ln, err := p.createNewLogicalNode(element, ref)
			if err != nil {
				return nil, err
			}
			logicalNodes = append(logicalNodes, ln)
		}
	}

	lDevice := NewLogicalDevice(NewObjectReference(ref), logicalNodes)

	return lDevice, nil
}

func (p *SclParser) createNewLogicalNode(lnXmlNode *etree.Element, parentRef string) (ln *LogicalNode, err error) {
	// attributes not needed: desc

	inst := ""
	lnClass := ""
	lnType := ""
	prefix := ""

	attributes := lnXmlNode.Attr

	for _, node := range attributes {
		nodeName := node.Key
		if nodeName == ("inst") {
			inst = node.Value
		} else if nodeName == ("lnType") {
			lnType = node.Value
		} else if nodeName == ("lnClass") {
			lnClass = node.Value
		} else if nodeName == ("prefix") {
			prefix = node.Value
		}
	}

	//if inst == "" {
	//	err = errors.New("Required attribute \"inst\" not found!")
	//	return nil, err
	//}
	//if lnType == "" {
	//	err = errors.New("Required attribute \"lnType\" not found!")
	//	return nil, err
	//}
	//if lnClass == "" {
	//	err = errors.New("Required attribute \"lnClass\" not found!")
	//	return nil, err
	//}

	ref := parentRef + "/" + prefix + lnClass + inst

	lnTypeDef := p.typeDefinitions.getLNodeType(lnType)

	dataObjects := make([]FcDataObjectI, 0)

	if lnTypeDef == nil {
		err = errors.New("LNType " + lnType + " not defined!")
		return nil, err
	}
	for _, dobject := range lnTypeDef.dos {
		// look for DOI node with the name of the DO
		var doiNodeFound *etree.Element = nil
		for _, childNode := range lnXmlNode.ChildElements() {

			if "DOI" == (childNode.Tag) {

				nameAttribute := childNode.SelectAttr("name")
				if nameAttribute != nil && nameAttribute.Value == (dobject.getName()) {
					doiNodeFound = childNode
				}
			}
		}

		fc, err := p.createFcDataObjects(dobject.getName(), ref, dobject.getType(), doiNodeFound)
		if err != nil {
			return nil, err
		}
		dataObjects = append(dataObjects, fc...)

	}

	// look for ReportControl
	for _, childNode := range lnXmlNode.ChildElements() {
		if "ReportControl" == (childNode.Tag) {
			rcb, err := p.createReportControlBlocks(childNode, ref)
			if err != nil {
				return nil, err
			}
			for _, item := range rcb {
				dataObjects = append(dataObjects, item)
			}
		}
	}

	lNode := NewLogicalNode(NewObjectReference(ref), dataObjects)

	// look for DataSet definitions
	for _, childNode := range lnXmlNode.ChildElements() {
		if "DataSet" == (childNode.Tag) {
			p.dataSetDefs = append(p.dataSetDefs, NewLnSubDef(childNode, lNode))
		}
	}

	return lNode, nil
}

func (p *SclParser) createFcDataObjects(name string, parentRef string, doTypeID string, doiNode *etree.Element) (fc []FcDataObjectI, err error) {

	doType := p.typeDefinitions.getDOType(doTypeID)

	if doType == nil {
		err = errors.New("DO type " + doTypeID + " not defined!")
		return
	}

	ref := parentRef + "." + name

	childNodes := make([]ModelNodeI, 0)

	for _, dattr := range doType.das {
		// look for DAI node with the name of the DA
		iNodeFound := p.findINode(doiNode, dattr.getName())

		if dattr.getCount() >= 1 {
			attributes, err := p.createArrayOfDataAttributes(ref+"."+dattr.getName(), dattr, iNodeFound)
			if err != nil {
				return nil, err
			}
			childNodes = append(childNodes, attributes)
		} else {
			m, err := p.createDataAttribute(
				ref+"."+dattr.getName(),
				dattr.getFc(),
				dattr,
				iNodeFound,
				false,
				false,
				false)
			if err != nil {
				return nil, err
			}
			childNodes = append(childNodes, m)

		}
	}

	for _, sdo := range doType.sdos {
		// parsing Arrays of SubDataObjects is ignored for now because no SCL file was found to test
		// against. The
		// only DO that contains an Array of SDOs is Harmonic Value (HMV). The Kalkitech SCL Manager
		// handles the
		// array of SDOs in HMV as an array of DAs.

		iNodeFound := p.findINode(doiNode, sdo.getName())

		dataObjects, err := p.createFcDataObjects(sdo.getName(), ref, sdo.getType(), iNodeFound)

		if err != nil {
			return nil, err
		}

		for _, object := range dataObjects {

			childNodes = append(childNodes, object)
		}

	}

	subFCDataMap := make(map[string][]ModelNodeI)

	for _, childNode := range childNodes {
		subFCDataMap[childNode.(FcModelNodeI).getFc()] = append(subFCDataMap[childNode.(FcModelNodeI).getFc()], childNode.(FcModelNodeI))
	}

	fcDataObjects := make([]FcDataObjectI, 0)
	objectReference := NewObjectReference(ref)

	for s, is := range subFCDataMap {
		fcDataObjects = append(fcDataObjects, NewFcDataObject(objectReference, s, is))

	}

	return fcDataObjects, nil
}

func (p *SclParser) createReportControlBlocks(xmlNode *etree.Element, parentRef string) (rcb []RcbI, err error) {

	fc := RP

	attribute := xmlNode.SelectAttr("buffered")
	if attribute != nil && "true" == strings.ToLower(attribute.Value) {
		fc = BR
	}

	nameAttribute := xmlNode.SelectAttr("name")
	if nameAttribute == nil {
		err = errors.New("Report Control Block has no name attribute.")
		return nil, err
	}
	maxInstances := 1
	for _, childNode := range xmlNode.ChildElements() {
		if "RptEnabled" == childNode.Tag {

			rptEnabledMaxAttr := childNode.SelectAttr("max")
			if rptEnabledMaxAttr != nil {
				atoi, err := strconv.Atoi(rptEnabledMaxAttr.Value)
				if err != nil {
					return nil, err
				}
				maxInstances = atoi
				if maxInstances < 1 || maxInstances > 99 {
					err = errors.New(
						"Report Control Block max instances should be between 1 and 99 but is: " + strconv.Itoa(maxInstances))
					return nil, err
				}
			}
		}
	}

	rcbInstances := make([]RcbI, 0)

	for z := 1; z <= maxInstances; z++ {

		var reportObjRef *ObjectReference

		if maxInstances == 1 {

			reportObjRef = NewObjectReference(parentRef + "." + nameAttribute.Value)
		} else {
			reportObjRef =
				NewObjectReference(
					parentRef + "." + nameAttribute.Value + fmt.Sprintf("%02d", z))
		}

		trigOps := NewBdaTriggerConditions(NewObjectReference(reportObjRef.toString()+".TrgOps"), fc)

		optFields := NewBdaOptFlds(NewObjectReference(reportObjRef.toString()+".OptFlds"), fc)

		for _, childNode := range xmlNode.ChildElements() {
			if childNode.Tag == "TrgOps" {

				attributes := childNode.Attr

				if attributes != nil {
					for _, node := range attributes {

						nodeName := node.Key

						if "dchg" == nodeName {
							trigOps.setDataChange(strings.ToLower(node.Value) == "true")
						} else if "qchg" == nodeName {
							trigOps.setQualityChange(strings.ToLower(node.Value) == "true")

						} else if "dupd" == nodeName {
							trigOps.setDataUpdate(strings.ToLower(node.Value) == "true")

						} else if "period" == nodeName {
							trigOps.setIntegrity(strings.ToLower(node.Value) == "true")

						} else if "gi" == nodeName {
							trigOps.setGeneralInterrogation(strings.ToLower(node.Value) == "true")
						}
					}
				}
			} else if "OptFields" == childNode.Tag {

				attributes := childNode.Attr

				if attributes != nil {
					for _, node := range attributes {

						nodeName := node.Key

						if "seqNum" == nodeName {
							optFields.setSequenceNumber(strings.ToLower(node.Value) == ("true"))
						} else if "timeStamp" == nodeName {
							optFields.setReportTimestamp(strings.ToLower(node.Value) == ("true"))

						} else if "reasonCode" == nodeName {
							optFields.setReasonForInclusion(strings.ToLower(node.Value) == ("true"))

						} else if "dataSet" == nodeName {
							optFields.setDataSetName(strings.ToLower(node.Value) == ("true"))

						} else if nodeName == ("bufOvfl") {
							optFields.setBufferOverflow(strings.ToLower(node.Value) == ("true"))

						} else if nodeName == ("entryID") {
							optFields.setEntryId(strings.ToLower(node.Value) == ("true"))
						}
						// not supported for now:
						// else if (nodeName== "configRef")) {
						// optFields.setConfigRevision(node.Value== "true"));
						// }
					}
				}
			} else if "RptEnabled" == childNode.Tag {
				rptEnabledMaxAttr := childNode.SelectAttr("max")
				if rptEnabledMaxAttr != nil {
					atoi, err := strconv.Atoi(rptEnabledMaxAttr.Value)
					if err != nil {
						return nil, err
					}
					maxInstances = atoi
					if maxInstances < 1 || maxInstances > 99 {
						err = errors.New(
							"Report Control Block max instances should be between 1 and 99 but is: " + strconv.Itoa(maxInstances))
						return nil, err
					}
				}
			}
		}

		if fc == RP {
			optFields.setEntryId(false)
			optFields.setBufferOverflow(false)
		}

		children := make([]ModelNodeI, 0)

		rptId := NewBdaVisibleString(NewObjectReference(reportObjRef.toString()+".RptID"), fc, "", 129, false, false)
		attribute = xmlNode.SelectAttr("rptID")

		if attribute != nil {
			rptId.setValue(attribute.Value)
		} else {
			rptId.setValue(reportObjRef.toString())
		}

		children = append(children, rptId)

		children = append(children, NewBdaBoolean(NewObjectReference(reportObjRef.toString()+".RptEna"), fc, "", false, false))

		if fc == RP {
			children = append(children, NewBdaBoolean(NewObjectReference(reportObjRef.toString()+".Resv"), fc, "", false, false))
		}

		datSet := NewBdaVisibleString(NewObjectReference(reportObjRef.toString()+".DatSet"), fc, "", 129, false, false)

		attribute = xmlNode.SelectAttr("datSet")
		if attribute != nil {
			nodeValue := attribute.Value
			dataSetName := parentRef + "$" + nodeValue
			datSet.setValue(dataSetName)
		}
		children = append(children, datSet)

		confRef := NewBdaInt32U(NewObjectReference(reportObjRef.toString()+".ConfRev"), fc, "", false, false)
		attribute = xmlNode.SelectAttr("confRev")
		if attribute == nil {
			err = errors.New(
				"Report Control Block does not contain mandatory attribute confRev")
			return nil, err
		}
		atoi, err := strconv.Atoi(attribute.Value)
		if err != nil {
			return nil, err
		}
		confRef.setValue(atoi)
		children = append(children, confRef)
		children = append(children, optFields)

		bufTm := NewBdaInt32U(NewObjectReference(reportObjRef.toString()+".BufTm"), fc, "", false, false)
		attribute = xmlNode.SelectAttr("bufTime")
		if attribute != nil {
			i, err := strconv.Atoi(attribute.Value)
			if err != nil {
				return nil, err
			}
			bufTm.setValue(i)
		}
		children = append(children, bufTm)
		children = append(children, NewBdaInt8U(NewObjectReference(reportObjRef.toString()+".SqNum"), fc, "", false, false))

		children = append(children, trigOps)

		intgPd := NewBdaInt32U(NewObjectReference(reportObjRef.toString()+".IntgPd"), fc, "", false, false)
		attribute = xmlNode.SelectAttr("intgPd")
		if attribute != nil {
			i, err := strconv.Atoi(attribute.Value)
			if err != nil {
				return nil, err
			}
			intgPd.setValue(i)
		}
		children = append(children, intgPd)

		children = append(children, NewBdaBoolean(NewObjectReference(reportObjRef.toString()+".GI"), fc, "", false, false))

		var rcb RcbI

		if fc == BR {

			children = append(children,
				NewBdaBoolean(
					NewObjectReference(reportObjRef.toString()+".PurgeBuf"), fc, "", false, false))

			children = append(children,
				NewBdaOctetString(
					NewObjectReference(reportObjRef.toString()+".EntryID"),
					fc,
					"",
					8,
					false,
					false))

			children = append(children,
				NewBdaEntryTime(
					NewObjectReference(reportObjRef.toString()+".TimeOfEntry"),
					fc,
					"",
					false,
					false))

			if p.useResvTmsAttributes {
				children = append(children,
					NewBdaInt16(
						NewObjectReference(reportObjRef.toString()+".ResvTms"), fc, "", false, false))
			}

			children = append(children,
				NewBdaOctetString(
					NewObjectReference(reportObjRef.toString()+".Owner"), fc, "", 64, false, false))

			rcb = NewBrcb(reportObjRef, children)

		} else {
			children = append(children,
				NewBdaOctetString(
					NewObjectReference(reportObjRef.toString()+".Owner"), fc, "", 64, false, false))

			rcb = NewUrcb(reportObjRef, children)
		}

		rcbInstances = append(rcbInstances, rcb)
	}

	return rcbInstances, nil
}

func (p *SclParser) createDataSet(model *ServerModel, lNode *LogicalNode, dsXmlNode *etree.Element) (ds *DataSet, err error) {

	nameAttribute := dsXmlNode.SelectAttr("name")
	if nameAttribute == nil {
		err = errors.New("DataSet must have a name")
		return nil, err
	}

	name := nameAttribute.Value

	dsMembers := make([]FcModelNodeI, 0)

	for _, fcdaXmlNode := range dsXmlNode.ChildElements() {

		if "FCDA" == fcdaXmlNode.Tag {

			// For the definition of FCDA see Table 22 part6 ed2

			ldInst := ""
			prefix := ""
			lnClass := ""
			lnInst := ""
			doName := ""
			daName := ""
			desc := ""
			fc := ""

			attributes := fcdaXmlNode.Attr
			for _, node := range attributes {

				nodeName := node.Key

				if nodeName == "ldInst" {
					ldInst = node.Value
				} else if nodeName == "lnInst" {
					lnInst = node.Value
				} else if nodeName == "lnClass" {
					lnClass = node.Value
				} else if nodeName == "prefix" {
					prefix = node.Value
				} else if nodeName == "doName" {
					doName = node.Value
				} else if nodeName == "daName" {
					if node.Value != "" {
						daName = "." + node.Value
					}
				} else if nodeName == "fc" {
					fc = node.Value
					if fc == "" {
						err = errors.New("FCDA contains invalid FC: " + node.Value)
						return nil, err
					}
				} else if nodeName == "desc" {
					desc = node.Value
				}
			}

			if ldInst == "" {
				err = errors.New(
					"Required attribute \"ldInst\" not found in FCDA: " + nameAttribute.Key + "!")
				return nil, err
			}

			if lnClass == "" {
				err = errors.New("Required attribute \"lnClass\" not found in FCDA!")
				return nil, err
			}
			if fc == "" {
				err = errors.New("Required attribute \"fc\" not found in FCDA!")
				return nil, err
			}
			if doName != "" {

				objectReference := p.iedName + ldInst + "/" + prefix + lnClass + lnInst + "." + doName + daName

				fcdaNode := model.findModelNode(objectReference, fc)
				fcdaNode.setDesc(desc)

				if fcdaNode == nil {
					err = errors.New(
						"Specified FCDA: " + objectReference + " in DataSet: " + nameAttribute.Key + " not found in Model.")
					return nil, err
				}
				dsMembers = append(dsMembers, fcdaNode.(FcModelNodeI))

			} else {

				objectReference := p.iedName + ldInst + "/" + prefix + lnClass + lnInst

				logicalNode := model.findModelNode(objectReference, "")
				if logicalNode == nil {
					err = errors.New("Specified FCDA: " + objectReference + " in DataSet: " + nameAttribute.Key + " not found in Model.")
					return nil, err
				}
				fcDataObjects := logicalNode.getChildren()[fc]
				for _, dataObj := range fcDataObjects.getChildren() {
					dsMembers = append(dsMembers, dataObj.(FcModelNodeI))
				}

			}
		}
	}

	dataSet := NewDataSet(lNode.getObjectReference().toString()+"."+name, dsMembers, false)
	return dataSet, nil
}

func (p *SclParser) createArrayOfDataAttributes(ref string, dataAttribute *Da, iXmlNode *etree.Element) (m *FCArray, err error) {

	fc := dataAttribute.getFc()
	size := dataAttribute.getCount()
	arrayItems := make([]ModelNodeI, 0)
	for i := 0; i < size; i++ {
		// TODO go down the iXmlNode using the ix attribute?
		attribute, err := p.createDataAttribute(
			ref+"("+strconv.Itoa(i)+")",
			fc,
			dataAttribute,
			iXmlNode,
			dataAttribute.isDchg(),
			dataAttribute.isDupd(),
			dataAttribute.isQchg())
		if err != nil {
			return nil, err
		}
		arrayItems = append(arrayItems, attribute)

	}

	return NewFCArray(NewObjectReference(ref), fc, arrayItems), nil
}

func (p *SclParser) createDataAttribute(ref string, fc string, dattr AbstractDataAttributeI, iXmlNode *etree.Element, dchg bool, dupd bool, qchg bool) (m FcModelNodeI, err error) {

	dataAttribute, ok := dattr.(*Da)
	if ok {
		dchg = dataAttribute.isDchg()
		dupd = dataAttribute.isDupd()
		qchg = dataAttribute.isQchg()
	}

	bType := dattr.getbType()

	if bType == "Struct" {
		datype := p.typeDefinitions.getDaType(dattr.getType())

		if datype == nil {
			err = errors.New("DAType " + dattr.getbType() + " not declared!")
			return nil, err
		}

		subDataAttributes := make([]ModelNodeI, 0)
		for _, bda := range datype.bdas {
			iNodeFound := p.findINode(iXmlNode, bda.getName())
			attribute, err := p.createDataAttribute(ref+"."+bda.getName(), fc, bda, iNodeFound, dchg, dupd, qchg)
			if err != nil {
				return nil, err
			}
			subDataAttributes = append(subDataAttributes, attribute)

		}

		return NewConstructedDataAttribute(NewObjectReference(ref), fc, subDataAttributes), nil
	}

	val := ""
	sAddr := ""
	if iXmlNode != nil {
		sAddrAttribute := iXmlNode.SelectAttr("sAddr")

		if sAddrAttribute != nil {
			sAddr = sAddrAttribute.Value
		}

		for _, node := range iXmlNode.ChildElements() {
			if node.Tag == "Val" {
				val = node.Text()
			}
		}

		if val == "" {
			// insert value from DA element
			val = dattr.getValue()
		}
	}

	if bType == "BOOLEAN" {
		bda := NewBdaBoolean(NewObjectReference(ref), fc, sAddr, dchg, dupd)
		if val != "" {
			if strings.ToLower(val) == ("true") || val == "1" {
				bda.setValue(true)
			} else if strings.ToLower(val) == ("false") || val == "0" {
				bda.setValue(false)
			} else {
				err = errors.New("invalid boolean configured value: " + val)
				return nil, err
			}
		}
		return bda, nil
	} else if bType == "INT8" {
		bda := NewBdaInt8(NewObjectReference(ref), fc, sAddr, dchg, dupd)
		if val != "" {

			atoi, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			bda.setValue(atoi)

		}
		return bda, nil
	} else if bType == "INT16" {
		bda := NewBdaInt16(NewObjectReference(ref), fc, sAddr, dchg, dupd)
		if val != "" {
			atoi, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			bda.setValue(atoi)

		}
		return bda, nil
	} else if bType == "INT32" {
		bda := NewBdaInt32(NewObjectReference(ref), fc, sAddr, dchg, dupd)
		if val != "" {
			atoi, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			bda.setValue(atoi)
		}
		return bda, nil
	} else if bType == "INT64" {
		bda := NewBdaInt64(NewObjectReference(ref), fc, sAddr, dchg, dupd)
		if val != "" {
			atoi, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			bda.setValue(atoi)
		}
		return bda, nil
	} else if bType == "INT128" {
		bda := NewBdaInt128(NewObjectReference(ref), fc, sAddr, dchg, dupd)
		if val != "" {
			atoi, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			bda.setValue(atoi)
		}
		return bda, nil
	} else if bType == "INT8U" {
		bda := NewBdaInt8U(NewObjectReference(ref), fc, sAddr, dchg, dupd)
		if val != "" {
			atoi, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			bda.setValue(atoi)
		}
		return bda, nil
	} else if bType == "INT16U" {
		bda := NewBdaInt16U(NewObjectReference(ref), fc, sAddr, dchg, dupd)
		if val != "" {
			atoi, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			bda.setValue(atoi)
		}
		return bda, nil
	} else if bType == "INT32U" {
		bda := NewBdaInt32U(NewObjectReference(ref), fc, sAddr, dchg, dupd)
		if val != "" {
			atoi, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			bda.setValue(atoi)
		}
		return bda, nil
	} else if bType == "FLOAT32" {
		bda := NewBdaFloat32(NewObjectReference(ref), fc, sAddr, dchg, dupd)
		if val != "" {
			bda.setValue([]byte(val))
		}
		return bda, nil
	} else if bType == "FLOAT64" {
		bda := NewBdaFloat64(NewObjectReference(ref), fc, sAddr, dchg, dupd)
		if val != "" {
			bda.setValue([]byte(val))
		}
		return bda, nil
	} else if strings.Index(bType, "VisString") == 0 {

		atoi, err := strconv.Atoi((dattr.getbType()[9:]))
		if err != nil {
			return nil, err
		}
		bda :=
			NewBdaVisibleString(
				NewObjectReference(ref),
				fc,
				sAddr,
				atoi,
				dchg,
				dupd)
		if val != "" {
			bda.setValue(val)
		}
		return bda, nil
	} else if strings.Index(bType, "Unicode") == 0 {
		atoi, err := strconv.Atoi((dattr.getbType()[7:]))
		if err != nil {
			return nil, err
		}
		bda :=
			NewBdaUnicodeString(
				NewObjectReference(ref),
				fc,
				sAddr,
				atoi,
				dchg,
				dupd)
		if val != "" {
			bda.setValue([]byte(val))
		}
		return bda, nil
	} else if strings.Index(bType, "Octet") == 0 {
		atoi, err := strconv.Atoi((dattr.getbType()[5:]))
		if err != nil {
			return nil, err
		}
		bda :=
			NewBdaOctetString(
				NewObjectReference(ref),
				fc,
				sAddr,
				atoi,
				dchg,
				dupd)
		if val != "" {
			// TODO
			// err = errors.New("parsing configured value for octet string is not supported
			// yet.");
		}
		return bda, nil
	} else if bType == "Quality" {
		return NewBdaQuality(NewObjectReference(ref), fc, sAddr, qchg), nil
	} else if bType == "Check" {
		return NewBdaCheck(NewObjectReference(ref)), nil
	} else if bType == "Dbpos" {
		return NewBdaDoubleBitPos(NewObjectReference(ref), fc, sAddr, dchg, dupd), nil
	} else if bType == "Tcmd" {
		return NewBdaTapCommand(NewObjectReference(ref), fc, sAddr, dchg, dupd), nil
	} else if bType == "OptFlds" {
		return NewBdaOptFlds(NewObjectReference(ref), fc), nil
	} else if bType == "TrgOps" {
		return NewBdaTriggerConditions(NewObjectReference(ref), fc), nil
	} else if bType == "EntryID" {
		return NewBdaOctetString(NewObjectReference(ref), fc, sAddr, 8, dchg, dupd), nil
	} else if bType == "EntryTime" {
		return NewBdaEntryTime(NewObjectReference(ref), fc, sAddr, dchg, dupd), nil
	} else if bType == "PhyComAddr" {
		// TODO not correct!
		return NewBdaOctetString(NewObjectReference(ref), fc, sAddr, 6, dchg, dupd), nil
	} else if bType == "Timestamp" {
		bda := NewBdaTimestamp(NewObjectReference(ref), fc, sAddr, dchg, dupd)
		if val != "" {
			// TODO
			err = errors.New("parsing configured value for TIMESTAMP is not supported yet.")
			return nil, err
		}
		return bda, nil
	} else if bType == "Enum" {
		atype := dattr.getType()
		if atype == "" {
			err = errors.New("The exact type of the enumeration is not set.")
			return nil, err
		}

		enumType := p.typeDefinitions.getEnumType(atype)

		if enumType == nil {
			err = errors.New("Definition of enum type: " + atype + " not found.")
			return nil, err
		}

		if enumType.max > 127 || enumType.min < -128 {

			bda := NewBdaInt16(NewObjectReference(ref), fc, sAddr, dchg, dupd)
			if val != "" {
				for _, enumVal := range enumType.getValues() {
					if val == (enumVal.getId()) {
						bda.setValue(enumVal.getOrd())
						return bda, nil
					}
				}

				err = errors.New("unknown enum value: " + val)
				return nil, err
			}
			return bda, nil
		} else {

			bda := NewBdaInt8(NewObjectReference(ref), fc, sAddr, dchg, dupd)
			for _, enumVal := range enumType.getValues() {
				if val == (enumVal.getId()) {
					bda.setValue(
						enumVal.getOrd())
					return bda, nil
				}
			}
			err = errors.New("unknown enum value: " + val)
			return bda, nil
		}
	} else if bType == "ObjRef" {

		bda :=
			NewBdaVisibleString(NewObjectReference(ref), fc, sAddr, 129, dchg, dupd)
		if val != "" {
			bda.setValue(val)
		}
		return bda, nil
	} else {
		err = errors.New("Invalid bType: " + bType)
		return nil, err
	}

}

func (p *SclParser) findINode(node *etree.Element, dattrName string) *etree.Element {
	if node == nil {
		return nil
	}
	for _, childNode := range node.ChildElements() {
		if childNode.Attr != nil {
			attr := childNode.SelectAttr("name")
			if attr != nil && attr.Value == dattrName {
				return childNode
			}
		}
	}

	return nil
}
