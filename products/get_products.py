#! /usr/bin/env python3


from bs4 import BeautifulSoup
import os
import json 
from collections import OrderedDict
import re


def get_products():

    NamesFromPrdid = {}
    NamesFromPrdidAndFeature = {}

    # Does not work with requests for some reason
    url = "https://www.ibm.com/docs/en/entitled-systems-support?topic=migration-inventory-software-products-eligible"
    path_to_output = "products.html"
    os.system(f"curl -s '{url}' -o '{path_to_output}'")
    with open(path_to_output, "r") as f:
        content = f.read()

    soup = BeautifulSoup(content, "lxml")
    body = soup.find("div", attrs={"class": ["body", "conbody"]})

    for tr in body.find_all("tr"):
        tds = tr.find_all("td")
        if len(tds) == 7:
            ProductID = tds[0].text.strip().upper()
            ProductName = tds[1].text.strip().replace("\n", " ")
            FunctionTitle = tds[2].text.strip().replace("\n", " ")
            FunctionID = tds[3].text.strip()
            EntitlementType = tds[4].text.strip()
            EnabledForTransfer = tds[5].text.strip()
            ListedUnder = tds[6].text.strip()
            
            if ProductID not in NamesFromPrdidAndFeature:
                NamesFromPrdidAndFeature[ProductID] = {}

            if ProductID not in NamesFromPrdid:
                NamesFromPrdid[ProductID] = ProductName
            
            if FunctionID not in NamesFromPrdidAndFeature[ProductID]:
                if ProductName == FunctionTitle:
                    NamesFromPrdidAndFeature[ProductID][FunctionID] = ProductName
                else:
                    NamesFromPrdidAndFeature[ProductID][FunctionID] = f"{ProductName} - {FunctionTitle}"
            
    return NamesFromPrdid, NamesFromPrdidAndFeature


template = """
package products

// Sources: 
// - https://www.ibm.com/docs/en/i/7.4.0?topic=reference-media-labels-their-contents
// - https://www.ibm.com/docs/en/entitled-systems-support?topic=migration-inventory-software-products-eligible

var NamesFromPrdid = map[string]string%s

var NamesFromPrdidAndFeature = map[string]map[string]string%s

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
"""


if __name__ == "__main__":
    NamesFromPrdid, NamesFromPrdidAndFeature = get_products()

    # Sort the products and features
    OrderedNamesFromPrdidAndFeature = OrderedDict()
    for ProductID in sorted(NamesFromPrdidAndFeature.keys()):
        OrderedNamesFromPrdidAndFeature[ProductID] = {}
        for FunctionID in sorted(NamesFromPrdidAndFeature[ProductID].keys()):
            OrderedNamesFromPrdidAndFeature[ProductID][FunctionID] = NamesFromPrdidAndFeature[ProductID][FunctionID]
    OrderedNamesFromPrdidAndFeatureString = json.dumps(OrderedNamesFromPrdidAndFeature, indent=4)
    OrderedNamesFromPrdidAndFeatureString = re.sub(r"\"\n", "\",\n", OrderedNamesFromPrdidAndFeatureString)
    OrderedNamesFromPrdidAndFeatureString = re.sub(r"}\n", "},\n", OrderedNamesFromPrdidAndFeatureString)
    
    # Sort the products
    OrderedNamesFromPrdid = OrderedDict()
    for ProductID in sorted(NamesFromPrdid.keys()):
        OrderedNamesFromPrdid[ProductID] = NamesFromPrdid[ProductID]
    OrderedNamesFromPrdidString = json.dumps(OrderedNamesFromPrdid, indent=4)
    OrderedNamesFromPrdidString = re.sub(r"\"\n", "\",\n", OrderedNamesFromPrdidString)
    OrderedNamesFromPrdidString = re.sub(r"}\n", "},\n", OrderedNamesFromPrdidString)
    
    # Write the products and features to the output file
    with open("products.go", "w") as f:
        f.write(template % (OrderedNamesFromPrdidString, OrderedNamesFromPrdidAndFeatureString))