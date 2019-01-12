void main(List<String> args) {
  // Person person = new Person(name:"张三");
  // person.run();

  // Employee emp = new Employee();
  // emp.run();  

  // person = new Employee(name:"李四",phone: "15198989999");
  // person.run();

  greetBob(new Person(name:"川普"));
  greetBob(new Employee(name:"李四",phone: "15198989999"));
}

greetBob(Person person) => person.run();


class Person{
  final String name;
  Person({this.name});

  void run(){
    print("Person[$name] runing...");
  }
}

///继承
class Employee implements Person{
  final String name;
  final String phone;
  Employee({
    this.name,
    this.phone
    }
  ):super();

  void run(){
    print("Employee [$name] runing...");
  }
}
