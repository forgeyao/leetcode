#!/usr/bin/python3
# -*- coding: UTF-8 -*-

import os
import shutil

dir_list = os.listdir()
for name in dir_list:
    if os.path.isfile(name):
        continue

    strs = name.split("-")
    if len(strs) < 1:
        print ("error:",name)
        break

    try:
        num = int(strs[0])
    except ValueError as e:
        print (name,"is not dst dir")
        continue

    dst = ""
    if num < 201:
        dst = "1-200"
    elif num < 501:
        dst = "201-500"
    elif num < 1001:
        dst = "501-1000"
    elif num < 1501:
        dst = "1001-1500"
    elif num < 2001:
        dst = "1501-2000"
    else:
        dst = "2001-"

    print (name,dst)
    shutil.move(name, dst)
