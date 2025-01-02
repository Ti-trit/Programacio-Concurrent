import threading

THREADS = 4


def thread():
    print()


def main():
    threads = []
    for i in range(THREADS):
        t = threading.Thread(target=thread)
        threads.append(t)
        t.start()


    for t in threads:
        t.join()    


if __name__== "__main__":
    main()        