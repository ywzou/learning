void main(List<String> args) {
  int a = 1;
  var arr = [1, 2, 3];
  try {
    if(a == 1){
      throw new Exception("a can't eq 1");
    }

    if(a == 2){
      throw new NumException("自定义异常......");
    }

    print(arr[3]);
  } catch (e) {
    print(e);
  }
}

class NumException implements Exception{
  NumException([this.msg=""]);
  final String msg;
  String toString() {
    if (msg == null) return "NumException";
    return "Exception: $msg";
  }
}