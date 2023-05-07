void main() {
  print(findDifference("abcd", "abcde"));
  print(findDifference("abcd", "abcd"));
  print(findDifferenceDifferently("abcd", "abcdf"));
  print(findDifferenceDifferently("abcd", "abcde"));

  print(getHowManyTimes("hahahahahahahahahahahahahahahahaha", "ha"));
}

String findDifference(String a, String b) {
  String result = "";

  // create a map with key as character and value as count
  Map<String, int> mapA = {};
  Map<String, int> mapB = {};
  // loop through string a and add to map
  for (int i = 0; i < a.length; i++) {
    if (mapA.containsKey(a[i])) {
      mapA[a[i]] = mapA[a[i]]! + 1;
    } else {
      mapA[a[i]] = 1;
    }
  }
  // loop through string b and add to map
  for (int i = 0; i < b.length; i++) {
    if (mapB.containsKey(b[i])) {
      mapB[b[i]] = mapB[b[i]]! + 1;
    } else {
      mapB[b[i]] = 1;
    }
  }

  // loop through mapB and check if key is present in mapA or the count is different
  for (String key in mapB.keys) {
    if (!mapA.containsKey(key) || mapA[key] != mapB[key]) {
      return key;
    }
  }
  return result;
}

// another approach by iterating on the second string and removing characters from b string that match
String findDifferenceDifferently(String a, String b) {
  for (int i = 0; i < b.length;) {
    if (a.contains(b[i])) {
      b = b.replaceFirst(b[i], "");
      print(b);
      continue;
    }
    i++;
  }
  return b;
}

int getHowManyTimes(String a, String b) {
  int count = 0;
  int length = a.length;
  for (int i = 0; i < length; i++) {
    if (a.contains(b)) {
      a = a.replaceFirst(b, "");
      count++;
    }
    print('a');
  }
  return count;
}
