package cn.ywzou.thread;

/**
 * 线程常用你的API方法
 */
public class ThreadAPI {

    public static void main(String[] args) {
        MyThread my = new ThreadAPI().new MyThread();

        //启动线程
        Thread t1 = new Thread(my);
        t1.start();

        //优先级
        t1.setPriority(Thread.MAX_PRIORITY);//高
        t1.setPriority(Thread.MIN_PRIORITY);//低
        t1.setPriority(Thread.NORM_PRIORITY);//中等

        //线程是否启动判断
        System.out.println("线程是否活着：" + t1.isAlive());

        //强制运行
        try {
            t1.join();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        //休眠
        try {
            Thread.sleep(500);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        //中断
        t1.interrupt();



        // 创建时指定线程名称
        new Thread(my,"我的线程").start();
    }

    public class MyThread implements Runnable {
        @Override
        public void run() {
            for (int i = 0; i < 3; i++) {
                //获取线程名称
                System.out.println(Thread.currentThread().getName() + "运行，i=" + i);
            }
        }
    }
}
