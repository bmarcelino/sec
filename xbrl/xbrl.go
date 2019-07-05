package xbrl

import (
	"bytes"
	"github.com/mjibson/goread/_third_party/code.google.com/p/go-charset/charset"
	_ "github.com/rogpeppe/go-charset/data"
	x "encoding/xml"
	"fmt"
	xml "gopkg.in/xmlpath.v1"
	"io/ioutil"
	"log"
	//"strconv"
	"time"
)

const dateLayout = "2006-01-02"

type Xbrl struct {
	XBRLInstanceLocation string
	Fields               map[string]string
	Root                 *xml.Node
}

func (self *Xbrl) Init(xbrlInstanceLocation string) {
	self.XBRLInstanceLocation = xbrlInstanceLocation
	self.Fields = make(map[string]string)

	bytesXml := loadFile(xbrlInstanceLocation)

	d := x.NewDecoder(bytes.NewReader(bytesXml))
	d.CharsetReader = charset.NewReader

	if root, err := xml.ParseDecoder(d); err != nil {
		//if root, err := xml.Parse(bytes.NewBuffer(bytesXml)); err != nil {
		log.Fatal(err)
	} else {
		self.Root = root
	}

	self.GetBaseInformation()
	self.loadYear(0)
}

func parseDate(dateString string) (date time.Time) {
	if date, err := time.Parse(dateLayout, dateString); err == nil {
		return date
	} else {
		fmt.Printf("%v is not a date: %v", dateString, err)
		panic(err)
	}
}

func (self *Xbrl) loadYear(yearMinus int) {
	var currentEnd string

	if node := self.getNode("//DocumentPeriodEndDate", nil); node != nil {
		currentEnd = node.String()

		asDate := parseDate(currentEnd)
		thisEnd := time.Date(asDate.Year()-yearMinus, asDate.Month(), asDate.Day(), 0, 0, 0, 0, time.UTC)

		//fmt.Printf("thisEnd %v\n", thisEnd.Format(dateLayout))

		self.GetCurrentPeriodAndContextInformation(thisEnd.Format(dateLayout))

		fin := new(FundamentantalAccountingConcepts)
		fin.Init(self)
		/*
			if asDate, err := time.Parse(dateLayout, currentEnd); err == nil {
				thisEnd := time.Date(asDate.Year()-yearMinus, asDate.Month(), asDate.Day(), 0, 0, 0, 0, time.UTC)

				fmt.Printf("thisEnd: %v\n", thisEnd)
				self.GetCurrentPeriodAndContextInformation(thisEnd)

				fin := new(FundamentantalAccountingConcepts)
				fin.Init(self)
			} else {
				fmt.Printf("%v is not a date: %v", currentEnd, err)
			}
		*/
	}

}

func loadFile(filePath string) (xmlContents []byte) {
	// read whole the file
	xmlContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return xmlContents
}

func (self *Xbrl) getNodeList(xPath string, root *xml.Node) (nodes *xml.Iter) {
	if root == nil {
		root = self.Root
	}
	path := xml.MustCompile(xPath)

	if exists := path.Exists(root); exists {
		return path.Iter(root)
	}

	return nil
}

func (self *Xbrl) getNode(xPath string, root *xml.Node) (node *xml.Node) {
	if root == nil {
		root = self.Root
	}
	path := xml.MustCompile(xPath)

	if exists := path.Exists(root); exists {
		iterator := path.Iter(root)

		if iterator.Next() {
			return iterator.Node()
		}
	}
	return nil
}

func (self *Xbrl) GetFactValue(seekConcept, conceptPeriodType string) (factValue string) {
	factValue = ""
	var contextReference string

	switch conceptPeriodType {
	case "Instant":
		contextReference = self.Fields["ContextForInstants"]
	case "Duration":
		contextReference = self.Fields["ContextForDurations"]
	default: //An error occured
		return "CONTEXT ERROR"
	}

	if contextReference == "" {
		return ""
	}

	node := self.getNode("//"+seekConcept+"[@contextRef='"+contextReference+"']", nil)
	if node != nil {
		factValue = node.String()
	}

	return factValue
}

func (self *Xbrl) GetBaseInformation() {
	//Registered Name
	if node := self.getNode("//EntityRegistrantName", nil); node != nil {
		self.Fields["EntityRegistrantName"] = node.String()
	} else {
		self.Fields["EntityRegistrantName"] = "Registered name not found"
	}
	//fmt.Printf("EntityRegistrantName: %v\n", self.Fields["EntityRegistrantName"])

	//Fiscal year
	if node := self.getNode("//CurrentFiscalYearEndDate", nil); node != nil {
		self.Fields["FiscalYear"] = node.String()
	} else {
		self.Fields["FiscalYear"] = "Fiscal Year not found"
	}
	//fmt.Printf("FiscalYear: %v\n", self.Fields["FiscalYear"])

	//EntityCentralIndexKey
	if node := self.getNode("//EntityCentralIndexKey", nil); node != nil {
		self.Fields["EntityCentralIndexKey"] = node.String()
	} else {
		self.Fields["EntityCentralIndexKey"] = "CIK not found"
	}
	//fmt.Printf("EntityCentralIndexKey: %v\n", self.Fields["EntityCentralIndexKey"])

	//EntityFilerCategory
	if node := self.getNode("//EntityFilerCategory", nil); node != nil {
		self.Fields["EntityFilerCategory"] = node.String()
	} else {
		self.Fields["EntityFilerCategory"] = "Filer category not found"
	}
	//fmt.Printf("EntityFilerCategory: %v\n", self.Fields["EntityFilerCategory"])

	//TradingSymbol
	if node := self.getNode("//TradingSymbol", nil); node != nil {
		self.Fields["TradingSymbol"] = node.String()
	} else {
		self.Fields["TradingSymbol"] = "Not provided"
	}
	//fmt.Printf("TradingSymbol: %v\n", self.Fields["TradingSymbol"])

	//DocumentFiscalYearFocus
	if node := self.getNode("//DocumentFiscalYearFocus", nil); node != nil {
		self.Fields["DocumentFiscalYearFocus"] = node.String()
	} else {
		self.Fields["DocumentFiscalYearFocus"] = "Fiscal year focus not found"
	}
	//fmt.Printf("DocumentFiscalYearFocus: %v\n", self.Fields["DocumentFiscalYearFocus"])

	//DocumentFiscalPeriodFocus
	if node := self.getNode("//DocumentFiscalPeriodFocus", nil); node != nil {
		self.Fields["DocumentFiscalPeriodFocus"] = node.String()
	} else {
		self.Fields["DocumentFiscalPeriodFocus"] = "Fiscal period focus not found"
	}
	//fmt.Printf("DocumentFiscalPeriodFocus: %v\n", self.Fields["DocumentFiscalPeriodFocus"])

	//DocumentType
	if node := self.getNode("//DocumentType", nil); node != nil {
		self.Fields["DocumentType"] = node.String()
	} else {
		self.Fields["DocumentType"] = "Document type not found"
	}
	//fmt.Printf("DocumentType: %v\n", self.Fields["DocumentType"])
}

func iterToArray(iter *xml.Iter, nodesArray []*xml.Node) (nodes []*xml.Node) {
	if nodesArray == nil {
		nodesArray = make([]*xml.Node, 0)
	}

	if iter == nil {
		return nodesArray
	}

	for iter.Next() {
		nodesArray = append(nodesArray, iter.Node())
	}

	return nodesArray
}

func (self *Xbrl) GetCurrentPeriodAndContextInformation(endDate string) {
	//Figures out the current period and contexts for the current period instance/duration contexts

	self.Fields["BalanceSheetDate"] = "ERROR"
	self.Fields["IncomeStatementPeriodYTD"] = "ERROR"

	self.Fields["ContextForInstants"] = "ERROR"
	self.Fields["ContextForDurations"] = "ERROR"

	//This finds the period end date for the database table, and instant date (for balance sheet):
	useContext := "ERROR"

	//Uses the concept ASSETS to find the correct instance context
	//This finds the Context ID for that end date (has correct <instant> date plus has no dimensions):

	nodeList := self.getNodeList("//Assets/@contextRef", nil)
	nodeList2 := iterToArray(nodeList, nil)

	nodeList = self.getNodeList("//AssetsCurrent/@contextRef", nil)
	nodeList2 = iterToArray(nodeList, nodeList2)

	nodeList = self.getNodeList("//LiabilitiesAndStockholdersEquity/@contextRef", nil)
	nodeList2 = iterToArray(nodeList, nodeList2)

	//Nodelist of all the facts which are us-gaap:Assets
	for i := range nodeList2 {
		node := nodeList2[i]

		contextId := node.String()
		//fmt.Printf("\n\ncontextId: %v\n", contextId)
		//contextPeriod := self.getNode("//context[@id='"+contextId+"']/period/instant", nil)
		//fmt.Printf("contextPeriod: %v\n", contextPeriod)

		//Nodelist of all the contexts of the fact us-gaap:Assets
		nodeList3 := self.getNodeList("//context[@id='"+contextId+"']", nil)

		for nodeList3.Next() {
			j := nodeList3.Node()

			if node := self.getNode("period/instant", j); node != nil {
				instant := node.String()
				//fmt.Printf("instant: %v, endDate: %v\n", instant, endDate)

				if instant == endDate {
					//fmt.Println(instant == endDate)

					if node4 := self.getNodeList("entity/segment/explicitMember", j); node4 == nil {
						useContext = contextId

						//fmt.Println(useContext)
					}
				}
			}
		}
	}

	contextForInstants := useContext
	self.Fields["ContextForInstants"] = contextForInstants
	//fmt.Printf("self.Fields[\"ContextForInstants\"]: %v\n", contextForInstants)

	//This finds the duration context
	//This may work incorrectly for fiscal year ends because the dates cross calendar years
	//Get context ID of durations and the start date for the database table
	nodeList = self.getNodeList("//CashAndCashEquivalentsPeriodIncreaseDecrease/@contextRef", nil)
	nodeList2 = iterToArray(nodeList, nil)

	nodeList = self.getNodeList("//CashPeriodIncreaseDecrease/@contextRef", nil)
	nodeList2 = iterToArray(nodeList, nodeList2)

	nodeList = self.getNodeList("//NetIncomeLoss/@contextRef", nil)
	nodeList2 = iterToArray(nodeList, nodeList2)

	nodeList = self.getNodeList("//DocumentPeriodEndDate/@contextRef", nil)
	nodeList2 = iterToArray(nodeList, nodeList2)

	//startDate := "ERROR"
	var startDate time.Time
	startDateYTD := parseDate("2099-01-01")
	useContext = "ERROR"

	for i := range nodeList2 {
		node := nodeList2[i]

		contextId := node.String()
		//fmt.Printf("\n\ncontextId: %v\n", contextId)
		//contextPeriod := self.getNode("//context[@id='"+contextId+"']/period/endDate", nil)
		//fmt.Printf("contextPeriod: %v\n", contextPeriod)

		//Nodelist of all the contexts of the fact us-gaap:Assets
		nodeList3 := self.getNodeList("//context[@id='"+contextId+"']", nil)
		for nodeList3.Next() {
			j := nodeList3.Node()

			//Nodes with the right period
			if node := self.getNode("period/endDate", j); node != nil {
				if date := node.String(); date == endDate {
					node4 := self.getNodeList("entity/segment/explicitMember", j)

					//Making sure there are no dimensions. Is this the right way to do it?
					if node4 != nil {
						//Get the year-to-date context, not the current period
						if startDateNode := self.getNode("period/startDate", j); startDateNode != nil {
							startDate = parseDate(startDateNode.String())
						}

						fmt.Printf("Context start date: %v\n", startDate.Format(dateLayout))
						fmt.Printf("YTD start date: %v\n", startDateYTD.Format(dateLayout))

						//if startDate <= startDateYTD {
						if startDate.Before(startDateYTD) {
							//Start date is for quarter
							fmt.Println("Context start date is less than current year to date, replace")
							fmt.Printf("Context start date: %v\n", startDate.Format(dateLayout))
							fmt.Printf("Current min: %v\n", startDateYTD.Format(dateLayout))

							startDateYTD = startDate

							if idAttr := self.getNode("/@id", j); idAttr != nil {
								useContext = idAttr.String()
							}

						} else {
							//Start date is for year
							fmt.Println("Context start date is greater than YTD, keep current YTD")
							fmt.Printf("Context start date: %v\n", startDate.Format(dateLayout))

							startDateYTD = startDateYTD
						}

						fmt.Println("Use context ID: " + useContext)
						fmt.Printf("Current min: %v\n", startDateYTD.Format(dateLayout))
						fmt.Println(" ")

						fmt.Println("Use context: " + useContext)
					}
				}
			}
		}

		//Balance sheet date of current period
		self.Fields["BalanceSheetDate"] = endDate

		if contextForInstants == "ERROR" {
			contextForInstants = self.LookForAlternativeInstanceContext()
			self.Fields["ContextForInstants"] = contextForInstants
		}

		//Income statement date for current fiscal year, year to date
		self.Fields["IncomeStatementPeriodYTD"] = startDateYTD.Format(dateLayout)

		contextForDurations := useContext
		self.Fields["ContextForDurations"] = contextForDurations
	}
}

func (self *Xbrl) LookForAlternativeInstanceContext() (something string) {
	//This deals with the situation where no instance context has no dimensions

	//See if there are any nodes with the document period focus date
	//fmt.Println("//context[period/instant='" + self.Fields["BalanceSheetDate"] + "']")
	nodeList_Alt := self.getNodeList("//context[period/instant='"+self.Fields["BalanceSheetDate"]+"']", nil)
	for nodeList_Alt.Next() {
		node_Alt := nodeList_Alt.Node()

		//Found possible contexts
		node_AltId := self.getNode("//@id", node_Alt).String()

		if something := self.getNode("//Assets[@contextRef='"+node_AltId+"']", nil); something != nil {
			return node_AltId
		}
	}

	return something
}
