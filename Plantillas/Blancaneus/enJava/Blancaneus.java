package Blancaneus.enJava;

import java.util.Random;
import java.util.logging.Level;
import java.util.logging.Logger;

public class Blancaneus implements Runnable {

    public Blancaneus(){

    }
    @Override

    public void run() {
        try {
            Random rn = new Random();
            main.mutex_nans.acquire();
            while (main.nans_cabats<main.NUM_NANS) {
                main.mutex_nans.release();
                main.mutex.acquire();
                    if (main.volMenjar==0){
                        main.mutex.release();
                        System.out.println(" Blancaneus s'ha anat a fer una passejada");
                        Thread.sleep(rn.nextInt(6));
                    }else{
                        main.mutex.release();
                        System.out.println(" Blancaneus esta preparanet el menjar");
                        Thread.sleep(rn.nextInt(1));
                        //avisar al nan que pot menjar
                        System.out.println(" Blacanues: Ja pots menjar");
                        main.EsperarMenjar.release();

                    }


            }
            main.mutex_nans.release();

            System.out.println("Blancaneus: Tots els nans han anat a dormir, jo també m'aniré");

        }catch (InterruptedException ex){
            Logger.getLogger(nan.class.getName()).log(Level.SEVERE, null, ex);
        }

    }
}
