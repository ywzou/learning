void main(List<String> args) {
  List<Color> colors = Color.values;
  print(colors[0] == Color.red);

  Color aColor = Color.blue;
  switch (aColor) {
    case Color.red:
      print('Red as roses!');
      break;
    case Color.green:
      print('Green as grass!');
      break;
    default:
      print(aColor);
  }
}

enum Color {
  red,
  green,
  blue
}
