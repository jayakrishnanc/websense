#!/usr/bin/python

import pickle
import gib_detect_train
import sys

filepath = sys.argv[1]
outfile  = filepath + ".out"
model_data = pickle.load(open('gib_model.pki', 'rb'))

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
