package probelmaAbelles.Monitors.Implicits;

import java.util.Random;

public class os implements Runnable{

    monitorOsAbella monitor;
    public os(monitorOsAbella monitor) {
        this.monitor = monitor;
    }
    Random rn = new Random();
    @Override
    public void run() {
        try {
            System.out.println("Hola, som l'os");
            for (int i = 0; i < monitor.repeticions; i++) {
                monitor.esperarPotPle();
                Thread.sleep(rn.nextInt(20));
                monitor.consumir_mel();
            }

        } catch (InterruptedException e) {
            throw new RuntimeException(e);
        }


    }
}
