void main(List<String> args) {
  var cmd = "close_win";//close_io
  switch (cmd) {
    case "close_win":
    case "close_io":
      print("close");
      break;
      case "open":
      print("open");
      break;
    default:
      print("default");
  }
}