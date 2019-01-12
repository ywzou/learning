///列表
void main(List<String> args) {
  var arr = [2,3.3,"dart",true];
  print(arr);//[2, 3.3, dart, true]
  print("arr的长度是：${arr.length}");//4
  print("arr是否是空：${arr.isEmpty}");//false
  print("arr[2]:${arr[2]}");//dart
  print("arr拼接为字符串:${arr.join()}");//23.3darttrue
  print("arr中是否有dart元素：${arr.contains("dart")}");//true
  //遍历
  arr.forEach((v) => print(v));

  //固定长度
  List arr2 = new List(2);
  arr2[0] = 20;
  arr2[1] = "google";
  print(arr2);//[20, google]
  //遍历
  for(int i = 0; i < arr2.length; i++){
    print(arr[1]);
  }


  //固定类型
  List<String> arr3 = new List<String>();
  arr3.add("google");
  // arr3.add(10);//不能添加
  print(arr3);//[google]

}