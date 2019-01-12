///Map
void main(List<String> args) {
  var map = {
    "name": "dart",
    "company": "google",
    "age": 8
  };

  print(map);///{name: dart, company: google, age: 8}

  Map user = new Map();
  user["name"] = "dart";
  user["age"] = 8;
  user["age"] = 10;
  print(user);///{name: dart, age: 10}

  Map<String,String> user1 = new Map<String,String>();
  user1["name"] = "dart";
  /// user1["age"] = 8;/// error A value of type 'int' can't be assigned to a variable of type 'String'.
  print(user1);///{name: dart}

  ///获取值
  print(user["name"]);///dart

  ///修改值
  user["name"] = "Flutter";
  print(user);//{name: Flutter, age: 10}

  print(user.length);//长度 
  print(user.isEmpty);//是否为空 false

  ///添加合并两个map
  Map user2 = new Map();
  user2["email"] = "dart@gmail.com";
  user2["phone"] = "911";
  user.addAll(user2);
  print(user);///{name: Flutter, age: 10, email: dart@gmail.com, phone: 911}

  ///删除一个值
  user.remove("phone");
  print(user);///{name: Flutter, age: 10, email: dart@gmail.com}

  ///遍历 传入的是一个匿名函数
  user.forEach((key,value) => 
    print("key=$key,value=$value")
  );
  //key=name,value=Flutter
  // key=age,value=10
  // key=email,value=dart@gmail.com
}
