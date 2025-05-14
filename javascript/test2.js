function decipher(cipher, clue) {
  const lowerClue = clue.toLowerCase();
  for (let shift = 0; shift < 26; shift++) {
    let plain = "";
    for (let i = 0; i < cipher.length; i++) {
      const ch = cipher[i];
      if (ch >= "A" && ch <= "Z") {
        plain += String.fromCharCode(
          ((ch.charCodeAt(0) - 65 - shift + 26) % 26) + 65
        );
      } else if (ch >= "a" && ch <= "z") {
        plain += String.fromCharCode(
          ((ch.charCodeAt(0) - 97 - shift + 26) % 26) + 97
        );
      } else {
        plain += ch;
      }
    }
    if (plain.toLowerCase().includes(lowerClue)) return plain;
  }
  return null;
}
