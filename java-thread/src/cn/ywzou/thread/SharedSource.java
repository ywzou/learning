package cn.ywzou.thread;

/**
 * 共享资源
 */
public class SharedSource {

    public static void main(String args[]) {
        Integer ticket = 5;
        ThreadOne threadOne1 = new SharedSource().new ThreadOne("线程A1", ticket);
        ThreadOne threadOne2 = new SharedSource().new ThreadOne("线程A2", ticket);

        TheadTwo theadTwo = new SharedSource().new TheadTwo("线程B1", ticket);

        // 启动 两个线程 两个线程均执行了ticket次 资源没有共享
        threadOne1.start();
        threadOne2.start();
        System.out.println("================");

        //启动 两个线程 两个线程总共执行了ticket次 实现了资源的共享
        new Thread(theadTwo).start();
    }

    /**
     * 实现方式一 继承 Thread
     * 继承 Thread类不能资源共享
     */
    public class ThreadOne extends Thread {
        private String name;
        private Integer ticket;


        public ThreadOne(String name, Integer ticket) {
            this.name = name;
            this.ticket = ticket;
        }

        @Override
        public void run() {
            for (int i = 0; i < 50; i++) {
                if (this.ticket > 0)
                    System.out.println(name + "卖票， 剩余票数 = " + (ticket--));
            }
        }
    }

    /**
     * 实现方式二 实现 Runnable 接口
     * 实现 Runnable 接口可以实现资源的共享
     */
    public class TheadTwo implements Runnable {
        private String name;
        private Integer ticket;

        public TheadTwo(String name, Integer ticket) {
            this.name = name;
            this.ticket = ticket;
        }

        @Override
        public void run() {
            for (int i = 0; i < 50; i++) {
                if (this.ticket > 0)
                    System.out.println(name + "卖票， 剩余票数 = " + (ticket--));
            }
        }
    }

}


