void main(List<String> args) {
  for(int i = 1; i < 10; i++){
    for(int j = 1; j < 10; j++){
      print("$i * $j = ${i * j}");
    }
  }

  // continue 跳出当次循环 break结束循环
  var list = [1, "dart", 30, 12.6];
  for(int n = 0; n < list.length; n++){
    var val = list[n];
    if("dart" == val){
      continue;
    }

    if(30 == val){
      break;
    }
    print(val);
  }
}