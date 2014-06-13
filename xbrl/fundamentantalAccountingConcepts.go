package xbrl

import (
	"fmt"
)

type FundamentantalAccountingConcepts struct {
	Xbrl *Xbrl
}

func (self *FundamentantalAccountingConcepts) Init(xbrl *Xbrl) {
	self.Xbrl = xbrl

	fmt.Println(" ")
	fmt.Println("FUNDAMENTAL ACCOUNTING CONCEPTS CHECK REPORT:")
	fmt.Printf("XBRL instance: %v\n", self.Xbrl.XBRLInstanceLocation)
	fmt.Printf("XBRL Cloud Viewer: https://edgardashboard.xbrlcloud.com/flex/viewer/XBRLViewer.html#instance=%v\n", self.Xbrl.XBRLInstanceLocation)

	fmt.Printf("Entity regiant name: %v\n", self.Xbrl.Fields["EntityRegistrantName"])
	fmt.Printf("CIK: %v\n", self.Xbrl.Fields["EntityCentralIndexKey"])
	fmt.Printf("Entity filer category: %v\n", self.Xbrl.Fields["EntityFilerCategory"])
	fmt.Printf("Trading symbol: %v\n", self.Xbrl.Fields["TradingSymbol"])
	fmt.Printf("Fiscal year: %v\n", self.Xbrl.Fields["DocumentFiscalYearFocus"])
	fmt.Printf("Fiscal period: %v\n", self.Xbrl.Fields["DocumentFiscalPeriodFocus"])
	fmt.Printf("Document type: %v\n", self.Xbrl.Fields["DocumentType"])

	fmt.Printf("Balance Sheet Date (document period end date): %v\n", self.Xbrl.Fields["BalanceSheetDate"])
	fmt.Printf("Income Statement Period (YTD, current period, period start date): %v\n to %v\n", self.Xbrl.Fields["IncomeStatementPeriodYTD"], self.Xbrl.Fields["BalanceSheetDate"])

	fmt.Printf("Context ID for document period focus (instants): %v\n", self.Xbrl.Fields["ContextForInstants"])
	fmt.Printf("Context ID for YTD period (durations): %v\n", self.Xbrl.Fields["ContextForDurations"])
	fmt.Println(" ")
}
