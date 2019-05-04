package cn.ywzou.nacos.controller;

import java.util.HashMap;
import java.util.Map;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class TestController {

	private static final Logger log = LoggerFactory.getLogger(TestController.class);

	@GetMapping("/nacos/hello")
	public Map<String, String> hello(@RequestParam String name) {
		log.info(">>>>>> 服务提供方接收到消息 name = {}", name);
		Map<String, String> result = new HashMap<String, String>();
		result.put("code", "0000");
		result.put("msg", "Hello " + name);
		return result;
	}

}
