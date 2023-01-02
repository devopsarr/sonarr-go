
# Python program to read
# json file

import json
import re

# Opening JSON file
f = open('openapi.json','r+')

# returns JSON object as 
# a dictionary
data = json.load(f)

# Loop over paths
for path, methods in data['paths'].items():
    # for each method adjust the name
    for method, details in methods.items():
        method_name = re.sub("^\/api\/v3\/(.*)$","\\1", path).split("/")
        if method_name[-1].startswith("{"):
            method_name[-1] = method_name[-1].strip('{}')
            method_name.insert(-1,"by")
        final_method= method
        if method == "delete" and len(method_name) > 1 and method_name[-2] == "by":
            method_name = method_name[:-2]     
        if method == "put" and len(method_name) > 1 and method_name[-2] == "by":
            final_method = "update"
            method_name = method_name[:-2]
        if method == "post":
            final_method = "create"
            if method_name[-1].startswith("test"):
                final_method = method_name[-1]
                method_name = method_name[:-1]
        try:
            if method == "get" and details['responses']['200']['content']['application/json']['schema']['type'] == "array":
                final_method = "list"
        except:
            print("no array")
        method_name.insert(0,final_method)
        # print(''.join([str.capitalize(i) for i in method_name]))
        data['paths'][path][method]['operationId']= ''.join([str.capitalize(i) for i in method_name])

# Overwrire file content
f.seek(0)
f.write(json.dumps(data, indent=2))
# Closing file
f.close()