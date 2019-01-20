package cn.ywzou.thread;

import java.util.concurrent.Callable;

public class Ticket {

    public static void main(String[] args) {
        TicketTheard tic = new Ticket().new TicketTheard();

        for (int i = 0; i < 10; i++) {
            Thread t1 = new Thread(tic);
            t1.start();
        }

//        Thread t1 = new Thread(tic);
//        Thread t2 = new Thread(tic);
//        Thread t3 = new Thread(tic);
//
//        t1.start();
//        t2.start();
//        t3.start();

    }

    public class A implements Callable<Object>{
        @Override
        public Object call() throws Exception {
            return null;
        }
    }

    public class TicketTheard implements Runnable {
        private volatile int ticket = 5;


        @Override
        public void run() {
            sale();
        }

        public synchronized void sale() {
            if (ticket > 0) {
//                try {
//                    Thread.sleep(300);
//                } catch (InterruptedException e) {
//                    e.printStackTrace();
//                }
                System.out.println(Thread.currentThread().getName() + "买票：ticket = " + (ticket--));
            }
        }
    }
}
