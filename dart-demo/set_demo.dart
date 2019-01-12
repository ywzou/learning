///列表
void main(List<String> args) {
  var arr = new Set();//Set没有固定的元素，所以不能指定长度
  arr.add("dart");
  arr.add(2);
  arr.add(2);

  print(arr);//{dart, 2}
  print("arr的长度是：${arr.length}");//2
  print("arr是否是空：${arr.isEmpty}");//false
  print("arr拼接为字符串:${arr.join()}");//dart2
  print("arr中是否有dart元素：${arr.contains("dart")}");//true
  print("arr中索引为1的元素：${arr.elementAt(1)}");//2
  //遍历
  arr.forEach((v) => print(v));
}

