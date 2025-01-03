package probelmaAbelles.Monitors.Implicits;

import java.util.Random;

public class Abella implements Runnable{
    int id;
    monitorOsAbella monitor;

    public Abella(int id, monitorOsAbella monitor){
        this.id = id;
        this.monitor = monitor;
    }
    Random rn = new Random();
    @Override
    public void run() {
        for (int i = 0; i < monitor.repeticions; i++) {
            monitor.esperarPotBuid();
            try {
                Thread.sleep(rn.nextInt(100));
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
            monitor.posar_mel(this.id);
        }


    }
}
