import '../utils.dart';

dynamic part1(List<String> input) {}

dynamic part2(List<String> input) {}

main() {
  final input = readFile('input.txt');
  final testInput = readFile('test_input.txt');

  check(0, part1(testInput));
  print(part1(input));

  check(0, part2(testInput));
  print(part2(input));
}
