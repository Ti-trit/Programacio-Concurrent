import threading 
import time
import random

THREADS = 6
BUFFER_SIZE = 10
repeticions = 20
global monitor

class monitorLE():
    def __init__(self):
        #nombre de lectors
        
        self.readers = 0
        self.writing = False # si hi ha qualque escriptor a la SC
        #variables de condició
        self.mutex = threading.Lock() #per garantir l'exclusió mutua i rebre les operacions de wait i notify
        self.canRead = threading.Condition(self.mutex)
        self.canWrite = threading.Condition(self.mutex)

    #bloqueja als lectors si ha un esciptor a la SC
    def reader_lock(self):
        with self.mutex:
            while self.writing:
                self.canRead.wait()

            self.readers+=1
            self.canRead.notify() # notificar a la resta de lectors    

    def reader_unlock(self):
        with self.mutex:
            self.readers=-1
            #el darrer lector desbloqueja als escriptors
            if self.readers==0:
                self.canWrite.notify()

    def writer_lock(self):
        with self.mutex:
            while self.writing or self.readers:
                self.canWrite.wait()        
            self.writing=True

    #acabar l'escriptura 
    def writer_unlock (self):
        with self.mutex:
            self.writing= False
            #notifica als que vulguen entrar a la SC
            self.canWrite.notify()
            self.canRead.notify()   
            


class lector(threading.Thread):
    def __init__(self, id:int,  monitor):
        super().__init__()
        self.monitor = monitor
        self.id = id
    global monitor 
    def run(self):
        print(f"Hola, som el lector {self.id}")
        time.sleep(random.uniform(0.1, 3))
        for i in range (repeticions):
            self.monitor.reader_lock()
            print(f"<--- lector {self.id} esta llegint...")
            self.monitor.reader_unlock()
            print(f"---> lector {self.id} ha acabat ")
            time.sleep(random.uniform(0.1, 3))
            

class escriptor (threading.Thread):

    def __init__(self, id:int, monitor):
        super().__init__()
        self.id = id
        self.monitor = monitor

    def run(self):
        print(f"Hola, som l'escriptor {self.id}")
        time.sleep(random.uniform(0.1, 3))
        for i in range (repeticions):
            self.monitor.writer_lock()
            print(f"<---- escriptor {self.id} esta escrivint...")
            self.monitor.writer_unlock()
            print(f"---> escritpor {self.id} ha acabat ")
            print("")
            time.sleep(random.uniform(0.1, 3))

        
def main():
    threads = []
    monitor = monitorLE()

    for i in range (THREADS//2):
        threads.append(lector(i+1, monitor))
        threads.append(escriptor(i+1, monitor))

   

    for t in threads:
        t.start()

    for t in threads :
        t.join()


    print("Simulació acabada.")

if __name__== "__main__":
    main()        
