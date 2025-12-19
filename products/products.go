package products

// Sources:
// - https://www.ibm.com/docs/en/i/7.4.0?topic=reference-media-labels-their-contents
// - https://www.ibm.com/docs/en/entitled-systems-support?topic=migration-inventory-software-products-eligible

var NamesFromPrdid = map[string]string{
	"5621BGA": "SMART BUSINESS SW PACK FOR I",
	"5648F82": "PurifyPlus for AIX",
	"5722AC3": "CRYPTOGRAPH ACC PROV 128-BIT",
	"5722AF1": "AFP Utilities for iSeries",
	"5722AP1": "Advanced DBCS Printer Support for iSeries",
	"5722BR1": "Backup Recovery and Media Services for iSeries",
	"5722CE3": "Client Encryption 128-bit",
	"5722CM1": "Communications Utilities for i5/OS",
	"5722CR1": "IBM Cryptographic Support for iSeries",
	"5722DB1": "System/38 Utilities",
	"5722DE1": "DB2 Universal DB Extenders",
	"5722DFH": "CICS Transaction Server for i5/OS",
	"5722DP4": "DataPropagator for iSeries",
	"5722DR1": "Director for i5/OS",
	"5722DS1": "Business Graphics Utility for AS/400",
	"5722IMI": "SYS DIR VIRTUAL IMAGE MGT",
	"5722IP1": "Infoprint Server for iSeries",
	"5722JS1": "Job Scheduler for iSeries",
	"5722MG1": "Managed System Services for i5/OS",
	"5722NLV": "National Languages for i5/OS",
	"5722PD1": "APD",
	"5722PT1": "Performance Tools for iSeries",
	"5722QU1": "Query for iSeries",
	"5722SM1": "System Manager",
	"5722SS1": "i5/OS",
	"5722ST1": "DB2 and SQL Dev Kit",
	"5722SWD": "Director Premium Edition",
	"5722WDS": "WDS/400",
	"5722WE1": "Web Enablement",
	"5722WE2": "Web Enablement for i5/OS",
	"5722XW1": "iSeries Access",
	"5724I08": "XL Fortran Enterprise Ed V9",
	"5724I10": "XL C Enterprise Ed (AIX) V7",
	"5724I11": "XL C/C++ Enterprise Ed V7",
	"5724I15": "Tivoli Provisioning Manager",
	"5724K76": "XL Fortran Adv Ed Linux V9",
	"5724K77": "XL C/C++ Adv Ed for Linux",
	"5724L09": "Tivoli Provisioning Manager",
	"5724M11": "XL C Enterprise Ed V8 AIX",
	"5724M12": "XL C/C++ Enterprise Ed V8",
	"5724M13": "XL Fortran Enterprise Ed V10",
	"5724M16": "XL C/C++ Adv Ed for Linux V8",
	"5724M17": "XL Fortran Adv Ed Linux V10",
	"5724O33": "IBM Tivoli Usage Acct Mgr",
	"5724PL1": "IBM PL/I FOR AIX",
	"5724S70": "XL C EE for AIX V9.0",
	"5724S71": "XL C/C++ EE for AIX V9.0",
	"5724S72": "XL Fortran EE for AIX V11.1",
	"5724S73": "XL C/C++ AE for LINUX V9.0",
	"5724S74": "XL Fort AE for LINUX V11.1",
	"5724U80": "IBM XL C for AIX",
	"5724U81": "XL C/C++ for AIX",
	"5724U82": "XL Fortran for AIX",
	"5724U83": "XL C/C++ for Linux V10.1",
	"5724U84": "XL Fortran for Linux V12.1",
	"5724V62": "IBM COBOL for AIX V3.1",
	"5724X12": "XL C for AIX",
	"5724X13": "XL C/C++ for AIX",
	"5724X14": "XL C/C++ for Linux V11.1",
	"5724X15": "XL Fortran for AIX",
	"5724X16": "XL Fortran for Linux V13.1",
	"5724Z00": "Rational Team Concert Expr.",
	"5724Z01": "Rational Team Concert Std.",
	"5724Z87": "IBM COBOL for AIX V4.1",
	"5733AC1": "ARCAD RPG Converter",
	"5733AO1": "ARCAD Observer",
	"5733ARE": "IBM Application Runtime Expert for i",
	"5733B45": "IBM AFP Font Collection for i",
	"5733CO2": "Connect for iSeries",
	"5733CY1": "CRYPTOGRAPHIC DEVICE MANAGER",
	"5733CY2": "CRYPTOGRAPHIC DEVICE MANAGER",
	"5733CY3": "Cryptographic Device Manager",
	"5733DIR": "Director with VE Console",
	"5733DR1": "IBM Director with Virtualization Engine Console for i5/OS",
	"5733EWA": "VE Ent Wrkload Mgr / i5/OS",
	"5733EWM": "VE Enterprise Workload Mgr",
	"5733FXD": "Integrated Domino Fax for iSeries",
	"5733GT1": "IBM GRID TOOLBOX",
	"5733HAA": "HA ASSIST FOR I",
	"5733ICC": "Cloud Storage Solutions for i",
	"5733ICL": "DATAMIRROR ICLUSTER",
	"5733ICS": "DATAMIRROR ICLUSTER SMB",
	"5733ID1": "Infoprint Designer/400",
	"5733OAR": "IBM Rational Open Access RPG Edition",
	"5733OMF": "Find Text Search Server for DB2 for i5/OS",
	"5733PS1": "SECURE PERSPECTIVE",
	"5733QU2": "DB2 Web Query for System i",
	"5733QU3": "DB2 Web Query Report Broker",
	"5733QU4": "DB2 Web Query Software Development Kit",
	"5733RDA": "RD Power C/C++Development Tools for AIX",
	"5733RDB": "RD Power COBOL Development Tools for AIX",
	"5733RDC": "RD Power C/C++ Development Tools for Linux",
	"5733RDD": "RD Power Power Tools for i",
	"5733RDE": "RD Power Power Tools for AIX",
	"5733RDG": "RD Power RPG and COBOL Development Tools for i",
	"5733RDH": "RD Power C/C++Development Tools for AIX",
	"5733RDI": "Rational Developer for System i",
	"5733RDJ": "RD Power COBOL Development Tools for AIX",
	"5733RDP": "RD Power RPG and COBOL Development Tools for i",
	"5733RDS": "VE Resource Dependency Svc",
	"5733RDW": "Rational Developer for i",
	"5733RDX": "Rational Developer for AIX and Linux",
	"5733RTC": "IBM Rational Team Concert",
	"5733SC1": "PORTABLE UTILITIES FOR i5/OS",
	"5733SOA": "IBM Rational Developer for System i for SOA Construction",
	"5733SOC": "Rational Developer for i for SOA Construction",
	"5733VE1": "VIRTUALIZATION ENGINE",
	"5733WE3": "Web Enablement for i",
	"5733WQX": "IBM Db2 Web Query for i",
	"5733WSE": "Workplace Services Express",
	"5733XT1": "XML Toolkit for iSeries",
	"5733XT2": "XML Toolkit for IBM System i5",
	"5761AF1": "Advanced Function Printing Utilities",
	"5761AMT": "IBM Rational Application Management Tool Set for i",
	"5761AP1": "Advanced DBCS Printer Support for iSeries",
	"5761BR1": "Backup Recovery and Media Services for i5/OS",
	"5761CM1": "Communications Utilities for System i",
	"5761DB1": "System/38 Utilities for System i",
	"5761DE1": "DB2 Extenders Version 9.1 for i5/OS",
	"5761DFH": "CICS Transaction Server for i5/OS",
	"5761DP4": "DB2 DataPropagator for iSeries V8.1",
	"5761DS2": "Business Graphics Utility for System i",
	"5761HAS": "High Availability Solutions Manager",
	"5761JS1": "Advanced Job Scheduler for i5/OS",
	"5761MG1": "Managed System Services for i5/OS",
	"5761NLV": "National Languages for i5/OS",
	"5761PT1": "Performance Tools for i5/OS",
	"5761QU1": "Query for i5/OS",
	"5761SM1": "System Manager for i5/OS",
	"5761SS1": "IBM i",
	"5761ST1": "DB2 Query Manager and SQL Development Kit for i5/OS",
	"5761WDS": "WebSphere Development Studio for System i",
	"5761XW1": "System i Access Family",
	"57652B1": "AIX 7 Standard Edition Sub",
	"57652C1": "Pvt Cloud Ed with AIX 7 Sub",
	"57652E1": "AIX 7 Enterprise Edition Sub",
	"57656C1": "Private Cloud Edn Subs",
	"5765AEP": "IBM Systems Director Active Energy Manager for Power Systems",
	"5765AEZ": "AIX Enterprise Edition",
	"5765AIE": "IBM PowerAI Enterprise",
	"5765AME": "AIX Management Edition",
	"5765AMT": "AIX7.2 Monthly term offering",
	"5765ASC": "Platform Appl Svc Cntrl Pwr",
	"5765ASM": "IBM Systems Director Storage Control",
	"5765AVE": "PowerVM Lx86",
	"5765C41": "Parallel ESSL for AIX",
	"5765C42": "ESSL for AIX",
	"5765CBA": "Private Cloud Edn with AIX 7",
	"5765CD1": "AIX Enterprise Edition 7.1",
	"5765CD3": "AIX Enterprise Edition 7.2 and 7.3",
	"5765CMO": "Cloud Mgr w/OpenStack 4.2",
	"5765CMT": "CMC Monthly Term offering",
	"5765COB": "IBM COBOL for AIX V5",
	"5765D51": "PSSP for AIX V3",
	"5765D93": "Parallel Environment V3",
	"5765DR1": "IBM Director for pSystems",
	"5765DRG": "VM Recovery Manager DR",
	"5765DRP": "IBM Systems Director for Power Systems",
	"5765E61": "AIX 5L for POWER V5.1",
	"5765E62": "AIX 5L for POWER V5.2",
	"5765E68": "Performance Aide V3",
	"5765E69": "Loadleveler",
	"5765E72": "AIX Fast Connect for POWER",
	"5765E74": "AIX Performance Toolbox",
	"5765E82": "HAGEO/GeoRM V2",
	"5765E85": "AIXlink/X.25",
	"5765EAP": "ESSL for AIX on POWER",
	"5765ECB": "IBM Private Cloud Edition",
	"5765EEP": "IBM Systems Director Enterprise Edition for Power",
	"5765EGO": "Platform Resource Scheduler",
	"5765EL5": "Parallel ESSL for Linux on Power",
	"5765EMP": "IBM Systems Director VMControl Enterprise Edition",
	"5765EPL": "IBM Systems Director for PowerLinux",
	"5765ESL": "Parallel Engineering and Scientific Subroutine Library",
	"5765ESX": "Parallel Engineering and Scientific Subroutine Library",
	"5765EXP": "IBM Systems Director Express Edition for Power",
	"5765F56": "VisualAge C++ Pro for AIX V6",
	"5765F57": "C for AIX V6.0",
	"5765F62": "PowerHA V5",
	"5765F64": "GPFS for AIX 5L V2",
	"5765F67": "Cluster Systems Management",
	"5765F68": "IBM Infoprint Manager for AIX",
	"5765F70": "XL Fortran AIX",
	"5765F71": "XL Fortran RTE for Linux",
	"5765F82": "ESSL FOR AIX, V4",
	"5765F83": "PARALLEL ENVIRONMENT V4",
	"5765F84": "PARALLEL ESSL FOR AIX, V3",
	"5765FS2": "SAN File System Software V2",
	"5765G03": "AIX 5L for POWER V5.3",
	"5765G16": "Cluster Systems Management for Linux on POWER",
	"5765G17": "ESSL for Linux on POWER",
	"5765G18": "Parallel ESSL for Linux",
	"5765G20": "GPFS for Linux on pSeries",
	"5765G31": "Partition Load Manager",
	"5765G34": "Virtual I/O Server",
	"5765G62": "AIX Standard Edition",
	"5765G66": "General Parallel File System for POWER",
	"5765G67": "GPFS for Linux on Power",
	"5765G70": "PE for Linux on Power",
	"5765G71": "HACMP for Linux, V5",
	"5765G82": "IBM PowerSC Express Edition",
	"5765G83": "Workload Partitions Manager",
	"5765G90": "AIX Express Edition",
	"5765G97": "AIX Express Edition 7.1",
	"5765G98": "AIX Standard Edition 7.1, 7.2 and 7.3",
	"5765G99": "AIX Enterprise Edition",
	"5765H23": "PowerHA for AIX Standard Edition V6",
	"5765H24": "PowerHA for AIX Enterprise Edition V6",
	"5765H25": "Engineering and Scientific Subroutine Library for AIX",
	"5765H37": "PowerHA SystemMirror Enterprise Edition V7",
	"5765H38": "AIX 5.2 Workload Partitions for AIX 7",
	"5765H39": "PowerHA for AIX Standard Edition V7",
	"5765H43": "PHA SystemMirror Std ed subs",
	"5765H44": "PowerHA SM Ent ed subs",
	"5765H7E": "PowerHA SystemMirror Enterprise Edition Monthly Term V7.2",
	"5765H7S": "PowerHA SystemMirror Standard Edition Monthly Term V7.2",
	"5765HMA": "Power HMC Virtual Software Appliance for Power HW",
	"5765HMB": "HMC Virtual Appliance Power",
	"5765HMV": "Power HMC Virtual Software Appliance for x86 HW",
	"5765HMW": "HMC Virtual Appliance x86",
	"5765I2P": "IBM Systems Director VMControl Standard Edition",
	"5765IMP": "IBM Systems Director Virtual Image Management",
	"5765ISS": "ISS Service Sensor V7",
	"5765ITM": "ITM SE for System p",
	"5765J01": "IBM XL C for AIX",
	"5765J02": "IBM XL C/C++ for AIX",
	"5765J03": "IBM XL C/C++ for Linux",
	"5765J04": "IBM XL Fortran for AIX",
	"5765J05": "IBM XL Fortran for Linux",
	"5765J06": "IBM XL C for AIX",
	"5765J07": "IBM XL C/C++ for AIX",
	"5765J08": "IBM XL C/C++ for Linux",
	"5765J09": "IBM XL Fortran for AIX",
	"5765J10": "IBM XL Fortran for Linux",
	"5765J11": "IBM XL C for AIX V16",
	"5765J12": "IBM XL C/C++ for AIX V16",
	"5765J13": "IBM XL C/C++ for Linux V16",
	"5765J14": "IBM XL Fortran for AIX V16",
	"5765J15": "IBM XL Fortran for Linux V16",
	"5765J16": "XL C/C++ Mo Term",
	"5765J17": "Fortran for AIX Mo Term",
	"5765J18": "IBM Open XL C/C++ for AIX",
	"5765J19": "IBM Open XL Fortran for AIX",
	"5765J20": "Open XL C/C++ for Lnx on P",
	"5765J21": "Open XL Fortran for Lnx on P",
	"5765J22": "Open XL C/C++for Lnx Mo Term",
	"5765J23": "XL Fortran for Lnx Mo Term",
	"5765K01": "DB2 WSE 9.7 for HPC",
	"5765KV3": "IBM PowerKVM V3",
	"5765KVM": "IBM PowerKVM",
	"5765L22": "IBM PowerHA for Linux",
	"5765L40": "TWS LoadLeveler for AIX, V4",
	"5765L50": "Tivoli Workload Scheduler LoadLeveler for AIX",
	"5765L51": "ESSL for Linux on Power",
	"5765L61": "ESSL for Linux on Power V6",
	"5765LLP": "LoadLeveler for Linux on POWER",
	"5765LSF": "IBM Platform LSF - Standard Edition for Power",
	"5765M1M": "PowerSC MFA Monthly Term",
	"5765MCH": "Systems Director Management Console SDMC Hardware Appliance",
	"5765MCV": "Systems Director Management Console SDMC SVA",
	"5765MFA": "PowerSC MFA",
	"5765NCP": "IBM Systems Director Network Control",
	"5765OSF": "IBM Cloud Manager with OpenStack for IBM Flex Systems",
	"5765OSP": "IBM Cloud Manager with OpenStack for Power",
	"5765PAC": "Platform App Center Power V9",
	"5765PAI": "IBM PowerAI",
	"5765PCA": "IBM Spectrum Cluster Foundation for POWER",
	"5765PCS": "IBM Platform Cluster Manager - Standard Edition for POWER",
	"5765PD2": "Parallel Performance Toolkit for POWER",
	"5765PDP": "Parallel Environment Developer Edition for Linux on Power",
	"5765PEA": "Parallel Environment for AIX",
	"5765PED": "Parallel Environment Developer Edition for AIX",
	"5765PEL": "PE for Linux V5",
	"5765PEP": "System Dir AEM for POWER",
	"5765PER": "Parallel Environment Runtime Edition for AIX",
	"5765PHS": "IBM Power HPC Stack",
	"5765PIM": "PurePower Integrated Manager",
	"5765PPM": "Platform Process Manager for POWER",
	"5765PR2": "Parallel Environment Runtime Edition for Linux on Power",
	"5765PRP": "Parallel Environment Runtime Edition for Linux on Power",
	"5765PSA": "Platform Symphony - Advanced Edition for Power",
	"5765PSE": "IBM PowerSC Standard Edition",
	"5765PTS": "PowerSC Trusted Surveyor",
	"5765PVD": "PowerVM Enterprise Edition for Small Servers",
	"5765PVE": "PowerVM Enterprise Edition V2",
	"5765PVL": "PowerVM for IBM PowerLinux",
	"5765PVS": "PowerVM Standard Edition",
	"5765PVX": "PowerVM Express Edition",
	"5765PWO": "AIX Dynamic System Optimizer",
	"5765R11": "IBM Spectrum Computing Suite for High Performance Analytics",
	"5765R13": "ESSL on OpenPower",
	"5765R17": "IBM BOA",
	"5765RD2": "VM Recovery Manager DR subs",
	"5765RH2": "VM Recovery Manager HA subs",
	"5765RTM": "IBM Platform RTM for Power",
	"5765S1S": "IBM PowerSC Standard Edition Monthly Term V1.2",
	"5765S2S": "IBM PowerSC 2.1 Subscription",
	"5765SA7": "Platform Symphony Adv PWR V7",
	"5765SB3": "IBM SmartCloud Entry for Power bundle V3",
	"5765SC2": "IBM PowerSC V2",
	"5765SC3": "IBM SmartCloud Entry for Power",
	"5765SCP": "IBM SmartCloud Entry for Flex system",
	"5765SEP": "IBM Systems Director Standard Edition for Power",
	"5765SF3": "IBM SmartCloud Entry for IBM Flex System",
	"5765SKC": "IBM SmartCloud Entry for Power",
	"5765SKP": "SmartCloud Entry Pwr bundle",
	"5765SLE": "PowerVP Standard Edition",
	"5765SS7": "Platform Symphony Std PWR V7",
	"5765SSE": "Platform Symphony - Standard Edition for Power",
	"5765TAI": "IBM Visual Insights",
	"5765UAV": "IUAM Virtualization Edition",
	"5765VC2": "PowerVC for Private Cloud",
	"5765VCB": "PowerVC Base",
	"5765VCC": "IBM Cloud PowerVC Manager",
	"5765VCD": "Cloud PowerVC Mgr for SDI",
	"5765VCS": "PowerVC Standard Edition",
	"5765VCX": "PowerVC Express Edition",
	"5765VE3": "PowerVM Enterprise Edition V3",
	"5765VHP": "vHMC Version - Power based",
	"5765VHX": "vHMC Version - x86 based",
	"5765VL3": "IBM PowerVM Linux Edition",
	"5765VPL": "IBM Systems Director VMControl for PowerLinux",
	"5765VRM": "VM Recovery Manager HA",
	"5765VS2": "PowerVC",
	"5765VS3": "IBM PowerVM Standard Edition",
	"5765WP7": "AIX 5.3 Workload Partitions for AIX 7",
	"5765WPM": "Workload Partitions Mgr",
	"5765XA3": "General Parallel File System on x86 Architecture",
	"5765Z11": "Virtual Appliance x86 Sub",
	"5765Z12": "HMC Vir Appl Pwr Sub",
	"5770AF1": "Advanced Function Printing Utilities",
	"5770AMT": "Rational Application Management Toolset for i",
	"5770BR1": "Backup Recovery and Media Services for i",
	"5770DBM": "Db2 Mirror for i",
	"5770DE1": "DB2 Extenders Version 9.5 for i",
	"5770DFH": "CICS Transaction Server for i",
	"5770HAS": "PowerHA SystemMirror for i",
	"5770JS1": "Advanced Job Scheduler for i",
	"5770MG1": "Managed System Services for i",
	"5770NLV": "National Languages for i",
	"5770PT1": "Performance Tools for i",
	"5770QU1": "Query for i",
	"5770SM1": "System Manager for i",
	"5770SS1": "IBM i",
	"5770ST1": "DB2 Query Manager and SQL Development Kit for i",
	"5770WDS": "Rational Development Studio for i",
	"5770XW1": "IBM i Access Family",
	"5798FAX": "Facsimile Support for iSeries",
	"5799GTJ": "Websphere Development Studio for iSeries, Single Host",
	"5799OC2": "Red Hat Openshift Enablement for Power",
	"5799RP2": "Red Hat Enablement for Power Private Cloud",
	"5799WP7": "AIX 5.3 Workload Partitions for AIX 7",
}

var NamesFromPrdidAndFeature = map[string]map[string]string{
	"5621BGA": {
		"5050": "SMART BUSINESS SW PACK FOR I - IBM Smart Business SW Pack for i",
	},
	"5648F82": {
		"BASE": "PurifyPlus for AIX",
	},
	"5722AC3": {
		"5050": "CRYPTOGRAPH ACC PROV 128-BIT - Cryptograhic Access Provider",
	},
	"5722AF1": {
		"5050": "AFP Utilities for iSeries - AFP Utilities",
	},
	"5722AP1": {
		"5050": "Advanced DBCS Printer Support for iSeries - Advanced DBCS Printer Support",
		"5101": "Advanced DBCS Printer Support for iSeries - IPDS",
	},
	"5722BR1": {
		"5050": "Backup Recovery and Media Services for iSeries - BRMS",
		"5101": "Backup Recovery and Media Services for iSeries - BRMS Network Feature",
		"5102": "Backup Recovery and Media Services for iSeries - BRMS Advanced Functions",
	},
	"5722CE3": {
		"5050": "Client Encryption 128-bit - Client Encryption",
	},
	"5722CM1": {
		"5050": "Communications Utilities for i5/OS - Communication Utilities",
	},
	"5722CR1": {
		"5050": "IBM Cryptographic Support for iSeries - Cryptographic Support",
	},
	"5722DB1": {
		"5050": "System/38 Utilities - S/38 Utilities",
	},
	"5722DE1": {
		"5050": "DB2 Universal DB Extenders - DB2 UDB Extenders",
	},
	"5722DFH": {
		"5050": "CICS Transaction Server for i5/OS - CICS",
	},
	"5722DP4": {
		"5050": "DataPropagator for iSeries - DB2 DataPropagator",
	},
	"5722DR1": {
		"5050": "Director for i5/OS",
	},
	"5722DS1": {
		"5050": "Business Graphics Utility for AS/400 - Business Graphics Utility",
	},
	"5722IMI": {
		"5050": "SYS DIR VIRTUAL IMAGE MGT - Sys Dir Virtual Image Mgt",
	},
	"5722IP1": {
		"5050": "Infoprint Server for iSeries - InfoPrint Server for iSeries",
		"5101": "Infoprint Server for iSeries - PS to AFP Transform",
	},
	"5722JS1": {
		"5050": "Job Scheduler for iSeries - Advanced Job Scheduler",
	},
	"5722MG1": {
		"5050": "Managed System Services for i5/OS - Managed System",
	},
	"5722NLV": {
		"5050": "National Languages for i5/OS - Additional Primary NLV",
	},
	"5722PD1": {
		"5050": "APD",
	},
	"5722PT1": {
		"5050": "Performance Tools for iSeries - Performance Tools",
		"5101": "Performance Tools for iSeries - Performance Tools Manager",
		"5102": "Performance Tools for iSeries - Performance Tools Agent",
	},
	"5722QU1": {
		"5050": "Query for iSeries - Query",
	},
	"5722SM1": {
		"5050": "System Manager",
	},
	"5722SS1": {
		"5050": "i5/OS",
		"5051": "i5/OS - i5/OS,IBM i - Per Processor",
		"5052": "i5/OS - i5/OS,IBM i - Users",
		"5053": "i5/OS - i5/OS,IBM i - Appl Server Processor",
		"5103": "i5/OS - Media & Storage Extensions",
		"5112": "i5/OS - PSF/400 1-45 IPM Printer Support",
		"5113": "i5/OS - PSF/400 1-100 IPM Printer Support",
		"5114": "i5/OS - PSF/400 Anyspeed Printer Support",
		"5116": "i5/OS - HA Switchable Resources",
		"5117": "i5/OS - HA Journal Performance",
	},
	"5722ST1": {
		"5050": "DB2 and SQL Dev Kit",
	},
	"5722SWD": {
		"5050": "Director Premium Edition",
	},
	"5722WDS": {
		"5050": "WDS/400 - WebSphere Development Studio",
	},
	"5722WE1": {
		"5050": "Web Enablement",
	},
	"5722WE2": {
		"1001": "Web Enablement for i5/OS - Web Enablement for i5/OS WAS 6.0",
		"1002": "Web Enablement for i5/OS - Web Enablement for i5/OS WAS 5.1",
		"1003": "Web Enablement for i5/OS - Web Enablement for i5/OS WAS 6.1",
		"1004": "Web Enablement for i5/OS - Web Enablement for i5/OS WAS 1.1",
		"1005": "Web Enablement for i5/OS - Web Enablement for i5/OS WAS 1.2",
		"1006": "Web Enablement for i5/OS - Web Enablement for i5/OS WAS 7.0",
		"1007": "Web Enablement for i5/OS - WebSphere Express 8.0/8.5/8.5.5/8.5.5.15",
		"5050": "Web Enablement for i5/OS - Web Enablement i5/OS",
	},
	"5722XW1": {
		"5101": "iSeries Access - iSeries Access Family Per User",
	},
	"5724I08": {
		"BASE": "XL Fortran Enterprise Ed V9",
	},
	"5724I10": {
		"BASE": "XL C Enterprise Ed (AIX) V7",
	},
	"5724I11": {
		"BASE": "XL C/C++ Enterprise Ed V7",
	},
	"5724I15": {
		"BASE": "Tivoli Provisioning Manager",
	},
	"5724K76": {
		"BASE": "XL Fortran Adv Ed Linux V9",
	},
	"5724K77": {
		"BASE": "XL C/C++ Adv Ed for Linux - XL C/C++ Adv Ed for Linux V7",
	},
	"5724L09": {
		"BASE": "Tivoli Provisioning Manager",
	},
	"5724M11": {
		"BASE": "XL C Enterprise Ed V8 AIX",
	},
	"5724M12": {
		"BASE": "XL C/C++ Enterprise Ed V8",
	},
	"5724M13": {
		"BASE": "XL Fortran Enterprise Ed V10",
	},
	"5724M16": {
		"BASE": "XL C/C++ Adv Ed for Linux V8",
	},
	"5724M17": {
		"BASE": "XL Fortran Adv Ed Linux V10",
	},
	"5724O33": {
		"BASE": "IBM Tivoli Usage Acct Mgr - Tivoli Usage Accounting Manager",
	},
	"5724PL1": {
		"BASE": "IBM PL/I FOR AIX",
	},
	"5724S70": {
		"BASE": "XL C EE for AIX V9.0 - IBM XL C Ent Edition for AIX V9.0",
	},
	"5724S71": {
		"BASE": "XL C/C++ EE for AIX V9.0 - XL C/C++ Ent Ed for AIX V9.0",
	},
	"5724S72": {
		"BASE": "XL Fortran EE for AIX V11.1 - XL Fortran Ent Ed for AIX V11.1",
	},
	"5724S73": {
		"BASE": "XL C/C++ AE for LINUX V9.0 - XL C/C++ Adv Ed for LINUX V9.0",
	},
	"5724S74": {
		"BASE": "XL Fort AE for LINUX V11.1 - XL Fortran Adv Ed for LINUX V11.1",
	},
	"5724U80": {
		"BASE": "IBM XL C for AIX",
	},
	"5724U81": {
		"BASE": "XL C/C++ for AIX - IBM XL C/C++ for AIX",
	},
	"5724U82": {
		"BASE": "XL Fortran for AIX",
	},
	"5724U83": {
		"BASE": "XL C/C++ for Linux V10.1 - IBM XL C/C++ for Linux V10.1",
	},
	"5724U84": {
		"BASE": "XL Fortran for Linux V12.1 - IBM XL Fortran for Linux V12.1",
	},
	"5724V62": {
		"BASE": "IBM COBOL for AIX V3.1",
	},
	"5724X12": {
		"BASE": "XL C for AIX",
	},
	"5724X13": {
		"BASE": "XL C/C++ for AIX",
	},
	"5724X14": {
		"BASE": "XL C/C++ for Linux V11.1 - IBM XL C/C++ for Linux V11.1",
	},
	"5724X15": {
		"BASE": "XL Fortran for AIX",
	},
	"5724X16": {
		"BASE": "XL Fortran for Linux V13.1",
	},
	"5724Z00": {
		"0002": "Rational Team Concert Expr. - RTC Express Contrib Client",
		"0003": "Rational Team Concert Expr. - RTC Express Developer Client",
		"BASE": "Rational Team Concert Expr. - Rational Team Concert Expr",
	},
	"5724Z01": {
		"0002": "Rational Team Concert Std. - RTC Standard Contrib Client",
		"0003": "Rational Team Concert Std. - RTC Std Edn Developer Client",
		"BASE": "Rational Team Concert Std. - Rational Team Concert Std",
	},
	"5724Z87": {
		"BASE": "IBM COBOL for AIX V4.1",
	},
	"5733AC1": {
		"1001": "ARCAD RPG Converter - ARCAD RPG Converter for IBM i",
	},
	"5733AO1": {
		"1001": "ARCAD Observer - ARCAD Observer for IBM i",
	},
	"5733ARE": {
		"5050": "IBM Application Runtime Expert for i - Application Runtime Expert",
	},
	"5733B45": {
		"5050": "IBM AFP Font Collection for i - AFP Outline Fonts",
	},
	"5733CO2": {
		"5050": "Connect for iSeries - Connect",
	},
	"5733CY1": {
		"5050": "CRYPTOGRAPHIC DEVICE MANAGER - Cryptograhic Device Manager",
	},
	"5733CY2": {
		"5050": "CRYPTOGRAPHIC DEVICE MANAGER - System i Cryptograhic Device Manager",
	},
	"5733CY3": {
		"5050": "Cryptographic Device Manager - Cryptograhic Device Manager",
	},
	"5733DIR": {
		"5050": "Director with VE Console",
	},
	"5733DR1": {
		"5050": "IBM Director with Virtualization Engine Console for i5/OS - Director w/VE Console i5/OS",
	},
	"5733EWA": {
		"5050": "VE Ent Wrkload Mgr / i5/OS",
	},
	"5733EWM": {
		"5050": "VE Enterprise Workload Mgr",
	},
	"5733FXD": {
		"5050": "Integrated Domino Fax for iSeries - Domino Fax",
	},
	"5733GT1": {
		"5050": "IBM GRID TOOLBOX - Grid Toolbox",
	},
	"5733HAA": {
		"5050": "HA ASSIST FOR I - IBM DataMirror iCluster",
	},
	"5733ICC": {
		"5101": "Cloud Storage Solutions for i",
		"5102": "Cloud Storage Solutions for i - Cloud Storage Solutions advanced",
	},
	"5733ICL": {
		"5050": "DATAMIRROR ICLUSTER - IBM DataMirror iCluster",
		"5051": "DATAMIRROR ICLUSTER - IBM DataMirror iBalance",
	},
	"5733ICS": {
		"5050": "DATAMIRROR ICLUSTER SMB - IBM DataMirror iCluster SMB",
	},
	"5733ID1": {
		"5050": "Infoprint Designer/400 - InfoPrint Designer for iSeries",
	},
	"5733OAR": {
		"5050": "IBM Rational Open Access RPG Edition",
	},
	"5733OMF": {
		"5050": "Find Text Search Server for DB2 for i5/OS - OmniFind Text Search Server for DB2",
	},
	"5733PS1": {
		"5050": "SECURE PERSPECTIVE - Secure Perspective",
	},
	"5733QU2": {
		"5050": "DB2 Web Query for System i - IBM DB2 Web Query for System i",
		"5101": "DB2 Web Query for System i - Active Reports",
		"5102": "DB2 Web Query for System i - OLAP Enablement",
		"5104": "DB2 Web Query for System i - Run Time User Enablement",
		"5105": "DB2 Web Query for System i - Spreadsheet Client",
		"5106": "DB2 Web Query for System i - SQL Server Adapter",
		"5107": "DB2 Web Query for System i - DB2 Web Adpt for JD Edwards",
	},
	"5733QU3": {
		"5050": "DB2 Web Query Report Broker - IBM DB2 Web Query Report Broker",
	},
	"5733QU4": {
		"5050": "DB2 Web Query Software Development Kit - IBM DB2 Web Query SW Dev Kit",
	},
	"5733RDA": {
		"BASE": "RD Power C/C++Development Tools for AIX - RDPower C / C++ Dev Tool AIX",
	},
	"5733RDB": {
		"BASE": "RD Power COBOL Development Tools for AIX - RDPower Cobol Dev Tools AIX",
	},
	"5733RDC": {
		"BASE": "RD Power C/C++ Development Tools for Linux - RDPower C/C++ Dev for Linux",
	},
	"5733RDD": {
		"5050": "RD Power Power Tools for i - RDPower Tools for i",
	},
	"5733RDE": {
		"BASE": "RD Power Power Tools for AIX - RDPower Power Tools for AIX",
	},
	"5733RDG": {
		"5050": "RD Power RPG and COBOL Development Tools for i - RDPower RPG COBOL Dev for i",
	},
	"5733RDH": {
		"BASE": "RD Power C/C++Development Tools for AIX - RDPower C/C++ Dev for AIX",
	},
	"5733RDI": {
		"5050": "Rational Developer for System i - Rational Developer for System i V7.1",
	},
	"5733RDJ": {
		"BASE": "RD Power COBOL Development Tools for AIX - RDPower COBOL Dev for AIX",
	},
	"5733RDP": {
		"5050": "RD Power RPG and COBOL Development Tools for i - Rational Developer for Power Systems SW",
	},
	"5733RDS": {
		"5050": "VE Resource Dependency Svc",
	},
	"5733RDW": {
		"5101": "Rational Developer for i - Rational Developer RPG&COBOL Tools V9",
		"5102": "Rational Developer for i - Rational Developer RPG&COBOL Java Ed V9",
		"5103": "Rational Developer for i - Rational Developer RPG&COBOL EGL Ed V9",
	},
	"5733RDX": {
		"0002": "Rational Developer for AIX and Linux - Rational Dev for AIX C/C++ Edition",
		"BASE": "Rational Developer for AIX and Linux - Rational Dev for AIX-Linux AIX COBOL",
	},
	"5733RTC": {
		"0002": "IBM Rational Team Concert - IBM RTC Dev for Ent Platform",
		"0003": "IBM Rational Team Concert - IBM RTC Dev for Workgroups",
		"0004": "IBM Rational Team Concert - IBM RTC Developer",
		"0005": "IBM Rational Team Concert - IBM RTC Stakeholder",
		"BASE": "IBM Rational Team Concert - IBM RTC Contributor",
	},
	"5733SC1": {
		"5050": "PORTABLE UTILITIES FOR i5/OS - Portable Utilities for i5/OS",
	},
	"5733SOA": {
		"5050": "IBM Rational Developer for System i for SOA Construction - Rational Developer on System i for SOA",
	},
	"5733SOC": {
		"5050": "Rational Developer for i for SOA Construction - IBM Rational Developer for i for SOA Con",
	},
	"5733VE1": {
		"5050": "VIRTUALIZATION ENGINE - IBM Director Multiplatform",
	},
	"5733WE3": {
		"1001": "Web Enablement for i",
	},
	"5733WQX": {
		"5050": "IBM Db2 Web Query for i - Web Query for System i Base",
		"5101": "IBM Db2 Web Query for i - Web Query for System i Express",
		"5102": "IBM Db2 Web Query for i - Web Query for System i Standard",
		"5103": "IBM Db2 Web Query for i - Web Query for System i Enterprise",
		"5104": "IBM Db2 Web Query for i - Web Query User Registration",
		"5105": "IBM Db2 Web Query for i - Development Workbench User Registration",
		"5106": "IBM Db2 Web Query for i - Runtime Group Profile Registration",
		"5107": "IBM Db2 Web Query for i - JDE Adapter Registration",
		"5108": "IBM Db2 Web Query for i - IBM Web Query DataMigrator for i",
		"5109": "IBM Db2 Web Query for i - Web Query Scheduler Edition",
		"5110": "IBM Db2 Web Query for i - Web Query Runtime User Edition",
	},
	"5733WSE": {
		"5050": "Workplace Services Express",
	},
	"5733XT1": {
		"5050": "XML Toolkit for iSeries - XML Toolkit",
	},
	"5733XT2": {
		"5050": "XML Toolkit for IBM System i5 - XML Toolkit",
	},
	"5761AF1": {
		"5050": "Advanced Function Printing Utilities - AFP Utilities",
	},
	"5761AMT": {
		"5050": "IBM Rational Application Management Tool Set for i - Rational Appl Mgmt Tool Set for i",
	},
	"5761AP1": {
		"5050": "Advanced DBCS Printer Support for iSeries - Advanced DBCS Printer Support",
		"5101": "Advanced DBCS Printer Support for iSeries - IPDS",
	},
	"5761BR1": {
		"5050": "Backup Recovery and Media Services for i5/OS - BRMS",
		"5101": "Backup Recovery and Media Services for i5/OS - BRMS Network Feature",
		"5102": "Backup Recovery and Media Services for i5/OS - BRMS Advanced Functions",
	},
	"5761CM1": {
		"5050": "Communications Utilities for System i - Communication Utilities",
	},
	"5761DB1": {
		"5050": "System/38 Utilities for System i - S/38 Utilities",
	},
	"5761DE1": {
		"5050": "DB2 Extenders Version 9.1 for i5/OS - DB2 Extenders V9.1 for iSeries",
	},
	"5761DFH": {
		"5050": "CICS Transaction Server for i5/OS - CICS",
	},
	"5761DP4": {
		"5050": "DB2 DataPropagator for iSeries V8.1",
	},
	"5761DS2": {
		"5050": "Business Graphics Utility for System i - Business Graphics Utility",
	},
	"5761HAS": {
		"5050": "High Availability Solutions Manager - System i High Avail Solutions Manager",
	},
	"5761JS1": {
		"5050": "Advanced Job Scheduler for i5/OS - Advanced Job Scheduler",
	},
	"5761MG1": {
		"5050": "Managed System Services for i5/OS - Managed System",
	},
	"5761NLV": {
		"5050": "National Languages for i5/OS - Additional Primary NLV",
	},
	"5761PT1": {
		"5050": "Performance Tools for i5/OS - Performance Tools",
		"5101": "Performance Tools for i5/OS - Performance Tools Manager",
		"5102": "Performance Tools for i5/OS - Performance Tools Agent",
		"5103": "Performance Tools for i5/OS - Performance Tools Job Watcher",
	},
	"5761QU1": {
		"5050": "Query for i5/OS - Query",
	},
	"5761SM1": {
		"5050": "System Manager for i5/OS - System Manager",
	},
	"5761SS1": {
		"5050": "IBM i",
		"5051": "IBM i - IBM i - Per Processor",
		"5052": "IBM i - IBM i Users",
		"5053": "IBM i - IBM i Appl Server Processor",
		"5103": "IBM i - Media & Storage Extensions",
		"5112": "IBM i - Print 1-55 Impressions/Minute",
		"5113": "IBM i - Print 1-100 Impressions/Minute",
		"5114": "IBM i - Print Anyspeed Impressions/Minute",
		"5116": "IBM i - HA Switchable Resources",
		"5117": "IBM i - HA Journal Performance",
	},
	"5761ST1": {
		"5050": "DB2 Query Manager and SQL Development Kit for i5/OS - DB2 and SQL Dev Kit",
	},
	"5761WDS": {
		"5101": "WebSphere Development Studio for System i - ILE Compiler",
		"5102": "WebSphere Development Studio for System i - Heritage Compiler",
		"5103": "WebSphere Development Studio for System i - Application Development ToolSet",
	},
	"5761XW1": {
		"5101": "System i Access Family - iSeries Access Family Per User",
	},
	"57652B1": {
		"0012": "AIX 7 Standard Edition Sub - AIX 7 Standard Edition Sub`1YR",
	},
	"57652C1": {
		"0012": "Pvt Cloud Ed with AIX 7 Sub - IBM Private Cloud Edition with AIX 7",
	},
	"57652E1": {
		"0012": "AIX 7 Enterprise Edition Sub - AIX 7 Enterprise Edition Sub 1YR",
	},
	"57656C1": {
		"BASE": "Private Cloud Edn Subs - Private Cloud Edition Subscription",
	},
	"5765AEP": {
		"BASE": "IBM Systems Director Active Energy Manager for Power Systems - System Dir AEM Power V4",
	},
	"5765AEZ": {
		"BASE": "AIX Enterprise Edition - AIX Enterprise Edition V6",
	},
	"5765AIE": {
		"BASE": "IBM PowerAI Enterprise",
	},
	"5765AME": {
		"BASE": "AIX Management Edition",
	},
	"5765AMT": {
		"BASE": "AIX7.2 Monthly term offering",
	},
	"5765ASC": {
		"BASE": "Platform Appl Svc Cntrl Pwr",
	},
	"5765ASM": {
		"BASE": "IBM Systems Director Storage Control - IBM Sys Dir Storage Control",
	},
	"5765AVE": {
		"BASE": "PowerVM Lx86",
	},
	"5765C41": {
		"BASE": "Parallel ESSL for AIX - PESSL V2.3 for AIX",
	},
	"5765C42": {
		"BASE": "ESSL for AIX - ESSL FOR AIX V3",
	},
	"5765CBA": {
		"BASE": "Private Cloud Edn with AIX 7 - IBM Cloud Edition with AIX V1",
	},
	"5765CD1": {
		"BASE": "AIX Enterprise Edition 7.1 - AIX V7.1 Enterprise Edition",
	},
	"5765CD3": {
		"BASE": "AIX Enterprise Edition 7.2 and 7.3 - AIX V7.2 Enterprise Edition",
	},
	"5765CMO": {
		"BASE": "Cloud Mgr w/OpenStack 4.2 - Cloud Manager with Openstack V4.2",
	},
	"5765CMT": {
		"003M": "CMC Monthly Term offering - IBM Cloud Mgmt Console Monthly 3 month",
		"006M": "CMC Monthly Term offering - IBM Cloud Mgmt Console Monthly 6 month",
		"012M": "CMC Monthly Term offering - IBM Cloud Mgmt Console Monthly 12 month",
		"036M": "CMC Monthly Term offering - IBM Cloud Mgmt Console Monthly 36 month",
		"048M": "CMC Monthly Term offering - IBM Cloud Mgmt Console Monthly 48 month",
		"060M": "CMC Monthly Term offering - IBM Cloud Mgmt Console Monthly 60 month",
	},
	"5765COB": {
		"BASE": "IBM COBOL for AIX V5 - IBM COBOL for AIX V5.1",
		"OPT1": "IBM COBOL for AIX V5 - COBOL Runtime for AIX",
	},
	"5765D51": {
		"BASE": "PSSP for AIX V3 - PSSP for AIX",
	},
	"5765D93": {
		"BASE": "Parallel Environment V3",
	},
	"5765DR1": {
		"0002": "IBM Director for pSystems - IBM Director Premium Edition FC",
		"BASE": "IBM Director for pSystems",
	},
	"5765DRG": {
		"BASE": "VM Recovery Manager DR - Geographically Dispersed Resiliency",
	},
	"5765DRP": {
		"BASE": "IBM Systems Director for Power Systems - Systems Director for Power V6",
	},
	"5765E61": {
		"BASE": "AIX 5L for POWER V5.1",
	},
	"5765E62": {
		"BASE": "AIX 5L for POWER V5.2",
	},
	"5765E68": {
		"BASE": "Performance Aide V3",
	},
	"5765E69": {
		"BASE": "Loadleveler - Loadleveler V3",
	},
	"5765E72": {
		"BASE": "AIX Fast Connect for POWER - Fast Connect V3.1 for AIX",
	},
	"5765E74": {
		"BASE": "AIX Performance Toolbox - Performance Toolbox V3",
	},
	"5765E82": {
		"0002": "HAGEO/GeoRM V2 - HAGEO V2",
		"BASE": "HAGEO/GeoRM V2 - GeoRM V2",
	},
	"5765E85": {
		"BASE": "AIXlink/X.25 - AIX Link/X.25 V2",
	},
	"5765EAP": {
		"BASE": "ESSL for AIX on POWER - ESSL for AIX on Power",
	},
	"5765ECB": {
		"BASE": "IBM Private Cloud Edition - IBM Private Cloud Edition V1",
	},
	"5765EEP": {
		"BASE": "IBM Systems Director Enterprise Edition for Power - System Dir Enterprise Ed V6.2",
	},
	"5765EGO": {
		"BASE": "Platform Resource Scheduler - Platform Resource Scheduler for Power V2",
	},
	"5765EL5": {
		"BASE": "Parallel ESSL for Linux on Power - Parallel ESSL for Linux on Power V5.1",
	},
	"5765EMP": {
		"BASE": "IBM Systems Director VMControl Enterprise Edition - IBM Sys Dir VMCntl System Pools",
	},
	"5765EPL": {
		"BASE": "IBM Systems Director for PowerLinux - Sys Dir for PowerLinux",
	},
	"5765ESL": {
		"BASE": "Parallel Engineering and Scientific Subroutine Library - Parallel ESSL for Linux POWER V4",
	},
	"5765ESX": {
		"BASE": "Parallel Engineering and Scientific Subroutine Library - Parallel ESSL for AIX V4.2",
	},
	"5765EXP": {
		"BASE": "IBM Systems Director Express Edition for Power - Systems Dir Express Ed V6.2",
	},
	"5765F56": {
		"BASE": "VisualAge C++ Pro for AIX V6",
	},
	"5765F57": {
		"BASE": "C for AIX V6.0",
	},
	"5765F62": {
		"0002": "PowerHA V5 - PowerHA V5 Extended Distance",
		"0003": "PowerHA V5 - PowerHA V5 Smart Assist",
		"BASE": "PowerHA V5",
	},
	"5765F64": {
		"BASE": "GPFS for AIX 5L V2",
	},
	"5765F67": {
		"0002": "Cluster Systems Management - CSM HAMS option",
		"BASE": "Cluster Systems Management - CSM for AIX",
	},
	"5765F68": {
		"BASE": "IBM Infoprint Manager for AIX - Infoprint Manager/AIX V4",
	},
	"5765F70": {
		"BASE": "XL Fortran AIX - XL Fortran Compiler for AIX V8.1",
	},
	"5765F71": {
		"BASE": "XL Fortran RTE for Linux - XL Fortran RTE for Linux 8.1",
	},
	"5765F82": {
		"BASE": "ESSL FOR AIX, V4 - ESSL FOR AIX V4",
	},
	"5765F83": {
		"BASE": "PARALLEL ENVIRONMENT V4 - Parallel Environment V4",
	},
	"5765F84": {
		"BASE": "PARALLEL ESSL FOR AIX, V3 - Parallel ESSL for AIX V3",
	},
	"5765FS2": {
		"BASE": "SAN File System Software V2",
	},
	"5765G03": {
		"BASE": "AIX 5L for POWER V5.3",
	},
	"5765G16": {
		"0002": "Cluster Systems Management for Linux on POWER - CSM HAMS option",
		"BASE": "Cluster Systems Management for Linux on POWER - CSM for Linux",
	},
	"5765G17": {
		"BASE": "ESSL for Linux on POWER - ESSL for Linux V4",
	},
	"5765G18": {
		"BASE": "Parallel ESSL for Linux",
	},
	"5765G20": {
		"BASE": "GPFS for Linux on pSeries",
	},
	"5765G31": {
		"BASE": "Partition Load Manager",
	},
	"5765G34": {
		"BASE": "Virtual I/O Server",
	},
	"5765G62": {
		"BASE": "AIX Standard Edition - AIX Standard Edition V6",
	},
	"5765G66": {
		"0002": "General Parallel File System for POWER - GPFS for POWER Client",
		"0003": "General Parallel File System for POWER - GPFS for POWER Server",
		"BASE": "General Parallel File System for POWER - GPFS File Placement Optimizer",
	},
	"5765G67": {
		"BASE": "GPFS for Linux on Power",
	},
	"5765G70": {
		"BASE": "PE for Linux on Power",
	},
	"5765G71": {
		"BASE": "HACMP for Linux, V5 - HACMP V5.4 for Linux on POWER",
	},
	"5765G82": {
		"BASE": "IBM PowerSC Express Edition - PowerSC Express Edition V1.1",
	},
	"5765G83": {
		"BASE": "Workload Partitions Manager - WPAR Manager for AIX V2.2",
	},
	"5765G90": {
		"BASE": "AIX Express Edition - AIX Express Edition V6",
	},
	"5765G97": {
		"BASE": "AIX Express Edition 7.1",
	},
	"5765G98": {
		"BASE": "AIX Standard Edition 7.1, 7.2 and 7.3 - AIX Standard Edition V7",
	},
	"5765G99": {
		"BASE": "AIX Enterprise Edition - AIX Enterprise Edition V7",
	},
	"5765H23": {
		"BASE": "PowerHA for AIX Standard Edition V6 - PowerHA Standard Edition V6",
	},
	"5765H24": {
		"BASE": "PowerHA for AIX Enterprise Edition V6 - PowerHA Enterprise Edition V6",
	},
	"5765H25": {
		"BASE": "Engineering and Scientific Subroutine Library for AIX - ESSL for AIX V5",
	},
	"5765H37": {
		"BASE": "PowerHA SystemMirror Enterprise Edition V7 - PowerHA SystemMirror Ent Ed V7",
	},
	"5765H38": {
		"BASE": "AIX 5.2 Workload Partitions for AIX 7",
	},
	"5765H39": {
		"BASE": "PowerHA for AIX Standard Edition V7 - PowerHA SystemMirror Std Ed V7",
	},
	"5765H43": {
		"BASE": "PHA SystemMirror Std ed subs",
	},
	"5765H44": {
		"BASE": "PowerHA SM Ent ed subs",
	},
	"5765H7E": {
		"003M": "PowerHA SystemMirror Enterprise Edition Monthly Term V7.2 - PowerHA SysMirror Ent Ed Monthly 3 month",
		"006M": "PowerHA SystemMirror Enterprise Edition Monthly Term V7.2 - PowerHA SysMirror Ent Ed Monthly 6 month",
		"012M": "PowerHA SystemMirror Enterprise Edition Monthly Term V7.2 - PowerHA SysMirror Ent Ed Monthly 12month",
		"036M": "PowerHA SystemMirror Enterprise Edition Monthly Term V7.2 - PowerHA SysMirror Ent Ed Monthly 36month",
	},
	"5765H7S": {
		"003M": "PowerHA SystemMirror Standard Edition Monthly Term V7.2 - PowerHA SysMir Std Ed Monthly 3 month",
		"006M": "PowerHA SystemMirror Standard Edition Monthly Term V7.2 - PowerHA SysMir Std Ed Monthly 6 month",
		"012M": "PowerHA SystemMirror Standard Edition Monthly Term V7.2 - PowerHA SysMir Std Ed Monthly 12 month",
		"036M": "PowerHA SystemMirror Standard Edition Monthly Term V7.2 - PowerHA SysMir Std Ed Monthly 36 month",
	},
	"5765HMA": {
		"BASE": "Power HMC Virtual Software Appliance for Power HW - IBM Cloud PowerVC Manager",
	},
	"5765HMB": {
		"BASE": "HMC Virtual Appliance Power",
	},
	"5765HMV": {
		"BASE": "Power HMC Virtual Software Appliance for x86 HW - Power HMC Virtual Software Appliance",
	},
	"5765HMW": {
		"BASE": "HMC Virtual Appliance x86 - IBM HMC Virtual Appliance x86",
	},
	"5765I2P": {
		"BASE": "IBM Systems Director VMControl Standard Edition - Virtual Image Manager V2.1",
	},
	"5765IMP": {
		"BASE": "IBM Systems Director Virtual Image Management - Systems Director Virtual Image Manager",
	},
	"5765ISS": {
		"BASE": "ISS Service Sensor V7",
	},
	"5765ITM": {
		"BASE": "ITM SE for System p - Tivoli Monitoring V6.1 for System P",
	},
	"5765J01": {
		"BASE": "IBM XL C for AIX - IBM XL C for AIX V12.1",
	},
	"5765J02": {
		"BASE": "IBM XL C/C++ for AIX - IBM XL C/C++ for AIX V12.1",
	},
	"5765J03": {
		"BASE": "IBM XL C/C++ for Linux - IBM XL C/C++ for Linux V12.1",
	},
	"5765J04": {
		"BASE": "IBM XL Fortran for AIX - IBM XL Fortran for AIX V14.1",
	},
	"5765J05": {
		"BASE": "IBM XL Fortran for Linux - IBM XL Fortran for Linux V14.1",
	},
	"5765J06": {
		"BASE": "IBM XL C for AIX",
	},
	"5765J07": {
		"BASE": "IBM XL C/C++ for AIX",
	},
	"5765J08": {
		"BASE": "IBM XL C/C++ for Linux",
	},
	"5765J09": {
		"BASE": "IBM XL Fortran for AIX",
	},
	"5765J10": {
		"BASE": "IBM XL Fortran for Linux",
	},
	"5765J11": {
		"BASE": "IBM XL C for AIX V16 - IBM XL C for AIX",
	},
	"5765J12": {
		"BASE": "IBM XL C/C++ for AIX V16 - IBM XL C/C++ for AIX",
	},
	"5765J13": {
		"BASE": "IBM XL C/C++ for Linux V16 - IBM XL C/C++ for Linux",
	},
	"5765J14": {
		"BASE": "IBM XL Fortran for AIX V16 - IBM XL Fortran for AIX",
	},
	"5765J15": {
		"BASE": "IBM XL Fortran for Linux V16 - IBM XL Fortran for Linux",
	},
	"5765J16": {
		"BASE": "XL C/C++ Mo Term",
	},
	"5765J17": {
		"BASE": "Fortran for AIX Mo Term",
	},
	"5765J18": {
		"BASE": "IBM Open XL C/C++ for AIX",
	},
	"5765J19": {
		"BASE": "IBM Open XL Fortran for AIX",
	},
	"5765J20": {
		"BASE": "Open XL C/C++ for Lnx on P - IBM Open XL C/C++ for Linux on Power",
	},
	"5765J21": {
		"BASE": "Open XL Fortran for Lnx on P - IBM OpenXLFortran for Linux",
	},
	"5765J22": {
		"BASE": "Open XL C/C++for Lnx Mo Term - Open XL C/C++ for LinuxMonthly Term",
	},
	"5765J23": {
		"BASE": "XL Fortran for Lnx Mo Term - Open XL Fortran for Linux17Monthly Term",
	},
	"5765K01": {
		"BASE": "DB2 WSE 9.7 for HPC - DB2 Workgroup Server Edition 9.7 for HPC",
	},
	"5765KV3": {
		"BASE": "IBM PowerKVM V3",
	},
	"5765KVM": {
		"BASE": "IBM PowerKVM",
	},
	"5765L22": {
		"BASE": "IBM PowerHA for Linux - IBM PowerHA SystemMirror for Linux",
	},
	"5765L40": {
		"BASE": "TWS LoadLeveler for AIX, V4 - TWS LoadLeveler for AIX V4",
	},
	"5765L50": {
		"BASE": "Tivoli Workload Scheduler LoadLeveler for AIX - TWS LoadLeveler for AIX V5",
	},
	"5765L51": {
		"BASE": "ESSL for Linux on Power - ESSL for Linux V5",
	},
	"5765L61": {
		"BASE": "ESSL for Linux on Power V6 - ESSL for Linux V6",
	},
	"5765LLP": {
		"BASE": "LoadLeveler for Linux on POWER - LoadLeveler for Linux V5.1",
	},
	"5765LSF": {
		"BASE": "IBM Platform LSF - Standard Edition for Power - IBM Platform LSF Std Ed. V9",
	},
	"5765M1M": {
		"003M": "PowerSC MFA Monthly Term - PowerSC MFA Monthly Term 3 month",
		"006M": "PowerSC MFA Monthly Term - PowerSC MFA Monthly Term 6 month",
		"012M": "PowerSC MFA Monthly Term - PowerSC MFA Monthly Term 12 month",
		"036M": "PowerSC MFA Monthly Term - PowerSC MFA Monthly Term 36 month",
	},
	"5765MCH": {
		"BASE": "Systems Director Management Console SDMC Hardware Appliance - SDMC Hardware Appliance",
	},
	"5765MCV": {
		"BASE": "Systems Director Management Console SDMC SVA - Systems Director Management Console",
	},
	"5765MFA": {
		"BASE": "PowerSC MFA",
	},
	"5765NCP": {
		"BASE": "IBM Systems Director Network Control",
	},
	"5765OSF": {
		"BASE": "IBM Cloud Manager with OpenStack for IBM Flex Systems - Cloud Mgr w/OpenStack FLX V4",
	},
	"5765OSP": {
		"BASE": "IBM Cloud Manager with OpenStack for Power - Cloud Mgr w/OpenStack Pwr V4",
	},
	"5765PAC": {
		"BASE": "Platform App Center Power V9 - Platform App Center Std Ed Power V9",
	},
	"5765PAI": {
		"BASE": "IBM PowerAI",
	},
	"5765PCA": {
		"BASE": "IBM Spectrum Cluster Foundation for POWER - Platform Cluster Mgr Adv PWR",
	},
	"5765PCS": {
		"BASE": "IBM Platform Cluster Manager - Standard Edition for POWER - Platform Cluster Mgr Std Ed PWR V4.1",
	},
	"5765PD2": {
		"0002": "Parallel Performance Toolkit for POWER - PE Dev Ed for Linux V2 HPC Toolkit",
		"BASE": "Parallel Performance Toolkit for POWER - PE Dev Ed for Linux V2 Workbench",
	},
	"5765PDP": {
		"0002": "Parallel Environment Developer Edition for Linux on Power - PE Dev Ed for Linux HPC Toolkit",
		"BASE": "Parallel Environment Developer Edition for Linux on Power - PE Dev Ed for Linux Workbench",
	},
	"5765PEA": {
		"BASE": "Parallel Environment for AIX - PE for AIX V5",
	},
	"5765PED": {
		"0002": "Parallel Environment Developer Edition for AIX - PE Dev Ed for AIX Toolkit",
		"BASE": "Parallel Environment Developer Edition for AIX - PE Dev Ed for AIX Workbench",
	},
	"5765PEL": {
		"BASE": "PE for Linux V5",
	},
	"5765PEP": {
		"BASE": "System Dir AEM for POWER - PowerExecutive V3.1",
	},
	"5765PER": {
		"BASE": "Parallel Environment Runtime Edition for AIX - PE Runtime Edition for AIX V1.2",
	},
	"5765PHS": {
		"BASE": "IBM Power HPC Stack",
	},
	"5765PIM": {
		"BASE": "PurePower Integrated Manager - PurePower Integrated Mgr",
	},
	"5765PPM": {
		"BASE": "Platform Process Manager for POWER - Platform Process Manager for POWER V9.1",
	},
	"5765PR2": {
		"BASE": "Parallel Environment Runtime Edition for Linux on Power - PE Runtime Ed for Linux V2",
	},
	"5765PRP": {
		"BASE": "Parallel Environment Runtime Edition for Linux on Power - PE Runtime Ed. Linux Power V1",
	},
	"5765PSA": {
		"BASE": "Platform Symphony - Advanced Edition for Power - Platform Symphony - Adv Ed for PWR V6.1",
	},
	"5765PSE": {
		"BASE": "IBM PowerSC Standard Edition",
	},
	"5765PTS": {
		"BASE": "PowerSC Trusted Surveyor - PowerSC Trusted Surveyor V1.1",
	},
	"5765PVD": {
		"BASE": "PowerVM Enterprise Edition for Small Servers - PowerVM Enterprise Ed for Small Servers",
	},
	"5765PVE": {
		"BASE": "PowerVM Enterprise Edition V2 - PowerVM Enterprise Edition",
	},
	"5765PVL": {
		"BASE": "PowerVM for IBM PowerLinux - PowerVM PowerLinux",
	},
	"5765PVS": {
		"BASE": "PowerVM Standard Edition",
	},
	"5765PVX": {
		"BASE": "PowerVM Express Edition - PowerVM Express",
	},
	"5765PWO": {
		"BASE": "AIX Dynamic System Optimizer",
	},
	"5765R11": {
		"BASE": "IBM Spectrum Computing Suite for High Performance Analytics - IBM Spectrum Computing Suite for HPA",
	},
	"5765R13": {
		"BASE": "ESSL on OpenPower - ESSL on OpenPower V6R2",
	},
	"5765R17": {
		"BASE": "IBM BOA - IBM Bayesian Optimization Accelerator",
	},
	"5765RD2": {
		"BASE": "VM Recovery Manager DR subs",
	},
	"5765RH2": {
		"BASE": "VM Recovery Manager HA subs",
	},
	"5765RTM": {
		"0002": "IBM Platform RTM for Power - Platform RTM Data Collectors for Power",
		"BASE": "IBM Platform RTM for Power - Platform RTM Server for Power",
	},
	"5765S1S": {
		"003M": "IBM PowerSC Standard Edition Monthly Term V1.2 - IBM PowerSC Std Ed Monthly Term 3 month",
		"006M": "IBM PowerSC Standard Edition Monthly Term V1.2 - IBM PowerSC Std Ed Monthly Term 6 month",
		"012M": "IBM PowerSC Standard Edition Monthly Term V1.2 - IBM PowerSC Std Ed Monthly Term 12 month",
		"036M": "IBM PowerSC Standard Edition Monthly Term V1.2 - IBM PowerSC Std Ed Monthly Term 36 month",
	},
	"5765S2S": {
		"012M": "IBM PowerSC 2.1 Subscription - IBM PowerSC Subscription 2.1 term 1Y",
		"024M": "IBM PowerSC 2.1 Subscription - IBM PowerSC Subscription 2.1 term 2Y",
		"036M": "IBM PowerSC 2.1 Subscription - IBM PowerSC Subscription 2.1 term 3Y",
		"048M": "IBM PowerSC 2.1 Subscription - IBM PowerSC Subscription 2.1 term 4Y",
		"060M": "IBM PowerSC 2.1 Subscription - IBM PowerSC Subscription 2.1 term 5y",
	},
	"5765SA7": {
		"BASE": "Platform Symphony Adv PWR V7",
	},
	"5765SB3": {
		"BASE": "IBM SmartCloud Entry for Power bundle V3 - SmartCloud Entry for Power BundlePlus V3",
	},
	"5765SC2": {
		"BASE": "IBM PowerSC V2 - IBM PowerSC",
	},
	"5765SC3": {
		"BASE": "IBM SmartCloud Entry for Power - SmartCloud Entry for Power V3",
	},
	"5765SCP": {
		"BASE": "IBM SmartCloud Entry for Flex system - SmartCloud Entry for Flex V2",
	},
	"5765SEP": {
		"BASE": "IBM Systems Director Standard Edition for Power - System Dir Standard Ed V6.2",
	},
	"5765SF3": {
		"BASE": "IBM SmartCloud Entry for IBM Flex System - SmartCloud Entry for Flex System V3",
	},
	"5765SKC": {
		"BASE": "IBM SmartCloud Entry for Power - SmartCloud Entry for Power V2",
	},
	"5765SKP": {
		"BASE": "SmartCloud Entry Pwr bundle - SmartCloud Entry for Power bundle V2.4",
	},
	"5765SLE": {
		"BASE": "PowerVP Standard Edition - PowerVP Standard Edition V1",
	},
	"5765SS7": {
		"BASE": "Platform Symphony Std PWR V7",
	},
	"5765SSE": {
		"BASE": "Platform Symphony - Standard Edition for Power - Platform Symphony - Std Ed for POWER",
	},
	"5765TAI": {
		"BASE": "IBM Visual Insights - IBM PowerAI Vision",
	},
	"5765UAV": {
		"BASE": "IUAM Virtualization Edition",
	},
	"5765VC2": {
		"BASE": "PowerVC for Private Cloud - IBM PowerVC for Private Cloud",
	},
	"5765VCB": {
		"BASE": "PowerVC Base - PowerVC Base Edition V1.1",
	},
	"5765VCC": {
		"BASE": "IBM Cloud PowerVC Manager",
	},
	"5765VCD": {
		"BASE": "Cloud PowerVC Mgr for SDI",
	},
	"5765VCS": {
		"BASE": "PowerVC Standard Edition - PowerVC Standard Edition V1",
	},
	"5765VCX": {
		"BASE": "PowerVC Express Edition - PowerVC Express Edition V1",
	},
	"5765VE3": {
		"BASE": "PowerVM Enterprise Edition V3",
	},
	"5765VHP": {
		"BASE": "vHMC Version - Power based",
	},
	"5765VHX": {
		"BASE": "vHMC Version - x86 based",
	},
	"5765VL3": {
		"BASE": "IBM PowerVM Linux Edition - PowerVM for Linux V3",
	},
	"5765VPL": {
		"BASE": "IBM Systems Director VMControl for PowerLinux - ISD VMControl for PowerLinux",
	},
	"5765VRM": {
		"BASE": "VM Recovery Manager HA - IBM VM Recovery Manager HA for Power",
	},
	"5765VS2": {
		"BASE": "PowerVC - IBM PowerVC V2",
	},
	"5765VS3": {
		"BASE": "IBM PowerVM Standard Edition - PowerVM Standard edition V3",
	},
	"5765WP7": {
		"BASE": "AIX 5.3 Workload Partitions for AIX 7 - AIX Workload Partitions Mgr",
	},
	"5765WPM": {
		"BASE": "Workload Partitions Mgr - AIX Workload Partitions Mgr",
	},
	"5765XA3": {
		"0002": "General Parallel File System on x86 Architecture - GPFS on IBM X-Architecture Client",
		"0003": "General Parallel File System on x86 Architecture - GPFS on IBM X-Architecture Server",
		"BASE": "General Parallel File System on x86 Architecture - GPFS File Placement Optimizer",
	},
	"5765Z11": {
		"BASE": "Virtual Appliance x86 Sub",
	},
	"5765Z12": {
		"BASE": "HMC Vir Appl Pwr Sub",
	},
	"5770AF1": {
		"5050": "Advanced Function Printing Utilities - AFP Utilities",
	},
	"5770AMT": {
		"5050": "Rational Application Management Toolset for i - Rational Appl Mgmt Tool Set for i",
	},
	"5770BR1": {
		"5050": "Backup Recovery and Media Services for i - BRMS",
		"5101": "Backup Recovery and Media Services for i - BRMS Network Feature",
		"5102": "Backup Recovery and Media Services for i - BRMS Advanced Functions",
	},
	"5770DBM": {
		"5101": "Db2 Mirror for i - IBM Database 2 Mirror for i",
	},
	"5770DE1": {
		"5050": "DB2 Extenders Version 9.5 for i - DB2 Extenders V9.1 for iSeries",
	},
	"5770DFH": {
		"5050": "CICS Transaction Server for i - CICS",
	},
	"5770HAS": {
		"5050": "PowerHA SystemMirror for i - PowerHA SystemMirror i - Standard",
		"5051": "PowerHA SystemMirror for i - PowerHA SystemMirror i (no edition)",
		"5052": "PowerHA SystemMirror for i - PowerHA SystemMirror i - Standard",
		"5053": "PowerHA SystemMirror for i - PowerHA SystemMirror i - Express",
	},
	"5770JS1": {
		"5050": "Advanced Job Scheduler for i - Advanced Job Scheduler",
	},
	"5770MG1": {
		"5050": "Managed System Services for i - Managed System",
	},
	"5770NLV": {
		"5050": "National Languages for i - Additional Primary NLV",
	},
	"5770PT1": {
		"5050": "Performance Tools for i - Performance Tools",
		"5101": "Performance Tools for i - Performance Tools Manager",
		"5102": "Performance Tools for i - Performance Tools Agent",
		"5103": "Performance Tools for i - Performance Tools Job Watcher",
	},
	"5770QU1": {
		"5050": "Query for i - Query",
	},
	"5770SM1": {
		"5050": "System Manager for i - System Manager",
	},
	"5770SS1": {
		"5050": "IBM i",
		"5051": "IBM i - IBM i - Per Processor",
		"5052": "IBM i - IBM i Users",
		"5053": "IBM i - IBM i Appl Server Processor",
		"5103": "IBM i - Media & Storage Extensions",
		"5112": "IBM i - Print 1-55 Impressions/Minute",
		"5113": "IBM i - Print 1-100 Impressions/Minute",
		"5114": "IBM i - Print Anyspeed Impressions/Minute",
		"5116": "IBM i - HA Switchable Resources",
		"5117": "IBM i - HA Journal Performance",
		"5120": "IBM i - Db2 Data Mirroring",
	},
	"5770ST1": {
		"5050": "DB2 Query Manager and SQL Development Kit for i - DB2 and SQL Dev Kit",
	},
	"5770WDS": {
		"5101": "Rational Development Studio for i - ILE Compiler",
		"5102": "Rational Development Studio for i - Heritage Compiler",
		"5103": "Rational Development Studio for i - Application Development ToolSet",
	},
	"5770XW1": {
		"5101": "IBM i Access Family - iSeries Access Family Per User",
	},
	"5798FAX": {
		"5050": "Facsimile Support for iSeries - Facsimile Support",
	},
	"5799GTJ": {
		"5050": "Websphere Development Studio for iSeries, Single Host - WDS for iSeries Single Host Cmp",
	},
	"5799OC2": {
		"BASE": "Red Hat Openshift Enablement for Power - Private Cloud Openshift Enablement for Pools",
	},
	"5799RP2": {
		"BASE": "Red Hat Enablement for Power Private Cloud - Red Hat Enablement for Pools",
	},
	"5799WP7": {
		"BASE": "AIX 5.3 Workload Partitions for AIX 7",
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
