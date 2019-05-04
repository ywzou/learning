package cn.ywzou.nacos.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cloud.client.ServiceInstance;
import org.springframework.cloud.client.loadbalancer.LoadBalancerClient;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

@RestController
public class TestController {

	// 接口在获取服务实例的时候，已经实现了对服务提供方实例的负载均衡
	@Autowired
	LoadBalancerClient loadBalancerClient;

	@GetMapping("/nacos/client")
	public String test() {
		// 通过spring cloud common中的负载均衡接口选取服务提供节点实现接口调用
		ServiceInstance serviceInstance = loadBalancerClient.choose("nacos-discovery-server");
		String url = serviceInstance.getUri() + "/nacos/hello?name=" + "ywzou";
		RestTemplate restTemplate = new RestTemplate();
		String result = restTemplate.getForObject(url, String.class);
		return "Invoke : " + url + ", return : " + result;
	}
}
