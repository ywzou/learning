void main(List<String> args) {
    dynamic age = 20;
    dynamic name = "dart";
    dynamic height = 10.2;
    dynamic success = true;
    print(age is int);
    age = "google";
    print(age is int);
    print(name is String);
    print(height is double);
    print(success is bool);
}