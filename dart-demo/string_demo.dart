///字符串
void main(List<String> args) {
  String str = "this is a dart String";
  print(str);

  ///使用三个单引号或者双引号也可以 创建多行字符串对象：
  str = """这是第一行
  这是第二行""";
  String str2 = '''这是第一行
  这是第二行''';


  str = "Hello";
  str2 = "Dart";
  ///字符串拼接
  str = str + str2;
  print(str);///HelloDart

  ///常用方法
  print(str.toUpperCase());//转大写 HELLODART
  print(str.toLowerCase());//转小写 hellodart
  print(str.length);///长度 9
  print(str.startsWith("He"));///以He开头 true
  print(str.endsWith("He"));///以He结尾 false
  print(str.contains("He"));///字符传中是否有He true
  print(str.replaceFirst("l", "W"));///将第一个l替换为W
  print("".isEmpty);///是否为空 true `isNotEmpty`
  print(str.substring(2));///从第3个字符开始截取 下标索引 'lloDart'
  print(str.substring(2,6));///从第3个字符开始截取到第6个字符为止 [2,6) 下标索引 'lloD'
  print(str.indexOf("l"));///第一个'l'出现尾位置 索引 2
  print(str.replaceRange(2, 6, "Google"));///将2-6位置字符替换为Google  HelloDart => HeGoogleart
  
  String str3 = "\tDart is fun\n"; 
  /// trim() trimLeft() trimRight()
  str3 = str3.trim();
  print(str3);///去除首尾的空 'Dart is fun'

  String a = "10";
  int b = int.parse(a);
  print(b);
}