package problema_LE.monitors.Implicits;

import problema_LE.monitors.Implicits.LectorEscriptor_implicit;

import java.util.logging.Level;
import java.util.logging.Logger;

public class monitorLE {
    volatile int readers = 0;
    volatile boolean writers = false;

    synchronized void readerLock() {
        while (writers){
            try{
                wait();
            }catch (InterruptedException ex){
                Logger.getLogger(monitorLE.class.getName()).log(Level.SEVERE, null, ex);
            }
        }
        readers++;
        notifyAll(); //desperta a tots els processos bloquejats

    }
    synchronized void readerUnlock() {
        readers--;
        //si és el darrer notifica a tothom
        if (readers==0){
            notifyAll();
        }
    }
    synchronized void writerLock() {
        while (writers || readers>0){
            try{
                wait();

            }catch (InterruptedException ex){}
        }
        writers=true;
    }
    synchronized void writerUnlock() {
        writers=false;
        notifyAll();
    }

}


