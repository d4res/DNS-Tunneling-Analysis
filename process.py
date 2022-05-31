from cgi import test
from multiprocessing.managers import ValueProxy
from operator import mod
from re import T
import string
import time
from tracemalloc import start

import joblib

import numpy as np
import pandas as pd
from sklearn.metrics import accuracy_score
from sklearn import datasets
from sklearn import svm
from sklearn.feature_extraction import FeatureHasher


def predict(domain: str) -> bool :
    model = joblib.load(filename="SVM.Model")
    t_h=FeatureHasher(n_features=10,input_type='string')
    t_df=pd.Series(domain)
    t_f = t_h.transform(t_df)
    t_f.toarray()
    res = model.predict(t_f)
    return bool(res[0])

if __name__ == "__main__":
    model = joblib.load(filename="SVM.Model")
    t_h=FeatureHasher(n_features=10,input_type='string')
    t_df=pd.Series("r51646.tunnel.tuns.org.")
    t_f = t_h.transform(t_df)
    t_f.toarray()
    res = model.predict(t_f)
    print(res)
