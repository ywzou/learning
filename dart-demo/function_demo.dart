void main(List<String> args) {
  String name ="小明";
  int age = 10;
  String email = "xiaoming@qq.com";

  createUser(name, age, email);
  createUserOpt();
  createUserOptName(name: name,email: email);
  createUserPos(name, age);
  var f = (val){
    print("value : $val");
  };
  callFunction(f);
}
///参数必需
void createUser(String name,int age,String email) {
  print("name:$name,age:$age,email:$email");
}

///可选参数
void createUserOpt([String name,int age,String email]) {
  print("name:$name,age:$age,email:$email");
}

///可选命名参数  默认参数值
void createUserOptName({String name,int age = 20,String email}) {
  print("name:$name,age:$age,email:$email");
}
///可选位置参数
void createUserPos(String name,[int age,String email]) {
  print("name:$name,age:$age,email:$email");
}

///方法作为参数
void callFunction(void f(int age)){
  f(100);//执行传入的方法
}