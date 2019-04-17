package cn.ywzou.thread;

import java.io.BufferedOutputStream;
import java.io.BufferedReader;
import java.io.File;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.UnsupportedEncodingException;
import java.net.URLEncoder;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.Map.Entry;

import org.apache.commons.lang3.StringUtils;
import org.apache.http.NameValuePair;
import org.apache.http.client.config.RequestConfig;
import org.apache.http.client.config.RequestConfig.Builder;
import org.apache.http.client.entity.UrlEncodedFormEntity;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.entity.ContentType;
import org.apache.http.entity.StringEntity;
import org.apache.http.entity.mime.MultipartEntityBuilder;
import org.apache.http.entity.mime.content.FileBody;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.message.BasicNameValuePair;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

/**
 * 作者：ywzou <br>
 * 创建时间：2017年9月22日 <br>
 * 描述：http请求工具
 */
public class HttpClientUtil {
	private final static Logger logger = LoggerFactory.getLogger(HttpClientUtil.class);

	/**
	 * 作者：ywzou <br>
	 * 创建时间：2017年9月22日 <br>
	 * 描述： get请求
	 * 
	 * @param url
	 * @param headers
	 * @param querys
	 * @return
	 * @throws Exception
	 */
	public static String doGet(String url, Map<String, Object> headers, Map<String, Object> querys) throws Exception {
		CloseableHttpClient httpclient = HttpClients.createDefault();
		Builder builder = RequestConfig.custom().setConnectionRequestTimeout(6000).setConnectTimeout(6000);
		RequestConfig config = builder.build();
		CloseableHttpResponse response = null;
		try {
			HttpGet httpGet = new HttpGet(buildUrl(url, querys));
			httpGet.setConfig(config);
			if (headers != null && !headers.isEmpty()) {
				for (Map.Entry<String, Object> e : headers.entrySet()) {
					httpGet.addHeader(e.getKey(), e.getValue().toString());
				}
			}

			response = httpclient.execute(httpGet);
			if (null != response) {
				InputStream inputStream = response.getEntity().getContent();
				return convertStreamToString(inputStream);
			}
		} catch (Exception e) {
			logger.error("######【" + url + "】请求异常!", e);
		} finally {
			if (response != null) {
				try {
					response.close();
				} catch (IOException e) {
					e.printStackTrace();
				}
			}
		}
		return null;
	}
	

	/**
	 * 作者：ywzou <br>
	 * 创建时间：2017年9月22日 <br>
	 * 描述： post请求
	 * 
	 * @param url     请求路径
	 * @param headers 请求头
	 * @param querys  请求路径参数 url参数
	 * @param bodys   请求体 表单参数
	 * @return
	 */
	public static String doPost(String url, Map<String, Object> headers, Map<String, Object> querys,
			Map<String, Object> bodys) {
		CloseableHttpClient httpclient = HttpClients.createDefault();
		Builder builder = RequestConfig.custom().setConnectionRequestTimeout(6000).setConnectTimeout(6000);
		RequestConfig config = builder.build();
		CloseableHttpResponse response = null;
		try {
			HttpPost httpPost = new HttpPost(buildUrl(url, querys));
			httpPost.setConfig(config);
			if (headers != null && !headers.isEmpty()) {
				for (Map.Entry<String, Object> e : headers.entrySet()) {
					httpPost.addHeader(e.getKey(), e.getValue().toString());
				}
			}
			if (bodys != null) {
				List<NameValuePair> nameValuePairList = new ArrayList<NameValuePair>();
				for (String key : bodys.keySet()) {
					nameValuePairList.add(new BasicNameValuePair(key, bodys.get(key).toString()));
				}
				UrlEncodedFormEntity formEntity = new UrlEncodedFormEntity(nameValuePairList, "UTF-8");
				formEntity.setContentType("application/x-www-form-urlencoded; charset=UTF-8");
				httpPost.setEntity(formEntity);
			}

			response = httpclient.execute(httpPost);
			if (null != response) {
				InputStream inputStream = response.getEntity().getContent();
				return convertStreamToString(inputStream);
			}
		} catch (Exception e) {
			logger.error("######【" + url + "】请求异常!", e);
		} finally {
			if (response != null) {
				try {
					response.close();
				} catch (IOException e) {
					e.printStackTrace();
				}
			}
		}
		return null;
	}

	/**
	 * 作者：ywzou <br>
	 * 创建时间：2017年9月22日 <br>
	 * 描述： post请求
	 * 
	 * @param url     请求路径
	 * @param headers 请求头
	 * @param querys  请求路径参数 url参数
	 * @param body    请求体 表单参数
	 * @return
	 */
	public static String doPost(String url, Map<String, Object> headers, Map<String, Object> querys, String body) {
		CloseableHttpClient httpclient = HttpClients.createDefault();
		Builder builder = RequestConfig.custom().setConnectionRequestTimeout(6000).setConnectTimeout(6000);
		RequestConfig config = builder.build();
		CloseableHttpResponse response = null;
		try {
			HttpPost httpPost = new HttpPost(buildUrl(url, querys));
			httpPost.setConfig(config);
			if (headers != null && !headers.isEmpty()) {
				for (Map.Entry<String, Object> e : headers.entrySet()) {
					httpPost.addHeader(e.getKey(), e.getValue().toString());
				}
			}

			if (StringUtils.isNotBlank(body)) {
				httpPost.setEntity(new StringEntity(body, "UTF-8"));
			}

			response = httpclient.execute(httpPost);
			if (null != response) {
				InputStream inputStream = response.getEntity().getContent();
				return convertStreamToString(inputStream);
			}
		} catch (Exception e) {
			logger.error("######【" + url + "】请求异常!", e);
		} finally {
			if (response != null) {
				try {
					response.close();
				} catch (IOException e) {
					e.printStackTrace();
				}
			}
		}
		return null;
	}

	public static String doPostJson(String url, Map<String, Object> headers, Map<String, Object> querys,
			String jsonStr) {
		CloseableHttpClient httpclient = HttpClients.createDefault();
		Builder builder = RequestConfig.custom().setConnectionRequestTimeout(6000).setConnectTimeout(6000);
		RequestConfig config = builder.build();
		CloseableHttpResponse response = null;
		try {
			HttpPost httpPost = new HttpPost(buildUrl(url, querys));
			httpPost.setConfig(config);
			if (headers != null && !headers.isEmpty()) {
				for (Map.Entry<String, Object> e : headers.entrySet()) {
					httpPost.addHeader(e.getKey(), e.getValue().toString());
				}
			}

			if (StringUtils.isNotBlank(jsonStr)) {
				StringEntity entity = new StringEntity(jsonStr, "UTF-8");// 解决中文乱码问题
				entity.setContentEncoding("UTF-8");
				entity.setContentType("application/json;charset=UTF-8");
				httpPost.setEntity(entity);
			}

			response = httpclient.execute(httpPost);
			logger.info("############ 请求返回值：" + response.getStatusLine().getStatusCode());
			if (null != response) {
				InputStream inputStream = response.getEntity().getContent();
				return convertStreamToString(inputStream);
			}
		} catch (Exception e) {
			logger.error("######【" + url + "】请求异常!", e);
		} finally {
			if (response != null) {
				try {
					response.close();
				} catch (IOException e) {
					e.printStackTrace();
				}
			}
		}
		return null;
	}

	/**
	 * @Author:ywzou <br>
	 * @Date: 2019年4月17日 <br>
	 * @Description: 文件下载
	 * @param url      路径
	 * @param headers  请求头参数
	 * @param querys   url参数
	 * @param bodys    表单参数
	 * @param filepath 文件存储地址
	 */
	public static void download(String url, Map<String, Object> headers, Map<String, Object> querys,
			Map<String, Object> bodys, String filepath) {
		CloseableHttpClient httpclient = HttpClients.createDefault();
		Builder builder = RequestConfig.custom().setConnectionRequestTimeout(6000).setConnectTimeout(6000);
		RequestConfig config = builder.build();
		CloseableHttpResponse response = null;

		InputStream is = null;
		FileOutputStream fileout = null;
		BufferedOutputStream bufOut = null;
		try {
			HttpPost httpPost = new HttpPost(buildUrl(url, querys));
			httpPost.setConfig(config);
			if (headers != null && !headers.isEmpty()) {
				for (Map.Entry<String, Object> e : headers.entrySet()) {
					httpPost.addHeader(e.getKey(), e.getValue().toString());
				}
			}
			if (bodys != null) {
				List<NameValuePair> nameValuePairList = new ArrayList<NameValuePair>();
				for (String key : bodys.keySet()) {
					nameValuePairList.add(new BasicNameValuePair(key, bodys.get(key).toString()));
				}
				UrlEncodedFormEntity formEntity = new UrlEncodedFormEntity(nameValuePairList, "UTF-8");
				formEntity.setContentType("application/x-www-form-urlencoded; charset=UTF-8");
				httpPost.setEntity(formEntity);
			}

			response = httpclient.execute(httpPost);
			if (null != response) {
				is = response.getEntity().getContent();
				File file = new File(filepath);
				file.getParentFile().mkdirs();
				fileout = new FileOutputStream(file);
				bufOut = new BufferedOutputStream(fileout);
				byte[] buffer = new byte[1024];
				int ch = 0;
				while ((ch = is.read(buffer)) != -1) {
					bufOut.write(buffer, 0, ch);
				}
				bufOut.flush();
			}
		} catch (Exception e) {
			logger.error("######【" + url + "】请求异常!", e);
		} finally {
			if (null != bufOut) {
				try {
					bufOut.close();
				} catch (IOException e) {
					e.printStackTrace();
				}
			}

			if (null != fileout) {
				try {
					fileout.close();
				} catch (IOException e) {
					e.printStackTrace();
				}
			}

			if (null != is) {
				try {
					is.close();
				} catch (IOException e) {
					e.printStackTrace();
				}
			}

			if (response != null) {
				try {
					response.close();
				} catch (IOException e) {
					e.printStackTrace();
				}
			}
		}
	}

	/**
	 * 作者：ywzou <br>
	 * 创建时间：2017年9月22日 <br>
	 * 描述： post请求
	 * 
	 * @param url     请求路径
	 * @param headers 请求头
	 * @param querys  请求路径参数 url参数
	 * @param bodys   请求体 表单参数
	 * @return
	 */
	public static String upload(String url, Map<String, Object> headers, Map<String, Object> querys,
			Map<String, Object> bodys, Map<String, File> files) {
		CloseableHttpClient httpclient = HttpClients.createDefault();
		CloseableHttpResponse response = null;
		try {
			HttpPost httpPost = new HttpPost(buildUrl(url, querys));
			if (headers != null && !headers.isEmpty()) {
				for (Map.Entry<String, Object> e : headers.entrySet()) {
					httpPost.addHeader(e.getKey(), e.getValue().toString());
				}
			}

			MultipartEntityBuilder reqEntity = MultipartEntityBuilder.create();
			if (files != null) {
				for (Map.Entry<String, File> en : files.entrySet()) {
					// reqEntity.addBinaryBody(en.getKey(), en.getValue());
					reqEntity.addPart(en.getKey(), new FileBody(en.getValue(), ContentType.MULTIPART_FORM_DATA));
				}
			}
			if (bodys != null) {
				for (Map.Entry<String, Object> key : bodys.entrySet()) {
					// reqEntity.addPart(key, new
					// StringBody(bodys.get(key).toString(),
					// ContentType.MULTIPART_FORM_DATA));
					// builder.addTextBody(p.getName(), p.getValue(),
					// ContentType.TEXT_PLAIN.withCharset("UTF-8"));
					reqEntity.addTextBody(key.getKey(), (String) key.getValue(),
							ContentType.TEXT_PLAIN.withCharset("UTF-8"));
				}
			}
			httpPost.setEntity(reqEntity.build());
			response = httpclient.execute(httpPost);
			if (null != response) {
				InputStream inputStream = response.getEntity().getContent();
				return convertStreamToString(inputStream);
			}
		} catch (Exception e) {
			logger.error("######【" + url + "】请求异常!", e);
		} finally {
			if (response != null) {
				try {
					response.close();
				} catch (IOException e) {
					e.printStackTrace();
				}
			}

			if (httpclient != null) {
				try {
					httpclient.close();
				} catch (IOException e) {
					e.printStackTrace();
				}
			}
		}
		return null;
	}

	/**
	 * 作者：ywzou <br>
	 * 创建时间：2017年9月22日 <br>
	 * 描述： post请求
	 * 
	 * @param url     请求路径
	 * @param headers 请求头
	 * @param querys  请求路径参数 url参数
	 * @param bodys   请求体 表单参数
	 * @return
	 */
	public static String uploadByInput(String url, Map<String, Object> headers, Map<String, Object> querys,
			Map<String, Object> bodys, Map<String, InputStream> files) {
		CloseableHttpClient httpclient = HttpClients.createDefault();
		CloseableHttpResponse response = null;
		try {
			HttpPost httpPost = new HttpPost(buildUrl(url, querys));
			if (headers != null && !headers.isEmpty()) {
				for (Map.Entry<String, Object> e : headers.entrySet()) {
					httpPost.addHeader(e.getKey(), e.getValue().toString());
				}
			}
			MultipartEntityBuilder reqEntity = MultipartEntityBuilder.create();
			if (files != null) {
				for (Map.Entry<String, InputStream> en : files.entrySet()) {
					// reqEntity.addPart(en.getKey(), new
					// InputStreamBody(en.getValue(),
					// ContentType.MULTIPART_FORM_DATA));
					reqEntity.addBinaryBody(en.getKey(), en.getValue(), ContentType.MULTIPART_FORM_DATA, en.getKey());
				}
			}

			if (bodys != null) {
				for (Entry<String, Object> key : bodys.entrySet()) {
					reqEntity.addTextBody(key.getKey(), (String) key.getValue(),
							ContentType.MULTIPART_FORM_DATA.withCharset("UTF-8"));
				}
			}
			httpPost.setEntity(reqEntity.build());
			response = httpclient.execute(httpPost);
			if (null != response) {
				InputStream inputStream = response.getEntity().getContent();
				return convertStreamToString(inputStream);
			}
		} catch (Exception e) {
			logger.error("######【" + url + "】请求异常!", e);
		} finally {
			if (response != null) {
				try {
					response.close();
				} catch (IOException e) {
					e.printStackTrace();
				}
			}

			if (httpclient != null) {
				try {
					httpclient.close();
				} catch (IOException e) {
					e.printStackTrace();
				}
			}
		}
		return null;
	}

	private static String buildUrl(String url, Map<String, Object> querys) throws UnsupportedEncodingException {
		StringBuilder sbUrl = new StringBuilder(url);
		if (null != querys) {
			StringBuilder sbQuery = new StringBuilder();
			for (Map.Entry<String, Object> query : querys.entrySet()) {
				if (0 < sbQuery.length()) {
					sbQuery.append("&");
				}
				if (StringUtils.isBlank(query.getKey()) && query.getValue() != null) {
					sbQuery.append(query.getValue());
				}
				if (!StringUtils.isBlank(query.getKey())) {
					sbQuery.append(query.getKey());
					if (query.getValue() != null) {
						sbQuery.append("=");
						sbQuery.append(URLEncoder.encode(query.getValue().toString(), "UTF-8"));
					}
				}
			}
			if (0 < sbQuery.length()) {
				sbUrl.append("?").append(sbQuery);
			}
		}

		return sbUrl.toString();
	}

	private static String convertStreamToString(InputStream is) throws UnsupportedEncodingException {
		BufferedReader reader = new BufferedReader(new InputStreamReader(is, "UTF-8"));
		StringBuilder sb = new StringBuilder();

		String line = null;
		try {
			while ((line = reader.readLine()) != null) {
				sb.append(line + "\n");
			}
		} catch (IOException e) {
			e.printStackTrace();
			logger.error("######请求结果处理流转String异常!", e);
		} finally {
			try {
				is.close();
				reader.close();
			} catch (IOException e) {
				e.printStackTrace();
			}
		}
		return sb.toString();
	}
}
