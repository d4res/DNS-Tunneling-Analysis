import time 

time_window = 15

class Table:
    def __init__(self,hostname:str):
        self.hostname=hostname
        self.time=time_window # 剩余时间
        self.num = 1
    def change(self,changetime: float):
        self.time=self.time - changetime
        if(self.time <0):
            return True
        else:
            return False
    def add(self):
        self.num= self.num+1
        if(self.num <=3):
            return True
        else:
            return False
    def equal(self,hostname:str):
        if(self.hostname == hostname):
            return True
        else:
            return False

class CheckList:
    def __init__(self, curTime: float):
        self.lastTime = curTime
        self.list = []

checkList = CheckList(time.time())

def fun(hostname: str, curTime:float) -> bool:
    midtime =  curTime - checkList.lastTime
    checkList.lastTime = curTime
    if not checkList.list:
        temp=Table(hostname)
        checkList.list.append(temp)
        return False 
    else:
        if(midtime >= time_window):
            checkList.list.clear()
            temp=Table(hostname)
            checkList.list.append(temp)
            return False 
        else:
            for numTable in checkList.list:
                numTable.change(midtime)
                if(numTable.change(midtime)):
                    # del numTable
                    checkList.list.remove(numTable)
            for numTable in checkList.list:
                if(numTable.hostname == hostname ):
                    if(numTable.add()):
                        return False 
                    else:
                        #print(numTable.hostname +" is banned now !")
                        checkList.list.remove(numTable)
                        # print(len(checkList.list))
                        return True 
            temp=Table(hostname)
            checkList.list.append(temp)
            return False

# test
if __name__ == "__main__":
    fun("111", 0)
    fun("111", 1)
    fun("111", 2)
    fun("111", 12)
