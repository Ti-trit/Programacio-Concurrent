import random
import threading
import time

NUM_ELFS = 6
NUM_RENS = 9
NUM_JOGUINES = 3

elfs_amb_dubtes = []  # llista dels elfs amb dubtes que representa la sala d'espera
mutex_elfs = threading.Semaphore(1)  # Semaphore per protegir la llista elfs_amb_dubtes

rens_arribats = 0  # comptador dels rens que van arribant
mutex_rens_arribants = threading.Semaphore(1)  # Semaphore per protegir la variable rens_arribants

dubtes_a_resoldre = threading.Semaphore(0)  # per notificar al Pare Noel que hi ha 3 elfs amb dubtes a la sala d'espera
rens_llestos = threading.Semaphore(0)  # per notificar al Pare Noel que tots els rens han arribat

joguines_construides = 0  # variable global compartida per comptar el total de joguines construides
mutex_joquines_construides = threading.Semaphore(1)  # per protegir la variable joguines_contruides

sem_espera_elfs = threading.Semaphore(3)  # semaphore per poder entrar a la sala d'espera
joguines_llestes = threading.Semaphore(0)  # per notificar al Pare Noel que totes les joguines estan llestes


threads_RENS = []


class ELF(threading.Thread):
    """
    Mètode constructor. Cada elf té un nom, un índex de joguina que va actualitzant
    i un semàfor que li notifica quan el seu dubte s'ha resolt pel Pare Noel.
    """

    def __init__(self, nom: str):
        super().__init__()
        self.nom = nom
        self.index_joguina = 0
        self.dubteResolta = threading.Semaphore(0)
        print(f"Hola som l'elf {self.nom} construiré {NUM_JOGUINES} joguines")
        # time.sleep(random.uniform(0.01, 0.2))


    def dubte(self):
        time.sleep(random.uniform(0.001, 0.03))
        sem_espera_elfs.acquire()  # Si ja hi ha 3 elfs a la sala d'espera no pot entrar
        mutex_elfs.acquire()  # obtenir permís per modificar la llista de elfs_amb_dubtes

        elfs_amb_dubtes.append(self)  # afegir l'elf que té dubtes en el moment actual
        # Se'l suma 1 a l'índex perquè aquest comença a 0, mentre que 
        # per l'exemple comencen per 1        
        print(f"{self.nom} diu: tinc dubtes amb la joguina {self.index_joguina + 1}")

        if len(elfs_amb_dubtes) == 3:
            print(f"{self.nom} diu: Som 3 que tenim dubtes, PARE NOEEEEEL!")
            dubtes_a_resoldre.release()  # Notificar al Pare Noel que hi ha 3 elfs amb dubtes

        mutex_elfs.release()

    def construir_joguina(self):
        global joguines_construides

        self.dubteResolta.acquire()  # l'elf espera que el seu dubte estigui resolt
        print(f"\n{self.nom} diu: Construeixo la joguina amb ajuda")
        self.index_joguina += 1

        mutex_joquines_construides.acquire()  # permís per incrementar el total de joguines construides
        joguines_construides += 1  # Incrementar en 1 per cada joguina
        if joguines_construides == NUM_ELFS * NUM_JOGUINES:

            joguines_llestes.release()  # Notificar al Pare Noel que totes les joguines estan llestes

        mutex_joquines_construides.release()  # donar permís a altres elfs per modificar/consultar la variable joguines_construides
        sem_espera_elfs.release()  # dona permís a la resta d'elfs per entrar a la sala d'espera


    def construccio_acabada(self):
        #Si l'elf ha constuit 3 joguines
        if self.index_joguina == NUM_JOGUINES:
            print(f"L'elf {self.nom} ha fet les seves joguines i acaba <---------")

    def run(self):
        for i in range(NUM_JOGUINES):
            self.dubte()
            self.construir_joguina()
            self.construccio_acabada()


class REN(threading.Thread):
    """
    Mètode constructor.
    Un ren només ve representat per un nom
    """

    def __init__(self, nom: str):
        super().__init__()
        self.nom = nom

    def pasturar(self):
        print(f"{self.nom} se'n va a pasturar")
        time.sleep(random.uniform(0.001, 0.003))

    def arribar(self):
        time.sleep(random.uniform(0.03, 0.07))
        # incrementar el total de rens que han arribat
        global rens_arribats
        mutex_rens_arribants.acquire()
        rens_arribats += 1

        # si ja hi són tots els rens, notifiquem al Pare Noel
        if rens_arribats == NUM_RENS:
            print(f"El ren {self.nom} diu: Som el darrer en voler podem partir")
            rens_llestos.release()

        else:
            # Si no es tracta del darrer ren
            print(f"\tEl ren {self.nom} arriba, {rens_arribats}")
            time.sleep(random.uniform(0.02, 0.3))

        mutex_rens_arribants.release()

    def run(self):
        self.pasturar()
        self.arribar()


class PareNoel(threading.Thread):


    def __init__(self):
        super().__init__()
        print("-------> El Pare Noel diu: estic despert però me'n torn a jeure")


    """
    El thread del Pare Noel es troba en un bucle, mentre que no hi hagi 3 elfs amb dubtes,
    no s'hagin construït totes les joguines o no hagin arribat tots els rens, el Pare Noel descansa.
    """

    @staticmethod
    def resoldre_dubtes():
        global joguines_construides
        while 1:
            dubtes_a_resoldre.acquire()  # Espera que li notifiquen que hi ha dubtes
            print("-------> El Pare Noel diu: Atendré els dubtes d'aquests 3")
            mutex_elfs.acquire()
            # simulació de resoldre els dubtes
            # Elimina tots els elfs de la llista i els notifica que el seu dubte està resolt
            for i in range(3):
                elfs_amb_dubtes[i].dubteResolta.release()
            elfs_amb_dubtes.clear()
            mutex_elfs.release()
            print("-------> El Pare Noel diu: estic cansat me'n torn a jeure")
            # esperar que els elfs acabin de contruir les joguines
            r = random.uniform(0.1, 0.2)
            time.sleep(r)
            # Condició de sortida del bucle: que s'hagin construit totes les joguines
            mutex_joquines_construides.acquire()

            if joguines_construides == NUM_JOGUINES * NUM_ELFS:
                mutex_joquines_construides.release()

                break
            mutex_joquines_construides.release()

    @staticmethod
    def preparar_per_enganxar_rens():
        joguines_llestes.acquire()
        print("-------> Pare Noel diu: Les joguines estan llestes. I Els rens?")

    @staticmethod
    def enganxar_rens():
        rens_llestos.acquire()  # Espera que arribin tots els rens
        mutex_rens_arribants.acquire()
        # simulació del procés d'enganxar als rens
        if rens_arribats == NUM_RENS:
            print("-------> Pare Noel diu: Enganxaré els rens i partiré")
            for ren in threads_RENS:
                print(f"El ren {ren.nom} està enganxat al trineu")
                r = random.uniform(0.1, 0.3)
                time.sleep(r)

            print("-------> El Pare Noel ha enganxat els rens, ha carregat les joguines i se'n va \n SIMULACIÓ ACABADA")

    def run(self):

        self.resoldre_dubtes()
        self.preparar_per_enganxar_rens()
        self.enganxar_rens()


def main():
    threads_elfs = []
    noms_elfs = ['Taleasin', 'Halafarin', 'Ailduin', 'Adamar', 'Galather', 'Estelar']

    global threads_RENS
    noms_rens = ['RUDOLPH', 'BLITZEN', 'DONDER', 'CUPID', 'COMET', 'VIXEN', 'PRANCER', 'DANCER', 'DASHER']
    print("SIMULACIÓ DEL PARE NOEL I ELS ELFS EN PRÀCTIQUES")

    pare_noel = PareNoel()
    pare_noel.start()

    #seleccionar els noms de forma aleatoria

    random.shuffle(noms_rens)
    for nom in noms_rens:
        ren = REN(nom)
        threads_RENS.append(ren)
        ren.start()

    random.shuffle(noms_elfs)
    for nom in noms_elfs:
        elf = ELF(nom)
        threads_elfs.append(elf)
        elf.start()


    pare_noel.join()

    for ren in threads_RENS:
        ren.join()
    for elf in threads_elfs:
        elf.join()



if __name__ == "__main__":
    main()
