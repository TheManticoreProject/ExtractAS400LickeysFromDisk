package products

// Source: https://www.ibm.com/docs/en/i/7.4.0?topic=reference-media-labels-their-contents

var NamesFromPrdid = map[string]string{
	"5770SS1": "IBM i Operating System",
	"5770BR1": "IBM Backup, Recovery, and Media Services",
	"5733QMG": "IBM QMGTOOLS",
	"5770DG1": "IBM HTTP Server for i",
	"5770JV1": "IBM Developer Kit for Java",
	"5733RDI": "IBM Rational Developer",
	"5770XE1": "IBM i Access for Windows",
	"5770WDS": "IBM Rational Development Studio for i",
	"5733WQX": "IBM DB2 Web Query for i",
	"5770PT1": "IBM Performance Tools for i",
	"5770ST1": "IBM Db2 Query Manager and SQL Development Kit for i",
	"5770HAS": "IBM PowerHA SystemMirror for i",
	"5733ICC": "IBM Cloud Storage Solutions for i",
	"5761CM1": "IBM Communications Utilities for System i",
	"5770DBM": "IBM Db2 Mirror for i",
	"5770UME": "IBM Universal Manageability Enablement",
	"5770SM1": "IBM System Manager for i",
	"5770MG1": "IBM Managed System Services for i",
}

var NamesFromPrdidAndFeature = map[string]map[string]string{
	"5722QU1": {
		"5050": "Query",
	},
	"5722SS1": {
		"5050": "i5/OS",
		"5051": "i5/OS, IBM i - Per Processor",
		"5052": "i5/OS, IBM i - Users",
		"5053": "i5/OS, IBM i - Appl Server Processor",
		"5103": "Media & Storage Extensions",
		"5112": "PSF/400 1-45 IPM Printer Support",
		"5113": "PSF/400 1-100 IPM Printer Support",
		"5114": "PSF/400 Anyspeed Printer Support",
		"5116": "HA Switchable Resources",
		"5117": "HA Journal Performance",
	},
	"5770SS1": {
		"5050": "IBM i",
		"5051": "IBM i - Per Processor",
		"5052": "IBM i Users",
		"5053": "IBM i Appl Server Processor",
		"5103": "Media & Storage Extensions",
		"5112": "Print 1-55 Impressions/Minute",
		"5113": "Print 1-100 Impressions/Minute",
		"5114": "Print Anyspeed Impressions/Minute",
		"5116": "HA Switchable Resources",
		"5117": "HA Journal Performance",
		"5120": "Db2 Data Mirroring",
	},
	"5770BR1": {
		"5050": "BRMS",
		"5101": "BRMS Network Feature",
		"5102": "BRMS Advanced Functions",
	},
	"5770HAS": {
		"5050": "PowerHA SystemMirror i - Standard",
		"5051": "PowerHA SystemMirror i (no edition) / Enterprise",
		"5052": "PowerHA SystemMirror i - Standard (additional)",
		"5053": "PowerHA SystemMirror i - Express",
	},
	"5770PT1": {
		"5050": "Performance Tools",
		"5101": "Performance Tools Manager",
		"5102": "Performance Tools Agent",
		"5103": "Performance Tools Job Watcher",
	},
	"5770ST1": {
		"5050": "DB2 and SQL Dev Kit",
	},
	"5733WQX": {
		"5050": "Web Query for System i Base",
		"5101": "Web Query for System i Express",
		"5102": "Web Query for System i Standard",
		"5104": "Web Query User Registration",
		"5105": "Development Workbench User Registration",
		"5106": "Runtime Group Profile Registration",
		"5107": "JDE Adapter",
		"5108": "IBM Web Query DataMigrator for i",
		"5109": "Web Query Scheduler Edition",
		"5110": "Web Query Runtime User Edition",
	},
	"5722AF1": {
		"5050": "AFP Utilities",
	},
	"5722AP1": {
		"5050": "Advanced DBCS Printer Support",
		"5101": "IPDS",
	},
	"5722BR1": {
		"5050": "BRMS",
		"5101": "BRMS Network Feature",
		"5102": "BRMS Advanced Functions",
	},
	"5722CM1": {
		"5050": "Communication Utilities",
	},
	"5722DE1": {
		"5050": "DB2 UDB Extenders",
	},
	"5722IP1": {
		"5050": "InfoPrint Server for iSeries",
		"5101": "PS to AFP Transform",
	},
	"5722JS1": {
		"5050": "Advanced Job Scheduler",
	},
	"5722PT1": {
		"5050": "Performance Tools",
		"5101": "Performance Tools Manager",
		"5102": "Performance Tools Agent",
	},
	"5722WDS": {
		"5050": "WebSphere Development Studio",
	},
	"5722XW1": {
		"5101": "iSeries Access Family Per User / Proc Based",
	},
	"5733ICC": {
		"5101": "Cloud Storage Solutions for i",
		"5102": "Cloud Storage Solutions advanced",
	},
	"5733ICL": {
		"5050": "IBM DataMirror iCluster",
		"5051": "IBM DataMirror iBalance",
	},
	"5733ID1": {
		"5050": "InfoPrint Designer for iSeries",
	},
	"5733OAR": {
		"5050": "IBM Rational Open Access RPG Edition",
	},
	"5733QU2": {
		"5050": "IBM DB2 Web Query for System i",
		"5101": "Active Reports",
		"5102": "OLAP Enablement",
		"5104": "Run Time User Enablement",
		"5105": "Spreadsheet Client",
		"5106": "SQL Server Adapter",
		"5107": "DB2 Web Adapter for JD Edwards",
	},
}

// Get the name of the product from the PRDID
func GetNameFromPrdid(prdid string) string {
	if name, ok := NamesFromPrdid[prdid]; ok {
		return name
	}
	return ""
}

// Get the name of the product from the PRDID and feature
func GetNameFromPrdidAndFeature(prdid string, feature string) string {
	if name, ok := NamesFromPrdidAndFeature[prdid]; ok {
		if featureName, ok := name[feature]; ok {
			return featureName
		}
	}
	return ""
}
