#!/usr/bin/python

import pickle
import gib_detect_train
import sys
import os

filepath = sys.argv[1]
outfile  = filepath + ".out"
bin_location = sys.argv[0]
bin_dir = os.path.dirname(bin_location)
model_file = bin_dir + "/gib_model.pki"
model_data = pickle.load(open(model_file, 'rb'))

with open(filepath) as fp:
   line = fp.readline()
   cnt = 1
   file = open(outfile,"w") 
   while line:
       cnt += 1
       model_mat = model_data['mat']
       threshold = model_data['thresh']
       if line:
           print line.strip(), ":", gib_detect_train.avg_transition_prob(line, model_mat) > threshold
       line = fp.readline()
