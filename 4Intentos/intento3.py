#pseudocode
#Tercer intent de Dekker

#loop forever
#non-critical section
#wantq = true
#await wantq = false
#critical section
#wantp = false


import threading

THREADS = 2
MAX_COUNT = 10000000

want = [False]*THREADS
n = 0  # Contador compartido
n_lock = threading.Lock()  # Lock para proteger el acceso a 'n'

class CounterDekker(threading.Thread):
    def __init__(self, id):
        threading.Thread.__init__(self)
        self.id = id

    def run(self):
        global torn, n
        max_count = MAX_COUNT // THREADS
        altre_proces = (self.id + 1) % THREADS

        print(f"Thread {self.id} started")

        for i in range(max_count):
            
            want[self.id]=True
            #espera activa
            while want[altre_proces]:
                pass 

            # Sección crítica

            with n_lock:  
                n += 1

            want[self.id]=False

def main():
    global n
    threads = []

    for i in range(THREADS):
        t = CounterDekker(i)
        #t = threading.Thread(target=CounterDekker(i))
        threads.append(t)
        t.start()

    for t in threads:
        t.join()

    error = (MAX_COUNT - n) / MAX_COUNT * 100
    print(f"Counter value: {n} Expected: {MAX_COUNT} Error: {error:.6f}%")

if __name__ == "__main__":
    main()
