package Blancaneus.enJava;
import java.util.logging.Level;
import java.util.logging.Logger;
import java.util.Random;

public class nan implements Runnable {

    private int id;
    public nan(int id){
        this.id= id;
    }
    @Override
    public void run() {
        try {
            Random rn = new Random();
            for (int i = 0; i < main.NUM_MENJADES; i++) {

                //va a la mina
                System.out.println("El nan " + this.id + " va a la mina");
                //retard
                int espera = rn.nextInt(100);
                Thread.sleep(espera);
                System.out.println("El nan " + this.id + " ha tornat de la mina");

                main.cadires.acquire();//hi ha qualque cadira lliure?
                System.out.println("---> nan "+ this.id+ " : Blancaneus, he trobat una cadira per seure!!");

                main.mutex.acquire();
                main.volMenjar++;
                main.mutex.release();

                //espera que le donen el menjar
                main.EsperarMenjar.acquire();
                System.out.println("El nan "+ this.id + " : ha estat servit");
                Thread.sleep(rn.nextInt(12));
                System.out.println("El nan "+ this.id + " : ha acabat de menjar --->");

                main.mutex.acquire();
                main.volMenjar--;
                main.mutex.release();

                //deixa la cadira
                main.cadires.release();
                System.out.println("El nan " + this.id + " : he deixat la cadira --->");
            }
            //incrementar numero de nans acabats
            main.mutex_nans.acquire();
            main.nans_cabats++;
            main.mutex_nans.release();
            System.out.println("----------->> El nan "+ this.id + " ha anat a dormir");
        }catch (InterruptedException ex){
            Logger.getLogger(nan.class.getName()).log(Level.SEVERE, null, ex);
        }

    }
}
