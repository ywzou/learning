void main(List<String> args) {
  List<String> list = ["A","B"];

  Result result = new Result(code: "0000",msg: "请求成功!",data: list);
  print(result.toString());//[code=0000,msg:请求成功!,data:[A, B]]
}

class Result<T>{
  String code;
  String msg;
  T data;
  Result({this.code,this.msg,this.data});

  String toString(){
    return "[code=$code,msg:$msg,data:$data]";
  }
}