void main(List<String> args) {
  int i = 1;
  while(i < 10){
    int j = 1;
    while(j < 10){
      print("$i * $j = ${i * j}");
      j++;
    }
    i++;
  }

  int a = 3;
  int step = 0;
  do{
    step ++;
    print("执行$step次");//执行1次
  }while(a < 2);
}