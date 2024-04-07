import json

with open("usb.ids", "r", encoding = "ISO-8859-1") as file:
	ids = file.readlines()
	data = {}
	curr_vendor = {}
	curr_vendor_key = ""
	for idx, val in enumerate(ids):
		if val[0] != "#" and val != "\n":
			if val[0:1] == "\t":
				val = val[1:len(val)]
				val = val[:len(val)-1]
				item = val.split(" ", maxsplit=2)
				curr_vendor[curr_vendor_key]["devices"][item[0]] = item[2]
			else:
				if curr_vendor != {}:
					data.update(curr_vendor)
					curr_vendor = {}
				if val[0] == "C":
					break
				val = val[:len(val)-1]
				item = val.split(" ", maxsplit=2)
				curr_vendor = {
					item[0]: {
						"id": item[0],
						"name": item[2],
						"devices": {}
					}
				}
				curr_vendor_key = item[0]
			
	with open("result.json", "w") as resultantfile:
		json.dump(data, resultantfile, sort_keys=False, indent=4)