from cgi import test
from multiprocessing.managers import ValueProxy
from operator import mod
from re import T, fullmatch
import string
import time
from tracemalloc import start
import os 
import joblib

import numpy as np
import pandas as pd
from sklearn.metrics import accuracy_score
from sklearn import datasets
from sklearn import svm
from sklearn.feature_extraction import FeatureHasher
import pymongo
from check import fun 


def predict(domain: str) -> bool :
    model = joblib.load(filename="SVM.Model")
    t_h=FeatureHasher(n_features=10,input_type='string')
    t_df=pd.Series(domain)
    t_f = t_h.transform(t_df)
    t_f.toarray()
    res = model.predict(t_f)
    myclient = pymongo.MongoClient('mongodb://localhost:27017/')
    mydb = myclient["dns_pcap"]
    mycol = mydb["blackList"]
    d = domain[domain.find('.') + 1: ]
    res = mycol.find_one({"domain": d})
    if bool(res[0]) and res != None:
        if (fun(d, time.time())):
            mycol.insert_one({"domain": d})
    return bool(res[0])

if __name__ == "__main__":
    # model = joblib.load(filename="SVM.Model")
    # t_h=FeatureHasher(n_features=10,input_type='string')
    # t_df=pd.Series("r51646.tunnel.tuns.org.")
    # t_f = t_h.transform(t_df)
    # t_f.toarray()
    # res = model.predict(t_f)

    for i in range(0, 5):
        time.sleep(1)
        predict("r51646.tunnel.tuns.org.")
    # print(res)
