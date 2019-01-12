void main(List<String> args) {
  Person emp = new Employee("张三", "男", 18.4, "15188889999", "xx@qq.com", 10.2);
  print(emp.getName);
  emp.speak("你好！");
}

class Person{
  String name;
  String gender;
  double height;
  ///构造方法
  ///方法一
  /// Person({this.name,this.gender="男",this.height});
  // Person(this.name,this.gender,this.height);

  Person(String name,String gender,double height){
    this.name = name;
    this.gender = gender;
    this.height = height;
  }

  //get set
  String get getName => this.name;
         set setName(String name) => this.name;


  void speak(String text){
    print("Person speak $text");
  }

}

///继承
class Employee extends Person{
  
  String phone;
  String email;
  double salary;

  Employee(
    String name,
    String gender,
    double height,
    this.phone,
    this.email,
    this.salary
  ):super(name,gender,height);

  @override
  void speak(String text){
    print("Employee speak $text");
  }
}
