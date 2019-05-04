package cn.ywzou.nacos;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;


//开启Spring Cloud的服务注册与发现 有alibaba-nacos 实现
@EnableDiscoveryClient
@SpringBootApplication
public class NacosDiscoveryServerApplication {

	public static void main(String[] args) {
		SpringApplication.run(NacosDiscoveryServerApplication.class, args);
	}

}
