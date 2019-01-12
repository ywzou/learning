int modInverse(int a, int m) {
    bool result = false;
    a %= m;
    print("#a=$a");
    for(int x = 1; x < m; x++) {
        if((a*x) % m == 1) {
          result = true;
          return x;
        }
    }

    if(!result){
      throw new Exception("没有符合条件的数据");
    }
    return 0;
}

///数值型
void main(List<String> args) {
  int age = 100;
  print("age:$age");

  num height = 120;
  print("height:$height");

  num money =  100.00;
  print("money:$money");

  ///一些简单的运算
  int x = 10;
  int y = 3;
  print("x + y = ${x + y}");///加 13
  print("x - y = ${x - y}");///减 7
  print("x * y = ${x * y}");///乘 30
  print("x / y = ${x / y}");///除 3.333....
  print("x ~/ y = ${x ~/ y}");///除 结果取整 3
  print("x % y = ${x % y}");///余 1

  ///一些简单的方法和属性
  int z = -5;
  print("3是基数:${x.isEven}");// 是否是奇数 true
  print("10是偶数:${y.isOdd}");// 是否是偶数 true
  print("-5是负数:${z.isNegative}");// 是否为负数 5
  print("-5的绝对值:${z.abs()}");// 绝对值 5
  print("24和36的最大公约数：${24.gcd(36)}");//返回此整数和[other]的最大公约数。
  print("10的2次幂对30取模：${10.modPow(2, 30)}"); /// 10^2 % 30 = 10
  print("10大于2:${10.compareTo(2)}");/// 数值比较 1大于 0等于 -1小于
  print("10转换为double:${x.toDouble()}");/// 10.0
  print("10转换为string:${x.toString() is String}");/// is判断是否是某一个类型 结果true
  print(7.modInverse(20));///等同于modInverse()  模乘逆元 3


  double dx = 10.5;
  print(dx.ceil());///返回不小于this的最小整数。 11
  print(dx.ceilToDouble());///返回不小于this的最小整数double值。 11.0
  print(dx.floor());///返回不大于this的最小整数。 10
  print(dx.floorToDouble());///返回不大于this的最小整数double值。 10.0

  /// round() roundToDouble() 返回最接近此的整数
  print(3.5.round());// roundToDouble() 返回最接近此的整数。 4
  print(-3.5.round());//返回最接近此的整数。-4
}


