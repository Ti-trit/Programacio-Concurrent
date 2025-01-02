
import random
import time
import threading

num_abelles = 10
mida_pot = 10
notPle = threading.Semaphore(mida_pot) #posar mel al pot
avisarOs = threading.Semaphore(0) #avisar a l'os
mutex = threading.Lock() #protegir el pot
pot = 0 
repeticions = 20

class Os (threading.Thread):
    def __init__(self):
        super().__init__()

    def run(self):
        global pot, repeticions
    
        for  _ in range (repeticions):
        #esta dormint
            time.sleep(random.uniform(0.1, 0.3))

            #esperar que l'avisen
            avisarOs.acquire()
            print("yummy. Començaré a menjar la mel")
            mutex.acquire()
            pot = 0
            time.sleep(random.uniform(0.1, 0.3))
            mutex.release()
            #alliberar a les abelles
            notPle.release(mida_pot)
            print("L'os ha menjat tot el mel, i se'n va a dormir")
        




class Abella (threading.Thread):
    def __init__(self, id:int):
        super().__init__()
        self.id = id


    def run(self):
        global pot, repeticions

        for i in range (repeticions):
            time.sleep(random.uniform(0.001, 0.03))

            #agafar permis per produir
            notPle.acquire()
            mutex.acquire()
            pot = pot+1
            print(f"L'abella {self.id} ha posat una porció de mel. Ara n'hi ha {pot}" )
            if (pot == mida_pot):
                print(f"L'abella {self.id}: El pot está ple.  Despertaré a l'ós")
                #avisar a l'os
                avisarOs.release()
            mutex.release()   
            time.sleep(0.1)
             


        


def main():
    
    os = Os()
    os.start()

    threads_abelles = []
    for i in range(num_abelles):
        t = Abella(i+1)
        threads_abelles.append(t)
        t.start()


    os.join()
    for t in threads_abelles:
        t.join()    
   # os.join()
    print("simulació acabada")


if __name__== "__main__":
    main()        