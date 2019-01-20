package cn.ywzou.thread;


/***
 *
 * 线程基础
 */
public class ThreadDemo {

    public static void main(String args[]) {
        ThreadOne threadOne = new ThreadDemo().new ThreadOne("线程A");
        TheadTwo theadTwo = new ThreadDemo().new TheadTwo("线程B");

        //启动
        threadOne.start();
        new Thread(theadTwo).start();
    }

    /**
     * 实现方式一 继承 Thread
     */
    public class ThreadOne extends Thread {
        private String name;

        public ThreadOne(String name) {
            this.name = name;
        }

        @Override
        public void run() {
            for (int i = 0; i < 10; i++) {
                System.out.println(name + "运行， i = " + i);
            }
        }
    }

    /**
     * 实现方式二 实现 Runnable 接口
     */
    public class TheadTwo implements Runnable {
        private String name;

        public TheadTwo(String name) {
            this.name = name;
        }

        @Override
        public void run() {
            for (int i = 0; i < 10; i++) {
                System.out.println(name + "运行， i = " + i);
            }
        }
    }
}

