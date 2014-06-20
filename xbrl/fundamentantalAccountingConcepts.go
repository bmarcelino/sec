package xbrl

import (
	"fmt"
	"strconv"
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

	self.LoadFinFacts()
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 32)
}

func (self *FundamentantalAccountingConcepts) LoadFinFacts() {
	//Assets
	if fact := self.Xbrl.GetFactValue("Assets", "Instant"); fact != "" {
		self.Xbrl.Fields["Assets"] = fact
	}

	//Current Assets
	self.Xbrl.Fields["CurrentAssets"] = self.Xbrl.GetFactValue("AssetsCurrent", "Instant")
	if self.Xbrl.Fields["CurrentAssets"] == "" {
		self.Xbrl.Fields["CurrentAssets"] = "0"
	}

	//Noncurrent Assets
	self.Xbrl.Fields["NoncurrentAssets"] = self.Xbrl.GetFactValue("AssetsNoncurrent", "Instant")
	if self.Xbrl.Fields["NoncurrentAssets"] == "" {
		if self.Xbrl.Fields["Assets"] != "" && self.Xbrl.Fields["CurrentAssets"] != "" {
			fAssets, _ := strconv.ParseFloat(self.Xbrl.Fields["Assets"], 32)
			fCurrentAssets, _ := strconv.ParseFloat(self.Xbrl.Fields["CurrentAssets"], 32)
			fNoncurrentAssets := fAssets - fCurrentAssets

			self.Xbrl.Fields["NoncurrentAssets"] = FloatToString(fNoncurrentAssets)
		} else {
			self.Xbrl.Fields["NoncurrentAssets"] = "0"
		}
	}
	//LiabilitiesAndEquity
	self.Xbrl.Fields["LiabilitiesAndEquity"] = self.Xbrl.GetFactValue("LiabilitiesAndStockholdersEquity", "Instant")
	if self.Xbrl.Fields["LiabilitiesAndEquity"] == "" {
		self.Xbrl.Fields["LiabilitiesAndEquity"] = self.Xbrl.GetFactValue("LiabilitiesAndPartnersCapital", "Instant")
		if self.Xbrl.Fields["LiabilitiesAndEquity"] == "" {
			self.Xbrl.Fields["LiabilitiesAndEquity"] = "0"
		}
	}
	//Liabilities
	self.Xbrl.Fields["Liabilities"] = self.Xbrl.GetFactValue("Liabilities", "Instant")
	if self.Xbrl.Fields["Liabilities"] == "" {
		self.Xbrl.Fields["Liabilities"] = "0"
	}
	//CurrentLiabilities
	self.Xbrl.Fields["CurrentLiabilities"] = self.Xbrl.GetFactValue("LiabilitiesCurrent", "Instant")
	if self.Xbrl.Fields["CurrentLiabilities"] == "" {
		self.Xbrl.Fields["CurrentLiabilities"] = "0"
	}
	//Noncurrent Liabilities
	self.Xbrl.Fields["NoncurrentLiabilities"] = self.Xbrl.GetFactValue("LiabilitiesNoncurrent", "Instant")
	if self.Xbrl.Fields["NoncurrentLiabilities"] == "" {
		if self.Xbrl.Fields["Liabilities"] != "" && self.Xbrl.Fields["CurrentLiabilities"] != "" {
			fLiabilities, _ := strconv.ParseFloat(self.Xbrl.Fields["Liabilities"], 32)
			fCurrentLiabilities, _ := strconv.ParseFloat(self.Xbrl.Fields["CurrentLiabilities"], 32)
			fNoncurrentLiabilities := fLiabilities - fCurrentLiabilities

			self.Xbrl.Fields["NoncurrentLiabilities"] = FloatToString(fNoncurrentLiabilities)

		} else {
			self.Xbrl.Fields["NoncurrentLiabilities"] = "0"
		}
	}
	//CommitmentsAndContingencies
	self.Xbrl.Fields["CommitmentsAndContingencies"] = self.Xbrl.GetFactValue("CommitmentsAndContingencies", "Instant")
	if self.Xbrl.Fields["CommitmentsAndContingencies"] == "" {
		self.Xbrl.Fields["CommitmentsAndContingencies"] = "0"
	}

	//TemporaryEquity
	self.Xbrl.Fields["TemporaryEquity"] = self.Xbrl.GetFactValue("TemporaryEquityRedemptionValue", "Instant")
	if self.Xbrl.Fields["TemporaryEquity"] == "" {
		self.Xbrl.Fields["TemporaryEquity"] = self.Xbrl.GetFactValue("RedeemablePreferredStockCarryingAmount", "Instant")
		if self.Xbrl.Fields["TemporaryEquity"] == "" {
			self.Xbrl.Fields["TemporaryEquity"] = self.Xbrl.GetFactValue("TemporaryEquityCarryingAmount", "Instant")
			if self.Xbrl.Fields["TemporaryEquity"] == "" {
				self.Xbrl.Fields["TemporaryEquity"] = self.Xbrl.GetFactValue("TemporaryEquityValueExcludingAdditionalPaidInCapital", "Instant")
				if self.Xbrl.Fields["TemporaryEquity"] == "" {
					self.Xbrl.Fields["TemporaryEquity"] = self.Xbrl.GetFactValue("TemporaryEquityCarryingAmountAttributableToParent", "Instant")
					if self.Xbrl.Fields["TemporaryEquity"] == "" {
						self.Xbrl.Fields["TemporaryEquity"] = self.Xbrl.GetFactValue("RedeemableNoncontrollingInterestEquityFairValue", "Instant")
						if self.Xbrl.Fields["TemporaryEquity"] == "" {
							self.Xbrl.Fields["TemporaryEquity"] = "0"
						}
					}
				}
			}
		}
	}

	//RedeemableNoncontrollingInterest (added to temporary equity)
	RedeemableNoncontrollingInterest := ""

	RedeemableNoncontrollingInterest = self.Xbrl.GetFactValue("RedeemableNoncontrollingInterestEquityCarryingAmount", "Instant")
	if RedeemableNoncontrollingInterest == "" {
		RedeemableNoncontrollingInterest = self.Xbrl.GetFactValue("RedeemableNoncontrollingInterestEquityCommonCarryingAmount", "Instant")
		if RedeemableNoncontrollingInterest == "" {
			RedeemableNoncontrollingInterest = "0"
		}
	}

	//This adds redeemable noncontrolling interest and temporary equity which are rare, but can be reported seperately
	if self.Xbrl.Fields["TemporaryEquity"] != "" {
		fTemporaryEquity, _ := strconv.ParseFloat(self.Xbrl.Fields["TemporaryEquity"], 32)
		fRedeemableNoncontrollingInterest, _ := strconv.ParseFloat(RedeemableNoncontrollingInterest, 32)

		self.Xbrl.Fields["TemporaryEquity"] = FloatToString(fTemporaryEquity + fRedeemableNoncontrollingInterest)
	}

	//Equity
	self.Xbrl.Fields["Equity"] = self.Xbrl.GetFactValue("StockholdersEquityIncludingPortionAttributableToNoncontrollingInterest", "Instant")
	if self.Xbrl.Fields["Equity"] == "" {
		self.Xbrl.Fields["Equity"] = self.Xbrl.GetFactValue("StockholdersEquity", "Instant")
		if self.Xbrl.Fields["Equity"] == "" {
			self.Xbrl.Fields["Equity"] = self.Xbrl.GetFactValue("PartnersCapitalIncludingPortionAttributableToNoncontrollingInterest", "Instant")
			if self.Xbrl.Fields["Equity"] == "" {
				self.Xbrl.Fields["Equity"] = self.Xbrl.GetFactValue("PartnersCapital", "Instant")
				if self.Xbrl.Fields["Equity"] == "" {
					self.Xbrl.Fields["Equity"] = self.Xbrl.GetFactValue("CommonStockholdersEquity", "Instant")
					if self.Xbrl.Fields["Equity"] == "" {
						self.Xbrl.Fields["Equity"] = self.Xbrl.GetFactValue("MemberEquity", "Instant")
						if self.Xbrl.Fields["Equity"] == "" {
							self.Xbrl.Fields["Equity"] = self.Xbrl.GetFactValue("AssetsNet", "Instant")
							if self.Xbrl.Fields["Equity"] == "" {
								self.Xbrl.Fields["Equity"] = "0"
							}
						}
					}
				}
			}
		}
	}

	//EquityAttributableToNoncontrollingInterest
	self.Xbrl.Fields["EquityAttributableToNoncontrollingInterest"] = self.Xbrl.GetFactValue("MinorityInterest", "Instant")
	if self.Xbrl.Fields["EquityAttributableToNoncontrollingInterest"] == "" {
		self.Xbrl.Fields["EquityAttributableToNoncontrollingInterest"] = self.Xbrl.GetFactValue("PartnersCapitalAttributableToNoncontrollingInterest", "Instant")
		if self.Xbrl.Fields["EquityAttributableToNoncontrollingInterest"] == "" {
			self.Xbrl.Fields["EquityAttributableToNoncontrollingInterest"] = "0"
		}
	}

	//EquityAttributableToParent
	self.Xbrl.Fields["EquityAttributableToParent"] = self.Xbrl.GetFactValue("StockholdersEquity", "Instant")
	if self.Xbrl.Fields["EquityAttributableToParent"] == "" {
		self.Xbrl.Fields["EquityAttributableToParent"] = self.Xbrl.GetFactValue("LiabilitiesAndPartnersCapital", "Instant")
		if self.Xbrl.Fields["EquityAttributableToParent"] == "" {
			self.Xbrl.Fields["EquityAttributableToParent"] = "0"
		}
	}

	//BS Adjustments
	//if total assets is missing, try using current assets
	if self.Xbrl.Fields["Assets"] == "0" && self.Xbrl.Fields["Assets"] == self.Xbrl.Fields["LiabilitiesAndEquity"] && self.Xbrl.Fields["CurrentAssets"] == self.Xbrl.Fields["LiabilitiesAndEquity"] {
		self.Xbrl.Fields["Assets"] = self.Xbrl.Fields["CurrentAssets"]
	}
	//Added to fix Assets
	if self.Xbrl.Fields["Assets"] == "0" && self.Xbrl.Fields["LiabilitiesAndEquity"] != "0" && (self.Xbrl.Fields["CurrentAssets"] == self.Xbrl.Fields["LiabilitiesAndEquity"]) {
		self.Xbrl.Fields["Assets"] = self.Xbrl.Fields["CurrentAssets"]
	}
	//Added to fix Assets even more
	if self.Xbrl.Fields["Assets"] == "0" && self.Xbrl.Fields["NoncurrentAssets"] == "0" && self.Xbrl.Fields["LiabilitiesAndEquity"] != "0" && (self.Xbrl.Fields["LiabilitiesAndEquity"] == self.Xbrl.Fields["Liabilities"]+self.Xbrl.Fields["Equity"]) {
		self.Xbrl.Fields["Assets"] = self.Xbrl.Fields["CurrentAssets"]
	}
	if self.Xbrl.Fields["Assets"] != "0" && self.Xbrl.Fields["CurrentAssets"] != "0" {
		fAssets, _ := strconv.ParseFloat(self.Xbrl.Fields["Assets"], 32)
		fCurrentAssets, _ := strconv.ParseFloat(self.Xbrl.Fields["CurrentAssets"], 32)

		self.Xbrl.Fields["NoncurrentAssets"] = FloatToString(fAssets - fCurrentAssets) //self.Xbrl.Fields["Assets"] - self.Xbrl.Fields["CurrentAssets"]
	}
	if self.Xbrl.Fields["LiabilitiesAndEquity"] == "0" && self.Xbrl.Fields["Assets"] != "0" {
		self.Xbrl.Fields["LiabilitiesAndEquity"] = self.Xbrl.Fields["Assets"]
	}
	//Impute: Equity based no parent && noncontrolling interest being present
	if self.Xbrl.Fields["EquityAttributableToNoncontrollingInterest"] != "0" && self.Xbrl.Fields["EquityAttributableToParent"] != "0" {
		fEquityAttributableToParent, _ := strconv.ParseFloat(self.Xbrl.Fields["EquityAttributableToParent"], 32)
		fEquityAttributableToNoncontrollingInterest, _ := strconv.ParseFloat(self.Xbrl.Fields["EquityAttributableToNoncontrollingInterest"], 32)

		self.Xbrl.Fields["Equity"] = FloatToString(fEquityAttributableToParent + fEquityAttributableToNoncontrollingInterest) //self.Xbrl.Fields["EquityAttributableToParent"] + self.Xbrl.Fields["EquityAttributableToNoncontrollingInterest"]
	}
	if self.Xbrl.Fields["Equity"] == "0" && self.Xbrl.Fields["EquityAttributableToNoncontrollingInterest"] == "0" && self.Xbrl.Fields["EquityAttributableToParent"] != "0" {
		self.Xbrl.Fields["Equity"] = self.Xbrl.Fields["EquityAttributableToParent"]
	}
	if self.Xbrl.Fields["Equity"] == "0" {
		fEquityAttributableToParent, _ := strconv.ParseFloat(self.Xbrl.Fields["EquityAttributableToParent"], 32)
		fEquityAttributableToNoncontrollingInterest, _ := strconv.ParseFloat(self.Xbrl.Fields["EquityAttributableToNoncontrollingInterest"], 32)

		self.Xbrl.Fields["Equity"] = FloatToString(fEquityAttributableToParent + fEquityAttributableToNoncontrollingInterest) //self.Xbrl.Fields["EquityAttributableToParent"] + self.Xbrl.Fields["EquityAttributableToNoncontrollingInterest"]
	}
	//Added{ Impute Equity attributable to parent based on existence of equity && noncontrolling interest.
	if self.Xbrl.Fields["Equity"] != "0" && self.Xbrl.Fields["EquityAttributableToNoncontrollingInterest"] != "0" && self.Xbrl.Fields["EquityAttributableToParent"] == "0" {
		fEquity, _ := strconv.ParseFloat(self.Xbrl.Fields["Equity"], 32)
		fEquityAttributableToNoncontrollingInterest, _ := strconv.ParseFloat(self.Xbrl.Fields["EquityAttributableToNoncontrollingInterest"], 32)

		self.Xbrl.Fields["EquityAttributableToParent"] = FloatToString(fEquity - fEquityAttributableToNoncontrollingInterest) //self.Xbrl.Fields["Equity"] - self.Xbrl.Fields["EquityAttributableToNoncontrollingInterest"]
	}
	//Added{ Impute Equity attributable to parent based on existence of equity && noncontrolling interest.
	if self.Xbrl.Fields["Equity"] != "0" && self.Xbrl.Fields["EquityAttributableToNoncontrollingInterest"] == "0" && self.Xbrl.Fields["EquityAttributableToParent"] == "0" {
		self.Xbrl.Fields["EquityAttributableToParent"] = self.Xbrl.Fields["Equity"]
	}
	//if total liabilities is missing, figure it out based on liabilities && equity
	if self.Xbrl.Fields["Liabilities"] == "0" && self.Xbrl.Fields["Equity"] != "0" {
		fLiabilitiesAndEquity, _ := strconv.ParseFloat(self.Xbrl.Fields["LiabilitiesAndEquity"], 32)
		fCommitmentsAndContingencies, _ := strconv.ParseFloat(self.Xbrl.Fields["CommitmentsAndContingencies"], 32)
		fTemporaryEquity, _ := strconv.ParseFloat(self.Xbrl.Fields["TemporaryEquity"], 32)
		fEquity, _ := strconv.ParseFloat(self.Xbrl.Fields["Equity"], 32)

		self.Xbrl.Fields["Liabilities"] = FloatToString(fLiabilitiesAndEquity - (fCommitmentsAndContingencies + fTemporaryEquity + fEquity)) //self.Xbrl.Fields["LiabilitiesAndEquity"] - (self.Xbrl.Fields["CommitmentsAndContingencies"] + self.Xbrl.Fields["TemporaryEquity"] + self.Xbrl.Fields["Equity"])
	}
	//This seems incorrect because liabilities might not be reported
	if self.Xbrl.Fields["Liabilities"] != "0" && self.Xbrl.Fields["CurrentLiabilities"] != "0" {
		fLiabilities, _ := strconv.ParseFloat(self.Xbrl.Fields["Liabilities"], 32)
		fCurrentLiabilities, _ := strconv.ParseFloat(self.Xbrl.Fields["CurrentLiabilities"], 32)

		self.Xbrl.Fields["NoncurrentLiabilities"] = FloatToString(fLiabilities - fCurrentLiabilities) //self.Xbrl.Fields["Liabilities"] - self.Xbrl.Fields["CurrentLiabilities"]
	}
	//Added to fix liabilities based on current liabilities
	if self.Xbrl.Fields["Liabilities"] == "0" && self.Xbrl.Fields["CurrentLiabilities"] != "0" && self.Xbrl.Fields["NoncurrentLiabilities"] == "0" {
		self.Xbrl.Fields["Liabilities"] = self.Xbrl.Fields["CurrentLiabilities"]
	}

	fEquity, _ := strconv.ParseFloat(self.Xbrl.Fields["Equity"], 32)
	fEquityAttributableToParent, _ := strconv.ParseFloat(self.Xbrl.Fields["EquityAttributableToParent"], 32)
	fEquityAttributableToNoncontrollingInterest, _ := strconv.ParseFloat(self.Xbrl.Fields["EquityAttributableToNoncontrollingInterest"], 32)
	fAssets, _ := strconv.ParseFloat(self.Xbrl.Fields["Assets"], 32)
	fLiabilitiesAndEquity, _ := strconv.ParseFloat(self.Xbrl.Fields["LiabilitiesAndEquity"], 32)

	lngBSCheck1 := FloatToString(fEquity - (fEquityAttributableToParent + fEquityAttributableToNoncontrollingInterest)) //self.Xbrl.Fields["Equity"] - (self.Xbrl.Fields["EquityAttributableToParent"] + self.Xbrl.Fields["EquityAttributableToNoncontrollingInterest"])
	lngBSCheck2 := FloatToString(fAssets - fLiabilitiesAndEquity)                                                       //self.Xbrl.Fields["Assets"] - self.Xbrl.Fields["LiabilitiesAndEquity"]

	if self.Xbrl.Fields["CurrentAssets"] == "0" && self.Xbrl.Fields["NoncurrentAssets"] == "0" && self.Xbrl.Fields["CurrentLiabilities"] == "0" && self.Xbrl.Fields["NoncurrentLiabilities"] == "0" {
		//if current assets/liabilities are zero && noncurrent assets/liabilities;{ don"t do this test because the balance sheet is not classified
		lngBSCheck3 := "0"
		lngBSCheck4 := "0"
	} else {
		//balance sheet IS classified
		fAssets, _ := strconv.ParseFloat(self.Xbrl.Fields["Assets"], 32)
		fCurrentAssets, _ := strconv.ParseFloat(self.Xbrl.Fields["CurrentAssets"], 32)
		fNoncurrentAssets, _ := strconv.ParseFloat(self.Xbrl.Fields["NoncurrentAssets"], 32)
		fLiabilities, _ := strconv.ParseFloat(self.Xbrl.Fields["Liabilities"], 32)
		fCurrentLiabilities, _ := strconv.ParseFloat(self.Xbrl.Fields["CurrentLiabilities"], 32)
		fNoncurrentLiabilities, _ := strconv.ParseFloat(self.Xbrl.Fields["NoncurrentLiabilities"], 32)

		lngBSCheck3 := FloatToString(fAssets - (fCurrentAssets + fNoncurrentAssets))                //self.Xbrl.Fields["Assets"] - (self.Xbrl.Fields["CurrentAssets"] + self.Xbrl.Fields["NoncurrentAssets"])
		lngBSCheck4 := FloatToString(fLiabilities - (fCurrentLiabilities + fNoncurrentLiabilities)) //self.Xbrl.Fields["Liabilities"] - (self.Xbrl.Fields["CurrentLiabilities"] + self.Xbrl.Fields["NoncurrentLiabilities"])
	}

	fLiabilitiesAndEquity, _ = strconv.ParseFloat(self.Xbrl.Fields["LiabilitiesAndEquity"], 32)
	fLiabilities, _ := strconv.ParseFloat(self.Xbrl.Fields["Liabilities"], 32)
	fCommitmentsAndContingencies, _ := strconv.ParseFloat(self.Xbrl.Fields["CommitmentsAndContingencies"], 32)
	fTemporaryEquity, _ := strconv.ParseFloat(self.Xbrl.Fields["TemporaryEquity"], 32)
	fEquity, _ = strconv.ParseFloat(self.Xbrl.Fields["Equity"], 32)

	lngBSCheck5 := FloatToString(fLiabilitiesAndEquity - (fLiabilities + fCommitmentsAndContingencies + fTemporaryEquity + fEquity)) //self.Xbrl.Fields["LiabilitiesAndEquity"] - (self.Xbrl.Fields["Liabilities"] + self.Xbrl.Fields["CommitmentsAndContingencies"] + self.Xbrl.Fields["TemporaryEquity"] + self.Xbrl.Fields["Equity"])
	/*
	   if lngBSCheck1{
	       print "BS1{ Equity(" , self.Xbrl.Fields["Equity"] , ") = EquityAttributableToParent(" , self.Xbrl.Fields["EquityAttributableToParent"] , ") , EquityAttributableToNoncontrollingInterest(" , self.Xbrl.Fields["EquityAttributableToNoncontrollingInterest"] , "){ " , lngBSCheck1
	   if lngBSCheck2{
	       print "BS2{ Assets(" , self.Xbrl.Fields["Assets"] , ") = LiabilitiesAndEquity(" , self.Xbrl.Fields["LiabilitiesAndEquity"] , "){ " , lngBSCheck2
	   if lngBSCheck3{
	       print "BS3{ Assets(" , self.Xbrl.Fields["Assets"] , ") = CurrentAssets(" , self.Xbrl.Fields["CurrentAssets"] , ") , NoncurrentAssets(" , self.Xbrl.Fields["NoncurrentAssets"] , "){ " , lngBSCheck3
	   if lngBSCheck4{
	       print "BS4{ Liabilities(" , self.Xbrl.Fields["Liabilities"] , ")= CurrentLiabilities(" , self.Xbrl.Fields["CurrentLiabilities"] , ") , NoncurrentLiabilities(" , self.Xbrl.Fields["NoncurrentLiabilities"] , "){ " , lngBSCheck4
	   if lngBSCheck5{
	       print "BS5{ Liabilities && Equity(" , self.Xbrl.Fields["LiabilitiesAndEquity"] , ")= Liabilities(" , self.Xbrl.Fields["Liabilities"] , ") , CommitmentsAndContingencies(" , self.Xbrl.Fields["CommitmentsAndContingencies"] , "), TemporaryEquity(" , self.Xbrl.Fields["TemporaryEquity"] , "), Equity(" , self.Xbrl.Fields["Equity"] , "){ " , lngBSCheck5
	*/

	//Income statement

	//Revenues
	self.Xbrl.Fields["Revenues"] = self.Xbrl.GetFactValue("Revenues", "Duration")
	if self.Xbrl.Fields["Revenues"] == "" {
		self.Xbrl.Fields["Revenues"] = self.Xbrl.GetFactValue("SalesRevenueNet", "Duration")
		if self.Xbrl.Fields["Revenues"] == "" {
			self.Xbrl.Fields["Revenues"] = self.Xbrl.GetFactValue("SalesRevenueServicesNet", "Duration")
			if self.Xbrl.Fields["Revenues"] == "" {
				self.Xbrl.Fields["Revenues"] = self.Xbrl.GetFactValue("RevenuesNetOfInterestExpense", "Duration")
				if self.Xbrl.Fields["Revenues"] == "" {
					self.Xbrl.Fields["Revenues"] = self.Xbrl.GetFactValue("RegulatedAndUnregulatedOperatingRevenue", "Duration")
					if self.Xbrl.Fields["Revenues"] == "" {
						self.Xbrl.Fields["Revenues"] = self.Xbrl.GetFactValue("HealthCareOrganizationRevenue", "Duration")
						if self.Xbrl.Fields["Revenues"] == "" {
							self.Xbrl.Fields["Revenues"] = self.Xbrl.GetFactValue("InterestAndDividendIncomeOperating", "Duration")
							if self.Xbrl.Fields["Revenues"] == "" {
								self.Xbrl.Fields["Revenues"] = self.Xbrl.GetFactValue("RealEstateRevenueNet", "Duration")
								if self.Xbrl.Fields["Revenues"] == "" {
									self.Xbrl.Fields["Revenues"] = self.Xbrl.GetFactValue("RevenueMineralSales", "Duration")
									if self.Xbrl.Fields["Revenues"] == "" {
										self.Xbrl.Fields["Revenues"] = self.Xbrl.GetFactValue("OilAndGasRevenue", "Duration")
										if self.Xbrl.Fields["Revenues"] == "" {
											self.Xbrl.Fields["Revenues"] = self.Xbrl.GetFactValue("FinancialServicesRevenue", "Duration")
											if self.Xbrl.Fields["Revenues"] == "" {
												self.Xbrl.Fields["Revenues"] = self.Xbrl.GetFactValue("RegulatedAndUnregulatedOperatingRevenue", "Duration")
												if self.Xbrl.Fields["Revenues"] == "" {
													self.Xbrl.Fields["Revenues"] = "0"
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	//CostOfRevenue
	self.Xbrl.Fields["CostOfRevenue"] = self.Xbrl.GetFactValue("CostOfRevenue", "Duration")
	if self.Xbrl.Fields["CostOfRevenue"] == "" {
		self.Xbrl.Fields["CostOfRevenue"] = self.Xbrl.GetFactValue("CostOfServices", "Duration")
		if self.Xbrl.Fields["CostOfRevenue"] == "" {
			self.Xbrl.Fields["CostOfRevenue"] = self.Xbrl.GetFactValue("CostOfGoodsSold", "Duration")
			if self.Xbrl.Fields["CostOfRevenue"] == "" {
				self.Xbrl.Fields["CostOfRevenue"] = self.Xbrl.GetFactValue("CostOfGoodsAndServicesSold", "Duration")
				if self.Xbrl.Fields["CostOfRevenue"] == "" {
					self.Xbrl.Fields["CostOfRevenue"] = "0"
				}
			}
		}
	}
	//GrossProfit
	self.Xbrl.Fields["GrossProfit"] = self.Xbrl.GetFactValue("GrossProfit", "Duration")
	if self.Xbrl.Fields["GrossProfit"] == "" {
		self.Xbrl.Fields["GrossProfit"] = self.Xbrl.GetFactValue("GrossProfit", "Duration")
		if self.Xbrl.Fields["GrossProfit"] == "" {
			self.Xbrl.Fields["GrossProfit"] = "0"
		}
	}
	//OperatingExpenses
	self.Xbrl.Fields["OperatingExpenses"] = self.Xbrl.GetFactValue("OperatingExpenses", "Duration")
	if self.Xbrl.Fields["OperatingExpenses"] == "" {
		self.Xbrl.Fields["OperatingExpenses"] = self.Xbrl.GetFactValue("OperatingCostsAndExpenses", "Duration") //This concept seems incorrect.
		if self.Xbrl.Fields["OperatingExpenses"] == "" {
			self.Xbrl.Fields["OperatingExpenses"] = "0"
		}
	}

	//CostsAndExpenses
	self.Xbrl.Fields["CostsAndExpenses"] = self.Xbrl.GetFactValue("CostsAndExpenses", "Duration")
	if self.Xbrl.Fields["CostsAndExpenses"] == "" {
		self.Xbrl.Fields["CostsAndExpenses"] = self.Xbrl.GetFactValue("CostsAndExpenses", "Duration")
		if self.Xbrl.Fields["CostsAndExpenses"] == "" {
			self.Xbrl.Fields["CostsAndExpenses"] = "0"
		}
	}
	//OtherOperatingIncome
	self.Xbrl.Fields["OtherOperatingIncome"] = self.Xbrl.GetFactValue("OtherOperatingIncome", "Duration")
	if self.Xbrl.Fields["OtherOperatingIncome"] == "" {
		self.Xbrl.Fields["OtherOperatingIncome"] = self.Xbrl.GetFactValue("OtherOperatingIncome", "Duration")
		if self.Xbrl.Fields["OtherOperatingIncome"] == "" {
			self.Xbrl.Fields["OtherOperatingIncome"] = "0"
		}
	}
	//OperatingIncomeLoss
	self.Xbrl.Fields["OperatingIncomeLoss"] = self.Xbrl.GetFactValue("OperatingIncomeLoss", "Duration")
	if self.Xbrl.Fields["OperatingIncomeLoss"] == "" {
		self.Xbrl.Fields["OperatingIncomeLoss"] = self.Xbrl.GetFactValue("OperatingIncomeLoss", "Duration")
		if self.Xbrl.Fields["OperatingIncomeLoss"] == "" {
			self.Xbrl.Fields["OperatingIncomeLoss"] = "0"
		}
	}
	//NonoperatingIncomeLoss
	self.Xbrl.Fields["NonoperatingIncomeLoss"] = self.Xbrl.GetFactValue("NonoperatingIncomeExpense", "Duration")
	if self.Xbrl.Fields["NonoperatingIncomeLoss"] == "" {
		self.Xbrl.Fields["NonoperatingIncomeLoss"] = self.Xbrl.GetFactValue("NonoperatingIncomeExpense", "Duration")
		if self.Xbrl.Fields["NonoperatingIncomeLoss"] == "" {
			self.Xbrl.Fields["NonoperatingIncomeLoss"] = "0"
		}
	}

	//InterestAndDebtExpense
	self.Xbrl.Fields["InterestAndDebtExpense"] = self.Xbrl.GetFactValue("InterestAndDebtExpense", "Duration")
	if self.Xbrl.Fields["InterestAndDebtExpense"] == "" {
		self.Xbrl.Fields["InterestAndDebtExpense"] = self.Xbrl.GetFactValue("InterestAndDebtExpense", "Duration")
		if self.Xbrl.Fields["InterestAndDebtExpense"] == "" {
			self.Xbrl.Fields["InterestAndDebtExpense"] = "0"
		}
	}

	//IncomeBeforeEquityMethodInvestments
	self.Xbrl.Fields["IncomeBeforeEquityMethodInvestments"] = self.Xbrl.GetFactValue("IncomeLossFromContinuingOperationsBeforeIncomeTaxesMinorityInterestAndIncomeLossFromEquityMethodInvestments", "Duration")
	if self.Xbrl.Fields["IncomeBeforeEquityMethodInvestments"] == "" {
		self.Xbrl.Fields["IncomeBeforeEquityMethodInvestments"] = self.Xbrl.GetFactValue("IncomeLossFromContinuingOperationsBeforeIncomeTaxesMinorityInterestAndIncomeLossFromEquityMethodInvestments", "Duration")
		if self.Xbrl.Fields["IncomeBeforeEquityMethodInvestments"] == "" {
			self.Xbrl.Fields["IncomeBeforeEquityMethodInvestments"] = "0"
		}
	}
	//IncomeFromEquityMethodInvestments
	self.Xbrl.Fields["IncomeFromEquityMethodInvestments"] = self.Xbrl.GetFactValue("IncomeLossFromEquityMethodInvestments", "Duration")
	if self.Xbrl.Fields["IncomeFromEquityMethodInvestments"] == "" {
		self.Xbrl.Fields["IncomeFromEquityMethodInvestments"] = self.Xbrl.GetFactValue("IncomeLossFromEquityMethodInvestments", "Duration")
		if self.Xbrl.Fields["IncomeFromEquityMethodInvestments"] == "" {
			self.Xbrl.Fields["IncomeFromEquityMethodInvestments"] = "0"
		}
	}
	//IncomeFromContinuingOperationsBeforeTax
	self.Xbrl.Fields["IncomeFromContinuingOperationsBeforeTax"] = self.Xbrl.GetFactValue("IncomeLossFromContinuingOperationsBeforeIncomeTaxesMinorityInterestAndIncomeLossFromEquityMethodInvestments", "Duration")
	if self.Xbrl.Fields["IncomeFromContinuingOperationsBeforeTax"] == "" {
		self.Xbrl.Fields["IncomeFromContinuingOperationsBeforeTax"] = self.Xbrl.GetFactValue("IncomeLossFromContinuingOperationsBeforeIncomeTaxesExtraordinaryItemsNoncontrollingInterest", "Duration")
		if self.Xbrl.Fields["IncomeFromContinuingOperationsBeforeTax"] == "" {
			self.Xbrl.Fields["IncomeFromContinuingOperationsBeforeTax"] = "0"
		}
	}
	//IncomeTaxExpenseBenefit
	self.Xbrl.Fields["IncomeTaxExpenseBenefit"] = self.Xbrl.GetFactValue("IncomeTaxExpenseBenefit", "Duration")
	if self.Xbrl.Fields["IncomeTaxExpenseBenefit"] == "" {
		self.Xbrl.Fields["IncomeTaxExpenseBenefit"] = self.Xbrl.GetFactValue("IncomeTaxExpenseBenefitContinuingOperations", "Duration")
		if self.Xbrl.Fields["IncomeTaxExpenseBenefit"] == "" {
			self.Xbrl.Fields["IncomeTaxExpenseBenefit"] = "0"
		}
	}
	//IncomeFromContinuingOperationsAfterTax
	self.Xbrl.Fields["IncomeFromContinuingOperationsAfterTax"] = self.Xbrl.GetFactValue("IncomeLossBeforeExtraordinaryItemsAndCumulativeEffectOfChangeInAccountingPrinciple", "Duration")
	if self.Xbrl.Fields["IncomeFromContinuingOperationsAfterTax"] == "" {
		self.Xbrl.Fields["IncomeFromContinuingOperationsAfterTax"] = self.Xbrl.GetFactValue("IncomeLossBeforeExtraordinaryItemsAndCumulativeEffectOfChangeInAccountingPrinciple", "Duration")
		if self.Xbrl.Fields["IncomeFromContinuingOperationsAfterTax"] == "" {
			self.Xbrl.Fields["IncomeFromContinuingOperationsAfterTax"] = "0"

		}
	}
	//IncomeFromDiscontinuedOperations
	self.Xbrl.Fields["IncomeFromDiscontinuedOperations"] = self.Xbrl.GetFactValue("IncomeLossFromDiscontinuedOperationsNetOfTax", "Duration")
	if self.Xbrl.Fields["IncomeFromDiscontinuedOperations"] == "" {
		self.Xbrl.Fields["IncomeFromDiscontinuedOperations"] = self.Xbrl.GetFactValue("DiscontinuedOperationGainLossOnDisposalOfDiscontinuedOperationNetOfTax", "Duration")
		if self.Xbrl.Fields["IncomeFromDiscontinuedOperations"] == "" {
			self.Xbrl.Fields["IncomeFromDiscontinuedOperations"] = self.Xbrl.GetFactValue("IncomeLossFromDiscontinuedOperationsNetOfTaxAttributableToReportingEntity", "Duration")
			if self.Xbrl.Fields["IncomeFromDiscontinuedOperations"] == "" {
				self.Xbrl.Fields["IncomeFromDiscontinuedOperations"] = "0"
			}
		}
	}

	//ExtraordaryItemsGainLoss
	self.Xbrl.Fields["ExtraordaryItemsGainLoss"] = self.Xbrl.GetFactValue("ExtraordinaryItemNetOfTax", "Duration")
	if self.Xbrl.Fields["ExtraordaryItemsGainLoss"] == "" {
		self.Xbrl.Fields["ExtraordaryItemsGainLoss"] = self.Xbrl.GetFactValue("ExtraordinaryItemNetOfTax", "Duration")
		if self.Xbrl.Fields["ExtraordaryItemsGainLoss"] == "" {
			self.Xbrl.Fields["ExtraordaryItemsGainLoss"] = "0"

		}
	}

	//NetIncomeLoss
	self.Xbrl.Fields["NetIncomeLoss"] = self.Xbrl.GetFactValue("ProfitLoss", "Duration")
	if self.Xbrl.Fields["NetIncomeLoss"] == "" {
		self.Xbrl.Fields["NetIncomeLoss"] = self.Xbrl.GetFactValue("NetIncomeLoss", "Duration")
		if self.Xbrl.Fields["NetIncomeLoss"] == "" {
			self.Xbrl.Fields["NetIncomeLoss"] = self.Xbrl.GetFactValue("NetIncomeLossAvailableToCommonStockholdersBasic", "Duration")
			if self.Xbrl.Fields["NetIncomeLoss"] == "" {
				self.Xbrl.Fields["NetIncomeLoss"] = self.Xbrl.GetFactValue("IncomeLossFromContinuingOperations", "Duration")
				if self.Xbrl.Fields["NetIncomeLoss"] == "" {
					self.Xbrl.Fields["NetIncomeLoss"] = self.Xbrl.GetFactValue("IncomeLossAttributableToParent", "Duration")
					if self.Xbrl.Fields["NetIncomeLoss"] == "" {
						self.Xbrl.Fields["NetIncomeLoss"] = self.Xbrl.GetFactValue("IncomeLossFromContinuingOperationsIncludingPortionAttributableToNoncontrollingInterest", "Duration")
						if self.Xbrl.Fields["NetIncomeLoss"] == "" {
							self.Xbrl.Fields["NetIncomeLoss"] = "0"
						}
					}
				}
			}
		}
	}

	//NetIncomeAvailableToCommonStockholdersBasic
	self.Xbrl.Fields["NetIncomeAvailableToCommonStockholdersBasic"] = self.Xbrl.GetFactValue("NetIncomeLossAvailableToCommonStockholdersBasic", "Duration")
	if self.Xbrl.Fields["NetIncomeAvailableToCommonStockholdersBasic"] == "" {
		self.Xbrl.Fields["NetIncomeAvailableToCommonStockholdersBasic"] = "0"
	}
	//PreferredStockDividendsAndOtherAdjustments
	self.Xbrl.Fields["PreferredStockDividendsAndOtherAdjustments"] = self.Xbrl.GetFactValue("PreferredStockDividendsAndOtherAdjustments", "Duration")
	if self.Xbrl.Fields["PreferredStockDividendsAndOtherAdjustments"] == "" {
		self.Xbrl.Fields["PreferredStockDividendsAndOtherAdjustments"] = "0"
	}
	//NetIncomeAttributableToNoncontrollingInterest
	self.Xbrl.Fields["NetIncomeAttributableToNoncontrollingInterest"] = self.Xbrl.GetFactValue("NetIncomeLossAttributableToNoncontrollingInterest", "Duration")
	if self.Xbrl.Fields["NetIncomeAttributableToNoncontrollingInterest"] == "" {
		self.Xbrl.Fields["NetIncomeAttributableToNoncontrollingInterest"] = "0"
	}
	//NetIncomeAttributableToParent
	self.Xbrl.Fields["NetIncomeAttributableToParent"] = self.Xbrl.GetFactValue("NetIncomeLoss", "Duration")
	if self.Xbrl.Fields["NetIncomeAttributableToParent"] == "" {
		self.Xbrl.Fields["NetIncomeAttributableToParent"] = "0"
	}

	//OtherComprehensiveIncome
	self.Xbrl.Fields["OtherComprehensiveIncome"] = self.Xbrl.GetFactValue("OtherComprehensiveIncomeLossNetOfTax", "Duration")
	if self.Xbrl.Fields["OtherComprehensiveIncome"] == "" {
		self.Xbrl.Fields["OtherComprehensiveIncome"] = self.Xbrl.GetFactValue("OtherComprehensiveIncomeLossNetOfTax", "Duration")
		if self.Xbrl.Fields["OtherComprehensiveIncome"] == "" {
			self.Xbrl.Fields["OtherComprehensiveIncome"] = "0"
		}
	}
	//ComprehensiveIncome
	self.Xbrl.Fields["ComprehensiveIncome"] = self.Xbrl.GetFactValue("ComprehensiveIncomeNetOfTaxIncludingPortionAttributableToNoncontrollingInterest", "Duration")
	if self.Xbrl.Fields["ComprehensiveIncome"] == "" {
		self.Xbrl.Fields["ComprehensiveIncome"] = self.Xbrl.GetFactValue("ComprehensiveIncomeNetOfTax", "Duration")
		if self.Xbrl.Fields["ComprehensiveIncome"] == "" {
			self.Xbrl.Fields["ComprehensiveIncome"] = "0"

		}
	}
	//ComprehensiveIncomeAttributableToParent
	self.Xbrl.Fields["ComprehensiveIncomeAttributableToParent"] = self.Xbrl.GetFactValue("ComprehensiveIncomeNetOfTax", "Duration")
	if self.Xbrl.Fields["ComprehensiveIncomeAttributableToParent"] == "" {
		self.Xbrl.Fields["ComprehensiveIncomeAttributableToParent"] = self.Xbrl.GetFactValue("ComprehensiveIncomeNetOfTax", "Duration")
		if self.Xbrl.Fields["ComprehensiveIncomeAttributableToParent"] == "" {
			self.Xbrl.Fields["ComprehensiveIncomeAttributableToParent"] = "0"
		}
	}
	//ComprehensiveIncomeAttributableToNoncontrollingInterest
	self.Xbrl.Fields["ComprehensiveIncomeAttributableToNoncontrollingInterest"] = self.Xbrl.GetFactValue("ComprehensiveIncomeNetOfTaxAttributableToNoncontrollingInterest", "Duration")
	if self.Xbrl.Fields["ComprehensiveIncomeAttributableToNoncontrollingInterest"] == "" {
		self.Xbrl.Fields["ComprehensiveIncomeAttributableToNoncontrollingInterest"] = self.Xbrl.GetFactValue("ComprehensiveIncomeNetOfTaxAttributableToNoncontrollingInterest", "Duration")
		if self.Xbrl.Fields["ComprehensiveIncomeAttributableToNoncontrollingInterest"] == "" {
			self.Xbrl.Fields["ComprehensiveIncomeAttributableToNoncontrollingInterest"] = "0"

		}
	}

	//Adjustments to income statement information
	//Impute: NonoperatingIncomeLossPlusInterestAndDebtExpense
	fNonoperatingIncomeLoss, _ := strconv.ParseFloat(self.Xbrl.Fields["NonoperatingIncomeLoss"], 32)
	fInterestAndDebtExpense, _ := strconv.ParseFloat(self.Xbrl.Fields["InterestAndDebtExpense"], 32)

	self.Xbrl.Fields["NonoperatingIncomeLossPlusInterestAndDebtExpense"] = FloatToString(fNonoperatingIncomeLoss + fInterestAndDebtExpense) //self.Xbrl.Fields["NonoperatingIncomeLoss"] + self.Xbrl.Fields["InterestAndDebtExpense"]

	//Impute: Net income available to common stockholders  (if it does not exist)
	if self.Xbrl.Fields["NetIncomeAvailableToCommonStockholdersBasic"] == "0" && self.Xbrl.Fields["PreferredStockDividendsAndOtherAdjustments"] == "0" && self.Xbrl.Fields["NetIncomeAttributableToParent"] != "0" {
		self.Xbrl.Fields["NetIncomeAvailableToCommonStockholdersBasic"] = self.Xbrl.Fields["NetIncomeAttributableToParent"]
	}
	//Impute NetIncomeLoss
	if self.Xbrl.Fields["NetIncomeLoss"] != "0" && self.Xbrl.Fields["IncomeFromContinuingOperationsAfterTax"] == "0" {
		fNetIncomeLoss, _ := strconv.ParseFloat(self.Xbrl.Fields["NetIncomeLoss"], 32)
		fIncomeFromDiscontinuedOperations, _ := strconv.ParseFloat(self.Xbrl.Fields["IncomeFromDiscontinuedOperations"], 32)
		fExtraordaryItemsGainLoss, _ := strconv.ParseFloat(self.Xbrl.Fields["ExtraordaryItemsGainLoss"], 32)

		self.Xbrl.Fields["IncomeFromContinuingOperationsAfterTax"] = FloatToString(fNetIncomeLoss - fIncomeFromDiscontinuedOperations - fExtraordaryItemsGainLoss) //self.Xbrl.Fields["NetIncomeLoss"] - self.Xbrl.Fields["IncomeFromDiscontinuedOperations"] - self.Xbrl.Fields["ExtraordaryItemsGainLoss"]
	}

	//Impute: Net income attributable to parent if it does not exist
	if self.Xbrl.Fields["NetIncomeAttributableToParent"] == "0" && self.Xbrl.Fields["NetIncomeAttributableToNoncontrollingInterest"] == "0" && self.Xbrl.Fields["NetIncomeLoss"] != "0" {
		self.Xbrl.Fields["NetIncomeAttributableToParent"] = self.Xbrl.Fields["NetIncomeLoss"]

	}
	//Inpute: PreferredStockDividendsAndOtherAdjustments
	if self.Xbrl.Fields["PreferredStockDividendsAndOtherAdjustments"] == "0" && self.Xbrl.Fields["NetIncomeAttributableToParent"] != "0" && self.Xbrl.Fields["NetIncomeAvailableToCommonStockholdersBasic"] != "0" {
		fNetIncomeAttributableToParent, _ := strconv.ParseFloat(self.Xbrl.Fields["NetIncomeAttributableToParent"], 32)
		fNetIncomeAvailableToCommonStockholdersBasic, _ := strconv.ParseFloat(self.Xbrl.Fields["NetIncomeAvailableToCommonStockholdersBasic"], 32)

		self.Xbrl.Fields["PreferredStockDividendsAndOtherAdjustments"] = FloatToString(fNetIncomeAttributableToParent - fNetIncomeAvailableToCommonStockholdersBasic) //self.Xbrl.Fields["NetIncomeAttributableToParent"] - self.Xbrl.Fields["NetIncomeAvailableToCommonStockholdersBasic"]
	}

	//Impute: comprehensive income
	if self.Xbrl.Fields["ComprehensiveIncomeAttributableToParent"] == "0" && self.Xbrl.Fields["ComprehensiveIncomeAttributableToNoncontrollingInterest"] == "0" && self.Xbrl.Fields["ComprehensiveIncome"] == "0" && self.Xbrl.Fields["OtherComprehensiveIncome"] == "0" {
		self.Xbrl.Fields["ComprehensiveIncome"] = self.Xbrl.Fields["NetIncomeLoss"]
	}
	//Inpute: other comprehensive income
	if self.Xbrl.Fields["ComprehensiveIncome"] != "0" && self.Xbrl.Fields["OtherComprehensiveIncome"] == "0" {
		fComprehensiveIncome, _ := strconv.ParseFloat(self.Xbrl.Fields["ComprehensiveIncome"], 32)
		fNetIncomeLoss, _ := strconv.ParseFloat(self.Xbrl.Fields["NetIncomeLoss"], 32)

		self.Xbrl.Fields["OtherComprehensiveIncome"] = FloatToString(fComprehensiveIncome - fNetIncomeLoss) //self.Xbrl.Fields["ComprehensiveIncome"] - self.Xbrl.Fields["NetIncomeLoss"]
	}

	//Inpute: comprehensive income attributable to parent if it does not exist
	if self.Xbrl.Fields["ComprehensiveIncomeAttributableToParent"] == "0" && self.Xbrl.Fields["ComprehensiveIncomeAttributableToNoncontrollingInterest"] == "0" && self.Xbrl.Fields["ComprehensiveIncome"] != "0" {
		self.Xbrl.Fields["ComprehensiveIncomeAttributableToParent"] = self.Xbrl.Fields["ComprehensiveIncome"]
	}

	//Inpute: IncomeFromContinuingOperations*Before*Tax
	if self.Xbrl.Fields["IncomeBeforeEquityMethodInvestments"] != "0" && self.Xbrl.Fields["IncomeFromEquityMethodInvestments"] != "0" && self.Xbrl.Fields["IncomeFromContinuingOperationsBeforeTax"] == "0" {
		fIncomeBeforeEquityMethodInvestments, _ := strconv.ParseFloat(self.Xbrl.Fields["IncomeBeforeEquityMethodInvestments"], 32)
		fIncomeFromEquityMethodInvestments, _ := strconv.ParseFloat(self.Xbrl.Fields["IncomeFromEquityMethodInvestments"], 32)

		self.Xbrl.Fields["IncomeFromContinuingOperationsBeforeTax"] = FloatToString(fIncomeBeforeEquityMethodInvestments + fIncomeFromEquityMethodInvestments) //self.Xbrl.Fields["IncomeBeforeEquityMethodInvestments"] + self.Xbrl.Fields["IncomeFromEquityMethodInvestments"]
	}
	//Inpute: IncomeFromContinuingOperations*Before*Tax2 (if income before tax is missing)
	if self.Xbrl.Fields["IncomeFromContinuingOperationsBeforeTax"] == "0" && self.Xbrl.Fields["IncomeFromContinuingOperationsAfterTax"] != "0" {
		fIncomeFromContinuingOperationsAfterTax, _ := strconv.ParseFloat(self.Xbrl.Fields["IncomeFromContinuingOperationsAfterTax"], 32)
		fIncomeTaxExpenseBenefit, _ := strconv.ParseFloat(self.Xbrl.Fields["IncomeTaxExpenseBenefit"], 32)

		self.Xbrl.Fields["IncomeFromContinuingOperationsBeforeTax"] = FloatToString(fIncomeFromContinuingOperationsAfterTax + fIncomeTaxExpenseBenefit) //self.Xbrl.Fields["IncomeFromContinuingOperationsAfterTax"] + self.Xbrl.Fields["IncomeTaxExpenseBenefit"]
	}
	//Inpute: IncomeFromContinuingOperations*After*Tax
	if self.Xbrl.Fields["IncomeFromContinuingOperationsAfterTax"] == "0" && (self.Xbrl.Fields["IncomeTaxExpenseBenefit"] != "0" || self.Xbrl.Fields["IncomeTaxExpenseBenefit"] == "0") && self.Xbrl.Fields["IncomeFromContinuingOperationsBeforeTax"] != "0" {
		fIncomeFromContinuingOperationsBeforeTax, _ := strconv.ParseFloat(self.Xbrl.Fields["IncomeFromContinuingOperationsBeforeTax"], 32)
		fIncomeTaxExpenseBenefit, _ := strconv.ParseFloat(self.Xbrl.Fields["IncomeTaxExpenseBenefit"], 32)

		self.Xbrl.Fields["IncomeFromContinuingOperationsAfterTax"] = FloatToString(fIncomeFromContinuingOperationsBeforeTax - fIncomeTaxExpenseBenefit) // self.Xbrl.Fields["IncomeFromContinuingOperationsBeforeTax"] - self.Xbrl.Fields["IncomeTaxExpenseBenefit"]
	}

	//Inpute: GrossProfit
	if self.Xbrl.Fields["GrossProfit"] == "0" && (self.Xbrl.Fields["Revenues"] != "0" && self.Xbrl.Fields["CostOfRevenue"] != "0") {
		fRevenues, _ := strconv.ParseFloat(self.Xbrl.Fields["Revenues"], 32)
		fCostOfRevenue, _ := strconv.ParseFloat(self.Xbrl.Fields["CostOfRevenue"], 32)

		self.Xbrl.Fields["GrossProfit"] = FloatToString(fRevenues - fCostOfRevenue) //self.Xbrl.Fields["Revenues"] - self.Xbrl.Fields["CostOfRevenue"]
	}

	//Inpute: Revenues
	if self.Xbrl.Fields["GrossProfit"] != "0" && (self.Xbrl.Fields["Revenues"] == "0" && self.Xbrl.Fields["CostOfRevenue"] != "0") {
		fGrossProfit, _ := strconv.ParseFloat(self.Xbrl.Fields["GrossProfit"], 32)
		fCostOfRevenue, _ := strconv.ParseFloat(self.Xbrl.Fields["CostOfRevenue"], 32)

		self.Xbrl.Fields["Revenues"] = FloatToString(fGrossProfit + fCostOfRevenue) //self.Xbrl.Fields["GrossProfit"] + self.Xbrl.Fields["CostOfRevenue"]
	}
	//Inpute: CostOfRevenue
	if self.Xbrl.Fields["GrossProfit"] != "0" && (self.Xbrl.Fields["Revenues"] != "0" && self.Xbrl.Fields["CostOfRevenue"] == "0") {
		fGrossProfit, _ := strconv.ParseFloat(self.Xbrl.Fields["GrossProfit"], 32)
		fRevenues, _ := strconv.ParseFloat(self.Xbrl.Fields["Revenues"], 32)

		self.Xbrl.Fields["CostOfRevenue"] = FloatToString(fGrossProfit + fRevenues) //self.Xbrl.Fields["GrossProfit"] + self.Xbrl.Fields["Revenues"]
	}
	//Inpute: CostsAndExpenses (would NEVER have costs && expenses if has gross profit, gross profit is multi-step && costs && expenses is single-step)
	if self.Xbrl.Fields["GrossProfit"] == "0" && self.Xbrl.Fields["CostsAndExpenses"] == "0" && (self.Xbrl.Fields["CostOfRevenue"] != "0" && self.Xbrl.Fields["OperatingExpenses"] != "0") {
		fCostOfRevenue, _ := strconv.ParseFloat(self.Xbrl.Fields["CostOfRevenue"], 32)
		fOperatingExpenses, _ := strconv.ParseFloat(self.Xbrl.Fields["OperatingExpenses"], 32)

		self.Xbrl.Fields["CostsAndExpenses"] = FloatToString(fCostOfRevenue + fOperatingExpenses) //self.Xbrl.Fields["CostOfRevenue"] + self.Xbrl.Fields["OperatingExpenses"]
	}
	//Inpute: CostsAndExpenses based on existance of both costs of revenues && operating expenses
	if self.Xbrl.Fields["CostsAndExpenses"] == "0" && self.Xbrl.Fields["OperatingExpenses"] != "0" && (self.Xbrl.Fields["CostOfRevenue"] != "0") {
		fCostOfRevenue, _ := strconv.ParseFloat(self.Xbrl.Fields["CostOfRevenue"], 32)
		fOperatingExpenses, _ := strconv.ParseFloat(self.Xbrl.Fields["OperatingExpenses"], 32)

		self.Xbrl.Fields["CostsAndExpenses"] = FloatToString(fCostOfRevenue + fOperatingExpenses) //self.Xbrl.Fields["CostOfRevenue"] + self.Xbrl.Fields["OperatingExpenses"]
	}
	//Inpute: CostsAndExpenses
	if self.Xbrl.Fields["GrossProfit"] == "0" && self.Xbrl.Fields["CostsAndExpenses"] == "0" && self.Xbrl.Fields["Revenues"] != "0" && self.Xbrl.Fields["OperatingIncomeLoss"] != "0" && self.Xbrl.Fields["OtherOperatingIncome"] != "0" {
		fRevenues, _ := strconv.ParseFloat(self.Xbrl.Fields["Revenues"], 32)
		fOperatingIncomeLoss, _ := strconv.ParseFloat(self.Xbrl.Fields["OperatingIncomeLoss"], 32)
		fOtherOperatingIncome, _ := strconv.ParseFloat(self.Xbrl.Fields["OtherOperatingIncome"], 32)

		self.Xbrl.Fields["CostsAndExpenses"] = FloatToString(fRevenues - fOperatingIncomeLoss - fOtherOperatingIncome) //self.Xbrl.Fields["Revenues"] - self.Xbrl.Fields["OperatingIncomeLoss"] - self.Xbrl.Fields["OtherOperatingIncome"]
	}
	//Inpute: OperatingExpenses based on existance of costs && expenses && cost of revenues
	if self.Xbrl.Fields["CostOfRevenue"] != "0" && self.Xbrl.Fields["CostsAndExpenses"] != "0" && self.Xbrl.Fields["OperatingExpenses"] == "0" {
		fCostsAndExpenses, _ := strconv.ParseFloat(self.Xbrl.Fields["CostsAndExpenses"], 32)
		fCostOfRevenue, _ := strconv.ParseFloat(self.Xbrl.Fields["CostOfRevenue"], 32)

		self.Xbrl.Fields["OperatingExpenses"] = FloatToString(fCostsAndExpenses - fCostOfRevenue) //self.Xbrl.Fields["CostsAndExpenses"] - self.Xbrl.Fields["CostOfRevenue"]
	}

	//Inpute: CostOfRevenues single-step method
	fRevenues, _ := strconv.ParseFloat(self.Xbrl.Fields["Revenues"], 32)
	fCostsAndExpenses, _ := strconv.ParseFloat(self.Xbrl.Fields["CostsAndExpenses"], 32)

	if self.Xbrl.Fields["Revenues"] != "0" && self.Xbrl.Fields["GrossProfit"] == "0" && (FloatToString(fRevenues-fCostsAndExpenses) == self.Xbrl.Fields["OperatingIncomeLoss"]) && self.Xbrl.Fields["OperatingExpenses"] == "0" && self.Xbrl.Fields["OtherOperatingIncome"] == "0" {
		fCostsAndExpenses, _ := strconv.ParseFloat(self.Xbrl.Fields["CostsAndExpenses"], 32)
		fOperatingExpenses, _ := strconv.ParseFloat(self.Xbrl.Fields["OperatingExpenses"], 32)

		self.Xbrl.Fields["CostOfRevenue"] = FloatToString(fCostsAndExpenses - fOperatingExpenses) //self.Xbrl.Fields["CostsAndExpenses"] - self.Xbrl.Fields["OperatingExpenses"]
	}

	//Inpute: IncomeBeforeEquityMethodInvestments
	if self.Xbrl.Fields["IncomeBeforeEquityMethodInvestments"] == "0" && self.Xbrl.Fields["IncomeFromContinuingOperationsBeforeTax"] != "0" {
		fIncomeFromContinuingOperationsBeforeTax, _ := strconv.ParseFloat(self.Xbrl.Fields["IncomeFromContinuingOperationsBeforeTax"], 32)
		fIncomeFromEquityMethodInvestments, _ := strconv.ParseFloat(self.Xbrl.Fields["IncomeFromEquityMethodInvestments"], 32)

		self.Xbrl.Fields["IncomeBeforeEquityMethodInvestments"] = FloatToString(fIncomeFromContinuingOperationsBeforeTax - fIncomeFromEquityMethodInvestments) //self.Xbrl.Fields["IncomeFromContinuingOperationsBeforeTax"] - self.Xbrl.Fields["IncomeFromEquityMethodInvestments"]
	}
	//Inpute: IncomeBeforeEquityMethodInvestments
	if self.Xbrl.Fields["OperatingIncomeLoss"] != "0" && (self.Xbrl.Fields["NonoperatingIncomeLoss"] != "0" && self.Xbrl.Fields["InterestAndDebtExpense"] == "0" && self.Xbrl.Fields["IncomeBeforeEquityMethodInvestments"] != "0") {
		fIncomeBeforeEquityMethodInvestments, _ := strconv.ParseFloat(self.Xbrl.Fields["IncomeBeforeEquityMethodInvestments"], 32)
		fOperatingIncomeLoss, _ := strconv.ParseFloat(self.Xbrl.Fields["OperatingIncomeLoss"], 32)
		fNonoperatingIncomeLoss, _ := strconv.ParseFloat(self.Xbrl.Fields["NonoperatingIncomeLoss"], 32)

		self.Xbrl.Fields["InterestAndDebtExpense"] = FloatToString(fIncomeBeforeEquityMethodInvestments - (fOperatingIncomeLoss + fNonoperatingIncomeLoss)) //self.Xbrl.Fields["IncomeBeforeEquityMethodInvestments"] - (self.Xbrl.Fields["OperatingIncomeLoss"] + self.Xbrl.Fields["NonoperatingIncomeLoss"])
	}
	//Inpute: OtherOperatingIncome
	if self.Xbrl.Fields["GrossProfit"] != "0" && (self.Xbrl.Fields["OperatingExpenses"] != "0" && self.Xbrl.Fields["OperatingIncomeLoss"] != "0") {
		fOperatingIncomeLoss, _ := strconv.ParseFloat(self.Xbrl.Fields["OperatingIncomeLoss"], 32)
		fGrossProfit, _ := strconv.ParseFloat(self.Xbrl.Fields["GrossProfit"], 32)
		fOperatingExpenses, _ := strconv.ParseFloat(self.Xbrl.Fields["OperatingExpenses"], 32)

		self.Xbrl.Fields["OtherOperatingIncome"] = FloatToString(fOperatingIncomeLoss - (fGrossProfit - fOperatingExpenses)) //self.Xbrl.Fields["OperatingIncomeLoss"] - (self.Xbrl.Fields["GrossProfit"] - self.Xbrl.Fields["OperatingExpenses"])
	}

	//Move IncomeFromEquityMethodInvestments
	if self.Xbrl.Fields["IncomeFromEquityMethodInvestments"] != "0" && self.Xbrl.Fields["IncomeBeforeEquityMethodInvestments"] != "0" && self.Xbrl.Fields["IncomeBeforeEquityMethodInvestments"] != self.Xbrl.Fields["IncomeFromContinuingOperationsBeforeTax"] {
		fIncomeFromContinuingOperationsBeforeTax, _ := strconv.ParseFloat(self.Xbrl.Fields["IncomeFromContinuingOperationsBeforeTax"], 32)
		fIncomeFromEquityMethodInvestments, _ := strconv.ParseFloat(self.Xbrl.Fields["IncomeFromEquityMethodInvestments"], 32)

		self.Xbrl.Fields["IncomeBeforeEquityMethodInvestments"] = FloatToString(fIncomeFromContinuingOperationsBeforeTax - fIncomeFromEquityMethodInvestments) //self.Xbrl.Fields["IncomeFromContinuingOperationsBeforeTax"] - self.Xbrl.Fields["IncomeFromEquityMethodInvestments"]

		fOperatingIncomeLoss, _ := strconv.ParseFloat(self.Xbrl.Fields["OperatingIncomeLoss"], 32)
		fIncomeFromEquityMethodInvestments, _ = strconv.ParseFloat(self.Xbrl.Fields["IncomeFromEquityMethodInvestments"], 32)

		self.Xbrl.Fields["OperatingIncomeLoss"] = FloatToString(fOperatingIncomeLoss - fIncomeFromEquityMethodInvestments) //self.Xbrl.Fields["OperatingIncomeLoss"] - self.Xbrl.Fields["IncomeFromEquityMethodInvestments"]
	}
	//DANGEROUS!!  May need to turn off. IS3 had 2"0"85 PASSES WITHOUT this imputing. if it is higher,{ keep the test
	//Inpute: OperatingIncomeLoss
	if self.Xbrl.Fields["OperatingIncomeLoss"] == "0" && self.Xbrl.Fields["IncomeBeforeEquityMethodInvestments"] != "0" {
		fIncomeBeforeEquityMethodInvestments, _ := strconv.ParseFloat(self.Xbrl.Fields["IncomeBeforeEquityMethodInvestments"], 32)
		fNonoperatingIncomeLoss, _ := strconv.ParseFloat(self.Xbrl.Fields["NonoperatingIncomeLoss"], 32)
		fInterestAndDebtExpense, _ := strconv.ParseFloat(self.Xbrl.Fields["InterestAndDebtExpense"], 32)

		self.Xbrl.Fields["OperatingIncomeLoss"] = FloatToString(fIncomeBeforeEquityMethodInvestments + fNonoperatingIncomeLoss - fInterestAndDebtExpense) //self.Xbrl.Fields["IncomeBeforeEquityMethodInvestments"] + self.Xbrl.Fields["NonoperatingIncomeLoss"] - self.Xbrl.Fields["InterestAndDebtExpense"]
	}

	fIncomeFromContinuingOperationsBeforeTax, _ := strconv.ParseFloat(self.Xbrl.Fields["IncomeFromContinuingOperationsBeforeTax"], 32)
	fOperatingIncomeLoss, _ := strconv.ParseFloat(self.Xbrl.Fields["OperatingIncomeLoss"], 32)

	self.Xbrl.Fields["NonoperatingIncomePlusInterestAndDebtExpensePlusIncomeFromEquityMethodInvestments"] = FloatToString(fIncomeFromContinuingOperationsBeforeTax - fOperatingIncomeLoss) //self.Xbrl.Fields["IncomeFromContinuingOperationsBeforeTax"] - self.Xbrl.Fields["OperatingIncomeLoss"]

	//NonoperatingIncomeLossPlusInterestAndDebtExpense
	if self.Xbrl.Fields["NonoperatingIncomeLossPlusInterestAndDebtExpense"] == "0" && self.Xbrl.Fields["NonoperatingIncomePlusInterestAndDebtExpensePlusIncomeFromEquityMethodInvestments"] != "0" {
		fNonoperatingIncomePlusInterestAndDebtExpensePlusIncomeFromEquityMethodInvestments, _ := strconv.ParseFloat(self.Xbrl.Fields["NonoperatingIncomePlusInterestAndDebtExpensePlusIncomeFromEquityMethodInvestments"], 32)
		fIncomeFromEquityMethodInvestments, _ := strconv.ParseFloat(self.Xbrl.Fields["IncomeFromEquityMethodInvestments"], 32)

		self.Xbrl.Fields["NonoperatingIncomeLossPlusInterestAndDebtExpense"] = FloatToString(fNonoperatingIncomePlusInterestAndDebtExpensePlusIncomeFromEquityMethodInvestments - fIncomeFromEquityMethodInvestments) //self.Xbrl.Fields["NonoperatingIncomePlusInterestAndDebtExpensePlusIncomeFromEquityMethodInvestments"] - self.Xbrl.Fields["IncomeFromEquityMethodInvestments"]
	}

	lngIS1 := (self.Xbrl.Fields["Revenues"] - self.Xbrl.Fields["CostOfRevenue"]) - self.Xbrl.Fields["GrossProfit"]
	lngIS2 := (self.Xbrl.Fields["GrossProfit"] - self.Xbrl.Fields["OperatingExpenses"] + self.Xbrl.Fields["OtherOperatingIncome"]) - self.Xbrl.Fields["OperatingIncomeLoss"]
	lngIS3 := (self.Xbrl.Fields["OperatingIncomeLoss"] + self.Xbrl.Fields["NonoperatingIncomeLossPlusInterestAndDebtExpense"]) - self.Xbrl.Fields["IncomeBeforeEquityMethodInvestments"]
	lngIS4 := (self.Xbrl.Fields["IncomeBeforeEquityMethodInvestments"] + self.Xbrl.Fields["IncomeFromEquityMethodInvestments"]) - self.Xbrl.Fields["IncomeFromContinuingOperationsBeforeTax"]
	lngIS5 := (self.Xbrl.Fields["IncomeFromContinuingOperationsBeforeTax"] - self.Xbrl.Fields["IncomeTaxExpenseBenefit"]) - self.Xbrl.Fields["IncomeFromContinuingOperationsAfterTax"]
	lngIS6 := (self.Xbrl.Fields["IncomeFromContinuingOperationsAfterTax"] + self.Xbrl.Fields["IncomeFromDiscontinuedOperations"] + self.Xbrl.Fields["ExtraordaryItemsGainLoss"]) - self.Xbrl.Fields["NetIncomeLoss"]
	lngIS7 := (self.Xbrl.Fields["NetIncomeAttributableToParent"] + self.Xbrl.Fields["NetIncomeAttributableToNoncontrollingInterest"]) - self.Xbrl.Fields["NetIncomeLoss"]
	lngIS8 := (self.Xbrl.Fields["NetIncomeAttributableToParent"] - self.Xbrl.Fields["PreferredStockDividendsAndOtherAdjustments"]) - self.Xbrl.Fields["NetIncomeAvailableToCommonStockholdersBasic"]
	lngIS9 := (self.Xbrl.Fields["ComprehensiveIncomeAttributableToParent"] + self.Xbrl.Fields["ComprehensiveIncomeAttributableToNoncontrollingInterest"]) - self.Xbrl.Fields["ComprehensiveIncome"]
	lngIS10 := (self.Xbrl.Fields["NetIncomeLoss"] + self.Xbrl.Fields["OtherComprehensiveIncome"]) - self.Xbrl.Fields["ComprehensiveIncome"]
	lngIS11 := self.Xbrl.Fields["OperatingIncomeLoss"] - (self.Xbrl.Fields["Revenues"] - self.Xbrl.Fields["CostsAndExpenses"] + self.Xbrl.Fields["OtherOperatingIncome"])

	/*
	   if lngIS1 != ""{
	       print "IS1{ GrossProfit(" , self.Xbrl.Fields["GrossProfit"] , ") = Revenues(" , self.Xbrl.Fields["Revenues"] , ") - CostOfRevenue(" , self.Xbrl.Fields["CostOfRevenue"] , "){ " , lngIS1
	   if lngIS2!= ""{
	       print "IS2{ OperatingIncomeLoss(" , self.Xbrl.Fields["OperatingIncomeLoss"] , ") = GrossProfit(" , self.Xbrl.Fields["GrossProfit"] , ") - OperatingExpenses(" , self.Xbrl.Fields["OperatingExpenses"] , ") , OtherOperatingIncome(" , self.Xbrl.Fields["OtherOperatingIncome"] , "){ " , lngIS2
	   if lngIS3!= ""{
	       print "IS3{ IncomeBeforeEquityMethodInvestments(" , self.Xbrl.Fields["IncomeBeforeEquityMethodInvestments"] , ") = OperatingIncomeLoss(" , self.Xbrl.Fields["OperatingIncomeLoss"] , ") - NonoperatingIncomeLoss(" , self.Xbrl.Fields["NonoperatingIncomeLoss"] , "), InterestAndDebtExpense(" , self.Xbrl.Fields["InterestAndDebtExpense"] , "){ " , lngIS3
	   if lngIS4!= ""{
	       print "IS4{ IncomeFromContinuingOperationsBeforeTax(" , self.Xbrl.Fields["IncomeFromContinuingOperationsBeforeTax"] , ") = IncomeBeforeEquityMethodInvestments(" , self.Xbrl.Fields["IncomeBeforeEquityMethodInvestments"] , ") , IncomeFromEquityMethodInvestments(" , self.Xbrl.Fields["IncomeFromEquityMethodInvestments"] , "){ " , lngIS4

	   if lngIS5{
	       print "IS5{ IncomeFromContinuingOperationsAfterTax(" , self.Xbrl.Fields["IncomeFromContinuingOperationsAfterTax"] , ") = IncomeFromContinuingOperationsBeforeTax(" , self.Xbrl.Fields["IncomeFromContinuingOperationsBeforeTax"] , ") - IncomeTaxExpenseBenefit(" , self.Xbrl.Fields["IncomeTaxExpenseBenefit"] , "){ " , lngIS5
	   if  lngIS6{
	       print "IS6{ NetIncomeLoss(" , self.Xbrl.Fields["NetIncomeLoss"] , ") = IncomeFromContinuingOperationsAfterTax(" , self.Xbrl.Fields["IncomeFromContinuingOperationsAfterTax"] , ") , IncomeFromDiscontinuedOperations(" , self.Xbrl.Fields["IncomeFromDiscontinuedOperations"] , ") , ExtraordaryItemsGainLoss(" , self.Xbrl.Fields["ExtraordaryItemsGainLoss"] , "){ " , lngIS6
	   if lngIS7{
	       print "IS7{ NetIncomeLoss(" , self.Xbrl.Fields["NetIncomeLoss"] , ") = NetIncomeAttributableToParent(" , self.Xbrl.Fields["NetIncomeAttributableToParent"] , ") , NetIncomeAttributableToNoncontrollingInterest(" , self.Xbrl.Fields["NetIncomeAttributableToNoncontrollingInterest"] , "){ " , lngIS7
	   if lngIS8{
	       print "IS8{ NetIncomeAvailableToCommonStockholdersBasic(" , self.Xbrl.Fields["NetIncomeAvailableToCommonStockholdersBasic"] , ") = NetIncomeAttributableToParent(" , self.Xbrl.Fields["NetIncomeAttributableToParent"] , ") - PreferredStockDividendsAndOtherAdjustments(" , self.Xbrl.Fields["PreferredStockDividendsAndOtherAdjustments"] , "){ " , lngIS8
	   if lngIS9{
	       print "IS9{ ComprehensiveIncome(" , self.Xbrl.Fields["ComprehensiveIncome"] , ") = ComprehensiveIncomeAttributableToParent(" , self.Xbrl.Fields["ComprehensiveIncomeAttributableToParent"] , ") , ComprehensiveIncomeAttributableToNoncontrollingInterest(" , self.Xbrl.Fields["ComprehensiveIncomeAttributableToNoncontrollingInterest"] , "){ " , lngIS9
	   if lngIS1"0"{
	       print "IS1"0"{ ComprehensiveIncome(" , self.Xbrl.Fields["ComprehensiveIncome"] , ") = NetIncomeLoss(" , self.Xbrl.Fields["NetIncomeLoss"] , ") , OtherComprehensiveIncome(" , self.Xbrl.Fields["OtherComprehensiveIncome"] , "){ " , lngIS1"0"
	   if lngIS11{
	       print "IS11{ OperatingIncomeLoss(" , self.Xbrl.Fields["OperatingIncomeLoss"] , ") = Revenues(" , self.Xbrl.Fields["Revenues"] , ") - CostsAndExpenses(" , self.Xbrl.Fields["CostsAndExpenses"] , ") , OtherOperatingIncome(" , self.Xbrl.Fields["OtherOperatingIncome"] , "){ " , lngIS11
	*/

	//////Cash flow statement

	//NetCashFlow
	self.Xbrl.Fields["NetCashFlow"] = self.Xbrl.GetFactValue("CashAndCashEquivalentsPeriodIncreaseDecrease", "Duration")
	if self.Xbrl.Fields["NetCashFlow"] == "" {
		self.Xbrl.Fields["NetCashFlow"] = self.Xbrl.GetFactValue("CashPeriodIncreaseDecrease", "Duration")
		if self.Xbrl.Fields["NetCashFlow"] == "" {
			self.Xbrl.Fields["NetCashFlow"] = self.Xbrl.GetFactValue("NetCashProvidedByUsedInContinuingOperations", "Duration")
			if self.Xbrl.Fields["NetCashFlow"] == "" {
				self.Xbrl.Fields["NetCashFlow"] = "0"
			}
		}
	}
	//NetCashFlowsOperating
	self.Xbrl.Fields["NetCashFlowsOperating"] = self.Xbrl.GetFactValue("NetCashProvidedByUsedInOperatingActivities", "Duration")
	if self.Xbrl.Fields["NetCashFlowsOperating"] == "" {
		self.Xbrl.Fields["NetCashFlowsOperating"] = "0"
	}
	//NetCashFlowsInvesting
	self.Xbrl.Fields["NetCashFlowsInvesting"] = self.Xbrl.GetFactValue("NetCashProvidedByUsedInInvestingActivities", "Duration")
	if self.Xbrl.Fields["NetCashFlowsInvesting"] == "" {
		self.Xbrl.Fields["NetCashFlowsInvesting"] = "0"
	}
	//NetCashFlowsFinancing
	self.Xbrl.Fields["NetCashFlowsFinancing"] = self.Xbrl.GetFactValue("NetCashProvidedByUsedInFinancingActivities", "Duration")
	if self.Xbrl.Fields["NetCashFlowsFinancing"] == "" {
		self.Xbrl.Fields["NetCashFlowsFinancing"] = "0"
	}
	//NetCashFlowsOperatingContinuing
	self.Xbrl.Fields["NetCashFlowsOperatingContinuing"] = self.Xbrl.GetFactValue("NetCashProvidedByUsedInOperatingActivitiesContinuingOperations", "Duration")
	if self.Xbrl.Fields["NetCashFlowsOperatingContinuing"] == "" {
		self.Xbrl.Fields["NetCashFlowsOperatingContinuing"] = "0"
	}
	//NetCashFlowsInvestingContinuing
	self.Xbrl.Fields["NetCashFlowsInvestingContinuing"] = self.Xbrl.GetFactValue("NetCashProvidedByUsedInInvestingActivitiesContinuingOperations", "Duration")
	if self.Xbrl.Fields["NetCashFlowsInvestingContinuing"] == "" {
		self.Xbrl.Fields["NetCashFlowsInvestingContinuing"] = "0"
	}
	//NetCashFlowsFinancingContinuing
	self.Xbrl.Fields["NetCashFlowsFinancingContinuing"] = self.Xbrl.GetFactValue("NetCashProvidedByUsedInFinancingActivitiesContinuingOperations", "Duration")
	if self.Xbrl.Fields["NetCashFlowsFinancingContinuing"] == "" {
		self.Xbrl.Fields["NetCashFlowsFinancingContinuing"] = "0"
	}
	//NetCashFlowsOperatingDiscontinued
	self.Xbrl.Fields["NetCashFlowsOperatingDiscontinued"] = self.Xbrl.GetFactValue("CashProvidedByUsedInOperatingActivitiesDiscontinuedOperations", "Duration")
	if self.Xbrl.Fields["NetCashFlowsOperatingDiscontinued"] == "" {
		self.Xbrl.Fields["NetCashFlowsOperatingDiscontinued"] = "0"
	}
	//NetCashFlowsInvestingDiscontinued
	self.Xbrl.Fields["NetCashFlowsInvestingDiscontinued"] = self.Xbrl.GetFactValue("CashProvidedByUsedInInvestingActivitiesDiscontinuedOperations", "Duration")
	if self.Xbrl.Fields["NetCashFlowsInvestingDiscontinued"] == "" {
		self.Xbrl.Fields["NetCashFlowsInvestingDiscontinued"] = "0"
	}
	//NetCashFlowsFinancingDiscontinued
	self.Xbrl.Fields["NetCashFlowsFinancingDiscontinued"] = self.Xbrl.GetFactValue("CashProvidedByUsedInFinancingActivitiesDiscontinuedOperations", "Duration")
	if self.Xbrl.Fields["NetCashFlowsFinancingDiscontinued"] == "" {
		self.Xbrl.Fields["NetCashFlowsFinancingDiscontinued"] = "0"
	}
	//NetCashFlowsDiscontinued
	self.Xbrl.Fields["NetCashFlowsDiscontinued"] = self.Xbrl.GetFactValue("NetCashProvidedByUsedInDiscontinuedOperations", "Duration")
	if self.Xbrl.Fields["NetCashFlowsDiscontinued"] == "" {
		self.Xbrl.Fields["NetCashFlowsDiscontinued"] = "0"
	}
	//ExchangeGainsLosses
	self.Xbrl.Fields["ExchangeGainsLosses"] = self.Xbrl.GetFactValue("EffectOfExchangeRateOnCashAndCashEquivalents", "Duration")
	if self.Xbrl.Fields["ExchangeGainsLosses"] == "" {
		self.Xbrl.Fields["ExchangeGainsLosses"] = self.Xbrl.GetFactValue("EffectOfExchangeRateOnCashAndCashEquivalentsContinuingOperations", "Duration")
		if self.Xbrl.Fields["ExchangeGainsLosses"] == "" {
			self.Xbrl.Fields["ExchangeGainsLosses"] = self.Xbrl.GetFactValue("CashProvidedByUsedInFinancingActivitiesDiscontinuedOperations", "Duration")
			if self.Xbrl.Fields["ExchangeGainsLosses"] == "" {
				self.Xbrl.Fields["ExchangeGainsLosses"] = "0"
			}
		}
	}

	////////Adjustments
	//Inpute: total net cash flows discontinued if not reported
	if self.Xbrl.Fields["NetCashFlowsDiscontinued"] == "0" {
		self.Xbrl.Fields["NetCashFlowsDiscontinued"] = self.Xbrl.Fields["NetCashFlowsOperatingDiscontinued"] + self.Xbrl.Fields["NetCashFlowsInvestingDiscontinued"] + self.Xbrl.Fields["NetCashFlowsFinancingDiscontinued"]

	}
	//Inpute: cash flows from continuing
	if self.Xbrl.Fields["NetCashFlowsOperating"] != "0" && self.Xbrl.Fields["NetCashFlowsOperatingContinuing"] == "0" {
		self.Xbrl.Fields["NetCashFlowsOperatingContinuing"] = self.Xbrl.Fields["NetCashFlowsOperating"] - self.Xbrl.Fields["NetCashFlowsOperatingDiscontinued"]
	}
	if self.Xbrl.Fields["NetCashFlowsInvesting"] != "0" && self.Xbrl.Fields["NetCashFlowsInvestingContinuing"] == "0" {
		self.Xbrl.Fields["NetCashFlowsInvestingContinuing"] = self.Xbrl.Fields["NetCashFlowsInvesting"] - self.Xbrl.Fields["NetCashFlowsInvestingDiscontinued"]
	}
	if self.Xbrl.Fields["NetCashFlowsFinancing"] != "0" && self.Xbrl.Fields["NetCashFlowsFinancingContinuing"] == "0" {
		self.Xbrl.Fields["NetCashFlowsFinancingContinuing"] = self.Xbrl.Fields["NetCashFlowsFinancing"] - self.Xbrl.Fields["NetCashFlowsFinancingDiscontinued"]
	}

	if self.Xbrl.Fields["NetCashFlowsOperating"] == "0" && self.Xbrl.Fields["NetCashFlowsOperatingContinuing"] != "0" && self.Xbrl.Fields["NetCashFlowsOperatingDiscontinued"] == "0" {
		self.Xbrl.Fields["NetCashFlowsOperating"] = self.Xbrl.Fields["NetCashFlowsOperatingContinuing"]
	}
	if self.Xbrl.Fields["NetCashFlowsInvesting"] == "0" && self.Xbrl.Fields["NetCashFlowsInvestingContinuing"] != "0" && self.Xbrl.Fields["NetCashFlowsInvestingDiscontinued"] == "0" {
		self.Xbrl.Fields["NetCashFlowsInvesting"] = self.Xbrl.Fields["NetCashFlowsInvestingContinuing"]
	}
	if self.Xbrl.Fields["NetCashFlowsFinancing"] == "0" && self.Xbrl.Fields["NetCashFlowsFinancingContinuing"] != "0" && self.Xbrl.Fields["NetCashFlowsFinancingDiscontinued"] == "0" {
		self.Xbrl.Fields["NetCashFlowsFinancing"] = self.Xbrl.Fields["NetCashFlowsFinancingContinuing"]

	}
	self.Xbrl.Fields["NetCashFlowsContinuing"] = self.Xbrl.Fields["NetCashFlowsOperatingContinuing"] + self.Xbrl.Fields["NetCashFlowsInvestingContinuing"] + self.Xbrl.Fields["NetCashFlowsFinancingContinuing"]

	//Inpute: if net cash flow is missing,{ this tries to figure out the value by adding up the detail
	if self.Xbrl.Fields["NetCashFlow"] == "0" && (self.Xbrl.Fields["NetCashFlowsOperating"] != "0" || self.Xbrl.Fields["NetCashFlowsInvesting"] != "0" || self.Xbrl.Fields["NetCashFlowsFinancing"] != "0") {
		self.Xbrl.Fields["NetCashFlow"] = self.Xbrl.Fields["NetCashFlowsOperating"] + self.Xbrl.Fields["NetCashFlowsInvesting"] + self.Xbrl.Fields["NetCashFlowsFinancing"]
	}

	lngCF1 := self.Xbrl.Fields["NetCashFlow"] - (self.Xbrl.Fields["NetCashFlowsOperating"] + self.Xbrl.Fields["NetCashFlowsInvesting"] + self.Xbrl.Fields["NetCashFlowsFinancing"] + self.Xbrl.Fields["ExchangeGainsLosses"])
	if lngCF1 != "0" && (self.Xbrl.Fields["NetCashFlow"]-(self.Xbrl.Fields["NetCashFlowsOperating"]+self.Xbrl.Fields["NetCashFlowsInvesting"]+self.Xbrl.Fields["NetCashFlowsFinancing"]+self.Xbrl.Fields["ExchangeGainsLosses"]) == (self.Xbrl.Fields["ExchangeGainsLosses"] * -1)) {
		lngCF1 := 888888
	}
	//What is going on here is that 171 filers compute net cash flow differently than everyone else.
	//What I am doing is marking these by setting the value of the test to a number 888888 which would never occur naturally, so that I can differentiate this from errors.
	lngCF2 := self.Xbrl.Fields["NetCashFlowsContinuing"] - (self.Xbrl.Fields["NetCashFlowsOperatingContinuing"] + self.Xbrl.Fields["NetCashFlowsInvestingContinuing"] + self.Xbrl.Fields["NetCashFlowsFinancingContinuing"])
	lngCF3 := self.Xbrl.Fields["NetCashFlowsDiscontinued"] - (self.Xbrl.Fields["NetCashFlowsOperatingDiscontinued"] + self.Xbrl.Fields["NetCashFlowsInvestingDiscontinued"] + self.Xbrl.Fields["NetCashFlowsFinancingDiscontinued"])
	lngCF4 := self.Xbrl.Fields["NetCashFlowsOperating"] - (self.Xbrl.Fields["NetCashFlowsOperatingContinuing"] + self.Xbrl.Fields["NetCashFlowsOperatingDiscontinued"])
	lngCF5 := self.Xbrl.Fields["NetCashFlowsInvesting"] - (self.Xbrl.Fields["NetCashFlowsInvestingContinuing"] + self.Xbrl.Fields["NetCashFlowsInvestingDiscontinued"])
	lngCF6 := self.Xbrl.Fields["NetCashFlowsFinancing"] - (self.Xbrl.Fields["NetCashFlowsFinancingContinuing"] + self.Xbrl.Fields["NetCashFlowsFinancingDiscontinued"])

	/*
	   if lngCF1{
	       print "CF1{ NetCashFlow(" , self.Xbrl.Fields["NetCashFlow"] , ") = (NetCashFlowsOperating(" , self.Xbrl.Fields["NetCashFlowsOperating"] , ") , (NetCashFlowsInvesting(" , self.Xbrl.Fields["NetCashFlowsInvesting"] , ") , (NetCashFlowsFinancing(" , self.Xbrl.Fields["NetCashFlowsFinancing"] , ") , ExchangeGainsLosses(" , self.Xbrl.Fields["ExchangeGainsLosses"] , "){ " , lngCF1
	   if lngCF2{
	       print "CF2{ NetCashFlowsContinuing(" , self.Xbrl.Fields["NetCashFlowsContinuing"] , ") = NetCashFlowsOperatingContinuing(" , self.Xbrl.Fields["NetCashFlowsOperatingContinuing"] , ") , NetCashFlowsInvestingContinuing(" , self.Xbrl.Fields["NetCashFlowsInvestingContinuing"] , ") , NetCashFlowsFinancingContinuing(" , self.Xbrl.Fields["NetCashFlowsFinancingContinuing"] , "){ " , lngCF2
	   if lngCF3{
	       print "CF3{ NetCashFlowsDiscontinued(" , self.Xbrl.Fields["NetCashFlowsDiscontinued"] , ") = NetCashFlowsOperatingDiscontinued(" , self.Xbrl.Fields["NetCashFlowsOperatingDiscontinued"] , ") , NetCashFlowsInvestingDiscontinued(" , self.Xbrl.Fields["NetCashFlowsInvestingDiscontinued"] , ") , NetCashFlowsFinancingDiscontinued(" , self.Xbrl.Fields["NetCashFlowsFinancingDiscontinued"] , "){ " , lngCF3
	   if lngCF4{
	       print "CF4{ NetCashFlowsOperating(" , self.Xbrl.Fields["NetCashFlowsOperating"] , ") = NetCashFlowsOperatingContinuing(" , self.Xbrl.Fields["NetCashFlowsOperatingContinuing"] , ") , NetCashFlowsOperatingDiscontinued(" , self.Xbrl.Fields["NetCashFlowsOperatingDiscontinued"] , "){ " , lngCF4
	   if lngCF5{
	       print "CF5{ NetCashFlowsInvesting(" , self.Xbrl.Fields["NetCashFlowsInvesting"] , ") = NetCashFlowsInvestingContinuing(" , self.Xbrl.Fields["NetCashFlowsInvestingContinuing"] , ") , NetCashFlowsInvestingDiscontinued(" , self.Xbrl.Fields["NetCashFlowsInvestingDiscontinued"] , "){ " , lngCF5
	   if lngCF6{
	       print "CF6{ NetCashFlowsFinancing(" , self.Xbrl.Fields["NetCashFlowsFinancing"] , ") = NetCashFlowsFinancingContinuing(" , self.Xbrl.Fields["NetCashFlowsFinancingContinuing"] , ") , NetCashFlowsFinancingDiscontinued(" , self.Xbrl.Fields["NetCashFlowsFinancingDiscontinued"] , "){ " , lngCF6
	*/

	/*

	   //Key ratios
	   try{
	       self.Xbrl.Fields["SGR"] = ((self.Xbrl.Fields["NetIncomeLoss"] / self.Xbrl.Fields["Revenues"]) * (1 + ((self.Xbrl.Fields["Assets"] - self.Xbrl.Fields["Equity"]) / self.Xbrl.Fields["Equity"]))) / ((1 / (self.Xbrl.Fields["Revenues"] / self.Xbrl.Fields["Assets"])) - (((self.Xbrl.Fields["NetIncomeLoss"] / self.Xbrl.Fields["Revenues"]) * (1 + (((self.Xbrl.Fields["Assets"] - self.Xbrl.Fields["Equity"]) / self.Xbrl.Fields["Equity"]))))))
	   except{
	       pass

	   try{
	       self.Xbrl.Fields["ROA"] = self.Xbrl.Fields["NetIncomeLoss"] / self.Xbrl.Fields["Assets"]
	   except{
	       pass

	   try{
	       self.Xbrl.Fields["ROE"] = self.Xbrl.Fields["NetIncomeLoss"] / self.Xbrl.Fields["Equity"]
	   except{
	       pass

	   try{
	       self.Xbrl.Fields["ROS"] = self.Xbrl.Fields["NetIncomeLoss"] / self.Xbrl.Fields["Revenues"]
	   except{
	       pass
	*/
}
