package probelmaAbelles.EnJava.src;

import java.util.Random;
import java.util.logging.Level;
import java.util.logging.Logger;

public class os implements Runnable{


    public os(){
        
    }

    @Override
    public void run() {
       try {
        Random rn = new Random();
        for (int i = 0; i < osAbelles.repeticions; i++) {
               //es troba dormint
               Thread.sleep(10);
               //RETARD DE MENJAR
               int espera = rn.nextInt(10);
               Thread.sleep(espera);
               //esperar que l'avisen
               osAbelles.avisarOs.acquire();
               System.out.println("yummy. Començaré a menjar la mel");
               osAbelles.mutex.acquire();
               osAbelles.pot = 0; //buidar
               osAbelles.mutex.release();
               System.out.println("L'os ha menjat tot el mel, i se'n va a dormir");
               //allibera  les abelles
               osAbelles.notPle.release(osAbelles.maxim);
           }
       } catch (InterruptedException ex) {
        Logger.getLogger(os.class.getName()).log(Level.SEVERE, null, ex);
    }
    }


    
}
