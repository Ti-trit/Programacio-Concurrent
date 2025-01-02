import threading
import time
import random
THREADS = 4


class fil (threading.Thread):

    def __init__(self, id:int, nom_str):
        super().__init__()

    def run(self):
        #fer coses
       print(f"som el fil {self.id}") 
    

class pare (threading.Thread):
    def __init__(self, ):
        super().__init__()

    def run(self):
        print("nover gonna give you up")    


def main():
    #fil pare
    filePare = fil
    filePare.start()

    threads = []
    for i in range(THREADS):
        t = fil(i+1, "fil1")
        threads.append(t)
        t.start()


    filePare.join()
    for t in threads:
        t.join()    

    print("simulació acabada")


if __name__== "__main__":
    main()        

    


